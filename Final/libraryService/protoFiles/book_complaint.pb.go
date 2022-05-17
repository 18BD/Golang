// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: protoFiles/book_complaint.proto

package libraryService

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

// Model
type BookComplaint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BookID    int32  `protobuf:"varint,2,opt,name=BookID,proto3" json:"BookID,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Complaint string `protobuf:"bytes,4,opt,name=Complaint,proto3" json:"Complaint,omitempty"`
}

func (x *BookComplaint) Reset() {
	*x = BookComplaint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_book_complaint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookComplaint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookComplaint) ProtoMessage() {}

func (x *BookComplaint) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_book_complaint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookComplaint.ProtoReflect.Descriptor instead.
func (*BookComplaint) Descriptor() ([]byte, []int) {
	return file_protoFiles_book_complaint_proto_rawDescGZIP(), []int{0}
}

func (x *BookComplaint) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BookComplaint) GetBookID() int32 {
	if x != nil {
		return x.BookID
	}
	return 0
}

func (x *BookComplaint) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookComplaint) GetComplaint() string {
	if x != nil {
		return x.Complaint
	}
	return ""
}

// POST
type AddBookComplaintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookComplaint *BookComplaint `protobuf:"bytes,1,opt,name=BookComplaint,proto3" json:"BookComplaint,omitempty"`
}

func (x *AddBookComplaintRequest) Reset() {
	*x = AddBookComplaintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_book_complaint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookComplaintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookComplaintRequest) ProtoMessage() {}

func (x *AddBookComplaintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_book_complaint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookComplaintRequest.ProtoReflect.Descriptor instead.
func (*AddBookComplaintRequest) Descriptor() ([]byte, []int) {
	return file_protoFiles_book_complaint_proto_rawDescGZIP(), []int{1}
}

func (x *AddBookComplaintRequest) GetBookComplaint() *BookComplaint {
	if x != nil {
		return x.BookComplaint
	}
	return nil
}

type AddBookComplaintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AddBookComplaintResponse) Reset() {
	*x = AddBookComplaintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_book_complaint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookComplaintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookComplaintResponse) ProtoMessage() {}

func (x *AddBookComplaintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_book_complaint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookComplaintResponse.ProtoReflect.Descriptor instead.
func (*AddBookComplaintResponse) Descriptor() ([]byte, []int) {
	return file_protoFiles_book_complaint_proto_rawDescGZIP(), []int{2}
}

func (x *AddBookComplaintResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// PUT
type EditBookComplaintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookComplaint *BookComplaint `protobuf:"bytes,1,opt,name=BookComplaint,proto3" json:"BookComplaint,omitempty"`
}

func (x *EditBookComplaintRequest) Reset() {
	*x = EditBookComplaintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_book_complaint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditBookComplaintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditBookComplaintRequest) ProtoMessage() {}

func (x *EditBookComplaintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_book_complaint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditBookComplaintRequest.ProtoReflect.Descriptor instead.
func (*EditBookComplaintRequest) Descriptor() ([]byte, []int) {
	return file_protoFiles_book_complaint_proto_rawDescGZIP(), []int{3}
}

func (x *EditBookComplaintRequest) GetBookComplaint() *BookComplaint {
	if x != nil {
		return x.BookComplaint
	}
	return nil
}

type EditBookComplaintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *EditBookComplaintResponse) Reset() {
	*x = EditBookComplaintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_book_complaint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditBookComplaintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditBookComplaintResponse) ProtoMessage() {}

func (x *EditBookComplaintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_book_complaint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditBookComplaintResponse.ProtoReflect.Descriptor instead.
func (*EditBookComplaintResponse) Descriptor() ([]byte, []int) {
	return file_protoFiles_book_complaint_proto_rawDescGZIP(), []int{4}
}

func (x *EditBookComplaintResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_protoFiles_book_complaint_proto protoreflect.FileDescriptor

var file_protoFiles_book_complaint_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x6b, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x5e,
	0x0a, 0x17, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x42, 0x6f, 0x6f,
	0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52,
	0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x36,
	0x0a, 0x18, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x0a, 0x18, 0x45, 0x64, 0x69, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x19, 0x45, 0x64, 0x69, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xe7, 0x01, 0x0a, 0x14, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x27, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x68, 0x0a, 0x11, 0x45, 0x64, 0x69, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoFiles_book_complaint_proto_rawDescOnce sync.Once
	file_protoFiles_book_complaint_proto_rawDescData = file_protoFiles_book_complaint_proto_rawDesc
)

func file_protoFiles_book_complaint_proto_rawDescGZIP() []byte {
	file_protoFiles_book_complaint_proto_rawDescOnce.Do(func() {
		file_protoFiles_book_complaint_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoFiles_book_complaint_proto_rawDescData)
	})
	return file_protoFiles_book_complaint_proto_rawDescData
}

var file_protoFiles_book_complaint_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protoFiles_book_complaint_proto_goTypes = []interface{}{
	(*BookComplaint)(nil),             // 0: libraryService.BookComplaint
	(*AddBookComplaintRequest)(nil),   // 1: libraryService.AddBookComplaintRequest
	(*AddBookComplaintResponse)(nil),  // 2: libraryService.AddBookComplaintResponse
	(*EditBookComplaintRequest)(nil),  // 3: libraryService.EditBookComplaintRequest
	(*EditBookComplaintResponse)(nil), // 4: libraryService.EditBookComplaintResponse
}
var file_protoFiles_book_complaint_proto_depIdxs = []int32{
	0, // 0: libraryService.AddBookComplaintRequest.BookComplaint:type_name -> libraryService.BookComplaint
	0, // 1: libraryService.EditBookComplaintRequest.BookComplaint:type_name -> libraryService.BookComplaint
	1, // 2: libraryService.BookComplaintService.AddBookComplaint:input_type -> libraryService.AddBookComplaintRequest
	3, // 3: libraryService.BookComplaintService.EditBookComplaint:input_type -> libraryService.EditBookComplaintRequest
	2, // 4: libraryService.BookComplaintService.AddBookComplaint:output_type -> libraryService.AddBookComplaintResponse
	4, // 5: libraryService.BookComplaintService.EditBookComplaint:output_type -> libraryService.EditBookComplaintResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protoFiles_book_complaint_proto_init() }
func file_protoFiles_book_complaint_proto_init() {
	if File_protoFiles_book_complaint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoFiles_book_complaint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookComplaint); i {
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
		file_protoFiles_book_complaint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookComplaintRequest); i {
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
		file_protoFiles_book_complaint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookComplaintResponse); i {
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
		file_protoFiles_book_complaint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditBookComplaintRequest); i {
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
		file_protoFiles_book_complaint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditBookComplaintResponse); i {
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
			RawDescriptor: file_protoFiles_book_complaint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoFiles_book_complaint_proto_goTypes,
		DependencyIndexes: file_protoFiles_book_complaint_proto_depIdxs,
		MessageInfos:      file_protoFiles_book_complaint_proto_msgTypes,
	}.Build()
	File_protoFiles_book_complaint_proto = out.File
	file_protoFiles_book_complaint_proto_rawDesc = nil
	file_protoFiles_book_complaint_proto_goTypes = nil
	file_protoFiles_book_complaint_proto_depIdxs = nil
}
