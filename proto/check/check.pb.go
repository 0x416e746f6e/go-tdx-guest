// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: check.proto

// Package check represents an attestation validation policy.

package check

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

// Policy is a representation of an attestation quote validation policy.
// Each field corresponds to a field on validate.Options. This format
// is useful for providing programmatic inputs to the `check` CLI tool.
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HeaderPolicy is representation of Header of an attestation quote validation
	// policy.
	HeaderPolicy *HeaderPolicy `protobuf:"bytes,1,opt,name=header_policy,json=headerPolicy,proto3" json:"header_policy,omitempty"` // should be 20 bytes
	// TDQuoteBodyPolicy is representation of TdQuoteBody of an attestation quote
	// validation policy.
	TdQuoteBodyPolicy *TDQuoteBodyPolicy `protobuf:"bytes,2,opt,name=td_quote_body_policy,json=tdQuoteBodyPolicy,proto3" json:"td_quote_body_policy,omitempty"` // should be 528 bytes
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetHeaderPolicy() *HeaderPolicy {
	if x != nil {
		return x.HeaderPolicy
	}
	return nil
}

func (x *Policy) GetTdQuoteBodyPolicy() *TDQuoteBodyPolicy {
	if x != nil {
		return x.TdQuoteBodyPolicy
	}
	return nil
}

type HeaderPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinmumQeSvn   []byte `protobuf:"bytes,1,opt,name=minmum_qe_svn,json=minmumQeSvn,proto3" json:"minmum_qe_svn,omitempty"`       // should be 2 bytes
	MinimumPceSvn []byte `protobuf:"bytes,2,opt,name=minimum_pce_svn,json=minimumPceSvn,proto3" json:"minimum_pce_svn,omitempty"` // should be 2 bytes
	// Unique vendor id of QE vendor
	QeVendorId []byte `protobuf:"bytes,3,opt,name=qe_vendor_id,json=qeVendorId,proto3" json:"qe_vendor_id,omitempty"` // should be 16 bytes
}

func (x *HeaderPolicy) Reset() {
	*x = HeaderPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderPolicy) ProtoMessage() {}

func (x *HeaderPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderPolicy.ProtoReflect.Descriptor instead.
func (*HeaderPolicy) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderPolicy) GetMinmumQeSvn() []byte {
	if x != nil {
		return x.MinmumQeSvn
	}
	return nil
}

func (x *HeaderPolicy) GetMinimumPceSvn() []byte {
	if x != nil {
		return x.MinimumPceSvn
	}
	return nil
}

func (x *HeaderPolicy) GetQeVendorId() []byte {
	if x != nil {
		return x.QeVendorId
	}
	return nil
}

type TDQuoteBodyPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinimumTeeTcbSvn []byte `protobuf:"bytes,1,opt,name=minimum_tee_tcb_svn,json=minimumTeeTcbSvn,proto3" json:"minimum_tee_tcb_svn,omitempty"` // should be 16 bytes
	MrSeam           []byte `protobuf:"bytes,2,opt,name=mr_seam,json=mrSeam,proto3" json:"mr_seam,omitempty"`                                   // should be 48 bytes
	TdAttributes     []byte `protobuf:"bytes,3,opt,name=td_attributes,json=tdAttributes,proto3" json:"td_attributes,omitempty"`                 // should be 8 bytes
	Xfam             []byte `protobuf:"bytes,4,opt,name=xfam,proto3" json:"xfam,omitempty"`                                                     // should be 8 bytes
	MrTd             []byte `protobuf:"bytes,5,opt,name=mr_td,json=mrTd,proto3" json:"mr_td,omitempty"`                                         // should be 48 bytes
	MrConfigId       []byte `protobuf:"bytes,6,opt,name=mr_config_id,json=mrConfigId,proto3" json:"mr_config_id,omitempty"`                     // should be 48 bytes
	MrOwner          []byte `protobuf:"bytes,7,opt,name=mr_owner,json=mrOwner,proto3" json:"mr_owner,omitempty"`                                // should be 48 bytes
	MrOwnerConfig    []byte `protobuf:"bytes,8,opt,name=mr_owner_config,json=mrOwnerConfig,proto3" json:"mr_owner_config,omitempty"`            // should be 48 bytes
	RtMr0            []byte `protobuf:"bytes,9,opt,name=rt_mr0,json=rtMr0,proto3" json:"rt_mr0,omitempty"`                                      // should be 48 bytes
	RtMr1            []byte `protobuf:"bytes,10,opt,name=rt_mr1,json=rtMr1,proto3" json:"rt_mr1,omitempty"`                                     // should be 48 bytes
	RtMr2            []byte `protobuf:"bytes,11,opt,name=rt_mr2,json=rtMr2,proto3" json:"rt_mr2,omitempty"`                                     // should be 48 bytes
	RtMr3            []byte `protobuf:"bytes,12,opt,name=rt_mr3,json=rtMr3,proto3" json:"rt_mr3,omitempty"`                                     // should be 48 bytes
	ReportData       []byte `protobuf:"bytes,13,opt,name=report_data,json=reportData,proto3" json:"report_data,omitempty"`                      // should be 64 bytes
}

func (x *TDQuoteBodyPolicy) Reset() {
	*x = TDQuoteBodyPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDQuoteBodyPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDQuoteBodyPolicy) ProtoMessage() {}

func (x *TDQuoteBodyPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDQuoteBodyPolicy.ProtoReflect.Descriptor instead.
func (*TDQuoteBodyPolicy) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{2}
}

func (x *TDQuoteBodyPolicy) GetMinimumTeeTcbSvn() []byte {
	if x != nil {
		return x.MinimumTeeTcbSvn
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetMrSeam() []byte {
	if x != nil {
		return x.MrSeam
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetTdAttributes() []byte {
	if x != nil {
		return x.TdAttributes
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetXfam() []byte {
	if x != nil {
		return x.Xfam
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetMrTd() []byte {
	if x != nil {
		return x.MrTd
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetMrConfigId() []byte {
	if x != nil {
		return x.MrConfigId
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetMrOwner() []byte {
	if x != nil {
		return x.MrOwner
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetMrOwnerConfig() []byte {
	if x != nil {
		return x.MrOwnerConfig
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetRtMr0() []byte {
	if x != nil {
		return x.RtMr0
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetRtMr1() []byte {
	if x != nil {
		return x.RtMr1
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetRtMr2() []byte {
	if x != nil {
		return x.RtMr2
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetRtMr3() []byte {
	if x != nil {
		return x.RtMr3
	}
	return nil
}

func (x *TDQuoteBodyPolicy) GetReportData() []byte {
	if x != nil {
		return x.ReportData
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The quote validation policy.
	Policy *Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_check_proto protoreflect.FileDescriptor

var file_check_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x22, 0x8d, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x38, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0c, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x49, 0x0a, 0x14, 0x74, 0x64, 0x5f,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x54, 0x44, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x11, 0x74, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x22, 0x7c, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x6d, 0x75, 0x6d, 0x5f, 0x71,
	0x65, 0x5f, 0x73, 0x76, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6d, 0x69, 0x6e,
	0x6d, 0x75, 0x6d, 0x51, 0x65, 0x53, 0x76, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x69,
	0x6d, 0x75, 0x6d, 0x5f, 0x70, 0x63, 0x65, 0x5f, 0x73, 0x76, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x63, 0x65, 0x53, 0x76, 0x6e,
	0x12, 0x20, 0x0a, 0x0c, 0x71, 0x65, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x71, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x49, 0x64, 0x22, 0x8b, 0x03, 0x0a, 0x11, 0x54, 0x44, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x6f,
	0x64, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x69,
	0x6d, 0x75, 0x6d, 0x5f, 0x74, 0x65, 0x65, 0x5f, 0x74, 0x63, 0x62, 0x5f, 0x73, 0x76, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x54, 0x65,
	0x65, 0x54, 0x63, 0x62, 0x53, 0x76, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x72, 0x5f, 0x73, 0x65,
	0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6d, 0x72, 0x53, 0x65, 0x61, 0x6d,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x74, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x66, 0x61, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x78, 0x66, 0x61, 0x6d, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x72, 0x5f,
	0x74, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x72, 0x54, 0x64, 0x12, 0x20,
	0x0a, 0x0c, 0x6d, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x72, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x6d, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x6d,
	0x72, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6d, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x74, 0x5f, 0x6d, 0x72, 0x30, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x74, 0x4d, 0x72, 0x30, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x74,
	0x5f, 0x6d, 0x72, 0x31, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x74, 0x4d, 0x72,
	0x31, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x74, 0x5f, 0x6d, 0x72, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x72, 0x74, 0x4d, 0x72, 0x32, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x74, 0x5f, 0x6d,
	0x72, 0x33, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x74, 0x4d, 0x72, 0x33, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x2f, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x64, 0x78, 0x2d, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_proto_rawDescOnce sync.Once
	file_check_proto_rawDescData = file_check_proto_rawDesc
)

func file_check_proto_rawDescGZIP() []byte {
	file_check_proto_rawDescOnce.Do(func() {
		file_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_proto_rawDescData)
	})
	return file_check_proto_rawDescData
}

var file_check_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_check_proto_goTypes = []interface{}{
	(*Policy)(nil),            // 0: check.Policy
	(*HeaderPolicy)(nil),      // 1: check.HeaderPolicy
	(*TDQuoteBodyPolicy)(nil), // 2: check.TDQuoteBodyPolicy
	(*Config)(nil),            // 3: check.Config
}
var file_check_proto_depIdxs = []int32{
	1, // 0: check.Policy.header_policy:type_name -> check.HeaderPolicy
	2, // 1: check.Policy.td_quote_body_policy:type_name -> check.TDQuoteBodyPolicy
	0, // 2: check.Config.policy:type_name -> check.Policy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_check_proto_init() }
func file_check_proto_init() {
	if File_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_check_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderPolicy); i {
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
		file_check_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDQuoteBodyPolicy); i {
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
		file_check_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_check_proto_goTypes,
		DependencyIndexes: file_check_proto_depIdxs,
		MessageInfos:      file_check_proto_msgTypes,
	}.Build()
	File_check_proto = out.File
	file_check_proto_rawDesc = nil
	file_check_proto_goTypes = nil
	file_check_proto_depIdxs = nil
}
