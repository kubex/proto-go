// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

/*
Package users is a generated protocol buffer package.

It is generated from these files:
	users.proto

It has these top-level messages:
	RetrieveRequest
	UserResponse
	UsersResponse
*/
package users

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

type RetrieveRequest struct {
	Ids    []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Emails []string `protobuf:"bytes,2,rep,name=emails" json:"emails,omitempty"`
}

func (m *RetrieveRequest) Reset()                    { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()               {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RetrieveRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *RetrieveRequest) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

type UserResponse struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type UsersResponse struct {
	Users map[string]*UserResponse `protobuf:"bytes,1,rep,name=users" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UsersResponse) Reset()                    { *m = UsersResponse{} }
func (m *UsersResponse) String() string            { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()               {}
func (*UsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UsersResponse) GetUsers() map[string]*UserResponse {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*RetrieveRequest)(nil), "users.RetrieveRequest")
	proto.RegisterType((*UserResponse)(nil), "users.UserResponse")
	proto.RegisterType((*UsersResponse)(nil), "users.UsersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Users service

type UsersClient interface {
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*UsersResponse, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := grpc.Invoke(ctx, "/users.Users/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersServer interface {
	Retrieve(context.Context, *RetrieveRequest) (*UsersResponse, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Retrieve",
			Handler:    _Users_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4b, 0x4b, 0xfb, 0x40,
	0x10, 0x27, 0xe9, 0x3f, 0xa5, 0x99, 0xfe, 0x7d, 0xb0, 0x96, 0x12, 0x2a, 0xd2, 0xda, 0x53, 0x3d,
	0x98, 0x40, 0x44, 0x28, 0x7a, 0x52, 0xf0, 0xa8, 0x87, 0x05, 0x2f, 0x5e, 0x24, 0x31, 0x63, 0x5d,
	0x9a, 0x97, 0xbb, 0x9b, 0x62, 0x3f, 0x88, 0xdf, 0x57, 0x76, 0x36, 0xf1, 0x85, 0xb7, 0xfd, 0x3d,
	0x66, 0xe6, 0x37, 0xb3, 0x30, 0x6c, 0x14, 0x4a, 0x15, 0xd6, 0xb2, 0xd2, 0x15, 0xf3, 0x08, 0xcc,
	0x2f, 0x61, 0x8f, 0xa3, 0x96, 0x02, 0x37, 0xc8, 0xf1, 0xb5, 0x41, 0xa5, 0xd9, 0x3e, 0xf4, 0x44,
	0xa6, 0x02, 0x67, 0xd6, 0x5b, 0xf8, 0xdc, 0x3c, 0xd9, 0x18, 0xfa, 0x58, 0x24, 0x22, 0x57, 0x81,
	0x4b, 0x64, 0x8b, 0xe6, 0x35, 0xfc, 0xbf, 0x57, 0x28, 0x39, 0xaa, 0xba, 0x2a, 0x15, 0xb2, 0x5d,
	0x70, 0x45, 0x16, 0x38, 0x33, 0x67, 0xe1, 0x73, 0x57, 0x64, 0x6c, 0x04, 0x1e, 0x39, 0x03, 0x97,
	0x28, 0x0b, 0xd8, 0x11, 0xc0, 0xb3, 0x90, 0x4a, 0x3f, 0x96, 0x49, 0x81, 0x41, 0x8f, 0x24, 0x9f,
	0x98, 0xbb, 0xa4, 0x40, 0x76, 0x08, 0x7e, 0x9e, 0x74, 0xea, 0x3f, 0x52, 0x07, 0x86, 0x30, 0xe2,
	0xfc, 0xdd, 0x81, 0x1d, 0x33, 0x52, 0x7d, 0xce, 0x3c, 0x07, 0xbb, 0x09, 0xe5, 0x1d, 0xc6, 0xd3,
	0xd0, 0x2e, 0xf9, 0xc3, 0x64, 0xd1, 0x4d, 0xa9, 0xe5, 0x96, 0x5b, 0xf7, 0xe4, 0x16, 0xe0, 0x8b,
	0x34, 0x2b, 0xaf, 0x71, 0xdb, 0x26, 0x37, 0x4f, 0x76, 0x02, 0xde, 0x26, 0xc9, 0x1b, 0xa4, 0xe8,
	0xc3, 0xf8, 0xe0, 0x5b, 0xdb, 0xae, 0x2b, 0xb7, 0x8e, 0x0b, 0x77, 0xe9, 0xc4, 0x57, 0xe0, 0x51,
	0x3b, 0xb6, 0x84, 0x41, 0x77, 0x4f, 0x36, 0x6e, 0x8b, 0x7e, 0x1d, 0x78, 0x32, 0xfa, 0x2b, 0xe3,
	0xf5, 0xf1, 0xc3, 0x74, 0x25, 0xf4, 0x4b, 0x93, 0x86, 0x4f, 0x55, 0x11, 0xad, 0x9b, 0x14, 0xdf,
	0x22, 0xfa, 0xaa, 0xd3, 0x55, 0x15, 0x51, 0x41, 0xda, 0x27, 0x7c, 0xf6, 0x11, 0x00, 0x00, 0xff,
	0xff, 0xb0, 0xb4, 0xc9, 0xb0, 0xc9, 0x01, 0x00, 0x00,
}
