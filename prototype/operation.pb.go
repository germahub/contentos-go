// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototype/operation.proto

package prototype

import (
	fmt "fmt"
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

type AccountCreateOperation struct {
	Fee                  *Coin        `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee,omitempty"`
	Creator              *AccountName `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	NewAccountName       *AccountName `protobuf:"bytes,3,opt,name=new_account_name,json=newAccountName,proto3" json:"new_account_name,omitempty"`
	Owner                *Authority   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	JsonMetadata         string       `protobuf:"bytes,8,opt,name=json_metadata,json=jsonMetadata,proto3" json:"json_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AccountCreateOperation) Reset()         { *m = AccountCreateOperation{} }
func (m *AccountCreateOperation) String() string { return proto.CompactTextString(m) }
func (*AccountCreateOperation) ProtoMessage()    {}
func (*AccountCreateOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{0}
}

func (m *AccountCreateOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateOperation.Unmarshal(m, b)
}
func (m *AccountCreateOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateOperation.Marshal(b, m, deterministic)
}
func (m *AccountCreateOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateOperation.Merge(m, src)
}
func (m *AccountCreateOperation) XXX_Size() int {
	return xxx_messageInfo_AccountCreateOperation.Size(m)
}
func (m *AccountCreateOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateOperation proto.InternalMessageInfo

func (m *AccountCreateOperation) GetFee() *Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *AccountCreateOperation) GetCreator() *AccountName {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *AccountCreateOperation) GetNewAccountName() *AccountName {
	if m != nil {
		return m.NewAccountName
	}
	return nil
}

func (m *AccountCreateOperation) GetOwner() *Authority {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AccountCreateOperation) GetJsonMetadata() string {
	if m != nil {
		return m.JsonMetadata
	}
	return ""
}

type TransferOperation struct {
	From                 *AccountName `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *AccountName `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               *Coin        `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Memo                 string       `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TransferOperation) Reset()         { *m = TransferOperation{} }
func (m *TransferOperation) String() string { return proto.CompactTextString(m) }
func (*TransferOperation) ProtoMessage()    {}
func (*TransferOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{1}
}

func (m *TransferOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferOperation.Unmarshal(m, b)
}
func (m *TransferOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferOperation.Marshal(b, m, deterministic)
}
func (m *TransferOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferOperation.Merge(m, src)
}
func (m *TransferOperation) XXX_Size() int {
	return xxx_messageInfo_TransferOperation.Size(m)
}
func (m *TransferOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferOperation.DiscardUnknown(m)
}

var xxx_messageInfo_TransferOperation proto.InternalMessageInfo

func (m *TransferOperation) GetFrom() *AccountName {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransferOperation) GetTo() *AccountName {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferOperation) GetAmount() *Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *TransferOperation) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

type TransferToVestingOperation struct {
	From                 *AccountName `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *AccountName `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               *Coin        `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TransferToVestingOperation) Reset()         { *m = TransferToVestingOperation{} }
func (m *TransferToVestingOperation) String() string { return proto.CompactTextString(m) }
func (*TransferToVestingOperation) ProtoMessage()    {}
func (*TransferToVestingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{2}
}

func (m *TransferToVestingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferToVestingOperation.Unmarshal(m, b)
}
func (m *TransferToVestingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferToVestingOperation.Marshal(b, m, deterministic)
}
func (m *TransferToVestingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferToVestingOperation.Merge(m, src)
}
func (m *TransferToVestingOperation) XXX_Size() int {
	return xxx_messageInfo_TransferToVestingOperation.Size(m)
}
func (m *TransferToVestingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferToVestingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_TransferToVestingOperation proto.InternalMessageInfo

func (m *TransferToVestingOperation) GetFrom() *AccountName {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransferToVestingOperation) GetTo() *AccountName {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferToVestingOperation) GetAmount() *Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

type VoteOperation struct {
	Voter                *AccountName `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Idx                  uint64       `protobuf:"varint,3,opt,name=idx,proto3" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VoteOperation) Reset()         { *m = VoteOperation{} }
func (m *VoteOperation) String() string { return proto.CompactTextString(m) }
func (*VoteOperation) ProtoMessage()    {}
func (*VoteOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{3}
}

func (m *VoteOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteOperation.Unmarshal(m, b)
}
func (m *VoteOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteOperation.Marshal(b, m, deterministic)
}
func (m *VoteOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteOperation.Merge(m, src)
}
func (m *VoteOperation) XXX_Size() int {
	return xxx_messageInfo_VoteOperation.Size(m)
}
func (m *VoteOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteOperation.DiscardUnknown(m)
}

var xxx_messageInfo_VoteOperation proto.InternalMessageInfo

func (m *VoteOperation) GetVoter() *AccountName {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *VoteOperation) GetIdx() uint64 {
	if m != nil {
		return m.Idx
	}
	return 0
}

type BpRegisterOperation struct {
	Owner                *AccountName     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Url                  string           `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Desc                 string           `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	BlockSigningKey      *PublicKeyType   `protobuf:"bytes,4,opt,name=block_signing_key,json=blockSigningKey,proto3" json:"block_signing_key,omitempty"`
	Props                *ChainProperties `protobuf:"bytes,5,opt,name=props,proto3" json:"props,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BpRegisterOperation) Reset()         { *m = BpRegisterOperation{} }
func (m *BpRegisterOperation) String() string { return proto.CompactTextString(m) }
func (*BpRegisterOperation) ProtoMessage()    {}
func (*BpRegisterOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{4}
}

func (m *BpRegisterOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BpRegisterOperation.Unmarshal(m, b)
}
func (m *BpRegisterOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BpRegisterOperation.Marshal(b, m, deterministic)
}
func (m *BpRegisterOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BpRegisterOperation.Merge(m, src)
}
func (m *BpRegisterOperation) XXX_Size() int {
	return xxx_messageInfo_BpRegisterOperation.Size(m)
}
func (m *BpRegisterOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BpRegisterOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BpRegisterOperation proto.InternalMessageInfo

func (m *BpRegisterOperation) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *BpRegisterOperation) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *BpRegisterOperation) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *BpRegisterOperation) GetBlockSigningKey() *PublicKeyType {
	if m != nil {
		return m.BlockSigningKey
	}
	return nil
}

func (m *BpRegisterOperation) GetProps() *ChainProperties {
	if m != nil {
		return m.Props
	}
	return nil
}

type BpUnregisterOperation struct {
	Owner                *AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BpUnregisterOperation) Reset()         { *m = BpUnregisterOperation{} }
func (m *BpUnregisterOperation) String() string { return proto.CompactTextString(m) }
func (*BpUnregisterOperation) ProtoMessage()    {}
func (*BpUnregisterOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{5}
}

func (m *BpUnregisterOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BpUnregisterOperation.Unmarshal(m, b)
}
func (m *BpUnregisterOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BpUnregisterOperation.Marshal(b, m, deterministic)
}
func (m *BpUnregisterOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BpUnregisterOperation.Merge(m, src)
}
func (m *BpUnregisterOperation) XXX_Size() int {
	return xxx_messageInfo_BpUnregisterOperation.Size(m)
}
func (m *BpUnregisterOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BpUnregisterOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BpUnregisterOperation proto.InternalMessageInfo

func (m *BpUnregisterOperation) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type BpVoteOperation struct {
	Voter                *AccountName `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Witness              *AccountName `protobuf:"bytes,2,opt,name=witness,proto3" json:"witness,omitempty"`
	Cancel               bool         `protobuf:"varint,3,opt,name=cancel,proto3" json:"cancel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BpVoteOperation) Reset()         { *m = BpVoteOperation{} }
func (m *BpVoteOperation) String() string { return proto.CompactTextString(m) }
func (*BpVoteOperation) ProtoMessage()    {}
func (*BpVoteOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{6}
}

func (m *BpVoteOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BpVoteOperation.Unmarshal(m, b)
}
func (m *BpVoteOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BpVoteOperation.Marshal(b, m, deterministic)
}
func (m *BpVoteOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BpVoteOperation.Merge(m, src)
}
func (m *BpVoteOperation) XXX_Size() int {
	return xxx_messageInfo_BpVoteOperation.Size(m)
}
func (m *BpVoteOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BpVoteOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BpVoteOperation proto.InternalMessageInfo

func (m *BpVoteOperation) GetVoter() *AccountName {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *BpVoteOperation) GetWitness() *AccountName {
	if m != nil {
		return m.Witness
	}
	return nil
}

func (m *BpVoteOperation) GetCancel() bool {
	if m != nil {
		return m.Cancel
	}
	return false
}

type FollowOperation struct {
	Account              *AccountName `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	FAccount             *AccountName `protobuf:"bytes,2,opt,name=f_account,json=fAccount,proto3" json:"f_account,omitempty"`
	Cancel               bool         `protobuf:"varint,3,opt,name=cancel,proto3" json:"cancel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FollowOperation) Reset()         { *m = FollowOperation{} }
func (m *FollowOperation) String() string { return proto.CompactTextString(m) }
func (*FollowOperation) ProtoMessage()    {}
func (*FollowOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{7}
}

func (m *FollowOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowOperation.Unmarshal(m, b)
}
func (m *FollowOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowOperation.Marshal(b, m, deterministic)
}
func (m *FollowOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowOperation.Merge(m, src)
}
func (m *FollowOperation) XXX_Size() int {
	return xxx_messageInfo_FollowOperation.Size(m)
}
func (m *FollowOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FollowOperation proto.InternalMessageInfo

func (m *FollowOperation) GetAccount() *AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *FollowOperation) GetFAccount() *AccountName {
	if m != nil {
		return m.FAccount
	}
	return nil
}

func (m *FollowOperation) GetCancel() bool {
	if m != nil {
		return m.Cancel
	}
	return false
}

type ContractDeployOperation struct {
	Owner                *AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Contract             string       `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Abi                  string       `protobuf:"bytes,3,opt,name=abi,proto3" json:"abi,omitempty"`
	Code                 []byte       `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractDeployOperation) Reset()         { *m = ContractDeployOperation{} }
func (m *ContractDeployOperation) String() string { return proto.CompactTextString(m) }
func (*ContractDeployOperation) ProtoMessage()    {}
func (*ContractDeployOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{8}
}

func (m *ContractDeployOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractDeployOperation.Unmarshal(m, b)
}
func (m *ContractDeployOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractDeployOperation.Marshal(b, m, deterministic)
}
func (m *ContractDeployOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractDeployOperation.Merge(m, src)
}
func (m *ContractDeployOperation) XXX_Size() int {
	return xxx_messageInfo_ContractDeployOperation.Size(m)
}
func (m *ContractDeployOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractDeployOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ContractDeployOperation proto.InternalMessageInfo

func (m *ContractDeployOperation) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ContractDeployOperation) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ContractDeployOperation) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *ContractDeployOperation) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type ContractApplyOperation struct {
	Caller               *AccountName `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Owner                *AccountName `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Contract             string       `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	Method               string       `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Params               string       `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
	Amount               *Coin        `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Gas                  *Coin        `protobuf:"bytes,7,opt,name=gas,proto3" json:"gas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractApplyOperation) Reset()         { *m = ContractApplyOperation{} }
func (m *ContractApplyOperation) String() string { return proto.CompactTextString(m) }
func (*ContractApplyOperation) ProtoMessage()    {}
func (*ContractApplyOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{9}
}

func (m *ContractApplyOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractApplyOperation.Unmarshal(m, b)
}
func (m *ContractApplyOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractApplyOperation.Marshal(b, m, deterministic)
}
func (m *ContractApplyOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractApplyOperation.Merge(m, src)
}
func (m *ContractApplyOperation) XXX_Size() int {
	return xxx_messageInfo_ContractApplyOperation.Size(m)
}
func (m *ContractApplyOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractApplyOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ContractApplyOperation proto.InternalMessageInfo

func (m *ContractApplyOperation) GetCaller() *AccountName {
	if m != nil {
		return m.Caller
	}
	return nil
}

func (m *ContractApplyOperation) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ContractApplyOperation) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ContractApplyOperation) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ContractApplyOperation) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *ContractApplyOperation) GetAmount() *Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *ContractApplyOperation) GetGas() *Coin {
	if m != nil {
		return m.Gas
	}
	return nil
}

type ContractEstimateApplyOperation struct {
	Caller               *AccountName `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Owner                *AccountName `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Contract             string       `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	Params               string       `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractEstimateApplyOperation) Reset()         { *m = ContractEstimateApplyOperation{} }
func (m *ContractEstimateApplyOperation) String() string { return proto.CompactTextString(m) }
func (*ContractEstimateApplyOperation) ProtoMessage()    {}
func (*ContractEstimateApplyOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{10}
}

func (m *ContractEstimateApplyOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractEstimateApplyOperation.Unmarshal(m, b)
}
func (m *ContractEstimateApplyOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractEstimateApplyOperation.Marshal(b, m, deterministic)
}
func (m *ContractEstimateApplyOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractEstimateApplyOperation.Merge(m, src)
}
func (m *ContractEstimateApplyOperation) XXX_Size() int {
	return xxx_messageInfo_ContractEstimateApplyOperation.Size(m)
}
func (m *ContractEstimateApplyOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractEstimateApplyOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ContractEstimateApplyOperation proto.InternalMessageInfo

func (m *ContractEstimateApplyOperation) GetCaller() *AccountName {
	if m != nil {
		return m.Caller
	}
	return nil
}

func (m *ContractEstimateApplyOperation) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ContractEstimateApplyOperation) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ContractEstimateApplyOperation) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

type PostOperation struct {
	Uuid                 uint64                  `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Owner                *AccountName            `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Title                string                  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string                  `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Tags                 []string                `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Beneficiaries        []*BeneficiaryRouteType `protobuf:"bytes,6,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PostOperation) Reset()         { *m = PostOperation{} }
func (m *PostOperation) String() string { return proto.CompactTextString(m) }
func (*PostOperation) ProtoMessage()    {}
func (*PostOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{11}
}

func (m *PostOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostOperation.Unmarshal(m, b)
}
func (m *PostOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostOperation.Marshal(b, m, deterministic)
}
func (m *PostOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostOperation.Merge(m, src)
}
func (m *PostOperation) XXX_Size() int {
	return xxx_messageInfo_PostOperation.Size(m)
}
func (m *PostOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_PostOperation.DiscardUnknown(m)
}

var xxx_messageInfo_PostOperation proto.InternalMessageInfo

func (m *PostOperation) GetUuid() uint64 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

func (m *PostOperation) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PostOperation) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PostOperation) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PostOperation) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PostOperation) GetBeneficiaries() []*BeneficiaryRouteType {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

type ReplyOperation struct {
	Uuid                 uint64                  `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Owner                *AccountName            `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Content              string                  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	ParentUuid           uint64                  `protobuf:"varint,4,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"`
	Beneficiaries        []*BeneficiaryRouteType `protobuf:"bytes,6,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ReplyOperation) Reset()         { *m = ReplyOperation{} }
func (m *ReplyOperation) String() string { return proto.CompactTextString(m) }
func (*ReplyOperation) ProtoMessage()    {}
func (*ReplyOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{12}
}

func (m *ReplyOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyOperation.Unmarshal(m, b)
}
func (m *ReplyOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyOperation.Marshal(b, m, deterministic)
}
func (m *ReplyOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyOperation.Merge(m, src)
}
func (m *ReplyOperation) XXX_Size() int {
	return xxx_messageInfo_ReplyOperation.Size(m)
}
func (m *ReplyOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyOperation proto.InternalMessageInfo

func (m *ReplyOperation) GetUuid() uint64 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

func (m *ReplyOperation) GetOwner() *AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ReplyOperation) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReplyOperation) GetParentUuid() uint64 {
	if m != nil {
		return m.ParentUuid
	}
	return 0
}

func (m *ReplyOperation) GetBeneficiaries() []*BeneficiaryRouteType {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

type ClaimAllOperation struct {
	Account              *AccountName `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ClaimAllOperation) Reset()         { *m = ClaimAllOperation{} }
func (m *ClaimAllOperation) String() string { return proto.CompactTextString(m) }
func (*ClaimAllOperation) ProtoMessage()    {}
func (*ClaimAllOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{13}
}

func (m *ClaimAllOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimAllOperation.Unmarshal(m, b)
}
func (m *ClaimAllOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimAllOperation.Marshal(b, m, deterministic)
}
func (m *ClaimAllOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimAllOperation.Merge(m, src)
}
func (m *ClaimAllOperation) XXX_Size() int {
	return xxx_messageInfo_ClaimAllOperation.Size(m)
}
func (m *ClaimAllOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimAllOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimAllOperation proto.InternalMessageInfo

func (m *ClaimAllOperation) GetAccount() *AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

type ClaimOperation struct {
	Account              *AccountName `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Amount               uint64       `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ClaimOperation) Reset()         { *m = ClaimOperation{} }
func (m *ClaimOperation) String() string { return proto.CompactTextString(m) }
func (*ClaimOperation) ProtoMessage()    {}
func (*ClaimOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c964c0e078f560bc, []int{14}
}

func (m *ClaimOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimOperation.Unmarshal(m, b)
}
func (m *ClaimOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimOperation.Marshal(b, m, deterministic)
}
func (m *ClaimOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimOperation.Merge(m, src)
}
func (m *ClaimOperation) XXX_Size() int {
	return xxx_messageInfo_ClaimOperation.Size(m)
}
func (m *ClaimOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimOperation.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimOperation proto.InternalMessageInfo

func (m *ClaimOperation) GetAccount() *AccountName {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *ClaimOperation) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountCreateOperation)(nil), "prototype.account_create_operation")
	proto.RegisterType((*TransferOperation)(nil), "prototype.transfer_operation")
	proto.RegisterType((*TransferToVestingOperation)(nil), "prototype.transfer_to_vesting_operation")
	proto.RegisterType((*VoteOperation)(nil), "prototype.vote_operation")
	proto.RegisterType((*BpRegisterOperation)(nil), "prototype.bp_register_operation")
	proto.RegisterType((*BpUnregisterOperation)(nil), "prototype.bp_unregister_operation")
	proto.RegisterType((*BpVoteOperation)(nil), "prototype.bp_vote_operation")
	proto.RegisterType((*FollowOperation)(nil), "prototype.follow_operation")
	proto.RegisterType((*ContractDeployOperation)(nil), "prototype.contract_deploy_operation")
	proto.RegisterType((*ContractApplyOperation)(nil), "prototype.contract_apply_operation")
	proto.RegisterType((*ContractEstimateApplyOperation)(nil), "prototype.contract_estimate_apply_operation")
	proto.RegisterType((*PostOperation)(nil), "prototype.post_operation")
	proto.RegisterType((*ReplyOperation)(nil), "prototype.reply_operation")
	proto.RegisterType((*ClaimAllOperation)(nil), "prototype.claim_all_operation")
	proto.RegisterType((*ClaimOperation)(nil), "prototype.claim_operation")
}

func init() { proto.RegisterFile("prototype/operation.proto", fileDescriptor_c964c0e078f560bc) }

var fileDescriptor_c964c0e078f560bc = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x97, 0x63, 0x27, 0x69, 0x5e, 0x77, 0xdb, 0xee, 0x50, 0xba, 0xde, 0x22, 0x44, 0x6b, 0x0e,
	0x5b, 0x2d, 0xb4, 0x51, 0x17, 0xbe, 0xc0, 0x72, 0x80, 0x95, 0x10, 0x48, 0x18, 0x71, 0x41, 0x48,
	0xd6, 0x78, 0xf2, 0xe2, 0x0c, 0x6b, 0xcf, 0x58, 0x33, 0xe3, 0x0d, 0xb9, 0x73, 0x84, 0x0b, 0xdc,
	0xb9, 0xf2, 0x11, 0xf8, 0x2a, 0x1c, 0xb9, 0xf1, 0x35, 0xd0, 0x8c, 0x27, 0xae, 0x03, 0xea, 0x46,
	0x4b, 0xf6, 0xd0, 0x4b, 0xf4, 0xfe, 0xbf, 0xdf, 0xfb, 0xcd, 0xcb, 0x33, 0x3c, 0xaa, 0x95, 0x34,
	0xd2, 0xac, 0x6a, 0x9c, 0xca, 0x1a, 0x15, 0x35, 0x5c, 0x8a, 0x2b, 0x67, 0x23, 0x93, 0xce, 0x75,
	0x7a, 0x7c, 0x13, 0x65, 0x7f, 0xda, 0x80, 0xe4, 0xc7, 0x01, 0xc4, 0x94, 0x31, 0xd9, 0x08, 0x93,
	0x31, 0x85, 0xd4, 0x60, 0xd6, 0xd5, 0x20, 0xe7, 0x10, 0xce, 0x11, 0xe3, 0xe0, 0x2c, 0xb8, 0xd8,
	0x7f, 0x7a, 0x78, 0xd5, 0x15, 0xb8, 0x62, 0x92, 0x8b, 0xd4, 0xfa, 0xc8, 0x35, 0x8c, 0x5d, 0x9a,
	0x54, 0xf1, 0xc0, 0x85, 0x3d, 0xec, 0x85, 0xad, 0x0b, 0x0b, 0x5a, 0x61, 0xba, 0x8e, 0x23, 0xcf,
	0xe0, 0x48, 0xe0, 0x32, 0xeb, 0x3b, 0xe3, 0xf0, 0xd5, 0xb9, 0x07, 0x02, 0x97, 0xcf, 0x5a, 0xc3,
	0x97, 0xb4, 0x42, 0xf2, 0x04, 0x86, 0x72, 0x29, 0x50, 0xc5, 0x91, 0xcb, 0x3b, 0xee, 0xe7, 0x35,
	0x66, 0x21, 0x15, 0x37, 0xab, 0xb4, 0x0d, 0x21, 0xef, 0xc3, 0xfd, 0xef, 0xb5, 0x14, 0x59, 0x85,
	0x86, 0xce, 0xa8, 0xa1, 0xf1, 0xde, 0x59, 0x70, 0x31, 0x49, 0xef, 0x59, 0xe3, 0x17, 0xde, 0x96,
	0xfc, 0x1e, 0x00, 0x31, 0x8a, 0x0a, 0x3d, 0x47, 0xd5, 0x23, 0xe0, 0x03, 0x88, 0xe6, 0x4a, 0x56,
	0x9e, 0x81, 0x5b, 0xe1, 0xb9, 0x20, 0xf2, 0x18, 0x06, 0x46, 0x6e, 0x63, 0x61, 0x60, 0x24, 0x79,
	0x0c, 0x23, 0x5a, 0x59, 0x93, 0x1f, 0xfb, 0x3f, 0xcc, 0x7a, 0x37, 0x21, 0x10, 0x55, 0x58, 0x49,
	0x37, 0xe5, 0x24, 0x75, 0x72, 0xf2, 0x5b, 0x00, 0xef, 0x76, 0x48, 0x8d, 0xcc, 0x5e, 0xa2, 0x36,
	0x5c, 0x14, 0x77, 0x06, 0x74, 0xf2, 0x15, 0x1c, 0xbc, 0x94, 0x1b, 0x6b, 0x74, 0x09, 0x43, 0x6b,
	0x51, 0xdb, 0x10, 0xb5, 0x51, 0xe4, 0x08, 0x42, 0x3e, 0xfb, 0xc1, 0xb5, 0x89, 0x52, 0x2b, 0x26,
	0x7f, 0x07, 0xf0, 0x76, 0x5e, 0x67, 0x0a, 0x0b, 0xae, 0xcd, 0xc6, 0x03, 0x5d, 0xae, 0x17, 0x61,
	0x5b, 0xe9, 0x76, 0x17, 0x8e, 0x20, 0x6c, 0x54, 0xe9, 0xc6, 0x9d, 0xa4, 0x56, 0xb4, 0x14, 0xcf,
	0x50, 0x33, 0xd7, 0x6d, 0x92, 0x3a, 0x99, 0x7c, 0x0a, 0x0f, 0xf2, 0x52, 0xb2, 0x17, 0x99, 0xe6,
	0x85, 0xb0, 0xdc, 0xbe, 0xc0, 0x95, 0xdf, 0xb4, 0xd3, 0x5e, 0x83, 0xba, 0xc9, 0x4b, 0xce, 0xac,
	0x33, 0xb3, 0x7a, 0x7a, 0xe8, 0x92, 0xbe, 0x6e, 0x73, 0x3e, 0xc7, 0x15, 0xb9, 0x86, 0x61, 0xad,
	0x64, 0xad, 0xe3, 0xa1, 0xcb, 0x7d, 0xa7, 0xcf, 0xd8, 0x82, 0x72, 0x91, 0x59, 0x2f, 0x2a, 0xc3,
	0x51, 0xa7, 0x6d, 0x64, 0xf2, 0x1c, 0x1e, 0xe6, 0x75, 0xd6, 0x88, 0x9d, 0x47, 0x4d, 0x7e, 0x0e,
	0xe0, 0x41, 0x5e, 0x67, 0xbb, 0x3d, 0xc5, 0x35, 0x8c, 0x97, 0xdc, 0x08, 0xd4, 0x7a, 0xeb, 0xbf,
	0xdb, 0xc7, 0x91, 0x13, 0x18, 0x31, 0x2a, 0x18, 0x96, 0x8e, 0xd2, 0xbd, 0xd4, 0x6b, 0xc9, 0xaf,
	0x01, 0x1c, 0xcd, 0x65, 0x59, 0xca, 0x65, 0x0f, 0xce, 0x35, 0x8c, 0x7d, 0x95, 0x6d, 0x80, 0xd6,
	0x71, 0xe4, 0x63, 0x98, 0xcc, 0xd7, 0xb7, 0x63, 0x1b, 0xa8, 0xbd, 0xb9, 0x3f, 0x1a, 0xb7, 0xa2,
	0xfa, 0x29, 0x80, 0x47, 0x4c, 0x0a, 0xa3, 0x28, 0x33, 0xd9, 0x0c, 0xeb, 0x52, 0xae, 0xfe, 0xff,
	0x76, 0x9d, 0xc2, 0xde, 0xba, 0x96, 0x5f, 0xb1, 0x4e, 0xb7, 0x9b, 0x47, 0x73, 0xee, 0xd7, 0xcc,
	0x8a, 0x76, 0xf3, 0x98, 0x9c, 0xa1, 0x5b, 0xac, 0x7b, 0xa9, 0x93, 0x93, 0x5f, 0x06, 0x10, 0x77,
	0x70, 0x68, 0x5d, 0x97, 0x7d, 0x34, 0x53, 0x3b, 0x43, 0x59, 0x6e, 0x87, 0xe3, 0xc3, 0x6e, 0xe0,
	0x0f, 0x5e, 0x1b, 0x7e, 0xf8, 0x2f, 0xf8, 0x27, 0x30, 0xaa, 0xd0, 0x2c, 0xe4, 0xcc, 0xdf, 0x22,
	0xaf, 0x59, 0x7b, 0x4d, 0x15, 0xad, 0xda, 0x1d, 0x9f, 0xa4, 0x5e, 0xeb, 0x5d, 0x8b, 0xd1, 0xab,
	0x4f, 0xdc, 0x39, 0x84, 0x05, 0xd5, 0xf1, 0xf8, 0x96, 0x4f, 0x4c, 0x41, 0x75, 0xf2, 0x47, 0x00,
	0xe7, 0x1d, 0x29, 0xf6, 0xda, 0x55, 0xf6, 0x2b, 0x75, 0xc7, 0xd8, 0xf1, 0x2c, 0x44, 0x7d, 0x16,
	0x92, 0xbf, 0x02, 0x38, 0xa8, 0xa5, 0x36, 0x3d, 0x98, 0x04, 0xa2, 0xa6, 0xe1, 0x33, 0x07, 0x32,
	0x4a, 0x9d, 0xfc, 0xba, 0x48, 0x8e, 0x61, 0x68, 0xb8, 0x29, 0xd1, 0xc3, 0x68, 0x15, 0x12, 0xc3,
	0xd8, 0xe2, 0x41, 0x61, 0x3c, 0x88, 0xb5, 0x6a, 0x5b, 0x1a, 0x5a, 0xd8, 0x17, 0x0a, 0xed, 0x89,
	0xb3, 0x32, 0xf9, 0x0c, 0xee, 0xe7, 0x28, 0x70, 0xce, 0x19, 0xa7, 0x8a, 0xa3, 0x8e, 0x47, 0x67,
	0xe1, 0xc5, 0xfe, 0xd3, 0xf3, 0x5e, 0xeb, 0x1b, 0xff, 0x2a, 0x53, 0xb2, 0x31, 0xd8, 0x5e, 0xb9,
	0xcd, 0xbc, 0xe4, 0xcf, 0x00, 0x0e, 0x15, 0x6e, 0x3e, 0xc5, 0x1b, 0x98, 0xb1, 0x37, 0x4d, 0xb8,
	0x39, 0xcd, 0x7b, 0xb0, 0x5f, 0x53, 0x85, 0xc2, 0x64, 0xae, 0x47, 0xe4, 0x7a, 0x40, 0x6b, 0xfa,
	0xc6, 0x76, 0x7a, 0x63, 0xa3, 0x3d, 0x87, 0xb7, 0x58, 0x49, 0x79, 0x95, 0xd1, 0xb2, 0xdc, 0xe9,
	0x66, 0x25, 0xdf, 0xc1, 0x61, 0x5b, 0x69, 0xa7, 0xcb, 0x77, 0xd2, 0xfd, 0xa7, 0x06, 0x6e, 0x68,
	0xaf, 0x7d, 0xf2, 0xe1, 0xb7, 0x4f, 0x0a, 0x6e, 0x16, 0x4d, 0x7e, 0xc5, 0x64, 0x35, 0x65, 0x52,
	0xbb, 0x8f, 0xcb, 0xd4, 0x13, 0x26, 0xf5, 0x65, 0x21, 0xa7, 0x5d, 0xed, 0x7c, 0xe4, 0xc4, 0x8f,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x5e, 0x2c, 0x4f, 0x35, 0x0a, 0x00, 0x00,
}
