// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compiler.proto

package compiler_grpc

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

// Compilation unit.
type CompilationUnit struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompilationUnit) Reset()         { *m = CompilationUnit{} }
func (m *CompilationUnit) String() string { return proto.CompactTextString(m) }
func (*CompilationUnit) ProtoMessage()    {}
func (*CompilationUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5727dbaeb66833, []int{0}
}

func (m *CompilationUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompilationUnit.Unmarshal(m, b)
}
func (m *CompilationUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompilationUnit.Marshal(b, m, deterministic)
}
func (m *CompilationUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompilationUnit.Merge(m, src)
}
func (m *CompilationUnit) XXX_Size() int {
	return xxx_messageInfo_CompilationUnit.Size(m)
}
func (m *CompilationUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_CompilationUnit.DiscardUnknown(m)
}

var xxx_messageInfo_CompilationUnit proto.InternalMessageInfo

func (m *CompilationUnit) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CompilationUnit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Compiler API
type SourceFiles struct {
	Units                []*CompilationUnit `protobuf:"bytes,1,rep,name=units,proto3" json:"units,omitempty"`
	Address              []byte             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SourceFiles) Reset()         { *m = SourceFiles{} }
func (m *SourceFiles) String() string { return proto.CompactTextString(m) }
func (*SourceFiles) ProtoMessage()    {}
func (*SourceFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5727dbaeb66833, []int{1}
}

func (m *SourceFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceFiles.Unmarshal(m, b)
}
func (m *SourceFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceFiles.Marshal(b, m, deterministic)
}
func (m *SourceFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceFiles.Merge(m, src)
}
func (m *SourceFiles) XXX_Size() int {
	return xxx_messageInfo_SourceFiles.Size(m)
}
func (m *SourceFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceFiles.DiscardUnknown(m)
}

var xxx_messageInfo_SourceFiles proto.InternalMessageInfo

func (m *SourceFiles) GetUnits() []*CompilationUnit {
	if m != nil {
		return m.Units
	}
	return nil
}

func (m *SourceFiles) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// Compiled source.
type CompiledUnit struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bytecode             []byte   `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompiledUnit) Reset()         { *m = CompiledUnit{} }
func (m *CompiledUnit) String() string { return proto.CompactTextString(m) }
func (*CompiledUnit) ProtoMessage()    {}
func (*CompiledUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5727dbaeb66833, []int{2}
}

func (m *CompiledUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompiledUnit.Unmarshal(m, b)
}
func (m *CompiledUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompiledUnit.Marshal(b, m, deterministic)
}
func (m *CompiledUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompiledUnit.Merge(m, src)
}
func (m *CompiledUnit) XXX_Size() int {
	return xxx_messageInfo_CompiledUnit.Size(m)
}
func (m *CompiledUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_CompiledUnit.DiscardUnknown(m)
}

var xxx_messageInfo_CompiledUnit proto.InternalMessageInfo

func (m *CompiledUnit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CompiledUnit) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

type CompilationResult struct {
	Units                []*CompiledUnit `protobuf:"bytes,1,rep,name=units,proto3" json:"units,omitempty"`
	Errors               []string        `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CompilationResult) Reset()         { *m = CompilationResult{} }
func (m *CompilationResult) String() string { return proto.CompactTextString(m) }
func (*CompilationResult) ProtoMessage()    {}
func (*CompilationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5727dbaeb66833, []int{3}
}

func (m *CompilationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompilationResult.Unmarshal(m, b)
}
func (m *CompilationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompilationResult.Marshal(b, m, deterministic)
}
func (m *CompilationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompilationResult.Merge(m, src)
}
func (m *CompilationResult) XXX_Size() int {
	return xxx_messageInfo_CompilationResult.Size(m)
}
func (m *CompilationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CompilationResult.DiscardUnknown(m)
}

var xxx_messageInfo_CompilationResult proto.InternalMessageInfo

func (m *CompilationResult) GetUnits() []*CompiledUnit {
	if m != nil {
		return m.Units
	}
	return nil
}

func (m *CompilationResult) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*CompilationUnit)(nil), "compiler_grpc.CompilationUnit")
	proto.RegisterType((*SourceFiles)(nil), "compiler_grpc.SourceFiles")
	proto.RegisterType((*CompiledUnit)(nil), "compiler_grpc.CompiledUnit")
	proto.RegisterType((*CompilationResult)(nil), "compiler_grpc.CompilationResult")
}

func init() {
	proto.RegisterFile("compiler.proto", fileDescriptor_6a5727dbaeb66833)
}

var fileDescriptor_6a5727dbaeb66833 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0xd5, 0xd6, 0x4e, 0xaa, 0xe2, 0x1c, 0x24, 0x44, 0x90, 0xb0, 0xa7, 0x9c, 0x02,
	0x56, 0x2f, 0x5e, 0xbc, 0x28, 0x82, 0xd7, 0x15, 0xc1, 0x8b, 0x4a, 0x9a, 0x0c, 0xb2, 0x90, 0x64,
	0xc3, 0xec, 0x46, 0xf4, 0xdf, 0x4b, 0x76, 0x13, 0x6d, 0x8b, 0xde, 0xe6, 0x3d, 0xf6, 0xcd, 0xfb,
	0x98, 0x85, 0xa3, 0x42, 0xd7, 0xad, 0xaa, 0x88, 0xb3, 0x96, 0xb5, 0xd5, 0x78, 0x38, 0xea, 0xb7,
	0x77, 0x6e, 0x0b, 0x71, 0x0d, 0xc7, 0xb7, 0xce, 0xc8, 0xad, 0xd2, 0xcd, 0x53, 0xa3, 0x2c, 0x22,
	0xec, 0x59, 0xfa, 0xb4, 0x51, 0x90, 0x04, 0xe9, 0x5c, 0xba, 0xb9, 0xf7, 0x9a, 0xbc, 0xa6, 0x68,
	0xd7, 0x7b, 0xfd, 0x2c, 0x5e, 0x20, 0x7c, 0xd4, 0x1d, 0x17, 0x74, 0xaf, 0x2a, 0x32, 0x78, 0x05,
	0xfb, 0x5d, 0xa3, 0xac, 0x89, 0x82, 0x64, 0x92, 0x86, 0xcb, 0xf3, 0x6c, 0xa3, 0x28, 0xdb, 0x6a,
	0x91, 0xfe, 0x31, 0x46, 0x30, 0xcb, 0xcb, 0x92, 0xc9, 0x18, 0xb7, 0x7b, 0x21, 0x47, 0x29, 0x6e,
	0x60, 0xe1, 0x33, 0x54, 0x8e, 0x58, 0x0e, 0x21, 0xf8, 0x45, 0xc0, 0x18, 0x0e, 0x56, 0x5f, 0x96,
	0x0a, 0x5d, 0xd2, 0x10, 0xff, 0xd1, 0xe2, 0x15, 0x4e, 0xd6, 0x3a, 0x25, 0x99, 0xae, 0xb2, 0x78,
	0xb1, 0x09, 0x79, 0xf6, 0x27, 0xa4, 0x2f, 0x1c, 0x09, 0x4f, 0x61, 0x4a, 0xcc, 0x9a, 0x7b, 0xc0,
	0x49, 0x3a, 0x97, 0x83, 0x5a, 0x3e, 0x43, 0x78, 0xf7, 0x51, 0x0f, 0x09, 0xc6, 0x07, 0x98, 0x0d,
	0x33, 0xc6, 0x5b, 0x5b, 0xd7, 0xae, 0x14, 0x27, 0xff, 0x9f, 0xc5, 0x23, 0x8a, 0x9d, 0xd5, 0xd4,
	0xfd, 0xd4, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x0f, 0x0c, 0x0c, 0xbb, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DvmCompilerClient is the client API for DvmCompiler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DvmCompilerClient interface {
	Compile(ctx context.Context, in *SourceFiles, opts ...grpc.CallOption) (*CompilationResult, error)
}

type dvmCompilerClient struct {
	cc grpc.ClientConnInterface
}

func NewDvmCompilerClient(cc grpc.ClientConnInterface) DvmCompilerClient {
	return &dvmCompilerClient{cc}
}

func (c *dvmCompilerClient) Compile(ctx context.Context, in *SourceFiles, opts ...grpc.CallOption) (*CompilationResult, error) {
	out := new(CompilationResult)
	err := c.cc.Invoke(ctx, "/compiler_grpc.DvmCompiler/Compile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DvmCompilerServer is the server API for DvmCompiler service.
type DvmCompilerServer interface {
	Compile(context.Context, *SourceFiles) (*CompilationResult, error)
}

// UnimplementedDvmCompilerServer can be embedded to have forward compatible implementations.
type UnimplementedDvmCompilerServer struct {
}

func (*UnimplementedDvmCompilerServer) Compile(ctx context.Context, req *SourceFiles) (*CompilationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}

func RegisterDvmCompilerServer(s *grpc.Server, srv DvmCompilerServer) {
	s.RegisterService(&_DvmCompiler_serviceDesc, srv)
}

func _DvmCompiler_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SourceFiles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DvmCompilerServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compiler_grpc.DvmCompiler/Compile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DvmCompilerServer).Compile(ctx, req.(*SourceFiles))
	}
	return interceptor(ctx, in, info, handler)
}

var _DvmCompiler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "compiler_grpc.DvmCompiler",
	HandlerType: (*DvmCompilerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _DvmCompiler_Compile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compiler.proto",
}
