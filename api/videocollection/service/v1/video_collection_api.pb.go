// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/videocollection/service/v1/video_collection_api.proto

package v1

import (
	_ "github.com/castbox/go-guru/pkg/goguru/annotations"
	_ "github.com/castbox/go-guru/pkg/goguru/query"
	_ "github.com/castbox/go-guru/pkg/goguru/types"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	proto "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_videocollection_service_v1_video_collection_api_proto protoreflect.FileDescriptor

var file_api_videocollection_service_v1_video_collection_api_proto_rawDesc = []byte{
	0x0a, 0x39, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x67, 0x6f, 0x67, 0x75, 0x72, 0x75, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x67, 0x75, 0x72, 0x75, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x67, 0x6f, 0x67, 0x75, 0x72, 0x75, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x67, 0x75, 0x72, 0x75, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x61, 0x70, 0x70, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfa, 0x08, 0x0a, 0x0f, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8a, 0x01, 0x0a,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x41, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01,
	0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x82, 0xc9, 0x90,
	0xbe, 0x02, 0x16, 0x12, 0x0f, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x03, 0x33, 0x30, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x03, 0x4f, 0x6e,
	0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x22, 0x3f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f,
	0x6e, 0x65, 0x82, 0xc9, 0x90, 0xbe, 0x02, 0x16, 0x12, 0x0f, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x03, 0x33, 0x30, 0x73, 0x12, 0x86,
	0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a,
	0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x82, 0xc9, 0x90, 0xbe, 0x02,
	0x16, 0x12, 0x0f, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x03, 0x33, 0x30, 0x73, 0x12, 0x6a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x82, 0xc9, 0x90,
	0xbe, 0x02, 0x00, 0x12, 0x71, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x82, 0xc9, 0x90, 0xbe, 0x02, 0x00, 0x12, 0x76, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x32, 0x19, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x82, 0xc9, 0x90, 0xbe, 0x02, 0x00, 0x12, 0x76,
	0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x1a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x82, 0xc9, 0x90, 0xbe, 0x02, 0x00, 0x12, 0x73, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x82, 0xc9, 0x90, 0xbe, 0x02, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65,
	0x71, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x73, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a,
	0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x82,
	0xc9, 0x90, 0xbe, 0x02, 0x00, 0x42, 0x7a, 0x0a, 0x18, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x42, 0x11, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x73, 0x6c, 0x65, 0x79, 0x77, 0x75, 0x2f, 0x72, 0x69, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_videocollection_service_v1_video_collection_api_proto_goTypes = []interface{}{
	(*proto.VideoCollectionCountReq)(nil),       // 0: proto.VideoCollectionCountReq
	(*proto.VideoCollectionOneReq)(nil),         // 1: proto.VideoCollectionOneReq
	(*proto.VideoCollectionListReq)(nil),        // 2: proto.VideoCollectionListReq
	(*proto.VideoCollectionGetReq)(nil),         // 3: proto.VideoCollectionGetReq
	(*proto.VideoCollectionCreateReq)(nil),      // 4: proto.VideoCollectionCreateReq
	(*proto.VideoCollectionUpdateReq)(nil),      // 5: proto.VideoCollectionUpdateReq
	(*proto.VideoCollectionUpsertReq)(nil),      // 6: proto.VideoCollectionUpsertReq
	(*proto.VideoCollectionDeleteReq)(nil),      // 7: proto.VideoCollectionDeleteReq
	(*proto.VideoCollectionDeleteMultiReq)(nil), // 8: proto.VideoCollectionDeleteMultiReq
	(*proto.VideoCollectionCountRes)(nil),       // 9: proto.VideoCollectionCountRes
	(*proto.VideoCollectionOneRes)(nil),         // 10: proto.VideoCollectionOneRes
	(*proto.VideoCollectionListRes)(nil),        // 11: proto.VideoCollectionListRes
	(*proto.VideoCollectionGetRes)(nil),         // 12: proto.VideoCollectionGetRes
	(*proto.VideoCollectionCreateRes)(nil),      // 13: proto.VideoCollectionCreateRes
	(*proto.VideoCollectionUpdateRes)(nil),      // 14: proto.VideoCollectionUpdateRes
	(*proto.VideoCollectionUpsertRes)(nil),      // 15: proto.VideoCollectionUpsertRes
	(*proto.VideoCollectionDeleteRes)(nil),      // 16: proto.VideoCollectionDeleteRes
	(*proto.VideoCollectionDeleteMultiRes)(nil), // 17: proto.VideoCollectionDeleteMultiRes
}
var file_api_videocollection_service_v1_video_collection_api_proto_depIdxs = []int32{
	0,  // 0: v1.VideoCollection.Count:input_type -> proto.VideoCollectionCountReq
	1,  // 1: v1.VideoCollection.One:input_type -> proto.VideoCollectionOneReq
	2,  // 2: v1.VideoCollection.List:input_type -> proto.VideoCollectionListReq
	3,  // 3: v1.VideoCollection.Get:input_type -> proto.VideoCollectionGetReq
	4,  // 4: v1.VideoCollection.Create:input_type -> proto.VideoCollectionCreateReq
	5,  // 5: v1.VideoCollection.Update:input_type -> proto.VideoCollectionUpdateReq
	6,  // 6: v1.VideoCollection.Upsert:input_type -> proto.VideoCollectionUpsertReq
	7,  // 7: v1.VideoCollection.Delete:input_type -> proto.VideoCollectionDeleteReq
	8,  // 8: v1.VideoCollection.DeleteMulti:input_type -> proto.VideoCollectionDeleteMultiReq
	9,  // 9: v1.VideoCollection.Count:output_type -> proto.VideoCollectionCountRes
	10, // 10: v1.VideoCollection.One:output_type -> proto.VideoCollectionOneRes
	11, // 11: v1.VideoCollection.List:output_type -> proto.VideoCollectionListRes
	12, // 12: v1.VideoCollection.Get:output_type -> proto.VideoCollectionGetRes
	13, // 13: v1.VideoCollection.Create:output_type -> proto.VideoCollectionCreateRes
	14, // 14: v1.VideoCollection.Update:output_type -> proto.VideoCollectionUpdateRes
	15, // 15: v1.VideoCollection.Upsert:output_type -> proto.VideoCollectionUpsertRes
	16, // 16: v1.VideoCollection.Delete:output_type -> proto.VideoCollectionDeleteRes
	17, // 17: v1.VideoCollection.DeleteMulti:output_type -> proto.VideoCollectionDeleteMultiRes
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_videocollection_service_v1_video_collection_api_proto_init() }
func file_api_videocollection_service_v1_video_collection_api_proto_init() {
	if File_api_videocollection_service_v1_video_collection_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_videocollection_service_v1_video_collection_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_videocollection_service_v1_video_collection_api_proto_goTypes,
		DependencyIndexes: file_api_videocollection_service_v1_video_collection_api_proto_depIdxs,
	}.Build()
	File_api_videocollection_service_v1_video_collection_api_proto = out.File
	file_api_videocollection_service_v1_video_collection_api_proto_rawDesc = nil
	file_api_videocollection_service_v1_video_collection_api_proto_goTypes = nil
	file_api_videocollection_service_v1_video_collection_api_proto_depIdxs = nil
}