// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multiav.comodo.proto

package comodo_api

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

// The scan file request message containing the file path to scan.
type ScanFileRequest struct {
	Filepath             string   `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanFileRequest) Reset()         { *m = ScanFileRequest{} }
func (m *ScanFileRequest) String() string { return proto.CompactTextString(m) }
func (*ScanFileRequest) ProtoMessage()    {}
func (*ScanFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20dee691326e42ac, []int{0}
}

func (m *ScanFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanFileRequest.Unmarshal(m, b)
}
func (m *ScanFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanFileRequest.Marshal(b, m, deterministic)
}
func (m *ScanFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanFileRequest.Merge(m, src)
}
func (m *ScanFileRequest) XXX_Size() int {
	return xxx_messageInfo_ScanFileRequest.Size(m)
}
func (m *ScanFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanFileRequest proto.InternalMessageInfo

func (m *ScanFileRequest) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

// The scan response message containing detection results of the AntiVirus.
type ScanResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Infected             bool     `protobuf:"varint,2,opt,name=infected,proto3" json:"infected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanResponse) Reset()         { *m = ScanResponse{} }
func (m *ScanResponse) String() string { return proto.CompactTextString(m) }
func (*ScanResponse) ProtoMessage()    {}
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20dee691326e42ac, []int{1}
}

func (m *ScanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanResponse.Unmarshal(m, b)
}
func (m *ScanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanResponse.Marshal(b, m, deterministic)
}
func (m *ScanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanResponse.Merge(m, src)
}
func (m *ScanResponse) XXX_Size() int {
	return xxx_messageInfo_ScanResponse.Size(m)
}
func (m *ScanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScanResponse proto.InternalMessageInfo

func (m *ScanResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *ScanResponse) GetInfected() bool {
	if m != nil {
		return m.Infected
	}
	return false
}

// The version request message ask for version.
type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20dee691326e42ac, []int{2}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

// The response message containing program/VPS version.
type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20dee691326e42ac, []int{3}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*ScanFileRequest)(nil), "comodo.api.ScanFileRequest")
	proto.RegisterType((*ScanResponse)(nil), "comodo.api.ScanResponse")
	proto.RegisterType((*VersionRequest)(nil), "comodo.api.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "comodo.api.VersionResponse")
}

func init() { proto.RegisterFile("multiav.comodo.proto", fileDescriptor_20dee691326e42ac) }

var fileDescriptor_20dee691326e42ac = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x5d, 0x0f, 0x75, 0x1d, 0xd4, 0x4a, 0x10, 0x59, 0xd2, 0x4b, 0xc9, 0xa9, 0x20, 0xe6,
	0xa0, 0xff, 0xc0, 0x82, 0xbd, 0x47, 0xf0, 0x1e, 0xb7, 0x53, 0x0c, 0x6c, 0x33, 0x71, 0x33, 0xe9,
	0xbf, 0xf1, 0xbf, 0xca, 0xd6, 0x64, 0xd5, 0xc5, 0xe3, 0x37, 0xf3, 0x78, 0xf3, 0xde, 0xc0, 0xcd,
	0x3e, 0x75, 0xec, 0xec, 0x41, 0xb7, 0xb4, 0xa7, 0x2d, 0xe9, 0xd0, 0x13, 0x93, 0x80, 0x4c, 0x36,
	0x38, 0x75, 0x0f, 0xf3, 0x97, 0xd6, 0xfa, 0x67, 0xd7, 0xa1, 0xc1, 0x8f, 0x84, 0x91, 0x85, 0x84,
	0x7a, 0xe7, 0x3a, 0x0c, 0x96, 0xdf, 0x9b, 0x6a, 0x59, 0xad, 0xce, 0xcd, 0xc8, 0xea, 0x09, 0x2e,
	0x06, 0xb9, 0xc1, 0x18, 0xc8, 0x47, 0x14, 0xb7, 0x30, 0xa3, 0xc4, 0x21, 0x71, 0x56, 0x66, 0x1a,
	0x3c, 0x9c, 0xdf, 0x61, 0xcb, 0xb8, 0x6d, 0x4e, 0x97, 0xd5, 0xaa, 0x36, 0x23, 0xab, 0x6b, 0xb8,
	0x7a, 0xc5, 0x3e, 0x3a, 0xf2, 0xf9, 0xa2, 0xba, 0x83, 0xf9, 0x38, 0xc9, 0xc6, 0x0d, 0x9c, 0x1d,
	0xbe, 0x47, 0xd9, 0xb9, 0xe0, 0xc3, 0x67, 0x05, 0x97, 0xeb, 0x63, 0x81, 0x21, 0x89, 0xc7, 0x5e,
	0xac, 0xa1, 0x2e, 0x1d, 0xc4, 0x42, 0xff, 0x94, 0xd3, 0x93, 0x66, 0xb2, 0x99, 0x2e, 0xcb, 0x39,
	0x75, 0x22, 0x36, 0x00, 0x1b, 0xe4, 0x1c, 0x43, 0xc8, 0xdf, 0xca, 0xbf, 0x69, 0xe5, 0xe2, 0xdf,
	0x5d, 0x31, 0x7a, 0x9b, 0x1d, 0x9f, 0xfc, 0xf8, 0x15, 0x00, 0x00, 0xff, 0xff, 0x46, 0x79, 0x9a,
	0x61, 0x7c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ComodoScannerClient is the client API for ComodoScanner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComodoScannerClient interface {
	// Scan a file
	ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error)
	// Get program version
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type comodoScannerClient struct {
	cc *grpc.ClientConn
}

func NewComodoScannerClient(cc *grpc.ClientConn) ComodoScannerClient {
	return &comodoScannerClient{cc}
}

func (c *comodoScannerClient) ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error) {
	out := new(ScanResponse)
	err := c.cc.Invoke(ctx, "/comodo.api.ComodoScanner/ScanFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comodoScannerClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/comodo.api.ComodoScanner/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComodoScannerServer is the server API for ComodoScanner service.
type ComodoScannerServer interface {
	// Scan a file
	ScanFile(context.Context, *ScanFileRequest) (*ScanResponse, error)
	// Get program version
	GetVersion(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedComodoScannerServer can be embedded to have forward compatible implementations.
type UnimplementedComodoScannerServer struct {
}

func (*UnimplementedComodoScannerServer) ScanFile(ctx context.Context, req *ScanFileRequest) (*ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanFile not implemented")
}
func (*UnimplementedComodoScannerServer) GetVersion(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterComodoScannerServer(s *grpc.Server, srv ComodoScannerServer) {
	s.RegisterService(&_ComodoScanner_serviceDesc, srv)
}

func _ComodoScanner_ScanFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComodoScannerServer).ScanFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comodo.api.ComodoScanner/ScanFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComodoScannerServer).ScanFile(ctx, req.(*ScanFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComodoScanner_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComodoScannerServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comodo.api.ComodoScanner/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComodoScannerServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComodoScanner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comodo.api.ComodoScanner",
	HandlerType: (*ComodoScannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanFile",
			Handler:    _ComodoScanner_ScanFile_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ComodoScanner_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multiav.comodo.proto",
}
