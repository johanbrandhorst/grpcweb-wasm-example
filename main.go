// Copyright 2018 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"crypto/tls"
	"net/http"
	"strings"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/johanbrandhorst/grpcweb-wasm-example/backend"
	"github.com/johanbrandhorst/grpcweb-wasm-example/frontend/html"
	web "github.com/johanbrandhorst/grpcweb-wasm-example/proto"
)

var logger *logrus.Logger

func init() {
	logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})
	// Should only be done from init functions
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))
}

func main() {
	gs := grpc.NewServer()
	web.RegisterBackendServer(gs, &backend.Backend{})
	wrappedServer := grpcweb.WrapServer(gs)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		// Redirect gRPC and gRPC-Web requests to the gRPC-Web Websocket Proxy server
		if req.ProtoMajor == 2 && strings.Contains(req.Header.Get("Content-Type"), "application/grpc") {
			wrappedServer.ServeHTTP(resp, req)
			return
		}

		// Serve the WASM client
		wasmContentTypeSetter(http.FileServer(http.FS(html.Assets))).ServeHTTP(resp, req)
	}

	addr := "localhost:10000"
	httpsSrv := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(handler),
		// Some security settings
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
		},
	}

	logger.Info("Serving on https://" + addr)
	logger.Fatal(httpsSrv.ListenAndServeTLS("./insecure/cert.pem", "./insecure/key.pem"))
}

func wasmContentTypeSetter(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.URL.Path, ".wasm") {
			w.Header().Set("content-type", "application/wasm")
		}
		fn.ServeHTTP(w, req)
	}
}
