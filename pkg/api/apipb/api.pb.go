// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: pkg/api/apipb/api.proto

package apipb

import (
	healthpb "github.com/slntopp/ione-go/pkg/health/healthpb"
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

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Endpoint   string `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ExpireTime string `protobuf:"bytes,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_apipb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_apipb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_apipb_api_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *TokenRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *TokenRequest) GetExpireTime() string {
	if x != nil {
		return x.ExpireTime
	}
	return ""
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_apipb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_apipb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_apipb_api_proto_rawDescGZIP(), []int{1}
}

func (x *TokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_pkg_api_apipb_api_proto protoreflect.FileDescriptor

var file_pkg_api_apipb_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x32, 0x48, 0x0a, 0x0a, 0x4f, 0x4e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x69, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x51, 0x0a, 0x0d,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a,
	0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x69, 0x6f, 0x6e, 0x65, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x6f, 0x6e, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c,
	0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x69, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_apipb_api_proto_rawDescOnce sync.Once
	file_pkg_api_apipb_api_proto_rawDescData = file_pkg_api_apipb_api_proto_rawDesc
)

func file_pkg_api_apipb_api_proto_rawDescGZIP() []byte {
	file_pkg_api_apipb_api_proto_rawDescOnce.Do(func() {
		file_pkg_api_apipb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_apipb_api_proto_rawDescData)
	})
	return file_pkg_api_apipb_api_proto_rawDescData
}

var file_pkg_api_apipb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_api_apipb_api_proto_goTypes = []interface{}{
	(*TokenRequest)(nil),           // 0: ione.api.TokenRequest
	(*TokenResponse)(nil),          // 1: ione.api.TokenResponse
	(*healthpb.ProbeRequest)(nil),  // 2: ione.health.ProbeRequest
	(*healthpb.ProbeResponse)(nil), // 3: ione.health.ProbeResponse
}
var file_pkg_api_apipb_api_proto_depIdxs = []int32{
	0, // 0: ione.api.ONeService.Token:input_type -> ione.api.TokenRequest
	2, // 1: ione.api.HealthService.Probe:input_type -> ione.health.ProbeRequest
	1, // 2: ione.api.ONeService.Token:output_type -> ione.api.TokenResponse
	3, // 3: ione.api.HealthService.Probe:output_type -> ione.health.ProbeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_apipb_api_proto_init() }
func file_pkg_api_apipb_api_proto_init() {
	if File_pkg_api_apipb_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_apipb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
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
		file_pkg_api_apipb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
			RawDescriptor: file_pkg_api_apipb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_api_apipb_api_proto_goTypes,
		DependencyIndexes: file_pkg_api_apipb_api_proto_depIdxs,
		MessageInfos:      file_pkg_api_apipb_api_proto_msgTypes,
	}.Build()
	File_pkg_api_apipb_api_proto = out.File
	file_pkg_api_apipb_api_proto_rawDesc = nil
	file_pkg_api_apipb_api_proto_goTypes = nil
	file_pkg_api_apipb_api_proto_depIdxs = nil
}
