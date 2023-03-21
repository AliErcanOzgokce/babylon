// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/v1/genesis.proto

package types

import (
	fmt "fmt"
	ed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the checkpointing module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to checkpointing
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// genesis_keys defines the public keys for the genesis validators
	GenesisKeys []*GenesisKey `protobuf:"bytes,2,rep,name=genesis_keys,json=genesisKeys,proto3" json:"genesis_keys,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c524ebc9800de, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGenesisKeys() []*GenesisKey {
	if m != nil {
		return m.GenesisKeys
	}
	return nil
}

// GenesisKey defines public key information about the genesis validators
type GenesisKey struct {
	// validator_address is the address corresponding to a validator
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// bls_key defines the BLS key of the validator at genesis
	BlsKey *BlsKey `protobuf:"bytes,2,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	// val_pubkey defines the ed25519 public key of the validator at genesis
	ValPubkey *ed25519.PubKey `protobuf:"bytes,3,opt,name=val_pubkey,json=valPubkey,proto3" json:"val_pubkey,omitempty"`
}

func (m *GenesisKey) Reset()         { *m = GenesisKey{} }
func (m *GenesisKey) String() string { return proto.CompactTextString(m) }
func (*GenesisKey) ProtoMessage()    {}
func (*GenesisKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c524ebc9800de, []int{1}
}
func (m *GenesisKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisKey.Merge(m, src)
}
func (m *GenesisKey) XXX_Size() int {
	return m.Size()
}
func (m *GenesisKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisKey.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisKey proto.InternalMessageInfo

func (m *GenesisKey) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *GenesisKey) GetBlsKey() *BlsKey {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *GenesisKey) GetValPubkey() *ed25519.PubKey {
	if m != nil {
		return m.ValPubkey
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "babylon.checkpointing.v1.GenesisState")
	proto.RegisterType((*GenesisKey)(nil), "babylon.checkpointing.v1.GenesisKey")
}

func init() {
	proto.RegisterFile("babylon/checkpointing/v1/genesis.proto", fileDescriptor_bf2c524ebc9800de)
}

var fileDescriptor_bf2c524ebc9800de = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x6a, 0xea, 0x40,
	0x18, 0x85, 0x33, 0x7a, 0xf1, 0xe2, 0xe8, 0xe2, 0xde, 0x70, 0x17, 0x41, 0xb8, 0x69, 0x90, 0xb6,
	0x08, 0x85, 0x19, 0xb4, 0xb8, 0x10, 0x4a, 0xa1, 0x6e, 0x5c, 0x74, 0x51, 0xb1, 0xbb, 0x6e, 0x64,
	0x26, 0x19, 0x62, 0x30, 0x66, 0x42, 0x66, 0x0c, 0xcd, 0x5b, 0x74, 0xd7, 0xf7, 0xe8, 0x53, 0xb8,
	0x74, 0xd9, 0x55, 0x29, 0xfa, 0x22, 0x25, 0x33, 0xa3, 0xa5, 0x85, 0xd0, 0x55, 0x26, 0x99, 0xef,
	0xfc, 0xe7, 0xfc, 0x39, 0xf0, 0x9c, 0x12, 0x5a, 0xc4, 0x3c, 0xc1, 0xfe, 0x82, 0xf9, 0xcb, 0x94,
	0x47, 0x89, 0x8c, 0x92, 0x10, 0xe7, 0x7d, 0x1c, 0xb2, 0x84, 0x89, 0x48, 0xa0, 0x34, 0xe3, 0x92,
	0xdb, 0x8e, 0xe1, 0xd0, 0x17, 0x0e, 0xe5, 0xfd, 0xce, 0xbf, 0x90, 0x87, 0x5c, 0x41, 0xb8, 0x3c,
	0x69, 0xbe, 0xe3, 0xf9, 0x5c, 0xac, 0xb8, 0xc0, 0x7e, 0x56, 0xa4, 0x92, 0x63, 0x16, 0x0c, 0x86,
	0xc3, 0xfe, 0x08, 0x2f, 0x59, 0x61, 0x26, 0x76, 0xce, 0x2a, 0x9d, 0x53, 0x92, 0x91, 0xd5, 0x01,
	0xab, 0x0e, 0x48, 0x63, 0x31, 0x5f, 0xb2, 0x42, 0x73, 0xdd, 0x67, 0x00, 0xdb, 0x13, 0x1d, 0xf9,
	0x5e, 0x12, 0xc9, 0xec, 0x6b, 0xd8, 0xd0, 0x83, 0x1c, 0xe0, 0x81, 0x5e, 0x6b, 0xe0, 0xa1, 0xaa,
	0x15, 0xd0, 0x54, 0x71, 0xe3, 0x5f, 0x9b, 0xb7, 0x13, 0x6b, 0x66, 0x54, 0xf6, 0x04, 0xb6, 0xcd,
	0x2f, 0x28, 0x5d, 0x84, 0x53, 0xf3, 0xea, 0xbd, 0xd6, 0xe0, 0xb4, 0x7a, 0x8a, 0x71, 0xbf, 0x65,
	0xc5, 0xac, 0x15, 0x1e, 0xcf, 0xa2, 0xfb, 0x02, 0x20, 0xfc, 0xbc, 0xb3, 0x2f, 0xe0, 0xdf, 0x9c,
	0xc4, 0x51, 0x40, 0x24, 0xcf, 0xe6, 0x24, 0x08, 0x32, 0x26, 0x74, 0xc4, 0xe6, 0xec, 0xcf, 0xf1,
	0xe2, 0x46, 0x7f, 0xb7, 0x47, 0xf0, 0xb7, 0x59, 0xd3, 0xa9, 0xfd, 0xb4, 0xc5, 0x38, 0x56, 0xde,
	0x0d, 0xaa, 0x9e, 0xf6, 0x15, 0x84, 0x39, 0x89, 0xe7, 0xe9, 0x9a, 0x96, 0xea, 0xba, 0x52, 0xff,
	0x47, 0xba, 0x16, 0xa4, 0x6b, 0x41, 0xa6, 0x16, 0x34, 0x5d, 0xd3, 0x52, 0xda, 0xcc, 0x49, 0x3c,
	0x55, 0xfc, 0xf8, 0x6e, 0xb3, 0x73, 0xc1, 0x76, 0xe7, 0x82, 0xf7, 0x9d, 0x0b, 0x9e, 0xf6, 0xae,
	0xb5, 0xdd, 0xbb, 0xd6, 0xeb, 0xde, 0xb5, 0x1e, 0x86, 0x61, 0x24, 0x17, 0x6b, 0x8a, 0x7c, 0xbe,
	0xc2, 0x26, 0x8b, 0xbf, 0x20, 0x51, 0x72, 0x78, 0xc1, 0x8f, 0xdf, 0xaa, 0x92, 0x45, 0xca, 0x04,
	0x6d, 0xa8, 0x9a, 0x2e, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x9e, 0x6b, 0xd7, 0x71, 0x02,
	0x00, 0x00,
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
	if len(m.GenesisKeys) > 0 {
		for iNdEx := len(m.GenesisKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValPubkey != nil {
		{
			size, err := m.ValPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlsKey != nil {
		{
			size, err := m.BlsKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.GenesisKeys) > 0 {
		for _, e := range m.GenesisKeys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.BlsKey != nil {
		l = m.BlsKey.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ValPubkey != nil {
		l = m.ValPubkey.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisKeys", wireType)
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
			m.GenesisKeys = append(m.GenesisKeys, &GenesisKey{})
			if err := m.GenesisKeys[len(m.GenesisKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsKey", wireType)
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
			if m.BlsKey == nil {
				m.BlsKey = &BlsKey{}
			}
			if err := m.BlsKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValPubkey", wireType)
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
			if m.ValPubkey == nil {
				m.ValPubkey = &ed25519.PubKey{}
			}
			if err := m.ValPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
