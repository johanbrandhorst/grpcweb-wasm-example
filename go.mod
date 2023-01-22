module github.com/johanbrandhorst/grpcweb-wasm-example

go 1.19

require (
	github.com/golang/protobuf v1.3.0
	github.com/improbable-eng/grpc-web v0.9.1
	github.com/shurcooL/vfsgen v0.0.0-20180414153424-bded413e56ef
	github.com/sirupsen/logrus v1.0.5
	google.golang.org/genproto v0.0.0-20190227213309-4f5b463f9597
	google.golang.org/grpc v1.19.0
)

require (
	github.com/golang/glog v1.0.0 // indirect
	github.com/gorilla/websocket v1.2.0 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/rs/cors v1.4.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20171119174359-809beceb2371 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20180621125126-a49355c7e3f8 // indirect
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd // indirect
	golang.org/x/sys v0.0.0-20180707002001-3c6ecd8f22c6 // indirect
	golang.org/x/text v0.3.0 // indirect
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
)

replace google.golang.org/grpc => github.com/johanbrandhorst/grpc-go v1.2.1-0.20180625151142-1f109e898476
