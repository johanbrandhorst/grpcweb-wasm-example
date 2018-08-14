# gRPC-Web WASM example
Small example server using gRPC-Web with WASM using `go1.11rc1`

## Requirements

1. `go1.11rc1`. See https://golang.org/dl/#go1.11rc1 for installation
instructions.
1. The Google protobuf compiler, `protoc`.

## Development

Run `make generate` to regenerate the protofiles and the frontend.

Run `make serve` to start the web server.
