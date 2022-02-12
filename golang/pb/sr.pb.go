// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pb/sr.proto

package pb

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

type PackageState int32

const (
	PackageState_UPLOADING         PackageState = 0          // package is currently in the process of being uploaded
	PackageState_UPLOADED          PackageState = 1          // upload has completed; the package has not yet started validataion
	PackageState_VALIDATING        PackageState = 2          // the package is in the process of being validated
	PackageState_INVALID           PackageState = 3          // the package failed validation checks and cannot be deployed
	PackageState_DEPLOYING         PackageState = 5          // the package is in the process of being deployed to production
	PackageState_PRODUCTION        PackageState = 6          // the package is active in production
	PackageState_DEACTIVATING      PackageState = 7          // the package is in the process of being deactivated from production
	PackageState_DELETING          PackageState = 9          // the package is in the process of being deleted
	PackageState_UNKNOWN_PKG_STATE PackageState = 2147483647 // MAX_INT
)

// Enum value maps for PackageState.
var (
	PackageState_name = map[int32]string{
		0:          "UPLOADING",
		1:          "UPLOADED",
		2:          "VALIDATING",
		3:          "INVALID",
		5:          "DEPLOYING",
		6:          "PRODUCTION",
		7:          "DEACTIVATING",
		9:          "DELETING",
		2147483647: "UNKNOWN_PKG_STATE",
	}
	PackageState_value = map[string]int32{
		"UPLOADING":         0,
		"UPLOADED":          1,
		"VALIDATING":        2,
		"INVALID":           3,
		"DEPLOYING":         5,
		"PRODUCTION":        6,
		"DEACTIVATING":      7,
		"DELETING":          9,
		"UNKNOWN_PKG_STATE": 2147483647,
	}
)

func (x PackageState) Enum() *PackageState {
	p := new(PackageState)
	*p = x
	return p
}

func (x PackageState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageState) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_sr_proto_enumTypes[0].Descriptor()
}

func (PackageState) Type() protoreflect.EnumType {
	return &file_pb_sr_proto_enumTypes[0]
}

func (x PackageState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageState.Descriptor instead.
func (PackageState) EnumDescriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{0}
}

type PackageDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName        string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`               // name of the Bopmatic project; must be unique
	PackageId          string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`                   // hex string of first 4 bytes of packageXsum
	PackageXsum        []byte `protobuf:"bytes,3,opt,name=packageXsum,proto3" json:"packageXsum,omitempty"`               // sha256 checksum of packageTarballData
	PackageTarballData []byte `protobuf:"bytes,4,opt,name=packageTarballData,proto3" json:"packageTarballData,omitempty"` // package tarball content in .tar.xz format
	PackageName        string `protobuf:"bytes,5,opt,name=packageName,proto3" json:"packageName,omitempty"`               // user prescribed name; may not be unique; may be empty string
}

func (x *PackageDescription) Reset() {
	*x = PackageDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageDescription) ProtoMessage() {}

func (x *PackageDescription) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageDescription.ProtoReflect.Descriptor instead.
func (*PackageDescription) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{0}
}

func (x *PackageDescription) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *PackageDescription) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *PackageDescription) GetPackageXsum() []byte {
	if x != nil {
		return x.PackageXsum
	}
	return nil
}

func (x *PackageDescription) GetPackageTarballData() []byte {
	if x != nil {
		return x.PackageTarballData
	}
	return nil
}

func (x *PackageDescription) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

type DeployPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc *PackageDescription `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *DeployPackageRequest) Reset() {
	*x = DeployPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployPackageRequest) ProtoMessage() {}

func (x *DeployPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployPackageRequest.ProtoReflect.Descriptor instead.
func (*DeployPackageRequest) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{1}
}

func (x *DeployPackageRequest) GetDesc() *PackageDescription {
	if x != nil {
		return x.Desc
	}
	return nil
}

type DeployPackageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State PackageState `protobuf:"varint,1,opt,name=state,proto3,enum=PackageState" json:"state,omitempty"`
}

func (x *DeployPackageReply) Reset() {
	*x = DeployPackageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployPackageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployPackageReply) ProtoMessage() {}

func (x *DeployPackageReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployPackageReply.ProtoReflect.Descriptor instead.
func (*DeployPackageReply) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{2}
}

func (x *DeployPackageReply) GetState() PackageState {
	if x != nil {
		return x.State
	}
	return PackageState_UPLOADING
}

type DescribePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	PackageId   string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
}

func (x *DescribePackageRequest) Reset() {
	*x = DescribePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePackageRequest) ProtoMessage() {}

func (x *DescribePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePackageRequest.ProtoReflect.Descriptor instead.
func (*DescribePackageRequest) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePackageRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *DescribePackageRequest) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

type DescribePackageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc         *PackageDescription `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	PackageState PackageState        `protobuf:"varint,2,opt,name=packageState,proto3,enum=PackageState" json:"packageState,omitempty"`
	// the following fields are only valid when packageState==PRODUCTION
	SiteEndpoint string   `protobuf:"bytes,3,opt,name=siteEndpoint,proto3" json:"siteEndpoint,omitempty"`
	RpcEndpoints []string `protobuf:"bytes,4,rep,name=rpcEndpoints,proto3" json:"rpcEndpoints,omitempty"`
}

func (x *DescribePackageReply) Reset() {
	*x = DescribePackageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePackageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePackageReply) ProtoMessage() {}

func (x *DescribePackageReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePackageReply.ProtoReflect.Descriptor instead.
func (*DescribePackageReply) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{4}
}

func (x *DescribePackageReply) GetDesc() *PackageDescription {
	if x != nil {
		return x.Desc
	}
	return nil
}

func (x *DescribePackageReply) GetPackageState() PackageState {
	if x != nil {
		return x.PackageState
	}
	return PackageState_UPLOADING
}

func (x *DescribePackageReply) GetSiteEndpoint() string {
	if x != nil {
		return x.SiteEndpoint
	}
	return ""
}

func (x *DescribePackageReply) GetRpcEndpoints() []string {
	if x != nil {
		return x.RpcEndpoints
	}
	return nil
}

type ListPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"` // leave empty for all projects
}

func (x *ListPackagesRequest) Reset() {
	*x = ListPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPackagesRequest) ProtoMessage() {}

func (x *ListPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPackagesRequest.ProtoReflect.Descriptor instead.
func (*ListPackagesRequest) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{5}
}

func (x *ListPackagesRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type ListPackagesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ListPackagesReply_ListPackagesItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPackagesReply) Reset() {
	*x = ListPackagesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPackagesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPackagesReply) ProtoMessage() {}

func (x *ListPackagesReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPackagesReply.ProtoReflect.Descriptor instead.
func (*ListPackagesReply) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{6}
}

func (x *ListPackagesReply) GetItems() []*ListPackagesReply_ListPackagesItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeletePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	PackageId   string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
}

func (x *DeletePackageRequest) Reset() {
	*x = DeletePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageRequest) ProtoMessage() {}

func (x *DeletePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageRequest.ProtoReflect.Descriptor instead.
func (*DeletePackageRequest) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePackageRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *DeletePackageRequest) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

type DeletePackageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State PackageState `protobuf:"varint,1,opt,name=state,proto3,enum=PackageState" json:"state,omitempty"`
}

func (x *DeletePackageReply) Reset() {
	*x = DeletePackageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageReply) ProtoMessage() {}

func (x *DeletePackageReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageReply.ProtoReflect.Descriptor instead.
func (*DeletePackageReply) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePackageReply) GetState() PackageState {
	if x != nil {
		return x.State
	}
	return PackageState_UPLOADING
}

type ListPackagesReply_ListPackagesItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	PackageId   string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
}

func (x *ListPackagesReply_ListPackagesItem) Reset() {
	*x = ListPackagesReply_ListPackagesItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sr_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPackagesReply_ListPackagesItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPackagesReply_ListPackagesItem) ProtoMessage() {}

func (x *ListPackagesReply_ListPackagesItem) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sr_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPackagesReply_ListPackagesItem.ProtoReflect.Descriptor instead.
func (*ListPackagesReply_ListPackagesItem) Descriptor() ([]byte, []int) {
	return file_pb_sr_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListPackagesReply_ListPackagesItem) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ListPackagesReply_ListPackagesItem) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

var File_pb_sr_proto protoreflect.FileDescriptor

var file_pb_sr_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x62, 0x2f, 0x73, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01,
	0x0a, 0x12, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x58,
	0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x58, 0x73, 0x75, 0x6d, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x54, 0x61, 0x72, 0x62, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x12, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x61, 0x72, 0x62, 0x61,
	0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x39, 0x0a, 0x12, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xba,
	0x01, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x31, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x74, 0x65, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x70, 0x63, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x70, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x37, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x52, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x22, 0x39, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0xa2, 0x01, 0x0a,
	0x0c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x45, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x50, 0x4b, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0xff, 0xff, 0xff, 0xff,
	0x07, 0x32, 0x8e, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x62, 0x6f, 0x70, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_sr_proto_rawDescOnce sync.Once
	file_pb_sr_proto_rawDescData = file_pb_sr_proto_rawDesc
)

func file_pb_sr_proto_rawDescGZIP() []byte {
	file_pb_sr_proto_rawDescOnce.Do(func() {
		file_pb_sr_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_sr_proto_rawDescData)
	})
	return file_pb_sr_proto_rawDescData
}

var file_pb_sr_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_sr_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pb_sr_proto_goTypes = []interface{}{
	(PackageState)(0),                          // 0: PackageState
	(*PackageDescription)(nil),                 // 1: PackageDescription
	(*DeployPackageRequest)(nil),               // 2: DeployPackageRequest
	(*DeployPackageReply)(nil),                 // 3: DeployPackageReply
	(*DescribePackageRequest)(nil),             // 4: DescribePackageRequest
	(*DescribePackageReply)(nil),               // 5: DescribePackageReply
	(*ListPackagesRequest)(nil),                // 6: ListPackagesRequest
	(*ListPackagesReply)(nil),                  // 7: ListPackagesReply
	(*DeletePackageRequest)(nil),               // 8: DeletePackageRequest
	(*DeletePackageReply)(nil),                 // 9: DeletePackageReply
	(*ListPackagesReply_ListPackagesItem)(nil), // 10: ListPackagesReply.ListPackagesItem
}
var file_pb_sr_proto_depIdxs = []int32{
	1,  // 0: DeployPackageRequest.desc:type_name -> PackageDescription
	0,  // 1: DeployPackageReply.state:type_name -> PackageState
	1,  // 2: DescribePackageReply.desc:type_name -> PackageDescription
	0,  // 3: DescribePackageReply.packageState:type_name -> PackageState
	10, // 4: ListPackagesReply.items:type_name -> ListPackagesReply.ListPackagesItem
	0,  // 5: DeletePackageReply.state:type_name -> PackageState
	4,  // 6: ServiceRunner.DescribePackage:input_type -> DescribePackageRequest
	8,  // 7: ServiceRunner.DeletePackage:input_type -> DeletePackageRequest
	6,  // 8: ServiceRunner.ListPackages:input_type -> ListPackagesRequest
	2,  // 9: ServiceRunner.DeployPackage:input_type -> DeployPackageRequest
	5,  // 10: ServiceRunner.DescribePackage:output_type -> DescribePackageReply
	9,  // 11: ServiceRunner.DeletePackage:output_type -> DeletePackageReply
	7,  // 12: ServiceRunner.ListPackages:output_type -> ListPackagesReply
	3,  // 13: ServiceRunner.DeployPackage:output_type -> DeployPackageReply
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pb_sr_proto_init() }
func file_pb_sr_proto_init() {
	if File_pb_sr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_sr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageDescription); i {
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
		file_pb_sr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployPackageRequest); i {
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
		file_pb_sr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployPackageReply); i {
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
		file_pb_sr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePackageRequest); i {
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
		file_pb_sr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePackageReply); i {
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
		file_pb_sr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPackagesRequest); i {
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
		file_pb_sr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPackagesReply); i {
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
		file_pb_sr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageRequest); i {
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
		file_pb_sr_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageReply); i {
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
		file_pb_sr_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPackagesReply_ListPackagesItem); i {
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
			RawDescriptor: file_pb_sr_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_sr_proto_goTypes,
		DependencyIndexes: file_pb_sr_proto_depIdxs,
		EnumInfos:         file_pb_sr_proto_enumTypes,
		MessageInfos:      file_pb_sr_proto_msgTypes,
	}.Build()
	File_pb_sr_proto = out.File
	file_pb_sr_proto_rawDesc = nil
	file_pb_sr_proto_goTypes = nil
	file_pb_sr_proto_depIdxs = nil
}
