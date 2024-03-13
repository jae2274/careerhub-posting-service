// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: careerhub/processor/scanner_grpc/scanner_grpc/grpc.proto

package scanner_grpc

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

type ScanComplete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsScanComplete bool `protobuf:"varint,1,opt,name=isScanComplete,proto3" json:"isScanComplete,omitempty"`
}

func (x *ScanComplete) Reset() {
	*x = ScanComplete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanComplete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanComplete) ProtoMessage() {}

func (x *ScanComplete) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanComplete.ProtoReflect.Descriptor instead.
func (*ScanComplete) Descriptor() ([]byte, []int) {
	return file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *ScanComplete) GetIsScanComplete() bool {
	if x != nil {
		return x.IsScanComplete
	}
	return false
}

type JobPostingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site           string   `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	PostingId      string   `protobuf:"bytes,2,opt,name=postingId,proto3" json:"postingId,omitempty"`
	Title          string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Qualifications string   `protobuf:"bytes,4,opt,name=qualifications,proto3" json:"qualifications,omitempty"`
	Preferred      string   `protobuf:"bytes,5,opt,name=preferred,proto3" json:"preferred,omitempty"`
	RequiredSkill  []string `protobuf:"bytes,6,rep,name=requiredSkill,proto3" json:"requiredSkill,omitempty"`
}

func (x *JobPostingInfo) Reset() {
	*x = JobPostingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobPostingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobPostingInfo) ProtoMessage() {}

func (x *JobPostingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobPostingInfo.ProtoReflect.Descriptor instead.
func (*JobPostingInfo) Descriptor() ([]byte, []int) {
	return file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *JobPostingInfo) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *JobPostingInfo) GetPostingId() string {
	if x != nil {
		return x.PostingId
	}
	return ""
}

func (x *JobPostingInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *JobPostingInfo) GetQualifications() string {
	if x != nil {
		return x.Qualifications
	}
	return ""
}

func (x *JobPostingInfo) GetPreferred() string {
	if x != nil {
		return x.Preferred
	}
	return ""
}

func (x *JobPostingInfo) GetRequiredSkill() []string {
	if x != nil {
		return x.RequiredSkill
	}
	return nil
}

type SetRequiredSkillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site          string           `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	PostingId     string           `protobuf:"bytes,2,opt,name=postingId,proto3" json:"postingId,omitempty"`
	RequiredSkill []*RequiredSkill `protobuf:"bytes,3,rep,name=requiredSkill,proto3" json:"requiredSkill,omitempty"`
}

func (x *SetRequiredSkillsRequest) Reset() {
	*x = SetRequiredSkillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequiredSkillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequiredSkillsRequest) ProtoMessage() {}

func (x *SetRequiredSkillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequiredSkillsRequest.ProtoReflect.Descriptor instead.
func (*SetRequiredSkillsRequest) Descriptor() ([]byte, []int) {
	return file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *SetRequiredSkillsRequest) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *SetRequiredSkillsRequest) GetPostingId() string {
	if x != nil {
		return x.PostingId
	}
	return ""
}

func (x *SetRequiredSkillsRequest) GetRequiredSkill() []*RequiredSkill {
	if x != nil {
		return x.RequiredSkill
	}
	return nil
}

type Skills struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillNames []string `protobuf:"bytes,1,rep,name=skillNames,proto3" json:"skillNames,omitempty"`
}

func (x *Skills) Reset() {
	*x = Skills{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skills) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skills) ProtoMessage() {}

func (x *Skills) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skills.ProtoReflect.Descriptor instead.
func (*Skills) Descriptor() ([]byte, []int) {
	return file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Skills) GetSkillNames() []string {
	if x != nil {
		return x.SkillNames
	}
	return nil
}

type RequiredSkill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillName string `protobuf:"bytes,1,opt,name=skillName,proto3" json:"skillName,omitempty"`
	SkillFrom string `protobuf:"bytes,2,opt,name=skillFrom,proto3" json:"skillFrom,omitempty"`
}

func (x *RequiredSkill) Reset() {
	*x = RequiredSkill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredSkill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredSkill) ProtoMessage() {}

func (x *RequiredSkill) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredSkill.ProtoReflect.Descriptor instead.
func (*RequiredSkill) Descriptor() ([]byte, []int) {
	return file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *RequiredSkill) GetSkillName() string {
	if x != nil {
		return x.SkillName
	}
	return ""
}

func (x *RequiredSkill) GetSkillFrom() string {
	if x != nil {
		return x.SkillFrom
	}
	return ""
}

type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *BoolResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto protoreflect.FileDescriptor

var file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x36, 0x0a, 0x0c,
	0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x69, 0x73, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x22, 0xa3, 0x01, 0x0a, 0x18,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x55, 0x0a, 0x0d, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x22, 0x28, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x0d, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x22, 0x28, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0xdb, 0x03, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x72,
	0x70, 0x63, 0x12, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x1a, 0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01, 0x12, 0x65, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x28, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12,
	0x81, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x3a, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x12, 0x6b, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68,
	0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x1a, 0x2e, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x22, 0x5a, 0x20, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescOnce sync.Once
	file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescData = file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDesc
)

func file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescGZIP() []byte {
	file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescOnce.Do(func() {
		file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescData)
	})
	return file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDescData
}

var file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_goTypes = []interface{}{
	(*ScanComplete)(nil),             // 0: careerhub.processor.scanner_grpc.ScanComplete
	(*JobPostingInfo)(nil),           // 1: careerhub.processor.scanner_grpc.JobPostingInfo
	(*SetRequiredSkillsRequest)(nil), // 2: careerhub.processor.scanner_grpc.SetRequiredSkillsRequest
	(*Skills)(nil),                   // 3: careerhub.processor.scanner_grpc.Skills
	(*RequiredSkill)(nil),            // 4: careerhub.processor.scanner_grpc.RequiredSkill
	(*BoolResponse)(nil),             // 5: careerhub.processor.scanner_grpc.BoolResponse
}
var file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_depIdxs = []int32{
	4, // 0: careerhub.processor.scanner_grpc.SetRequiredSkillsRequest.requiredSkill:type_name -> careerhub.processor.scanner_grpc.RequiredSkill
	0, // 1: careerhub.processor.scanner_grpc.ScannerGrpc.GetJobPostings:input_type -> careerhub.processor.scanner_grpc.ScanComplete
	0, // 2: careerhub.processor.scanner_grpc.ScannerGrpc.GetSkills:input_type -> careerhub.processor.scanner_grpc.ScanComplete
	2, // 3: careerhub.processor.scanner_grpc.ScannerGrpc.SetRequiredSkills:input_type -> careerhub.processor.scanner_grpc.SetRequiredSkillsRequest
	3, // 4: careerhub.processor.scanner_grpc.ScannerGrpc.SetScanComplete:input_type -> careerhub.processor.scanner_grpc.Skills
	1, // 5: careerhub.processor.scanner_grpc.ScannerGrpc.GetJobPostings:output_type -> careerhub.processor.scanner_grpc.JobPostingInfo
	3, // 6: careerhub.processor.scanner_grpc.ScannerGrpc.GetSkills:output_type -> careerhub.processor.scanner_grpc.Skills
	5, // 7: careerhub.processor.scanner_grpc.ScannerGrpc.SetRequiredSkills:output_type -> careerhub.processor.scanner_grpc.BoolResponse
	5, // 8: careerhub.processor.scanner_grpc.ScannerGrpc.SetScanComplete:output_type -> careerhub.processor.scanner_grpc.BoolResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_init() }
func file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_init() {
	if File_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanComplete); i {
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
		file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobPostingInfo); i {
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
		file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequiredSkillsRequest); i {
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
		file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skills); i {
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
		file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredSkill); i {
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
		file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
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
			RawDescriptor: file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_goTypes,
		DependencyIndexes: file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_depIdxs,
		MessageInfos:      file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_msgTypes,
	}.Build()
	File_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto = out.File
	file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_rawDesc = nil
	file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_goTypes = nil
	file_careerhub_processor_scanner_grpc_scanner_grpc_grpc_proto_depIdxs = nil
}
