// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: avatar_service.proto

package avatar_go

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

type SetAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Filename string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *SetAvatarRequest) Reset() {
	*x = SetAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarRequest) ProtoMessage() {}

func (x *SetAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarRequest.ProtoReflect.Descriptor instead.
func (*SetAvatarRequest) Descriptor() ([]byte, []int) {
	return file_avatar_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetAvatarRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SetAvatarRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetAvatarRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type SetAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *SetAvatarResponse) Reset() {
	*x = SetAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarResponse) ProtoMessage() {}

func (x *SetAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarResponse.ProtoReflect.Descriptor instead.
func (*SetAvatarResponse) Descriptor() ([]byte, []int) {
	return file_avatar_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetAvatarResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SetAvatarResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SetAvatarResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type GetAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetAvatarRequest) Reset() {
	*x = GetAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarRequest) ProtoMessage() {}

func (x *GetAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarRequest.ProtoReflect.Descriptor instead.
func (*GetAvatarRequest) Descriptor() ([]byte, []int) {
	return file_avatar_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAvatarRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *GetAvatarResponse) Reset() {
	*x = GetAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarResponse) ProtoMessage() {}

func (x *GetAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarResponse.ProtoReflect.Descriptor instead.
func (*GetAvatarResponse) Descriptor() ([]byte, []int) {
	return file_avatar_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAvatarResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetAvatarResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetAvatarResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_avatar_service_proto protoreflect.FileDescriptor

var file_avatar_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x65,
	0x0a, 0x11, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x65, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x32, 0xb9, 0x01, 0x0a, 0x0d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x20, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x52, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x2e, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_avatar_service_proto_rawDescOnce sync.Once
	file_avatar_service_proto_rawDescData = file_avatar_service_proto_rawDesc
)

func file_avatar_service_proto_rawDescGZIP() []byte {
	file_avatar_service_proto_rawDescOnce.Do(func() {
		file_avatar_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_avatar_service_proto_rawDescData)
	})
	return file_avatar_service_proto_rawDescData
}

var file_avatar_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_avatar_service_proto_goTypes = []interface{}{
	(*SetAvatarRequest)(nil),  // 0: avatar_service.SetAvatarRequest
	(*SetAvatarResponse)(nil), // 1: avatar_service.SetAvatarResponse
	(*GetAvatarRequest)(nil),  // 2: avatar_service.GetAvatarRequest
	(*GetAvatarResponse)(nil), // 3: avatar_service.GetAvatarResponse
}
var file_avatar_service_proto_depIdxs = []int32{
	0, // 0: avatar_service.AvatarService.SetAvatar:input_type -> avatar_service.SetAvatarRequest
	2, // 1: avatar_service.AvatarService.GetAvatar:input_type -> avatar_service.GetAvatarRequest
	1, // 2: avatar_service.AvatarService.SetAvatar:output_type -> avatar_service.SetAvatarResponse
	3, // 3: avatar_service.AvatarService.GetAvatar:output_type -> avatar_service.GetAvatarResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_avatar_service_proto_init() }
func file_avatar_service_proto_init() {
	if File_avatar_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_avatar_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvatarRequest); i {
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
		file_avatar_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvatarResponse); i {
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
		file_avatar_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvatarRequest); i {
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
		file_avatar_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvatarResponse); i {
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
			RawDescriptor: file_avatar_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avatar_service_proto_goTypes,
		DependencyIndexes: file_avatar_service_proto_depIdxs,
		MessageInfos:      file_avatar_service_proto_msgTypes,
	}.Build()
	File_avatar_service_proto = out.File
	file_avatar_service_proto_rawDesc = nil
	file_avatar_service_proto_goTypes = nil
	file_avatar_service_proto_depIdxs = nil
}
