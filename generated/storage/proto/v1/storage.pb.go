// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/storage.proto

package squzy_storage_v1_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Type int32

const (
	Type_Tcp     Type = 0
	Type_Grpc    Type = 1
	Type_Http    Type = 2
	Type_SiteMap Type = 3
)

var Type_name = map[int32]string{
	0: "Tcp",
	1: "Grpc",
	2: "Http",
	3: "SiteMap",
}

var Type_value = map[string]int32{
	"Tcp":     0,
	"Grpc":    1,
	"Http":    2,
	"SiteMap": 3,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65bd54e225cb21b3, []int{0}
}

type StatusCode int32

const (
	StatusCode_OK    StatusCode = 0
	StatusCode_Error StatusCode = 1
)

var StatusCode_name = map[int32]string{
	0: "OK",
	1: "Error",
}

var StatusCode_value = map[string]int32{
	"OK":    0,
	"Error": 1,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65bd54e225cb21b3, []int{1}
}

type Log struct {
	Code                 StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=squzy.storage.v1.service.StatusCode" json:"code,omitempty"`
	Description          string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Meta                 *MetaData  `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bd54e225cb21b3, []int{0}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_OK
}

func (m *Log) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Log) GetMeta() *MetaData {
	if m != nil {
		return m.Meta
	}
	return nil
}

type MetaData struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location             string               `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Port                 int32                `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Type                 Type                 `protobuf:"varint,5,opt,name=type,proto3,enum=squzy.storage.v1.service.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MetaData) Reset()         { *m = MetaData{} }
func (m *MetaData) String() string { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()    {}
func (*MetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bd54e225cb21b3, []int{1}
}

func (m *MetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaData.Unmarshal(m, b)
}
func (m *MetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaData.Marshal(b, m, deterministic)
}
func (m *MetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaData.Merge(m, src)
}
func (m *MetaData) XXX_Size() int {
	return xxx_messageInfo_MetaData.Size(m)
}
func (m *MetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetaData proto.InternalMessageInfo

func (m *MetaData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MetaData) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *MetaData) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *MetaData) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *MetaData) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_Tcp
}

type SendLogMessageRequest struct {
	Log                  *Log     `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	SchedulerId          string   `protobuf:"bytes,2,opt,name=schedulerId,proto3" json:"schedulerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendLogMessageRequest) Reset()         { *m = SendLogMessageRequest{} }
func (m *SendLogMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendLogMessageRequest) ProtoMessage()    {}
func (*SendLogMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bd54e225cb21b3, []int{2}
}

func (m *SendLogMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendLogMessageRequest.Unmarshal(m, b)
}
func (m *SendLogMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendLogMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendLogMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendLogMessageRequest.Merge(m, src)
}
func (m *SendLogMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendLogMessageRequest.Size(m)
}
func (m *SendLogMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendLogMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendLogMessageRequest proto.InternalMessageInfo

func (m *SendLogMessageRequest) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *SendLogMessageRequest) GetSchedulerId() string {
	if m != nil {
		return m.SchedulerId
	}
	return ""
}

type SendLogMessageResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendLogMessageResponse) Reset()         { *m = SendLogMessageResponse{} }
func (m *SendLogMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendLogMessageResponse) ProtoMessage()    {}
func (*SendLogMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bd54e225cb21b3, []int{3}
}

func (m *SendLogMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendLogMessageResponse.Unmarshal(m, b)
}
func (m *SendLogMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendLogMessageResponse.Marshal(b, m, deterministic)
}
func (m *SendLogMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendLogMessageResponse.Merge(m, src)
}
func (m *SendLogMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendLogMessageResponse.Size(m)
}
func (m *SendLogMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendLogMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendLogMessageResponse proto.InternalMessageInfo

func (m *SendLogMessageResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("squzy.storage.v1.service.Type", Type_name, Type_value)
	proto.RegisterEnum("squzy.storage.v1.service.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*Log)(nil), "squzy.storage.v1.service.Log")
	proto.RegisterType((*MetaData)(nil), "squzy.storage.v1.service.MetaData")
	proto.RegisterType((*SendLogMessageRequest)(nil), "squzy.storage.v1.service.SendLogMessageRequest")
	proto.RegisterType((*SendLogMessageResponse)(nil), "squzy.storage.v1.service.SendLogMessageResponse")
}

func init() { proto.RegisterFile("proto/v1/storage.proto", fileDescriptor_65bd54e225cb21b3) }

var fileDescriptor_65bd54e225cb21b3 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdb, 0x8a, 0x13, 0x41,
	0x10, 0xdd, 0xb9, 0xe4, 0x56, 0x81, 0x30, 0x14, 0xb8, 0x0c, 0x01, 0xdd, 0x30, 0xf8, 0x10, 0xf6,
	0x61, 0x66, 0x77, 0x04, 0xf1, 0x5d, 0x45, 0xc5, 0x04, 0xa1, 0x93, 0x1f, 0xe8, 0xed, 0x29, 0xdb,
	0x91, 0x24, 0xdd, 0xe9, 0xee, 0x09, 0x44, 0xf0, 0x43, 0xfc, 0x09, 0xbf, 0x51, 0xa6, 0xc3, 0x68,
	0x14, 0x23, 0xbe, 0x75, 0x1d, 0x4e, 0xd5, 0x39, 0x75, 0xaa, 0xe1, 0x5a, 0x1b, 0xe5, 0x54, 0x71,
	0xb8, 0x2f, 0xac, 0x53, 0x86, 0x4b, 0xca, 0x3d, 0x80, 0xa9, 0xdd, 0x37, 0x5f, 0x8e, 0x79, 0x07,
	0x1e, 0xee, 0x73, 0x4b, 0xe6, 0x50, 0x0b, 0x9a, 0xde, 0x48, 0xa5, 0xe4, 0x86, 0x0a, 0xcf, 0x7b,
	0x68, 0x3e, 0x16, 0xae, 0xde, 0x92, 0x75, 0x7c, 0xab, 0x4f, 0xad, 0xd9, 0xb7, 0x00, 0xa2, 0x85,
	0x92, 0xf8, 0x02, 0x62, 0xa1, 0x2a, 0x4a, 0x83, 0x59, 0x30, 0x9f, 0x94, 0x4f, 0xf3, 0x4b, 0x13,
	0xf3, 0x95, 0xe3, 0xae, 0xb1, 0x2f, 0x55, 0x45, 0xcc, 0x77, 0xe0, 0x0c, 0xc6, 0x15, 0x59, 0x61,
	0x6a, 0xed, 0x6a, 0xb5, 0x4b, 0xc3, 0x59, 0x30, 0x1f, 0xb1, 0x73, 0x08, 0x9f, 0x43, 0xbc, 0x25,
	0xc7, 0xd3, 0x68, 0x16, 0xcc, 0xc7, 0x65, 0x76, 0x79, 0xf6, 0x92, 0x1c, 0x7f, 0xc5, 0x1d, 0x67,
	0x9e, 0x9f, 0x7d, 0x0f, 0x60, 0xd8, 0x41, 0x38, 0x81, 0xb0, 0xae, 0xbc, 0xbd, 0x11, 0x0b, 0xeb,
	0x0a, 0xa7, 0x30, 0xdc, 0x28, 0xc1, 0xcf, 0x34, 0x7f, 0xd6, 0x88, 0x10, 0x6b, 0x65, 0x9c, 0x17,
	0xec, 0x31, 0xff, 0xc6, 0x1c, 0xe2, 0x76, 0xf7, 0x34, 0xf6, 0x26, 0xa6, 0xf9, 0x29, 0x98, 0xbc,
	0x0b, 0x26, 0x5f, 0x77, 0xc1, 0x30, 0xcf, 0xc3, 0x12, 0x62, 0x77, 0xd4, 0x94, 0xf6, 0x7c, 0x20,
	0x4f, 0x2e, 0x9b, 0x5e, 0x1f, 0x35, 0x31, 0xcf, 0xcd, 0x3e, 0xc3, 0xa3, 0x15, 0xed, 0xaa, 0x85,
	0x92, 0x4b, 0xb2, 0x96, 0x4b, 0x62, 0xb4, 0x6f, 0xc8, 0x3a, 0x2c, 0x20, 0xda, 0x28, 0xe9, 0xdd,
	0x8f, 0xcb, 0xc7, 0x97, 0x67, 0x2d, 0x94, 0x64, 0x2d, 0xb3, 0x0d, 0xd5, 0x8a, 0x4f, 0x54, 0x35,
	0x1b, 0x32, 0xef, 0xaa, 0x2e, 0xd4, 0x33, 0x28, 0x2b, 0xe1, 0xfa, 0x4f, 0x2d, 0xab, 0xd5, 0xce,
	0x12, 0xa6, 0x30, 0xb0, 0x8d, 0x10, 0x64, 0xad, 0x17, 0x1c, 0xb2, 0xae, 0xbc, 0xbd, 0x83, 0xb8,
	0x75, 0x8b, 0x03, 0x88, 0xd6, 0x42, 0x27, 0x57, 0x38, 0x84, 0xf8, 0x8d, 0xd1, 0x22, 0x09, 0xda,
	0xd7, 0x5b, 0xe7, 0x74, 0x12, 0xe2, 0x18, 0x06, 0xab, 0xda, 0xd1, 0x92, 0xeb, 0x24, 0xba, 0xbd,
	0x01, 0xf8, 0x75, 0x70, 0xec, 0x43, 0xf8, 0xe1, 0x7d, 0x72, 0x85, 0x23, 0xe8, 0xbd, 0x36, 0x46,
	0x99, 0x24, 0x28, 0xbf, 0x42, 0x7f, 0xa1, 0xa4, 0x24, 0x83, 0x16, 0x26, 0xbf, 0x1b, 0xc2, 0xe2,
	0x1f, 0xbf, 0xe8, 0x6f, 0x31, 0x4d, 0xef, 0xfe, 0xbf, 0xe1, 0xb4, 0xeb, 0x43, 0xdf, 0xdf, 0xef,
	0xd9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x69, 0x00, 0xa9, 0x1a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggerClient interface {
	SendLogMessage(ctx context.Context, in *SendLogMessageRequest, opts ...grpc.CallOption) (*SendLogMessageResponse, error)
}

type loggerClient struct {
	cc *grpc.ClientConn
}

func NewLoggerClient(cc *grpc.ClientConn) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) SendLogMessage(ctx context.Context, in *SendLogMessageRequest, opts ...grpc.CallOption) (*SendLogMessageResponse, error) {
	out := new(SendLogMessageResponse)
	err := c.cc.Invoke(ctx, "/squzy.storage.v1.service.Logger/SendLogMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
type LoggerServer interface {
	SendLogMessage(context.Context, *SendLogMessageRequest) (*SendLogMessageResponse, error)
}

// UnimplementedLoggerServer can be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (*UnimplementedLoggerServer) SendLogMessage(ctx context.Context, req *SendLogMessageRequest) (*SendLogMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogMessage not implemented")
}

func RegisterLoggerServer(s *grpc.Server, srv LoggerServer) {
	s.RegisterService(&_Logger_serviceDesc, srv)
}

func _Logger_SendLogMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendLogMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).SendLogMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.storage.v1.service.Logger/SendLogMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).SendLogMessage(ctx, req.(*SendLogMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squzy.storage.v1.service.Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLogMessage",
			Handler:    _Logger_SendLogMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/storage.proto",
}
