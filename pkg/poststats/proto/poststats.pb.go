// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/poststats/proto/poststats.proto

package poststats

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPostStatsRequest struct {
	PostUid              string   `protobuf:"bytes,1,opt,name=postUid,proto3" json:"postUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostStatsRequest) Reset()         { *m = GetPostStatsRequest{} }
func (m *GetPostStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPostStatsRequest) ProtoMessage()    {}
func (*GetPostStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{0}
}
func (m *GetPostStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostStatsRequest.Unmarshal(m, b)
}
func (m *GetPostStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostStatsRequest.Marshal(b, m, deterministic)
}
func (dst *GetPostStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostStatsRequest.Merge(dst, src)
}
func (m *GetPostStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPostStatsRequest.Size(m)
}
func (m *GetPostStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostStatsRequest proto.InternalMessageInfo

func (m *GetPostStatsRequest) GetPostUid() string {
	if m != nil {
		return m.PostUid
	}
	return ""
}

type SinglePostStats struct {
	PostUid              string   `protobuf:"bytes,1,opt,name=postUid,proto3" json:"postUid,omitempty"`
	NumLikes             int32    `protobuf:"varint,2,opt,name=numLikes,proto3" json:"numLikes,omitempty"`
	NumDislikes          int32    `protobuf:"varint,3,opt,name=numDislikes,proto3" json:"numDislikes,omitempty"`
	NumViews             int32    `protobuf:"varint,4,opt,name=numViews,proto3" json:"numViews,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SinglePostStats) Reset()         { *m = SinglePostStats{} }
func (m *SinglePostStats) String() string { return proto.CompactTextString(m) }
func (*SinglePostStats) ProtoMessage()    {}
func (*SinglePostStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{1}
}
func (m *SinglePostStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SinglePostStats.Unmarshal(m, b)
}
func (m *SinglePostStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SinglePostStats.Marshal(b, m, deterministic)
}
func (dst *SinglePostStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SinglePostStats.Merge(dst, src)
}
func (m *SinglePostStats) XXX_Size() int {
	return xxx_messageInfo_SinglePostStats.Size(m)
}
func (m *SinglePostStats) XXX_DiscardUnknown() {
	xxx_messageInfo_SinglePostStats.DiscardUnknown(m)
}

var xxx_messageInfo_SinglePostStats proto.InternalMessageInfo

func (m *SinglePostStats) GetPostUid() string {
	if m != nil {
		return m.PostUid
	}
	return ""
}

func (m *SinglePostStats) GetNumLikes() int32 {
	if m != nil {
		return m.NumLikes
	}
	return 0
}

func (m *SinglePostStats) GetNumDislikes() int32 {
	if m != nil {
		return m.NumDislikes
	}
	return 0
}

func (m *SinglePostStats) GetNumViews() int32 {
	if m != nil {
		return m.NumViews
	}
	return 0
}

type CreatePostStatsRequest struct {
	PostUid              string   `protobuf:"bytes,1,opt,name=postUid,proto3" json:"postUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostStatsRequest) Reset()         { *m = CreatePostStatsRequest{} }
func (m *CreatePostStatsRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostStatsRequest) ProtoMessage()    {}
func (*CreatePostStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{2}
}
func (m *CreatePostStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostStatsRequest.Unmarshal(m, b)
}
func (m *CreatePostStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostStatsRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePostStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostStatsRequest.Merge(dst, src)
}
func (m *CreatePostStatsRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostStatsRequest.Size(m)
}
func (m *CreatePostStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostStatsRequest proto.InternalMessageInfo

func (m *CreatePostStatsRequest) GetPostUid() string {
	if m != nil {
		return m.PostUid
	}
	return ""
}

type LikePostRequest struct {
	PostUid              string   `protobuf:"bytes,1,opt,name=postUid,proto3" json:"postUid,omitempty"`
	UserUid              string   `protobuf:"bytes,2,opt,name=userUid,proto3" json:"userUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LikePostRequest) Reset()         { *m = LikePostRequest{} }
func (m *LikePostRequest) String() string { return proto.CompactTextString(m) }
func (*LikePostRequest) ProtoMessage()    {}
func (*LikePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{3}
}
func (m *LikePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LikePostRequest.Unmarshal(m, b)
}
func (m *LikePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LikePostRequest.Marshal(b, m, deterministic)
}
func (dst *LikePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LikePostRequest.Merge(dst, src)
}
func (m *LikePostRequest) XXX_Size() int {
	return xxx_messageInfo_LikePostRequest.Size(m)
}
func (m *LikePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LikePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LikePostRequest proto.InternalMessageInfo

func (m *LikePostRequest) GetPostUid() string {
	if m != nil {
		return m.PostUid
	}
	return ""
}

func (m *LikePostRequest) GetUserUid() string {
	if m != nil {
		return m.UserUid
	}
	return ""
}

type LikePostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	FirstTime            bool     `protobuf:"varint,2,opt,name=firstTime,proto3" json:"firstTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LikePostResponse) Reset()         { *m = LikePostResponse{} }
func (m *LikePostResponse) String() string { return proto.CompactTextString(m) }
func (*LikePostResponse) ProtoMessage()    {}
func (*LikePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{4}
}
func (m *LikePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LikePostResponse.Unmarshal(m, b)
}
func (m *LikePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LikePostResponse.Marshal(b, m, deterministic)
}
func (dst *LikePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LikePostResponse.Merge(dst, src)
}
func (m *LikePostResponse) XXX_Size() int {
	return xxx_messageInfo_LikePostResponse.Size(m)
}
func (m *LikePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LikePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LikePostResponse proto.InternalMessageInfo

func (m *LikePostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LikePostResponse) GetFirstTime() bool {
	if m != nil {
		return m.FirstTime
	}
	return false
}

type DislikePostRequest struct {
	PostUid              string   `protobuf:"bytes,1,opt,name=postUid,proto3" json:"postUid,omitempty"`
	UserUid              string   `protobuf:"bytes,2,opt,name=userUid,proto3" json:"userUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DislikePostRequest) Reset()         { *m = DislikePostRequest{} }
func (m *DislikePostRequest) String() string { return proto.CompactTextString(m) }
func (*DislikePostRequest) ProtoMessage()    {}
func (*DislikePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{5}
}
func (m *DislikePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DislikePostRequest.Unmarshal(m, b)
}
func (m *DislikePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DislikePostRequest.Marshal(b, m, deterministic)
}
func (dst *DislikePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DislikePostRequest.Merge(dst, src)
}
func (m *DislikePostRequest) XXX_Size() int {
	return xxx_messageInfo_DislikePostRequest.Size(m)
}
func (m *DislikePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DislikePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DislikePostRequest proto.InternalMessageInfo

func (m *DislikePostRequest) GetPostUid() string {
	if m != nil {
		return m.PostUid
	}
	return ""
}

func (m *DislikePostRequest) GetUserUid() string {
	if m != nil {
		return m.UserUid
	}
	return ""
}

type DislikePostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	FirstTime            bool     `protobuf:"varint,2,opt,name=firstTime,proto3" json:"firstTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DislikePostResponse) Reset()         { *m = DislikePostResponse{} }
func (m *DislikePostResponse) String() string { return proto.CompactTextString(m) }
func (*DislikePostResponse) ProtoMessage()    {}
func (*DislikePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{6}
}
func (m *DislikePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DislikePostResponse.Unmarshal(m, b)
}
func (m *DislikePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DislikePostResponse.Marshal(b, m, deterministic)
}
func (dst *DislikePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DislikePostResponse.Merge(dst, src)
}
func (m *DislikePostResponse) XXX_Size() int {
	return xxx_messageInfo_DislikePostResponse.Size(m)
}
func (m *DislikePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DislikePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DislikePostResponse proto.InternalMessageInfo

func (m *DislikePostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DislikePostResponse) GetFirstTime() bool {
	if m != nil {
		return m.FirstTime
	}
	return false
}

type IncreaseViewsRequest struct {
	PostUid              string   `protobuf:"bytes,1,opt,name=postUid,proto3" json:"postUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncreaseViewsRequest) Reset()         { *m = IncreaseViewsRequest{} }
func (m *IncreaseViewsRequest) String() string { return proto.CompactTextString(m) }
func (*IncreaseViewsRequest) ProtoMessage()    {}
func (*IncreaseViewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{7}
}
func (m *IncreaseViewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncreaseViewsRequest.Unmarshal(m, b)
}
func (m *IncreaseViewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncreaseViewsRequest.Marshal(b, m, deterministic)
}
func (dst *IncreaseViewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncreaseViewsRequest.Merge(dst, src)
}
func (m *IncreaseViewsRequest) XXX_Size() int {
	return xxx_messageInfo_IncreaseViewsRequest.Size(m)
}
func (m *IncreaseViewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncreaseViewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncreaseViewsRequest proto.InternalMessageInfo

func (m *IncreaseViewsRequest) GetPostUid() string {
	if m != nil {
		return m.PostUid
	}
	return ""
}

type IncreaseViewsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncreaseViewsResponse) Reset()         { *m = IncreaseViewsResponse{} }
func (m *IncreaseViewsResponse) String() string { return proto.CompactTextString(m) }
func (*IncreaseViewsResponse) ProtoMessage()    {}
func (*IncreaseViewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{8}
}
func (m *IncreaseViewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncreaseViewsResponse.Unmarshal(m, b)
}
func (m *IncreaseViewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncreaseViewsResponse.Marshal(b, m, deterministic)
}
func (dst *IncreaseViewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncreaseViewsResponse.Merge(dst, src)
}
func (m *IncreaseViewsResponse) XXX_Size() int {
	return xxx_messageInfo_IncreaseViewsResponse.Size(m)
}
func (m *IncreaseViewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncreaseViewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncreaseViewsResponse proto.InternalMessageInfo

type DeletePostStatsRequest struct {
	PostUid              string   `protobuf:"bytes,1,opt,name=postUid,proto3" json:"postUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostStatsRequest) Reset()         { *m = DeletePostStatsRequest{} }
func (m *DeletePostStatsRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostStatsRequest) ProtoMessage()    {}
func (*DeletePostStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{9}
}
func (m *DeletePostStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostStatsRequest.Unmarshal(m, b)
}
func (m *DeletePostStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostStatsRequest.Marshal(b, m, deterministic)
}
func (dst *DeletePostStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostStatsRequest.Merge(dst, src)
}
func (m *DeletePostStatsRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostStatsRequest.Size(m)
}
func (m *DeletePostStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostStatsRequest proto.InternalMessageInfo

func (m *DeletePostStatsRequest) GetPostUid() string {
	if m != nil {
		return m.PostUid
	}
	return ""
}

type DeletePostStatsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostStatsResponse) Reset()         { *m = DeletePostStatsResponse{} }
func (m *DeletePostStatsResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePostStatsResponse) ProtoMessage()    {}
func (*DeletePostStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_poststats_cbd0d844e1ec1f80, []int{10}
}
func (m *DeletePostStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostStatsResponse.Unmarshal(m, b)
}
func (m *DeletePostStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostStatsResponse.Marshal(b, m, deterministic)
}
func (dst *DeletePostStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostStatsResponse.Merge(dst, src)
}
func (m *DeletePostStatsResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePostStatsResponse.Size(m)
}
func (m *DeletePostStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostStatsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetPostStatsRequest)(nil), "poststats.GetPostStatsRequest")
	proto.RegisterType((*SinglePostStats)(nil), "poststats.SinglePostStats")
	proto.RegisterType((*CreatePostStatsRequest)(nil), "poststats.CreatePostStatsRequest")
	proto.RegisterType((*LikePostRequest)(nil), "poststats.LikePostRequest")
	proto.RegisterType((*LikePostResponse)(nil), "poststats.LikePostResponse")
	proto.RegisterType((*DislikePostRequest)(nil), "poststats.DislikePostRequest")
	proto.RegisterType((*DislikePostResponse)(nil), "poststats.DislikePostResponse")
	proto.RegisterType((*IncreaseViewsRequest)(nil), "poststats.IncreaseViewsRequest")
	proto.RegisterType((*IncreaseViewsResponse)(nil), "poststats.IncreaseViewsResponse")
	proto.RegisterType((*DeletePostStatsRequest)(nil), "poststats.DeletePostStatsRequest")
	proto.RegisterType((*DeletePostStatsResponse)(nil), "poststats.DeletePostStatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostStatsClient is the client API for PostStats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostStatsClient interface {
	GetPostStats(ctx context.Context, in *GetPostStatsRequest, opts ...grpc.CallOption) (*SinglePostStats, error)
	CreatePostStats(ctx context.Context, in *CreatePostStatsRequest, opts ...grpc.CallOption) (*SinglePostStats, error)
	LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error)
	DislikePost(ctx context.Context, in *DislikePostRequest, opts ...grpc.CallOption) (*DislikePostResponse, error)
	IncreaseViews(ctx context.Context, in *IncreaseViewsRequest, opts ...grpc.CallOption) (*IncreaseViewsResponse, error)
	DeletePostStats(ctx context.Context, in *DeletePostStatsRequest, opts ...grpc.CallOption) (*DeletePostStatsResponse, error)
}

type postStatsClient struct {
	cc *grpc.ClientConn
}

func NewPostStatsClient(cc *grpc.ClientConn) PostStatsClient {
	return &postStatsClient{cc}
}

func (c *postStatsClient) GetPostStats(ctx context.Context, in *GetPostStatsRequest, opts ...grpc.CallOption) (*SinglePostStats, error) {
	out := new(SinglePostStats)
	err := c.cc.Invoke(ctx, "/poststats.PostStats/GetPostStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStatsClient) CreatePostStats(ctx context.Context, in *CreatePostStatsRequest, opts ...grpc.CallOption) (*SinglePostStats, error) {
	out := new(SinglePostStats)
	err := c.cc.Invoke(ctx, "/poststats.PostStats/CreatePostStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStatsClient) LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error) {
	out := new(LikePostResponse)
	err := c.cc.Invoke(ctx, "/poststats.PostStats/LikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStatsClient) DislikePost(ctx context.Context, in *DislikePostRequest, opts ...grpc.CallOption) (*DislikePostResponse, error) {
	out := new(DislikePostResponse)
	err := c.cc.Invoke(ctx, "/poststats.PostStats/DislikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStatsClient) IncreaseViews(ctx context.Context, in *IncreaseViewsRequest, opts ...grpc.CallOption) (*IncreaseViewsResponse, error) {
	out := new(IncreaseViewsResponse)
	err := c.cc.Invoke(ctx, "/poststats.PostStats/IncreaseViews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStatsClient) DeletePostStats(ctx context.Context, in *DeletePostStatsRequest, opts ...grpc.CallOption) (*DeletePostStatsResponse, error) {
	out := new(DeletePostStatsResponse)
	err := c.cc.Invoke(ctx, "/poststats.PostStats/DeletePostStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostStatsServer is the server API for PostStats service.
type PostStatsServer interface {
	GetPostStats(context.Context, *GetPostStatsRequest) (*SinglePostStats, error)
	CreatePostStats(context.Context, *CreatePostStatsRequest) (*SinglePostStats, error)
	LikePost(context.Context, *LikePostRequest) (*LikePostResponse, error)
	DislikePost(context.Context, *DislikePostRequest) (*DislikePostResponse, error)
	IncreaseViews(context.Context, *IncreaseViewsRequest) (*IncreaseViewsResponse, error)
	DeletePostStats(context.Context, *DeletePostStatsRequest) (*DeletePostStatsResponse, error)
}

func RegisterPostStatsServer(s *grpc.Server, srv PostStatsServer) {
	s.RegisterService(&_PostStats_serviceDesc, srv)
}

func _PostStats_GetPostStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStatsServer).GetPostStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poststats.PostStats/GetPostStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStatsServer).GetPostStats(ctx, req.(*GetPostStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStats_CreatePostStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStatsServer).CreatePostStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poststats.PostStats/CreatePostStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStatsServer).CreatePostStats(ctx, req.(*CreatePostStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStats_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStatsServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poststats.PostStats/LikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStatsServer).LikePost(ctx, req.(*LikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStats_DislikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DislikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStatsServer).DislikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poststats.PostStats/DislikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStatsServer).DislikePost(ctx, req.(*DislikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStats_IncreaseViews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseViewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStatsServer).IncreaseViews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poststats.PostStats/IncreaseViews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStatsServer).IncreaseViews(ctx, req.(*IncreaseViewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStats_DeletePostStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStatsServer).DeletePostStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poststats.PostStats/DeletePostStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStatsServer).DeletePostStats(ctx, req.(*DeletePostStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostStats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poststats.PostStats",
	HandlerType: (*PostStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostStats",
			Handler:    _PostStats_GetPostStats_Handler,
		},
		{
			MethodName: "CreatePostStats",
			Handler:    _PostStats_CreatePostStats_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _PostStats_LikePost_Handler,
		},
		{
			MethodName: "DislikePost",
			Handler:    _PostStats_DislikePost_Handler,
		},
		{
			MethodName: "IncreaseViews",
			Handler:    _PostStats_IncreaseViews_Handler,
		},
		{
			MethodName: "DeletePostStats",
			Handler:    _PostStats_DeletePostStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/poststats/proto/poststats.proto",
}

func init() {
	proto.RegisterFile("pkg/poststats/proto/poststats.proto", fileDescriptor_poststats_cbd0d844e1ec1f80)
}

var fileDescriptor_poststats_cbd0d844e1ec1f80 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x8e, 0x9a, 0x40,
	0x18, 0x0d, 0xda, 0x1f, 0xf8, 0x6c, 0x43, 0x33, 0xb6, 0x95, 0xd2, 0xd6, 0x52, 0x7a, 0xe3, 0x95,
	0x36, 0xf6, 0x11, 0xb4, 0x69, 0x6b, 0x6c, 0x62, 0x70, 0x77, 0xb3, 0xb7, 0xae, 0xfb, 0xad, 0x21,
	0x2a, 0xb0, 0x7c, 0x43, 0xf6, 0x09, 0xf6, 0x65, 0xf7, 0x29, 0x36, 0x33, 0x20, 0x8c, 0x88, 0x1a,
	0xe3, 0xe5, 0x99, 0xf3, 0xe3, 0x38, 0xe7, 0x04, 0xf8, 0x11, 0x2d, 0x17, 0xbd, 0x28, 0x24, 0x4e,
	0x7c, 0xc6, 0xa9, 0x17, 0xc5, 0x21, 0x0f, 0x0b, 0xdc, 0x95, 0x98, 0x19, 0xf9, 0x81, 0xdb, 0x83,
	0xe6, 0x1f, 0xe4, 0x93, 0x90, 0xf8, 0x54, 0x60, 0x0f, 0xef, 0x13, 0x24, 0xce, 0x2c, 0x78, 0x2d,
	0x34, 0x97, 0xfe, 0xad, 0xa5, 0x39, 0x5a, 0xc7, 0xf0, 0x36, 0xd0, 0x7d, 0xd4, 0xc0, 0x9c, 0xfa,
	0xc1, 0x62, 0x85, 0xb9, 0x69, 0xbf, 0x9a, 0xd9, 0xa0, 0x07, 0xc9, 0x7a, 0xec, 0x2f, 0x91, 0xac,
	0x9a, 0xa3, 0x75, 0x5e, 0x7a, 0x39, 0x66, 0x0e, 0x34, 0x82, 0x64, 0x3d, 0xf4, 0x69, 0x25, 0xe9,
	0xba, 0xa4, 0xd5, 0xa3, 0xcc, 0x7d, 0xe5, 0xe3, 0x03, 0x59, 0x2f, 0x72, 0xb7, 0xc4, 0x6e, 0x1f,
	0x3e, 0x0e, 0x62, 0x9c, 0x71, 0x3c, 0xe1, 0xee, 0xbf, 0xc1, 0x14, 0x3f, 0x2d, 0x1c, 0x47, 0xc5,
	0x82, 0x49, 0x08, 0x63, 0xc1, 0xd4, 0x52, 0x26, 0x83, 0xee, 0x08, 0xde, 0x15, 0x31, 0x14, 0x85,
	0x01, 0xa1, 0x50, 0x53, 0x32, 0x9f, 0x23, 0x91, 0xcc, 0xd1, 0xbd, 0x0d, 0x64, 0x5f, 0xc0, 0xb8,
	0xf3, 0x63, 0xe2, 0x17, 0xfe, 0x1a, 0x65, 0x92, 0xee, 0x15, 0x07, 0xee, 0x5f, 0x60, 0xd9, 0xdf,
	0x3d, 0xf7, 0x56, 0xff, 0xa1, 0xb9, 0x95, 0x74, 0xe6, 0xc5, 0x7e, 0xc2, 0xfb, 0x7f, 0xc1, 0x3c,
	0xc6, 0x19, 0xa1, 0x7c, 0xf0, 0xe3, 0xaf, 0xdb, 0x82, 0x0f, 0x25, 0x47, 0x7a, 0x05, 0x51, 0xd5,
	0x10, 0x57, 0x78, 0x52, 0x55, 0x9f, 0xa0, 0xb5, 0xe3, 0x49, 0xe3, 0xfa, 0x4f, 0x75, 0x30, 0x8a,
	0xed, 0x8d, 0xe0, 0x8d, 0x3a, 0x60, 0xd6, 0xee, 0x16, 0x6b, 0xaf, 0x58, 0xb6, 0x6d, 0x2b, 0x7c,
	0x79, 0xc7, 0x13, 0x30, 0x4b, 0x9b, 0x62, 0xdf, 0x15, 0x79, 0xf5, 0xde, 0x0e, 0x26, 0x0e, 0x40,
	0xdf, 0x4c, 0x85, 0xa9, 0xba, 0xd2, 0x0c, 0xed, 0xcf, 0x95, 0x5c, 0x56, 0xe1, 0x18, 0x1a, 0x4a,
	0xb3, 0xec, 0xab, 0xa2, 0xdd, 0xdd, 0x8e, 0xdd, 0xde, 0x47, 0x67, 0x69, 0x1e, 0xbc, 0xdd, 0xaa,
	0x89, 0x7d, 0x53, 0x0c, 0x55, 0x95, 0xdb, 0xce, 0x7e, 0x41, 0x96, 0x79, 0x0d, 0x66, 0xa9, 0xad,
	0xad, 0x87, 0xab, 0x6e, 0xdf, 0x76, 0x0f, 0x49, 0xd2, 0xe4, 0x9b, 0x57, 0xf2, 0x8b, 0xf5, 0xeb,
	0x39, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xa7, 0x91, 0xeb, 0xd8, 0x04, 0x00, 0x00,
}
