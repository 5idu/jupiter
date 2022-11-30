// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: testproto/v1/hello.proto

package testproto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the user
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_v1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_v1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_testproto_v1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request message containing the user's name.
type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error is the error code.
	Error uint32 `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	// msg is the message.
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// data is the data.
	Data *SayHelloResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_v1_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_v1_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_testproto_v1_hello_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *SayHelloResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SayHelloResponse) GetData() *SayHelloResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// Data is the data to be sent.
type SayHelloResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the user
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloResponse_Data) Reset() {
	*x = SayHelloResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_v1_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse_Data) ProtoMessage() {}

func (x *SayHelloResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_v1_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse_Data.ProtoReflect.Descriptor instead.
func (*SayHelloResponse_Data) Descriptor() ([]byte, []int) {
	return file_testproto_v1_hello_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SayHelloResponse_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_testproto_v1_hello_proto protoreflect.FileDescriptor

var file_testproto_v1_hello_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8f, 0x01,
	0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x1a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32,
	0x87, 0x01, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x75, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1d,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x2f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x42, 0xa6, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0a,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x75, 0x2f, 0x6a,
	0x75, 0x70, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x54, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x54, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testproto_v1_hello_proto_rawDescOnce sync.Once
	file_testproto_v1_hello_proto_rawDescData = file_testproto_v1_hello_proto_rawDesc
)

func file_testproto_v1_hello_proto_rawDescGZIP() []byte {
	file_testproto_v1_hello_proto_rawDescOnce.Do(func() {
		file_testproto_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_testproto_v1_hello_proto_rawDescData)
	})
	return file_testproto_v1_hello_proto_rawDescData
}

var file_testproto_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_testproto_v1_hello_proto_goTypes = []interface{}{
	(*SayHelloRequest)(nil),       // 0: testproto.v1.SayHelloRequest
	(*SayHelloResponse)(nil),      // 1: testproto.v1.SayHelloResponse
	(*SayHelloResponse_Data)(nil), // 2: testproto.v1.SayHelloResponse.Data
}
var file_testproto_v1_hello_proto_depIdxs = []int32{
	2, // 0: testproto.v1.SayHelloResponse.data:type_name -> testproto.v1.SayHelloResponse.Data
	0, // 1: testproto.v1.GreeterService.SayHello:input_type -> testproto.v1.SayHelloRequest
	1, // 2: testproto.v1.GreeterService.SayHello:output_type -> testproto.v1.SayHelloResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_testproto_v1_hello_proto_init() }
func file_testproto_v1_hello_proto_init() {
	if File_testproto_v1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testproto_v1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_testproto_v1_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_testproto_v1_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse_Data); i {
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
			RawDescriptor: file_testproto_v1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testproto_v1_hello_proto_goTypes,
		DependencyIndexes: file_testproto_v1_hello_proto_depIdxs,
		MessageInfos:      file_testproto_v1_hello_proto_msgTypes,
	}.Build()
	File_testproto_v1_hello_proto = out.File
	file_testproto_v1_hello_proto_rawDesc = nil
	file_testproto_v1_hello_proto_goTypes = nil
	file_testproto_v1_hello_proto_depIdxs = nil
}
