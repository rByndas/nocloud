//
//Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pkg/states/proto/states.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NoCloudState int32

const (
	NoCloudState_INIT      NoCloudState = 0
	NoCloudState_UNKNOWN   NoCloudState = 1
	NoCloudState_STOPPED   NoCloudState = 2
	NoCloudState_RUNNING   NoCloudState = 3
	NoCloudState_FAILURE   NoCloudState = 4
	NoCloudState_DELETED   NoCloudState = 5
	NoCloudState_SUSPENDED NoCloudState = 6
	NoCloudState_OPERATION NoCloudState = 7
)

// Enum value maps for NoCloudState.
var (
	NoCloudState_name = map[int32]string{
		0: "INIT",
		1: "UNKNOWN",
		2: "STOPPED",
		3: "RUNNING",
		4: "FAILURE",
		5: "DELETED",
		6: "SUSPENDED",
		7: "OPERATION",
	}
	NoCloudState_value = map[string]int32{
		"INIT":      0,
		"UNKNOWN":   1,
		"STOPPED":   2,
		"RUNNING":   3,
		"FAILURE":   4,
		"DELETED":   5,
		"SUSPENDED": 6,
		"OPERATION": 7,
	}
)

func (x NoCloudState) Enum() *NoCloudState {
	p := new(NoCloudState)
	*p = x
	return p
}

func (x NoCloudState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NoCloudState) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_states_proto_states_proto_enumTypes[0].Descriptor()
}

func (NoCloudState) Type() protoreflect.EnumType {
	return &file_pkg_states_proto_states_proto_enumTypes[0]
}

func (x NoCloudState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NoCloudState.Descriptor instead.
func (NoCloudState) EnumDescriptor() ([]byte, []int) {
	return file_pkg_states_proto_states_proto_rawDescGZIP(), []int{0}
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State NoCloudState               `protobuf:"varint,1,opt,name=state,proto3,enum=nocloud.states.NoCloudState" json:"state,omitempty"`                                                     // NoCloud Instance State
	Meta  map[string]*structpb.Value `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Driver(Provider/Hypervisor) State data
	Ts    *int64                     `protobuf:"varint,3,opt,name=ts,proto3,oneof" json:"ts,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_states_proto_states_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_states_proto_states_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_pkg_states_proto_states_proto_rawDescGZIP(), []int{0}
}

func (x *State) GetState() NoCloudState {
	if x != nil {
		return x.State
	}
	return NoCloudState_INIT
}

func (x *State) GetMeta() map[string]*structpb.Value {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *State) GetTs() int64 {
	if x != nil && x.Ts != nil {
		return *x.Ts
	}
	return 0
}

type ObjectState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	State *State `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ObjectState) Reset() {
	*x = ObjectState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_states_proto_states_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectState) ProtoMessage() {}

func (x *ObjectState) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_states_proto_states_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectState.ProtoReflect.Descriptor instead.
func (*ObjectState) Descriptor() ([]byte, []int) {
	return file_pkg_states_proto_states_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectState) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ObjectState) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

var File_pkg_states_proto_states_proto protoreflect.FileDescriptor

var file_pkg_states_proto_states_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x13, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02,
	0x74, 0x73, 0x88, 0x01, 0x01, 0x1a, 0x4f, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x74, 0x73, 0x22, 0x4e, 0x0a,
	0x0b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x77, 0x0a,
	0x0c, 0x4e, 0x6f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50,
	0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x42, 0xa7, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x42, 0x0b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70,
	0x2f, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x4e, 0x53, 0x58, 0xaa,
	0x02, 0x0e, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0xca, 0x02, 0x0e, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0xe2, 0x02, 0x1a, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0f, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_states_proto_states_proto_rawDescOnce sync.Once
	file_pkg_states_proto_states_proto_rawDescData = file_pkg_states_proto_states_proto_rawDesc
)

func file_pkg_states_proto_states_proto_rawDescGZIP() []byte {
	file_pkg_states_proto_states_proto_rawDescOnce.Do(func() {
		file_pkg_states_proto_states_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_states_proto_states_proto_rawDescData)
	})
	return file_pkg_states_proto_states_proto_rawDescData
}

var file_pkg_states_proto_states_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_states_proto_states_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_states_proto_states_proto_goTypes = []interface{}{
	(NoCloudState)(0),      // 0: nocloud.states.NoCloudState
	(*State)(nil),          // 1: nocloud.states.State
	(*ObjectState)(nil),    // 2: nocloud.states.ObjectState
	nil,                    // 3: nocloud.states.State.MetaEntry
	(*structpb.Value)(nil), // 4: google.protobuf.Value
}
var file_pkg_states_proto_states_proto_depIdxs = []int32{
	0, // 0: nocloud.states.State.state:type_name -> nocloud.states.NoCloudState
	3, // 1: nocloud.states.State.meta:type_name -> nocloud.states.State.MetaEntry
	1, // 2: nocloud.states.ObjectState.state:type_name -> nocloud.states.State
	4, // 3: nocloud.states.State.MetaEntry.value:type_name -> google.protobuf.Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_states_proto_states_proto_init() }
func file_pkg_states_proto_states_proto_init() {
	if File_pkg_states_proto_states_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_states_proto_states_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_pkg_states_proto_states_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectState); i {
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
	file_pkg_states_proto_states_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_states_proto_states_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_states_proto_states_proto_goTypes,
		DependencyIndexes: file_pkg_states_proto_states_proto_depIdxs,
		EnumInfos:         file_pkg_states_proto_states_proto_enumTypes,
		MessageInfos:      file_pkg_states_proto_states_proto_msgTypes,
	}.Build()
	File_pkg_states_proto_states_proto = out.File
	file_pkg_states_proto_states_proto_rawDesc = nil
	file_pkg_states_proto_states_proto_goTypes = nil
	file_pkg_states_proto_states_proto_depIdxs = nil
}
