// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: employee/employee.proto

package employee

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EmployeeId  string `protobuf:"bytes,2,opt,name=employeeId,proto3" json:"employeeId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Level       string `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Employee) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Employee) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *CreateEmployeeRequest) Reset() {
	*x = CreateEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeRequest) ProtoMessage() {}

func (x *CreateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*CreateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmployeeRequest) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type CreateEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"` // will have a employee id
}

func (x *CreateEmployeeResponse) Reset() {
	*x = CreateEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeResponse) ProtoMessage() {}

func (x *CreateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*CreateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEmployeeResponse) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type ReadEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string `protobuf:"bytes,1,opt,name=employeeId,proto3" json:"employeeId,omitempty"`
}

func (x *ReadEmployeeRequest) Reset() {
	*x = ReadEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEmployeeRequest) ProtoMessage() {}

func (x *ReadEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEmployeeRequest.ProtoReflect.Descriptor instead.
func (*ReadEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{3}
}

func (x *ReadEmployeeRequest) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

type ReadEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *ReadEmployeeResponse) Reset() {
	*x = ReadEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEmployeeResponse) ProtoMessage() {}

func (x *ReadEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEmployeeResponse.ProtoReflect.Descriptor instead.
func (*ReadEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{4}
}

func (x *ReadEmployeeResponse) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type UpdateEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *UpdateEmployeeRequest) Reset() {
	*x = UpdateEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeRequest) ProtoMessage() {}

func (x *UpdateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEmployeeRequest) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type UpdateEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *UpdateEmployeeResponse) Reset() {
	*x = UpdateEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeResponse) ProtoMessage() {}

func (x *UpdateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEmployeeResponse) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type DeleteEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string `protobuf:"bytes,1,opt,name=employeeId,proto3" json:"employeeId,omitempty"`
}

func (x *DeleteEmployeeRequest) Reset() {
	*x = DeleteEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmployeeRequest) ProtoMessage() {}

func (x *DeleteEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmployeeRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEmployeeRequest) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

type DeleteEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string `protobuf:"bytes,1,opt,name=employeeId,proto3" json:"employeeId,omitempty"`
}

func (x *DeleteEmployeeResponse) Reset() {
	*x = DeleteEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmployeeResponse) ProtoMessage() {}

func (x *DeleteEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmployeeResponse.ProtoReflect.Descriptor instead.
func (*DeleteEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEmployeeResponse) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

type ListEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *ListEmployeeRequest) Reset() {
	*x = ListEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmployeeRequest) ProtoMessage() {}

func (x *ListEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmployeeRequest.ProtoReflect.Descriptor instead.
func (*ListEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{9}
}

func (x *ListEmployeeRequest) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type ListEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *ListEmployeeResponse) Reset() {
	*x = ListEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmployeeResponse) ProtoMessage() {}

func (x *ListEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmployeeResponse.ProtoReflect.Descriptor instead.
func (*ListEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{10}
}

func (x *ListEmployeeResponse) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

var File_employee_employee_proto protoreflect.FileDescriptor

var file_employee_employee_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22,
	0x35, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x47,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x22, 0x37, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x46, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x32, 0xb0, 0x03, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1d, 0x2e, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1f, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x62,
	0x69, 0x72, 0x61, 0x32, 0x30, 0x30, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63,
	0x72, 0x75, 0x64, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_employee_employee_proto_rawDescOnce sync.Once
	file_employee_employee_proto_rawDescData = file_employee_employee_proto_rawDesc
)

func file_employee_employee_proto_rawDescGZIP() []byte {
	file_employee_employee_proto_rawDescOnce.Do(func() {
		file_employee_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_employee_proto_rawDescData)
	})
	return file_employee_employee_proto_rawDescData
}

var file_employee_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_employee_employee_proto_goTypes = []interface{}{
	(*Employee)(nil),               // 0: employee.Employee
	(*CreateEmployeeRequest)(nil),  // 1: employee.CreateEmployeeRequest
	(*CreateEmployeeResponse)(nil), // 2: employee.CreateEmployeeResponse
	(*ReadEmployeeRequest)(nil),    // 3: employee.ReadEmployeeRequest
	(*ReadEmployeeResponse)(nil),   // 4: employee.ReadEmployeeResponse
	(*UpdateEmployeeRequest)(nil),  // 5: employee.UpdateEmployeeRequest
	(*UpdateEmployeeResponse)(nil), // 6: employee.UpdateEmployeeResponse
	(*DeleteEmployeeRequest)(nil),  // 7: employee.DeleteEmployeeRequest
	(*DeleteEmployeeResponse)(nil), // 8: employee.DeleteEmployeeResponse
	(*ListEmployeeRequest)(nil),    // 9: employee.ListEmployeeRequest
	(*ListEmployeeResponse)(nil),   // 10: employee.ListEmployeeResponse
}
var file_employee_employee_proto_depIdxs = []int32{
	0,  // 0: employee.CreateEmployeeRequest.employee:type_name -> employee.Employee
	0,  // 1: employee.CreateEmployeeResponse.employee:type_name -> employee.Employee
	0,  // 2: employee.ReadEmployeeResponse.employee:type_name -> employee.Employee
	0,  // 3: employee.UpdateEmployeeRequest.employee:type_name -> employee.Employee
	0,  // 4: employee.UpdateEmployeeResponse.employee:type_name -> employee.Employee
	0,  // 5: employee.ListEmployeeRequest.employee:type_name -> employee.Employee
	0,  // 6: employee.ListEmployeeResponse.employee:type_name -> employee.Employee
	1,  // 7: employee.EmployeeService.CreateEmployee:input_type -> employee.CreateEmployeeRequest
	3,  // 8: employee.EmployeeService.ReadEmployee:input_type -> employee.ReadEmployeeRequest
	5,  // 9: employee.EmployeeService.UpdateEmployee:input_type -> employee.UpdateEmployeeRequest
	7,  // 10: employee.EmployeeService.DeleteEmployee:input_type -> employee.DeleteEmployeeRequest
	9,  // 11: employee.EmployeeService.ListEmployee:input_type -> employee.ListEmployeeRequest
	2,  // 12: employee.EmployeeService.CreateEmployee:output_type -> employee.CreateEmployeeResponse
	4,  // 13: employee.EmployeeService.ReadEmployee:output_type -> employee.ReadEmployeeResponse
	6,  // 14: employee.EmployeeService.UpdateEmployee:output_type -> employee.UpdateEmployeeResponse
	8,  // 15: employee.EmployeeService.DeleteEmployee:output_type -> employee.DeleteEmployeeResponse
	10, // 16: employee.EmployeeService.ListEmployee:output_type -> employee.ListEmployeeResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_employee_employee_proto_init() }
func file_employee_employee_proto_init() {
	if File_employee_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_employee_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmployeeRequest); i {
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
		file_employee_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmployeeResponse); i {
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
		file_employee_employee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEmployeeRequest); i {
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
		file_employee_employee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEmployeeResponse); i {
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
		file_employee_employee_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmployeeRequest); i {
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
		file_employee_employee_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmployeeResponse); i {
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
		file_employee_employee_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmployeeRequest); i {
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
		file_employee_employee_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmployeeResponse); i {
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
		file_employee_employee_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmployeeRequest); i {
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
		file_employee_employee_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmployeeResponse); i {
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
			RawDescriptor: file_employee_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employee_employee_proto_goTypes,
		DependencyIndexes: file_employee_employee_proto_depIdxs,
		MessageInfos:      file_employee_employee_proto_msgTypes,
	}.Build()
	File_employee_employee_proto = out.File
	file_employee_employee_proto_rawDesc = nil
	file_employee_employee_proto_goTypes = nil
	file_employee_employee_proto_depIdxs = nil
}