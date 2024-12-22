// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/cdp/v1beta1/cdp.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CDP defines the state of a single collateralized debt position.
type CDP struct {
	ID              uint64                                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner           github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty"`
	Type            string                                        `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Collateral      types.Coin                                    `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral"`
	Principal       types.Coin                                    `protobuf:"bytes,5,opt,name=principal,proto3" json:"principal"`
	AccumulatedFees types.Coin                                    `protobuf:"bytes,6,opt,name=accumulated_fees,json=accumulatedFees,proto3" json:"accumulated_fees"`
	FeesUpdated     time.Time                                     `protobuf:"bytes,7,opt,name=fees_updated,json=feesUpdated,proto3,stdtime" json:"fees_updated"`
	InterestFactor  github_com_cosmos_cosmos_sdk_types.Dec        `protobuf:"bytes,8,opt,name=interest_factor,json=interestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"interest_factor"`
}

func (m *CDP) Reset()         { *m = CDP{} }
func (m *CDP) String() string { return proto.CompactTextString(m) }
func (*CDP) ProtoMessage()    {}
func (*CDP) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a9ab097fb7be40, []int{0}
}
func (m *CDP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CDP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CDP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CDP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDP.Merge(m, src)
}
func (m *CDP) XXX_Size() int {
	return m.Size()
}
func (m *CDP) XXX_DiscardUnknown() {
	xxx_messageInfo_CDP.DiscardUnknown(m)
}

var xxx_messageInfo_CDP proto.InternalMessageInfo

// Deposit defines an amount of coins deposited by an account to a cdp
type Deposit struct {
	CdpID     uint64                                        `protobuf:"varint,1,opt,name=cdp_id,json=cdpId,proto3" json:"cdp_id,omitempty"`
	Depositor github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=depositor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor,omitempty"`
	Amount    types.Coin                                    `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a9ab097fb7be40, []int{1}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

// TotalPrincipal defines the total principal of a given collateral type
type TotalPrincipal struct {
	CollateralType string     `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	Amount         types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *TotalPrincipal) Reset()         { *m = TotalPrincipal{} }
func (m *TotalPrincipal) String() string { return proto.CompactTextString(m) }
func (*TotalPrincipal) ProtoMessage()    {}
func (*TotalPrincipal) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a9ab097fb7be40, []int{2}
}
func (m *TotalPrincipal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalPrincipal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalPrincipal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalPrincipal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalPrincipal.Merge(m, src)
}
func (m *TotalPrincipal) XXX_Size() int {
	return m.Size()
}
func (m *TotalPrincipal) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalPrincipal.DiscardUnknown(m)
}

var xxx_messageInfo_TotalPrincipal proto.InternalMessageInfo

// TotalCollateral defines the total collateral of a given collateral type
type TotalCollateral struct {
	CollateralType string     `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	Amount         types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *TotalCollateral) Reset()         { *m = TotalCollateral{} }
func (m *TotalCollateral) String() string { return proto.CompactTextString(m) }
func (*TotalCollateral) ProtoMessage()    {}
func (*TotalCollateral) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a9ab097fb7be40, []int{3}
}
func (m *TotalCollateral) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalCollateral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalCollateral.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalCollateral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalCollateral.Merge(m, src)
}
func (m *TotalCollateral) XXX_Size() int {
	return m.Size()
}
func (m *TotalCollateral) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalCollateral.DiscardUnknown(m)
}

var xxx_messageInfo_TotalCollateral proto.InternalMessageInfo

// OwnerCDPIndex defines the cdp ids for a single cdp owner
type OwnerCDPIndex struct {
	CdpIDs []uint64 `protobuf:"varint,1,rep,packed,name=cdp_ids,json=cdpIds,proto3" json:"cdp_ids,omitempty"`
}

func (m *OwnerCDPIndex) Reset()         { *m = OwnerCDPIndex{} }
func (m *OwnerCDPIndex) String() string { return proto.CompactTextString(m) }
func (*OwnerCDPIndex) ProtoMessage()    {}
func (*OwnerCDPIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a9ab097fb7be40, []int{4}
}
func (m *OwnerCDPIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnerCDPIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnerCDPIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnerCDPIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerCDPIndex.Merge(m, src)
}
func (m *OwnerCDPIndex) XXX_Size() int {
	return m.Size()
}
func (m *OwnerCDPIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerCDPIndex.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerCDPIndex proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CDP)(nil), "kava.cdp.v1beta1.CDP")
	proto.RegisterType((*Deposit)(nil), "kava.cdp.v1beta1.Deposit")
	proto.RegisterType((*TotalPrincipal)(nil), "kava.cdp.v1beta1.TotalPrincipal")
	proto.RegisterType((*TotalCollateral)(nil), "kava.cdp.v1beta1.TotalCollateral")
	proto.RegisterType((*OwnerCDPIndex)(nil), "kava.cdp.v1beta1.OwnerCDPIndex")
}

func init() { proto.RegisterFile("kava/cdp/v1beta1/cdp.proto", fileDescriptor_68a9ab097fb7be40) }

var fileDescriptor_68a9ab097fb7be40 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xae, 0xbb, 0x36, 0x5b, 0xbd, 0xb1, 0x4e, 0x06, 0xa1, 0xac, 0x87, 0xa4, 0x1a, 0x02, 0x7a,
	0x69, 0xa2, 0x01, 0x12, 0x17, 0x10, 0x5a, 0x1a, 0x0d, 0xca, 0x85, 0x29, 0x1a, 0x17, 0x0e, 0x54,
	0xae, 0xed, 0x96, 0x68, 0x69, 0x1c, 0xc5, 0xee, 0xd8, 0x1e, 0x02, 0x69, 0x0f, 0xb3, 0x87, 0xd8,
	0x81, 0xc3, 0xb4, 0x13, 0xe2, 0x10, 0x20, 0x7b, 0x0b, 0x4e, 0xc8, 0x4e, 0xba, 0xec, 0x58, 0x24,
	0x38, 0xf5, 0xf7, 0xc7, 0xdf, 0xf7, 0x73, 0x7f, 0xdf, 0x17, 0xc3, 0xce, 0x11, 0x3e, 0xc6, 0x2e,
	0xa1, 0x89, 0x7b, 0xbc, 0x3b, 0x66, 0x12, 0xef, 0xaa, 0xd8, 0x49, 0x52, 0x2e, 0x39, 0xda, 0x52,
	0x3d, 0x47, 0xe5, 0x65, 0xaf, 0x63, 0x11, 0x2e, 0x66, 0x5c, 0xb8, 0x63, 0x2c, 0x58, 0x05, 0xe0,
	0x61, 0x5c, 0x20, 0x3a, 0xdb, 0x45, 0x7f, 0xa4, 0x33, 0xb7, 0x48, 0xca, 0xd6, 0xbd, 0x29, 0x9f,
	0xf2, 0xa2, 0xae, 0xa2, 0xb2, 0x6a, 0x4f, 0x39, 0x9f, 0x46, 0xcc, 0xd5, 0xd9, 0x78, 0x3e, 0x71,
	0x65, 0x38, 0x63, 0x42, 0xe2, 0x59, 0x79, 0x87, 0x9d, 0x2f, 0x0d, 0xb8, 0x32, 0xf0, 0x0f, 0xd0,
	0x7d, 0x58, 0x0f, 0xa9, 0x09, 0xba, 0xa0, 0xd7, 0xf0, 0x8c, 0x3c, 0xb3, 0xeb, 0x43, 0x3f, 0xa8,
	0x87, 0x14, 0x7d, 0x84, 0x4d, 0xfe, 0x39, 0x66, 0xa9, 0x59, 0xef, 0x82, 0xde, 0x86, 0xf7, 0xe6,
	0x77, 0x66, 0xf7, 0xa7, 0xa1, 0xfc, 0x34, 0x1f, 0x3b, 0x84, 0xcf, 0xca, 0x2b, 0x94, 0x3f, 0x7d,
	0x41, 0x8f, 0x5c, 0x79, 0x9a, 0x30, 0xe1, 0xec, 0x11, 0xb2, 0x47, 0x69, 0xca, 0x84, 0xb8, 0x3a,
	0xef, 0xdf, 0x2d, 0x2f, 0x5a, 0x56, 0xbc, 0x53, 0xc9, 0x44, 0x50, 0xd0, 0x22, 0x04, 0x1b, 0x0a,
	0x61, 0xae, 0x74, 0x41, 0xaf, 0x15, 0xe8, 0x18, 0xbd, 0x82, 0x90, 0xf0, 0x28, 0xc2, 0x92, 0xa5,
	0x38, 0x32, 0x1b, 0x5d, 0xd0, 0x5b, 0x7f, 0xb2, 0xed, 0x94, 0x24, 0x6a, 0x35, 0x8b, 0x7d, 0x39,
	0x03, 0x1e, 0xc6, 0x5e, 0xe3, 0x22, 0xb3, 0x6b, 0xc1, 0x2d, 0x08, 0x7a, 0x09, 0x5b, 0x49, 0x1a,
	0xc6, 0x24, 0x4c, 0x70, 0x64, 0x36, 0x97, 0xc3, 0x57, 0x08, 0xf4, 0x16, 0x6e, 0x61, 0x42, 0xe6,
	0xb3, 0xb9, 0xe2, 0xa3, 0xa3, 0x09, 0x63, 0xc2, 0x34, 0x96, 0x63, 0x69, 0xdf, 0x02, 0xee, 0x33,
	0x26, 0xd0, 0x6b, 0xb8, 0xa1, 0xf0, 0xa3, 0x79, 0x42, 0x55, 0xcd, 0x5c, 0xd5, 0x3c, 0x1d, 0xa7,
	0xd0, 0xc5, 0x59, 0xe8, 0xe2, 0x1c, 0x2e, 0x74, 0xf1, 0xd6, 0x14, 0xd1, 0xd9, 0x0f, 0x1b, 0x04,
	0xeb, 0x0a, 0xf9, 0xbe, 0x00, 0x22, 0x06, 0xdb, 0x61, 0x2c, 0x59, 0xca, 0x84, 0x1c, 0x4d, 0x30,
	0x91, 0x3c, 0x35, 0xd7, 0xd4, 0xce, 0xbc, 0x17, 0xea, 0xfc, 0xf7, 0xcc, 0x7e, 0xb4, 0x84, 0x2c,
	0x3e, 0x23, 0x57, 0xe7, 0x7d, 0x58, 0xfe, 0x09, 0x9f, 0x91, 0x60, 0x73, 0x41, 0xba, 0xaf, 0x39,
	0x77, 0xbe, 0x02, 0xb8, 0xea, 0xb3, 0x84, 0x8b, 0x50, 0xa2, 0x2e, 0x34, 0x08, 0x4d, 0x46, 0x37,
	0xbe, 0x68, 0xe5, 0x99, 0xdd, 0x1c, 0xd0, 0x64, 0xe8, 0x07, 0x4d, 0x42, 0x93, 0x21, 0x45, 0x13,
	0xd8, 0xa2, 0xc5, 0x61, 0x5e, 0x38, 0xa4, 0xf5, 0x0f, 0x1d, 0x52, 0x51, 0xa3, 0xe7, 0xd0, 0xc0,
	0x33, 0x3e, 0x8f, 0xa5, 0xf6, 0xc9, 0x12, 0x3a, 0x94, 0xc7, 0x77, 0x52, 0xb8, 0x79, 0xc8, 0x25,
	0x8e, 0x0e, 0x6e, 0xc4, 0x7d, 0x0c, 0xdb, 0x95, 0x53, 0x46, 0xda, 0x7b, 0x40, 0x7b, 0x6f, 0xb3,
	0x2a, 0x1f, 0x2a, 0x17, 0x56, 0x33, 0xeb, 0x7f, 0x37, 0x53, 0xc0, 0xb6, 0x9e, 0x39, 0xa8, 0x0c,
	0xf9, 0xff, 0x87, 0x3e, 0x83, 0x77, 0xde, 0xa9, 0x0f, 0x6a, 0xe0, 0x1f, 0x0c, 0x63, 0xca, 0x4e,
	0xd0, 0x03, 0xb8, 0x5a, 0x88, 0x27, 0x4c, 0xd0, 0x5d, 0xe9, 0x35, 0x3c, 0x98, 0x67, 0xb6, 0xa1,
	0xd5, 0x13, 0x81, 0xa1, 0xe5, 0x13, 0xde, 0xe0, 0xe2, 0x97, 0x55, 0xbb, 0xc8, 0x2d, 0x70, 0x99,
	0x5b, 0xe0, 0x67, 0x6e, 0x81, 0xb3, 0x6b, 0xab, 0x76, 0x79, 0x6d, 0xd5, 0xbe, 0x5d, 0x5b, 0xb5,
	0x0f, 0x0f, 0x6f, 0xc9, 0xa8, 0x9e, 0xaa, 0x7e, 0x84, 0xc7, 0x42, 0x47, 0xee, 0x89, 0x7e, 0xd2,
	0xb4, 0x92, 0x63, 0x43, 0x9b, 0xf8, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xe3, 0x65,
	0xc4, 0xeb, 0x04, 0x00, 0x00,
}

func (m *CDP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CDP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CDP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InterestFactor.Size()
		i -= size
		if _, err := m.InterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.FeesUpdated, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FeesUpdated):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCdp(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.AccumulatedFees.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Principal.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Collateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintCdp(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.CdpID != 0 {
		i = encodeVarintCdp(dAtA, i, uint64(m.CdpID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TotalPrincipal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalPrincipal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalPrincipal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TotalCollateral) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalCollateral) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalCollateral) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OwnerCDPIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnerCDPIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnerCDPIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CdpIDs) > 0 {
		dAtA9 := make([]byte, len(m.CdpIDs)*10)
		var j8 int
		for _, num := range m.CdpIDs {
			for num >= 1<<7 {
				dAtA9[j8] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j8++
			}
			dAtA9[j8] = uint8(num)
			j8++
		}
		i -= j8
		copy(dAtA[i:], dAtA9[:j8])
		i = encodeVarintCdp(dAtA, i, uint64(j8))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCdp(dAtA []byte, offset int, v uint64) int {
	offset -= sovCdp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CDP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovCdp(uint64(m.ID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = m.Collateral.Size()
	n += 1 + l + sovCdp(uint64(l))
	l = m.Principal.Size()
	n += 1 + l + sovCdp(uint64(l))
	l = m.AccumulatedFees.Size()
	n += 1 + l + sovCdp(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FeesUpdated)
	n += 1 + l + sovCdp(uint64(l))
	l = m.InterestFactor.Size()
	n += 1 + l + sovCdp(uint64(l))
	return n
}

func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CdpID != 0 {
		n += 1 + sovCdp(uint64(m.CdpID))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovCdp(uint64(l))
	return n
}

func (m *TotalPrincipal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovCdp(uint64(l))
	return n
}

func (m *TotalCollateral) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovCdp(uint64(l))
	return n
}

func (m *OwnerCDPIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CdpIDs) > 0 {
		l = 0
		for _, e := range m.CdpIDs {
			l += sovCdp(uint64(e))
		}
		n += 1 + sovCdp(uint64(l)) + l
	}
	return n
}

func sovCdp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCdp(x uint64) (n int) {
	return sovCdp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CDP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CDP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CDP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Principal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Principal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumulatedFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccumulatedFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeesUpdated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.FeesUpdated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdpID", wireType)
			}
			m.CdpID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CdpID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TotalPrincipal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TotalPrincipal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalPrincipal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TotalCollateral) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TotalCollateral: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalCollateral: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OwnerCDPIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OwnerCDPIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnerCDPIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCdp
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CdpIDs = append(m.CdpIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCdp
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCdp
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCdp
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CdpIDs) == 0 {
					m.CdpIDs = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCdp
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CdpIDs = append(m.CdpIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CdpIDs", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCdp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCdp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCdp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCdp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCdp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCdp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCdp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCdp = fmt.Errorf("proto: unexpected end of group")
)
