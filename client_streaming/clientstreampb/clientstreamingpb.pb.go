// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: client_streaming/clientstreampb/clientstreamingpb.proto

package clientstreampb

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

// greet request
type ClientStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ClientStreamRequest) Reset() {
	*x = ClientStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamRequest.ProtoReflect.Descriptor instead.
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescGZIP(), []int{0}
}

func (x *ClientStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// greet response
type ClientStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *ClientStreamResponse) Reset() {
	*x = ClientStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamResponse) ProtoMessage() {}

func (x *ClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamResponse.ProtoReflect.Descriptor instead.
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescGZIP(), []int{1}
}

func (x *ClientStreamResponse) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_client_streaming_clientstreampb_clientstreamingpb_proto protoreflect.FileDescriptor

var file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x70,
	0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x22, 0x29, 0x0a, 0x13,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65,
	0x73, 0x32, 0x6c, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5c, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x22, 0x5a, 0x20, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescOnce sync.Once
	file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescData = file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDesc
)

func file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescGZIP() []byte {
	file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescOnce.Do(func() {
		file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescData)
	})
	return file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDescData
}

var file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_streaming_clientstreampb_clientstreamingpb_proto_goTypes = []interface{}{
	(*ClientStreamRequest)(nil),  // 0: clientstreamingpb.ClientStreamRequest
	(*ClientStreamResponse)(nil), // 1: clientstreamingpb.ClientStreamResponse
}
var file_client_streaming_clientstreampb_clientstreamingpb_proto_depIdxs = []int32{
	0, // 0: clientstreamingpb.GreetService.Greet:input_type -> clientstreamingpb.ClientStreamRequest
	1, // 1: clientstreamingpb.GreetService.Greet:output_type -> clientstreamingpb.ClientStreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_streaming_clientstreampb_clientstreamingpb_proto_init() }
func file_client_streaming_clientstreampb_clientstreamingpb_proto_init() {
	if File_client_streaming_clientstreampb_clientstreamingpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamRequest); i {
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
		file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamResponse); i {
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
			RawDescriptor: file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_streaming_clientstreampb_clientstreamingpb_proto_goTypes,
		DependencyIndexes: file_client_streaming_clientstreampb_clientstreamingpb_proto_depIdxs,
		MessageInfos:      file_client_streaming_clientstreampb_clientstreamingpb_proto_msgTypes,
	}.Build()
	File_client_streaming_clientstreampb_clientstreamingpb_proto = out.File
	file_client_streaming_clientstreampb_clientstreamingpb_proto_rawDesc = nil
	file_client_streaming_clientstreampb_clientstreamingpb_proto_goTypes = nil
	file_client_streaming_clientstreampb_clientstreamingpb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// unary API
	Greet(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/clientstreamingpb.GreetService/Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetClient{stream}
	return x, nil
}

type GreetService_GreetClient interface {
	Send(*ClientStreamRequest) error
	CloseAndRecv() (*ClientStreamResponse, error)
	grpc.ClientStream
}

type greetServiceGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetClient) Send(m *ClientStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetClient) CloseAndRecv() (*ClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// unary API
	Greet(GreetService_GreetServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(GreetService_GreetServer) error {
	return status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).Greet(&greetServiceGreetServer{stream})
}

type GreetService_GreetServer interface {
	SendAndClose(*ClientStreamResponse) error
	Recv() (*ClientStreamRequest, error)
	grpc.ServerStream
}

type greetServiceGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetServer) SendAndClose(m *ClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetServer) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clientstreamingpb.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet",
			Handler:       _GreetService_Greet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client_streaming/clientstreampb/clientstreamingpb.proto",
}
