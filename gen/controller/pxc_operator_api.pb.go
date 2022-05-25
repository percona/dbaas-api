// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: controller/pxc_operator_api.proto

package controllerv1beta1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InstallPXCOperatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes auth.
	KubeAuth *KubeAuth `protobuf:"bytes,1,opt,name=kube_auth,json=kubeAuth,proto3" json:"kube_auth,omitempty"`
	// Operator version to be installed.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *InstallPXCOperatorRequest) Reset() {
	*x = InstallPXCOperatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_pxc_operator_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallPXCOperatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallPXCOperatorRequest) ProtoMessage() {}

func (x *InstallPXCOperatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_pxc_operator_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallPXCOperatorRequest.ProtoReflect.Descriptor instead.
func (*InstallPXCOperatorRequest) Descriptor() ([]byte, []int) {
	return file_controller_pxc_operator_api_proto_rawDescGZIP(), []int{0}
}

func (x *InstallPXCOperatorRequest) GetKubeAuth() *KubeAuth {
	if x != nil {
		return x.KubeAuth
	}
	return nil
}

func (x *InstallPXCOperatorRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type InstallPXCOperatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InstallPXCOperatorResponse) Reset() {
	*x = InstallPXCOperatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_pxc_operator_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallPXCOperatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallPXCOperatorResponse) ProtoMessage() {}

func (x *InstallPXCOperatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_pxc_operator_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallPXCOperatorResponse.ProtoReflect.Descriptor instead.
func (*InstallPXCOperatorResponse) Descriptor() ([]byte, []int) {
	return file_controller_pxc_operator_api_proto_rawDescGZIP(), []int{1}
}

var File_controller_pxc_operator_api_proto protoreflect.FileDescriptor

var file_controller_pxc_operator_api_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x78, 0x63,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x29, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x17,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x97, 0x01, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x58, 0x43, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x58, 0x0a,
	0x09, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x6b,
	0x75, 0x62, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1c, 0x0a, 0x1a, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x50, 0x58, 0x43, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb4, 0x01, 0x0a, 0x0e, 0x50, 0x58, 0x43, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x50, 0x49, 0x12, 0xa1, 0x01, 0x0a, 0x12, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x58, 0x43, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x44, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x58, 0x43, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x50, 0x58, 0x43, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e,
	0x5a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x3b, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_pxc_operator_api_proto_rawDescOnce sync.Once
	file_controller_pxc_operator_api_proto_rawDescData = file_controller_pxc_operator_api_proto_rawDesc
)

func file_controller_pxc_operator_api_proto_rawDescGZIP() []byte {
	file_controller_pxc_operator_api_proto_rawDescOnce.Do(func() {
		file_controller_pxc_operator_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_pxc_operator_api_proto_rawDescData)
	})
	return file_controller_pxc_operator_api_proto_rawDescData
}

var file_controller_pxc_operator_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_pxc_operator_api_proto_goTypes = []interface{}{
	(*InstallPXCOperatorRequest)(nil),  // 0: percona.platform.dbaas.controller.v1beta1.InstallPXCOperatorRequest
	(*InstallPXCOperatorResponse)(nil), // 1: percona.platform.dbaas.controller.v1beta1.InstallPXCOperatorResponse
	(*KubeAuth)(nil),                   // 2: percona.platform.dbaas.controller.v1beta1.KubeAuth
}
var file_controller_pxc_operator_api_proto_depIdxs = []int32{
	2, // 0: percona.platform.dbaas.controller.v1beta1.InstallPXCOperatorRequest.kube_auth:type_name -> percona.platform.dbaas.controller.v1beta1.KubeAuth
	0, // 1: percona.platform.dbaas.controller.v1beta1.PXCOperatorAPI.InstallPXCOperator:input_type -> percona.platform.dbaas.controller.v1beta1.InstallPXCOperatorRequest
	1, // 2: percona.platform.dbaas.controller.v1beta1.PXCOperatorAPI.InstallPXCOperator:output_type -> percona.platform.dbaas.controller.v1beta1.InstallPXCOperatorResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_controller_pxc_operator_api_proto_init() }
func file_controller_pxc_operator_api_proto_init() {
	if File_controller_pxc_operator_api_proto != nil {
		return
	}
	file_controller_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_pxc_operator_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallPXCOperatorRequest); i {
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
		file_controller_pxc_operator_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallPXCOperatorResponse); i {
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
			RawDescriptor: file_controller_pxc_operator_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_pxc_operator_api_proto_goTypes,
		DependencyIndexes: file_controller_pxc_operator_api_proto_depIdxs,
		MessageInfos:      file_controller_pxc_operator_api_proto_msgTypes,
	}.Build()
	File_controller_pxc_operator_api_proto = out.File
	file_controller_pxc_operator_api_proto_rawDesc = nil
	file_controller_pxc_operator_api_proto_goTypes = nil
	file_controller_pxc_operator_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PXCOperatorAPIClient is the client API for PXCOperatorAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PXCOperatorAPIClient interface {
	// InstallPXCOperator installs PXC operator.
	InstallPXCOperator(ctx context.Context, in *InstallPXCOperatorRequest, opts ...grpc.CallOption) (*InstallPXCOperatorResponse, error)
}

type pXCOperatorAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPXCOperatorAPIClient(cc grpc.ClientConnInterface) PXCOperatorAPIClient {
	return &pXCOperatorAPIClient{cc}
}

func (c *pXCOperatorAPIClient) InstallPXCOperator(ctx context.Context, in *InstallPXCOperatorRequest, opts ...grpc.CallOption) (*InstallPXCOperatorResponse, error) {
	out := new(InstallPXCOperatorResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.dbaas.controller.v1beta1.PXCOperatorAPI/InstallPXCOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PXCOperatorAPIServer is the server API for PXCOperatorAPI service.
type PXCOperatorAPIServer interface {
	// InstallPXCOperator installs PXC operator.
	InstallPXCOperator(context.Context, *InstallPXCOperatorRequest) (*InstallPXCOperatorResponse, error)
}

// UnimplementedPXCOperatorAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPXCOperatorAPIServer struct {
}

func (*UnimplementedPXCOperatorAPIServer) InstallPXCOperator(context.Context, *InstallPXCOperatorRequest) (*InstallPXCOperatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallPXCOperator not implemented")
}

func RegisterPXCOperatorAPIServer(s *grpc.Server, srv PXCOperatorAPIServer) {
	s.RegisterService(&_PXCOperatorAPI_serviceDesc, srv)
}

func _PXCOperatorAPI_InstallPXCOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallPXCOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PXCOperatorAPIServer).InstallPXCOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.dbaas.controller.v1beta1.PXCOperatorAPI/InstallPXCOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PXCOperatorAPIServer).InstallPXCOperator(ctx, req.(*InstallPXCOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PXCOperatorAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "percona.platform.dbaas.controller.v1beta1.PXCOperatorAPI",
	HandlerType: (*PXCOperatorAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstallPXCOperator",
			Handler:    _PXCOperatorAPI_InstallPXCOperator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/pxc_operator_api.proto",
}
