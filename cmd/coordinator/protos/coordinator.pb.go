// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.2
// source: cmd/coordinator/protos/coordinator.proto

package protos

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

// ClearResultsRequest specifies the data needed to clear a result.
type ClearResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// builder is the builder to clear results.
	Builder string `protobuf:"bytes,1,opt,name=builder,proto3" json:"builder,omitempty"`
	// hash is the commit hash to clear results.
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ClearResultsRequest) Reset() {
	*x = ClearResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_coordinator_protos_coordinator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearResultsRequest) ProtoMessage() {}

func (x *ClearResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_coordinator_protos_coordinator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearResultsRequest.ProtoReflect.Descriptor instead.
func (*ClearResultsRequest) Descriptor() ([]byte, []int) {
	return file_cmd_coordinator_protos_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *ClearResultsRequest) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

func (x *ClearResultsRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type ClearResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearResultsResponse) Reset() {
	*x = ClearResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_coordinator_protos_coordinator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearResultsResponse) ProtoMessage() {}

func (x *ClearResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_coordinator_protos_coordinator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearResultsResponse.ProtoReflect.Descriptor instead.
func (*ClearResultsResponse) Descriptor() ([]byte, []int) {
	return file_cmd_coordinator_protos_coordinator_proto_rawDescGZIP(), []int{1}
}

var File_cmd_coordinator_protos_coordinator_proto protoreflect.FileDescriptor

var file_cmd_coordinator_protos_coordinator_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x5a, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x4b,
	0x0a, 0x0c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x78, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_coordinator_protos_coordinator_proto_rawDescOnce sync.Once
	file_cmd_coordinator_protos_coordinator_proto_rawDescData = file_cmd_coordinator_protos_coordinator_proto_rawDesc
)

func file_cmd_coordinator_protos_coordinator_proto_rawDescGZIP() []byte {
	file_cmd_coordinator_protos_coordinator_proto_rawDescOnce.Do(func() {
		file_cmd_coordinator_protos_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_coordinator_protos_coordinator_proto_rawDescData)
	})
	return file_cmd_coordinator_protos_coordinator_proto_rawDescData
}

var file_cmd_coordinator_protos_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_coordinator_protos_coordinator_proto_goTypes = []interface{}{
	(*ClearResultsRequest)(nil),  // 0: protos.ClearResultsRequest
	(*ClearResultsResponse)(nil), // 1: protos.ClearResultsResponse
}
var file_cmd_coordinator_protos_coordinator_proto_depIdxs = []int32{
	0, // 0: protos.Coordinator.ClearResults:input_type -> protos.ClearResultsRequest
	1, // 1: protos.Coordinator.ClearResults:output_type -> protos.ClearResultsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_coordinator_protos_coordinator_proto_init() }
func file_cmd_coordinator_protos_coordinator_proto_init() {
	if File_cmd_coordinator_protos_coordinator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_coordinator_protos_coordinator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearResultsRequest); i {
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
		file_cmd_coordinator_protos_coordinator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearResultsResponse); i {
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
			RawDescriptor: file_cmd_coordinator_protos_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_coordinator_protos_coordinator_proto_goTypes,
		DependencyIndexes: file_cmd_coordinator_protos_coordinator_proto_depIdxs,
		MessageInfos:      file_cmd_coordinator_protos_coordinator_proto_msgTypes,
	}.Build()
	File_cmd_coordinator_protos_coordinator_proto = out.File
	file_cmd_coordinator_protos_coordinator_proto_rawDesc = nil
	file_cmd_coordinator_protos_coordinator_proto_goTypes = nil
	file_cmd_coordinator_protos_coordinator_proto_depIdxs = nil
}
