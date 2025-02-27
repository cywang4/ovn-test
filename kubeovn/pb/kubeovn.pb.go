// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: kubeovn.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateVPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *CreateVPCRequest) Reset() {
	*x = CreateVPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubeovn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVPCRequest) ProtoMessage() {}

func (x *CreateVPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kubeovn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVPCRequest.ProtoReflect.Descriptor instead.
func (*CreateVPCRequest) Descriptor() ([]byte, []int) {
	return file_kubeovn_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVPCRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVPCRequest) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type CreateVPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateVPCResponse) Reset() {
	*x = CreateVPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubeovn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVPCResponse) ProtoMessage() {}

func (x *CreateVPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kubeovn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVPCResponse.ProtoReflect.Descriptor instead.
func (*CreateVPCResponse) Descriptor() ([]byte, []int) {
	return file_kubeovn_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVPCResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateVPCResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CreateSubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VpcName   string `protobuf:"bytes,2,opt,name=vpc_name,json=vpcName,proto3" json:"vpc_name,omitempty"`
	CidrBlock string `protobuf:"bytes,3,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	Gateway   string `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Protocol  string `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *CreateSubnetRequest) Reset() {
	*x = CreateSubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubeovn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubnetRequest) ProtoMessage() {}

func (x *CreateSubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kubeovn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubnetRequest.ProtoReflect.Descriptor instead.
func (*CreateSubnetRequest) Descriptor() ([]byte, []int) {
	return file_kubeovn_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSubnetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSubnetRequest) GetVpcName() string {
	if x != nil {
		return x.VpcName
	}
	return ""
}

func (x *CreateSubnetRequest) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *CreateSubnetRequest) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *CreateSubnetRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

type CreateSubnetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateSubnetResponse) Reset() {
	*x = CreateSubnetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubeovn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubnetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubnetResponse) ProtoMessage() {}

func (x *CreateSubnetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kubeovn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubnetResponse.ProtoReflect.Descriptor instead.
func (*CreateSubnetResponse) Descriptor() ([]byte, []int) {
	return file_kubeovn_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSubnetResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateSubnetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_kubeovn_proto protoreflect.FileDescriptor

var file_kubeovn_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x6f, 0x76, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6b, 0x75, 0x62, 0x65, 0x6f, 0x76, 0x6e, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x22, 0x41, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x50, 0x43, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x99, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x70, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x70, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69,
	0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22,
	0x44, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xa1, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x50, 0x43, 0x12, 0x19, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x6f, 0x76, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x6f, 0x76, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x6f, 0x76, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x6f, 0x76, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kubeovn_proto_rawDescOnce sync.Once
	file_kubeovn_proto_rawDescData = file_kubeovn_proto_rawDesc
)

func file_kubeovn_proto_rawDescGZIP() []byte {
	file_kubeovn_proto_rawDescOnce.Do(func() {
		file_kubeovn_proto_rawDescData = protoimpl.X.CompressGZIP(file_kubeovn_proto_rawDescData)
	})
	return file_kubeovn_proto_rawDescData
}

var file_kubeovn_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kubeovn_proto_goTypes = []interface{}{
	(*CreateVPCRequest)(nil),     // 0: kubeovn.CreateVPCRequest
	(*CreateVPCResponse)(nil),    // 1: kubeovn.CreateVPCResponse
	(*CreateSubnetRequest)(nil),  // 2: kubeovn.CreateSubnetRequest
	(*CreateSubnetResponse)(nil), // 3: kubeovn.CreateSubnetResponse
}
var file_kubeovn_proto_depIdxs = []int32{
	0, // 0: kubeovn.NetworkService.CreateVPC:input_type -> kubeovn.CreateVPCRequest
	2, // 1: kubeovn.NetworkService.CreateSubnet:input_type -> kubeovn.CreateSubnetRequest
	1, // 2: kubeovn.NetworkService.CreateVPC:output_type -> kubeovn.CreateVPCResponse
	3, // 3: kubeovn.NetworkService.CreateSubnet:output_type -> kubeovn.CreateSubnetResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kubeovn_proto_init() }
func file_kubeovn_proto_init() {
	if File_kubeovn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kubeovn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVPCRequest); i {
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
		file_kubeovn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVPCResponse); i {
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
		file_kubeovn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubnetRequest); i {
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
		file_kubeovn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubnetResponse); i {
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
			RawDescriptor: file_kubeovn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kubeovn_proto_goTypes,
		DependencyIndexes: file_kubeovn_proto_depIdxs,
		MessageInfos:      file_kubeovn_proto_msgTypes,
	}.Build()
	File_kubeovn_proto = out.File
	file_kubeovn_proto_rawDesc = nil
	file_kubeovn_proto_goTypes = nil
	file_kubeovn_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	CreateVPC(ctx context.Context, in *CreateVPCRequest, opts ...grpc.CallOption) (*CreateVPCResponse, error)
	CreateSubnet(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*CreateSubnetResponse, error)
}

type networkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServiceClient(cc grpc.ClientConnInterface) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) CreateVPC(ctx context.Context, in *CreateVPCRequest, opts ...grpc.CallOption) (*CreateVPCResponse, error) {
	out := new(CreateVPCResponse)
	err := c.cc.Invoke(ctx, "/kubeovn.NetworkService/CreateVPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) CreateSubnet(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*CreateSubnetResponse, error) {
	out := new(CreateSubnetResponse)
	err := c.cc.Invoke(ctx, "/kubeovn.NetworkService/CreateSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	CreateVPC(context.Context, *CreateVPCRequest) (*CreateVPCResponse, error)
	CreateSubnet(context.Context, *CreateSubnetRequest) (*CreateSubnetResponse, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) CreateVPC(context.Context, *CreateVPCRequest) (*CreateVPCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVPC not implemented")
}
func (*UnimplementedNetworkServiceServer) CreateSubnet(context.Context, *CreateSubnetRequest) (*CreateSubnetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubnet not implemented")
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_CreateVPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).CreateVPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeovn.NetworkService/CreateVPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).CreateVPC(ctx, req.(*CreateVPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_CreateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).CreateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeovn.NetworkService/CreateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).CreateSubnet(ctx, req.(*CreateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubeovn.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVPC",
			Handler:    _NetworkService_CreateVPC_Handler,
		},
		{
			MethodName: "CreateSubnet",
			Handler:    _NetworkService_CreateSubnet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeovn.proto",
}
