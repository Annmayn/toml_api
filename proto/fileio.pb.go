// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileio.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FileName struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileName) Reset()         { *m = FileName{} }
func (m *FileName) String() string { return proto.CompactTextString(m) }
func (*FileName) ProtoMessage()    {}
func (*FileName) Descriptor() ([]byte, []int) {
	return fileDescriptor_524f11e5f161eed5, []int{0}
}

func (m *FileName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileName.Unmarshal(m, b)
}
func (m *FileName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileName.Marshal(b, m, deterministic)
}
func (m *FileName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileName.Merge(m, src)
}
func (m *FileName) XXX_Size() int {
	return xxx_messageInfo_FileName.Size(m)
}
func (m *FileName) XXX_DiscardUnknown() {
	xxx_messageInfo_FileName.DiscardUnknown(m)
}

var xxx_messageInfo_FileName proto.InternalMessageInfo

func (m *FileName) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type FileData struct {
	FileName             string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileData) Reset()         { *m = FileData{} }
func (m *FileData) String() string { return proto.CompactTextString(m) }
func (*FileData) ProtoMessage()    {}
func (*FileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_524f11e5f161eed5, []int{1}
}

func (m *FileData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileData.Unmarshal(m, b)
}
func (m *FileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileData.Marshal(b, m, deterministic)
}
func (m *FileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileData.Merge(m, src)
}
func (m *FileData) XXX_Size() int {
	return xxx_messageInfo_FileData.Size(m)
}
func (m *FileData) XXX_DiscardUnknown() {
	xxx_messageInfo_FileData.DiscardUnknown(m)
}

var xxx_messageInfo_FileData proto.InternalMessageInfo

func (m *FileData) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileData) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type FileContent struct {
	Content              []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileContent) Reset()         { *m = FileContent{} }
func (m *FileContent) String() string { return proto.CompactTextString(m) }
func (*FileContent) ProtoMessage()    {}
func (*FileContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_524f11e5f161eed5, []int{2}
}

func (m *FileContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileContent.Unmarshal(m, b)
}
func (m *FileContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileContent.Marshal(b, m, deterministic)
}
func (m *FileContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileContent.Merge(m, src)
}
func (m *FileContent) XXX_Size() int {
	return xxx_messageInfo_FileContent.Size(m)
}
func (m *FileContent) XXX_DiscardUnknown() {
	xxx_messageInfo_FileContent.DiscardUnknown(m)
}

var xxx_messageInfo_FileContent proto.InternalMessageInfo

func (m *FileContent) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_524f11e5f161eed5, []int{3}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FileName)(nil), "proto.FileName")
	proto.RegisterType((*FileData)(nil), "proto.FileData")
	proto.RegisterType((*FileContent)(nil), "proto.FileContent")
	proto.RegisterType((*EmptyResponse)(nil), "proto.EmptyResponse")
}

func init() { proto.RegisterFile("fileio.proto", fileDescriptor_524f11e5f161eed5) }

var fileDescriptor_524f11e5f161eed5 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0xcc, 0x49,
	0xcd, 0xcc, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6a, 0x5c, 0x1c,
	0x6e, 0x99, 0x39, 0xa9, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x69, 0x50, 0xb6, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0xaf, 0xe4, 0x00, 0x51, 0xe7, 0x92, 0x58, 0x92, 0x88,
	0xa2, 0x8e, 0x09, 0x55, 0x9d, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89,
	0x04, 0x33, 0x58, 0x0a, 0xc6, 0x55, 0x52, 0xe7, 0xe2, 0x06, 0x99, 0xe0, 0x0c, 0xe1, 0x22, 0x2b,
	0x64, 0x51, 0x60, 0xd4, 0xe0, 0x41, 0x28, 0xe4, 0xe7, 0xe2, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x0c,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0xb2, 0xe5, 0xe2, 0x72, 0x4c, 0x49, 0x09, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xd2, 0xe7, 0xe2, 0x08, 0x4a, 0x4d, 0x4c, 0x01, 0x99, 0x25,
	0xc4, 0x0f, 0xf1, 0x8c, 0x1e, 0xcc, 0x0b, 0x52, 0x42, 0x48, 0x02, 0x50, 0x9b, 0x92, 0xd8, 0xc0,
	0x42, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x09, 0x28, 0x03, 0x00, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	ReadFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileContent, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) ReadFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileContent, error) {
	out := new(FileContent)
	err := c.cc.Invoke(ctx, "/proto.AddService/ReadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	ReadFile(context.Context, *FileName) (*FileContent, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) ReadFile(ctx context.Context, req *FileName) (*FileContent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFile not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/ReadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).ReadFile(ctx, req.(*FileName))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadFile",
			Handler:    _AddService_ReadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fileio.proto",
}
