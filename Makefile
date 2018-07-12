generate:
	protoc -I. ./proto/web.proto \
		--go_out=plugins=grpc:$$GOPATH/src
	go1.11beta1 generate -x ./frontend/

serve:
	go1.11beta1 run main.go
