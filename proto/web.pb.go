// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/web.proto

/*
Package web is a generated protocol buffer package.

Web exposes a backend server over gRPC.

It is generated from these files:
	proto/web.proto

It has these top-level messages:
	GetUserRequest
	User
	GetUsersRequest
*/
package web

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetUserRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type User struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUsersRequest struct {
	NumUsers int64 `protobuf:"varint,1,opt,name=num_users,json=numUsers" json:"num_users,omitempty"`
}

func (m *GetUsersRequest) Reset()                    { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()               {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUsersRequest) GetNumUsers() int64 {
	if m != nil {
		return m.NumUsers
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "web.GetUserRequest")
	proto.RegisterType((*User)(nil), "web.User")
	proto.RegisterType((*GetUsersRequest)(nil), "web.GetUsersRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Backend service

type BackendClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (Backend_GetUsersClient, error)
}

type backendClient struct {
	cc *grpc.ClientConn
}

func NewBackendClient(cc *grpc.ClientConn) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/web.Backend/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (Backend_GetUsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Backend_serviceDesc.Streams[0], c.cc, "/web.Backend/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &backendGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backend_GetUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type backendGetUsersClient struct {
	grpc.ClientStream
}

func (x *backendGetUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Backend service

type BackendServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	GetUsers(*GetUsersRequest, Backend_GetUsersServer) error
}

func RegisterBackendServer(s *grpc.Server, srv BackendServer) {
	s.RegisterService(&_Backend_serviceDesc, srv)
}

func _Backend_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Backend/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServer).GetUsers(m, &backendGetUsersServer{stream})
}

type Backend_GetUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type backendGetUsersServer struct {
	grpc.ServerStream
}

func (x *backendGetUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

var _Backend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "web.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Backend_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _Backend_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/web.proto",
}

func init() { proto.RegisterFile("proto/web.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4b, 0x03, 0x31,
	0x1c, 0xc5, 0xbd, 0x56, 0x7a, 0xbd, 0xef, 0xd0, 0x42, 0x14, 0x15, 0x5d, 0xe4, 0x26, 0x45, 0x9a,
	0x88, 0x4e, 0xd2, 0xad, 0x8b, 0xb8, 0x1e, 0xb8, 0xb8, 0x94, 0xe4, 0xf2, 0xe5, 0xee, 0xd4, 0x24,
	0x67, 0x7e, 0x10, 0xff, 0x7c, 0x49, 0x68, 0x45, 0xa1, 0xdb, 0xcb, 0xcb, 0xe7, 0xbd, 0x3c, 0x02,
	0xcb, 0xd1, 0x1a, 0x6f, 0x58, 0x44, 0x41, 0xb3, 0x22, 0xd3, 0x88, 0xa2, 0xbe, 0x85, 0xc5, 0x33,
	0xfa, 0x57, 0x87, 0xb6, 0xc1, 0xaf, 0x80, 0xce, 0x93, 0x73, 0x28, 0x83, 0x43, 0xbb, 0x1d, 0xe4,
	0x45, 0x71, 0x5d, 0xdc, 0x54, 0xcd, 0x2c, 0x1d, 0x5f, 0x64, 0x7d, 0x06, 0xc7, 0x89, 0x23, 0x0b,
	0x98, 0xfc, 0xde, 0x4d, 0x06, 0x59, 0x53, 0x58, 0xee, 0x2a, 0xdc, 0xbe, 0xe3, 0x0a, 0x2a, 0x1d,
	0xd4, 0x36, 0x05, 0x5d, 0x26, 0xa7, 0xcd, 0x5c, 0x07, 0x95, 0x99, 0x87, 0x0e, 0xca, 0x0d, 0x6f,
	0x3f, 0x50, 0x4b, 0x72, 0x07, 0xe5, 0x2e, 0x4a, 0x4e, 0x68, 0x5a, 0xf6, 0x7f, 0xcb, 0x65, 0x95,
	0xcd, 0xe4, 0xd4, 0x47, 0x84, 0xc1, 0x7c, 0xff, 0x0e, 0x39, 0xfd, 0x4b, 0xbb, 0x43, 0xf8, 0x7d,
	0xb1, 0x59, 0xbf, 0x3d, 0x75, 0x83, 0xef, 0x83, 0xa0, 0xad, 0x51, 0xec, 0xdd, 0xf4, 0x5c, 0x0b,
	0xcb, 0xb5, 0xec, 0x8d, 0x75, 0x9e, 0x75, 0x76, 0x6c, 0x23, 0x8a, 0x55, 0xe4, 0x4e, 0xad, 0xf0,
	0x9b, 0xab, 0xf1, 0x13, 0x59, 0xfe, 0x99, 0x75, 0x44, 0x21, 0x66, 0x59, 0x3e, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe7, 0x50, 0x1b, 0x08, 0x37, 0x01, 0x00, 0x00,
}