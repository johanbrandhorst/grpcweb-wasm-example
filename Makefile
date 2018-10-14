generate:
	protoc -I. ./proto/web.proto \
		--go_out=plugins=grpc:/home/johan/go/src
	go generate -x ./frontend/

serve:
	go run main.go
