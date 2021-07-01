// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: xjwt.proto

package xjwt

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

type CreateXjwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateXjwtRequest) Reset() {
	*x = CreateXjwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateXjwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateXjwtRequest) ProtoMessage() {}

func (x *CreateXjwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateXjwtRequest.ProtoReflect.Descriptor instead.
func (*CreateXjwtRequest) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{0}
}

type CreateXjwtReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateXjwtReply) Reset() {
	*x = CreateXjwtReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateXjwtReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateXjwtReply) ProtoMessage() {}

func (x *CreateXjwtReply) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateXjwtReply.ProtoReflect.Descriptor instead.
func (*CreateXjwtReply) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{1}
}

type UpdateXjwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateXjwtRequest) Reset() {
	*x = UpdateXjwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateXjwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateXjwtRequest) ProtoMessage() {}

func (x *UpdateXjwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateXjwtRequest.ProtoReflect.Descriptor instead.
func (*UpdateXjwtRequest) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{2}
}

type UpdateXjwtReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateXjwtReply) Reset() {
	*x = UpdateXjwtReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateXjwtReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateXjwtReply) ProtoMessage() {}

func (x *UpdateXjwtReply) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateXjwtReply.ProtoReflect.Descriptor instead.
func (*UpdateXjwtReply) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{3}
}

type DeleteXjwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteXjwtRequest) Reset() {
	*x = DeleteXjwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteXjwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteXjwtRequest) ProtoMessage() {}

func (x *DeleteXjwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteXjwtRequest.ProtoReflect.Descriptor instead.
func (*DeleteXjwtRequest) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{4}
}

type DeleteXjwtReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteXjwtReply) Reset() {
	*x = DeleteXjwtReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteXjwtReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteXjwtReply) ProtoMessage() {}

func (x *DeleteXjwtReply) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteXjwtReply.ProtoReflect.Descriptor instead.
func (*DeleteXjwtReply) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{5}
}

type GetXjwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetXjwtRequest) Reset() {
	*x = GetXjwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetXjwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetXjwtRequest) ProtoMessage() {}

func (x *GetXjwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetXjwtRequest.ProtoReflect.Descriptor instead.
func (*GetXjwtRequest) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{6}
}

type GetXjwtReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetXjwtReply) Reset() {
	*x = GetXjwtReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetXjwtReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetXjwtReply) ProtoMessage() {}

func (x *GetXjwtReply) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetXjwtReply.ProtoReflect.Descriptor instead.
func (*GetXjwtReply) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{7}
}

func (x *GetXjwtReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListXjwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListXjwtRequest) Reset() {
	*x = ListXjwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListXjwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListXjwtRequest) ProtoMessage() {}

func (x *ListXjwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListXjwtRequest.ProtoReflect.Descriptor instead.
func (*ListXjwtRequest) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{8}
}

type ListXjwtReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListXjwtReply) Reset() {
	*x = ListXjwtReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xjwt_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListXjwtReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListXjwtReply) ProtoMessage() {}

func (x *ListXjwtReply) ProtoReflect() protoreflect.Message {
	mi := &file_xjwt_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListXjwtReply.ProtoReflect.Descriptor instead.
func (*ListXjwtReply) Descriptor() ([]byte, []int) {
	return file_xjwt_proto_rawDescGZIP(), []int{9}
}

var File_xjwt_proto protoreflect.FileDescriptor

var file_xjwt_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x78, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6f,
	0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a, 0x77, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x58, 0x6a, 0x77, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xfd,
	0x02, 0x0a, 0x04, 0x58, 0x6a, 0x77, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x58, 0x6a, 0x77, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x78, 0x6a, 0x77, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x78, 0x6a, 0x77, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58,
	0x6a, 0x77, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a,
	0x77, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78,
	0x6a, 0x77, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x58, 0x6a, 0x77,
	0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a, 0x77, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a, 0x77,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x58, 0x6a, 0x77, 0x74, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a, 0x77, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x58,
	0x6a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x5f,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a, 0x77, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x58, 0x6a, 0x77,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x58, 0x6a,
	0x77, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a, 0x77,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x78, 0x6a, 0x77, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x58, 0x6a, 0x77, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x18,
	0x0a, 0x04, 0x78, 0x6a, 0x77, 0x74, 0x50, 0x01, 0x5a, 0x0e, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x78,
	0x6a, 0x77, 0x74, 0x3b, 0x78, 0x6a, 0x77, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xjwt_proto_rawDescOnce sync.Once
	file_xjwt_proto_rawDescData = file_xjwt_proto_rawDesc
)

func file_xjwt_proto_rawDescGZIP() []byte {
	file_xjwt_proto_rawDescOnce.Do(func() {
		file_xjwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_xjwt_proto_rawDescData)
	})
	return file_xjwt_proto_rawDescData
}

var file_xjwt_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_xjwt_proto_goTypes = []interface{}{
	(*CreateXjwtRequest)(nil), // 0: go_demo.xjwt.CreateXjwtRequest
	(*CreateXjwtReply)(nil),   // 1: go_demo.xjwt.CreateXjwtReply
	(*UpdateXjwtRequest)(nil), // 2: go_demo.xjwt.UpdateXjwtRequest
	(*UpdateXjwtReply)(nil),   // 3: go_demo.xjwt.UpdateXjwtReply
	(*DeleteXjwtRequest)(nil), // 4: go_demo.xjwt.DeleteXjwtRequest
	(*DeleteXjwtReply)(nil),   // 5: go_demo.xjwt.DeleteXjwtReply
	(*GetXjwtRequest)(nil),    // 6: go_demo.xjwt.GetXjwtRequest
	(*GetXjwtReply)(nil),      // 7: go_demo.xjwt.GetXjwtReply
	(*ListXjwtRequest)(nil),   // 8: go_demo.xjwt.ListXjwtRequest
	(*ListXjwtReply)(nil),     // 9: go_demo.xjwt.ListXjwtReply
}
var file_xjwt_proto_depIdxs = []int32{
	0, // 0: go_demo.xjwt.Xjwt.CreateXjwt:input_type -> go_demo.xjwt.CreateXjwtRequest
	2, // 1: go_demo.xjwt.Xjwt.UpdateXjwt:input_type -> go_demo.xjwt.UpdateXjwtRequest
	4, // 2: go_demo.xjwt.Xjwt.DeleteXjwt:input_type -> go_demo.xjwt.DeleteXjwtRequest
	6, // 3: go_demo.xjwt.Xjwt.GetXjwt:input_type -> go_demo.xjwt.GetXjwtRequest
	8, // 4: go_demo.xjwt.Xjwt.ListXjwt:input_type -> go_demo.xjwt.ListXjwtRequest
	1, // 5: go_demo.xjwt.Xjwt.CreateXjwt:output_type -> go_demo.xjwt.CreateXjwtReply
	3, // 6: go_demo.xjwt.Xjwt.UpdateXjwt:output_type -> go_demo.xjwt.UpdateXjwtReply
	5, // 7: go_demo.xjwt.Xjwt.DeleteXjwt:output_type -> go_demo.xjwt.DeleteXjwtReply
	7, // 8: go_demo.xjwt.Xjwt.GetXjwt:output_type -> go_demo.xjwt.GetXjwtReply
	9, // 9: go_demo.xjwt.Xjwt.ListXjwt:output_type -> go_demo.xjwt.ListXjwtReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_xjwt_proto_init() }
func file_xjwt_proto_init() {
	if File_xjwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xjwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateXjwtRequest); i {
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
		file_xjwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateXjwtReply); i {
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
		file_xjwt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateXjwtRequest); i {
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
		file_xjwt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateXjwtReply); i {
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
		file_xjwt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteXjwtRequest); i {
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
		file_xjwt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteXjwtReply); i {
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
		file_xjwt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetXjwtRequest); i {
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
		file_xjwt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetXjwtReply); i {
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
		file_xjwt_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListXjwtRequest); i {
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
		file_xjwt_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListXjwtReply); i {
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
			RawDescriptor: file_xjwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xjwt_proto_goTypes,
		DependencyIndexes: file_xjwt_proto_depIdxs,
		MessageInfos:      file_xjwt_proto_msgTypes,
	}.Build()
	File_xjwt_proto = out.File
	file_xjwt_proto_rawDesc = nil
	file_xjwt_proto_goTypes = nil
	file_xjwt_proto_depIdxs = nil
}
