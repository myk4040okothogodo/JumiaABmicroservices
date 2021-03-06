// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/service_A/service_A.proto

package service_A

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

// ServiceA_APIClient is the client API for ServiceA_API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceA_APIClient interface {
	GetCsvData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type serviceA_APIClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceA_APIClient(cc grpc.ClientConnInterface) ServiceA_APIClient {
	return &serviceA_APIClient{cc}
}

func (c *serviceA_APIClient) GetCsvData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/JumiaABmicroservices.grpc.service_A.v1.ServiceA_API/GetCsvData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceA_APIServer is the server API for ServiceA_API service.
// All implementations must embed UnimplementedServiceA_APIServer
// for forward compatibility
type ServiceA_APIServer interface {
	GetCsvData(context.Context, *DataRequest) (*DataResponse, error)
	//mustEmbedUnimplementedServiceA_APIServer()
}

// UnimplementedServiceA_APIServer must be embedded to have forward compatible implementations.
type UnimplementedServiceA_APIServer struct {
}

func (UnimplementedServiceA_APIServer) GetCsvData(context.Context, *DataRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCsvData not implemented")
}
//func (UnimplementedServiceA_APIServer) mustEmbedUnimplementedServiceA_APIServer() {}

// UnsafeServiceA_APIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceA_APIServer will
// result in compilation errors.
//type UnsafeServiceA_APIServer interface {
//	mustEmbedUnimplementedServiceA_APIServer()
//}

func RegisterServiceA_APIServer(s grpc.ServiceRegistrar, srv ServiceA_APIServer) {
	s.RegisterService(&ServiceA_API_ServiceDesc, srv)
}

func _ServiceA_API_GetCsvData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceA_APIServer).GetCsvData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JumiaABmicroservices.grpc.service_A.v1.ServiceA_API/GetCsvData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceA_APIServer).GetCsvData(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceA_API_ServiceDesc is the grpc.ServiceDesc for ServiceA_API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceA_API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JumiaABmicroservices.grpc.service_A.v1.ServiceA_API",
	HandlerType: (*ServiceA_APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCsvData",
			Handler:    _ServiceA_API_GetCsvData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service_A/service_A.proto",
}
