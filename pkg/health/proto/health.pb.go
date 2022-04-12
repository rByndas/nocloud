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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: pkg/health/proto/health.proto

package proto

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

type Status int32

const (
	Status_UNKNOWN Status = 0
	Status_SERVING Status = 1 // Service is up and running
	Status_OFFLINE Status = 2 // Service is offline(down)
	Status_RUNNING Status = 3 // Routine is running
	Status_STOPPED Status = 4 // Routine is stopped
	Status_INTENAL Status = 5 // Internal error while making status
	Status_HASERRS Status = 6 // Check has errors
	Status_NOEXIST Status = 7 // Service has no Routines
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "OFFLINE",
		3: "RUNNING",
		4: "STOPPED",
		5: "INTENAL",
		6: "HASERRS",
		7: "NOEXIST",
	}
	Status_value = map[string]int32{
		"UNKNOWN": 0,
		"SERVING": 1,
		"OFFLINE": 2,
		"RUNNING": 3,
		"STOPPED": 4,
		"INTENAL": 5,
		"HASERRS": 6,
		"NOEXIST": 7,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_health_proto_health_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_health_proto_health_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_health_proto_health_proto_rawDescGZIP(), []int{0}
}

type ProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProbeType string `protobuf:"bytes,1,opt,name=probe_type,json=probeType,proto3" json:"probe_type,omitempty"`
	Payload   string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ProbeRequest) Reset() {
	*x = ProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_health_proto_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeRequest) ProtoMessage() {}

func (x *ProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_health_proto_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeRequest.ProtoReflect.Descriptor instead.
func (*ProbeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_health_proto_health_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeRequest) GetProbeType() string {
	if x != nil {
		return x.ProbeType
	}
	return ""
}

func (x *ProbeRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string           `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Status   Status           `protobuf:"varint,2,opt,name=status,proto3,enum=nocloud.health.Status" json:"status,omitempty"`
	Serving  []*ServingStatus `protobuf:"bytes,3,rep,name=serving,proto3" json:"serving,omitempty"`
	Routines []*RoutineStatus `protobuf:"bytes,4,rep,name=routines,proto3" json:"routines,omitempty"`
}

func (x *ProbeResponse) Reset() {
	*x = ProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_health_proto_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeResponse) ProtoMessage() {}

func (x *ProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_health_proto_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeResponse.ProtoReflect.Descriptor instead.
func (*ProbeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_health_proto_health_proto_rawDescGZIP(), []int{1}
}

func (x *ProbeResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *ProbeResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *ProbeResponse) GetServing() []*ServingStatus {
	if x != nil {
		return x.Serving
	}
	return nil
}

func (x *ProbeResponse) GetRoutines() []*RoutineStatus {
	if x != nil {
		return x.Routines
	}
	return nil
}

type ServingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string  `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Status  Status  `protobuf:"varint,2,opt,name=status,proto3,enum=nocloud.health.Status" json:"status,omitempty"`
	Error   *string `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *ServingStatus) Reset() {
	*x = ServingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_health_proto_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServingStatus) ProtoMessage() {}

func (x *ServingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_health_proto_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServingStatus.ProtoReflect.Descriptor instead.
func (*ServingStatus) Descriptor() ([]byte, []int) {
	return file_pkg_health_proto_health_proto_rawDescGZIP(), []int{2}
}

func (x *ServingStatus) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ServingStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *ServingStatus) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type RoutineStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routine       string         `protobuf:"bytes,1,opt,name=routine,proto3" json:"routine,omitempty"`
	Status        *ServingStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	LastExecution string         `protobuf:"bytes,3,opt,name=last_execution,json=lastExecution,proto3" json:"last_execution,omitempty"`
}

func (x *RoutineStatus) Reset() {
	*x = RoutineStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_health_proto_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutineStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutineStatus) ProtoMessage() {}

func (x *RoutineStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_health_proto_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutineStatus.ProtoReflect.Descriptor instead.
func (*RoutineStatus) Descriptor() ([]byte, []int) {
	return file_pkg_health_proto_health_proto_rawDescGZIP(), []int{3}
}

func (x *RoutineStatus) GetRoutine() string {
	if x != nil {
		return x.Routine
	}
	return ""
}

func (x *RoutineStatus) GetStatus() *ServingStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RoutineStatus) GetLastExecution() string {
	if x != nil {
		return x.LastExecution
	}
	return ""
}

type RoutinesStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routines []*RoutineStatus `protobuf:"bytes,1,rep,name=routines,proto3" json:"routines,omitempty"`
}

func (x *RoutinesStatus) Reset() {
	*x = RoutinesStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_health_proto_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutinesStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutinesStatus) ProtoMessage() {}

func (x *RoutinesStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_health_proto_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutinesStatus.ProtoReflect.Descriptor instead.
func (*RoutinesStatus) Descriptor() ([]byte, []int) {
	return file_pkg_health_proto_health_proto_rawDescGZIP(), []int{4}
}

func (x *RoutinesStatus) GetRoutines() []*RoutineStatus {
	if x != nil {
		return x.Routines
	}
	return nil
}

var File_pkg_health_proto_health_proto protoreflect.FileDescriptor

var file_pkg_health_proto_health_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a,
	0x08, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x2a,
	0x70, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e,
	0x54, 0x45, 0x4e, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x41, 0x53, 0x45, 0x52,
	0x52, 0x53, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10,
	0x07, 0x32, 0x6f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5e, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x6f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x3a,
	0x01, 0x2a, 0x32, 0xa7, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x47, 0x0a, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x12, 0x1c,
	0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e,
	0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0xa7, 0x01, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x42, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2,
	0x02, 0x03, 0x4e, 0x48, 0x58, 0xaa, 0x02, 0x0e, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0xca, 0x02, 0x0e, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0xe2, 0x02, 0x1a, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_health_proto_health_proto_rawDescOnce sync.Once
	file_pkg_health_proto_health_proto_rawDescData = file_pkg_health_proto_health_proto_rawDesc
)

func file_pkg_health_proto_health_proto_rawDescGZIP() []byte {
	file_pkg_health_proto_health_proto_rawDescOnce.Do(func() {
		file_pkg_health_proto_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_health_proto_health_proto_rawDescData)
	})
	return file_pkg_health_proto_health_proto_rawDescData
}

var file_pkg_health_proto_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_health_proto_health_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_health_proto_health_proto_goTypes = []interface{}{
	(Status)(0),            // 0: nocloud.health.Status
	(*ProbeRequest)(nil),   // 1: nocloud.health.ProbeRequest
	(*ProbeResponse)(nil),  // 2: nocloud.health.ProbeResponse
	(*ServingStatus)(nil),  // 3: nocloud.health.ServingStatus
	(*RoutineStatus)(nil),  // 4: nocloud.health.RoutineStatus
	(*RoutinesStatus)(nil), // 5: nocloud.health.RoutinesStatus
}
var file_pkg_health_proto_health_proto_depIdxs = []int32{
	0, // 0: nocloud.health.ProbeResponse.status:type_name -> nocloud.health.Status
	3, // 1: nocloud.health.ProbeResponse.serving:type_name -> nocloud.health.ServingStatus
	4, // 2: nocloud.health.ProbeResponse.routines:type_name -> nocloud.health.RoutineStatus
	0, // 3: nocloud.health.ServingStatus.status:type_name -> nocloud.health.Status
	3, // 4: nocloud.health.RoutineStatus.status:type_name -> nocloud.health.ServingStatus
	4, // 5: nocloud.health.RoutinesStatus.routines:type_name -> nocloud.health.RoutineStatus
	1, // 6: nocloud.health.HealthService.Probe:input_type -> nocloud.health.ProbeRequest
	1, // 7: nocloud.health.InternalProbeService.Service:input_type -> nocloud.health.ProbeRequest
	1, // 8: nocloud.health.InternalProbeService.Routine:input_type -> nocloud.health.ProbeRequest
	2, // 9: nocloud.health.HealthService.Probe:output_type -> nocloud.health.ProbeResponse
	3, // 10: nocloud.health.InternalProbeService.Service:output_type -> nocloud.health.ServingStatus
	5, // 11: nocloud.health.InternalProbeService.Routine:output_type -> nocloud.health.RoutinesStatus
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_health_proto_health_proto_init() }
func file_pkg_health_proto_health_proto_init() {
	if File_pkg_health_proto_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_health_proto_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeRequest); i {
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
		file_pkg_health_proto_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeResponse); i {
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
		file_pkg_health_proto_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServingStatus); i {
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
		file_pkg_health_proto_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutineStatus); i {
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
		file_pkg_health_proto_health_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutinesStatus); i {
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
	file_pkg_health_proto_health_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_health_proto_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_health_proto_health_proto_goTypes,
		DependencyIndexes: file_pkg_health_proto_health_proto_depIdxs,
		EnumInfos:         file_pkg_health_proto_health_proto_enumTypes,
		MessageInfos:      file_pkg_health_proto_health_proto_msgTypes,
	}.Build()
	File_pkg_health_proto_health_proto = out.File
	file_pkg_health_proto_health_proto_rawDesc = nil
	file_pkg_health_proto_health_proto_goTypes = nil
	file_pkg_health_proto_health_proto_depIdxs = nil
}
