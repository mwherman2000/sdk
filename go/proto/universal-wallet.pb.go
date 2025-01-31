// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: services/universal-wallet/v1/universal-wallet.proto

package sdk

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

// Search request object
type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query             string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	ContinuationToken string `protobuf:"bytes,2,opt,name=continuation_token,json=continuationToken,proto3" json:"continuation_token,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetContinuationToken() string {
	if x != nil {
		return x.ContinuationToken
	}
	return ""
}

// Search response object
type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items             []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	HasMore           bool     `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Count             int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	ContinuationToken string   `protobuf:"bytes,4,opt,name=continuation_token,json=continuationToken,proto3" json:"continuation_token,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SearchResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *SearchResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SearchResponse) GetContinuationToken() string {
	if x != nil {
		return x.ContinuationToken
	}
	return ""
}

// Get item request object
type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The item identifier
	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *GetItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

// Get item response object
type GetItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The item data represented as stringified JSON
	ItemJson string `protobuf:"bytes,1,opt,name=item_json,json=itemJson,proto3" json:"item_json,omitempty"`
	// User set item type that described the content of this item
	ItemType string `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemResponse) GetItemJson() string {
	if x != nil {
		return x.ItemJson
	}
	return ""
}

func (x *GetItemResponse) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

// Update item request object
type UpdateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The item identifier
	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	// The item type that described the content of this item
	ItemType string `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
}

func (x *UpdateItemRequest) Reset() {
	*x = UpdateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemRequest) ProtoMessage() {}

func (x *UpdateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *UpdateItemRequest) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

// Update item response object
type UpdateItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response status
	Status ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=services.common.v1.ResponseStatus" json:"status,omitempty"`
}

func (x *UpdateItemResponse) Reset() {
	*x = UpdateItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemResponse) ProtoMessage() {}

func (x *UpdateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateItemResponse) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateItemResponse) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_SUCCESS
}

// Insert item request
type InsertItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the document to insert as stringified json
	ItemJson string `protobuf:"bytes,1,opt,name=item_json,json=itemJson,proto3" json:"item_json,omitempty"`
	// optional item type ex. "VerifiableCredential"
	ItemType string `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
}

func (x *InsertItemRequest) Reset() {
	*x = InsertItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertItemRequest) ProtoMessage() {}

func (x *InsertItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertItemRequest.ProtoReflect.Descriptor instead.
func (*InsertItemRequest) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *InsertItemRequest) GetItemJson() string {
	if x != nil {
		return x.ItemJson
	}
	return ""
}

func (x *InsertItemRequest) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

// Insert item response
type InsertItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=services.common.v1.ResponseStatus" json:"status,omitempty"`
	// The item identifier of the inserted record
	ItemId string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *InsertItemResponse) Reset() {
	*x = InsertItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertItemResponse) ProtoMessage() {}

func (x *InsertItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertItemResponse.ProtoReflect.Descriptor instead.
func (*InsertItemResponse) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *InsertItemResponse) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_SUCCESS
}

func (x *InsertItemResponse) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

// Delete item request
type DeleteItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// item identifier of the record to delete
	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

// Delete item response
type DeleteItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=services.common.v1.ResponseStatus" json:"status,omitempty"`
}

func (x *DeleteItemResponse) Reset() {
	*x = DeleteItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemResponse) ProtoMessage() {}

func (x *DeleteItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteItemResponse) Descriptor() ([]byte, []int) {
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteItemResponse) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_SUCCESS
}

var File_services_universal_wallet_v1_universal_wallet_proto protoreflect.FileDescriptor

var file_services_universal_wallet_v1_universal_wallet_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x6c, 0x2d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x2d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x4b, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x49, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4d, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x69, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22,
	0x50, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xa7, 0x04, 0x0a, 0x0f, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x64, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d,
	0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x66, 0x0a, 0x23, 0x74,
	0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x2d, 0x69, 0x64, 0x2f, 0x73, 0x64, 0x6b, 0xaa, 0x02, 0x23,
	0x54, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_universal_wallet_v1_universal_wallet_proto_rawDescOnce sync.Once
	file_services_universal_wallet_v1_universal_wallet_proto_rawDescData = file_services_universal_wallet_v1_universal_wallet_proto_rawDesc
)

func file_services_universal_wallet_v1_universal_wallet_proto_rawDescGZIP() []byte {
	file_services_universal_wallet_v1_universal_wallet_proto_rawDescOnce.Do(func() {
		file_services_universal_wallet_v1_universal_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_universal_wallet_v1_universal_wallet_proto_rawDescData)
	})
	return file_services_universal_wallet_v1_universal_wallet_proto_rawDescData
}

var file_services_universal_wallet_v1_universal_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_universal_wallet_v1_universal_wallet_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),      // 0: services.universalwallet.v1.SearchRequest
	(*SearchResponse)(nil),     // 1: services.universalwallet.v1.SearchResponse
	(*GetItemRequest)(nil),     // 2: services.universalwallet.v1.GetItemRequest
	(*GetItemResponse)(nil),    // 3: services.universalwallet.v1.GetItemResponse
	(*UpdateItemRequest)(nil),  // 4: services.universalwallet.v1.UpdateItemRequest
	(*UpdateItemResponse)(nil), // 5: services.universalwallet.v1.UpdateItemResponse
	(*InsertItemRequest)(nil),  // 6: services.universalwallet.v1.InsertItemRequest
	(*InsertItemResponse)(nil), // 7: services.universalwallet.v1.InsertItemResponse
	(*DeleteItemRequest)(nil),  // 8: services.universalwallet.v1.DeleteItemRequest
	(*DeleteItemResponse)(nil), // 9: services.universalwallet.v1.DeleteItemResponse
	(ResponseStatus)(0),        // 10: services.common.v1.ResponseStatus
}
var file_services_universal_wallet_v1_universal_wallet_proto_depIdxs = []int32{
	10, // 0: services.universalwallet.v1.UpdateItemResponse.status:type_name -> services.common.v1.ResponseStatus
	10, // 1: services.universalwallet.v1.InsertItemResponse.status:type_name -> services.common.v1.ResponseStatus
	10, // 2: services.universalwallet.v1.DeleteItemResponse.status:type_name -> services.common.v1.ResponseStatus
	2,  // 3: services.universalwallet.v1.UniversalWallet.GetItem:input_type -> services.universalwallet.v1.GetItemRequest
	0,  // 4: services.universalwallet.v1.UniversalWallet.Search:input_type -> services.universalwallet.v1.SearchRequest
	6,  // 5: services.universalwallet.v1.UniversalWallet.InsertItem:input_type -> services.universalwallet.v1.InsertItemRequest
	4,  // 6: services.universalwallet.v1.UniversalWallet.UpdateItem:input_type -> services.universalwallet.v1.UpdateItemRequest
	8,  // 7: services.universalwallet.v1.UniversalWallet.DeleteItem:input_type -> services.universalwallet.v1.DeleteItemRequest
	3,  // 8: services.universalwallet.v1.UniversalWallet.GetItem:output_type -> services.universalwallet.v1.GetItemResponse
	1,  // 9: services.universalwallet.v1.UniversalWallet.Search:output_type -> services.universalwallet.v1.SearchResponse
	7,  // 10: services.universalwallet.v1.UniversalWallet.InsertItem:output_type -> services.universalwallet.v1.InsertItemResponse
	5,  // 11: services.universalwallet.v1.UniversalWallet.UpdateItem:output_type -> services.universalwallet.v1.UpdateItemResponse
	9,  // 12: services.universalwallet.v1.UniversalWallet.DeleteItem:output_type -> services.universalwallet.v1.DeleteItemResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_services_universal_wallet_v1_universal_wallet_proto_init() }
func file_services_universal_wallet_v1_universal_wallet_proto_init() {
	if File_services_universal_wallet_v1_universal_wallet_proto != nil {
		return
	}
	file_services_common_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemResponse); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemRequest); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemResponse); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertItemRequest); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertItemResponse); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemRequest); i {
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
		file_services_universal_wallet_v1_universal_wallet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemResponse); i {
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
			RawDescriptor: file_services_universal_wallet_v1_universal_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_universal_wallet_v1_universal_wallet_proto_goTypes,
		DependencyIndexes: file_services_universal_wallet_v1_universal_wallet_proto_depIdxs,
		MessageInfos:      file_services_universal_wallet_v1_universal_wallet_proto_msgTypes,
	}.Build()
	File_services_universal_wallet_v1_universal_wallet_proto = out.File
	file_services_universal_wallet_v1_universal_wallet_proto_rawDesc = nil
	file_services_universal_wallet_v1_universal_wallet_proto_goTypes = nil
	file_services_universal_wallet_v1_universal_wallet_proto_depIdxs = nil
}
