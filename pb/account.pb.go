// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: account.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Account struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner         string                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Balance       float32                `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Currency      string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Account) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Account) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Account) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currency      string                 `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       *Account               `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	mi := &file_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

type GetAccountDetailsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       *Account               `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountDetailsResponse) Reset() {
	*x = GetAccountDetailsResponse{}
	mi := &file_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailsResponse) ProtoMessage() {}

func (x *GetAccountDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetAccountDetailsResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccountDetailsResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type ListAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageId        int32                  `protobuf:"varint,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAccountRequest) Reset() {
	*x = ListAccountRequest{}
	mi := &file_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountRequest) ProtoMessage() {}

func (x *ListAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountRequest.ProtoReflect.Descriptor instead.
func (*ListAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

func (x *ListAccountRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *ListAccountRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Accounts      []*Account             `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAccountResponse) Reset() {
	*x = ListAccountResponse{}
	mi := &file_account_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountResponse) ProtoMessage() {}

func (x *ListAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountResponse.ProtoReflect.Descriptor instead.
func (*ListAccountResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{6}
}

func (x *ListAccountResponse) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_account_proto protoreflect.FileDescriptor

const file_account_proto_rawDesc = "" +
	"\n" +
	"\raccount.proto\x12\x02pb\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa0\x01\n" +
	"\aAccount\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x14\n" +
	"\x05owner\x18\x02 \x01(\tR\x05owner\x12\x18\n" +
	"\abalance\x18\x03 \x01(\x02R\abalance\x12\x1a\n" +
	"\bcurrency\x18\x04 \x01(\tR\bcurrency\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"2\n" +
	"\x14CreateAccountRequest\x12\x1a\n" +
	"\bcurrency\x18\x01 \x01(\tR\bcurrency\">\n" +
	"\x15CreateAccountResponse\x12%\n" +
	"\aaccount\x18\x01 \x01(\v2\v.pb.AccountR\aaccount\"\a\n" +
	"\x05Empty\"B\n" +
	"\x19GetAccountDetailsResponse\x12%\n" +
	"\aaccount\x18\x01 \x01(\v2\v.pb.AccountR\aaccount\"H\n" +
	"\x12ListAccountRequest\x12\x16\n" +
	"\x06pageId\x18\x01 \x01(\x05R\x06pageId\x12\x1a\n" +
	"\bpageSize\x18\x02 \x01(\x05R\bpageSize\">\n" +
	"\x13ListAccountResponse\x12'\n" +
	"\baccounts\x18\x01 \x03(\v2\v.pb.AccountR\baccountsB-Z+github.com/suhailmuhammed157/simple_bank/pbb\x06proto3"

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData []byte
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_proto_rawDesc), len(file_account_proto_rawDesc)))
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_account_proto_goTypes = []any{
	(*Account)(nil),                   // 0: pb.Account
	(*CreateAccountRequest)(nil),      // 1: pb.CreateAccountRequest
	(*CreateAccountResponse)(nil),     // 2: pb.CreateAccountResponse
	(*Empty)(nil),                     // 3: pb.Empty
	(*GetAccountDetailsResponse)(nil), // 4: pb.GetAccountDetailsResponse
	(*ListAccountRequest)(nil),        // 5: pb.ListAccountRequest
	(*ListAccountResponse)(nil),       // 6: pb.ListAccountResponse
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_account_proto_depIdxs = []int32{
	7, // 0: pb.Account.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: pb.CreateAccountResponse.account:type_name -> pb.Account
	0, // 2: pb.GetAccountDetailsResponse.account:type_name -> pb.Account
	0, // 3: pb.ListAccountResponse.accounts:type_name -> pb.Account
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_proto_rawDesc), len(file_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
