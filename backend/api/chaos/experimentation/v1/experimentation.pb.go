// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chaos/experimentation/v1/experimentation.proto

package experimentationv1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Experiment struct {
	Id                   uint64   `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	TestConfig           *any.Any `protobuf:"bytes,2,opt,name=test_config,json=testConfig,proto3" json:"test_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Experiment) Reset()         { *m = Experiment{} }
func (m *Experiment) String() string { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()    {}
func (*Experiment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90200f91683161a, []int{0}
}

func (m *Experiment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experiment.Unmarshal(m, b)
}
func (m *Experiment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experiment.Marshal(b, m, deterministic)
}
func (m *Experiment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experiment.Merge(m, src)
}
func (m *Experiment) XXX_Size() int {
	return xxx_messageInfo_Experiment.Size(m)
}
func (m *Experiment) XXX_DiscardUnknown() {
	xxx_messageInfo_Experiment.DiscardUnknown(m)
}

var xxx_messageInfo_Experiment proto.InternalMessageInfo

func (m *Experiment) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Experiment) GetTestConfig() *any.Any {
	if m != nil {
		return m.TestConfig
	}
	return nil
}

type CreateExperimentsRequest struct {
	Experiments          []*Experiment `protobuf:"bytes,1,rep,name=experiments,proto3" json:"experiments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateExperimentsRequest) Reset()         { *m = CreateExperimentsRequest{} }
func (m *CreateExperimentsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateExperimentsRequest) ProtoMessage()    {}
func (*CreateExperimentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90200f91683161a, []int{1}
}

func (m *CreateExperimentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExperimentsRequest.Unmarshal(m, b)
}
func (m *CreateExperimentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExperimentsRequest.Marshal(b, m, deterministic)
}
func (m *CreateExperimentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExperimentsRequest.Merge(m, src)
}
func (m *CreateExperimentsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateExperimentsRequest.Size(m)
}
func (m *CreateExperimentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExperimentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExperimentsRequest proto.InternalMessageInfo

func (m *CreateExperimentsRequest) GetExperiments() []*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

type CreateExperimentsResponse struct {
	Experiments          []*Experiment `protobuf:"bytes,1,rep,name=experiments,proto3" json:"experiments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateExperimentsResponse) Reset()         { *m = CreateExperimentsResponse{} }
func (m *CreateExperimentsResponse) String() string { return proto.CompactTextString(m) }
func (*CreateExperimentsResponse) ProtoMessage()    {}
func (*CreateExperimentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90200f91683161a, []int{2}
}

func (m *CreateExperimentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExperimentsResponse.Unmarshal(m, b)
}
func (m *CreateExperimentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExperimentsResponse.Marshal(b, m, deterministic)
}
func (m *CreateExperimentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExperimentsResponse.Merge(m, src)
}
func (m *CreateExperimentsResponse) XXX_Size() int {
	return xxx_messageInfo_CreateExperimentsResponse.Size(m)
}
func (m *CreateExperimentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExperimentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExperimentsResponse proto.InternalMessageInfo

func (m *CreateExperimentsResponse) GetExperiments() []*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

type GetExperimentsRequest struct {
	Ids                  []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExperimentsRequest) Reset()         { *m = GetExperimentsRequest{} }
func (m *GetExperimentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetExperimentsRequest) ProtoMessage()    {}
func (*GetExperimentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90200f91683161a, []int{3}
}

func (m *GetExperimentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExperimentsRequest.Unmarshal(m, b)
}
func (m *GetExperimentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExperimentsRequest.Marshal(b, m, deterministic)
}
func (m *GetExperimentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExperimentsRequest.Merge(m, src)
}
func (m *GetExperimentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetExperimentsRequest.Size(m)
}
func (m *GetExperimentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExperimentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExperimentsRequest proto.InternalMessageInfo

func (m *GetExperimentsRequest) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetExperimentsResponse struct {
	Experiments          []*Experiment `protobuf:"bytes,3,rep,name=experiments,proto3" json:"experiments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetExperimentsResponse) Reset()         { *m = GetExperimentsResponse{} }
func (m *GetExperimentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetExperimentsResponse) ProtoMessage()    {}
func (*GetExperimentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90200f91683161a, []int{4}
}

func (m *GetExperimentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExperimentsResponse.Unmarshal(m, b)
}
func (m *GetExperimentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExperimentsResponse.Marshal(b, m, deterministic)
}
func (m *GetExperimentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExperimentsResponse.Merge(m, src)
}
func (m *GetExperimentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetExperimentsResponse.Size(m)
}
func (m *GetExperimentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExperimentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetExperimentsResponse proto.InternalMessageInfo

func (m *GetExperimentsResponse) GetExperiments() []*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

type DeleteExperimentsRequest struct {
	Ids                  []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteExperimentsRequest) Reset()         { *m = DeleteExperimentsRequest{} }
func (m *DeleteExperimentsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteExperimentsRequest) ProtoMessage()    {}
func (*DeleteExperimentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90200f91683161a, []int{5}
}

func (m *DeleteExperimentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteExperimentsRequest.Unmarshal(m, b)
}
func (m *DeleteExperimentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteExperimentsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteExperimentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteExperimentsRequest.Merge(m, src)
}
func (m *DeleteExperimentsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteExperimentsRequest.Size(m)
}
func (m *DeleteExperimentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteExperimentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteExperimentsRequest proto.InternalMessageInfo

func (m *DeleteExperimentsRequest) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type DeleteExperimentsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteExperimentsResponse) Reset()         { *m = DeleteExperimentsResponse{} }
func (m *DeleteExperimentsResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteExperimentsResponse) ProtoMessage()    {}
func (*DeleteExperimentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90200f91683161a, []int{6}
}

func (m *DeleteExperimentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteExperimentsResponse.Unmarshal(m, b)
}
func (m *DeleteExperimentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteExperimentsResponse.Marshal(b, m, deterministic)
}
func (m *DeleteExperimentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteExperimentsResponse.Merge(m, src)
}
func (m *DeleteExperimentsResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteExperimentsResponse.Size(m)
}
func (m *DeleteExperimentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteExperimentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteExperimentsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Experiment)(nil), "clutch.chaos.experimentation.v1.Experiment")
	proto.RegisterType((*CreateExperimentsRequest)(nil), "clutch.chaos.experimentation.v1.CreateExperimentsRequest")
	proto.RegisterType((*CreateExperimentsResponse)(nil), "clutch.chaos.experimentation.v1.CreateExperimentsResponse")
	proto.RegisterType((*GetExperimentsRequest)(nil), "clutch.chaos.experimentation.v1.GetExperimentsRequest")
	proto.RegisterType((*GetExperimentsResponse)(nil), "clutch.chaos.experimentation.v1.GetExperimentsResponse")
	proto.RegisterType((*DeleteExperimentsRequest)(nil), "clutch.chaos.experimentation.v1.DeleteExperimentsRequest")
	proto.RegisterType((*DeleteExperimentsResponse)(nil), "clutch.chaos.experimentation.v1.DeleteExperimentsResponse")
}

func init() {
	proto.RegisterFile("chaos/experimentation/v1/experimentation.proto", fileDescriptor_b90200f91683161a)
}

var fileDescriptor_b90200f91683161a = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0xb4, 0x2e, 0xe5, 0x04, 0x8a, 0x3b, 0xab, 0x6b, 0x9a, 0x5d, 0xb0, 0x04, 0xc1,
	0xba, 0xca, 0x84, 0x54, 0x54, 0xec, 0xdd, 0x76, 0x15, 0xf1, 0x42, 0x90, 0x78, 0x21, 0x78, 0x23,
	0xb3, 0xc9, 0x6c, 0x76, 0x24, 0xce, 0xc4, 0xce, 0x34, 0xb8, 0xb7, 0x3e, 0x82, 0xbe, 0x83, 0x82,
	0xb7, 0x3e, 0x8a, 0x4f, 0x20, 0xf8, 0x14, 0x5e, 0x49, 0xfe, 0xac, 0xdd, 0xe6, 0x0f, 0xad, 0xdb,
	0xbb, 0xc9, 0xcc, 0xf9, 0xbe, 0xf3, 0x9b, 0xef, 0x24, 0x01, 0x12, 0x9c, 0x52, 0xa9, 0x5c, 0xf6,
	0x31, 0x61, 0x33, 0xfe, 0x9e, 0x09, 0x4d, 0x35, 0x97, 0xc2, 0x4d, 0xbd, 0xea, 0x16, 0x49, 0x66,
	0x52, 0x4b, 0x7c, 0x33, 0x88, 0xe7, 0x3a, 0x38, 0x2d, 0x64, 0xa4, 0x5a, 0x93, 0x7a, 0xf6, 0x7e,
	0x24, 0x65, 0x14, 0x33, 0x97, 0x26, 0xdc, 0xa5, 0x42, 0xc8, 0xe2, 0x44, 0x15, 0x72, 0xfb, 0x46,
	0x4a, 0x63, 0x1e, 0x52, 0xcd, 0xdc, 0xf3, 0x45, 0x79, 0x30, 0x28, 0x65, 0xf9, 0xd3, 0xf1, 0xfc,
	0xc4, 0xa5, 0xe2, 0xac, 0x3c, 0xb2, 0x32, 0xab, 0xd4, 0xab, 0xbb, 0x39, 0xaf, 0x00, 0x9e, 0xfe,
	0x23, 0xc0, 0x7d, 0x30, 0x78, 0x68, 0xa1, 0x21, 0x1a, 0x6d, 0xf9, 0x06, 0x0f, 0xf1, 0x03, 0x30,
	0x35, 0x53, 0xfa, 0x6d, 0x20, 0xc5, 0x09, 0x8f, 0x2c, 0x63, 0x88, 0x46, 0xe6, 0xf8, 0x1a, 0x29,
	0x1a, 0x91, 0xf3, 0x46, 0xe4, 0x50, 0x9c, 0xf9, 0x90, 0x15, 0x1e, 0xe5, 0x75, 0x8e, 0x02, 0xeb,
	0x68, 0xc6, 0xa8, 0x66, 0x0b, 0x6b, 0xe5, 0xb3, 0x0f, 0x73, 0xa6, 0x34, 0x7e, 0x0d, 0xe6, 0xe2,
	0xca, 0xca, 0x42, 0xc3, 0xce, 0xc8, 0x1c, 0xdf, 0x25, 0x2b, 0x32, 0x21, 0x0b, 0xa7, 0x69, 0xef,
	0xcf, 0xf4, 0xca, 0x67, 0x64, 0xf4, 0x90, 0x7f, 0xd1, 0xc9, 0x79, 0x07, 0x83, 0x86, 0xa6, 0x2a,
	0x91, 0x42, 0x31, 0xfc, 0x62, 0xd3, 0xae, 0xcb, 0xbd, 0xee, 0xc0, 0xf5, 0x67, 0x4c, 0x37, 0xdc,
	0xee, 0x2a, 0x74, 0x78, 0x58, 0xf8, 0x77, 0xfd, 0x6c, 0xe9, 0x44, 0xb0, 0x5b, 0x2d, 0x6d, 0x66,
	0xea, 0x6c, 0xc8, 0x74, 0x0f, 0xac, 0x27, 0x2c, 0x66, 0x8d, 0xa1, 0xd7, 0xb1, 0xf6, 0x60, 0xd0,
	0x50, 0x5d, 0x90, 0x8d, 0xbf, 0x76, 0xa1, 0x7f, 0x61, 0xff, 0xf0, 0xe5, 0x73, 0xfc, 0x03, 0xc1,
	0x76, 0x2d, 0x5e, 0xfc, 0x78, 0x25, 0x6d, 0xdb, 0x7b, 0x60, 0x4f, 0x2e, 0x23, 0x2d, 0xf8, 0x9c,
	0xdb, 0xdf, 0x7f, 0xed, 0x1b, 0x3d, 0xf4, 0xe9, 0xe7, 0xef, 0x2f, 0xc6, 0x9e, 0xb3, 0xbb, 0xfc,
	0xa9, 0x29, 0x37, 0xc8, 0x85, 0x13, 0x74, 0x80, 0xbf, 0x21, 0xe8, 0x2f, 0xa7, 0x8f, 0x1f, 0xae,
	0xec, 0xdb, 0x38, 0x59, 0xfb, 0xd1, 0x7f, 0xeb, 0x4a, 0xd8, 0x5b, 0x39, 0xac, 0x91, 0xc3, 0x5a,
	0xce, 0x4e, 0x15, 0x36, 0x62, 0x3a, 0x23, 0xcd, 0xf2, 0xad, 0x0d, 0x64, 0x8d, 0x7c, 0xdb, 0x46,
	0xbe, 0x46, 0xbe, 0xad, 0xf3, 0x2f, 0xf3, 0xed, 0xb6, 0xe5, 0x1b, 0xe6, 0xc2, 0x09, 0x3a, 0x98,
	0xee, 0xbc, 0xd9, 0xae, 0x18, 0xa7, 0xde, 0xf1, 0x56, 0xfe, 0x5f, 0xb8, 0xff, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xdf, 0x30, 0x3a, 0x6e, 0x18, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExperimentsAPIClient is the client API for ExperimentsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExperimentsAPIClient interface {
	CreateExperiments(ctx context.Context, in *CreateExperimentsRequest, opts ...grpc.CallOption) (*CreateExperimentsResponse, error)
	GetExperiments(ctx context.Context, in *GetExperimentsRequest, opts ...grpc.CallOption) (*GetExperimentsResponse, error)
	DeleteExperiments(ctx context.Context, in *DeleteExperimentsRequest, opts ...grpc.CallOption) (*DeleteExperimentsResponse, error)
}

type experimentsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentsAPIClient(cc grpc.ClientConnInterface) ExperimentsAPIClient {
	return &experimentsAPIClient{cc}
}

func (c *experimentsAPIClient) CreateExperiments(ctx context.Context, in *CreateExperimentsRequest, opts ...grpc.CallOption) (*CreateExperimentsResponse, error) {
	out := new(CreateExperimentsResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/CreateExperiments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentsAPIClient) GetExperiments(ctx context.Context, in *GetExperimentsRequest, opts ...grpc.CallOption) (*GetExperimentsResponse, error) {
	out := new(GetExperimentsResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetExperiments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentsAPIClient) DeleteExperiments(ctx context.Context, in *DeleteExperimentsRequest, opts ...grpc.CallOption) (*DeleteExperimentsResponse, error) {
	out := new(DeleteExperimentsResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/DeleteExperiments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentsAPIServer is the server API for ExperimentsAPI service.
type ExperimentsAPIServer interface {
	CreateExperiments(context.Context, *CreateExperimentsRequest) (*CreateExperimentsResponse, error)
	GetExperiments(context.Context, *GetExperimentsRequest) (*GetExperimentsResponse, error)
	DeleteExperiments(context.Context, *DeleteExperimentsRequest) (*DeleteExperimentsResponse, error)
}

// UnimplementedExperimentsAPIServer can be embedded to have forward compatible implementations.
type UnimplementedExperimentsAPIServer struct {
}

func (*UnimplementedExperimentsAPIServer) CreateExperiments(ctx context.Context, req *CreateExperimentsRequest) (*CreateExperimentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExperiments not implemented")
}
func (*UnimplementedExperimentsAPIServer) GetExperiments(ctx context.Context, req *GetExperimentsRequest) (*GetExperimentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperiments not implemented")
}
func (*UnimplementedExperimentsAPIServer) DeleteExperiments(ctx context.Context, req *DeleteExperimentsRequest) (*DeleteExperimentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExperiments not implemented")
}

func RegisterExperimentsAPIServer(s *grpc.Server, srv ExperimentsAPIServer) {
	s.RegisterService(&_ExperimentsAPI_serviceDesc, srv)
}

func _ExperimentsAPI_CreateExperiments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExperimentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).CreateExperiments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/CreateExperiments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).CreateExperiments(ctx, req.(*CreateExperimentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentsAPI_GetExperiments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperimentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).GetExperiments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetExperiments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).GetExperiments(ctx, req.(*GetExperimentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentsAPI_DeleteExperiments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExperimentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).DeleteExperiments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/DeleteExperiments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).DeleteExperiments(ctx, req.(*DeleteExperimentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExperimentsAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.chaos.experimentation.v1.ExperimentsAPI",
	HandlerType: (*ExperimentsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExperiments",
			Handler:    _ExperimentsAPI_CreateExperiments_Handler,
		},
		{
			MethodName: "GetExperiments",
			Handler:    _ExperimentsAPI_GetExperiments_Handler,
		},
		{
			MethodName: "DeleteExperiments",
			Handler:    _ExperimentsAPI_DeleteExperiments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chaos/experimentation/v1/experimentation.proto",
}
