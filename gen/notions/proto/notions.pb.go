// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/notions.proto

package notions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionRequest) Reset() {
	*x = NotionRequest{}
	mi := &file_proto_notions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionRequest) ProtoMessage() {}

func (x *NotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionRequest.ProtoReflect.Descriptor instead.
func (*NotionRequest) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{0}
}

type NotionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionResponse) Reset() {
	*x = NotionResponse{}
	mi := &file_proto_notions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionResponse) ProtoMessage() {}

func (x *NotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionResponse.ProtoReflect.Descriptor instead.
func (*NotionResponse) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{1}
}

type NotionRequest_Get struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionRequest_Get) Reset() {
	*x = NotionRequest_Get{}
	mi := &file_proto_notions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionRequest_Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionRequest_Get) ProtoMessage() {}

func (x *NotionRequest_Get) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionRequest_Get.ProtoReflect.Descriptor instead.
func (*NotionRequest_Get) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{0, 0}
}

type NotionRequest_Add struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionRequest_Add) Reset() {
	*x = NotionRequest_Add{}
	mi := &file_proto_notions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionRequest_Add) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionRequest_Add) ProtoMessage() {}

func (x *NotionRequest_Add) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionRequest_Add.ProtoReflect.Descriptor instead.
func (*NotionRequest_Add) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{0, 1}
}

func (x *NotionRequest_Add) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotionRequest_Add) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type NotionRequest_Update struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Id            int64                  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionRequest_Update) Reset() {
	*x = NotionRequest_Update{}
	mi := &file_proto_notions_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionRequest_Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionRequest_Update) ProtoMessage() {}

func (x *NotionRequest_Update) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionRequest_Update.ProtoReflect.Descriptor instead.
func (*NotionRequest_Update) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{0, 2}
}

func (x *NotionRequest_Update) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotionRequest_Update) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NotionRequest_Update) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NotionRequest_Delete struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionRequest_Delete) Reset() {
	*x = NotionRequest_Delete{}
	mi := &file_proto_notions_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionRequest_Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionRequest_Delete) ProtoMessage() {}

func (x *NotionRequest_Delete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionRequest_Delete.ProtoReflect.Descriptor instead.
func (*NotionRequest_Delete) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{0, 3}
}

func (x *NotionRequest_Delete) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NotionResponse_Get struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Id            int64                  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionResponse_Get) Reset() {
	*x = NotionResponse_Get{}
	mi := &file_proto_notions_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionResponse_Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionResponse_Get) ProtoMessage() {}

func (x *NotionResponse_Get) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionResponse_Get.ProtoReflect.Descriptor instead.
func (*NotionResponse_Get) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NotionResponse_Get) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotionResponse_Get) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NotionResponse_Get) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *NotionResponse_Get) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NotionResponse_Add struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionResponse_Add) Reset() {
	*x = NotionResponse_Add{}
	mi := &file_proto_notions_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionResponse_Add) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionResponse_Add) ProtoMessage() {}

func (x *NotionResponse_Add) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionResponse_Add.ProtoReflect.Descriptor instead.
func (*NotionResponse_Add) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{1, 1}
}

func (x *NotionResponse_Add) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NotionResponse_Update struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionResponse_Update) Reset() {
	*x = NotionResponse_Update{}
	mi := &file_proto_notions_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionResponse_Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionResponse_Update) ProtoMessage() {}

func (x *NotionResponse_Update) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionResponse_Update.ProtoReflect.Descriptor instead.
func (*NotionResponse_Update) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{1, 2}
}

func (x *NotionResponse_Update) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NotionResponse_Delete struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotionResponse_Delete) Reset() {
	*x = NotionResponse_Delete{}
	mi := &file_proto_notions_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotionResponse_Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionResponse_Delete) ProtoMessage() {}

func (x *NotionResponse_Delete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notions_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionResponse_Delete.ProtoReflect.Descriptor instead.
func (*NotionResponse_Delete) Descriptor() ([]byte, []int) {
	return file_proto_notions_proto_rawDescGZIP(), []int{1, 3}
}

func (x *NotionResponse_Delete) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_notions_proto protoreflect.FileDescriptor

const file_proto_notions_proto_rawDesc = "" +
	"\n" +
	"\x13proto/notions.proto\"\xb1\x01\n" +
	"\rNotionRequest\x1a\x05\n" +
	"\x03Get\x1a5\n" +
	"\x03Add\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x1aH\n" +
	"\x06Update\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x0e\n" +
	"\x02id\x18\x03 \x01(\x03R\x02id\x1a\x18\n" +
	"\x06Delete\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\xba\x01\n" +
	"\x0eNotionResponse\x1a]\n" +
	"\x03Get\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x16\n" +
	"\x06author\x18\x03 \x01(\tR\x06author\x12\x0e\n" +
	"\x02id\x18\x04 \x01(\x03R\x02id\x1a\x15\n" +
	"\x03Add\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x1a\x18\n" +
	"\x06Update\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x1a\x18\n" +
	"\x06Delete\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id2\xdd\x01\n" +
	"\aNotions\x120\n" +
	"\x03Get\x12\x12.NotionRequest.Get\x1a\x13.NotionResponse.Get0\x01\x12.\n" +
	"\x03Add\x12\x12.NotionRequest.Add\x1a\x13.NotionResponse.Add\x127\n" +
	"\x06Update\x12\x15.NotionRequest.Update\x1a\x16.NotionResponse.Update\x127\n" +
	"\x06Delete\x12\x15.NotionRequest.Delete\x1a\x16.NotionResponse.DeleteB\n" +
	"Z\bnotions/b\x06proto3"

var (
	file_proto_notions_proto_rawDescOnce sync.Once
	file_proto_notions_proto_rawDescData []byte
)

func file_proto_notions_proto_rawDescGZIP() []byte {
	file_proto_notions_proto_rawDescOnce.Do(func() {
		file_proto_notions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_notions_proto_rawDesc), len(file_proto_notions_proto_rawDesc)))
	})
	return file_proto_notions_proto_rawDescData
}

var file_proto_notions_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_notions_proto_goTypes = []any{
	(*NotionRequest)(nil),         // 0: NotionRequest
	(*NotionResponse)(nil),        // 1: NotionResponse
	(*NotionRequest_Get)(nil),     // 2: NotionRequest.Get
	(*NotionRequest_Add)(nil),     // 3: NotionRequest.Add
	(*NotionRequest_Update)(nil),  // 4: NotionRequest.Update
	(*NotionRequest_Delete)(nil),  // 5: NotionRequest.Delete
	(*NotionResponse_Get)(nil),    // 6: NotionResponse.Get
	(*NotionResponse_Add)(nil),    // 7: NotionResponse.Add
	(*NotionResponse_Update)(nil), // 8: NotionResponse.Update
	(*NotionResponse_Delete)(nil), // 9: NotionResponse.Delete
}
var file_proto_notions_proto_depIdxs = []int32{
	2, // 0: Notions.Get:input_type -> NotionRequest.Get
	3, // 1: Notions.Add:input_type -> NotionRequest.Add
	4, // 2: Notions.Update:input_type -> NotionRequest.Update
	5, // 3: Notions.Delete:input_type -> NotionRequest.Delete
	6, // 4: Notions.Get:output_type -> NotionResponse.Get
	7, // 5: Notions.Add:output_type -> NotionResponse.Add
	8, // 6: Notions.Update:output_type -> NotionResponse.Update
	9, // 7: Notions.Delete:output_type -> NotionResponse.Delete
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_notions_proto_init() }
func file_proto_notions_proto_init() {
	if File_proto_notions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_notions_proto_rawDesc), len(file_proto_notions_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notions_proto_goTypes,
		DependencyIndexes: file_proto_notions_proto_depIdxs,
		MessageInfos:      file_proto_notions_proto_msgTypes,
	}.Build()
	File_proto_notions_proto = out.File
	file_proto_notions_proto_goTypes = nil
	file_proto_notions_proto_depIdxs = nil
}
