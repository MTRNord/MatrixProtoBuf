// Code generated by protoc-gen-go. DO NOT EDIT.
// source: matrix.proto

/*
Package matrixProtos is a generated protocol buffer package.

It is generated from these files:
	matrix.proto

It has these top-level messages:
	VersionsResponse
	Version
	LoginRequest
	LoginResponse
*/
package matrixProtos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
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

type VersionsResponse struct {
	Versions []*Version `protobuf:"bytes,1,rep,name=versions" json:"versions,omitempty"`
}

func (m *VersionsResponse) Reset()                    { *m = VersionsResponse{} }
func (m *VersionsResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionsResponse) ProtoMessage()               {}
func (*VersionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VersionsResponse) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

type Version struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type LoginRequest struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Medium   string `protobuf:"bytes,4,opt,name=medium" json:"medium,omitempty"`
	Address  string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LoginRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetMedium() string {
	if m != nil {
		return m.Medium
	}
	return ""
}

func (m *LoginRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type LoginResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	HomeServer   string `protobuf:"bytes,2,opt,name=home_server,json=homeServer" json:"home_server,omitempty"`
	UserId       string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	Errcode      string `protobuf:"bytes,5,opt,name=errcode" json:"errcode,omitempty"`
	Error        string `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginResponse) GetHomeServer() string {
	if m != nil {
		return m.HomeServer
	}
	return ""
}

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *LoginResponse) GetErrcode() string {
	if m != nil {
		return m.Errcode
	}
	return ""
}

func (m *LoginResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionsResponse)(nil), "matrixProtos.VersionsResponse")
	proto.RegisterType((*Version)(nil), "matrixProtos.Version")
	proto.RegisterType((*LoginRequest)(nil), "matrixProtos.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "matrixProtos.LoginResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoginService service

type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/matrixProtos.LoginService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoginService service

type LoginServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrixProtos.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matrixProtos.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "matrix.proto",
}

func init() { proto.RegisterFile("matrix.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x80, 0x53, 0x81, 0x16, 0x87, 0x92, 0x98, 0x8d, 0x3f, 0x1b, 0x3c, 0x88, 0xe5, 0xc2, 0x89,
	0x44, 0x7c, 0x01, 0x2f, 0x1e, 0x4c, 0x3c, 0x90, 0x6a, 0xbc, 0x92, 0xda, 0x8e, 0xd2, 0x98, 0xb2,
	0x75, 0xa6, 0xa0, 0xde, 0x7d, 0x31, 0xdf, 0xcc, 0xec, 0xee, 0x94, 0x60, 0xe2, 0x6d, 0xbf, 0x6f,
	0x7f, 0xfa, 0x75, 0x5b, 0x88, 0xab, 0xac, 0xa1, 0xf2, 0x73, 0x56, 0x93, 0x69, 0x8c, 0x12, 0x5a,
	0x58, 0xe0, 0xe4, 0x16, 0x8e, 0x9e, 0x90, 0xb8, 0x34, 0x6b, 0x4e, 0x91, 0x6b, 0xb3, 0x66, 0x54,
	0x57, 0xd0, 0xdf, 0x8a, 0xd3, 0xc1, 0xb8, 0x33, 0x1d, 0xcc, 0x4f, 0x66, 0xfb, 0x9b, 0x66, 0xb2,
	0x23, 0xdd, 0x2d, 0x4b, 0x26, 0x10, 0x89, 0x54, 0x1a, 0x22, 0xd1, 0x3a, 0x18, 0x07, 0xd3, 0xc3,
	0xb4, 0xc5, 0xe4, 0x3b, 0x80, 0xf8, 0xde, 0xbc, 0x96, 0xeb, 0x14, 0xdf, 0x37, 0xc8, 0x8d, 0x52,
	0xd0, 0x6d, 0xbe, 0x6a, 0x94, 0x75, 0x6e, 0x6c, 0xdd, 0x86, 0x91, 0xf4, 0x81, 0x77, 0x76, 0xac,
	0x46, 0xd0, 0xaf, 0x33, 0xe6, 0x0f, 0x43, 0x85, 0xee, 0x38, 0xbf, 0x63, 0x75, 0x0a, 0x61, 0x85,
	0x45, 0xb9, 0xa9, 0x74, 0xd7, 0xcd, 0x08, 0xd9, 0x8c, 0xac, 0x28, 0x08, 0x99, 0x75, 0xcf, 0x67,
	0x08, 0x26, 0x3f, 0x01, 0x0c, 0x25, 0x43, 0x5e, 0xf8, 0x12, 0xe2, 0x2c, 0xcf, 0x91, 0x79, 0xd9,
	0x98, 0x37, 0x6c, 0xbb, 0x07, 0xde, 0x3d, 0x5a, 0xa5, 0x2e, 0x60, 0xb0, 0x32, 0x15, 0x2e, 0x19,
	0x69, 0xbb, 0xab, 0x03, 0xab, 0x1e, 0x9c, 0x51, 0x67, 0x10, 0xd9, 0xd6, 0x65, 0xd9, 0x26, 0x86,
	0x16, 0xef, 0x0a, 0x35, 0x81, 0x21, 0xe1, 0x0b, 0x21, 0xaf, 0xe4, 0x74, 0xdf, 0x19, 0x8b, 0xf4,
	0xc7, 0x6b, 0x88, 0x90, 0x28, 0x37, 0x05, 0xb6, 0xb5, 0x82, 0xea, 0x18, 0x7a, 0x48, 0x64, 0x48,
	0x87, 0xce, 0x7b, 0x98, 0x2f, 0xe4, 0x26, 0xed, 0xc3, 0xcb, 0x1c, 0xd5, 0x0d, 0xf4, 0x1c, 0xab,
	0xd1, 0xdf, 0x2f, 0xb5, 0x7f, 0xdd, 0xa3, 0xf3, 0x7f, 0xe7, 0xfc, 0x1d, 0x3c, 0x87, 0xee, 0xef,
	0xb8, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x37, 0x21, 0x76, 0xf7, 0x2d, 0x02, 0x00, 0x00,
}
