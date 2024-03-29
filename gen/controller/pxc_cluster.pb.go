// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: controller/pxc_cluster.proto

package controllerv1beta1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// PXCBackupState represents PXC backup CR state.
type PXCBackupState int32

const (
	// PXC_BACKUP_STATE_INVALID represents unknown state.
	PXCBackupState_PXC_BACKUP_STATE_INVALID PXCBackupState = 0
	// PXC_BACKUP_STATE_RUNNING represents running backup (Starting, Running, FIXME check it).
	PXCBackupState_PXC_BACKUP_STATE_RUNNING PXCBackupState = 1
	// PXC_BACKUP_STATE_SUCCEEDED represents succeeded backup (Succeeded, FIXME check it).
	PXCBackupState_PXC_BACKUP_STATE_SUCCEEDED PXCBackupState = 2
	// PXC_BACKUP_STATE_FAILED represents failed backup (Failed, FIXME check it).
	PXCBackupState_PXC_BACKUP_STATE_FAILED PXCBackupState = 3
)

// Enum value maps for PXCBackupState.
var (
	PXCBackupState_name = map[int32]string{
		0: "PXC_BACKUP_STATE_INVALID",
		1: "PXC_BACKUP_STATE_RUNNING",
		2: "PXC_BACKUP_STATE_SUCCEEDED",
		3: "PXC_BACKUP_STATE_FAILED",
	}
	PXCBackupState_value = map[string]int32{
		"PXC_BACKUP_STATE_INVALID":   0,
		"PXC_BACKUP_STATE_RUNNING":   1,
		"PXC_BACKUP_STATE_SUCCEEDED": 2,
		"PXC_BACKUP_STATE_FAILED":    3,
	}
)

func (x PXCBackupState) Enum() *PXCBackupState {
	p := new(PXCBackupState)
	*p = x
	return p
}

func (x PXCBackupState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PXCBackupState) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_pxc_cluster_proto_enumTypes[0].Descriptor()
}

func (PXCBackupState) Type() protoreflect.EnumType {
	return &file_controller_pxc_cluster_proto_enumTypes[0]
}

func (x PXCBackupState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PXCBackupState.Descriptor instead.
func (PXCBackupState) EnumDescriptor() ([]byte, []int) {
	return file_controller_pxc_cluster_proto_rawDescGZIP(), []int{0}
}

// PXCClusterParams represents PXC cluster parameters that can be updated.
type PXCClusterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster size.
	ClusterSize int32 `protobuf:"varint,1,opt,name=cluster_size,json=clusterSize,proto3" json:"cluster_size,omitempty"`
	// PXC container parameters.
	Pxc *PXCClusterParams_PXC `protobuf:"bytes,2,opt,name=pxc,proto3" json:"pxc,omitempty"`
	// NOTE: Only one of fields proxysql or haproxy has to be set, we check this in the code.
	// ProxySQL container parameters.
	Proxysql *PXCClusterParams_ProxySQL `protobuf:"bytes,3,opt,name=proxysql,proto3" json:"proxysql,omitempty"`
	// HAProxy container parameters.
	Haproxy *PXCClusterParams_HAProxy `protobuf:"bytes,4,opt,name=haproxy,proto3" json:"haproxy,omitempty"`
	// Version service URL. We need to pass the URL because operators use it to fetch information about versions during upgrade.
	// We want the URL to match the one used in pmm-managed because we can use custom version service.
	VersionServiceUrl string `protobuf:"bytes,5,opt,name=version_service_url,json=versionServiceUrl,proto3" json:"version_service_url,omitempty"`
}

func (x *PXCClusterParams) Reset() {
	*x = PXCClusterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_pxc_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PXCClusterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PXCClusterParams) ProtoMessage() {}

func (x *PXCClusterParams) ProtoReflect() protoreflect.Message {
	mi := &file_controller_pxc_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PXCClusterParams.ProtoReflect.Descriptor instead.
func (*PXCClusterParams) Descriptor() ([]byte, []int) {
	return file_controller_pxc_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *PXCClusterParams) GetClusterSize() int32 {
	if x != nil {
		return x.ClusterSize
	}
	return 0
}

func (x *PXCClusterParams) GetPxc() *PXCClusterParams_PXC {
	if x != nil {
		return x.Pxc
	}
	return nil
}

func (x *PXCClusterParams) GetProxysql() *PXCClusterParams_ProxySQL {
	if x != nil {
		return x.Proxysql
	}
	return nil
}

func (x *PXCClusterParams) GetHaproxy() *PXCClusterParams_HAProxy {
	if x != nil {
		return x.Haproxy
	}
	return nil
}

func (x *PXCClusterParams) GetVersionServiceUrl() string {
	if x != nil {
		return x.VersionServiceUrl
	}
	return ""
}

// PXCCredentials is cluster connection credentials.
type PXCCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PXC host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// PXC port.
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// PXC username.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// PXC password.
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PXCCredentials) Reset() {
	*x = PXCCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_pxc_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PXCCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PXCCredentials) ProtoMessage() {}

func (x *PXCCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_controller_pxc_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PXCCredentials.ProtoReflect.Descriptor instead.
func (*PXCCredentials) Descriptor() ([]byte, []int) {
	return file_controller_pxc_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *PXCCredentials) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PXCCredentials) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PXCCredentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PXCCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// PXC container parameters.
type PXCClusterParams_PXC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Docker image used for PXC.
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	// Requested compute resources.
	ComputeResources *ComputeResources `protobuf:"bytes,1,opt,name=compute_resources,json=computeResources,proto3" json:"compute_resources,omitempty"`
	// Disk size in bytes.
	DiskSize int64 `protobuf:"varint,2,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
}

func (x *PXCClusterParams_PXC) Reset() {
	*x = PXCClusterParams_PXC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_pxc_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PXCClusterParams_PXC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PXCClusterParams_PXC) ProtoMessage() {}

func (x *PXCClusterParams_PXC) ProtoReflect() protoreflect.Message {
	mi := &file_controller_pxc_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PXCClusterParams_PXC.ProtoReflect.Descriptor instead.
func (*PXCClusterParams_PXC) Descriptor() ([]byte, []int) {
	return file_controller_pxc_cluster_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PXCClusterParams_PXC) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PXCClusterParams_PXC) GetComputeResources() *ComputeResources {
	if x != nil {
		return x.ComputeResources
	}
	return nil
}

func (x *PXCClusterParams_PXC) GetDiskSize() int64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

// ProxySQL container parameters.
type PXCClusterParams_ProxySQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Docker image used for ProxySQL.
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	// Requested compute resources.
	ComputeResources *ComputeResources `protobuf:"bytes,1,opt,name=compute_resources,json=computeResources,proto3" json:"compute_resources,omitempty"`
	// Disk size in bytes.
	DiskSize int64 `protobuf:"varint,2,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
}

func (x *PXCClusterParams_ProxySQL) Reset() {
	*x = PXCClusterParams_ProxySQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_pxc_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PXCClusterParams_ProxySQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PXCClusterParams_ProxySQL) ProtoMessage() {}

func (x *PXCClusterParams_ProxySQL) ProtoReflect() protoreflect.Message {
	mi := &file_controller_pxc_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PXCClusterParams_ProxySQL.ProtoReflect.Descriptor instead.
func (*PXCClusterParams_ProxySQL) Descriptor() ([]byte, []int) {
	return file_controller_pxc_cluster_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PXCClusterParams_ProxySQL) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PXCClusterParams_ProxySQL) GetComputeResources() *ComputeResources {
	if x != nil {
		return x.ComputeResources
	}
	return nil
}

func (x *PXCClusterParams_ProxySQL) GetDiskSize() int64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

// HAProxy container parameters.
// NOTE: HAProxy does not need disk size as ProxySQL because the container does not require it.
type PXCClusterParams_HAProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Docker image used for HAProxy.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Requested compute resources.
	ComputeResources *ComputeResources `protobuf:"bytes,2,opt,name=compute_resources,json=computeResources,proto3" json:"compute_resources,omitempty"`
}

func (x *PXCClusterParams_HAProxy) Reset() {
	*x = PXCClusterParams_HAProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_pxc_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PXCClusterParams_HAProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PXCClusterParams_HAProxy) ProtoMessage() {}

func (x *PXCClusterParams_HAProxy) ProtoReflect() protoreflect.Message {
	mi := &file_controller_pxc_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PXCClusterParams_HAProxy.ProtoReflect.Descriptor instead.
func (*PXCClusterParams_HAProxy) Descriptor() ([]byte, []int) {
	return file_controller_pxc_cluster_proto_rawDescGZIP(), []int{0, 2}
}

func (x *PXCClusterParams_HAProxy) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PXCClusterParams_HAProxy) GetComputeResources() *ComputeResources {
	if x != nil {
		return x.ComputeResources
	}
	return nil
}

var File_controller_pxc_cluster_proto protoreflect.FileDescriptor

var file_controller_pxc_cluster_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x78, 0x63,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29,
	0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x07, 0x0a, 0x10, 0x50,
	0x58, 0x43, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x29, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x59, 0x0a, 0x03, 0x70, 0x78,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x58, 0x43, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x2e, 0x50, 0x58, 0x43, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01,
	0x52, 0x03, 0x70, 0x78, 0x63, 0x12, 0x60, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x71,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x58, 0x43, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x73, 0x71, 0x6c, 0x12, 0x5d, 0x0a, 0x07, 0x68, 0x61, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x58, 0x43, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x07, 0x68,
	0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x1a, 0xb2, 0x01, 0x0a, 0x03, 0x50, 0x58, 0x43, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10,
	0x00, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0xb7, 0x01, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x70,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x70, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61,
	0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x10,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x6b, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x89, 0x01, 0x0a, 0x07, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x22, 0x70, 0x0a, 0x0e, 0x50, 0x58, 0x43, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x2a, 0x89, 0x01, 0x0a, 0x0e, 0x50, 0x58, 0x43, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x58, 0x43, 0x5f, 0x42, 0x41,
	0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x58, 0x43, 0x5f, 0x42, 0x41, 0x43, 0x4b,
	0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x58, 0x43, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x58, 0x43, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42,
	0x1e, 0x5a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x3b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_pxc_cluster_proto_rawDescOnce sync.Once
	file_controller_pxc_cluster_proto_rawDescData = file_controller_pxc_cluster_proto_rawDesc
)

func file_controller_pxc_cluster_proto_rawDescGZIP() []byte {
	file_controller_pxc_cluster_proto_rawDescOnce.Do(func() {
		file_controller_pxc_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_pxc_cluster_proto_rawDescData)
	})
	return file_controller_pxc_cluster_proto_rawDescData
}

var file_controller_pxc_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controller_pxc_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_controller_pxc_cluster_proto_goTypes = []interface{}{
	(PXCBackupState)(0),               // 0: percona.platform.dbaas.controller.v1beta1.PXCBackupState
	(*PXCClusterParams)(nil),          // 1: percona.platform.dbaas.controller.v1beta1.PXCClusterParams
	(*PXCCredentials)(nil),            // 2: percona.platform.dbaas.controller.v1beta1.PXCCredentials
	(*PXCClusterParams_PXC)(nil),      // 3: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.PXC
	(*PXCClusterParams_ProxySQL)(nil), // 4: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.ProxySQL
	(*PXCClusterParams_HAProxy)(nil),  // 5: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.HAProxy
	(*ComputeResources)(nil),          // 6: percona.platform.dbaas.controller.v1beta1.ComputeResources
}
var file_controller_pxc_cluster_proto_depIdxs = []int32{
	3, // 0: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.pxc:type_name -> percona.platform.dbaas.controller.v1beta1.PXCClusterParams.PXC
	4, // 1: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.proxysql:type_name -> percona.platform.dbaas.controller.v1beta1.PXCClusterParams.ProxySQL
	5, // 2: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.haproxy:type_name -> percona.platform.dbaas.controller.v1beta1.PXCClusterParams.HAProxy
	6, // 3: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.PXC.compute_resources:type_name -> percona.platform.dbaas.controller.v1beta1.ComputeResources
	6, // 4: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.ProxySQL.compute_resources:type_name -> percona.platform.dbaas.controller.v1beta1.ComputeResources
	6, // 5: percona.platform.dbaas.controller.v1beta1.PXCClusterParams.HAProxy.compute_resources:type_name -> percona.platform.dbaas.controller.v1beta1.ComputeResources
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_controller_pxc_cluster_proto_init() }
func file_controller_pxc_cluster_proto_init() {
	if File_controller_pxc_cluster_proto != nil {
		return
	}
	file_controller_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_pxc_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PXCClusterParams); i {
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
		file_controller_pxc_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PXCCredentials); i {
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
		file_controller_pxc_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PXCClusterParams_PXC); i {
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
		file_controller_pxc_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PXCClusterParams_ProxySQL); i {
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
		file_controller_pxc_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PXCClusterParams_HAProxy); i {
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
			RawDescriptor: file_controller_pxc_cluster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_pxc_cluster_proto_goTypes,
		DependencyIndexes: file_controller_pxc_cluster_proto_depIdxs,
		EnumInfos:         file_controller_pxc_cluster_proto_enumTypes,
		MessageInfos:      file_controller_pxc_cluster_proto_msgTypes,
	}.Build()
	File_controller_pxc_cluster_proto = out.File
	file_controller_pxc_cluster_proto_rawDesc = nil
	file_controller_pxc_cluster_proto_goTypes = nil
	file_controller_pxc_cluster_proto_depIdxs = nil
}
