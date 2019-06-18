// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_witness.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SoWitness struct {
	Owner                *prototype.AccountName   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	CreatedTime          *prototype.TimePointSec  `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Url                  string                   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	VoteCount            *prototype.Vest          `protobuf:"bytes,4,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	SigningKey           *prototype.PublicKeyType `protobuf:"bytes,5,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	ProposedStaminaFree  uint64                   `protobuf:"varint,6,opt,name=proposed_stamina_free,json=proposedStaminaFree,proto3" json:"proposed_stamina_free,omitempty"`
	Active               bool                     `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	TpsExpected          uint64                   `protobuf:"varint,8,opt,name=tps_expected,json=tpsExpected,proto3" json:"tps_expected,omitempty"`
	AccountCreateFee     *prototype.Coin          `protobuf:"bytes,9,opt,name=account_create_fee,json=accountCreateFee,proto3" json:"account_create_fee,omitempty"`
	TopNAcquireFreeToken uint32                   `protobuf:"varint,10,opt,name=top_n_acquire_free_token,json=topNAcquireFreeToken,proto3" json:"top_n_acquire_free_token,omitempty"`
	EpochDuration        uint64                   `protobuf:"varint,11,opt,name=epoch_duration,json=epochDuration,proto3" json:"epoch_duration,omitempty"`
	PerTicketPrice       *prototype.Vest          `protobuf:"bytes,12,opt,name=per_ticket_price,json=perTicketPrice,proto3" json:"per_ticket_price,omitempty"`
	PerTicketWeight      uint64                   `protobuf:"varint,13,opt,name=per_ticket_weight,json=perTicketWeight,proto3" json:"per_ticket_weight,omitempty"`
	VoterList            []*prototype.AccountName `protobuf:"bytes,14,rep,name=voter_list,json=voterList,proto3" json:"voter_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SoWitness) Reset()         { *m = SoWitness{} }
func (m *SoWitness) String() string { return proto.CompactTextString(m) }
func (*SoWitness) ProtoMessage()    {}
func (*SoWitness) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{0}
}

func (m *SoWitness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoWitness.Unmarshal(m, b)
}
func (m *SoWitness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoWitness.Marshal(b, m, deterministic)
}
func (m *SoWitness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoWitness.Merge(m, src)
}
func (m *SoWitness) XXX_Size() int {
	return xxx_messageInfo_SoWitness.Size(m)
}
func (m *SoWitness) XXX_DiscardUnknown() {
	xxx_messageInfo_SoWitness.DiscardUnknown(m)
}

var xxx_messageInfo_SoWitness proto.InternalMessageInfo

func (m *SoWitness) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *SoWitness) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoWitness) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SoWitness) GetVoteCount() *prototype.Vest {
	if m != nil {
		return m.VoteCount
	}
	return nil
}

func (m *SoWitness) GetSigningKey() *prototype.PublicKeyType {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

func (m *SoWitness) GetProposedStaminaFree() uint64 {
	if m != nil {
		return m.ProposedStaminaFree
	}
	return 0
}

func (m *SoWitness) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *SoWitness) GetTpsExpected() uint64 {
	if m != nil {
		return m.TpsExpected
	}
	return 0
}

func (m *SoWitness) GetAccountCreateFee() *prototype.Coin {
	if m != nil {
		return m.AccountCreateFee
	}
	return nil
}

func (m *SoWitness) GetTopNAcquireFreeToken() uint32 {
	if m != nil {
		return m.TopNAcquireFreeToken
	}
	return 0
}

func (m *SoWitness) GetEpochDuration() uint64 {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func (m *SoWitness) GetPerTicketPrice() *prototype.Vest {
	if m != nil {
		return m.PerTicketPrice
	}
	return nil
}

func (m *SoWitness) GetPerTicketWeight() uint64 {
	if m != nil {
		return m.PerTicketWeight
	}
	return 0
}

func (m *SoWitness) GetVoterList() []*prototype.AccountName {
	if m != nil {
		return m.VoterList
	}
	return nil
}

type SoMemWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoMemWitnessByOwner) Reset()         { *m = SoMemWitnessByOwner{} }
func (m *SoMemWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByOwner) ProtoMessage()    {}
func (*SoMemWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{1}
}

func (m *SoMemWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByOwner.Unmarshal(m, b)
}
func (m *SoMemWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByOwner.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByOwner.Merge(m, src)
}
func (m *SoMemWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByOwner.Size(m)
}
func (m *SoMemWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByOwner proto.InternalMessageInfo

func (m *SoMemWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoMemWitnessByCreatedTime struct {
	CreatedTime          *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemWitnessByCreatedTime) Reset()         { *m = SoMemWitnessByCreatedTime{} }
func (m *SoMemWitnessByCreatedTime) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByCreatedTime) ProtoMessage()    {}
func (*SoMemWitnessByCreatedTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{2}
}

func (m *SoMemWitnessByCreatedTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByCreatedTime.Unmarshal(m, b)
}
func (m *SoMemWitnessByCreatedTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByCreatedTime.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByCreatedTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByCreatedTime.Merge(m, src)
}
func (m *SoMemWitnessByCreatedTime) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByCreatedTime.Size(m)
}
func (m *SoMemWitnessByCreatedTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByCreatedTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByCreatedTime proto.InternalMessageInfo

func (m *SoMemWitnessByCreatedTime) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

type SoMemWitnessByUrl struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByUrl) Reset()         { *m = SoMemWitnessByUrl{} }
func (m *SoMemWitnessByUrl) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByUrl) ProtoMessage()    {}
func (*SoMemWitnessByUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{3}
}

func (m *SoMemWitnessByUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByUrl.Unmarshal(m, b)
}
func (m *SoMemWitnessByUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByUrl.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByUrl.Merge(m, src)
}
func (m *SoMemWitnessByUrl) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByUrl.Size(m)
}
func (m *SoMemWitnessByUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByUrl.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByUrl proto.InternalMessageInfo

func (m *SoMemWitnessByUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type SoMemWitnessByVoteCount struct {
	VoteCount            *prototype.Vest `protobuf:"bytes,1,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SoMemWitnessByVoteCount) Reset()         { *m = SoMemWitnessByVoteCount{} }
func (m *SoMemWitnessByVoteCount) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByVoteCount) ProtoMessage()    {}
func (*SoMemWitnessByVoteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{4}
}

func (m *SoMemWitnessByVoteCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByVoteCount.Unmarshal(m, b)
}
func (m *SoMemWitnessByVoteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByVoteCount.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByVoteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByVoteCount.Merge(m, src)
}
func (m *SoMemWitnessByVoteCount) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByVoteCount.Size(m)
}
func (m *SoMemWitnessByVoteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByVoteCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByVoteCount proto.InternalMessageInfo

func (m *SoMemWitnessByVoteCount) GetVoteCount() *prototype.Vest {
	if m != nil {
		return m.VoteCount
	}
	return nil
}

type SoMemWitnessBySigningKey struct {
	SigningKey           *prototype.PublicKeyType `protobuf:"bytes,1,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SoMemWitnessBySigningKey) Reset()         { *m = SoMemWitnessBySigningKey{} }
func (m *SoMemWitnessBySigningKey) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessBySigningKey) ProtoMessage()    {}
func (*SoMemWitnessBySigningKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{5}
}

func (m *SoMemWitnessBySigningKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessBySigningKey.Unmarshal(m, b)
}
func (m *SoMemWitnessBySigningKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessBySigningKey.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessBySigningKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessBySigningKey.Merge(m, src)
}
func (m *SoMemWitnessBySigningKey) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessBySigningKey.Size(m)
}
func (m *SoMemWitnessBySigningKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessBySigningKey.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessBySigningKey proto.InternalMessageInfo

func (m *SoMemWitnessBySigningKey) GetSigningKey() *prototype.PublicKeyType {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

type SoMemWitnessByProposedStaminaFree struct {
	ProposedStaminaFree  uint64   `protobuf:"varint,1,opt,name=proposed_stamina_free,json=proposedStaminaFree,proto3" json:"proposed_stamina_free,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByProposedStaminaFree) Reset()         { *m = SoMemWitnessByProposedStaminaFree{} }
func (m *SoMemWitnessByProposedStaminaFree) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByProposedStaminaFree) ProtoMessage()    {}
func (*SoMemWitnessByProposedStaminaFree) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{6}
}

func (m *SoMemWitnessByProposedStaminaFree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByProposedStaminaFree.Unmarshal(m, b)
}
func (m *SoMemWitnessByProposedStaminaFree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByProposedStaminaFree.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByProposedStaminaFree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByProposedStaminaFree.Merge(m, src)
}
func (m *SoMemWitnessByProposedStaminaFree) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByProposedStaminaFree.Size(m)
}
func (m *SoMemWitnessByProposedStaminaFree) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByProposedStaminaFree.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByProposedStaminaFree proto.InternalMessageInfo

func (m *SoMemWitnessByProposedStaminaFree) GetProposedStaminaFree() uint64 {
	if m != nil {
		return m.ProposedStaminaFree
	}
	return 0
}

type SoMemWitnessByActive struct {
	Active               bool     `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByActive) Reset()         { *m = SoMemWitnessByActive{} }
func (m *SoMemWitnessByActive) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByActive) ProtoMessage()    {}
func (*SoMemWitnessByActive) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{7}
}

func (m *SoMemWitnessByActive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByActive.Unmarshal(m, b)
}
func (m *SoMemWitnessByActive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByActive.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByActive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByActive.Merge(m, src)
}
func (m *SoMemWitnessByActive) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByActive.Size(m)
}
func (m *SoMemWitnessByActive) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByActive.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByActive proto.InternalMessageInfo

func (m *SoMemWitnessByActive) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type SoMemWitnessByTpsExpected struct {
	TpsExpected          uint64   `protobuf:"varint,1,opt,name=tps_expected,json=tpsExpected,proto3" json:"tps_expected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByTpsExpected) Reset()         { *m = SoMemWitnessByTpsExpected{} }
func (m *SoMemWitnessByTpsExpected) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByTpsExpected) ProtoMessage()    {}
func (*SoMemWitnessByTpsExpected) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{8}
}

func (m *SoMemWitnessByTpsExpected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByTpsExpected.Unmarshal(m, b)
}
func (m *SoMemWitnessByTpsExpected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByTpsExpected.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByTpsExpected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByTpsExpected.Merge(m, src)
}
func (m *SoMemWitnessByTpsExpected) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByTpsExpected.Size(m)
}
func (m *SoMemWitnessByTpsExpected) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByTpsExpected.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByTpsExpected proto.InternalMessageInfo

func (m *SoMemWitnessByTpsExpected) GetTpsExpected() uint64 {
	if m != nil {
		return m.TpsExpected
	}
	return 0
}

type SoMemWitnessByAccountCreateFee struct {
	AccountCreateFee     *prototype.Coin `protobuf:"bytes,1,opt,name=account_create_fee,json=accountCreateFee,proto3" json:"account_create_fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SoMemWitnessByAccountCreateFee) Reset()         { *m = SoMemWitnessByAccountCreateFee{} }
func (m *SoMemWitnessByAccountCreateFee) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByAccountCreateFee) ProtoMessage()    {}
func (*SoMemWitnessByAccountCreateFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{9}
}

func (m *SoMemWitnessByAccountCreateFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByAccountCreateFee.Unmarshal(m, b)
}
func (m *SoMemWitnessByAccountCreateFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByAccountCreateFee.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByAccountCreateFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByAccountCreateFee.Merge(m, src)
}
func (m *SoMemWitnessByAccountCreateFee) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByAccountCreateFee.Size(m)
}
func (m *SoMemWitnessByAccountCreateFee) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByAccountCreateFee.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByAccountCreateFee proto.InternalMessageInfo

func (m *SoMemWitnessByAccountCreateFee) GetAccountCreateFee() *prototype.Coin {
	if m != nil {
		return m.AccountCreateFee
	}
	return nil
}

type SoMemWitnessByTopNAcquireFreeToken struct {
	TopNAcquireFreeToken uint32   `protobuf:"varint,1,opt,name=top_n_acquire_free_token,json=topNAcquireFreeToken,proto3" json:"top_n_acquire_free_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByTopNAcquireFreeToken) Reset()         { *m = SoMemWitnessByTopNAcquireFreeToken{} }
func (m *SoMemWitnessByTopNAcquireFreeToken) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByTopNAcquireFreeToken) ProtoMessage()    {}
func (*SoMemWitnessByTopNAcquireFreeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{10}
}

func (m *SoMemWitnessByTopNAcquireFreeToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByTopNAcquireFreeToken.Unmarshal(m, b)
}
func (m *SoMemWitnessByTopNAcquireFreeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByTopNAcquireFreeToken.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByTopNAcquireFreeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByTopNAcquireFreeToken.Merge(m, src)
}
func (m *SoMemWitnessByTopNAcquireFreeToken) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByTopNAcquireFreeToken.Size(m)
}
func (m *SoMemWitnessByTopNAcquireFreeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByTopNAcquireFreeToken.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByTopNAcquireFreeToken proto.InternalMessageInfo

func (m *SoMemWitnessByTopNAcquireFreeToken) GetTopNAcquireFreeToken() uint32 {
	if m != nil {
		return m.TopNAcquireFreeToken
	}
	return 0
}

type SoMemWitnessByEpochDuration struct {
	EpochDuration        uint64   `protobuf:"varint,1,opt,name=epoch_duration,json=epochDuration,proto3" json:"epoch_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByEpochDuration) Reset()         { *m = SoMemWitnessByEpochDuration{} }
func (m *SoMemWitnessByEpochDuration) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByEpochDuration) ProtoMessage()    {}
func (*SoMemWitnessByEpochDuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{11}
}

func (m *SoMemWitnessByEpochDuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByEpochDuration.Unmarshal(m, b)
}
func (m *SoMemWitnessByEpochDuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByEpochDuration.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByEpochDuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByEpochDuration.Merge(m, src)
}
func (m *SoMemWitnessByEpochDuration) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByEpochDuration.Size(m)
}
func (m *SoMemWitnessByEpochDuration) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByEpochDuration.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByEpochDuration proto.InternalMessageInfo

func (m *SoMemWitnessByEpochDuration) GetEpochDuration() uint64 {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

type SoMemWitnessByPerTicketPrice struct {
	PerTicketPrice       *prototype.Vest `protobuf:"bytes,1,opt,name=per_ticket_price,json=perTicketPrice,proto3" json:"per_ticket_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SoMemWitnessByPerTicketPrice) Reset()         { *m = SoMemWitnessByPerTicketPrice{} }
func (m *SoMemWitnessByPerTicketPrice) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByPerTicketPrice) ProtoMessage()    {}
func (*SoMemWitnessByPerTicketPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{12}
}

func (m *SoMemWitnessByPerTicketPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByPerTicketPrice.Unmarshal(m, b)
}
func (m *SoMemWitnessByPerTicketPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByPerTicketPrice.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByPerTicketPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByPerTicketPrice.Merge(m, src)
}
func (m *SoMemWitnessByPerTicketPrice) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByPerTicketPrice.Size(m)
}
func (m *SoMemWitnessByPerTicketPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByPerTicketPrice.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByPerTicketPrice proto.InternalMessageInfo

func (m *SoMemWitnessByPerTicketPrice) GetPerTicketPrice() *prototype.Vest {
	if m != nil {
		return m.PerTicketPrice
	}
	return nil
}

type SoMemWitnessByPerTicketWeight struct {
	PerTicketWeight      uint64   `protobuf:"varint,1,opt,name=per_ticket_weight,json=perTicketWeight,proto3" json:"per_ticket_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemWitnessByPerTicketWeight) Reset()         { *m = SoMemWitnessByPerTicketWeight{} }
func (m *SoMemWitnessByPerTicketWeight) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByPerTicketWeight) ProtoMessage()    {}
func (*SoMemWitnessByPerTicketWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{13}
}

func (m *SoMemWitnessByPerTicketWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByPerTicketWeight.Unmarshal(m, b)
}
func (m *SoMemWitnessByPerTicketWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByPerTicketWeight.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByPerTicketWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByPerTicketWeight.Merge(m, src)
}
func (m *SoMemWitnessByPerTicketWeight) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByPerTicketWeight.Size(m)
}
func (m *SoMemWitnessByPerTicketWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByPerTicketWeight.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByPerTicketWeight proto.InternalMessageInfo

func (m *SoMemWitnessByPerTicketWeight) GetPerTicketWeight() uint64 {
	if m != nil {
		return m.PerTicketWeight
	}
	return 0
}

type SoMemWitnessByVoterList struct {
	VoterList            []*prototype.AccountName `protobuf:"bytes,1,rep,name=voter_list,json=voterList,proto3" json:"voter_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SoMemWitnessByVoterList) Reset()         { *m = SoMemWitnessByVoterList{} }
func (m *SoMemWitnessByVoterList) String() string { return proto.CompactTextString(m) }
func (*SoMemWitnessByVoterList) ProtoMessage()    {}
func (*SoMemWitnessByVoterList) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{14}
}

func (m *SoMemWitnessByVoterList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemWitnessByVoterList.Unmarshal(m, b)
}
func (m *SoMemWitnessByVoterList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemWitnessByVoterList.Marshal(b, m, deterministic)
}
func (m *SoMemWitnessByVoterList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemWitnessByVoterList.Merge(m, src)
}
func (m *SoMemWitnessByVoterList) XXX_Size() int {
	return xxx_messageInfo_SoMemWitnessByVoterList.Size(m)
}
func (m *SoMemWitnessByVoterList) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemWitnessByVoterList.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemWitnessByVoterList proto.InternalMessageInfo

func (m *SoMemWitnessByVoterList) GetVoterList() []*prototype.AccountName {
	if m != nil {
		return m.VoterList
	}
	return nil
}

type SoListWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListWitnessByOwner) Reset()         { *m = SoListWitnessByOwner{} }
func (m *SoListWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoListWitnessByOwner) ProtoMessage()    {}
func (*SoListWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{15}
}

func (m *SoListWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListWitnessByOwner.Unmarshal(m, b)
}
func (m *SoListWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListWitnessByOwner.Marshal(b, m, deterministic)
}
func (m *SoListWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListWitnessByOwner.Merge(m, src)
}
func (m *SoListWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoListWitnessByOwner.Size(m)
}
func (m *SoListWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoListWitnessByOwner proto.InternalMessageInfo

func (m *SoListWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoListWitnessByVoteCount struct {
	VoteCount            *prototype.Vest        `protobuf:"bytes,1,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	Owner                *prototype.AccountName `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListWitnessByVoteCount) Reset()         { *m = SoListWitnessByVoteCount{} }
func (m *SoListWitnessByVoteCount) String() string { return proto.CompactTextString(m) }
func (*SoListWitnessByVoteCount) ProtoMessage()    {}
func (*SoListWitnessByVoteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{16}
}

func (m *SoListWitnessByVoteCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListWitnessByVoteCount.Unmarshal(m, b)
}
func (m *SoListWitnessByVoteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListWitnessByVoteCount.Marshal(b, m, deterministic)
}
func (m *SoListWitnessByVoteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListWitnessByVoteCount.Merge(m, src)
}
func (m *SoListWitnessByVoteCount) XXX_Size() int {
	return xxx_messageInfo_SoListWitnessByVoteCount.Size(m)
}
func (m *SoListWitnessByVoteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListWitnessByVoteCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListWitnessByVoteCount proto.InternalMessageInfo

func (m *SoListWitnessByVoteCount) GetVoteCount() *prototype.Vest {
	if m != nil {
		return m.VoteCount
	}
	return nil
}

func (m *SoListWitnessByVoteCount) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoUniqueWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoUniqueWitnessByOwner) Reset()         { *m = SoUniqueWitnessByOwner{} }
func (m *SoUniqueWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoUniqueWitnessByOwner) ProtoMessage()    {}
func (*SoUniqueWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{17}
}

func (m *SoUniqueWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Unmarshal(m, b)
}
func (m *SoUniqueWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Marshal(b, m, deterministic)
}
func (m *SoUniqueWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueWitnessByOwner.Merge(m, src)
}
func (m *SoUniqueWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Size(m)
}
func (m *SoUniqueWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueWitnessByOwner proto.InternalMessageInfo

func (m *SoUniqueWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterType((*SoWitness)(nil), "table.so_witness")
	proto.RegisterType((*SoMemWitnessByOwner)(nil), "table.so_mem_witness_by_owner")
	proto.RegisterType((*SoMemWitnessByCreatedTime)(nil), "table.so_mem_witness_by_created_time")
	proto.RegisterType((*SoMemWitnessByUrl)(nil), "table.so_mem_witness_by_url")
	proto.RegisterType((*SoMemWitnessByVoteCount)(nil), "table.so_mem_witness_by_vote_count")
	proto.RegisterType((*SoMemWitnessBySigningKey)(nil), "table.so_mem_witness_by_signing_key")
	proto.RegisterType((*SoMemWitnessByProposedStaminaFree)(nil), "table.so_mem_witness_by_proposed_stamina_free")
	proto.RegisterType((*SoMemWitnessByActive)(nil), "table.so_mem_witness_by_active")
	proto.RegisterType((*SoMemWitnessByTpsExpected)(nil), "table.so_mem_witness_by_tps_expected")
	proto.RegisterType((*SoMemWitnessByAccountCreateFee)(nil), "table.so_mem_witness_by_account_create_fee")
	proto.RegisterType((*SoMemWitnessByTopNAcquireFreeToken)(nil), "table.so_mem_witness_by_top_n_acquire_free_token")
	proto.RegisterType((*SoMemWitnessByEpochDuration)(nil), "table.so_mem_witness_by_epoch_duration")
	proto.RegisterType((*SoMemWitnessByPerTicketPrice)(nil), "table.so_mem_witness_by_per_ticket_price")
	proto.RegisterType((*SoMemWitnessByPerTicketWeight)(nil), "table.so_mem_witness_by_per_ticket_weight")
	proto.RegisterType((*SoMemWitnessByVoterList)(nil), "table.so_mem_witness_by_voter_list")
	proto.RegisterType((*SoListWitnessByOwner)(nil), "table.so_list_witness_by_owner")
	proto.RegisterType((*SoListWitnessByVoteCount)(nil), "table.so_list_witness_by_vote_count")
	proto.RegisterType((*SoUniqueWitnessByOwner)(nil), "table.so_unique_witness_by_owner")
}

func init() { proto.RegisterFile("app/table/so_witness.proto", fileDescriptor_00097e516fc05425) }

var fileDescriptor_00097e516fc05425 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0x5d, 0x4f, 0x2b, 0x45,
	0x18, 0xc7, 0x33, 0x87, 0x53, 0x3c, 0x3c, 0xe5, 0xcd, 0x11, 0x64, 0x6c, 0xd4, 0xd4, 0x51, 0x63,
	0x25, 0xd2, 0x26, 0x98, 0x90, 0x18, 0xf5, 0x42, 0x51, 0x22, 0xc1, 0x10, 0x5d, 0x89, 0x26, 0x46,
	0x9d, 0x6c, 0xa7, 0x0f, 0xed, 0x84, 0x76, 0x66, 0xd8, 0x99, 0x05, 0x7b, 0xe3, 0xa7, 0xf1, 0x83,
	0x9a, 0x9d, 0x5d, 0xda, 0xb5, 0xbb, 0xc5, 0x22, 0x37, 0x84, 0x7d, 0x5e, 0x67, 0x77, 0xfe, 0xbf,
	0x7f, 0x0a, 0xad, 0xd8, 0xda, 0x9e, 0x8f, 0xfb, 0x63, 0xec, 0x39, 0x23, 0xee, 0x95, 0xd7, 0xe8,
	0x5c, 0xd7, 0x26, 0xc6, 0x1b, 0xda, 0x08, 0xf1, 0xd6, 0x5e, 0x78, 0xf2, 0x53, 0x8b, 0xbd, 0xec,
	0x4f, 0x9e, 0xe4, 0x7f, 0x37, 0x00, 0xe6, 0x1d, 0xf4, 0x08, 0x1a, 0xe6, 0x5e, 0x63, 0xc2, 0x48,
	0x9b, 0x74, 0x9a, 0xc7, 0x07, 0xdd, 0x59, 0x53, 0x37, 0x96, 0xd2, 0xa4, 0xda, 0x0b, 0x1d, 0x4f,
	0x30, 0xca, 0xab, 0xe8, 0x17, 0xb0, 0x29, 0x13, 0x8c, 0x3d, 0x0e, 0x84, 0x57, 0x13, 0x64, 0x2f,
	0x42, 0xd7, 0x5b, 0xa5, 0xae, 0x2c, 0x2c, 0xac, 0x51, 0xda, 0x0b, 0x87, 0x32, 0x6a, 0x16, 0xe5,
	0x57, 0x6a, 0x82, 0x74, 0x17, 0xd6, 0xd2, 0x64, 0xcc, 0xd6, 0xda, 0xa4, 0xb3, 0x11, 0x65, 0xff,
	0xd2, 0x2e, 0xc0, 0x9d, 0xf1, 0x28, 0xc2, 0x26, 0xf6, 0x32, 0x4c, 0xdb, 0x29, 0x4d, 0xbb, 0x43,
	0xe7, 0xa3, 0x8d, 0xac, 0xe4, 0x34, 0xab, 0xa0, 0x9f, 0x43, 0xd3, 0xa9, 0xa1, 0x56, 0x7a, 0x28,
	0x6e, 0x70, 0xca, 0x1a, 0xa1, 0xa1, 0x55, 0x6a, 0xb0, 0x69, 0x7f, 0xac, 0x64, 0x96, 0x14, 0xd9,
	0x73, 0x04, 0x45, 0xf9, 0x05, 0x4e, 0xe9, 0x31, 0xec, 0xdb, 0xc4, 0x58, 0xe3, 0x70, 0x20, 0x9c,
	0x8f, 0x27, 0x4a, 0xc7, 0xe2, 0x3a, 0x41, 0x64, 0xeb, 0x6d, 0xd2, 0x79, 0x19, 0xbd, 0xf1, 0x90,
	0xfc, 0x29, 0xcf, 0x9d, 0x25, 0x88, 0xf4, 0x4d, 0x58, 0x8f, 0xa5, 0x57, 0x77, 0xc8, 0x5e, 0x6b,
	0x93, 0xce, 0xab, 0xa8, 0x78, 0xa2, 0xef, 0xc1, 0xa6, 0xb7, 0x4e, 0xe0, 0x9f, 0x16, 0xa5, 0xc7,
	0x01, 0x7b, 0x15, 0x46, 0x34, 0xbd, 0x75, 0xdf, 0x16, 0x21, 0xfa, 0x25, 0xd0, 0x87, 0x4f, 0x98,
	0x7f, 0x04, 0x71, 0x8d, 0xc8, 0x36, 0x2a, 0xef, 0x28, 0x8d, 0xd2, 0xd1, 0x6e, 0x51, 0x7a, 0x1a,
	0x2a, 0xcf, 0x10, 0xe9, 0x09, 0x30, 0x6f, 0xac, 0xd0, 0x22, 0x96, 0xb7, 0xa9, 0x4a, 0x30, 0x1c,
	0x55, 0x78, 0x73, 0x83, 0x9a, 0x41, 0x9b, 0x74, 0xb6, 0xa2, 0x3d, 0x6f, 0xec, 0xe5, 0x57, 0x79,
	0x36, 0x3b, 0xec, 0x55, 0x96, 0xa3, 0x1f, 0xc2, 0x36, 0x5a, 0x23, 0x47, 0x62, 0x90, 0x26, 0xb1,
	0x57, 0x46, 0xb3, 0x66, 0x38, 0xdb, 0x56, 0x88, 0x7e, 0x53, 0x04, 0xe9, 0x67, 0xb0, 0x6b, 0x31,
	0x11, 0x5e, 0xc9, 0x1b, 0xf4, 0xc2, 0x26, 0x4a, 0x22, 0xdb, 0xac, 0xff, 0xfe, 0xdb, 0x16, 0x93,
	0xab, 0x50, 0xf7, 0x43, 0x56, 0x46, 0x0f, 0xe1, 0xf5, 0x52, 0xeb, 0x3d, 0xaa, 0xe1, 0xc8, 0xb3,
	0xad, 0xb0, 0x64, 0x67, 0x56, 0xfa, 0x4b, 0x08, 0xd3, 0x93, 0xfc, 0x82, 0x13, 0x31, 0x56, 0xce,
	0xb3, 0xed, 0xf6, 0xda, 0x63, 0x22, 0x0b, 0x17, 0x9d, 0x7c, 0xaf, 0x9c, 0xe7, 0xdf, 0xc1, 0x81,
	0x33, 0x62, 0x82, 0x93, 0x07, 0xa5, 0x8a, 0xfe, 0x54, 0xe4, 0x1a, 0x7c, 0x9a, 0x64, 0xf9, 0x1f,
	0xf0, 0x6e, 0x75, 0x52, 0x59, 0xc4, 0x15, 0x51, 0x93, 0xa7, 0x88, 0x9a, 0x7f, 0x0c, 0xfb, 0xd5,
	0xf9, 0x99, 0xb6, 0x0b, 0xb5, 0x93, 0x99, 0xda, 0xf9, 0x25, 0xbc, 0x5d, 0x2d, 0x9d, 0xeb, 0x7f,
	0x81, 0x06, 0xf2, 0x5f, 0x34, 0xf0, 0xdf, 0xe0, 0x9d, 0xea, 0xbc, 0x12, 0x1f, 0x8b, 0xb8, 0x90,
	0xa7, 0xe0, 0xc2, 0x7f, 0x87, 0x8f, 0xaa, 0xd3, 0x6b, 0x01, 0x5a, 0x4e, 0x16, 0x59, 0x4a, 0x16,
	0x3f, 0x06, 0x56, 0x1d, 0x5f, 0xd0, 0x35, 0xa7, 0x8e, 0x94, 0xa9, 0xe3, 0xa7, 0x75, 0x77, 0x59,
	0xe6, 0xb0, 0xc2, 0x25, 0xa9, 0x70, 0xc9, 0x11, 0x3e, 0xa8, 0x5b, 0xbc, 0x48, 0xea, 0x12, 0x7e,
	0xc9, 0x8a, 0xfc, 0xf2, 0x01, 0x1c, 0xd6, 0x9c, 0x75, 0x09, 0xd1, 0x8f, 0xd2, 0x4e, 0x96, 0xd3,
	0xce, 0xcf, 0xa1, 0x5d, 0xdd, 0xf2, 0x6f, 0xfe, 0x6b, 0x1c, 0x81, 0xd4, 0x38, 0x02, 0x17, 0xc0,
	0x6b, 0xee, 0x7b, 0xc1, 0x23, 0x6a, 0x7d, 0x83, 0xac, 0xe4, 0x1b, 0xfc, 0x47, 0x78, 0xff, 0xd1,
	0x05, 0xb9, 0x93, 0xd4, 0xdb, 0x0b, 0xa9, 0xb5, 0x17, 0xfe, 0xf3, 0x32, 0xa2, 0x72, 0xc3, 0x59,
	0xb0, 0x1f, 0xb2, 0xb2, 0xfd, 0x9c, 0x07, 0x71, 0x66, 0x4d, 0xcf, 0xf6, 0x9f, 0xbf, 0x02, 0xa4,
	0x8b, 0xa3, 0xfe, 0x3f, 0xf5, 0xf3, 0xfd, 0x2f, 0x56, 0xda, 0x7f, 0x01, 0x2d, 0x67, 0x44, 0xaa,
	0xd5, 0x6d, 0x8a, 0xcf, 0x7d, 0x99, 0xaf, 0x3f, 0xf9, 0xf5, 0x70, 0xa8, 0xfc, 0x28, 0xed, 0x77,
	0xa5, 0x99, 0xf4, 0xa4, 0x71, 0x72, 0x14, 0x2b, 0xdd, 0x93, 0x46, 0x7b, 0xd4, 0xde, 0xb8, 0xa3,
	0xa1, 0xe9, 0xcd, 0x7e, 0x99, 0xf4, 0xd7, 0xc3, 0xb0, 0x4f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0x94, 0xa6, 0xb2, 0xad, 0x08, 0x00, 0x00,
}
