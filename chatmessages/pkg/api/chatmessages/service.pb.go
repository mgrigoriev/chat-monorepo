// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: api/service.proto

package chatmessages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_service_proto protoreflect.FileDescriptor

var file_api_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb6, 0x04, 0x0a, 0x13, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa8, 0x01, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x48, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x49, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72,
	0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e,
	0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb9, 0x01, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x50, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xb7, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x4f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x67, 0x72, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x76, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x63,
	0x68, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_api_service_proto_goTypes = []interface{}{
	(*SaveChatMessageRequest)(nil),         // 0: github.com.mgrigoriev.chat_monorepo.chatmessages.SaveChatMessageRequest
	(*ListPrivateChatMessagesRequest)(nil), // 1: github.com.mgrigoriev.chat_monorepo.chatmessages.ListPrivateChatMessagesRequest
	(*ListServerChatMessagesRequest)(nil),  // 2: github.com.mgrigoriev.chat_monorepo.chatmessages.ListServerChatMessagesRequest
	(*SaveChatMessageResponse)(nil),        // 3: github.com.mgrigoriev.chat_monorepo.chatmessages.SaveChatMessageResponse
	(*ListChatMessagesResponse)(nil),       // 4: github.com.mgrigoriev.chat_monorepo.chatmessages.ListChatMessagesResponse
}
var file_api_service_proto_depIdxs = []int32{
	0, // 0: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessagesService.SaveChatMessage:input_type -> github.com.mgrigoriev.chat_monorepo.chatmessages.SaveChatMessageRequest
	1, // 1: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessagesService.ListPrivateChatMessages:input_type -> github.com.mgrigoriev.chat_monorepo.chatmessages.ListPrivateChatMessagesRequest
	2, // 2: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessagesService.ListServerChatMessages:input_type -> github.com.mgrigoriev.chat_monorepo.chatmessages.ListServerChatMessagesRequest
	3, // 3: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessagesService.SaveChatMessage:output_type -> github.com.mgrigoriev.chat_monorepo.chatmessages.SaveChatMessageResponse
	4, // 4: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessagesService.ListPrivateChatMessages:output_type -> github.com.mgrigoriev.chat_monorepo.chatmessages.ListChatMessagesResponse
	4, // 5: github.com.mgrigoriev.chat_monorepo.chatmessages.ChatMessagesService.ListServerChatMessages:output_type -> github.com.mgrigoriev.chat_monorepo.chatmessages.ListChatMessagesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_service_proto_init() }
func file_api_service_proto_init() {
	if File_api_service_proto != nil {
		return
	}
	file_api_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_proto_goTypes,
		DependencyIndexes: file_api_service_proto_depIdxs,
	}.Build()
	File_api_service_proto = out.File
	file_api_service_proto_rawDesc = nil
	file_api_service_proto_goTypes = nil
	file_api_service_proto_depIdxs = nil
}
