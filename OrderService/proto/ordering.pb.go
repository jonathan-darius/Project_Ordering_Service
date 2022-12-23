// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ordering.proto

package OrderingService

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ProductID string `protobuf:"bytes,2,opt,name=productID,proto3" json:"productID,omitempty"`
	Qty       int32  `protobuf:"varint,3,opt,name=qty,proto3" json:"qty,omitempty"`
	Total     int64  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CartItem) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *CartItem) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *CartItem) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{1}
}

func (x *UserID) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string      `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Item   []*CartItem `protobuf:"bytes,2,rep,name=Item,proto3" json:"Item,omitempty"`
	Total  int64       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{2}
}

func (x *Cart) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Cart) GetItem() []*CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *Cart) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page          int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	RecordPerPage int32 `protobuf:"varint,2,opt,name=recordPerPage,proto3" json:"recordPerPage,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetRecordPerPage() int32 {
	if x != nil {
		return x.RecordPerPage
	}
	return 0
}

type TransactionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransactionID) Reset() {
	*x = TransactionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionID) ProtoMessage() {}

func (x *TransactionID) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionID.ProtoReflect.Descriptor instead.
func (*TransactionID) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid        string  `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	TransactionId string  `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	ProductId     string  `protobuf:"bytes,3,opt,name=productId,proto3" json:"productId,omitempty"`
	Rating        float64 `protobuf:"fixed64,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Desc          string  `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{5}
}

func (x *Rating) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *Rating) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Rating) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Rating) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Rating) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type DetailTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string  `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Qty       int32   `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
	Rating    float64 `protobuf:"fixed64,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Desc      string  `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Total     int64   `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *DetailTransaction) Reset() {
	*x = DetailTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailTransaction) ProtoMessage() {}

func (x *DetailTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailTransaction.ProtoReflect.Descriptor instead.
func (*DetailTransaction) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{6}
}

func (x *DetailTransaction) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *DetailTransaction) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *DetailTransaction) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *DetailTransaction) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *DetailTransaction) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string               `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	UserId        string               `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Total         int64                `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	CreateAt      int64                `protobuf:"varint,4,opt,name=createAt,proto3" json:"createAt,omitempty"`
	Detail        []*DetailTransaction `protobuf:"bytes,5,rep,name=detail,proto3" json:"detail,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{7}
}

func (x *Transaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Transaction) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Transaction) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Transaction) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Transaction) GetDetail() []*DetailTransaction {
	if x != nil {
		return x.Detail
	}
	return nil
}

type UserPurchase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     string      `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *UserPurchase) Reset() {
	*x = UserPurchase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPurchase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPurchase) ProtoMessage() {}

func (x *UserPurchase) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPurchase.ProtoReflect.Descriptor instead.
func (*UserPurchase) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{8}
}

func (x *UserPurchase) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserPurchase) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_ordering_proto protoreflect.FileDescriptor

var file_ordering_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x22, 0x5c, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x46, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x06, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x85, 0x01,
	0x0a, 0x11, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x71, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x5c, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xea, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0e,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x12, 0x35,
	0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x85, 0x02, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x42, 0x0a,
	0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x30,
	0x01, 0x12, 0x43, 0x0a, 0x11, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a,
	0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2f, 0x5a,
	0x2d, 0x4d, 0x69, 0x6e, 0x69, 0x5f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ordering_proto_rawDescOnce sync.Once
	file_ordering_proto_rawDescData = file_ordering_proto_rawDesc
)

func file_ordering_proto_rawDescGZIP() []byte {
	file_ordering_proto_rawDescOnce.Do(func() {
		file_ordering_proto_rawDescData = protoimpl.X.CompressGZIP(file_ordering_proto_rawDescData)
	})
	return file_ordering_proto_rawDescData
}

var file_ordering_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ordering_proto_goTypes = []interface{}{
	(*CartItem)(nil),          // 0: ordering.CartItem
	(*UserID)(nil),            // 1: ordering.UserID
	(*Cart)(nil),              // 2: ordering.Cart
	(*Pagination)(nil),        // 3: ordering.Pagination
	(*TransactionID)(nil),     // 4: ordering.TransactionID
	(*Rating)(nil),            // 5: ordering.Rating
	(*DetailTransaction)(nil), // 6: ordering.DetailTransaction
	(*Transaction)(nil),       // 7: ordering.Transaction
	(*UserPurchase)(nil),      // 8: ordering.UserPurchase
	(*empty.Empty)(nil),       // 9: google.protobuf.Empty
}
var file_ordering_proto_depIdxs = []int32{
	0,  // 0: ordering.Cart.Item:type_name -> ordering.CartItem
	6,  // 1: ordering.Transaction.detail:type_name -> ordering.DetailTransaction
	3,  // 2: ordering.UserPurchase.pagination:type_name -> ordering.Pagination
	0,  // 3: ordering.CartService.AddCartItem:input_type -> ordering.CartItem
	0,  // 4: ordering.CartService.RemoveCartItem:input_type -> ordering.CartItem
	1,  // 5: ordering.CartService.GetCart:input_type -> ordering.UserID
	1,  // 6: ordering.CartService.EmptyCart:input_type -> ordering.UserID
	1,  // 7: ordering.Purchase.PurchaseCart:input_type -> ordering.UserID
	8,  // 8: ordering.Purchase.PurchaseHistory:input_type -> ordering.UserPurchase
	4,  // 9: ordering.Purchase.DetailTransaction:input_type -> ordering.TransactionID
	5,  // 10: ordering.Purchase.AddRating:input_type -> ordering.Rating
	9,  // 11: ordering.CartService.AddCartItem:output_type -> google.protobuf.Empty
	9,  // 12: ordering.CartService.RemoveCartItem:output_type -> google.protobuf.Empty
	2,  // 13: ordering.CartService.GetCart:output_type -> ordering.Cart
	9,  // 14: ordering.CartService.EmptyCart:output_type -> google.protobuf.Empty
	4,  // 15: ordering.Purchase.PurchaseCart:output_type -> ordering.TransactionID
	7,  // 16: ordering.Purchase.PurchaseHistory:output_type -> ordering.Transaction
	7,  // 17: ordering.Purchase.DetailTransaction:output_type -> ordering.Transaction
	9,  // 18: ordering.Purchase.AddRating:output_type -> google.protobuf.Empty
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_ordering_proto_init() }
func file_ordering_proto_init() {
	if File_ordering_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ordering_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_ordering_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_ordering_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_ordering_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_ordering_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionID); i {
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
		file_ordering_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_ordering_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailTransaction); i {
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
		file_ordering_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_ordering_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPurchase); i {
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
			RawDescriptor: file_ordering_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_ordering_proto_goTypes,
		DependencyIndexes: file_ordering_proto_depIdxs,
		MessageInfos:      file_ordering_proto_msgTypes,
	}.Build()
	File_ordering_proto = out.File
	file_ordering_proto_rawDesc = nil
	file_ordering_proto_goTypes = nil
	file_ordering_proto_depIdxs = nil
}