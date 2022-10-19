// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: rpc_customer.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role           int32                  `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email          string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber    string                 `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DateOfBith     string                 `protobuf:"bytes,6,opt,name=date_of_bith,json=dateOfBith,proto3" json:"date_of_bith,omitempty"`
	IdentityCard   string                 `protobuf:"bytes,7,opt,name=identity_card,json=identityCard,proto3" json:"identity_card,omitempty"`
	Address        string                 `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	MembershipCard string                 `protobuf:"bytes,9,opt,name=membership_card,json=membershipCard,proto3" json:"membership_card,omitempty"`
	Password       string                 `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	Status         int32                  `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_rpc_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Customer) GetDateOfBith() string {
	if x != nil {
		return x.DateOfBith
	}
	return ""
}

func (x *Customer) GetIdentityCard() string {
	if x != nil {
		return x.IdentityCard
	}
	return ""
}

func (x *Customer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Customer) GetMembershipCard() string {
	if x != nil {
		return x.MembershipCard
	}
	return ""
}

func (x *Customer) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Customer) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Customer) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Customer) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type BookingHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	FlightId    string `protobuf:"bytes,2,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
	BookingCode string `protobuf:"bytes,3,opt,name=booking_code,json=bookingCode,proto3" json:"booking_code,omitempty"`
	FlightName  string `protobuf:"bytes,4,opt,name=flight_name,json=flightName,proto3" json:"flight_name,omitempty"`
}

func (x *BookingHistoryRequest) Reset() {
	*x = BookingHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingHistoryRequest) ProtoMessage() {}

func (x *BookingHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingHistoryRequest.ProtoReflect.Descriptor instead.
func (*BookingHistoryRequest) Descriptor() ([]byte, []int) {
	return file_rpc_customer_proto_rawDescGZIP(), []int{1}
}

func (x *BookingHistoryRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *BookingHistoryRequest) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *BookingHistoryRequest) GetBookingCode() string {
	if x != nil {
		return x.BookingCode
	}
	return ""
}

func (x *BookingHistoryRequest) GetFlightName() string {
	if x != nil {
		return x.FlightName
	}
	return ""
}

type BookingHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	History []*Booking `protobuf:"bytes,1,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *BookingHistoryResponse) Reset() {
	*x = BookingHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingHistoryResponse) ProtoMessage() {}

func (x *BookingHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingHistoryResponse.ProtoReflect.Descriptor instead.
func (*BookingHistoryResponse) Descriptor() ([]byte, []int) {
	return file_rpc_customer_proto_rawDescGZIP(), []int{2}
}

func (x *BookingHistoryResponse) GetHistory() []*Booking {
	if x != nil {
		return x.History
	}
	return nil
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	OldPassword     string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_customer_proto_rawDescGZIP(), []int{3}
}

func (x *ChangePasswordRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *ChangePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

type ChangePasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChangePasswordResponse) Reset() {
	*x = ChangePasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordResponse) ProtoMessage() {}

func (x *ChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_customer_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePasswordResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChangePasswordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_customer_proto protoreflect.FileDescriptor

var file_rpc_customer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x1a, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x08, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x66, 0x5f, 0x62, 0x69, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x15, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x16, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x74, 0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x22, 0xa9, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x46, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xdb, 0x02, 0x0a, 0x0b, 0x52, 0x50, 0x43, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x74, 0x75, 0x6e, 0x73,
	0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x74, 0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x44, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x18, 0x2e, 0x74, 0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x74, 0x75, 0x6e, 0x73,
	0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x74, 0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74,
	0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x74, 0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f,
	0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x74, 0x75, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_customer_proto_rawDescOnce sync.Once
	file_rpc_customer_proto_rawDescData = file_rpc_customer_proto_rawDesc
)

func file_rpc_customer_proto_rawDescGZIP() []byte {
	file_rpc_customer_proto_rawDescOnce.Do(func() {
		file_rpc_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_customer_proto_rawDescData)
	})
	return file_rpc_customer_proto_rawDescData
}

var file_rpc_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_customer_proto_goTypes = []interface{}{
	(*Customer)(nil),               // 0: tuns_go_flight.Customer
	(*BookingHistoryRequest)(nil),  // 1: tuns_go_flight.BookingHistoryRequest
	(*BookingHistoryResponse)(nil), // 2: tuns_go_flight.BookingHistoryResponse
	(*ChangePasswordRequest)(nil),  // 3: tuns_go_flight.ChangePasswordRequest
	(*ChangePasswordResponse)(nil), // 4: tuns_go_flight.ChangePasswordResponse
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*Booking)(nil),                // 6: tuns_go_flight.Booking
}
var file_rpc_customer_proto_depIdxs = []int32{
	5, // 0: tuns_go_flight.Customer.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: tuns_go_flight.Customer.updated_at:type_name -> google.protobuf.Timestamp
	6, // 2: tuns_go_flight.BookingHistoryResponse.history:type_name -> tuns_go_flight.Booking
	0, // 3: tuns_go_flight.RPCCustomer.CreateCustomer:input_type -> tuns_go_flight.Customer
	0, // 4: tuns_go_flight.RPCCustomer.UpdateCustomer:input_type -> tuns_go_flight.Customer
	3, // 5: tuns_go_flight.RPCCustomer.ChangePassword:input_type -> tuns_go_flight.ChangePasswordRequest
	1, // 6: tuns_go_flight.RPCCustomer.BookingHistory:input_type -> tuns_go_flight.BookingHistoryRequest
	0, // 7: tuns_go_flight.RPCCustomer.CreateCustomer:output_type -> tuns_go_flight.Customer
	0, // 8: tuns_go_flight.RPCCustomer.UpdateCustomer:output_type -> tuns_go_flight.Customer
	4, // 9: tuns_go_flight.RPCCustomer.ChangePassword:output_type -> tuns_go_flight.ChangePasswordResponse
	2, // 10: tuns_go_flight.RPCCustomer.BookingHistory:output_type -> tuns_go_flight.BookingHistoryResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_customer_proto_init() }
func file_rpc_customer_proto_init() {
	if File_rpc_customer_proto != nil {
		return
	}
	file_rpc_booking_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_rpc_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingHistoryRequest); i {
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
		file_rpc_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingHistoryResponse); i {
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
		file_rpc_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_rpc_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordResponse); i {
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
			RawDescriptor: file_rpc_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_customer_proto_goTypes,
		DependencyIndexes: file_rpc_customer_proto_depIdxs,
		MessageInfos:      file_rpc_customer_proto_msgTypes,
	}.Build()
	File_rpc_customer_proto = out.File
	file_rpc_customer_proto_rawDesc = nil
	file_rpc_customer_proto_goTypes = nil
	file_rpc_customer_proto_depIdxs = nil
}
