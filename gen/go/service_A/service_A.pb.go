// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: proto/service_A/service_A.proto

package service_A

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_A_service_A_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_A_service_A_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_A_service_A_proto_rawDescGZIP(), []int{0}
}

func (x *DataRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdersA []*Order_A `protobuf:"bytes,1,rep,name=orders_a,json=ordersA,proto3" json:"orders_a,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_A_service_A_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_A_service_A_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_A_service_A_proto_rawDescGZIP(), []int{1}
}

func (x *DataResponse) GetOrdersA() []*Order_A {
	if x != nil {
		return x.OrdersA
	}
	return nil
}

var File_proto_service_A_service_A_proto protoreflect.FileDescriptor

var file_proto_service_A_service_A_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x41, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x26, 0x4a, 0x75, 0x6d, 0x69, 0x61, 0x41, 0x42, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x41, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x41, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x41, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5a, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x4a, 0x75, 0x6d,
	0x69, 0x61, 0x41, 0x42, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x41,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x41, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x41, 0x32, 0x87, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x5f, 0x41, 0x50, 0x49, 0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x73, 0x76, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x33, 0x2e, 0x4a, 0x75, 0x6d, 0x69, 0x61, 0x41, 0x42, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x41, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x4a, 0x75, 0x6d, 0x69, 0x61,
	0x41, 0x42, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x41, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14,
	0x5a, 0x12, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_A_service_A_proto_rawDescOnce sync.Once
	file_proto_service_A_service_A_proto_rawDescData = file_proto_service_A_service_A_proto_rawDesc
)

func file_proto_service_A_service_A_proto_rawDescGZIP() []byte {
	file_proto_service_A_service_A_proto_rawDescOnce.Do(func() {
		file_proto_service_A_service_A_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_A_service_A_proto_rawDescData)
	})
	return file_proto_service_A_service_A_proto_rawDescData
}

var file_proto_service_A_service_A_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_service_A_service_A_proto_goTypes = []interface{}{
	(*DataRequest)(nil),  // 0: JumiaABmicroservices.grpc.service_A.v1.DataRequest
	(*DataResponse)(nil), // 1: JumiaABmicroservices.grpc.service_A.v1.DataResponse
	(*Order_A)(nil),      // 2: JumiaABmicroservices.grpc.service_A.v1.Order_A
}
var file_proto_service_A_service_A_proto_depIdxs = []int32{
	2, // 0: JumiaABmicroservices.grpc.service_A.v1.DataResponse.orders_a:type_name -> JumiaABmicroservices.grpc.service_A.v1.Order_A
	0, // 1: JumiaABmicroservices.grpc.service_A.v1.ServiceA_API.GetCsvData:input_type -> JumiaABmicroservices.grpc.service_A.v1.DataRequest
	1, // 2: JumiaABmicroservices.grpc.service_A.v1.ServiceA_API.GetCsvData:output_type -> JumiaABmicroservices.grpc.service_A.v1.DataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_service_A_service_A_proto_init() }
func file_proto_service_A_service_A_proto_init() {
	if File_proto_service_A_service_A_proto != nil {
		return
	}
	file_proto_service_A_service_A_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_service_A_service_A_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_A_service_A_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_A_service_A_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_A_service_A_proto_goTypes,
		DependencyIndexes: file_proto_service_A_service_A_proto_depIdxs,
		MessageInfos:      file_proto_service_A_service_A_proto_msgTypes,
	}.Build()
	File_proto_service_A_service_A_proto = out.File
	file_proto_service_A_service_A_proto_rawDesc = nil
	file_proto_service_A_service_A_proto_goTypes = nil
	file_proto_service_A_service_A_proto_depIdxs = nil
}
