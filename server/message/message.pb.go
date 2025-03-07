// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.14.0
// source: message.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//服务器类型
type SERVICE int32

const (
	SERVICE_NONE          SERVICE = 0
	SERVICE_CLIENT        SERVICE = 1
	SERVICE_GATESERVER    SERVICE = 2 //网关,转发服务
	SERVICE_ACCOUNTSERVER SERVICE = 3 //账号
	SERVICE_WORLDSERVER   SERVICE = 4 //世界
	SERVICE_ZONESERVER    SERVICE = 5 //地图
	SERVICE_WORLDDBSERVER SERVICE = 6 //db
)

// Enum value maps for SERVICE.
var (
	SERVICE_name = map[int32]string{
		0: "NONE",
		1: "CLIENT",
		2: "GATESERVER",
		3: "ACCOUNTSERVER",
		4: "WORLDSERVER",
		5: "ZONESERVER",
		6: "WORLDDBSERVER",
	}
	SERVICE_value = map[string]int32{
		"NONE":          0,
		"CLIENT":        1,
		"GATESERVER":    2,
		"ACCOUNTSERVER": 3,
		"WORLDSERVER":   4,
		"ZONESERVER":    5,
		"WORLDDBSERVER": 6,
	}
)

func (x SERVICE) Enum() *SERVICE {
	p := new(SERVICE)
	*p = x
	return p
}

func (x SERVICE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SERVICE) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (SERVICE) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x SERVICE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SERVICE.Descriptor instead.
func (SERVICE) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type CHAT int32

const (
	CHAT_MSG_TYPE_WORLD   CHAT = 0
	CHAT_MSG_TYPE_PRIVATE CHAT = 1
	CHAT_MSG_TYPE_ORG     CHAT = 2
	CHAT_MSG_TYPE_COUNT   CHAT = 3
)

// Enum value maps for CHAT.
var (
	CHAT_name = map[int32]string{
		0: "MSG_TYPE_WORLD",
		1: "MSG_TYPE_PRIVATE",
		2: "MSG_TYPE_ORG",
		3: "MSG_TYPE_COUNT",
	}
	CHAT_value = map[string]int32{
		"MSG_TYPE_WORLD":   0,
		"MSG_TYPE_PRIVATE": 1,
		"MSG_TYPE_ORG":     2,
		"MSG_TYPE_COUNT":   3,
	}
)

func (x CHAT) Enum() *CHAT {
	p := new(CHAT)
	*p = x
	return p
}

func (x CHAT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CHAT) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (CHAT) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x CHAT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CHAT.Descriptor instead.
func (CHAT) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

type PlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID   int64  `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	PlayerName string `protobuf:"bytes,2,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	PlayerGold int32  `protobuf:"varint,3,opt,name=PlayerGold,proto3" json:"PlayerGold,omitempty"`
}

func (x *PlayerData) Reset() {
	*x = PlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerData) ProtoMessage() {}

func (x *PlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerData.ProtoReflect.Descriptor instead.
func (*PlayerData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerData) GetPlayerID() int64 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *PlayerData) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerData) GetPlayerGold() int32 {
	if x != nil {
		return x.PlayerGold
	}
	return 0
}

//客户端包头
type Ipacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stx            int32   `protobuf:"varint,1,opt,name=Stx,proto3" json:"Stx,omitempty"`
	DestServerType SERVICE `protobuf:"varint,2,opt,name=DestServerType,proto3,enum=message.SERVICE" json:"DestServerType,omitempty"`
	Ckx            int32   `protobuf:"varint,3,opt,name=Ckx,proto3" json:"Ckx,omitempty"`
	Id             int64   `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Ipacket) Reset() {
	*x = Ipacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipacket) ProtoMessage() {}

func (x *Ipacket) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipacket.ProtoReflect.Descriptor instead.
func (*Ipacket) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Ipacket) GetStx() int32 {
	if x != nil {
		return x.Stx
	}
	return 0
}

func (x *Ipacket) GetDestServerType() SERVICE {
	if x != nil {
		return x.DestServerType
	}
	return SERVICE_NONE
}

func (x *Ipacket) GetCkx() int32 {
	if x != nil {
		return x.Ckx
	}
	return 0
}

func (x *Ipacket) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

//心跳包
type HeardPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeardPacket) Reset() {
	*x = HeardPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeardPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeardPacket) ProtoMessage() {}

func (x *HeardPacket) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeardPacket.ProtoReflect.Descriptor instead.
func (*HeardPacket) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x68, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x6f, 0x6c, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x6f,
	0x6c, 0x64, 0x22, 0x77, 0x0a, 0x07, 0x49, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x74, 0x78, 0x12,
	0x38, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x52, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6b, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x43, 0x6b, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x72, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2a, 0x76, 0x0a, 0x07, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x47,
	0x41, 0x54, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0f,
	0x0a, 0x0b, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x04, 0x12,
	0x0e, 0x0a, 0x0a, 0x5a, 0x4f, 0x4e, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12,
	0x11, 0x0a, 0x0d, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x44, 0x42, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x10, 0x06, 0x2a, 0x56, 0x0a, 0x04, 0x43, 0x48, 0x41, 0x54, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x53,
	0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4f, 0x52, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_proto_goTypes = []interface{}{
	(SERVICE)(0),        // 0: message.SERVICE
	(CHAT)(0),           // 1: message.CHAT
	(*PlayerData)(nil),  // 2: message.PlayerData
	(*Ipacket)(nil),     // 3: message.Ipacket
	(*HeardPacket)(nil), // 4: message.HeardPacket
}
var file_message_proto_depIdxs = []int32{
	0, // 0: message.Ipacket.DestServerType:type_name -> message.SERVICE
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerData); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipacket); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeardPacket); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
