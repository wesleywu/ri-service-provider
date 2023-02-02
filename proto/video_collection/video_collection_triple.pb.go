// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             (unknown)
// source: proto/video_collection/video_collection.proto

package proto_video_collection

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// VideoCollectionClient is the client API for VideoCollection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoCollectionClient interface {
	Count(ctx context.Context, in *VideoCollectionCountReq, opts ...grpc_go.CallOption) (*VideoCollectionCountRes, common.ErrorWithAttachment)
	One(ctx context.Context, in *VideoCollectionOneReq, opts ...grpc_go.CallOption) (*VideoCollectionOneRes, common.ErrorWithAttachment)
	List(ctx context.Context, in *VideoCollectionListReq, opts ...grpc_go.CallOption) (*VideoCollectionListRes, common.ErrorWithAttachment)
	Create(ctx context.Context, in *VideoCollectionCreateReq, opts ...grpc_go.CallOption) (*VideoCollectionCreateRes, common.ErrorWithAttachment)
	Update(ctx context.Context, in *VideoCollectionUpdateReq, opts ...grpc_go.CallOption) (*VideoCollectionUpdateRes, common.ErrorWithAttachment)
	Upsert(ctx context.Context, in *VideoCollectionUpsertReq, opts ...grpc_go.CallOption) (*VideoCollectionUpsertRes, common.ErrorWithAttachment)
	Delete(ctx context.Context, in *VideoCollectionDeleteReq, opts ...grpc_go.CallOption) (*VideoCollectionDeleteRes, common.ErrorWithAttachment)
}

type videoCollectionClient struct {
	cc *triple.TripleConn
}

type VideoCollectionClientImpl struct {
	Count  func(ctx context.Context, in *VideoCollectionCountReq) (*VideoCollectionCountRes, error)
	One    func(ctx context.Context, in *VideoCollectionOneReq) (*VideoCollectionOneRes, error)
	List   func(ctx context.Context, in *VideoCollectionListReq) (*VideoCollectionListRes, error)
	Create func(ctx context.Context, in *VideoCollectionCreateReq) (*VideoCollectionCreateRes, error)
	Update func(ctx context.Context, in *VideoCollectionUpdateReq) (*VideoCollectionUpdateRes, error)
	Upsert func(ctx context.Context, in *VideoCollectionUpsertReq) (*VideoCollectionUpsertRes, error)
	Delete func(ctx context.Context, in *VideoCollectionDeleteReq) (*VideoCollectionDeleteRes, error)
}

func (c *VideoCollectionClientImpl) GetDubboStub(cc *triple.TripleConn) VideoCollectionClient {
	return NewVideoCollectionClient(cc)
}

func (c *VideoCollectionClientImpl) XXX_InterfaceName() string {
	return "proto.video_collection.VideoCollection"
}

func NewVideoCollectionClient(cc *triple.TripleConn) VideoCollectionClient {
	return &videoCollectionClient{cc}
}

func (c *videoCollectionClient) Count(ctx context.Context, in *VideoCollectionCountReq, opts ...grpc_go.CallOption) (*VideoCollectionCountRes, common.ErrorWithAttachment) {
	out := new(VideoCollectionCountRes)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Count", in, out)
}

func (c *videoCollectionClient) One(ctx context.Context, in *VideoCollectionOneReq, opts ...grpc_go.CallOption) (*VideoCollectionOneRes, common.ErrorWithAttachment) {
	out := new(VideoCollectionOneRes)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/One", in, out)
}

func (c *videoCollectionClient) List(ctx context.Context, in *VideoCollectionListReq, opts ...grpc_go.CallOption) (*VideoCollectionListRes, common.ErrorWithAttachment) {
	out := new(VideoCollectionListRes)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/List", in, out)
}

func (c *videoCollectionClient) Create(ctx context.Context, in *VideoCollectionCreateReq, opts ...grpc_go.CallOption) (*VideoCollectionCreateRes, common.ErrorWithAttachment) {
	out := new(VideoCollectionCreateRes)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Create", in, out)
}

func (c *videoCollectionClient) Update(ctx context.Context, in *VideoCollectionUpdateReq, opts ...grpc_go.CallOption) (*VideoCollectionUpdateRes, common.ErrorWithAttachment) {
	out := new(VideoCollectionUpdateRes)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Update", in, out)
}

func (c *videoCollectionClient) Upsert(ctx context.Context, in *VideoCollectionUpsertReq, opts ...grpc_go.CallOption) (*VideoCollectionUpsertRes, common.ErrorWithAttachment) {
	out := new(VideoCollectionUpsertRes)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Upsert", in, out)
}

func (c *videoCollectionClient) Delete(ctx context.Context, in *VideoCollectionDeleteReq, opts ...grpc_go.CallOption) (*VideoCollectionDeleteRes, common.ErrorWithAttachment) {
	out := new(VideoCollectionDeleteRes)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Delete", in, out)
}

// VideoCollectionServer is the server API for VideoCollection service.
// All implementations must embed UnimplementedVideoCollectionServer
// for forward compatibility
type VideoCollectionServer interface {
	Count(context.Context, *VideoCollectionCountReq) (*VideoCollectionCountRes, error)
	One(context.Context, *VideoCollectionOneReq) (*VideoCollectionOneRes, error)
	List(context.Context, *VideoCollectionListReq) (*VideoCollectionListRes, error)
	Create(context.Context, *VideoCollectionCreateReq) (*VideoCollectionCreateRes, error)
	Update(context.Context, *VideoCollectionUpdateReq) (*VideoCollectionUpdateRes, error)
	Upsert(context.Context, *VideoCollectionUpsertReq) (*VideoCollectionUpsertRes, error)
	Delete(context.Context, *VideoCollectionDeleteReq) (*VideoCollectionDeleteRes, error)
	mustEmbedUnimplementedVideoCollectionServer()
}

// UnimplementedVideoCollectionServer must be embedded to have forward compatible implementations.
type UnimplementedVideoCollectionServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedVideoCollectionServer) Count(context.Context, *VideoCollectionCountReq) (*VideoCollectionCountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedVideoCollectionServer) One(context.Context, *VideoCollectionOneReq) (*VideoCollectionOneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method One not implemented")
}
func (UnimplementedVideoCollectionServer) List(context.Context, *VideoCollectionListReq) (*VideoCollectionListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedVideoCollectionServer) Create(context.Context, *VideoCollectionCreateReq) (*VideoCollectionCreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVideoCollectionServer) Update(context.Context, *VideoCollectionUpdateReq) (*VideoCollectionUpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVideoCollectionServer) Upsert(context.Context, *VideoCollectionUpsertReq) (*VideoCollectionUpsertRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedVideoCollectionServer) Delete(context.Context, *VideoCollectionDeleteReq) (*VideoCollectionDeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *UnimplementedVideoCollectionServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedVideoCollectionServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedVideoCollectionServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &VideoCollection_ServiceDesc
}
func (s *UnimplementedVideoCollectionServer) XXX_InterfaceName() string {
	return "proto.video_collection.VideoCollection"
}

func (UnimplementedVideoCollectionServer) mustEmbedUnimplementedVideoCollectionServer() {}

// UnsafeVideoCollectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoCollectionServer will
// result in compilation errors.
type UnsafeVideoCollectionServer interface {
	mustEmbedUnimplementedVideoCollectionServer()
}

func RegisterVideoCollectionServer(s grpc_go.ServiceRegistrar, srv VideoCollectionServer) {
	s.RegisterService(&VideoCollection_ServiceDesc, srv)
}

func _VideoCollection_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCollectionCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Count", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCollection_One_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCollectionOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("One", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCollection_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCollectionListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("List", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCollection_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCollectionCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Create", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCollection_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCollectionUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Update", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCollection_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCollectionUpsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Upsert", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCollection_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCollectionDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Delete", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// VideoCollection_ServiceDesc is the grpc_go.ServiceDesc for VideoCollection service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoCollection_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "proto.video_collection.VideoCollection",
	HandlerType: (*VideoCollectionServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Count",
			Handler:    _VideoCollection_Count_Handler,
		},
		{
			MethodName: "One",
			Handler:    _VideoCollection_One_Handler,
		},
		{
			MethodName: "List",
			Handler:    _VideoCollection_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _VideoCollection_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VideoCollection_Update_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _VideoCollection_Upsert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VideoCollection_Delete_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "proto/video_collection/video_collection.proto",
}