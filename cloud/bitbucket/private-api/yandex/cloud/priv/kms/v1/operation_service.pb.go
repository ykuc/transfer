// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/priv/kms/v1/operation_service.proto

package kms

import (
	_ "github.com/doublecloud/tross/cloud/bitbucket/private-api/yandex/cloud/priv"
	operation "github.com/doublecloud/tross/cloud/bitbucket/private-api/yandex/cloud/priv/operation"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *GetOperationRequest) Reset() {
	*x = GetOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_kms_v1_operation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationRequest) ProtoMessage() {}

func (x *GetOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_kms_v1_operation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationRequest.ProtoReflect.Descriptor instead.
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetOperationRequest) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

var File_yandex_cloud_priv_kms_v1_operation_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x18, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2b, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xca, 0x89, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x32, 0x70, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x50, 0x42, 0x03, 0x50, 0x4f, 0x53, 0x5a, 0x49, 0x61, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x6b, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescData = file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDesc
)

func file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescData)
	})
	return file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDescData
}

var file_yandex_cloud_priv_kms_v1_operation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_priv_kms_v1_operation_service_proto_goTypes = []interface{}{
	(*GetOperationRequest)(nil), // 0: yandex.cloud.priv.kms.v1.GetOperationRequest
	(*operation.Operation)(nil), // 1: yandex.cloud.priv.operation.Operation
}
var file_yandex_cloud_priv_kms_v1_operation_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.priv.kms.v1.OperationService.Get:input_type -> yandex.cloud.priv.kms.v1.GetOperationRequest
	1, // 1: yandex.cloud.priv.kms.v1.OperationService.Get:output_type -> yandex.cloud.priv.operation.Operation
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_priv_kms_v1_operation_service_proto_init() }
func file_yandex_cloud_priv_kms_v1_operation_service_proto_init() {
	if File_yandex_cloud_priv_kms_v1_operation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_priv_kms_v1_operation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOperationRequest); i {
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
			RawDescriptor: file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_priv_kms_v1_operation_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_priv_kms_v1_operation_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_priv_kms_v1_operation_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_priv_kms_v1_operation_service_proto = out.File
	file_yandex_cloud_priv_kms_v1_operation_service_proto_rawDesc = nil
	file_yandex_cloud_priv_kms_v1_operation_service_proto_goTypes = nil
	file_yandex_cloud_priv_kms_v1_operation_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperationServiceClient is the client API for OperationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationServiceClient interface {
	Get(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type operationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationServiceClient(cc grpc.ClientConnInterface) OperationServiceClient {
	return &operationServiceClient{cc}
}

func (c *operationServiceClient) Get(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.priv.kms.v1.OperationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServiceServer is the server API for OperationService service.
type OperationServiceServer interface {
	Get(context.Context, *GetOperationRequest) (*operation.Operation, error)
}

// UnimplementedOperationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperationServiceServer struct {
}

func (*UnimplementedOperationServiceServer) Get(context.Context, *GetOperationRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterOperationServiceServer(s *grpc.Server, srv OperationServiceServer) {
	s.RegisterService(&_OperationService_serviceDesc, srv)
}

func _OperationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.priv.kms.v1.OperationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Get(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.priv.kms.v1.OperationService",
	HandlerType: (*OperationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OperationService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/priv/kms/v1/operation_service.proto",
}