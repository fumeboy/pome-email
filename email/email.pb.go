// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email.proto

package email

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

type SendRequest struct {
	To                   []string `protobuf:"bytes,1,rep,name=To,proto3" json:"To,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	ContentType          string   `protobuf:"bytes,4,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *SendRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type PomeAsyncResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PomeAsyncResp) Reset()         { *m = PomeAsyncResp{} }
func (m *PomeAsyncResp) String() string { return proto.CompactTextString(m) }
func (*PomeAsyncResp) ProtoMessage()    {}
func (*PomeAsyncResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{1}
}

func (m *PomeAsyncResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PomeAsyncResp.Unmarshal(m, b)
}
func (m *PomeAsyncResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PomeAsyncResp.Marshal(b, m, deterministic)
}
func (m *PomeAsyncResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PomeAsyncResp.Merge(m, src)
}
func (m *PomeAsyncResp) XXX_Size() int {
	return xxx_messageInfo_PomeAsyncResp.Size(m)
}
func (m *PomeAsyncResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PomeAsyncResp.DiscardUnknown(m)
}

var xxx_messageInfo_PomeAsyncResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SendRequest)(nil), "email.SendRequest")
	proto.RegisterType((*PomeAsyncResp)(nil), "email.PomeAsyncResp")
}

func init() {
	proto.RegisterFile("email.proto", fileDescriptor_6175298cb4ed6faa)
}

var fileDescriptor_6175298cb4ed6faa = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xcd, 0xca, 0x82, 0x50,
	0x10, 0x40, 0x3f, 0x7f, 0xbe, 0xc2, 0xb1, 0x1f, 0x18, 0x5a, 0x5c, 0x5a, 0x89, 0x2b, 0x57, 0x2e,
	0xec, 0x09, 0x32, 0xda, 0xc7, 0xd5, 0x17, 0x48, 0x9d, 0x85, 0x91, 0x77, 0x4c, 0xaf, 0x81, 0x6f,
	0x1f, 0xde, 0x0a, 0x6c, 0x37, 0xe7, 0x0c, 0x0c, 0x67, 0xc0, 0xa7, 0xe6, 0x5a, 0xdf, 0xe3, 0xb6,
	0x63, 0xcd, 0xf8, 0x6f, 0x20, 0x6c, 0xc0, 0xcf, 0x48, 0x55, 0x92, 0x1e, 0x03, 0xf5, 0x1a, 0x37,
	0x60, 0xe7, 0x2c, 0xac, 0xc0, 0x89, 0x3c, 0x69, 0xe7, 0x8c, 0x02, 0x96, 0xd9, 0x50, 0xdc, 0xa8,
	0xd4, 0xc2, 0x0e, 0xac, 0xc8, 0x93, 0x5f, 0x44, 0x04, 0x37, 0xe5, 0x6a, 0x14, 0x8e, 0xd1, 0x66,
	0xc6, 0x00, 0xfc, 0x13, 0x2b, 0x4d, 0x4a, 0xe7, 0x63, 0x4b, 0xc2, 0x35, 0xab, 0xb9, 0x0a, 0xb7,
	0xb0, 0xbe, 0x70, 0x43, 0xc7, 0x7e, 0x54, 0xa5, 0xa4, 0xbe, 0x4d, 0x52, 0x58, 0x9d, 0xa7, 0x90,
	0x8c, 0xba, 0x67, 0x5d, 0x12, 0x26, 0xe0, 0x4e, 0x3d, 0x88, 0xf1, 0x3b, 0x76, 0x16, 0xb7, 0xdf,
	0x7d, 0xdc, 0xcf, 0x85, 0xf0, 0xaf, 0x58, 0x98, 0x8f, 0x0e, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x55, 0xe4, 0x7c, 0x0e, 0xe0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*PomeAsyncResp, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*PomeAsyncResp, error) {
	out := new(PomeAsyncResp)

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}
	md["pome-ifasync"] = []string{"1"}
	ctx = metadata.NewOutgoingContext(ctx, md)

	err := c.cc.Invoke(ctx, "/email.EmailService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	Send(context.Context, *SendRequest) (*PomeAsyncResp, error)
}

// UnimplementedEmailServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (*UnimplementedEmailServiceServer) Send(ctx context.Context, req *SendRequest) (*PomeAsyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.EmailService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "email.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _EmailService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
