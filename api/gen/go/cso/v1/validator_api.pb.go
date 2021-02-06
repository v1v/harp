// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cso/v1/validator_api.proto

package csov1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_validator_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_validator_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_cso_v1_validator_api_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// ValidateResponse returns the secret path evaluation.
type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret *Secret `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cso_v1_validator_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cso_v1_validator_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_cso_v1_validator_api_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateResponse) GetSecret() *Secret {
	if x != nil {
		return x.Secret
	}
	return nil
}

var File_cso_v1_validator_api_proto protoreflect.FileDescriptor

var file_cso_v1_validator_api_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x73, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x73,
	0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x63, 0x73, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x3a, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x32, 0x4d, 0x0a, 0x0c,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x50, 0x49, 0x12, 0x3d, 0x0a, 0x08,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x63, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x82, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x65, 0x63, 0x2e, 0x63, 0x73, 0x6f, 0x2e,
	0x76, 0x31, 0x42, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x68, 0x61, 0x72, 0x70,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x73, 0x6f, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x73, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02,
	0x06, 0x43, 0x73, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x43, 0x73, 0x6f, 0x5c, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cso_v1_validator_api_proto_rawDescOnce sync.Once
	file_cso_v1_validator_api_proto_rawDescData = file_cso_v1_validator_api_proto_rawDesc
)

func file_cso_v1_validator_api_proto_rawDescGZIP() []byte {
	file_cso_v1_validator_api_proto_rawDescOnce.Do(func() {
		file_cso_v1_validator_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cso_v1_validator_api_proto_rawDescData)
	})
	return file_cso_v1_validator_api_proto_rawDescData
}

var (
	file_cso_v1_validator_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_cso_v1_validator_api_proto_goTypes  = []interface{}{
		(*ValidateRequest)(nil),  // 0: cso.v1.ValidateRequest
		(*ValidateResponse)(nil), // 1: cso.v1.ValidateResponse
		(*Secret)(nil),           // 2: cso.v1.Secret
	}
)

var file_cso_v1_validator_api_proto_depIdxs = []int32{
	2, // 0: cso.v1.ValidateResponse.secret:type_name -> cso.v1.Secret
	0, // 1: cso.v1.ValidatorAPI.Validate:input_type -> cso.v1.ValidateRequest
	1, // 2: cso.v1.ValidatorAPI.Validate:output_type -> cso.v1.ValidateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cso_v1_validator_api_proto_init() }
func file_cso_v1_validator_api_proto_init() {
	if File_cso_v1_validator_api_proto != nil {
		return
	}
	file_cso_v1_secret_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cso_v1_validator_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_cso_v1_validator_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
			RawDescriptor: file_cso_v1_validator_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cso_v1_validator_api_proto_goTypes,
		DependencyIndexes: file_cso_v1_validator_api_proto_depIdxs,
		MessageInfos:      file_cso_v1_validator_api_proto_msgTypes,
	}.Build()
	File_cso_v1_validator_api_proto = out.File
	file_cso_v1_validator_api_proto_rawDesc = nil
	file_cso_v1_validator_api_proto_goTypes = nil
	file_cso_v1_validator_api_proto_depIdxs = nil
}
