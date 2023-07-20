// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	// leader earnings rates
	// 1,2,3 Comma separated, no space
	EarnRates string `protobuf:"bytes,1,opt,name=earn_rates,json=earnRates,proto3" json:"earn_rates,omitempty"`
	// pool burning rate
	BurnRate string `protobuf:"bytes,2,opt,name=burn_rate,json=burnRate,proto3" json:"burn_rate,omitempty"`
	// burn coin
	BurnCoin string `protobuf:"bytes,3,opt,name=burn_coin,json=burnCoin,proto3" json:"burn_coin,omitempty"`
	// pool limit is (parameter / 1000), 9999 representing as 99.99%
	// pool ratio
	PoolRatio string `protobuf:"bytes,4,opt,name=pool_ratio,json=poolRatio,proto3" json:"pool_ratio,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_43480567057de0ef, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEarnRates() string {
	if m != nil {
		return m.EarnRates
	}
	return ""
}

func (m *Params) GetBurnRate() string {
	if m != nil {
		return m.BurnRate
	}
	return ""
}

func (m *Params) GetBurnCoin() string {
	if m != nil {
		return m.BurnCoin
	}
	return ""
}

func (m *Params) GetPoolRatio() string {
	if m != nil {
		return m.PoolRatio
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "pendulumlabs.market.market.Params")
}

func init() { proto.RegisterFile("market/params.proto", fileDescriptor_43480567057de0ef) }

var fileDescriptor_43480567057de0ef = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0xd1, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x2a, 0x48, 0xcd, 0x4b, 0x29, 0xcd, 0x29, 0xcd, 0xcd, 0x49, 0x4c, 0x2a, 0xd6, 0x83, 0xa8,
	0x80, 0x52, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x65, 0xfa, 0x20, 0x16, 0x44, 0x87, 0x52,
	0x1b, 0x23, 0x17, 0x5b, 0x00, 0xd8, 0x08, 0x21, 0x59, 0x2e, 0xae, 0xd4, 0xc4, 0xa2, 0xbc, 0xf8,
	0xa2, 0xc4, 0x92, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0x90, 0x48, 0x10,
	0x48, 0x40, 0x48, 0x9a, 0x8b, 0x33, 0xa9, 0x14, 0x2a, 0x2d, 0xc1, 0x04, 0x96, 0xe5, 0x00, 0x09,
	0x80, 0x64, 0xe1, 0x92, 0xc9, 0xf9, 0x99, 0x79, 0x12, 0xcc, 0x08, 0x49, 0xe7, 0xfc, 0xcc, 0x3c,
	0x90, 0xc1, 0x05, 0xf9, 0xf9, 0x39, 0x20, 0x9d, 0x99, 0xf9, 0x12, 0x2c, 0x10, 0x83, 0x41, 0x22,
	0x41, 0x20, 0x01, 0x2b, 0x96, 0x19, 0x0b, 0xe4, 0x19, 0x9c, 0x3c, 0x4e, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x2f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0x1f, 0xe6, 0x3f, 0x5d, 0x90, 0x07, 0xf5, 0xa1, 0x41, 0x50, 0x01, 0x63, 0x94, 0x54, 0x16,
	0xa4, 0x16, 0x27, 0xb1, 0x81, 0x7d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xe0, 0xb0,
	0xe1, 0x22, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolRatio) > 0 {
		i -= len(m.PoolRatio)
		copy(dAtA[i:], m.PoolRatio)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PoolRatio)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BurnCoin) > 0 {
		i -= len(m.BurnCoin)
		copy(dAtA[i:], m.BurnCoin)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BurnCoin)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BurnRate) > 0 {
		i -= len(m.BurnRate)
		copy(dAtA[i:], m.BurnRate)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BurnRate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EarnRates) > 0 {
		i -= len(m.EarnRates)
		copy(dAtA[i:], m.EarnRates)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EarnRates)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EarnRates)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.BurnRate)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.BurnCoin)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.PoolRatio)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EarnRates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EarnRates = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnRate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnCoin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnCoin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
