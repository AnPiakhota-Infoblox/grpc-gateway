// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: examples/internal/proto/examplepb/non_standard_names.proto

package examplepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// NonStandardMessage has oddly named fields.
type NonStandardMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id        string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num       int64                     `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	LineNum   int64                     `protobuf:"varint,3,opt,name=line_num,json=lineNum,proto3" json:"line_num,omitempty"`
	LangIdent string                    `protobuf:"bytes,4,opt,name=langIdent,proto3" json:"langIdent,omitempty"`
	STATUS    string                    `protobuf:"bytes,5,opt,name=STATUS,proto3" json:"STATUS,omitempty"`
	En_GB     int64                     `protobuf:"varint,6,opt,name=en_GB,json=enGB,proto3" json:"en_GB,omitempty"`
	No        string                    `protobuf:"bytes,7,opt,name=no,proto3" json:"no,omitempty"`
	Thing     *NonStandardMessage_Thing `protobuf:"bytes,8,opt,name=thing,proto3" json:"thing,omitempty"`
}

func (x *NonStandardMessage) Reset() {
	*x = NonStandardMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardMessage) ProtoMessage() {}

func (x *NonStandardMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardMessage.ProtoReflect.Descriptor instead.
func (*NonStandardMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{0}
}

func (x *NonStandardMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NonStandardMessage) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *NonStandardMessage) GetLineNum() int64 {
	if x != nil {
		return x.LineNum
	}
	return 0
}

func (x *NonStandardMessage) GetLangIdent() string {
	if x != nil {
		return x.LangIdent
	}
	return ""
}

func (x *NonStandardMessage) GetSTATUS() string {
	if x != nil {
		return x.STATUS
	}
	return ""
}

func (x *NonStandardMessage) GetEn_GB() int64 {
	if x != nil {
		return x.En_GB
	}
	return 0
}

func (x *NonStandardMessage) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *NonStandardMessage) GetThing() *NonStandardMessage_Thing {
	if x != nil {
		return x.Thing
	}
	return nil
}

type NonStandardUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body       *NonStandardMessage    `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *NonStandardUpdateRequest) Reset() {
	*x = NonStandardUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardUpdateRequest) ProtoMessage() {}

func (x *NonStandardUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardUpdateRequest.ProtoReflect.Descriptor instead.
func (*NonStandardUpdateRequest) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{1}
}

func (x *NonStandardUpdateRequest) GetBody() *NonStandardMessage {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *NonStandardUpdateRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// NonStandardMessageWithJSONNames maps odd field names to odd JSON names for maximum confusion.
type NonStandardMessageWithJSONNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id        string                                 `protobuf:"bytes,1,opt,name=id,json=ID,proto3" json:"id,omitempty"`
	Num       int64                                  `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	LineNum   int64                                  `protobuf:"varint,3,opt,name=line_num,json=LineNum,proto3" json:"line_num,omitempty"`
	LangIdent string                                 `protobuf:"bytes,4,opt,name=langIdent,proto3" json:"langIdent,omitempty"`
	STATUS    string                                 `protobuf:"bytes,5,opt,name=STATUS,json=status,proto3" json:"STATUS,omitempty"`
	En_GB     int64                                  `protobuf:"varint,6,opt,name=en_GB,json=En_GB,proto3" json:"en_GB,omitempty"`
	No        string                                 `protobuf:"bytes,7,opt,name=no,json=yes,proto3" json:"no,omitempty"`
	Thing     *NonStandardMessageWithJSONNames_Thing `protobuf:"bytes,8,opt,name=thing,json=Thingy,proto3" json:"thing,omitempty"`
}

func (x *NonStandardMessageWithJSONNames) Reset() {
	*x = NonStandardMessageWithJSONNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardMessageWithJSONNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardMessageWithJSONNames) ProtoMessage() {}

func (x *NonStandardMessageWithJSONNames) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardMessageWithJSONNames.ProtoReflect.Descriptor instead.
func (*NonStandardMessageWithJSONNames) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{2}
}

func (x *NonStandardMessageWithJSONNames) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NonStandardMessageWithJSONNames) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *NonStandardMessageWithJSONNames) GetLineNum() int64 {
	if x != nil {
		return x.LineNum
	}
	return 0
}

func (x *NonStandardMessageWithJSONNames) GetLangIdent() string {
	if x != nil {
		return x.LangIdent
	}
	return ""
}

func (x *NonStandardMessageWithJSONNames) GetSTATUS() string {
	if x != nil {
		return x.STATUS
	}
	return ""
}

func (x *NonStandardMessageWithJSONNames) GetEn_GB() int64 {
	if x != nil {
		return x.En_GB
	}
	return 0
}

func (x *NonStandardMessageWithJSONNames) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *NonStandardMessageWithJSONNames) GetThing() *NonStandardMessageWithJSONNames_Thing {
	if x != nil {
		return x.Thing
	}
	return nil
}

type NonStandardWithJSONNamesUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body       *NonStandardMessageWithJSONNames `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask           `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *NonStandardWithJSONNamesUpdateRequest) Reset() {
	*x = NonStandardWithJSONNamesUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardWithJSONNamesUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardWithJSONNamesUpdateRequest) ProtoMessage() {}

func (x *NonStandardWithJSONNamesUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardWithJSONNamesUpdateRequest.ProtoReflect.Descriptor instead.
func (*NonStandardWithJSONNamesUpdateRequest) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{3}
}

func (x *NonStandardWithJSONNamesUpdateRequest) GetBody() *NonStandardMessageWithJSONNames {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *NonStandardWithJSONNamesUpdateRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type NonStandardMessage_Thing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubThing *NonStandardMessage_Thing_SubThing `protobuf:"bytes,1,opt,name=subThing,proto3" json:"subThing,omitempty"`
}

func (x *NonStandardMessage_Thing) Reset() {
	*x = NonStandardMessage_Thing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardMessage_Thing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardMessage_Thing) ProtoMessage() {}

func (x *NonStandardMessage_Thing) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardMessage_Thing.ProtoReflect.Descriptor instead.
func (*NonStandardMessage_Thing) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NonStandardMessage_Thing) GetSubThing() *NonStandardMessage_Thing_SubThing {
	if x != nil {
		return x.SubThing
	}
	return nil
}

type NonStandardMessage_Thing_SubThing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubValue string `protobuf:"bytes,1,opt,name=sub_value,json=subValue,proto3" json:"sub_value,omitempty"`
}

func (x *NonStandardMessage_Thing_SubThing) Reset() {
	*x = NonStandardMessage_Thing_SubThing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardMessage_Thing_SubThing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardMessage_Thing_SubThing) ProtoMessage() {}

func (x *NonStandardMessage_Thing_SubThing) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardMessage_Thing_SubThing.ProtoReflect.Descriptor instead.
func (*NonStandardMessage_Thing_SubThing) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *NonStandardMessage_Thing_SubThing) GetSubValue() string {
	if x != nil {
		return x.SubValue
	}
	return ""
}

type NonStandardMessageWithJSONNames_Thing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubThing *NonStandardMessageWithJSONNames_Thing_SubThing `protobuf:"bytes,1,opt,name=subThing,json=SubThing,proto3" json:"subThing,omitempty"`
}

func (x *NonStandardMessageWithJSONNames_Thing) Reset() {
	*x = NonStandardMessageWithJSONNames_Thing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardMessageWithJSONNames_Thing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardMessageWithJSONNames_Thing) ProtoMessage() {}

func (x *NonStandardMessageWithJSONNames_Thing) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardMessageWithJSONNames_Thing.ProtoReflect.Descriptor instead.
func (*NonStandardMessageWithJSONNames_Thing) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{2, 0}
}

func (x *NonStandardMessageWithJSONNames_Thing) GetSubThing() *NonStandardMessageWithJSONNames_Thing_SubThing {
	if x != nil {
		return x.SubThing
	}
	return nil
}

type NonStandardMessageWithJSONNames_Thing_SubThing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubValue string `protobuf:"bytes,1,opt,name=sub_value,json=sub_Value,proto3" json:"sub_value,omitempty"`
}

func (x *NonStandardMessageWithJSONNames_Thing_SubThing) Reset() {
	*x = NonStandardMessageWithJSONNames_Thing_SubThing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonStandardMessageWithJSONNames_Thing_SubThing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonStandardMessageWithJSONNames_Thing_SubThing) ProtoMessage() {}

func (x *NonStandardMessageWithJSONNames_Thing_SubThing) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonStandardMessageWithJSONNames_Thing_SubThing.ProtoReflect.Descriptor instead.
func (*NonStandardMessageWithJSONNames_Thing_SubThing) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *NonStandardMessageWithJSONNames_Thing_SubThing) GetSubValue() string {
	if x != nil {
		return x.SubValue
	}
	return ""
}

var File_examples_internal_proto_examplepb_non_standard_names_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_non_standard_names_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x03, 0x0a,
	0x12, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x6e, 0x5f, 0x47, 0x42, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x65, 0x6e, 0x47, 0x42, 0x12, 0x0e, 0x0a, 0x02, 0x6e,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x6f, 0x12, 0x5e, 0x0a, 0x05, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x9f, 0x01, 0x0a, 0x05,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x6d, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x75, 0x62, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x1a, 0x27, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x54, 0x68, 0x69, 0x6e, 0x67,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xaf, 0x01,
	0x0a, 0x18, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22,
	0xd9, 0x03, 0x0a, 0x1f, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4a, 0x53, 0x4f, 0x4e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x5f, 0x47, 0x42, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x45, 0x6e, 0x5f, 0x47, 0x42, 0x12, 0x0f, 0x0a, 0x02,
	0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x79, 0x65, 0x73, 0x12, 0x6c, 0x0a,
	0x05, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x4a, 0x53, 0x4f, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x52, 0x06, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x79, 0x1a, 0xad, 0x01, 0x0a, 0x05,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x7a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4a,
	0x53, 0x4f, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x75, 0x62, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x53, 0x75, 0x62, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x1a, 0x28, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x75, 0x62, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x25,
	0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x57, 0x69, 0x74, 0x68, 0x4a,
	0x53, 0x4f, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4a, 0x53, 0x4f, 0x4e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0xdb, 0x03, 0x0a, 0x12, 0x4e, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc5,
	0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x32,
	0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6e, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0xfc, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x4a, 0x53, 0x4f, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x55,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x57, 0x69, 0x74, 0x68, 0x4a,
	0x53, 0x4f, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4a, 0x53, 0x4f,
	0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x32, 0x2f,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6e, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x3a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescData = file_examples_internal_proto_examplepb_non_standard_names_proto_rawDesc
)

func file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_non_standard_names_proto_rawDescData
}

var file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_examples_internal_proto_examplepb_non_standard_names_proto_goTypes = []interface{}{
	(*NonStandardMessage)(nil),                             // 0: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage
	(*NonStandardUpdateRequest)(nil),                       // 1: grpc.gateway.examples.internal.proto.examplepb.NonStandardUpdateRequest
	(*NonStandardMessageWithJSONNames)(nil),                // 2: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames
	(*NonStandardWithJSONNamesUpdateRequest)(nil),          // 3: grpc.gateway.examples.internal.proto.examplepb.NonStandardWithJSONNamesUpdateRequest
	(*NonStandardMessage_Thing)(nil),                       // 4: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage.Thing
	(*NonStandardMessage_Thing_SubThing)(nil),              // 5: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage.Thing.SubThing
	(*NonStandardMessageWithJSONNames_Thing)(nil),          // 6: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames.Thing
	(*NonStandardMessageWithJSONNames_Thing_SubThing)(nil), // 7: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames.Thing.SubThing
	(*fieldmaskpb.FieldMask)(nil),                          // 8: google.protobuf.FieldMask
}
var file_examples_internal_proto_examplepb_non_standard_names_proto_depIdxs = []int32{
	4,  // 0: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage.thing:type_name -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage.Thing
	0,  // 1: grpc.gateway.examples.internal.proto.examplepb.NonStandardUpdateRequest.body:type_name -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage
	8,  // 2: grpc.gateway.examples.internal.proto.examplepb.NonStandardUpdateRequest.update_mask:type_name -> google.protobuf.FieldMask
	6,  // 3: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames.thing:type_name -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames.Thing
	2,  // 4: grpc.gateway.examples.internal.proto.examplepb.NonStandardWithJSONNamesUpdateRequest.body:type_name -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames
	8,  // 5: grpc.gateway.examples.internal.proto.examplepb.NonStandardWithJSONNamesUpdateRequest.update_mask:type_name -> google.protobuf.FieldMask
	5,  // 6: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage.Thing.subThing:type_name -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage.Thing.SubThing
	7,  // 7: grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames.Thing.subThing:type_name -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames.Thing.SubThing
	1,  // 8: grpc.gateway.examples.internal.proto.examplepb.NonStandardService.Update:input_type -> grpc.gateway.examples.internal.proto.examplepb.NonStandardUpdateRequest
	3,  // 9: grpc.gateway.examples.internal.proto.examplepb.NonStandardService.UpdateWithJSONNames:input_type -> grpc.gateway.examples.internal.proto.examplepb.NonStandardWithJSONNamesUpdateRequest
	0,  // 10: grpc.gateway.examples.internal.proto.examplepb.NonStandardService.Update:output_type -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessage
	2,  // 11: grpc.gateway.examples.internal.proto.examplepb.NonStandardService.UpdateWithJSONNames:output_type -> grpc.gateway.examples.internal.proto.examplepb.NonStandardMessageWithJSONNames
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_non_standard_names_proto_init() }
func file_examples_internal_proto_examplepb_non_standard_names_proto_init() {
	if File_examples_internal_proto_examplepb_non_standard_names_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardMessage); i {
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
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardUpdateRequest); i {
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
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardMessageWithJSONNames); i {
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
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardWithJSONNamesUpdateRequest); i {
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
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardMessage_Thing); i {
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
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardMessage_Thing_SubThing); i {
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
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardMessageWithJSONNames_Thing); i {
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
		file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonStandardMessageWithJSONNames_Thing_SubThing); i {
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
			RawDescriptor: file_examples_internal_proto_examplepb_non_standard_names_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_non_standard_names_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_non_standard_names_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_non_standard_names_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_non_standard_names_proto = out.File
	file_examples_internal_proto_examplepb_non_standard_names_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_non_standard_names_proto_goTypes = nil
	file_examples_internal_proto_examplepb_non_standard_names_proto_depIdxs = nil
}
