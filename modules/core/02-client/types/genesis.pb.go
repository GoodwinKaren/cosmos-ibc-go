// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/client/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the ibc client submodule's genesis state.
type GenesisState struct {
	// client states with their corresponding identifiers
	Clients IdentifiedClientStates `protobuf:"bytes,1,rep,name=clients,proto3,castrepeated=IdentifiedClientStates" json:"clients"`
	// consensus states from each client
	ClientsConsensus ClientsConsensusStates `protobuf:"bytes,2,rep,name=clients_consensus,json=clientsConsensus,proto3,castrepeated=ClientsConsensusStates" json:"clients_consensus" yaml:"clients_consensus"`
	// metadata from each client
	ClientsMetadata []IdentifiedGenesisMetadata `protobuf:"bytes,3,rep,name=clients_metadata,json=clientsMetadata,proto3" json:"clients_metadata" yaml:"clients_metadata"`
	Params          Params                      `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
	// create localhost on initialization
	CreateLocalhost bool `protobuf:"varint,5,opt,name=create_localhost,json=createLocalhost,proto3" json:"create_localhost,omitempty" yaml:"create_localhost"`
	// the sequence for the next generated client identifier
	NextClientSequence uint64 `protobuf:"varint,6,opt,name=next_client_sequence,json=nextClientSequence,proto3" json:"next_client_sequence,omitempty" yaml:"next_client_sequence"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcd0c0f1f2e6a91a, []int{0}
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

func (m *GenesisState) GetClients() IdentifiedClientStates {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *GenesisState) GetClientsConsensus() ClientsConsensusStates {
	if m != nil {
		return m.ClientsConsensus
	}
	return nil
}

func (m *GenesisState) GetClientsMetadata() []IdentifiedGenesisMetadata {
	if m != nil {
		return m.ClientsMetadata
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetCreateLocalhost() bool {
	if m != nil {
		return m.CreateLocalhost
	}
	return false
}

func (m *GenesisState) GetNextClientSequence() uint64 {
	if m != nil {
		return m.NextClientSequence
	}
	return 0
}

// GenesisMetadata defines the genesis type for metadata that clients may return
// with ExportMetadata
type GenesisMetadata struct {
	// store key of metadata without clientID-prefix
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// metadata value
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GenesisMetadata) Reset()         { *m = GenesisMetadata{} }
func (m *GenesisMetadata) String() string { return proto.CompactTextString(m) }
func (*GenesisMetadata) ProtoMessage()    {}
func (*GenesisMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcd0c0f1f2e6a91a, []int{1}
}
func (m *GenesisMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisMetadata.Merge(m, src)
}
func (m *GenesisMetadata) XXX_Size() int {
	return m.Size()
}
func (m *GenesisMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisMetadata proto.InternalMessageInfo

// IdentifiedGenesisMetadata has the client metadata with the corresponding
// client id.
type IdentifiedGenesisMetadata struct {
	ClientId       string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	ClientMetadata []GenesisMetadata `protobuf:"bytes,2,rep,name=client_metadata,json=clientMetadata,proto3" json:"client_metadata" yaml:"client_metadata"`
}

func (m *IdentifiedGenesisMetadata) Reset()         { *m = IdentifiedGenesisMetadata{} }
func (m *IdentifiedGenesisMetadata) String() string { return proto.CompactTextString(m) }
func (*IdentifiedGenesisMetadata) ProtoMessage()    {}
func (*IdentifiedGenesisMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcd0c0f1f2e6a91a, []int{2}
}
func (m *IdentifiedGenesisMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedGenesisMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedGenesisMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedGenesisMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedGenesisMetadata.Merge(m, src)
}
func (m *IdentifiedGenesisMetadata) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedGenesisMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedGenesisMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedGenesisMetadata proto.InternalMessageInfo

func (m *IdentifiedGenesisMetadata) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *IdentifiedGenesisMetadata) GetClientMetadata() []GenesisMetadata {
	if m != nil {
		return m.ClientMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.client.v1.GenesisState")
	proto.RegisterType((*GenesisMetadata)(nil), "ibc.core.client.v1.GenesisMetadata")
	proto.RegisterType((*IdentifiedGenesisMetadata)(nil), "ibc.core.client.v1.IdentifiedGenesisMetadata")
}

func init() { proto.RegisterFile("ibc/core/client/v1/genesis.proto", fileDescriptor_bcd0c0f1f2e6a91a) }

var fileDescriptor_bcd0c0f1f2e6a91a = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0xcd, 0x34, 0x69, 0x68, 0xa7, 0x15, 0x0d, 0xa3, 0xa8, 0x98, 0x54, 0xb2, 0x2d, 0xb3, 0x09,
	0x8b, 0xd8, 0x24, 0x08, 0xa9, 0xca, 0x06, 0xc9, 0x95, 0x40, 0x95, 0x40, 0x02, 0xb3, 0x63, 0x63,
	0x4d, 0xc6, 0x43, 0x3a, 0xc2, 0xf6, 0x84, 0xcc, 0x24, 0x22, 0x37, 0x60, 0x89, 0x38, 0x01, 0x6b,
	0xce, 0xc0, 0x01, 0xba, 0xec, 0xb2, 0xab, 0x80, 0x92, 0x1b, 0xe4, 0x04, 0xc8, 0x33, 0x63, 0xda,
	0xa6, 0x29, 0xbb, 0x9f, 0xe7, 0xf7, 0xde, 0x7f, 0x7a, 0x3f, 0x03, 0x5d, 0x36, 0x20, 0x01, 0xe1,
	0x63, 0x1a, 0x90, 0x94, 0xd1, 0x5c, 0x06, 0xd3, 0x6e, 0x30, 0xa4, 0x39, 0x15, 0x4c, 0xf8, 0xa3,
	0x31, 0x97, 0x1c, 0x21, 0x36, 0x20, 0x7e, 0xc1, 0xf0, 0x35, 0xc3, 0x9f, 0x76, 0x5b, 0xce, 0x06,
	0x95, 0xf9, 0xaa, 0x44, 0xad, 0xe6, 0x90, 0x0f, 0xb9, 0x1a, 0x83, 0x62, 0xd2, 0xa8, 0x77, 0x59,
	0x83, 0xfb, 0xaf, 0xb4, 0xf9, 0x7b, 0x89, 0x25, 0x45, 0x04, 0xde, 0xd3, 0x32, 0x61, 0x01, 0xb7,
	0xda, 0xde, 0xeb, 0x3d, 0xf1, 0x6f, 0x6f, 0xf3, 0x4f, 0x13, 0x9a, 0x4b, 0xf6, 0x91, 0xd1, 0xe4,
	0x44, 0x61, 0x4a, 0x1b, 0xda, 0xe7, 0x73, 0xa7, 0xf2, 0xf3, 0xb7, 0x73, 0xb8, 0xf1, 0xb3, 0x88,
	0x4a, 0x67, 0xf4, 0x1d, 0xc0, 0x07, 0x66, 0x8e, 0x09, 0xcf, 0x05, 0xcd, 0xc5, 0x44, 0x58, 0x5b,
	0x77, 0xef, 0xd3, 0x36, 0x27, 0x25, 0x55, 0xfb, 0x85, 0xfd, 0x62, 0xdf, 0x6a, 0xee, 0x58, 0x33,
	0x9c, 0xa5, 0x7d, 0xef, 0x96, 0xa3, 0x57, 0x64, 0xd1, 0x52, 0xb1, 0xa6, 0x8d, 0x1a, 0x64, 0x0d,
	0x47, 0x33, 0x58, 0x62, 0x71, 0x46, 0x25, 0x4e, 0xb0, 0xc4, 0x56, 0x55, 0x45, 0xea, 0xfc, 0xbf,
	0x02, 0xd3, 0xdf, 0x1b, 0x23, 0x0a, 0x1d, 0x13, 0xeb, 0xe1, 0xcd, 0x58, 0xa5, 0xa9, 0x17, 0x1d,
	0x18, 0xa8, 0x54, 0xa0, 0x63, 0x58, 0x1f, 0xe1, 0x31, 0xce, 0x84, 0x55, 0x73, 0x41, 0x7b, 0xaf,
	0xd7, 0xda, 0xb4, 0xf0, 0xad, 0x62, 0x84, 0xb5, 0xc2, 0x3d, 0x32, 0x7c, 0xf4, 0x12, 0x36, 0xc8,
	0x98, 0x62, 0x49, 0xe3, 0x94, 0x13, 0x9c, 0x9e, 0x71, 0x21, 0xad, 0x6d, 0x17, 0xb4, 0x77, 0xc2,
	0xa3, 0x6b, 0x09, 0xd6, 0x18, 0x45, 0x02, 0x05, 0xbd, 0x2e, 0x11, 0xf4, 0x0e, 0x36, 0x73, 0xfa,
	0x45, 0xc6, 0x7a, 0x5d, 0x2c, 0xe8, 0xe7, 0x09, 0xcd, 0x09, 0xb5, 0xea, 0x2e, 0x68, 0xd7, 0x42,
	0x67, 0x35, 0x77, 0x8e, 0xb4, 0xd7, 0x26, 0x96, 0x17, 0xa1, 0x02, 0x36, 0xb7, 0x2e, 0xc1, 0x17,
	0xf0, 0x60, 0xad, 0x19, 0xd4, 0x80, 0xd5, 0x4f, 0x74, 0x66, 0x01, 0x17, 0xb4, 0xf7, 0xa3, 0x62,
	0x44, 0x4d, 0xb8, 0x3d, 0xc5, 0xe9, 0x84, 0x5a, 0x5b, 0x0a, 0xd3, 0x3f, 0xfa, 0xb5, 0xaf, 0x3f,
	0x9c, 0x8a, 0xf7, 0x0b, 0xc0, 0x47, 0x77, 0xb6, 0x8c, 0xba, 0x70, 0xd7, 0xc4, 0x60, 0x89, 0x72,
	0xdc, 0x0d, 0x9b, 0xab, 0xb9, 0xd3, 0xb8, 0x5e, 0x7a, 0xcc, 0x12, 0x2f, 0xda, 0xd1, 0xf3, 0x69,
	0x82, 0x52, 0x68, 0x9a, 0xbf, 0x3a, 0xb0, 0xfe, 0xcf, 0x3d, 0xde, 0xd4, 0xf7, 0xfa, 0x59, 0x6d,
	0x73, 0xd6, 0xc3, 0x1b, 0x1b, 0xae, 0xae, 0x7a, 0x5f, 0x23, 0xff, 0xf8, 0xd1, 0xf9, 0xc2, 0x06,
	0x17, 0x0b, 0x1b, 0xfc, 0x59, 0xd8, 0xe0, 0xdb, 0xd2, 0xae, 0x5c, 0x2c, 0xed, 0xca, 0xe5, 0xd2,
	0xae, 0x7c, 0x38, 0x1e, 0x32, 0x79, 0x36, 0x19, 0xf8, 0x84, 0x67, 0x01, 0xe1, 0x22, 0xe3, 0x22,
	0x60, 0x03, 0xd2, 0x19, 0xf2, 0x60, 0xfa, 0x3c, 0xc8, 0x78, 0x32, 0x49, 0xa9, 0xd0, 0x6f, 0xf9,
	0x69, 0xaf, 0x63, 0x9e, 0xb3, 0x9c, 0x8d, 0xa8, 0x18, 0xd4, 0xd5, 0xab, 0x7d, 0xf6, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xe7, 0x4b, 0xf9, 0x6b, 0x24, 0x04, 0x00, 0x00,
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
	if m.NextClientSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextClientSequence))
		i--
		dAtA[i] = 0x30
	}
	if m.CreateLocalhost {
		i--
		if m.CreateLocalhost {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ClientsMetadata) > 0 {
		for iNdEx := len(m.ClientsMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientsMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ClientsConsensus) > 0 {
		for iNdEx := len(m.ClientsConsensus) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientsConsensus[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Clients) > 0 {
		for iNdEx := len(m.Clients) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Clients[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IdentifiedGenesisMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedGenesisMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedGenesisMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientMetadata) > 0 {
		for iNdEx := len(m.ClientMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientId)))
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
	if len(m.Clients) > 0 {
		for _, e := range m.Clients {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientsConsensus) > 0 {
		for _, e := range m.ClientsConsensus {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientsMetadata) > 0 {
		for _, e := range m.ClientsMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.CreateLocalhost {
		n += 2
	}
	if m.NextClientSequence != 0 {
		n += 1 + sovGenesis(uint64(m.NextClientSequence))
	}
	return n
}

func (m *GenesisMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *IdentifiedGenesisMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ClientMetadata) > 0 {
		for _, e := range m.ClientMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Clients", wireType)
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
			m.Clients = append(m.Clients, IdentifiedClientState{})
			if err := m.Clients[len(m.Clients)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientsConsensus", wireType)
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
			m.ClientsConsensus = append(m.ClientsConsensus, ClientConsensusStates{})
			if err := m.ClientsConsensus[len(m.ClientsConsensus)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientsMetadata", wireType)
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
			m.ClientsMetadata = append(m.ClientsMetadata, IdentifiedGenesisMetadata{})
			if err := m.ClientsMetadata[len(m.ClientsMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateLocalhost", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CreateLocalhost = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextClientSequence", wireType)
			}
			m.NextClientSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextClientSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *GenesisMetadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
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
func (m *IdentifiedGenesisMetadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IdentifiedGenesisMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedGenesisMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
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
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientMetadata", wireType)
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
			m.ClientMetadata = append(m.ClientMetadata, GenesisMetadata{})
			if err := m.ClientMetadata[len(m.ClientMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
