// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/storage.proto

package squzy_storage_v1_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
	Type_Tcp           Type = 0
	Type_Grpc          Type = 1
	Type_Http          Type = 2
	Type_SiteMap       Type = 3
	Type_MongoDb       Type = 4
	Type_PostgresDb    Type = 5
	Type_CassandraDb   Type = 6
	Type_MysqlDb       Type = 7
	Type_JsonHttpValue Type = 8
)

var Type_name = map[int32]string{
	0: "Tcp",
	1: "Grpc",
	2: "Http",
	3: "SiteMap",
	4: "MongoDb",
	5: "PostgresDb",
	6: "CassandraDb",
	7: "MysqlDb",
	8: "JsonHttpValue",
}

var Type_value = map[string]int32{
	"Tcp":           0,
	"Grpc":          1,
	"Http":          2,
	"SiteMap":       3,
	"MongoDb":       4,
	"PostgresDb":    5,
	"CassandraDb":   6,
	"MysqlDb":       7,
	"JsonHttpValue": 8,
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
	Code                 StatusCode     `protobuf:"varint,1,opt,name=code,proto3,enum=squzy.storage.v1.service.StatusCode" json:"code,omitempty"`
	Description          string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Meta                 *MetaData      `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	Value                *_struct.Value `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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

func (m *Log) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type MetaData struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location             string               `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Port                 int32                `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Type                 Type                 `protobuf:"varint,6,opt,name=type,proto3,enum=squzy.storage.v1.service.Type" json:"type,omitempty"`
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

func (m *MetaData) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *MetaData) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
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
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x13, 0xe7, 0x36, 0x11, 0xc1, 0x8c, 0x44, 0x65, 0x45, 0x40, 0xa3, 0x88, 0x87, 0xaa,
	0x42, 0x0e, 0x35, 0x08, 0xf5, 0xbd, 0x41, 0xdc, 0x12, 0x81, 0xdc, 0x88, 0xf7, 0xb5, 0x3d, 0x2c,
	0x46, 0x8e, 0xd7, 0xdd, 0x59, 0x47, 0x0a, 0x12, 0xfc, 0x1e, 0x7f, 0xc3, 0x37, 0x20, 0x6f, 0x64,
	0x5a, 0x0a, 0x01, 0xde, 0x76, 0x66, 0xcf, 0x99, 0x99, 0x73, 0x76, 0x16, 0x0e, 0x4b, 0xad, 0x8c,
	0x9a, 0x6d, 0x4e, 0x67, 0x6c, 0x94, 0x16, 0x92, 0x02, 0x9b, 0x40, 0x9f, 0x2f, 0xab, 0xcf, 0xdb,
	0xa0, 0x49, 0x6e, 0x4e, 0x03, 0x26, 0xbd, 0xc9, 0x12, 0x1a, 0x1f, 0x49, 0xa5, 0x64, 0x4e, 0x33,
	0x8b, 0x8b, 0xab, 0x0f, 0x33, 0x93, 0xad, 0x89, 0x8d, 0x58, 0x97, 0x3b, 0xea, 0xf8, 0xde, 0x4d,
	0x00, 0x1b, 0x5d, 0x25, 0x66, 0x77, 0x3b, 0xfd, 0xe6, 0x40, 0x7b, 0xa1, 0x24, 0x9e, 0x81, 0x9b,
	0xa8, 0x94, 0x7c, 0x67, 0xe2, 0x1c, 0x8f, 0xc2, 0x87, 0xc1, 0xbe, 0x7e, 0xc1, 0x85, 0x11, 0xa6,
	0xe2, 0x73, 0x95, 0x52, 0x64, 0x19, 0x38, 0x81, 0x61, 0x4a, 0x9c, 0xe8, 0xac, 0x34, 0x99, 0x2a,
	0xfc, 0xd6, 0xc4, 0x39, 0x1e, 0x44, 0xd7, 0x53, 0xf8, 0x0c, 0xdc, 0x35, 0x19, 0xe1, 0xb7, 0x27,
	0xce, 0xf1, 0x30, 0x9c, 0xee, 0xaf, 0xbd, 0x24, 0x23, 0xe6, 0xc2, 0x88, 0xc8, 0xe2, 0xf1, 0x11,
	0x74, 0x36, 0x22, 0xaf, 0xc8, 0x77, 0x2d, 0xf1, 0x30, 0xd8, 0x29, 0x09, 0x1a, 0x25, 0xc1, 0xfb,
	0xfa, 0x36, 0xda, 0x81, 0xa6, 0xdf, 0x1d, 0xe8, 0x37, 0x05, 0x70, 0x04, 0xad, 0x2c, 0xb5, 0x62,
	0x06, 0x51, 0x2b, 0x4b, 0x71, 0x0c, 0xfd, 0x5c, 0x25, 0xe2, 0xda, 0x84, 0x3f, 0x63, 0x44, 0x70,
	0x4b, 0xa5, 0x8d, 0x1d, 0xaf, 0x13, 0xd9, 0x33, 0x9e, 0xc1, 0x80, 0x8d, 0xd0, 0x66, 0x95, 0xad,
	0x9b, 0xf6, 0xe3, 0xdf, 0xda, 0xaf, 0x1a, 0xa7, 0xa3, 0x2b, 0x30, 0x3e, 0x85, 0x1e, 0x15, 0xa9,
	0xe5, 0x75, 0xfe, 0xc9, 0x6b, 0xa0, 0x18, 0x82, 0x6b, 0xb6, 0x25, 0xf9, 0x5d, 0x6b, 0xff, 0x83,
	0xfd, 0x16, 0xad, 0xb6, 0x25, 0x45, 0x16, 0x3b, 0xfd, 0x04, 0x77, 0x2f, 0xa8, 0x48, 0x17, 0x4a,
	0x2e, 0x89, 0x59, 0x48, 0x8a, 0xe8, 0xb2, 0x22, 0x36, 0x38, 0x83, 0x76, 0xae, 0xa4, 0x55, 0x3f,
	0x0c, 0xef, 0xef, 0xaf, 0xb5, 0x50, 0x32, 0xaa, 0x91, 0xf5, 0x13, 0x72, 0xf2, 0x91, 0xd2, 0x2a,
	0x27, 0xfd, 0x2a, 0x6d, 0x9e, 0xf0, 0x5a, 0x6a, 0x1a, 0xc2, 0xe1, 0xcd, 0x5e, 0x5c, 0xaa, 0x82,
	0x09, 0x7d, 0xe8, 0x71, 0x95, 0x24, 0xc4, 0x6c, 0x1b, 0xf6, 0xa3, 0x26, 0x3c, 0xf9, 0x0a, 0x6e,
	0x3d, 0x2d, 0xf6, 0xa0, 0xbd, 0x4a, 0x4a, 0xef, 0x00, 0xfb, 0xe0, 0xbe, 0xd0, 0x65, 0xe2, 0x39,
	0xf5, 0xe9, 0xa5, 0x31, 0xa5, 0xd7, 0xc2, 0x21, 0xf4, 0x2e, 0x32, 0x43, 0x4b, 0x51, 0x7a, 0xed,
	0x3a, 0x58, 0xaa, 0x42, 0xaa, 0x79, 0xec, 0xb9, 0x38, 0x02, 0x78, 0xa7, 0xd8, 0x48, 0x4d, 0x3c,
	0x8f, 0xbd, 0x0e, 0xde, 0x86, 0xe1, 0xb9, 0x60, 0x16, 0x45, 0xaa, 0xc5, 0x3c, 0xf6, 0xba, 0x16,
	0xbd, 0xe5, 0xcb, 0x7c, 0x1e, 0x7b, 0x3d, 0xbc, 0x03, 0xb7, 0x5e, 0xb3, 0x2a, 0xea, 0xaa, 0x76,
	0x2b, 0xbc, 0xfe, 0xc9, 0x11, 0xc0, 0xd5, 0xb2, 0x62, 0x17, 0x5a, 0x6f, 0xdf, 0x78, 0x07, 0x38,
	0x80, 0xce, 0x73, 0xad, 0x95, 0xf6, 0x9c, 0xf0, 0x0b, 0x74, 0x17, 0x4a, 0x4a, 0xd2, 0xc8, 0x30,
	0xfa, 0x55, 0x1e, 0xce, 0xfe, 0xf2, 0x03, 0xfe, 0x64, 0xfa, 0xf8, 0xf1, 0xff, 0x13, 0x76, 0xce,
	0xc5, 0x5d, 0xbb, 0x10, 0x4f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xd5, 0x94, 0x31, 0xf4,
	0x03, 0x00, 0x00,
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
