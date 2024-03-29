// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/v1/squzy_notification_server.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotificationManagerClient is the client API for NotificationManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationManagerClient interface {
	// protolint:disable:next MAX_LINE_LENGTH
	CreateNotificationMethod(ctx context.Context, in *CreateNotificationMethodRequest, opts ...grpc.CallOption) (*NotificationMethod, error)
	GetById(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error)
	GetNotificationMethods(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListResponse, error)
	DeleteById(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error)
	Activate(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error)
	Deactivate(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error)
	Add(ctx context.Context, in *NotificationMethodRequest, opts ...grpc.CallOption) (*NotificationMethod, error)
	Remove(ctx context.Context, in *NotificationMethodRequest, opts ...grpc.CallOption) (*NotificationMethod, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type notificationManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationManagerClient(cc grpc.ClientConnInterface) NotificationManagerClient {
	return &notificationManagerClient{cc}
}

func (c *notificationManagerClient) CreateNotificationMethod(ctx context.Context, in *CreateNotificationMethodRequest, opts ...grpc.CallOption) (*NotificationMethod, error) {
	out := new(NotificationMethod)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/CreateNotificationMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) GetById(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error) {
	out := new(NotificationMethod)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) GetNotificationMethods(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/GetNotificationMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) DeleteById(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error) {
	out := new(NotificationMethod)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) Activate(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error) {
	out := new(NotificationMethod)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) Deactivate(ctx context.Context, in *NotificationMethodIdRequest, opts ...grpc.CallOption) (*NotificationMethod, error) {
	out := new(NotificationMethod)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) Add(ctx context.Context, in *NotificationMethodRequest, opts ...grpc.CallOption) (*NotificationMethod, error) {
	out := new(NotificationMethod)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) Remove(ctx context.Context, in *NotificationMethodRequest, opts ...grpc.CallOption) (*NotificationMethod, error) {
	out := new(NotificationMethod)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagerClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/squzy.v1.notification.NotificationManager/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationManagerServer is the server API for NotificationManager service.
// All implementations must embed UnimplementedNotificationManagerServer
// for forward compatibility
type NotificationManagerServer interface {
	// protolint:disable:next MAX_LINE_LENGTH
	CreateNotificationMethod(context.Context, *CreateNotificationMethodRequest) (*NotificationMethod, error)
	GetById(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error)
	GetNotificationMethods(context.Context, *emptypb.Empty) (*GetListResponse, error)
	DeleteById(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error)
	Activate(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error)
	Deactivate(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error)
	Add(context.Context, *NotificationMethodRequest) (*NotificationMethod, error)
	Remove(context.Context, *NotificationMethodRequest) (*NotificationMethod, error)
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Notify(context.Context, *NotifyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNotificationManagerServer()
}

// UnimplementedNotificationManagerServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationManagerServer struct {
}

func (UnimplementedNotificationManagerServer) CreateNotificationMethod(context.Context, *CreateNotificationMethodRequest) (*NotificationMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotificationMethod not implemented")
}
func (UnimplementedNotificationManagerServer) GetById(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedNotificationManagerServer) GetNotificationMethods(context.Context, *emptypb.Empty) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationMethods not implemented")
}
func (UnimplementedNotificationManagerServer) DeleteById(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedNotificationManagerServer) Activate(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedNotificationManagerServer) Deactivate(context.Context, *NotificationMethodIdRequest) (*NotificationMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedNotificationManagerServer) Add(context.Context, *NotificationMethodRequest) (*NotificationMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedNotificationManagerServer) Remove(context.Context, *NotificationMethodRequest) (*NotificationMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedNotificationManagerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedNotificationManagerServer) Notify(context.Context, *NotifyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedNotificationManagerServer) mustEmbedUnimplementedNotificationManagerServer() {}

// UnsafeNotificationManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationManagerServer will
// result in compilation errors.
type UnsafeNotificationManagerServer interface {
	mustEmbedUnimplementedNotificationManagerServer()
}

func RegisterNotificationManagerServer(s grpc.ServiceRegistrar, srv NotificationManagerServer) {
	s.RegisterService(&NotificationManager_ServiceDesc, srv)
}

func _NotificationManager_CreateNotificationMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).CreateNotificationMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/CreateNotificationMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).CreateNotificationMethod(ctx, req.(*CreateNotificationMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationMethodIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).GetById(ctx, req.(*NotificationMethodIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_GetNotificationMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).GetNotificationMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/GetNotificationMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).GetNotificationMethods(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationMethodIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).DeleteById(ctx, req.(*NotificationMethodIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationMethodIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).Activate(ctx, req.(*NotificationMethodIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationMethodIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).Deactivate(ctx, req.(*NotificationMethodIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).Add(ctx, req.(*NotificationMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).Remove(ctx, req.(*NotificationMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManager_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagerServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squzy.v1.notification.NotificationManager/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagerServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationManager_ServiceDesc is the grpc.ServiceDesc for NotificationManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "squzy.v1.notification.NotificationManager",
	HandlerType: (*NotificationManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotificationMethod",
			Handler:    _NotificationManager_CreateNotificationMethod_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _NotificationManager_GetById_Handler,
		},
		{
			MethodName: "GetNotificationMethods",
			Handler:    _NotificationManager_GetNotificationMethods_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _NotificationManager_DeleteById_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _NotificationManager_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _NotificationManager_Deactivate_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _NotificationManager_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _NotificationManager_Remove_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _NotificationManager_GetList_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _NotificationManager_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/squzy_notification_server.proto",
}
