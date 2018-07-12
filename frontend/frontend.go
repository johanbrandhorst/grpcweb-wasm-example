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

// Build with Go WASM fork
//go:generate rm -f ./html/*
//go:generate bash -c "GOOS=js GOARCH=wasm go1.11beta1 build -o ./html/test.wasm frontend.go"

//go:generate bash -c "cp $DOLLAR(go1.11beta1 env GOROOT)/misc/wasm/wasm_exec.html ./html/index.html"
//go:generate bash -c "cp $DOLLAR(go1.11beta1 env GOROOT)/misc/wasm/wasm_exec.js ./html/wasm_exec.js"
//go:generate bash -c "sed -i -e 's;</button>;</button>\\n\\t<div id=\"target\"></div>;' ./html/index.html"

// Integrate generated JS into a Go file for static loading.
//go:generate bash -c "go1.11beta1 run assets_generate.go"

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
