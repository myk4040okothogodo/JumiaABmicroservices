// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/service_B/service_B.proto

package service_B

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

// ServiceB_APIClient is the client API for ServiceB_API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceB_APIClient interface {
	// get order using id
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	// Get Order by email
	GetOrderByEmail(ctx context.Context, in *GetOrderByEmailRequest, opts ...grpc.CallOption) (*GetOrderByEmailResponse, error)
	// Get orders by weight
	GetOrdersByWeight(ctx context.Context, in *GetOrdersByWeightRequest, opts ...grpc.CallOption) (*GetOrdersByWeightResponse, error)
	// Get all orders
	GetAllOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	// Get all Country Orders
	GetAllCountryOrders(ctx context.Context, in *GetCountryOrdersRequest, opts ...grpc.CallOption) (*GetCountryOrdersResponse, error)
	// Get totalweight of Orders per Country
	GetWeightofAllOrdersPerCountry(ctx context.Context, in *GetCountryOrdersWeightRequest, opts ...grpc.CallOption) (*GetCountryOrdersWeightResponse, error)
	// Get Orders as per date
	GetOrdersAsPerDate(ctx context.Context, in *GetOrdersasPerDateRequest, opts ...grpc.CallOption) (*GetOrdersasPerDateResponse, error)
}

type serviceB_APIClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceB_APIClient(cc grpc.ClientConnInterface) ServiceB_APIClient {
	return &serviceB_APIClient{cc}
}

func (c *serviceB_APIClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceB_APIClient) GetOrderByEmail(ctx context.Context, in *GetOrderByEmailRequest, opts ...grpc.CallOption) (*GetOrderByEmailResponse, error) {
	out := new(GetOrderByEmailResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrderByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceB_APIClient) GetOrdersByWeight(ctx context.Context, in *GetOrdersByWeightRequest, opts ...grpc.CallOption) (*GetOrdersByWeightResponse, error) {
	out := new(GetOrdersByWeightResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrdersByWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceB_APIClient) GetAllOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetAllOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceB_APIClient) GetAllCountryOrders(ctx context.Context, in *GetCountryOrdersRequest, opts ...grpc.CallOption) (*GetCountryOrdersResponse, error) {
	out := new(GetCountryOrdersResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetAllCountryOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceB_APIClient) GetWeightofAllOrdersPerCountry(ctx context.Context, in *GetCountryOrdersWeightRequest, opts ...grpc.CallOption) (*GetCountryOrdersWeightResponse, error) {
	out := new(GetCountryOrdersWeightResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetWeightofAllOrdersPerCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceB_APIClient) GetOrdersAsPerDate(ctx context.Context, in *GetOrdersasPerDateRequest, opts ...grpc.CallOption) (*GetOrdersasPerDateResponse, error) {
	out := new(GetOrdersasPerDateResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrdersAsPerDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceB_APIServer is the server API for ServiceB_API service.
// All implementations must embed UnimplementedServiceB_APIServer
// for forward compatibility
type ServiceB_APIServer interface {
	// get order using id
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	// Get Order by email
	GetOrderByEmail(context.Context, *GetOrderByEmailRequest) (*GetOrderByEmailResponse, error)
	// Get orders by weight
	GetOrdersByWeight(context.Context, *GetOrdersByWeightRequest) (*GetOrdersByWeightResponse, error)
	// Get all orders
	GetAllOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	// Get all Country Orders
	GetAllCountryOrders(context.Context, *GetCountryOrdersRequest) (*GetCountryOrdersResponse, error)
	// Get totalweight of Orders per Country
	GetWeightofAllOrdersPerCountry(context.Context, *GetCountryOrdersWeightRequest) (*GetCountryOrdersWeightResponse, error)
	// Get Orders as per date
	GetOrdersAsPerDate(context.Context, *GetOrdersasPerDateRequest) (*GetOrdersasPerDateResponse, error)
	//mustEmbedUnimplementedServiceB_APIServer()
}

// UnimplementedServiceB_APIServer must be embedded to have forward compatible implementations.
type UnimplementedServiceB_APIServer struct {
}

func (UnimplementedServiceB_APIServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedServiceB_APIServer) GetOrderByEmail(context.Context, *GetOrderByEmailRequest) (*GetOrderByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByEmail not implemented")
}
func (UnimplementedServiceB_APIServer) GetOrdersByWeight(context.Context, *GetOrdersByWeightRequest) (*GetOrdersByWeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByWeight not implemented")
}
func (UnimplementedServiceB_APIServer) GetAllOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
func (UnimplementedServiceB_APIServer) GetAllCountryOrders(context.Context, *GetCountryOrdersRequest) (*GetCountryOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCountryOrders not implemented")
}
func (UnimplementedServiceB_APIServer) GetWeightofAllOrdersPerCountry(context.Context, *GetCountryOrdersWeightRequest) (*GetCountryOrdersWeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeightofAllOrdersPerCountry not implemented")
}
func (UnimplementedServiceB_APIServer) GetOrdersAsPerDate(context.Context, *GetOrdersasPerDateRequest) (*GetOrdersasPerDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersAsPerDate not implemented")
}
//func (UnimplementedServiceB_APIServer) mustEmbedUnimplementedServiceB_APIServer() {}

// UnsafeServiceB_APIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceB_APIServer will
// result in compilation errors.
//type UnsafeServiceB_APIServer interface {
//	mustEmbedUnimplementedServiceB_APIServer()
//}

func RegisterServiceB_APIServer(s grpc.ServiceRegistrar, srv ServiceB_APIServer) {
	s.RegisterService(&ServiceB_API_ServiceDesc, srv)
}

func _ServiceB_API_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceB_APIServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceB_APIServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceB_API_GetOrderByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceB_APIServer).GetOrderByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrderByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceB_APIServer).GetOrderByEmail(ctx, req.(*GetOrderByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceB_API_GetOrdersByWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersByWeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceB_APIServer).GetOrdersByWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrdersByWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceB_APIServer).GetOrdersByWeight(ctx, req.(*GetOrdersByWeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceB_API_GetAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceB_APIServer).GetAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetAllOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceB_APIServer).GetAllOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceB_API_GetAllCountryOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountryOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceB_APIServer).GetAllCountryOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetAllCountryOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceB_APIServer).GetAllCountryOrders(ctx, req.(*GetCountryOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceB_API_GetWeightofAllOrdersPerCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountryOrdersWeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceB_APIServer).GetWeightofAllOrdersPerCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetWeightofAllOrdersPerCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceB_APIServer).GetWeightofAllOrdersPerCountry(ctx, req.(*GetCountryOrdersWeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceB_API_GetOrdersAsPerDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersasPerDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceB_APIServer).GetOrdersAsPerDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_B.v1.ServiceB_API/GetOrdersAsPerDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceB_APIServer).GetOrdersAsPerDate(ctx, req.(*GetOrdersasPerDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceB_API_ServiceDesc is the grpc.ServiceDesc for ServiceB_API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceB_API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JumiaABmicroservices.grpc.service_B.v1.ServiceB_API",
	HandlerType: (*ServiceB_APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _ServiceB_API_GetOrder_Handler,
		},
		{
			MethodName: "GetOrderByEmail",
			Handler:    _ServiceB_API_GetOrderByEmail_Handler,
		},
		{
			MethodName: "GetOrdersByWeight",
			Handler:    _ServiceB_API_GetOrdersByWeight_Handler,
		},
		{
			MethodName: "GetAllOrders",
			Handler:    _ServiceB_API_GetAllOrders_Handler,
		},
		{
			MethodName: "GetAllCountryOrders",
			Handler:    _ServiceB_API_GetAllCountryOrders_Handler,
		},
		{
			MethodName: "GetWeightofAllOrdersPerCountry",
			Handler:    _ServiceB_API_GetWeightofAllOrdersPerCountry_Handler,
		},
		{
			MethodName: "GetOrdersAsPerDate",
			Handler:    _ServiceB_API_GetOrdersAsPerDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service_B/service_B.proto",
}
