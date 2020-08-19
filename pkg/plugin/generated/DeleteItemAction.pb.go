// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DeleteItemAction.proto

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeleteItemActionExecuteRequest struct {
	Plugin string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	Item   []byte `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Backup []byte `protobuf:"bytes,3,opt,name=backup,proto3" json:"backup,omitempty"`
}

func (m *DeleteItemActionExecuteRequest) Reset()                    { *m = DeleteItemActionExecuteRequest{} }
func (m *DeleteItemActionExecuteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteItemActionExecuteRequest) ProtoMessage()               {}
func (*DeleteItemActionExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DeleteItemActionExecuteRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *DeleteItemActionExecuteRequest) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *DeleteItemActionExecuteRequest) GetBackup() []byte {
	if m != nil {
		return m.Backup
	}
	return nil
}

type DeleteItemActionAppliesToRequest struct {
	Plugin string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
}

func (m *DeleteItemActionAppliesToRequest) Reset()         { *m = DeleteItemActionAppliesToRequest{} }
func (m *DeleteItemActionAppliesToRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteItemActionAppliesToRequest) ProtoMessage()    {}
func (*DeleteItemActionAppliesToRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1}
}

func (m *DeleteItemActionAppliesToRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

type DeleteItemActionAppliesToResponse struct {
	ResourceSelector *ResourceSelector `protobuf:"bytes,1,opt,name=ResourceSelector" json:"ResourceSelector,omitempty"`
}

func (m *DeleteItemActionAppliesToResponse) Reset()         { *m = DeleteItemActionAppliesToResponse{} }
func (m *DeleteItemActionAppliesToResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteItemActionAppliesToResponse) ProtoMessage()    {}
func (*DeleteItemActionAppliesToResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{2}
}

func (m *DeleteItemActionAppliesToResponse) GetResourceSelector() *ResourceSelector {
	if m != nil {
		return m.ResourceSelector
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteItemActionExecuteRequest)(nil), "generated.DeleteItemActionExecuteRequest")
	proto.RegisterType((*DeleteItemActionAppliesToRequest)(nil), "generated.DeleteItemActionAppliesToRequest")
	proto.RegisterType((*DeleteItemActionAppliesToResponse)(nil), "generated.DeleteItemActionAppliesToResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeleteItemAction service

type DeleteItemActionClient interface {
	AppliesTo(ctx context.Context, in *DeleteItemActionAppliesToRequest, opts ...grpc.CallOption) (*DeleteItemActionAppliesToResponse, error)
	Execute(ctx context.Context, in *DeleteItemActionExecuteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type deleteItemActionClient struct {
	cc *grpc.ClientConn
}

func NewDeleteItemActionClient(cc *grpc.ClientConn) DeleteItemActionClient {
	return &deleteItemActionClient{cc}
}

func (c *deleteItemActionClient) AppliesTo(ctx context.Context, in *DeleteItemActionAppliesToRequest, opts ...grpc.CallOption) (*DeleteItemActionAppliesToResponse, error) {
	out := new(DeleteItemActionAppliesToResponse)
	err := grpc.Invoke(ctx, "/generated.DeleteItemAction/AppliesTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deleteItemActionClient) Execute(ctx context.Context, in *DeleteItemActionExecuteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/generated.DeleteItemAction/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeleteItemAction service

type DeleteItemActionServer interface {
	AppliesTo(context.Context, *DeleteItemActionAppliesToRequest) (*DeleteItemActionAppliesToResponse, error)
	Execute(context.Context, *DeleteItemActionExecuteRequest) (*Empty, error)
}

func RegisterDeleteItemActionServer(s *grpc.Server, srv DeleteItemActionServer) {
	s.RegisterService(&_DeleteItemAction_serviceDesc, srv)
}

func _DeleteItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemActionAppliesToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.DeleteItemAction/AppliesTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteItemActionServer).AppliesTo(ctx, req.(*DeleteItemActionAppliesToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeleteItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemActionExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.DeleteItemAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteItemActionServer).Execute(ctx, req.(*DeleteItemActionExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeleteItemAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.DeleteItemAction",
	HandlerType: (*DeleteItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _DeleteItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _DeleteItemAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "DeleteItemAction.proto",
}

func init() { proto.RegisterFile("DeleteItemAction.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0x89, 0x4a, 0x25, 0xcf, 0x1e, 0xc2, 0x1e, 0x4a, 0x88, 0x20, 0x31, 0xa7, 0x8a, 0x92,
	0x43, 0xbd, 0x79, 0x2b, 0x58, 0xc5, 0x6b, 0xea, 0x1f, 0x48, 0x37, 0x63, 0x0d, 0x6e, 0xb2, 0xeb,
	0xee, 0x5b, 0xd0, 0xbf, 0xe7, 0x2f, 0x13, 0x63, 0x28, 0x35, 0x42, 0xdb, 0xdb, 0xbe, 0xdd, 0x99,
	0xf7, 0x31, 0x3b, 0x34, 0xb9, 0x87, 0x02, 0xe3, 0x89, 0xd1, 0xcc, 0x25, 0xd7, 0xba, 0xcd, 0x8d,
	0xd5, 0xac, 0x45, 0xb8, 0x46, 0x0b, 0x5b, 0x32, 0xaa, 0x64, 0xbc, 0x7c, 0x2d, 0x2d, 0xaa, 0xdf,
	0x87, 0xac, 0xa2, 0x8b, 0xa1, 0x65, 0xf1, 0x01, 0xe9, 0x19, 0x05, 0xde, 0x3d, 0x1c, 0x8b, 0x09,
	0x8d, 0x8c, 0xf2, 0xeb, 0xba, 0x8d, 0x83, 0x34, 0x98, 0x86, 0x45, 0x3f, 0x09, 0x41, 0x27, 0x35,
	0xa3, 0x89, 0x8f, 0xd2, 0x60, 0x3a, 0x2e, 0xba, 0xf3, 0x8f, 0x76, 0x55, 0xca, 0x37, 0x6f, 0xe2,
	0xe3, 0xee, 0xb6, 0x9f, 0xb2, 0x3b, 0x4a, 0x87, 0x94, 0xb9, 0x31, 0xaa, 0x86, 0x7b, 0xd6, 0x7b,
	0x38, 0x99, 0xa2, 0xcb, 0x1d, 0x5e, 0x67, 0x74, 0xeb, 0x20, 0x1e, 0x29, 0x2a, 0xe0, 0xb4, 0xb7,
	0x12, 0x4b, 0x28, 0x48, 0xd6, 0xb6, 0x5b, 0x73, 0x36, 0x3b, 0xcf, 0x37, 0xd1, 0xf3, 0xa1, 0xa4,
	0xf8, 0x67, 0x9a, 0x7d, 0x05, 0x14, 0x0d, 0x71, 0xe2, 0x85, 0xc2, 0x0d, 0x52, 0x5c, 0x6f, 0x2d,
	0xdc, 0x17, 0x2a, 0xb9, 0x39, 0x4c, 0xdc, 0xa7, 0x78, 0xa0, 0xd3, 0xfe, 0xf3, 0xc5, 0xd5, 0x0e,
	0xe3, 0xdf, 0x82, 0x92, 0x68, 0x4b, 0xba, 0x68, 0x0c, 0x7f, 0xae, 0x46, 0x5d, 0xb7, 0xb7, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x31, 0x0b, 0xd3, 0x0e, 0x02, 0x00, 0x00,
}