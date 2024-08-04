// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/priv/iam/restriction.proto

package iam

import (
	_ "github.com/doublecloud/tross/cloud/bitbucket/private-api/yandex/cloud/priv"
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

type RestrictionKind int32

const (
	RestrictionKind_RESTRICTION_KIND_UNSPECIFIED RestrictionKind = 0
	RestrictionKind_BLOCK_PERMISSIONS            RestrictionKind = 1
)

// Enum value maps for RestrictionKind.
var (
	RestrictionKind_name = map[int32]string{
		0: "RESTRICTION_KIND_UNSPECIFIED",
		1: "BLOCK_PERMISSIONS",
	}
	RestrictionKind_value = map[string]int32{
		"RESTRICTION_KIND_UNSPECIFIED": 0,
		"BLOCK_PERMISSIONS":            1,
	}
)

func (x RestrictionKind) Enum() *RestrictionKind {
	p := new(RestrictionKind)
	*p = x
	return p
}

func (x RestrictionKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RestrictionKind) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_priv_iam_restriction_proto_enumTypes[0].Descriptor()
}

func (RestrictionKind) Type() protoreflect.EnumType {
	return &file_yandex_cloud_priv_iam_restriction_proto_enumTypes[0]
}

func (x RestrictionKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RestrictionKind.Descriptor instead.
func (RestrictionKind) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{0}
}

type Restriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestrictionKind   RestrictionKind        `protobuf:"varint,1,opt,name=restriction_kind,json=restrictionKind,proto3,enum=yandex.cloud.priv.iam.RestrictionKind" json:"restriction_kind,omitempty"`
	RestrictionTypeId string                 `protobuf:"bytes,2,opt,name=restriction_type_id,json=restrictionTypeId,proto3" json:"restriction_type_id,omitempty"`
	AddedAt           *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
	AddedBy           string                 `protobuf:"bytes,4,opt,name=added_by,json=addedBy,proto3" json:"added_by,omitempty"`
}

func (x *Restriction) Reset() {
	*x = Restriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restriction) ProtoMessage() {}

func (x *Restriction) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restriction.ProtoReflect.Descriptor instead.
func (*Restriction) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{0}
}

func (x *Restriction) GetRestrictionKind() RestrictionKind {
	if x != nil {
		return x.RestrictionKind
	}
	return RestrictionKind_RESTRICTION_KIND_UNSPECIFIED
}

func (x *Restriction) GetRestrictionTypeId() string {
	if x != nil {
		return x.RestrictionTypeId
	}
	return ""
}

func (x *Restriction) GetAddedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AddedAt
	}
	return nil
}

func (x *Restriction) GetAddedBy() string {
	if x != nil {
		return x.AddedBy
	}
	return ""
}

type ListRestrictionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId      string          `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RestrictionKind RestrictionKind `protobuf:"varint,2,opt,name=restriction_kind,json=restrictionKind,proto3,enum=yandex.cloud.priv.iam.RestrictionKind" json:"restriction_kind,omitempty"`
	PageSize        int64           `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken       string          `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListRestrictionsRequest) Reset() {
	*x = ListRestrictionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestrictionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestrictionsRequest) ProtoMessage() {}

func (x *ListRestrictionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestrictionsRequest.ProtoReflect.Descriptor instead.
func (*ListRestrictionsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{1}
}

func (x *ListRestrictionsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ListRestrictionsRequest) GetRestrictionKind() RestrictionKind {
	if x != nil {
		return x.RestrictionKind
	}
	return RestrictionKind_RESTRICTION_KIND_UNSPECIFIED
}

func (x *ListRestrictionsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRestrictionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListRestrictionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restrictions  []*Restriction `protobuf:"bytes,1,rep,name=restrictions,proto3" json:"restrictions,omitempty"`
	NextPageToken string         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListRestrictionsResponse) Reset() {
	*x = ListRestrictionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRestrictionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRestrictionsResponse) ProtoMessage() {}

func (x *ListRestrictionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRestrictionsResponse.ProtoReflect.Descriptor instead.
func (*ListRestrictionsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{2}
}

func (x *ListRestrictionsResponse) GetRestrictions() []*Restriction {
	if x != nil {
		return x.Restrictions
	}
	return nil
}

func (x *ListRestrictionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetRestrictionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId        string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RestrictionTypeId string `protobuf:"bytes,2,opt,name=restriction_type_id,json=restrictionTypeId,proto3" json:"restriction_type_id,omitempty"`
}

func (x *GetRestrictionRequest) Reset() {
	*x = GetRestrictionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRestrictionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRestrictionRequest) ProtoMessage() {}

func (x *GetRestrictionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRestrictionRequest.ProtoReflect.Descriptor instead.
func (*GetRestrictionRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{3}
}

func (x *GetRestrictionRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *GetRestrictionRequest) GetRestrictionTypeId() string {
	if x != nil {
		return x.RestrictionTypeId
	}
	return ""
}

type AddRestrictionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId        string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RestrictionTypeId string `protobuf:"bytes,2,opt,name=restriction_type_id,json=restrictionTypeId,proto3" json:"restriction_type_id,omitempty"`
}

func (x *AddRestrictionRequest) Reset() {
	*x = AddRestrictionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRestrictionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRestrictionRequest) ProtoMessage() {}

func (x *AddRestrictionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRestrictionRequest.ProtoReflect.Descriptor instead.
func (*AddRestrictionRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{4}
}

func (x *AddRestrictionRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AddRestrictionRequest) GetRestrictionTypeId() string {
	if x != nil {
		return x.RestrictionTypeId
	}
	return ""
}

type AddRestrictionMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId        string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RestrictionTypeId string `protobuf:"bytes,2,opt,name=restriction_type_id,json=restrictionTypeId,proto3" json:"restriction_type_id,omitempty"`
	// Information about operation cancellation
	CancelledBy string                 `protobuf:"bytes,3,opt,name=cancelled_by,json=cancelledBy,proto3" json:"cancelled_by,omitempty"`
	CancelledAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
}

func (x *AddRestrictionMetadata) Reset() {
	*x = AddRestrictionMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRestrictionMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRestrictionMetadata) ProtoMessage() {}

func (x *AddRestrictionMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRestrictionMetadata.ProtoReflect.Descriptor instead.
func (*AddRestrictionMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{5}
}

func (x *AddRestrictionMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AddRestrictionMetadata) GetRestrictionTypeId() string {
	if x != nil {
		return x.RestrictionTypeId
	}
	return ""
}

func (x *AddRestrictionMetadata) GetCancelledBy() string {
	if x != nil {
		return x.CancelledBy
	}
	return ""
}

func (x *AddRestrictionMetadata) GetCancelledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CancelledAt
	}
	return nil
}

type AddRestrictionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restriction *Restriction `protobuf:"bytes,1,opt,name=restriction,proto3" json:"restriction,omitempty"`
}

func (x *AddRestrictionResponse) Reset() {
	*x = AddRestrictionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRestrictionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRestrictionResponse) ProtoMessage() {}

func (x *AddRestrictionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRestrictionResponse.ProtoReflect.Descriptor instead.
func (*AddRestrictionResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{6}
}

func (x *AddRestrictionResponse) GetRestriction() *Restriction {
	if x != nil {
		return x.Restriction
	}
	return nil
}

type RemoveRestrictionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId        string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RestrictionTypeId string `protobuf:"bytes,2,opt,name=restriction_type_id,json=restrictionTypeId,proto3" json:"restriction_type_id,omitempty"`
}

func (x *RemoveRestrictionRequest) Reset() {
	*x = RemoveRestrictionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRestrictionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRestrictionRequest) ProtoMessage() {}

func (x *RemoveRestrictionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRestrictionRequest.ProtoReflect.Descriptor instead.
func (*RemoveRestrictionRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveRestrictionRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RemoveRestrictionRequest) GetRestrictionTypeId() string {
	if x != nil {
		return x.RestrictionTypeId
	}
	return ""
}

type RemoveRestrictionMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId        string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RestrictionTypeId string `protobuf:"bytes,2,opt,name=restriction_type_id,json=restrictionTypeId,proto3" json:"restriction_type_id,omitempty"`
}

func (x *RemoveRestrictionMetadata) Reset() {
	*x = RemoveRestrictionMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRestrictionMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRestrictionMetadata) ProtoMessage() {}

func (x *RemoveRestrictionMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRestrictionMetadata.ProtoReflect.Descriptor instead.
func (*RemoveRestrictionMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveRestrictionMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RemoveRestrictionMetadata) GetRestrictionTypeId() string {
	if x != nil {
		return x.RestrictionTypeId
	}
	return ""
}

type RemoveRestrictionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional, empty if there was no restriction
	Restriction *Restriction `protobuf:"bytes,1,opt,name=restriction,proto3" json:"restriction,omitempty"`
}

func (x *RemoveRestrictionResponse) Reset() {
	*x = RemoveRestrictionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRestrictionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRestrictionResponse) ProtoMessage() {}

func (x *RemoveRestrictionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_restriction_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRestrictionResponse.ProtoReflect.Descriptor instead.
func (*RemoveRestrictionResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveRestrictionResponse) GetRestriction() *Restriction {
	if x != nil {
		return x.Restriction
	}
	return nil
}

var File_yandex_cloud_priv_iam_restriction_proto protoreflect.FileDescriptor

var file_yandex_cloud_priv_iam_restriction_proto_rawDesc = []byte{
	0x0a, 0x27, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x65, 0x64, 0x42, 0x79, 0x22, 0xef, 0x01, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xa8, 0x89, 0x31,
	0x01, 0xca, 0x89, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x51, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xba, 0x89, 0x31,
	0x06, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x29, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0x89, 0x31, 0x06, 0x3c, 0x3d, 0x32, 0x30, 0x30,
	0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8a, 0x01, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xa8, 0x89, 0x31, 0x01, 0xca, 0x89,
	0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0d, 0xa8, 0x89, 0x31, 0x01, 0xca, 0x89, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x11,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xa8, 0x89, 0x31, 0x01, 0xca, 0x89, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x13, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xa8, 0x89, 0x31, 0x01, 0xca, 0x89, 0x31,
	0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x16, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x79, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5e, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xa8, 0x89, 0x31, 0x01, 0xca,
	0x89, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0d, 0xa8, 0x89, 0x31, 0x01, 0xca, 0x89, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x22, 0x6c, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x22, 0x61, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x4a, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x01, 0x42,
	0x4d, 0x42, 0x03, 0x50, 0x52, 0x4e, 0x5a, 0x46, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x3b, 0x69, 0x61, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_priv_iam_restriction_proto_rawDescOnce sync.Once
	file_yandex_cloud_priv_iam_restriction_proto_rawDescData = file_yandex_cloud_priv_iam_restriction_proto_rawDesc
)

func file_yandex_cloud_priv_iam_restriction_proto_rawDescGZIP() []byte {
	file_yandex_cloud_priv_iam_restriction_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_priv_iam_restriction_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_priv_iam_restriction_proto_rawDescData)
	})
	return file_yandex_cloud_priv_iam_restriction_proto_rawDescData
}

var file_yandex_cloud_priv_iam_restriction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_priv_iam_restriction_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_yandex_cloud_priv_iam_restriction_proto_goTypes = []interface{}{
	(RestrictionKind)(0),              // 0: yandex.cloud.priv.iam.RestrictionKind
	(*Restriction)(nil),               // 1: yandex.cloud.priv.iam.Restriction
	(*ListRestrictionsRequest)(nil),   // 2: yandex.cloud.priv.iam.ListRestrictionsRequest
	(*ListRestrictionsResponse)(nil),  // 3: yandex.cloud.priv.iam.ListRestrictionsResponse
	(*GetRestrictionRequest)(nil),     // 4: yandex.cloud.priv.iam.GetRestrictionRequest
	(*AddRestrictionRequest)(nil),     // 5: yandex.cloud.priv.iam.AddRestrictionRequest
	(*AddRestrictionMetadata)(nil),    // 6: yandex.cloud.priv.iam.AddRestrictionMetadata
	(*AddRestrictionResponse)(nil),    // 7: yandex.cloud.priv.iam.AddRestrictionResponse
	(*RemoveRestrictionRequest)(nil),  // 8: yandex.cloud.priv.iam.RemoveRestrictionRequest
	(*RemoveRestrictionMetadata)(nil), // 9: yandex.cloud.priv.iam.RemoveRestrictionMetadata
	(*RemoveRestrictionResponse)(nil), // 10: yandex.cloud.priv.iam.RemoveRestrictionResponse
	(*timestamppb.Timestamp)(nil),     // 11: google.protobuf.Timestamp
}
var file_yandex_cloud_priv_iam_restriction_proto_depIdxs = []int32{
	0,  // 0: yandex.cloud.priv.iam.Restriction.restriction_kind:type_name -> yandex.cloud.priv.iam.RestrictionKind
	11, // 1: yandex.cloud.priv.iam.Restriction.added_at:type_name -> google.protobuf.Timestamp
	0,  // 2: yandex.cloud.priv.iam.ListRestrictionsRequest.restriction_kind:type_name -> yandex.cloud.priv.iam.RestrictionKind
	1,  // 3: yandex.cloud.priv.iam.ListRestrictionsResponse.restrictions:type_name -> yandex.cloud.priv.iam.Restriction
	11, // 4: yandex.cloud.priv.iam.AddRestrictionMetadata.cancelled_at:type_name -> google.protobuf.Timestamp
	1,  // 5: yandex.cloud.priv.iam.AddRestrictionResponse.restriction:type_name -> yandex.cloud.priv.iam.Restriction
	1,  // 6: yandex.cloud.priv.iam.RemoveRestrictionResponse.restriction:type_name -> yandex.cloud.priv.iam.Restriction
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_yandex_cloud_priv_iam_restriction_proto_init() }
func file_yandex_cloud_priv_iam_restriction_proto_init() {
	if File_yandex_cloud_priv_iam_restriction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restriction); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRestrictionsRequest); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRestrictionsResponse); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRestrictionRequest); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRestrictionRequest); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRestrictionMetadata); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRestrictionResponse); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRestrictionRequest); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRestrictionMetadata); i {
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
		file_yandex_cloud_priv_iam_restriction_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRestrictionResponse); i {
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
			RawDescriptor: file_yandex_cloud_priv_iam_restriction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_priv_iam_restriction_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_priv_iam_restriction_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_priv_iam_restriction_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_priv_iam_restriction_proto_msgTypes,
	}.Build()
	File_yandex_cloud_priv_iam_restriction_proto = out.File
	file_yandex_cloud_priv_iam_restriction_proto_rawDesc = nil
	file_yandex_cloud_priv_iam_restriction_proto_goTypes = nil
	file_yandex_cloud_priv_iam_restriction_proto_depIdxs = nil
}