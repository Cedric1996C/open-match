// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/api/synchronizer.proto

package ipb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	pb "open-match.dev/open-match/pkg/pb"
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

type SynchronizeRequest struct {
	// A match returned by an mmf.
	Proposal             *pb.Match `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SynchronizeRequest) Reset()         { *m = SynchronizeRequest{} }
func (m *SynchronizeRequest) String() string { return proto.CompactTextString(m) }
func (*SynchronizeRequest) ProtoMessage()    {}
func (*SynchronizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ff6b85fea1c4b7, []int{0}
}

func (m *SynchronizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynchronizeRequest.Unmarshal(m, b)
}
func (m *SynchronizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynchronizeRequest.Marshal(b, m, deterministic)
}
func (m *SynchronizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynchronizeRequest.Merge(m, src)
}
func (m *SynchronizeRequest) XXX_Size() int {
	return xxx_messageInfo_SynchronizeRequest.Size(m)
}
func (m *SynchronizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SynchronizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SynchronizeRequest proto.InternalMessageInfo

func (m *SynchronizeRequest) GetProposal() *pb.Match {
	if m != nil {
		return m.Proposal
	}
	return nil
}

type SynchronizeResponse struct {
	// Instructs the backend call that it can start running the mmfs.
	StartMmfs bool `protobuf:"varint,1,opt,name=start_mmfs,json=startMmfs,proto3" json:"start_mmfs,omitempty"`
	// Instructs the backend call that it should cancel any RPC calls to the mmfs,
	// not send any more matches, and close the send stream.
	CancelMmfs bool `protobuf:"varint,2,opt,name=cancel_mmfs,json=cancelMmfs,proto3" json:"cancel_mmfs,omitempty"`
	// A match ID returned by the evaluator and should be returned to the FetchMatches
	// caller.
	MatchId              string   `protobuf:"bytes,4,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynchronizeResponse) Reset()         { *m = SynchronizeResponse{} }
func (m *SynchronizeResponse) String() string { return proto.CompactTextString(m) }
func (*SynchronizeResponse) ProtoMessage()    {}
func (*SynchronizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ff6b85fea1c4b7, []int{1}
}

func (m *SynchronizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynchronizeResponse.Unmarshal(m, b)
}
func (m *SynchronizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynchronizeResponse.Marshal(b, m, deterministic)
}
func (m *SynchronizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynchronizeResponse.Merge(m, src)
}
func (m *SynchronizeResponse) XXX_Size() int {
	return xxx_messageInfo_SynchronizeResponse.Size(m)
}
func (m *SynchronizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SynchronizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SynchronizeResponse proto.InternalMessageInfo

func (m *SynchronizeResponse) GetStartMmfs() bool {
	if m != nil {
		return m.StartMmfs
	}
	return false
}

func (m *SynchronizeResponse) GetCancelMmfs() bool {
	if m != nil {
		return m.CancelMmfs
	}
	return false
}

func (m *SynchronizeResponse) GetMatchId() string {
	if m != nil {
		return m.MatchId
	}
	return ""
}

func init() {
	proto.RegisterType((*SynchronizeRequest)(nil), "openmatch.internal.SynchronizeRequest")
	proto.RegisterType((*SynchronizeResponse)(nil), "openmatch.internal.SynchronizeResponse")
}

func init() { proto.RegisterFile("internal/api/synchronizer.proto", fileDescriptor_35ff6b85fea1c4b7) }

var fileDescriptor_35ff6b85fea1c4b7 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x16, 0x4d, 0x27, 0x1e, 0xca, 0x7a, 0xa9, 0x05, 0x69, 0xe9, 0xa1, 0xe6, 0xa0,
	0x1b, 0xa9, 0xdf, 0xa0, 0x37, 0x85, 0x5e, 0xe2, 0xcd, 0x4b, 0xd9, 0x24, 0x53, 0xbb, 0x90, 0xfd,
	0xe3, 0xce, 0x5a, 0xd0, 0x4f, 0x2f, 0xd9, 0xc5, 0xa6, 0xd2, 0x83, 0x97, 0x85, 0x37, 0xf3, 0xdb,
	0x37, 0xcc, 0x1b, 0x98, 0x4a, 0xed, 0xd1, 0x69, 0xd1, 0x16, 0xc2, 0xca, 0x82, 0xbe, 0x74, 0xbd,
	0x73, 0x46, 0xcb, 0x6f, 0x74, 0xdc, 0x3a, 0xe3, 0x0d, 0x63, 0xc6, 0xa2, 0x56, 0xc2, 0xd7, 0x3b,
	0xfe, 0x8b, 0x4e, 0x58, 0xc7, 0x2a, 0x24, 0x12, 0xef, 0x48, 0x91, 0x9b, 0xaf, 0x80, 0xbd, 0xf6,
	0xbf, 0x4b, 0xfc, 0xf8, 0x44, 0xf2, 0xec, 0x1e, 0x52, 0xeb, 0x8c, 0x35, 0x24, 0xda, 0x71, 0x32,
	0x4b, 0xf2, 0x6c, 0x39, 0xe2, 0xbd, 0xe1, 0xba, 0x7b, 0xcb, 0x03, 0x31, 0xdf, 0xc3, 0xf5, 0x1f,
	0x0f, 0xb2, 0x46, 0x13, 0xb2, 0x5b, 0x00, 0xf2, 0xc2, 0xf9, 0x8d, 0x52, 0x5b, 0x0a, 0x36, 0x69,
	0x39, 0x0c, 0x95, 0xb5, 0xda, 0x12, 0x9b, 0x42, 0x56, 0x0b, 0x5d, 0x63, 0x1b, 0xfb, 0x67, 0xa1,
	0x0f, 0xb1, 0x14, 0x80, 0x1b, 0x48, 0xc3, 0xbc, 0x8d, 0x6c, 0xc6, 0x83, 0x59, 0x92, 0x0f, 0xcb,
	0xcb, 0xa0, 0x9f, 0x9b, 0x97, 0x41, 0x7a, 0x3e, 0x1a, 0x2c, 0x1d, 0x5c, 0x1d, 0xcd, 0x75, 0xac,
	0x82, 0xec, 0x48, 0xb3, 0x05, 0x3f, 0xcd, 0x80, 0x9f, 0x2e, 0x3b, 0xb9, 0xfb, 0x97, 0x8b, 0x0b,
	0xe5, 0xc9, 0x63, 0xb2, 0xca, 0xdf, 0x16, 0x1d, 0xfd, 0x10, 0xf1, 0x06, 0xf7, 0x45, 0x2f, 0x8b,
	0xc3, 0x51, 0xa4, 0xad, 0xaa, 0x8b, 0x10, 0xf0, 0xd3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84,
	0xba, 0xf6, 0x41, 0xab, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynchronizerClient interface {
	// Synchronize signals the caller when it is safe to run mmfs, collects the
	// mmfs' proposals, and returns the evaluated matches.
	Synchronize(ctx context.Context, opts ...grpc.CallOption) (Synchronizer_SynchronizeClient, error)
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

func (c *synchronizerClient) Synchronize(ctx context.Context, opts ...grpc.CallOption) (Synchronizer_SynchronizeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synchronizer_serviceDesc.Streams[0], "/openmatch.internal.Synchronizer/Synchronize", opts...)
	if err != nil {
		return nil, err
	}
	x := &synchronizerSynchronizeClient{stream}
	return x, nil
}

type Synchronizer_SynchronizeClient interface {
	Send(*SynchronizeRequest) error
	Recv() (*SynchronizeResponse, error)
	grpc.ClientStream
}

type synchronizerSynchronizeClient struct {
	grpc.ClientStream
}

func (x *synchronizerSynchronizeClient) Send(m *SynchronizeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *synchronizerSynchronizeClient) Recv() (*SynchronizeResponse, error) {
	m := new(SynchronizeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SynchronizerServer is the server API for Synchronizer service.
type SynchronizerServer interface {
	// Synchronize signals the caller when it is safe to run mmfs, collects the
	// mmfs' proposals, and returns the evaluated matches.
	Synchronize(Synchronizer_SynchronizeServer) error
}

// UnimplementedSynchronizerServer can be embedded to have forward compatible implementations.
type UnimplementedSynchronizerServer struct {
}

func (*UnimplementedSynchronizerServer) Synchronize(srv Synchronizer_SynchronizeServer) error {
	return status.Errorf(codes.Unimplemented, "method Synchronize not implemented")
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

func _Synchronizer_Synchronize_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SynchronizerServer).Synchronize(&synchronizerSynchronizeServer{stream})
}

type Synchronizer_SynchronizeServer interface {
	Send(*SynchronizeResponse) error
	Recv() (*SynchronizeRequest, error)
	grpc.ServerStream
}

type synchronizerSynchronizeServer struct {
	grpc.ServerStream
}

func (x *synchronizerSynchronizeServer) Send(m *SynchronizeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *synchronizerSynchronizeServer) Recv() (*SynchronizeRequest, error) {
	m := new(SynchronizeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.internal.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Synchronize",
			Handler:       _Synchronizer_Synchronize_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/api/synchronizer.proto",
}
