// Code generated by protoc-gen-go.
// source: notify.proto
// DO NOT EDIT!

/*
Package notify is a generated protocol buffer package.

It is generated from these files:
	notify.proto

It has these top-level messages:
	NotificationRequest
	NotificationDataRequest
	NotificationResponse
	NotificationUpdateResponse
	NotificationReadRequest
	NotificationReadAllRequest
	NotificationReadAllResponse
	NotificationListRequest
	NotificationsResponse
*/
package notify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type NotificationRequest struct {
	NotificationUid string `protobuf:"bytes,1,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
}

func (m *NotificationRequest) Reset()                    { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()               {}
func (*NotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NotificationRequest) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

type NotificationDataRequest struct {
	NotificationId  string            `protobuf:"bytes,1,opt,name=notification_id,json=notificationId" json:"notification_id,omitempty"`
	NotificationUid string            `protobuf:"bytes,2,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
	ProjectId       string            `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	TargetId        []string          `protobuf:"bytes,4,rep,name=target_id,json=targetId" json:"target_id,omitempty"`
	Percentage      int32             `protobuf:"varint,5,opt,name=percentage" json:"percentage,omitempty"`
	Data            map[string]string `protobuf:"bytes,6,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NotificationDataRequest) Reset()                    { *m = NotificationDataRequest{} }
func (m *NotificationDataRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationDataRequest) ProtoMessage()               {}
func (*NotificationDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NotificationDataRequest) GetNotificationId() string {
	if m != nil {
		return m.NotificationId
	}
	return ""
}

func (m *NotificationDataRequest) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *NotificationDataRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *NotificationDataRequest) GetTargetId() []string {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *NotificationDataRequest) GetPercentage() int32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *NotificationDataRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type NotificationResponse struct {
	NotificationUid string                     `protobuf:"bytes,1,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
	NotificationId  string                     `protobuf:"bytes,2,opt,name=notification_id,json=notificationId" json:"notification_id,omitempty"`
	ProjectId       string                     `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Percentage      int32                      `protobuf:"varint,4,opt,name=percentage" json:"percentage,omitempty"`
	IsComplete      bool                       `protobuf:"varint,5,opt,name=is_complete,json=isComplete" json:"is_complete,omitempty"`
	Started         *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=started" json:"started,omitempty"`
	LastUpdate      *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=last_update,json=lastUpdate" json:"last_update,omitempty"`
	Read            bool                       `protobuf:"varint,8,opt,name=read" json:"read,omitempty"`
	Data            map[string]string          `protobuf:"bytes,9,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NotificationResponse) Reset()                    { *m = NotificationResponse{} }
func (m *NotificationResponse) String() string            { return proto.CompactTextString(m) }
func (*NotificationResponse) ProtoMessage()               {}
func (*NotificationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NotificationResponse) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *NotificationResponse) GetNotificationId() string {
	if m != nil {
		return m.NotificationId
	}
	return ""
}

func (m *NotificationResponse) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *NotificationResponse) GetPercentage() int32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *NotificationResponse) GetIsComplete() bool {
	if m != nil {
		return m.IsComplete
	}
	return false
}

func (m *NotificationResponse) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *NotificationResponse) GetLastUpdate() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

func (m *NotificationResponse) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *NotificationResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type NotificationUpdateResponse struct {
	NotificationUid string `protobuf:"bytes,1,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
	Message         string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Success         bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
}

func (m *NotificationUpdateResponse) Reset()                    { *m = NotificationUpdateResponse{} }
func (m *NotificationUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*NotificationUpdateResponse) ProtoMessage()               {}
func (*NotificationUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NotificationUpdateResponse) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *NotificationUpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NotificationUpdateResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type NotificationReadRequest struct {
	NotificationUid string `protobuf:"bytes,1,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
	MemberId        string `protobuf:"bytes,2,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Read            bool   `protobuf:"varint,3,opt,name=read" json:"read,omitempty"`
}

func (m *NotificationReadRequest) Reset()                    { *m = NotificationReadRequest{} }
func (m *NotificationReadRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationReadRequest) ProtoMessage()               {}
func (*NotificationReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NotificationReadRequest) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *NotificationReadRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

func (m *NotificationReadRequest) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

type NotificationReadAllRequest struct {
	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
}

func (m *NotificationReadAllRequest) Reset()                    { *m = NotificationReadAllRequest{} }
func (m *NotificationReadAllRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationReadAllRequest) ProtoMessage()               {}
func (*NotificationReadAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NotificationReadAllRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

type NotificationReadAllResponse struct {
	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Success  bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *NotificationReadAllResponse) Reset()                    { *m = NotificationReadAllResponse{} }
func (m *NotificationReadAllResponse) String() string            { return proto.CompactTextString(m) }
func (*NotificationReadAllResponse) ProtoMessage()               {}
func (*NotificationReadAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NotificationReadAllResponse) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

func (m *NotificationReadAllResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type NotificationListRequest struct {
	MemberId  string `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Limit     int32  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *NotificationListRequest) Reset()                    { *m = NotificationListRequest{} }
func (m *NotificationListRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationListRequest) ProtoMessage()               {}
func (*NotificationListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NotificationListRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

func (m *NotificationListRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *NotificationListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type NotificationsResponse struct {
	Notifications []*NotificationResponse `protobuf:"bytes,1,rep,name=notifications" json:"notifications,omitempty"`
}

func (m *NotificationsResponse) Reset()                    { *m = NotificationsResponse{} }
func (m *NotificationsResponse) String() string            { return proto.CompactTextString(m) }
func (*NotificationsResponse) ProtoMessage()               {}
func (*NotificationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *NotificationsResponse) GetNotifications() []*NotificationResponse {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func init() {
	proto.RegisterType((*NotificationRequest)(nil), "notify.NotificationRequest")
	proto.RegisterType((*NotificationDataRequest)(nil), "notify.NotificationDataRequest")
	proto.RegisterType((*NotificationResponse)(nil), "notify.NotificationResponse")
	proto.RegisterType((*NotificationUpdateResponse)(nil), "notify.NotificationUpdateResponse")
	proto.RegisterType((*NotificationReadRequest)(nil), "notify.NotificationReadRequest")
	proto.RegisterType((*NotificationReadAllRequest)(nil), "notify.NotificationReadAllRequest")
	proto.RegisterType((*NotificationReadAllResponse)(nil), "notify.NotificationReadAllResponse")
	proto.RegisterType((*NotificationListRequest)(nil), "notify.NotificationListRequest")
	proto.RegisterType((*NotificationsResponse)(nil), "notify.NotificationsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Notify service

type NotifyClient interface {
	Send(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error)
	Update(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error)
	Retrieve(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error)
	List(ctx context.Context, in *NotificationListRequest, opts ...grpc.CallOption) (*NotificationsResponse, error)
	Read(ctx context.Context, in *NotificationReadRequest, opts ...grpc.CallOption) (*NotificationResponse, error)
	ReadAll(ctx context.Context, in *NotificationReadAllRequest, opts ...grpc.CallOption) (*NotificationReadAllResponse, error)
}

type notifyClient struct {
	cc *grpc.ClientConn
}

func NewNotifyClient(cc *grpc.ClientConn) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) Send(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error) {
	out := new(NotificationUpdateResponse)
	err := grpc.Invoke(ctx, "/notify.Notify/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) Update(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error) {
	out := new(NotificationUpdateResponse)
	err := grpc.Invoke(ctx, "/notify.Notify/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) Retrieve(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := grpc.Invoke(ctx, "/notify.Notify/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) List(ctx context.Context, in *NotificationListRequest, opts ...grpc.CallOption) (*NotificationsResponse, error) {
	out := new(NotificationsResponse)
	err := grpc.Invoke(ctx, "/notify.Notify/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) Read(ctx context.Context, in *NotificationReadRequest, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := grpc.Invoke(ctx, "/notify.Notify/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) ReadAll(ctx context.Context, in *NotificationReadAllRequest, opts ...grpc.CallOption) (*NotificationReadAllResponse, error) {
	out := new(NotificationReadAllResponse)
	err := grpc.Invoke(ctx, "/notify.Notify/ReadAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notify service

type NotifyServer interface {
	Send(context.Context, *NotificationDataRequest) (*NotificationUpdateResponse, error)
	Update(context.Context, *NotificationDataRequest) (*NotificationUpdateResponse, error)
	Retrieve(context.Context, *NotificationRequest) (*NotificationResponse, error)
	List(context.Context, *NotificationListRequest) (*NotificationsResponse, error)
	Read(context.Context, *NotificationReadRequest) (*NotificationResponse, error)
	ReadAll(context.Context, *NotificationReadAllRequest) (*NotificationReadAllResponse, error)
}

func RegisterNotifyServer(s *grpc.Server, srv NotifyServer) {
	s.RegisterService(&_Notify_serviceDesc, srv)
}

func _Notify_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notify.Notify/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).Send(ctx, req.(*NotificationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notify.Notify/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).Update(ctx, req.(*NotificationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notify.Notify/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).Retrieve(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notify.Notify/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).List(ctx, req.(*NotificationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notify.Notify/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).Read(ctx, req.(*NotificationReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notify.Notify/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).ReadAll(ctx, req.(*NotificationReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notify.Notify",
	HandlerType: (*NotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Notify_Send_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Notify_Update_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Notify_Retrieve_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Notify_List_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Notify_Read_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _Notify_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notify.proto",
}

func init() { proto.RegisterFile("notify.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x56, 0x62, 0xe7, 0x35, 0xe1, 0x51, 0x2d, 0x45, 0x58, 0x2e, 0xa5, 0x91, 0x91, 0x20, 0x3d,
	0xe0, 0x48, 0x05, 0x09, 0x28, 0x42, 0x82, 0x42, 0x91, 0x22, 0x1e, 0x07, 0xd3, 0x5e, 0xe0, 0x10,
	0x6d, 0xec, 0x69, 0x58, 0xea, 0x17, 0xde, 0x75, 0x45, 0xc4, 0xaf, 0xe3, 0xc2, 0x91, 0xdf, 0x84,
	0xbc, 0x6b, 0xa7, 0x76, 0x71, 0x13, 0x8a, 0x7a, 0xdb, 0x1d, 0x7f, 0xfb, 0xcd, 0xcc, 0xf7, 0xcd,
	0x24, 0x70, 0x25, 0x8c, 0x04, 0x3b, 0x9a, 0xdb, 0x71, 0x12, 0x89, 0x88, 0xb4, 0xd5, 0xcd, 0xdc,
	0x9a, 0x45, 0xd1, 0xcc, 0xc7, 0x91, 0x8c, 0x4e, 0xd3, 0xa3, 0x91, 0x60, 0x01, 0x72, 0x41, 0x83,
	0x58, 0x01, 0xad, 0x17, 0x70, 0xe3, 0x43, 0x06, 0x65, 0x2e, 0x15, 0x2c, 0x0a, 0x1d, 0xfc, 0x96,
	0x22, 0x17, 0x64, 0x1b, 0xd6, 0xc2, 0x52, 0x78, 0x92, 0x32, 0xcf, 0x68, 0x0c, 0x1a, 0xc3, 0x9e,
	0x73, 0xbd, 0x1c, 0x3f, 0x64, 0x9e, 0xf5, 0xb3, 0x09, 0xb7, 0xca, 0x14, 0xaf, 0xa9, 0xa0, 0x05,
	0xcd, 0x7d, 0xa8, 0xc0, 0x27, 0x0b, 0x96, 0x6b, 0xe5, 0xf0, 0xd8, 0xab, 0xcd, 0xd7, 0xac, 0xcd,
	0x47, 0x36, 0x01, 0xe2, 0x24, 0xfa, 0x8a, 0xae, 0xc8, 0xe8, 0x34, 0x09, 0xea, 0xe5, 0x91, 0xb1,
	0x47, 0x36, 0xa0, 0x27, 0x68, 0x32, 0x43, 0xf9, 0x55, 0x1f, 0x68, 0xc3, 0x9e, 0xd3, 0x55, 0x81,
	0xb1, 0x47, 0xee, 0x00, 0xc4, 0x98, 0xb8, 0x18, 0x0a, 0x3a, 0x43, 0xa3, 0x35, 0x68, 0x0c, 0x5b,
	0x4e, 0x29, 0x42, 0x9e, 0x83, 0xee, 0x51, 0x41, 0x8d, 0xf6, 0x40, 0x1b, 0xf6, 0x77, 0xb6, 0xed,
	0x5c, 0xd3, 0x73, 0xda, 0xb3, 0xb3, 0xf3, 0x7e, 0x28, 0x92, 0xb9, 0x23, 0x9f, 0x99, 0x8f, 0xa1,
	0xb7, 0x08, 0x91, 0x35, 0xd0, 0x8e, 0x71, 0x9e, 0xf7, 0x9b, 0x1d, 0xc9, 0x3a, 0xb4, 0x4e, 0xa8,
	0x9f, 0x62, 0xde, 0x99, 0xba, 0xec, 0x36, 0x9f, 0x34, 0xac, 0x5f, 0x1a, 0xac, 0x57, 0x6d, 0xe0,
	0x71, 0x14, 0x72, 0xbc, 0x80, 0x0f, 0x75, 0x5a, 0x37, 0x6b, 0xb5, 0x5e, 0x21, 0x60, 0x55, 0x23,
	0xfd, 0x2f, 0x8d, 0xb6, 0xa0, 0xcf, 0xf8, 0xc4, 0x8d, 0x82, 0xd8, 0x47, 0xa1, 0x44, 0xec, 0x3a,
	0xc0, 0xf8, 0xab, 0x3c, 0x42, 0x1e, 0x41, 0x87, 0x0b, 0x9a, 0x08, 0xf4, 0x8c, 0xf6, 0xa0, 0x31,
	0xec, 0xef, 0x98, 0xb6, 0x9a, 0x42, 0xbb, 0x98, 0x42, 0xfb, 0xa0, 0x98, 0x42, 0xa7, 0x80, 0x92,
	0x67, 0xd0, 0xf7, 0x29, 0x17, 0x93, 0x34, 0xf6, 0xa8, 0x40, 0xa3, 0xb3, 0xf2, 0x25, 0x64, 0xf0,
	0x43, 0x89, 0x26, 0x04, 0xf4, 0x04, 0xa9, 0x67, 0x74, 0x65, 0x31, 0xf2, 0x4c, 0x76, 0x73, 0x2f,
	0x7b, 0xd2, 0xcb, 0x7b, 0x75, 0x5e, 0x16, 0x32, 0x5f, 0x9e, 0x91, 0x3f, 0xc0, 0x2c, 0x27, 0x50,
	0xe5, 0xfd, 0x8f, 0x9b, 0x06, 0x74, 0x02, 0xe4, 0x3c, 0xb3, 0x40, 0x25, 0x29, 0xae, 0xd9, 0x17,
	0x9e, 0xba, 0x2e, 0x72, 0x2e, 0xbd, 0xeb, 0x3a, 0xc5, 0xd5, 0x4a, 0xab, 0x8b, 0xe8, 0x20, 0xf5,
	0x2e, 0xbe, 0xcf, 0xd9, 0x02, 0x05, 0x18, 0x4c, 0x31, 0x39, 0x9d, 0xa0, 0xae, 0x0a, 0x8c, 0xbd,
	0x85, 0xd0, 0xda, 0xa9, 0xd0, 0xd6, 0xd3, 0x6a, 0xcf, 0x59, 0xda, 0x97, 0xbe, 0x5f, 0x64, 0xae,
	0xd0, 0x35, 0xaa, 0x74, 0xd6, 0x01, 0x6c, 0xd4, 0x3e, 0xcd, 0xf5, 0x5a, 0xf6, 0xb6, 0xac, 0x43,
	0xb3, 0xaa, 0xc3, 0x71, 0x55, 0x87, 0x77, 0x8c, 0x8b, 0x7f, 0xa9, 0xe6, 0xcc, 0x62, 0x34, 0xcf,
	0x2e, 0xc6, 0x3a, 0xb4, 0x7c, 0x16, 0x30, 0x21, 0x9b, 0x6f, 0x39, 0xea, 0x62, 0x7d, 0x86, 0x9b,
	0xe5, 0x64, 0x7c, 0x51, 0xfc, 0x1e, 0x5c, 0x2d, 0x4b, 0xcb, 0x8d, 0x86, 0x1c, 0xc4, 0xdb, 0xcb,
	0x06, 0xd1, 0xa9, 0x3e, 0xd9, 0xf9, 0xad, 0x41, 0x5b, 0xe2, 0xe6, 0xe4, 0x2d, 0xe8, 0x1f, 0x31,
	0xf4, 0xc8, 0xd6, 0x8a, 0x1f, 0x25, 0xd3, 0xaa, 0x03, 0x9c, 0x19, 0xc4, 0xf7, 0xd0, 0xce, 0x37,
	0xe7, 0x52, 0xe8, 0xf6, 0xa1, 0xeb, 0xa0, 0x48, 0x18, 0x9e, 0x20, 0xd9, 0xa8, 0xef, 0x4f, 0x91,
	0x2d, 0x6d, 0x9e, 0xbc, 0x01, 0x3d, 0xf3, 0xaa, 0xbe, 0xa6, 0x92, 0x8b, 0xe6, 0x66, 0x1d, 0x80,
	0x97, 0xca, 0xd1, 0xb3, 0x49, 0xaa, 0xe7, 0x29, 0x6d, 0xc5, 0x8a, 0x72, 0x1c, 0xe8, 0xe4, 0x03,
	0x49, 0xac, 0xf3, 0x98, 0x4e, 0x07, 0xdd, 0xbc, 0xbb, 0x14, 0xa3, 0x38, 0xf7, 0xac, 0x4f, 0x83,
	0x19, 0x13, 0x5f, 0xd2, 0xa9, 0xed, 0x46, 0xc1, 0xe8, 0x38, 0x9d, 0xe2, 0x77, 0xf5, 0xdf, 0xfc,
	0x60, 0x16, 0x8d, 0xd4, 0xfb, 0x69, 0x5b, 0x06, 0x1e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x37,
	0xe8, 0x0d, 0x30, 0xd2, 0x07, 0x00, 0x00,
}
