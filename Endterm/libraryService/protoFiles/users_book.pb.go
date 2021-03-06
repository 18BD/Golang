// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: protoFiles/users_book.proto

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

// Method
type UserBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID int32  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	BookID int32  `protobuf:"varint,3,opt,name=BookID,proto3" json:"BookID,omitempty"`
	Title  string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Status string `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Star   bool   `protobuf:"varint,6,opt,name=Star,proto3" json:"Star,omitempty"`
}

func (x *UserBook) Reset() {
	*x = UserBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBook) ProtoMessage() {}

func (x *UserBook) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBook.ProtoReflect.Descriptor instead.
func (*UserBook) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{0}
}

func (x *UserBook) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserBook) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserBook) GetBookID() int32 {
	if x != nil {
		return x.BookID
	}
	return 0
}

func (x *UserBook) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UserBook) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserBook) GetStar() bool {
	if x != nil {
		return x.Star
	}
	return false
}

// GET
type GetUserBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserBooksRequest) Reset() {
	*x = GetUserBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBooksRequest) ProtoMessage() {}

func (x *GetUserBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBooksRequest.ProtoReflect.Descriptor instead.
func (*GetUserBooksRequest) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{1}
}

type GetUserBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBooks []*UserBook `protobuf:"bytes,1,rep,name=UserBooks,proto3" json:"UserBooks,omitempty"`
}

func (x *GetUserBooksResponse) Reset() {
	*x = GetUserBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBooksResponse) ProtoMessage() {}

func (x *GetUserBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBooksResponse.ProtoReflect.Descriptor instead.
func (*GetUserBooksResponse) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserBooksResponse) GetUserBooks() []*UserBook {
	if x != nil {
		return x.UserBooks
	}
	return nil
}

// POST
type AddUserBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBook *UserBook `protobuf:"bytes,1,opt,name=UserBook,proto3" json:"UserBook,omitempty"`
}

func (x *AddUserBookRequest) Reset() {
	*x = AddUserBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserBookRequest) ProtoMessage() {}

func (x *AddUserBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserBookRequest.ProtoReflect.Descriptor instead.
func (*AddUserBookRequest) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{3}
}

func (x *AddUserBookRequest) GetUserBook() *UserBook {
	if x != nil {
		return x.UserBook
	}
	return nil
}

type AddUserBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AddUserBookResponse) Reset() {
	*x = AddUserBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserBookResponse) ProtoMessage() {}

func (x *AddUserBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserBookResponse.ProtoReflect.Descriptor instead.
func (*AddUserBookResponse) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{4}
}

func (x *AddUserBookResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// PUT
type EditUserBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBook *UserBook `protobuf:"bytes,1,opt,name=UserBook,proto3" json:"UserBook,omitempty"`
}

func (x *EditUserBookRequest) Reset() {
	*x = EditUserBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserBookRequest) ProtoMessage() {}

func (x *EditUserBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserBookRequest.ProtoReflect.Descriptor instead.
func (*EditUserBookRequest) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{5}
}

func (x *EditUserBookRequest) GetUserBook() *UserBook {
	if x != nil {
		return x.UserBook
	}
	return nil
}

type EditUserBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *EditUserBookResponse) Reset() {
	*x = EditUserBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserBookResponse) ProtoMessage() {}

func (x *EditUserBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserBookResponse.ProtoReflect.Descriptor instead.
func (*EditUserBookResponse) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{6}
}

func (x *EditUserBookResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// DELETE
type DeleteUserBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteUserBookRequest) Reset() {
	*x = DeleteUserBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserBookRequest) ProtoMessage() {}

func (x *DeleteUserBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserBookRequest) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserBookRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteUserBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteUserBookResponse) Reset() {
	*x = DeleteUserBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoFiles_users_book_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserBookResponse) ProtoMessage() {}

func (x *DeleteUserBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoFiles_users_book_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserBookResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserBookResponse) Descriptor() ([]byte, []int) {
	return file_protoFiles_users_book_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserBookResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_protoFiles_users_book_proto protoreflect.FileDescriptor

var file_protoFiles_users_book_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x8c, 0x01,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x53, 0x74, 0x61, 0x72, 0x22, 0x15, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x22, 0x4a, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x22,
	0x31, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x22,
	0x32, 0x0a, 0x14, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xff, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x23, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x22, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x23, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x25, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoFiles_users_book_proto_rawDescOnce sync.Once
	file_protoFiles_users_book_proto_rawDescData = file_protoFiles_users_book_proto_rawDesc
)

func file_protoFiles_users_book_proto_rawDescGZIP() []byte {
	file_protoFiles_users_book_proto_rawDescOnce.Do(func() {
		file_protoFiles_users_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoFiles_users_book_proto_rawDescData)
	})
	return file_protoFiles_users_book_proto_rawDescData
}

var file_protoFiles_users_book_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protoFiles_users_book_proto_goTypes = []interface{}{
	(*UserBook)(nil),               // 0: libraryService.UserBook
	(*GetUserBooksRequest)(nil),    // 1: libraryService.GetUserBooksRequest
	(*GetUserBooksResponse)(nil),   // 2: libraryService.GetUserBooksResponse
	(*AddUserBookRequest)(nil),     // 3: libraryService.AddUserBookRequest
	(*AddUserBookResponse)(nil),    // 4: libraryService.AddUserBookResponse
	(*EditUserBookRequest)(nil),    // 5: libraryService.EditUserBookRequest
	(*EditUserBookResponse)(nil),   // 6: libraryService.EditUserBookResponse
	(*DeleteUserBookRequest)(nil),  // 7: libraryService.DeleteUserBookRequest
	(*DeleteUserBookResponse)(nil), // 8: libraryService.DeleteUserBookResponse
}
var file_protoFiles_users_book_proto_depIdxs = []int32{
	0, // 0: libraryService.GetUserBooksResponse.UserBooks:type_name -> libraryService.UserBook
	0, // 1: libraryService.AddUserBookRequest.UserBook:type_name -> libraryService.UserBook
	0, // 2: libraryService.EditUserBookRequest.UserBook:type_name -> libraryService.UserBook
	1, // 3: libraryService.UserBookService.GetUserBook:input_type -> libraryService.GetUserBooksRequest
	3, // 4: libraryService.UserBookService.AddUserBook:input_type -> libraryService.AddUserBookRequest
	5, // 5: libraryService.UserBookService.EditUserBook:input_type -> libraryService.EditUserBookRequest
	7, // 6: libraryService.UserBookService.DeleteUserBook:input_type -> libraryService.DeleteUserBookRequest
	2, // 7: libraryService.UserBookService.GetUserBook:output_type -> libraryService.GetUserBooksResponse
	4, // 8: libraryService.UserBookService.AddUserBook:output_type -> libraryService.AddUserBookResponse
	6, // 9: libraryService.UserBookService.EditUserBook:output_type -> libraryService.EditUserBookResponse
	8, // 10: libraryService.UserBookService.DeleteUserBook:output_type -> libraryService.DeleteUserBookResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protoFiles_users_book_proto_init() }
func file_protoFiles_users_book_proto_init() {
	if File_protoFiles_users_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoFiles_users_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBook); i {
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
		file_protoFiles_users_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBooksRequest); i {
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
		file_protoFiles_users_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBooksResponse); i {
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
		file_protoFiles_users_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserBookRequest); i {
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
		file_protoFiles_users_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserBookResponse); i {
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
		file_protoFiles_users_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserBookRequest); i {
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
		file_protoFiles_users_book_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserBookResponse); i {
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
		file_protoFiles_users_book_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserBookRequest); i {
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
		file_protoFiles_users_book_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserBookResponse); i {
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
			RawDescriptor: file_protoFiles_users_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoFiles_users_book_proto_goTypes,
		DependencyIndexes: file_protoFiles_users_book_proto_depIdxs,
		MessageInfos:      file_protoFiles_users_book_proto_msgTypes,
	}.Build()
	File_protoFiles_users_book_proto = out.File
	file_protoFiles_users_book_proto_rawDesc = nil
	file_protoFiles_users_book_proto_goTypes = nil
	file_protoFiles_users_book_proto_depIdxs = nil
}
