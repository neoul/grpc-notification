// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/notification.proto

package notification

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Subscription struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb4fc010f5c4b1c, []int{0}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Notification struct {
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb4fc010f5c4b1c, []int{1}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Subscription)(nil), "Subscription")
	proto.RegisterType((*Notification)(nil), "Notification")
}

func init() { proto.RegisterFile("proto/notification.proto", fileDescriptor_beb4fc010f5c4b1c) }

var fileDescriptor_beb4fc010f5c4b1c = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03,
	0x0b, 0x29, 0x29, 0x71, 0xf1, 0x04, 0x97, 0x26, 0x15, 0x27, 0x17, 0x65, 0x16, 0x80, 0x44, 0x85,
	0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c,
	0x25, 0x0d, 0x2e, 0x1e, 0x3f, 0x24, 0x9d, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89,
	0xe9, 0xa9, 0x12, 0xcc, 0x60, 0x65, 0x30, 0xae, 0x91, 0x2d, 0x17, 0x0f, 0xb2, 0x1d, 0x42, 0xba,
	0x5c, 0x9c, 0x50, 0xd3, 0x93, 0x52, 0x85, 0x78, 0xf5, 0x90, 0x6d, 0x92, 0xe2, 0xd5, 0x43, 0x36,
	0x54, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xcb, 0x53, 0x66, 0x21, 0xaf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Notification_SubscribeClient, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Notification_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Notification_serviceDesc.Streams[0], "/notification/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationSubscribeClient{stream}
	return x, nil
}

type Notification_SubscribeClient interface {
	Send(*Subscription) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type notificationSubscribeClient struct {
	grpc.ClientStream
}

func (x *notificationSubscribeClient) Send(m *Subscription) error {
	return x.ClientStream.SendMsg(m)
}

func (x *notificationSubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServer is the server API for Notification service.
type NotificationServer interface {
	Subscribe(Notification_SubscribeServer) error
}

// UnimplementedNotificationServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (*UnimplementedNotificationServer) Subscribe(srv Notification_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterNotificationServer(s *grpc.Server, srv NotificationServer) {
	s.RegisterService(&_Notification_serviceDesc, srv)
}

func _Notification_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NotificationServer).Subscribe(&notificationSubscribeServer{stream})
}

type Notification_SubscribeServer interface {
	Send(*Notification) error
	Recv() (*Subscription, error)
	grpc.ServerStream
}

type notificationSubscribeServer struct {
	grpc.ServerStream
}

func (x *notificationSubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *notificationSubscribeServer) Recv() (*Subscription, error) {
	m := new(Subscription)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Notification_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notification",
	HandlerType: (*NotificationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Notification_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/notification.proto",
}
