// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: csrf_dispatcher.proto

package csrf

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CsrfToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid   string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CsrfToken) Reset() {
	*x = CsrfToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csrf_dispatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsrfToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsrfToken) ProtoMessage() {}

func (x *CsrfToken) ProtoReflect() protoreflect.Message {
	mi := &file_csrf_dispatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsrfToken.ProtoReflect.Descriptor instead.
func (*CsrfToken) Descriptor() ([]byte, []int) {
	return file_csrf_dispatcher_proto_rawDescGZIP(), []int{0}
}

func (x *CsrfToken) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *CsrfToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CsrfData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid       string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	TimeStamp int64  `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *CsrfData) Reset() {
	*x = CsrfData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csrf_dispatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsrfData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsrfData) ProtoMessage() {}

func (x *CsrfData) ProtoReflect() protoreflect.Message {
	mi := &file_csrf_dispatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsrfData.ProtoReflect.Descriptor instead.
func (*CsrfData) Descriptor() ([]byte, []int) {
	return file_csrf_dispatcher_proto_rawDescGZIP(), []int{1}
}

func (x *CsrfData) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *CsrfData) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type IsValid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IsValid) Reset() {
	*x = IsValid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csrf_dispatcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValid) ProtoMessage() {}

func (x *IsValid) ProtoReflect() protoreflect.Message {
	mi := &file_csrf_dispatcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValid.ProtoReflect.Descriptor instead.
func (*IsValid) Descriptor() ([]byte, []int) {
	return file_csrf_dispatcher_proto_rawDescGZIP(), []int{2}
}

func (x *IsValid) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_csrf_dispatcher_proto protoreflect.FileDescriptor

var file_csrf_dispatcher_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x73, 0x72, 0x66, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x73, 0x72, 0x66, 0x22, 0x33, 0x0a,
	0x09, 0x43, 0x73, 0x72, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x3a, 0x0a, 0x08, 0x43, 0x73, 0x72, 0x66, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x1f,
	0x0a, 0x07, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x68, 0x0a, 0x0e, 0x43, 0x73, 0x72, 0x66, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x63, 0x73,
	0x72, 0x66, 0x2e, 0x43, 0x73, 0x72, 0x66, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x63, 0x73,
	0x72, 0x66, 0x2e, 0x43, 0x73, 0x72, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x29,
	0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x63, 0x73, 0x72, 0x66, 0x2e, 0x43,
	0x73, 0x72, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0d, 0x2e, 0x63, 0x73, 0x72, 0x66, 0x2e,
	0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_csrf_dispatcher_proto_rawDescOnce sync.Once
	file_csrf_dispatcher_proto_rawDescData = file_csrf_dispatcher_proto_rawDesc
)

func file_csrf_dispatcher_proto_rawDescGZIP() []byte {
	file_csrf_dispatcher_proto_rawDescOnce.Do(func() {
		file_csrf_dispatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_csrf_dispatcher_proto_rawDescData)
	})
	return file_csrf_dispatcher_proto_rawDescData
}

var file_csrf_dispatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_csrf_dispatcher_proto_goTypes = []interface{}{
	(*CsrfToken)(nil), // 0: csrf.CsrfToken
	(*CsrfData)(nil),  // 1: csrf.CsrfData
	(*IsValid)(nil),   // 2: csrf.IsValid
}
var file_csrf_dispatcher_proto_depIdxs = []int32{
	1, // 0: csrf.CsrfDispatcher.Create:input_type -> csrf.CsrfData
	0, // 1: csrf.CsrfDispatcher.Check:input_type -> csrf.CsrfToken
	0, // 2: csrf.CsrfDispatcher.Create:output_type -> csrf.CsrfToken
	2, // 3: csrf.CsrfDispatcher.Check:output_type -> csrf.IsValid
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_csrf_dispatcher_proto_init() }
func file_csrf_dispatcher_proto_init() {
	if File_csrf_dispatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_csrf_dispatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsrfToken); i {
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
		file_csrf_dispatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsrfData); i {
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
		file_csrf_dispatcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValid); i {
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
			RawDescriptor: file_csrf_dispatcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_csrf_dispatcher_proto_goTypes,
		DependencyIndexes: file_csrf_dispatcher_proto_depIdxs,
		MessageInfos:      file_csrf_dispatcher_proto_msgTypes,
	}.Build()
	File_csrf_dispatcher_proto = out.File
	file_csrf_dispatcher_proto_rawDesc = nil
	file_csrf_dispatcher_proto_goTypes = nil
	file_csrf_dispatcher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CsrfDispatcherClient is the client API for CsrfDispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CsrfDispatcherClient interface {
	Create(ctx context.Context, in *CsrfData, opts ...grpc.CallOption) (*CsrfToken, error)
	Check(ctx context.Context, in *CsrfToken, opts ...grpc.CallOption) (*IsValid, error)
}

type csrfDispatcherClient struct {
	cc grpc.ClientConnInterface
}

func NewCsrfDispatcherClient(cc grpc.ClientConnInterface) CsrfDispatcherClient {
	return &csrfDispatcherClient{cc}
}

func (c *csrfDispatcherClient) Create(ctx context.Context, in *CsrfData, opts ...grpc.CallOption) (*CsrfToken, error) {
	out := new(CsrfToken)
	err := c.cc.Invoke(ctx, "/csrf.CsrfDispatcher/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *csrfDispatcherClient) Check(ctx context.Context, in *CsrfToken, opts ...grpc.CallOption) (*IsValid, error) {
	out := new(IsValid)
	err := c.cc.Invoke(ctx, "/csrf.CsrfDispatcher/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CsrfDispatcherServer is the server API for CsrfDispatcher service.
type CsrfDispatcherServer interface {
	Create(context.Context, *CsrfData) (*CsrfToken, error)
	Check(context.Context, *CsrfToken) (*IsValid, error)
}

// UnimplementedCsrfDispatcherServer can be embedded to have forward compatible implementations.
type UnimplementedCsrfDispatcherServer struct {
}

func (*UnimplementedCsrfDispatcherServer) Create(context.Context, *CsrfData) (*CsrfToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCsrfDispatcherServer) Check(context.Context, *CsrfToken) (*IsValid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterCsrfDispatcherServer(s *grpc.Server, srv CsrfDispatcherServer) {
	s.RegisterService(&_CsrfDispatcher_serviceDesc, srv)
}

func _CsrfDispatcher_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CsrfData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CsrfDispatcherServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CsrfDispatcher/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CsrfDispatcherServer).Create(ctx, req.(*CsrfData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CsrfDispatcher_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CsrfToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CsrfDispatcherServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrf.CsrfDispatcher/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CsrfDispatcherServer).Check(ctx, req.(*CsrfToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _CsrfDispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "csrf.CsrfDispatcher",
	HandlerType: (*CsrfDispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CsrfDispatcher_Create_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _CsrfDispatcher_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "csrf_dispatcher.proto",
}
