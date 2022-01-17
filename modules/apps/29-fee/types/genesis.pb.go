// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fee/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the fee middleware genesis state
type GenesisState struct {
	IdentifiedFees     []*IdentifiedPacketFee      `protobuf:"bytes,1,rep,name=identified_fees,json=identifiedFees,proto3" json:"identified_fees,omitempty" yaml:"identified_fees"`
	FeeEnabledChannels []*FeeEnabledChannel        `protobuf:"bytes,2,rep,name=fee_enabled_channels,json=feeEnabledChannels,proto3" json:"fee_enabled_channels,omitempty" yaml:"fee_enabled_channels"`
	RegisteredRelayers []*RegisteredRelayerAddress `protobuf:"bytes,3,rep,name=registered_relayers,json=registeredRelayers,proto3" json:"registered_relayers,omitempty" yaml:"registered_relayers"`
	ForwardRelayers    []*ForwardRelayerAddress    `protobuf:"bytes,4,rep,name=forward_relayers,json=forwardRelayers,proto3" json:"forward_relayers,omitempty" yaml:"forward_relayers"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7191992e856dff95, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetIdentifiedFees() []*IdentifiedPacketFee {
	if m != nil {
		return m.IdentifiedFees
	}
	return nil
}

func (m *GenesisState) GetFeeEnabledChannels() []*FeeEnabledChannel {
	if m != nil {
		return m.FeeEnabledChannels
	}
	return nil
}

func (m *GenesisState) GetRegisteredRelayers() []*RegisteredRelayerAddress {
	if m != nil {
		return m.RegisteredRelayers
	}
	return nil
}

func (m *GenesisState) GetForwardRelayers() []*ForwardRelayerAddress {
	if m != nil {
		return m.ForwardRelayers
	}
	return nil
}

// FeeEnabledChannel contains the PortID & ChannelID for a fee enabled channel
type FeeEnabledChannel struct {
	PortId    string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty" yaml:"port_id"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
}

func (m *FeeEnabledChannel) Reset()         { *m = FeeEnabledChannel{} }
func (m *FeeEnabledChannel) String() string { return proto.CompactTextString(m) }
func (*FeeEnabledChannel) ProtoMessage()    {}
func (*FeeEnabledChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_7191992e856dff95, []int{1}
}
func (m *FeeEnabledChannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeEnabledChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeEnabledChannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeEnabledChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeEnabledChannel.Merge(m, src)
}
func (m *FeeEnabledChannel) XXX_Size() int {
	return m.Size()
}
func (m *FeeEnabledChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeEnabledChannel.DiscardUnknown(m)
}

var xxx_messageInfo_FeeEnabledChannel proto.InternalMessageInfo

func (m *FeeEnabledChannel) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *FeeEnabledChannel) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// RegisteredRelayerAddress contains the address and counterparty address for a specific relayer (for distributing fees)
type RegisteredRelayerAddress struct {
	Address             string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CounterpartyAddress string `protobuf:"bytes,2,opt,name=counterparty_address,json=counterpartyAddress,proto3" json:"counterparty_address,omitempty" yaml:"counterparty_address"`
}

func (m *RegisteredRelayerAddress) Reset()         { *m = RegisteredRelayerAddress{} }
func (m *RegisteredRelayerAddress) String() string { return proto.CompactTextString(m) }
func (*RegisteredRelayerAddress) ProtoMessage()    {}
func (*RegisteredRelayerAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7191992e856dff95, []int{2}
}
func (m *RegisteredRelayerAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredRelayerAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredRelayerAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredRelayerAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredRelayerAddress.Merge(m, src)
}
func (m *RegisteredRelayerAddress) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredRelayerAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredRelayerAddress.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredRelayerAddress proto.InternalMessageInfo

func (m *RegisteredRelayerAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegisteredRelayerAddress) GetCounterpartyAddress() string {
	if m != nil {
		return m.CounterpartyAddress
	}
	return ""
}

// ForwardRelayerAddress contains the forward relayer address and packetId used for async acknowledgements
type ForwardRelayerAddress struct {
	Address  string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PacketId types.PacketId `protobuf:"bytes,2,opt,name=packet_id,json=packetId,proto3" json:"packet_id" yaml:"packet_id"`
}

func (m *ForwardRelayerAddress) Reset()         { *m = ForwardRelayerAddress{} }
func (m *ForwardRelayerAddress) String() string { return proto.CompactTextString(m) }
func (*ForwardRelayerAddress) ProtoMessage()    {}
func (*ForwardRelayerAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7191992e856dff95, []int{3}
}
func (m *ForwardRelayerAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForwardRelayerAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForwardRelayerAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForwardRelayerAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardRelayerAddress.Merge(m, src)
}
func (m *ForwardRelayerAddress) XXX_Size() int {
	return m.Size()
}
func (m *ForwardRelayerAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardRelayerAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardRelayerAddress proto.InternalMessageInfo

func (m *ForwardRelayerAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ForwardRelayerAddress) GetPacketId() types.PacketId {
	if m != nil {
		return m.PacketId
	}
	return types.PacketId{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.applications.fee.v1.GenesisState")
	proto.RegisterType((*FeeEnabledChannel)(nil), "ibc.applications.fee.v1.FeeEnabledChannel")
	proto.RegisterType((*RegisteredRelayerAddress)(nil), "ibc.applications.fee.v1.RegisteredRelayerAddress")
	proto.RegisterType((*ForwardRelayerAddress)(nil), "ibc.applications.fee.v1.ForwardRelayerAddress")
}

func init() {
	proto.RegisterFile("ibc/applications/fee/v1/genesis.proto", fileDescriptor_7191992e856dff95)
}

var fileDescriptor_7191992e856dff95 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xbf, 0x6f, 0xd4, 0x30,
	0x18, 0xbd, 0xb4, 0xa8, 0xa5, 0x2e, 0xea, 0x0f, 0xb7, 0xa5, 0xd1, 0x55, 0x24, 0xc5, 0x12, 0x52,
	0x05, 0x34, 0xd1, 0xb5, 0x30, 0xc0, 0xc6, 0x21, 0x8a, 0x6e, 0x02, 0x19, 0x26, 0x96, 0x53, 0x2e,
	0xf9, 0x92, 0x5a, 0xe4, 0xe2, 0x60, 0xbb, 0x41, 0x37, 0xb0, 0xb0, 0xc0, 0xc8, 0x9f, 0xd5, 0xb1,
	0x23, 0x53, 0x84, 0xda, 0xff, 0x20, 0x3b, 0x12, 0x4a, 0x9c, 0xb4, 0xc7, 0x71, 0x61, 0xfb, 0x62,
	0xbf, 0xf7, 0xbd, 0xe7, 0xe7, 0x7c, 0x46, 0x0f, 0xd8, 0xc8, 0x77, 0xbd, 0x34, 0x8d, 0x99, 0xef,
	0x29, 0xc6, 0x13, 0xe9, 0x86, 0x00, 0x6e, 0xd6, 0x73, 0x23, 0x48, 0x40, 0x32, 0xe9, 0xa4, 0x82,
	0x2b, 0x8e, 0x77, 0xd9, 0xc8, 0x77, 0xa6, 0x61, 0x4e, 0x08, 0xe0, 0x64, 0xbd, 0xee, 0x76, 0xc4,
	0x23, 0x5e, 0x61, 0xdc, 0xb2, 0xd2, 0xf0, 0xee, 0xfd, 0xb6, 0xae, 0x25, 0x6b, 0x0a, 0xe2, 0x73,
	0x01, 0xae, 0x7f, 0xea, 0x25, 0x09, 0xc4, 0xe5, 0x76, 0x5d, 0x6a, 0x08, 0xf9, 0xbd, 0x88, 0xee,
	0xbc, 0xd6, 0x36, 0xde, 0x29, 0x4f, 0x01, 0xfe, 0x84, 0xd6, 0x59, 0x00, 0x89, 0x62, 0x21, 0x83,
	0x60, 0x18, 0x02, 0x48, 0xd3, 0xd8, 0x5f, 0x3c, 0x58, 0x3d, 0x7a, 0xec, 0xb4, 0xf8, 0x73, 0x06,
	0xd7, 0xf8, 0xb7, 0x9e, 0xff, 0x11, 0xd4, 0x09, 0x40, 0xbf, 0x5b, 0xe4, 0xf6, 0xdd, 0x89, 0x37,
	0x8e, 0x9f, 0x93, 0x99, 0x76, 0x84, 0xae, 0xdd, 0xac, 0x9c, 0x00, 0x48, 0xfc, 0x05, 0x6d, 0x87,
	0x00, 0x43, 0x48, 0xbc, 0x51, 0x0c, 0xc1, 0xb0, 0x36, 0x28, 0xcd, 0x85, 0x4a, 0xf7, 0x61, 0xab,
	0xee, 0x09, 0xc0, 0x2b, 0xcd, 0x79, 0xa9, 0x29, 0x7d, 0xbb, 0xc8, 0xed, 0x3d, 0xad, 0x3a, 0xaf,
	0x23, 0xa1, 0x38, 0x9c, 0xe5, 0x48, 0xfc, 0xd5, 0x40, 0x5b, 0x02, 0x22, 0x26, 0x15, 0x08, 0x08,
	0x86, 0x02, 0x62, 0x6f, 0x02, 0x42, 0x9a, 0x8b, 0x95, 0x7c, 0xaf, 0x55, 0x9e, 0x5e, 0x73, 0xa8,
	0xa6, 0xbc, 0x08, 0x02, 0x01, 0x52, 0xf6, 0xad, 0x22, 0xb7, 0xbb, 0xda, 0xc5, 0x9c, 0xbe, 0x84,
	0x62, 0x31, 0xcb, 0x94, 0x38, 0x43, 0x1b, 0x21, 0x17, 0x9f, 0x3d, 0x31, 0x65, 0xe0, 0x56, 0x65,
	0xc0, 0x69, 0x3f, 0xbf, 0x26, 0xcc, 0xa8, 0xef, 0x15, 0xb9, 0xbd, 0x5b, 0x67, 0x30, 0xd3, 0x91,
	0xd0, 0xf5, 0xf0, 0x2f, 0x8e, 0x24, 0x19, 0xda, 0xfc, 0x27, 0x46, 0xfc, 0x08, 0x2d, 0xa7, 0x5c,
	0xa8, 0x21, 0x0b, 0x4c, 0x63, 0xdf, 0x38, 0x58, 0xe9, 0xe3, 0x22, 0xb7, 0xd7, 0x74, 0xcf, 0x7a,
	0x83, 0xd0, 0xa5, 0xb2, 0x1a, 0x04, 0xf8, 0x09, 0x42, 0x75, 0xbe, 0x25, 0x7e, 0xa1, 0xc2, 0xef,
	0x14, 0xb9, 0xbd, 0xa9, 0xf1, 0x37, 0x7b, 0x84, 0xae, 0xd4, 0x1f, 0x83, 0x80, 0x7c, 0x37, 0x90,
	0xd9, 0x16, 0x20, 0x36, 0xd1, 0xb2, 0xa7, 0x4b, 0xad, 0x4f, 0x9b, 0x4f, 0x4c, 0xd1, 0xb6, 0xcf,
	0xcf, 0x12, 0x05, 0x22, 0xf5, 0x84, 0x9a, 0x0c, 0x1b, 0x98, 0x96, 0x9d, 0xba, 0xfe, 0x79, 0x28,
	0x42, 0xb7, 0xa6, 0x97, 0x6b, 0x35, 0xf2, 0xcd, 0x40, 0x3b, 0x73, 0xa3, 0xfc, 0x8f, 0x8f, 0xf7,
	0x68, 0x25, 0xad, 0xfe, 0xf5, 0xe6, 0xcc, 0xab, 0x47, 0xf7, 0xaa, 0x7b, 0x2a, 0xa7, 0xcd, 0x69,
	0x46, 0x2c, 0xeb, 0x39, 0x7a, 0x22, 0x06, 0x41, 0xdf, 0x3c, 0xcf, 0xed, 0x4e, 0x91, 0xdb, 0x1b,
	0x75, 0x8c, 0x0d, 0x9b, 0xd0, 0xdb, 0x69, 0x83, 0x79, 0x73, 0x7e, 0x69, 0x19, 0x17, 0x97, 0x96,
	0xf1, 0xeb, 0xd2, 0x32, 0x7e, 0x5c, 0x59, 0x9d, 0x8b, 0x2b, 0xab, 0xf3, 0xf3, 0xca, 0xea, 0x7c,
	0x78, 0x1a, 0x31, 0x75, 0x7a, 0x36, 0x72, 0x7c, 0x3e, 0x76, 0x7d, 0x2e, 0xc7, 0x5c, 0xba, 0x6c,
	0xe4, 0x1f, 0x46, 0xdc, 0xcd, 0x8e, 0xdd, 0x31, 0x0f, 0xce, 0x62, 0x90, 0xe5, 0x63, 0x20, 0xdd,
	0xa3, 0x67, 0x87, 0xe5, 0x3b, 0xa0, 0x26, 0x29, 0xc8, 0xd1, 0x52, 0x35, 0xe4, 0xc7, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x5f, 0x63, 0x48, 0xc9, 0x82, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ForwardRelayers) > 0 {
		for iNdEx := len(m.ForwardRelayers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ForwardRelayers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RegisteredRelayers) > 0 {
		for iNdEx := len(m.RegisteredRelayers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisteredRelayers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FeeEnabledChannels) > 0 {
		for iNdEx := len(m.FeeEnabledChannels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeEnabledChannels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IdentifiedFees) > 0 {
		for iNdEx := len(m.IdentifiedFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IdentifiedFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FeeEnabledChannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeEnabledChannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeEnabledChannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisteredRelayerAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredRelayerAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredRelayerAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CounterpartyAddress) > 0 {
		i -= len(m.CounterpartyAddress)
		copy(dAtA[i:], m.CounterpartyAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CounterpartyAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ForwardRelayerAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardRelayerAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForwardRelayerAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PacketId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IdentifiedFees) > 0 {
		for _, e := range m.IdentifiedFees {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FeeEnabledChannels) > 0 {
		for _, e := range m.FeeEnabledChannels {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RegisteredRelayers) > 0 {
		for _, e := range m.RegisteredRelayers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ForwardRelayers) > 0 {
		for _, e := range m.ForwardRelayers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *FeeEnabledChannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *RegisteredRelayerAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.CounterpartyAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *ForwardRelayerAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.PacketId.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentifiedFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentifiedFees = append(m.IdentifiedFees, &IdentifiedPacketFee{})
			if err := m.IdentifiedFees[len(m.IdentifiedFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeEnabledChannels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeEnabledChannels = append(m.FeeEnabledChannels, &FeeEnabledChannel{})
			if err := m.FeeEnabledChannels[len(m.FeeEnabledChannels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredRelayers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisteredRelayers = append(m.RegisteredRelayers, &RegisteredRelayerAddress{})
			if err := m.RegisteredRelayers[len(m.RegisteredRelayers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwardRelayers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForwardRelayers = append(m.ForwardRelayers, &ForwardRelayerAddress{})
			if err := m.ForwardRelayers[len(m.ForwardRelayers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *FeeEnabledChannel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: FeeEnabledChannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeEnabledChannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *RegisteredRelayerAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: RegisteredRelayerAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredRelayerAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CounterpartyAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ForwardRelayerAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ForwardRelayerAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardRelayerAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
