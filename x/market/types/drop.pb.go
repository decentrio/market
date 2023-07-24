// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/drop.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Drop struct {
	Uid     uint64                                 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner   string                                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Pair    string                                 `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	Drops   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=drops,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"drops"`
	Product github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=product,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"product"`
	Active  bool                                   `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *Drop) Reset()         { *m = Drop{} }
func (m *Drop) String() string { return proto.CompactTextString(m) }
func (*Drop) ProtoMessage()    {}
func (*Drop) Descriptor() ([]byte, []int) {
	return fileDescriptor_3961bee11a1276cb, []int{0}
}
func (m *Drop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Drop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Drop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Drop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Drop.Merge(m, src)
}
func (m *Drop) XXX_Size() int {
	return m.Size()
}
func (m *Drop) XXX_DiscardUnknown() {
	xxx_messageInfo_Drop.DiscardUnknown(m)
}

var xxx_messageInfo_Drop proto.InternalMessageInfo

type Drops struct {
	Uids []uint64 `protobuf:"varint,1,rep,packed,name=uids,proto3" json:"uids,omitempty"`
}

func (m *Drops) Reset()         { *m = Drops{} }
func (m *Drops) String() string { return proto.CompactTextString(m) }
func (*Drops) ProtoMessage()    {}
func (*Drops) Descriptor() ([]byte, []int) {
	return fileDescriptor_3961bee11a1276cb, []int{1}
}
func (m *Drops) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Drops) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Drops.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Drops) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Drops.Merge(m, src)
}
func (m *Drops) XXX_Size() int {
	return m.Size()
}
func (m *Drops) XXX_DiscardUnknown() {
	xxx_messageInfo_Drops.DiscardUnknown(m)
}

var xxx_messageInfo_Drops proto.InternalMessageInfo

type DropsSum struct {
	Sum github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=sum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sum"`
}

func (m *DropsSum) Reset()         { *m = DropsSum{} }
func (m *DropsSum) String() string { return proto.CompactTextString(m) }
func (*DropsSum) ProtoMessage()    {}
func (*DropsSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3961bee11a1276cb, []int{2}
}
func (m *DropsSum) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DropsSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DropsSum.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DropsSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropsSum.Merge(m, src)
}
func (m *DropsSum) XXX_Size() int {
	return m.Size()
}
func (m *DropsSum) XXX_DiscardUnknown() {
	xxx_messageInfo_DropsSum.DiscardUnknown(m)
}

var xxx_messageInfo_DropsSum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Drop)(nil), "pendulumlabs.market.market.Drop")
	proto.RegisterType((*Drops)(nil), "pendulumlabs.market.market.Drops")
	proto.RegisterType((*DropsSum)(nil), "pendulumlabs.market.market.DropsSum")
}

func init() { proto.RegisterFile("market/drop.proto", fileDescriptor_3961bee11a1276cb) }

var fileDescriptor_3961bee11a1276cb = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xe3, 0x7f, 0x92, 0xfe, 0xcb, 0x9d, 0xc0, 0xaa, 0x90, 0x55, 0x24, 0x37, 0xea, 0x80,
	0xb2, 0x34, 0x19, 0x78, 0x01, 0x54, 0x75, 0x00, 0x89, 0x01, 0x85, 0x8d, 0x2d, 0x4d, 0xa2, 0x12,
	0xb5, 0xae, 0x2d, 0x7f, 0xf0, 0xf1, 0x16, 0x3c, 0x56, 0xc7, 0x8e, 0x15, 0x43, 0x45, 0xdb, 0x8d,
	0xa7, 0x40, 0x76, 0x52, 0x89, 0x15, 0xa6, 0xfb, 0xbb, 0xf7, 0xe6, 0x9c, 0x5c, 0xf9, 0xc0, 0x19,
	0xcb, 0xe5, 0xbc, 0xd2, 0x69, 0x29, 0xb9, 0x48, 0x84, 0xe4, 0x9a, 0xe3, 0xbe, 0xa8, 0x96, 0xa5,
	0x59, 0x18, 0xb6, 0xc8, 0xa7, 0x2a, 0x69, 0xf6, 0x6d, 0xe9, 0xf7, 0x66, 0x7c, 0xc6, 0xdd, 0x67,
	0xa9, 0xa5, 0x46, 0x31, 0xfc, 0x42, 0x10, 0x4c, 0x24, 0x17, 0xf8, 0x14, 0x7c, 0x53, 0x97, 0x04,
	0x45, 0x28, 0x0e, 0x32, 0x8b, 0xb8, 0x07, 0x21, 0x7f, 0x59, 0x56, 0x92, 0xfc, 0x8b, 0x50, 0x7c,
	0x92, 0x35, 0x0d, 0xc6, 0x10, 0x88, 0xbc, 0x96, 0xc4, 0x77, 0x43, 0xc7, 0x78, 0x02, 0xa1, 0x3d,
	0x42, 0x91, 0xc0, 0x0e, 0xc7, 0xc9, 0x6a, 0x3b, 0xf0, 0x3e, 0xb6, 0x83, 0xcb, 0x59, 0xad, 0x9f,
	0xcc, 0x34, 0x29, 0x38, 0x4b, 0x0b, 0xae, 0x18, 0x57, 0x6d, 0x19, 0xa9, 0x72, 0x9e, 0xea, 0x37,
	0x51, 0xa9, 0xe4, 0x76, 0xa9, 0xb3, 0x46, 0x8c, 0x6f, 0xe0, 0xbf, 0x90, 0xbc, 0x34, 0x85, 0x26,
	0xe1, 0x9f, 0x7c, 0x8e, 0x72, 0x7c, 0x0e, 0x9d, 0xbc, 0xd0, 0xf5, 0x73, 0x45, 0x3a, 0x11, 0x8a,
	0xbb, 0x59, 0xdb, 0x0d, 0x2f, 0x20, 0x9c, 0xb8, 0x5f, 0x61, 0x08, 0x4c, 0x5d, 0x2a, 0x82, 0x22,
	0x3f, 0x0e, 0x32, 0xc7, 0xc3, 0x3b, 0xe8, 0xba, 0xe5, 0x83, 0x61, 0xf8, 0x1a, 0x7c, 0x65, 0x98,
	0x7b, 0x8c, 0xdf, 0x9f, 0x61, 0xa5, 0xe3, 0xfb, 0xd5, 0x8e, 0x7a, 0x9b, 0x1d, 0x45, 0xab, 0x3d,
	0x45, 0xeb, 0x3d, 0x45, 0x9f, 0x7b, 0x8a, 0xde, 0x0f, 0xd4, 0x5b, 0x1f, 0xa8, 0xb7, 0x39, 0x50,
	0xef, 0x31, 0xf9, 0x61, 0x77, 0x8c, 0x6d, 0x64, 0x73, 0x4b, 0xdb, 0x5c, 0x5f, 0x8f, 0xe0, 0xac,
	0xa7, 0x1d, 0x17, 0xd8, 0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0xdd, 0x7a, 0x4e, 0xf7,
	0x01, 0x00, 0x00,
}

func (m *Drop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Drop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Drop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Product.Size()
		i -= size
		if _, err := m.Product.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDrop(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Drops.Size()
		i -= size
		if _, err := m.Drops.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDrop(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintDrop(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDrop(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Uid != 0 {
		i = encodeVarintDrop(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Drops) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Drops) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Drops) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Uids) > 0 {
		dAtA2 := make([]byte, len(m.Uids)*10)
		var j1 int
		for _, num := range m.Uids {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintDrop(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DropsSum) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DropsSum) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DropsSum) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Sum.Size()
		i -= size
		if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDrop(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDrop(dAtA []byte, offset int, v uint64) int {
	offset -= sovDrop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Drop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovDrop(uint64(m.Uid))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDrop(uint64(l))
	}
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovDrop(uint64(l))
	}
	l = m.Drops.Size()
	n += 1 + l + sovDrop(uint64(l))
	l = m.Product.Size()
	n += 1 + l + sovDrop(uint64(l))
	if m.Active {
		n += 2
	}
	return n
}

func (m *Drops) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Uids) > 0 {
		l = 0
		for _, e := range m.Uids {
			l += sovDrop(uint64(e))
		}
		n += 1 + sovDrop(uint64(l)) + l
	}
	return n
}

func (m *DropsSum) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Sum.Size()
	n += 1 + l + sovDrop(uint64(l))
	return n
}

func sovDrop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDrop(x uint64) (n int) {
	return sovDrop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Drop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDrop
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
			return fmt.Errorf("proto: Drop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Drop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Drops", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Drops.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Product", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Product.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDrop
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
func (m *Drops) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDrop
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
			return fmt.Errorf("proto: Drops: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Drops: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDrop
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
				m.Uids = append(m.Uids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDrop
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
					return ErrInvalidLengthDrop
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthDrop
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
				if elementCount != 0 && len(m.Uids) == 0 {
					m.Uids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDrop
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
					m.Uids = append(m.Uids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Uids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDrop
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
func (m *DropsSum) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDrop
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
			return fmt.Errorf("proto: DropsSum: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DropsSum: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDrop
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
func skipDrop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDrop
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
					return 0, ErrIntOverflowDrop
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
					return 0, ErrIntOverflowDrop
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
				return 0, ErrInvalidLengthDrop
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDrop
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDrop
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDrop        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDrop          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDrop = fmt.Errorf("proto: unexpected end of group")
)
