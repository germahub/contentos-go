// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_post.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoPost struct {
	PostId               uint64                            `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Category             string                            `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Author               *prototype.AccountName            `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Title                string                            `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string                            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Tags                 []string                          `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Created              *prototype.TimePointSec           `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	LastPayout           *prototype.TimePointSec           `protobuf:"bytes,8,opt,name=last_payout,json=lastPayout,proto3" json:"last_payout,omitempty"`
	Depth                uint32                            `protobuf:"varint,9,opt,name=depth,proto3" json:"depth,omitempty"`
	Children             uint32                            `protobuf:"varint,10,opt,name=children,proto3" json:"children,omitempty"`
	RootId               uint64                            `protobuf:"varint,11,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	ParentId             uint64                            `protobuf:"varint,12,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	VoteCnt              uint64                            `protobuf:"varint,13,opt,name=vote_cnt,json=voteCnt,proto3" json:"vote_cnt,omitempty"`
	Beneficiaries        []*prototype.BeneficiaryRouteType `protobuf:"bytes,14,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SoPost) Reset()         { *m = SoPost{} }
func (m *SoPost) String() string { return proto.CompactTextString(m) }
func (*SoPost) ProtoMessage()    {}
func (*SoPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{0}
}

func (m *SoPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoPost.Unmarshal(m, b)
}
func (m *SoPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoPost.Marshal(b, m, deterministic)
}
func (m *SoPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoPost.Merge(m, src)
}
func (m *SoPost) XXX_Size() int {
	return xxx_messageInfo_SoPost.Size(m)
}
func (m *SoPost) XXX_DiscardUnknown() {
	xxx_messageInfo_SoPost.DiscardUnknown(m)
}

var xxx_messageInfo_SoPost proto.InternalMessageInfo

func (m *SoPost) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *SoPost) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *SoPost) GetAuthor() *prototype.AccountName {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *SoPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SoPost) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SoPost) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SoPost) GetCreated() *prototype.TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *SoPost) GetLastPayout() *prototype.TimePointSec {
	if m != nil {
		return m.LastPayout
	}
	return nil
}

func (m *SoPost) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *SoPost) GetChildren() uint32 {
	if m != nil {
		return m.Children
	}
	return 0
}

func (m *SoPost) GetRootId() uint64 {
	if m != nil {
		return m.RootId
	}
	return 0
}

func (m *SoPost) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *SoPost) GetVoteCnt() uint64 {
	if m != nil {
		return m.VoteCnt
	}
	return 0
}

func (m *SoPost) GetBeneficiaries() []*prototype.BeneficiaryRouteType {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

type SoListPostByCreated struct {
	Created              *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	PostId               uint64                  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListPostByCreated) Reset()         { *m = SoListPostByCreated{} }
func (m *SoListPostByCreated) String() string { return proto.CompactTextString(m) }
func (*SoListPostByCreated) ProtoMessage()    {}
func (*SoListPostByCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{1}
}

func (m *SoListPostByCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListPostByCreated.Unmarshal(m, b)
}
func (m *SoListPostByCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListPostByCreated.Marshal(b, m, deterministic)
}
func (m *SoListPostByCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListPostByCreated.Merge(m, src)
}
func (m *SoListPostByCreated) XXX_Size() int {
	return xxx_messageInfo_SoListPostByCreated.Size(m)
}
func (m *SoListPostByCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListPostByCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SoListPostByCreated proto.InternalMessageInfo

func (m *SoListPostByCreated) GetCreated() *prototype.TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *SoListPostByCreated) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoUniquePostByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniquePostByPostId) Reset()         { *m = SoUniquePostByPostId{} }
func (m *SoUniquePostByPostId) String() string { return proto.CompactTextString(m) }
func (*SoUniquePostByPostId) ProtoMessage()    {}
func (*SoUniquePostByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{2}
}

func (m *SoUniquePostByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniquePostByPostId.Unmarshal(m, b)
}
func (m *SoUniquePostByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniquePostByPostId.Marshal(b, m, deterministic)
}
func (m *SoUniquePostByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniquePostByPostId.Merge(m, src)
}
func (m *SoUniquePostByPostId) XXX_Size() int {
	return xxx_messageInfo_SoUniquePostByPostId.Size(m)
}
func (m *SoUniquePostByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniquePostByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniquePostByPostId proto.InternalMessageInfo

func (m *SoUniquePostByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*SoPost)(nil), "table.so_post")
	proto.RegisterType((*SoListPostByCreated)(nil), "table.so_list_post_by_created")
	proto.RegisterType((*SoUniquePostByPostId)(nil), "table.so_unique_post_by_post_id")
}

func init() { proto.RegisterFile("app/table/so_post.proto", fileDescriptor_aaea493edfbbad11) }

var fileDescriptor_aaea493edfbbad11 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbf, 0x8f, 0xd3, 0x30,
	0x14, 0x56, 0xae, 0xbf, 0x5d, 0xca, 0x60, 0x9d, 0x54, 0xf7, 0x58, 0x42, 0x07, 0x94, 0x85, 0x46,
	0xba, 0x63, 0x62, 0x84, 0x01, 0x75, 0x43, 0x19, 0x59, 0x2c, 0xc7, 0x79, 0x24, 0x96, 0x52, 0xbf,
	0x60, 0xbf, 0x20, 0xe5, 0x5f, 0xe6, 0xaf, 0x40, 0x76, 0x7a, 0x25, 0x0c, 0x88, 0x5b, 0x12, 0x7f,
	0x3f, 0x9c, 0xe7, 0x7c, 0xfe, 0xd8, 0x5e, 0x75, 0x5d, 0x4e, 0xaa, 0x6c, 0x21, 0xf7, 0x28, 0x3b,
	0xf4, 0x74, 0xea, 0x1c, 0x12, 0xf2, 0x45, 0x24, 0x1f, 0xee, 0x23, 0xa2, 0xa1, 0x83, 0x3c, 0x3c,
	0x46, 0xf1, 0xf8, 0x6b, 0xc6, 0x56, 0x57, 0x3b, 0xdf, 0xb3, 0x55, 0x78, 0x4b, 0x53, 0x89, 0x24,
	0x4d, 0xb2, 0x79, 0xb1, 0x0c, 0xf0, 0x5c, 0xf1, 0x07, 0xb6, 0xd6, 0x8a, 0xa0, 0x46, 0x37, 0x88,
	0xbb, 0x34, 0xc9, 0x36, 0xc5, 0x0d, 0xf3, 0x9c, 0x2d, 0x55, 0x4f, 0x0d, 0x3a, 0x31, 0x4b, 0x93,
	0x6c, 0xfb, 0xb8, 0x3f, 0xdd, 0xe6, 0x9c, 0x94, 0xd6, 0xd8, 0x5b, 0x92, 0x56, 0x5d, 0xa0, 0xb8,
	0xda, 0xf8, 0x3d, 0x5b, 0x90, 0xa1, 0x16, 0xc4, 0x3c, 0x7e, 0x69, 0x04, 0x9c, 0xb3, 0x79, 0x89,
	0xd5, 0x20, 0x16, 0x91, 0x8c, 0xeb, 0xc0, 0x91, 0xaa, 0xbd, 0x58, 0xa6, 0xb3, 0xc0, 0x85, 0x35,
	0x7f, 0x62, 0x2b, 0xed, 0x40, 0x11, 0x54, 0x62, 0x15, 0xe7, 0x1d, 0x26, 0xf3, 0xc8, 0x5c, 0x40,
	0x76, 0x68, 0x2c, 0x49, 0x0f, 0xba, 0x78, 0x76, 0xf2, 0x8f, 0x6c, 0xdb, 0x2a, 0x4f, 0xb2, 0x53,
	0x03, 0xf6, 0x24, 0xd6, 0xff, 0xdb, 0xc8, 0x82, 0xfb, 0x6b, 0x34, 0x87, 0xe3, 0x56, 0xd0, 0x51,
	0x23, 0x36, 0x69, 0x92, 0xed, 0x8a, 0x11, 0xc4, 0x44, 0x1a, 0xd3, 0x56, 0x0e, 0xac, 0x60, 0x51,
	0xb8, 0xe1, 0x10, 0xa3, 0x43, 0x8c, 0x31, 0x6e, 0xc7, 0x18, 0x03, 0x3c, 0x57, 0xfc, 0x0d, 0xdb,
	0x74, 0xca, 0x81, 0x8d, 0xd2, 0xab, 0x28, 0xad, 0x47, 0xe2, 0x5c, 0xf1, 0x03, 0x5b, 0xff, 0x44,
	0x02, 0xa9, 0x2d, 0x89, 0x5d, 0xd4, 0x56, 0x01, 0x7f, 0xb6, 0xc4, 0xbf, 0xb0, 0x5d, 0x09, 0x16,
	0xbe, 0x1b, 0x6d, 0x94, 0x33, 0xe0, 0xc5, 0xeb, 0x74, 0x96, 0x6d, 0x1f, 0xdf, 0x4e, 0x7e, 0xe0,
	0x8f, 0x3e, 0x48, 0x87, 0x3d, 0x81, 0x0c, 0x74, 0xf1, 0xf7, 0xbe, 0x63, 0xcd, 0xf6, 0x1e, 0x65,
	0x6b, 0x42, 0x14, 0xe1, 0xa2, 0xcb, 0x41, 0x3e, 0x47, 0x34, 0xc9, 0x35, 0x79, 0x71, 0xae, 0x93,
	0xc2, 0xdc, 0x4d, 0x0b, 0x73, 0xfc, 0xc0, 0x0e, 0x1e, 0x65, 0x6f, 0xcd, 0x8f, 0x1e, 0x6e, 0xa3,
	0xae, 0xd6, 0x7f, 0xd6, 0xec, 0x53, 0xf6, 0xed, 0x5d, 0x6d, 0xa8, 0xe9, 0xcb, 0x93, 0xc6, 0x4b,
	0xae, 0xd1, 0xeb, 0x46, 0x19, 0x9b, 0x6b, 0xb4, 0x04, 0x96, 0xd0, 0xbf, 0xaf, 0x71, 0x2c, 0x78,
	0xb9, 0x8c, 0x67, 0x7b, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x39, 0xd4, 0x0b, 0xf4, 0x02,
	0x00, 0x00,
}
