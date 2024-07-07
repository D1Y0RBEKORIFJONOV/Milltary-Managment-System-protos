// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: storehouses/storehouses.proto

package storehouses1

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

type Storehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Amount        int32   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	TypeArtillery string  `protobuf:"bytes,5,opt,name=type_artillery,json=typeArtillery,proto3" json:"type_artillery,omitempty"`
	CreatedAt     string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     string  `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Storehouse) Reset() {
	*x = Storehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storehouse) ProtoMessage() {}

func (x *Storehouse) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storehouse.ProtoReflect.Descriptor instead.
func (*Storehouse) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{0}
}

func (x *Storehouse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Storehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Storehouse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Storehouse) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Storehouse) GetTypeArtillery() string {
	if x != nil {
		return x.TypeArtillery
	}
	return ""
}

func (x *Storehouse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Storehouse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Storehouse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CreateStorehouseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price         float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Amount        int32   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	TypeArtillery string  `protobuf:"bytes,4,opt,name=type_artillery,json=typeArtillery,proto3" json:"type_artillery,omitempty"`
}

func (x *CreateStorehouseReq) Reset() {
	*x = CreateStorehouseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStorehouseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStorehouseReq) ProtoMessage() {}

func (x *CreateStorehouseReq) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStorehouseReq.ProtoReflect.Descriptor instead.
func (*CreateStorehouseReq) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStorehouseReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStorehouseReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateStorehouseReq) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateStorehouseReq) GetTypeArtillery() string {
	if x != nil {
		return x.TypeArtillery
	}
	return ""
}

type UpdateStorehouseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price         float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Amount        int32   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	TypeArtillery string  `protobuf:"bytes,5,opt,name=type_artillery,json=typeArtillery,proto3" json:"type_artillery,omitempty"`
}

func (x *UpdateStorehouseReq) Reset() {
	*x = UpdateStorehouseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStorehouseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStorehouseReq) ProtoMessage() {}

func (x *UpdateStorehouseReq) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStorehouseReq.ProtoReflect.Descriptor instead.
func (*UpdateStorehouseReq) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateStorehouseReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStorehouseReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStorehouseReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateStorehouseReq) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateStorehouseReq) GetTypeArtillery() string {
	if x != nil {
		return x.TypeArtillery
	}
	return ""
}

type GetStorehouseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetStorehouseReq) Reset() {
	*x = GetStorehouseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorehouseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorehouseReq) ProtoMessage() {}

func (x *GetStorehouseReq) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorehouseReq.ProtoReflect.Descriptor instead.
func (*GetStorehouseReq) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{3}
}

func (x *GetStorehouseReq) GetFields() string {
	if x != nil {
		return x.Fields
	}
	return ""
}

func (x *GetStorehouseReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetAllStorehouseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Page   int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllStorehouseReq) Reset() {
	*x = GetAllStorehouseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStorehouseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStorehouseReq) ProtoMessage() {}

func (x *GetAllStorehouseReq) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStorehouseReq.ProtoReflect.Descriptor instead.
func (*GetAllStorehouseReq) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllStorehouseReq) GetFields() string {
	if x != nil {
		return x.Fields
	}
	return ""
}

func (x *GetAllStorehouseReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetAllStorehouseReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllStorehouseReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllStorehouseRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storehouses []*Storehouse `protobuf:"bytes,1,rep,name=storehouses,proto3" json:"storehouses,omitempty"`
	Count       int64         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllStorehouseRes) Reset() {
	*x = GetAllStorehouseRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStorehouseRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStorehouseRes) ProtoMessage() {}

func (x *GetAllStorehouseRes) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStorehouseRes.ProtoReflect.Descriptor instead.
func (*GetAllStorehouseRes) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllStorehouseRes) GetStorehouses() []*Storehouse {
	if x != nil {
		return x.Storehouses
	}
	return nil
}

func (x *GetAllStorehouseRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteStorehouseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStorehouseReq) Reset() {
	*x = DeleteStorehouseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStorehouseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStorehouseReq) ProtoMessage() {}

func (x *DeleteStorehouseReq) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStorehouseReq.ProtoReflect.Descriptor instead.
func (*DeleteStorehouseReq) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteStorehouseReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storehouses_storehouses_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_storehouses_storehouses_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_storehouses_storehouses_proto_rawDescGZIP(), []int{7}
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_storehouses_storehouses_proto protoreflect.FileDescriptor

var file_storehouses_storehouses_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe2, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x7e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x41, 0x72, 0x74, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x22, 0x8e, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a,
	0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52,
	0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa5, 0x02,
	0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x11, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12,
	0x14, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12,
	0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x61, 0x62, 0x64, 0x75, 0x6c, 0x6c, 0x6f,
	0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storehouses_storehouses_proto_rawDescOnce sync.Once
	file_storehouses_storehouses_proto_rawDescData = file_storehouses_storehouses_proto_rawDesc
)

func file_storehouses_storehouses_proto_rawDescGZIP() []byte {
	file_storehouses_storehouses_proto_rawDescOnce.Do(func() {
		file_storehouses_storehouses_proto_rawDescData = protoimpl.X.CompressGZIP(file_storehouses_storehouses_proto_rawDescData)
	})
	return file_storehouses_storehouses_proto_rawDescData
}

var file_storehouses_storehouses_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_storehouses_storehouses_proto_goTypes = []interface{}{
	(*Storehouse)(nil),          // 0: Storehouse
	(*CreateStorehouseReq)(nil), // 1: CreateStorehouseReq
	(*UpdateStorehouseReq)(nil), // 2: UpdateStorehouseReq
	(*GetStorehouseReq)(nil),    // 3: GetStorehouseReq
	(*GetAllStorehouseReq)(nil), // 4: GetAllStorehouseReq
	(*GetAllStorehouseRes)(nil), // 5: GetAllStorehouseRes
	(*DeleteStorehouseReq)(nil), // 6: DeleteStorehouseReq
	(*Status)(nil),              // 7: Status
}
var file_storehouses_storehouses_proto_depIdxs = []int32{
	0, // 0: GetAllStorehouseRes.storehouses:type_name -> Storehouse
	1, // 1: StorehouseService.CreateStorehouse:input_type -> CreateStorehouseReq
	3, // 2: StorehouseService.GetStorehouse:input_type -> GetStorehouseReq
	4, // 3: StorehouseService.GetAllStorehouse:input_type -> GetAllStorehouseReq
	2, // 4: StorehouseService.UpdateStorehouse:input_type -> UpdateStorehouseReq
	6, // 5: StorehouseService.DeleteStorehouse:input_type -> DeleteStorehouseReq
	0, // 6: StorehouseService.CreateStorehouse:output_type -> Storehouse
	0, // 7: StorehouseService.GetStorehouse:output_type -> Storehouse
	5, // 8: StorehouseService.GetAllStorehouse:output_type -> GetAllStorehouseRes
	0, // 9: StorehouseService.UpdateStorehouse:output_type -> Storehouse
	7, // 10: StorehouseService.DeleteStorehouse:output_type -> Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storehouses_storehouses_proto_init() }
func file_storehouses_storehouses_proto_init() {
	if File_storehouses_storehouses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storehouses_storehouses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storehouse); i {
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
		file_storehouses_storehouses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStorehouseReq); i {
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
		file_storehouses_storehouses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStorehouseReq); i {
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
		file_storehouses_storehouses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorehouseReq); i {
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
		file_storehouses_storehouses_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStorehouseReq); i {
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
		file_storehouses_storehouses_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStorehouseRes); i {
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
		file_storehouses_storehouses_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStorehouseReq); i {
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
		file_storehouses_storehouses_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_storehouses_storehouses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storehouses_storehouses_proto_goTypes,
		DependencyIndexes: file_storehouses_storehouses_proto_depIdxs,
		MessageInfos:      file_storehouses_storehouses_proto_msgTypes,
	}.Build()
	File_storehouses_storehouses_proto = out.File
	file_storehouses_storehouses_proto_rawDesc = nil
	file_storehouses_storehouses_proto_goTypes = nil
	file_storehouses_storehouses_proto_depIdxs = nil
}
