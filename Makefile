generate:
	protoc -I. ./proto/web.proto \
		--go_out=plugins=grpc:$$GOPATH/src
	go1.11beta3 generate -x ./frontend/

serve:
	go1.11beta3 run main.go
