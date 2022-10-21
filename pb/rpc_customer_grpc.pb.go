// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: rpc_customer.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RPCCustomerClient is the client API for RPCCustomer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCCustomerClient interface {
	FindById(ctx context.Context, in *CustomerParamId, opts ...grpc.CallOption) (*Customer, error)
	CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	UpdateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	BookingHistory(ctx context.Context, in *BookingHistoryRequest, opts ...grpc.CallOption) (*BookingHistoryResponse, error)
}

type rPCCustomerClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCCustomerClient(cc grpc.ClientConnInterface) RPCCustomerClient {
	return &rPCCustomerClient{cc}
}

func (c *rPCCustomerClient) FindById(ctx context.Context, in *CustomerParamId, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCCustomer/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCCustomerClient) CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCCustomer/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCCustomerClient) UpdateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCCustomer/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCCustomerClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCCustomer/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCCustomerClient) BookingHistory(ctx context.Context, in *BookingHistoryRequest, opts ...grpc.CallOption) (*BookingHistoryResponse, error) {
	out := new(BookingHistoryResponse)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCCustomer/BookingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCCustomerServer is the server API for RPCCustomer service.
// All implementations must embed UnimplementedRPCCustomerServer
// for forward compatibility
type RPCCustomerServer interface {
	FindById(context.Context, *CustomerParamId) (*Customer, error)
	CreateCustomer(context.Context, *Customer) (*Customer, error)
	UpdateCustomer(context.Context, *Customer) (*Customer, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	BookingHistory(context.Context, *BookingHistoryRequest) (*BookingHistoryResponse, error)
	mustEmbedUnimplementedRPCCustomerServer()
}

// UnimplementedRPCCustomerServer must be embedded to have forward compatible implementations.
type UnimplementedRPCCustomerServer struct {
}

func (UnimplementedRPCCustomerServer) FindById(context.Context, *CustomerParamId) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedRPCCustomerServer) CreateCustomer(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedRPCCustomerServer) UpdateCustomer(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedRPCCustomerServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedRPCCustomerServer) BookingHistory(context.Context, *BookingHistoryRequest) (*BookingHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookingHistory not implemented")
}
func (UnimplementedRPCCustomerServer) mustEmbedUnimplementedRPCCustomerServer() {}

// UnsafeRPCCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCCustomerServer will
// result in compilation errors.
type UnsafeRPCCustomerServer interface {
	mustEmbedUnimplementedRPCCustomerServer()
}

func RegisterRPCCustomerServer(s grpc.ServiceRegistrar, srv RPCCustomerServer) {
	s.RegisterService(&RPCCustomer_ServiceDesc, srv)
}

func _RPCCustomer_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerParamId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCustomerServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCCustomer/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCustomerServer).FindById(ctx, req.(*CustomerParamId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCCustomer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCCustomer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCustomerServer).CreateCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCCustomer_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCustomerServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCCustomer/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCustomerServer).UpdateCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCCustomer_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCustomerServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCCustomer/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCustomerServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCCustomer_BookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCCustomerServer).BookingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCCustomer/BookingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCCustomerServer).BookingHistory(ctx, req.(*BookingHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCCustomer_ServiceDesc is the grpc.ServiceDesc for RPCCustomer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCCustomer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tuns_go_flight.RPCCustomer",
	HandlerType: (*RPCCustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindById",
			Handler:    _RPCCustomer_FindById_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _RPCCustomer_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _RPCCustomer_UpdateCustomer_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _RPCCustomer_ChangePassword_Handler,
		},
		{
			MethodName: "BookingHistory",
			Handler:    _RPCCustomer_BookingHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_customer.proto",
}
