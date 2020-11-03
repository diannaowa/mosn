// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/plugin/proto/plugin.proto

package proto

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

type Request struct {
	Header               map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Trailer              map[string]string `protobuf:"bytes,2,rep,name=trailer,proto3" json:"trailer,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Type                 string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_5ee291f37a2d3192, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetTrailer() map[string]string {
	if m != nil {
		return m.Trailer
	}
	return nil
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Response struct {
	Header               map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Trailer              map[string]string `protobuf:"bytes,2,rep,name=trailer,proto3" json:"trailer,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Status               int32             `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_5ee291f37a2d3192, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetTrailer() map[string]string {
	if m != nil {
		return m.Trailer
	}
	return nil
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterMapType((map[string]string)(nil), "proto.Request.HeaderEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.Request.TrailerEntry")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterMapType((map[string]string)(nil), "proto.Response.HeaderEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.Response.TrailerEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type pluginClient struct {
	cc *grpc.ClientConn
}

func NewPluginClient(cc *grpc.ClientConn) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Plugin/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	Call(context.Context, *Request) (*Response, error)
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Plugin/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Plugin_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/plugin/proto/plugin.proto",
}

func init() {
	proto.RegisterFile("pkg/plugin/proto/plugin.proto", fileDescriptor_plugin_5ee291f37a2d3192)
}

var fileDescriptor_plugin_5ee291f37a2d3192 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x9a, 0xa4, 0x3a, 0x2d, 0x2a, 0x83, 0xc8, 0x12, 0x15, 0x42, 0x2f, 0xe6, 0x94,
	0x62, 0x8a, 0xa2, 0xbd, 0x8a, 0xe0, 0x51, 0x16, 0x5f, 0x20, 0xa5, 0x43, 0x2d, 0x0d, 0x49, 0xdc,
	0xdd, 0x08, 0x79, 0x0e, 0x6f, 0x3e, 0xad, 0x74, 0x36, 0x62, 0x5d, 0x3c, 0xe8, 0xc9, 0x53, 0xe6,
	0x9f, 0xf9, 0x7f, 0x86, 0xf9, 0xb2, 0x70, 0xde, 0x6c, 0x56, 0xd3, 0xa6, 0x6c, 0x57, 0xeb, 0x6a,
	0xda, 0xa8, 0xda, 0xd4, 0xbd, 0xc8, 0x58, 0x60, 0xc8, 0x9f, 0xc9, 0x9b, 0x0f, 0x43, 0x49, 0x2f,
	0x2d, 0x69, 0x83, 0x39, 0x44, 0xcf, 0x54, 0x2c, 0x49, 0x09, 0x2f, 0x19, 0xa4, 0xa3, 0x3c, 0xb6,
	0xd6, 0xac, 0x9f, 0x67, 0x0f, 0x3c, 0xbc, 0xaf, 0x8c, 0xea, 0x64, 0xef, 0xc4, 0x2b, 0x18, 0x1a,
	0x55, 0xac, 0x4b, 0x52, 0xc2, 0xe7, 0xd0, 0xa9, 0x13, 0x7a, 0xb2, 0x53, 0x9b, 0xfa, 0xf4, 0x22,
	0x42, 0xb0, 0xa8, 0x97, 0x9d, 0x18, 0x24, 0x5e, 0x3a, 0x96, 0x5c, 0x6f, 0x7b, 0xa6, 0x6b, 0x48,
	0x04, 0x89, 0x97, 0xee, 0x4b, 0xae, 0xe3, 0x5b, 0x18, 0xed, 0x6c, 0xc5, 0x23, 0x18, 0x6c, 0xa8,
	0x13, 0x1e, 0x3b, 0xb6, 0x25, 0x1e, 0x43, 0xf8, 0x5a, 0x94, 0x2d, 0x09, 0x9f, 0x7b, 0x56, 0xcc,
	0xfd, 0x1b, 0x2f, 0x9e, 0xc3, 0x78, 0x77, 0xf7, 0x5f, 0xb2, 0x93, 0x77, 0x1f, 0xf6, 0x24, 0xe9,
	0xa6, 0xae, 0x34, 0xe1, 0xcc, 0xc1, 0xf2, 0x75, 0xa1, 0x35, 0xfc, 0xc8, 0xe5, 0xda, 0xe5, 0x72,
	0xe6, 0xa6, 0x7e, 0x0f, 0xe6, 0x04, 0x22, 0x6d, 0x0a, 0xd3, 0x6a, 0x46, 0x13, 0xca, 0x5e, 0xfd,
	0x13, 0x9c, 0xfc, 0x12, 0xa2, 0x47, 0x7e, 0x49, 0x78, 0x01, 0xc1, 0x5d, 0x51, 0x96, 0x78, 0xf0,
	0xfd, 0x9f, 0xc7, 0x87, 0xce, 0xad, 0x8b, 0x88, 0xf5, 0xec, 0x23, 0x00, 0x00, 0xff, 0xff, 0x97,
	0x05, 0xe4, 0x14, 0x94, 0x02, 0x00, 0x00,
}
