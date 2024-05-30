// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: app/videocollection/service/enum/video_collection_enum.proto

package enum

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

type ContentType int32

const (
	ContentType_TextImage      ContentType = 0 // 图文
	ContentType_LandscapeVideo ContentType = 1 // 竖版短视频
	ContentType_PortraitVideo  ContentType = 2 // 横版短视频
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0: "TextImage",
		1: "LandscapeVideo",
		2: "PortraitVideo",
	}
	ContentType_value = map[string]int32{
		"TextImage":      0,
		"LandscapeVideo": 1,
		"PortraitVideo":  2,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_app_videocollection_service_enum_video_collection_enum_proto_enumTypes[0].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_app_videocollection_service_enum_video_collection_enum_proto_enumTypes[0]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_app_videocollection_service_enum_video_collection_enum_proto_rawDescGZIP(), []int{0}
}

type FilterType int32

const (
	FilterType_Ruled  FilterType = 0 // 规则筛选
	FilterType_Manual FilterType = 1 // 人工
)

// Enum value maps for FilterType.
var (
	FilterType_name = map[int32]string{
		0: "Ruled",
		1: "Manual",
	}
	FilterType_value = map[string]int32{
		"Ruled":  0,
		"Manual": 1,
	}
)

func (x FilterType) Enum() *FilterType {
	p := new(FilterType)
	*p = x
	return p
}

func (x FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_app_videocollection_service_enum_video_collection_enum_proto_enumTypes[1].Descriptor()
}

func (FilterType) Type() protoreflect.EnumType {
	return &file_app_videocollection_service_enum_video_collection_enum_proto_enumTypes[1]
}

func (x FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterType.Descriptor instead.
func (FilterType) EnumDescriptor() ([]byte, []int) {
	return file_app_videocollection_service_enum_video_collection_enum_proto_rawDescGZIP(), []int{1}
}

var File_app_videocollection_service_enum_video_collection_enum_proto protoreflect.FileDescriptor

var file_app_videocollection_service_enum_video_collection_enum_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x43, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x65, 0x78, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x6f, 0x72, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x02, 0x2a, 0x23, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65,
	0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x10, 0x01, 0x42,
	0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x73, 0x6c, 0x65, 0x79, 0x77, 0x75, 0x2f, 0x72, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_videocollection_service_enum_video_collection_enum_proto_rawDescOnce sync.Once
	file_app_videocollection_service_enum_video_collection_enum_proto_rawDescData = file_app_videocollection_service_enum_video_collection_enum_proto_rawDesc
)

func file_app_videocollection_service_enum_video_collection_enum_proto_rawDescGZIP() []byte {
	file_app_videocollection_service_enum_video_collection_enum_proto_rawDescOnce.Do(func() {
		file_app_videocollection_service_enum_video_collection_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_videocollection_service_enum_video_collection_enum_proto_rawDescData)
	})
	return file_app_videocollection_service_enum_video_collection_enum_proto_rawDescData
}

var file_app_videocollection_service_enum_video_collection_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_app_videocollection_service_enum_video_collection_enum_proto_goTypes = []interface{}{
	(ContentType)(0), // 0: proto.ContentType
	(FilterType)(0),  // 1: proto.FilterType
}
var file_app_videocollection_service_enum_video_collection_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_videocollection_service_enum_video_collection_enum_proto_init() }
func file_app_videocollection_service_enum_video_collection_enum_proto_init() {
	if File_app_videocollection_service_enum_video_collection_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_videocollection_service_enum_video_collection_enum_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_videocollection_service_enum_video_collection_enum_proto_goTypes,
		DependencyIndexes: file_app_videocollection_service_enum_video_collection_enum_proto_depIdxs,
		EnumInfos:         file_app_videocollection_service_enum_video_collection_enum_proto_enumTypes,
	}.Build()
	File_app_videocollection_service_enum_video_collection_enum_proto = out.File
	file_app_videocollection_service_enum_video_collection_enum_proto_rawDesc = nil
	file_app_videocollection_service_enum_video_collection_enum_proto_goTypes = nil
	file_app_videocollection_service_enum_video_collection_enum_proto_depIdxs = nil
}