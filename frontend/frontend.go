//go:build js && wasm

package main

import (
	"context"
	"io"
	"io/ioutil"
	"syscall/js"

	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	web "github.com/johanbrandhorst/grpcweb-wasm-example/proto"
)

var document js.Value

type DivWriter js.Value

func (d DivWriter) Write(p []byte) (n int, err error) {
	node := document.Call("createElement", "div")
	node.Set("innerHTML", string(p))
	js.Value(d).Call("appendChild", node)
	return len(p), nil
}

func init() {
	document = js.Global().Get("document")
	div := document.Call("getElementById", "target")
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(DivWriter(div), ioutil.Discard, ioutil.Discard))
}

func main() {
	// Dial the server which is hosting our wasm code.
	// This is equivalent to fetch("") in the browser.
	cc, err := grpc.Dial("")
	if err != nil {
		grpclog.Println(err)
		return
	}
	client := web.NewBackendClient(cc)
	resp, err := client.GetUser(context.Background(), &web.GetUserRequest{
		UserId: "1234",
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
	} else {
		grpclog.Println(resp)
	}
	resp, err = client.GetUser(context.Background(), &web.GetUserRequest{
		UserId: "123",
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
	} else {
		grpclog.Println(resp)
	}

	srv, err := client.GetUsers(context.Background(), &web.GetUsersRequest{
		NumUsers: 3,
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
	} else {
		for {
			user, err := srv.Recv()
			if err != nil {
				if err != io.EOF {
					st := status.Convert(err)
					grpclog.Println(st.Code(), st.Message(), st.Details())
				}
				break
			}

			grpclog.Println(user)
		}
	}

	grpclog.Println("finished")
}
