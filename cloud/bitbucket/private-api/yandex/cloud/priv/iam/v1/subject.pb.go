// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/priv/iam/v1/subject.proto

package iam

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubjectV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IAM subject id.
	SubjectId string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// Subject type (e.g. userAccount, serviceAccount, etc.).
	SubjectType string `protobuf:"bytes,2,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
	// Id of subject in external Identity Provider.
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// Id of subject's federation ("yandex" for passport users).
	FederationId string `protobuf:"bytes,4,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
	// Serialized JSON containing subject's attributes.
	Attributes string `protobuf:"bytes,5,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Serialized JSON containing subject's settings.
	Settings string `protobuf:"bytes,6,opt,name=settings,proto3" json:"settings,omitempty"`
	// When subject was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// When subject was last modified.
	ModifiedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
}

func (x *SubjectV2) Reset() {
	*x = SubjectV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_subject_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectV2) ProtoMessage() {}

func (x *SubjectV2) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_subject_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectV2.ProtoReflect.Descriptor instead.
func (*SubjectV2) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_subject_proto_rawDescGZIP(), []int{0}
}

func (x *SubjectV2) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *SubjectV2) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *SubjectV2) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *SubjectV2) GetFederationId() string {
	if x != nil {
		return x.FederationId
	}
	return ""
}

func (x *SubjectV2) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

func (x *SubjectV2) GetSettings() string {
	if x != nil {
		return x.Settings
	}
	return ""
}

func (x *SubjectV2) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SubjectV2) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

var File_yandex_cloud_priv_iam_v1_subject_proto protoreflect.FileDescriptor

var file_yandex_cloud_priv_iam_v1_subject_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x02, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56,
	0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x42, 0x4f, 0x42,
	0x02, 0x50, 0x53, 0x5a, 0x49, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_priv_iam_v1_subject_proto_rawDescOnce sync.Once
	file_yandex_cloud_priv_iam_v1_subject_proto_rawDescData = file_yandex_cloud_priv_iam_v1_subject_proto_rawDesc
)

func file_yandex_cloud_priv_iam_v1_subject_proto_rawDescGZIP() []byte {
	file_yandex_cloud_priv_iam_v1_subject_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_priv_iam_v1_subject_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_priv_iam_v1_subject_proto_rawDescData)
	})
	return file_yandex_cloud_priv_iam_v1_subject_proto_rawDescData
}

var file_yandex_cloud_priv_iam_v1_subject_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_priv_iam_v1_subject_proto_goTypes = []interface{}{
	(*SubjectV2)(nil),             // 0: yandex.cloud.priv.iam.v1.SubjectV2
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_priv_iam_v1_subject_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.priv.iam.v1.SubjectV2.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.priv.iam.v1.SubjectV2.modified_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_priv_iam_v1_subject_proto_init() }
func file_yandex_cloud_priv_iam_v1_subject_proto_init() {
	if File_yandex_cloud_priv_iam_v1_subject_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_priv_iam_v1_subject_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectV2); i {
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
			RawDescriptor: file_yandex_cloud_priv_iam_v1_subject_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_priv_iam_v1_subject_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_priv_iam_v1_subject_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_priv_iam_v1_subject_proto_msgTypes,
	}.Build()
	File_yandex_cloud_priv_iam_v1_subject_proto = out.File
	file_yandex_cloud_priv_iam_v1_subject_proto_rawDesc = nil
	file_yandex_cloud_priv_iam_v1_subject_proto_goTypes = nil
	file_yandex_cloud_priv_iam_v1_subject_proto_depIdxs = nil
}