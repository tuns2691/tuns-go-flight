// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: rpc_booking.proto

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

// RPCBookingClient is the client API for RPCBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCBookingClient interface {
	CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
	ViewBooking(ctx context.Context, in *ViewBookingRequest, opts ...grpc.CallOption) (*ViewBookingResponse, error)
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error)
}

type rPCBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCBookingClient(cc grpc.ClientConnInterface) RPCBookingClient {
	return &rPCBookingClient{cc}
}

func (c *rPCBookingClient) CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCBooking/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCBookingClient) ViewBooking(ctx context.Context, in *ViewBookingRequest, opts ...grpc.CallOption) (*ViewBookingResponse, error) {
	out := new(ViewBookingResponse)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCBooking/ViewBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCBookingClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error) {
	out := new(CancelBookingResponse)
	err := c.cc.Invoke(ctx, "/tuns_go_flight.RPCBooking/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCBookingServer is the server API for RPCBooking service.
// All implementations must embed UnimplementedRPCBookingServer
// for forward compatibility
type RPCBookingServer interface {
	CreateBooking(context.Context, *Booking) (*Booking, error)
	ViewBooking(context.Context, *ViewBookingRequest) (*ViewBookingResponse, error)
	CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error)
	mustEmbedUnimplementedRPCBookingServer()
}

// UnimplementedRPCBookingServer must be embedded to have forward compatible implementations.
type UnimplementedRPCBookingServer struct {
}

func (UnimplementedRPCBookingServer) CreateBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedRPCBookingServer) ViewBooking(context.Context, *ViewBookingRequest) (*ViewBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewBooking not implemented")
}
func (UnimplementedRPCBookingServer) CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedRPCBookingServer) mustEmbedUnimplementedRPCBookingServer() {}

// UnsafeRPCBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCBookingServer will
// result in compilation errors.
type UnsafeRPCBookingServer interface {
	mustEmbedUnimplementedRPCBookingServer()
}

func RegisterRPCBookingServer(s grpc.ServiceRegistrar, srv RPCBookingServer) {
	s.RegisterService(&RPCBooking_ServiceDesc, srv)
}

func _RPCBooking_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCBookingServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCBooking/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCBookingServer).CreateBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCBooking_ViewBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCBookingServer).ViewBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCBooking/ViewBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCBookingServer).ViewBooking(ctx, req.(*ViewBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCBooking_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCBookingServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tuns_go_flight.RPCBooking/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCBookingServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCBooking_ServiceDesc is the grpc.ServiceDesc for RPCBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tuns_go_flight.RPCBooking",
	HandlerType: (*RPCBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _RPCBooking_CreateBooking_Handler,
		},
		{
			MethodName: "ViewBooking",
			Handler:    _RPCBooking_ViewBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _RPCBooking_CancelBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_booking.proto",
}