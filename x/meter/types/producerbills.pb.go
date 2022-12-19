// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: electra/meter/producerbills.proto

package types

import (
	fmt "fmt"
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

type Producerbills struct {
	BillCycleID      uint64 `protobuf:"varint,1,opt,name=billCycleID,proto3" json:"billCycleID,omitempty"`
	ProducerDeviceID string `protobuf:"bytes,2,opt,name=producerDeviceID,proto3" json:"producerDeviceID,omitempty"`
	BillDate         uint64 `protobuf:"varint,3,opt,name=billDate,proto3" json:"billDate,omitempty"`
	BillTotalWh      uint64 `protobuf:"varint,4,opt,name=billTotalWh,proto3" json:"billTotalWh,omitempty"`
	BillTotalPrice   uint64 `protobuf:"varint,5,opt,name=billTotalPrice,proto3" json:"billTotalPrice,omitempty"`
	BillCurrency     string `protobuf:"bytes,6,opt,name=billCurrency,proto3" json:"billCurrency,omitempty"`
	BillValid        bool   `protobuf:"varint,7,opt,name=billValid,proto3" json:"billValid,omitempty"`
	Paid             bool   `protobuf:"varint,8,opt,name=paid,proto3" json:"paid,omitempty"`
	Creator          string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Producerbills) Reset()         { *m = Producerbills{} }
func (m *Producerbills) String() string { return proto.CompactTextString(m) }
func (*Producerbills) ProtoMessage()    {}
func (*Producerbills) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f5797d0f1ba4de0, []int{0}
}
func (m *Producerbills) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Producerbills) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Producerbills.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Producerbills) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Producerbills.Merge(m, src)
}
func (m *Producerbills) XXX_Size() int {
	return m.Size()
}
func (m *Producerbills) XXX_DiscardUnknown() {
	xxx_messageInfo_Producerbills.DiscardUnknown(m)
}

var xxx_messageInfo_Producerbills proto.InternalMessageInfo

func (m *Producerbills) GetBillCycleID() uint64 {
	if m != nil {
		return m.BillCycleID
	}
	return 0
}

func (m *Producerbills) GetProducerDeviceID() string {
	if m != nil {
		return m.ProducerDeviceID
	}
	return ""
}

func (m *Producerbills) GetBillDate() uint64 {
	if m != nil {
		return m.BillDate
	}
	return 0
}

func (m *Producerbills) GetBillTotalWh() uint64 {
	if m != nil {
		return m.BillTotalWh
	}
	return 0
}

func (m *Producerbills) GetBillTotalPrice() uint64 {
	if m != nil {
		return m.BillTotalPrice
	}
	return 0
}

func (m *Producerbills) GetBillCurrency() string {
	if m != nil {
		return m.BillCurrency
	}
	return ""
}

func (m *Producerbills) GetBillValid() bool {
	if m != nil {
		return m.BillValid
	}
	return false
}

func (m *Producerbills) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

func (m *Producerbills) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Producerbills)(nil), "electra.meter.Producerbills")
}

func init() { proto.RegisterFile("electra/meter/producerbills.proto", fileDescriptor_7f5797d0f1ba4de0) }

var fileDescriptor_7f5797d0f1ba4de0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x35, 0xb6, 0xc9, 0x6a, 0x45, 0x06, 0x84, 0x45, 0x64, 0x89, 0x3d, 0x48, 0xf0,
	0xd0, 0x1c, 0x7c, 0x03, 0xcd, 0xc5, 0x5b, 0x09, 0xa2, 0xe0, 0x6d, 0xbb, 0x19, 0x30, 0xb0, 0x9a,
	0x30, 0xdd, 0x8a, 0x79, 0x0b, 0xdf, 0xc5, 0x97, 0xf0, 0xd8, 0xa3, 0x47, 0x49, 0x5e, 0x44, 0xb2,
	0x24, 0x6d, 0xd5, 0xdb, 0xfc, 0xdf, 0xfe, 0xfb, 0xcf, 0x30, 0xc3, 0xcf, 0xd1, 0xa0, 0xb6, 0xa4,
	0xe2, 0x67, 0xb4, 0x48, 0x71, 0x49, 0x45, 0xb6, 0xd2, 0x48, 0x8b, 0xdc, 0x98, 0xe5, 0xac, 0xa4,
	0xc2, 0x16, 0x30, 0xe9, 0x2c, 0x33, 0x67, 0x99, 0x7e, 0x0c, 0xf9, 0x64, 0xbe, 0x6b, 0x83, 0x90,
	0x1f, 0xb4, 0xc5, 0x4d, 0xa5, 0x0d, 0xde, 0x26, 0x82, 0x85, 0x2c, 0xf2, 0xd2, 0x5d, 0x04, 0x97,
	0xfc, 0xb8, 0x4f, 0x4e, 0xf0, 0x35, 0xd7, 0xad, 0x6d, 0x18, 0xb2, 0x28, 0x48, 0xff, 0x71, 0x38,
	0xe5, 0x7e, 0xfb, 0x35, 0x51, 0x16, 0xc5, 0x9e, 0x8b, 0xda, 0xe8, 0xbe, 0xd3, 0x5d, 0x61, 0x95,
	0x79, 0x78, 0x12, 0xde, 0xb6, 0x53, 0x87, 0xe0, 0x82, 0x1f, 0x6d, 0xe4, 0x9c, 0x72, 0x8d, 0x62,
	0xdf, 0x99, 0xfe, 0x50, 0x98, 0xf2, 0x43, 0x37, 0xe0, 0x8a, 0x08, 0x5f, 0x74, 0x25, 0x46, 0x6e,
	0x9a, 0x5f, 0x0c, 0xce, 0x78, 0xd0, 0xea, 0x7b, 0x65, 0xf2, 0x4c, 0x8c, 0x43, 0x16, 0xf9, 0xe9,
	0x16, 0x00, 0x70, 0xaf, 0x54, 0x79, 0x26, 0x7c, 0xf7, 0xe0, 0x6a, 0x10, 0x7c, 0xac, 0x09, 0x95,
	0x2d, 0x48, 0x04, 0x2e, 0xb0, 0x97, 0xd7, 0xf1, 0x67, 0x2d, 0xd9, 0xba, 0x96, 0xec, 0xbb, 0x96,
	0xec, 0xbd, 0x91, 0x83, 0x75, 0x23, 0x07, 0x5f, 0x8d, 0x1c, 0x3c, 0x9e, 0xf4, 0x17, 0x78, 0xeb,
	0x6e, 0x60, 0xab, 0x12, 0x97, 0x8b, 0x91, 0x5b, 0xfe, 0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xae, 0x5e, 0x13, 0xd0, 0xa1, 0x01, 0x00, 0x00,
}

func (m *Producerbills) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Producerbills) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Producerbills) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProducerbills(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Paid {
		i--
		if m.Paid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.BillValid {
		i--
		if m.BillValid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.BillCurrency) > 0 {
		i -= len(m.BillCurrency)
		copy(dAtA[i:], m.BillCurrency)
		i = encodeVarintProducerbills(dAtA, i, uint64(len(m.BillCurrency)))
		i--
		dAtA[i] = 0x32
	}
	if m.BillTotalPrice != 0 {
		i = encodeVarintProducerbills(dAtA, i, uint64(m.BillTotalPrice))
		i--
		dAtA[i] = 0x28
	}
	if m.BillTotalWh != 0 {
		i = encodeVarintProducerbills(dAtA, i, uint64(m.BillTotalWh))
		i--
		dAtA[i] = 0x20
	}
	if m.BillDate != 0 {
		i = encodeVarintProducerbills(dAtA, i, uint64(m.BillDate))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ProducerDeviceID) > 0 {
		i -= len(m.ProducerDeviceID)
		copy(dAtA[i:], m.ProducerDeviceID)
		i = encodeVarintProducerbills(dAtA, i, uint64(len(m.ProducerDeviceID)))
		i--
		dAtA[i] = 0x12
	}
	if m.BillCycleID != 0 {
		i = encodeVarintProducerbills(dAtA, i, uint64(m.BillCycleID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProducerbills(dAtA []byte, offset int, v uint64) int {
	offset -= sovProducerbills(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Producerbills) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BillCycleID != 0 {
		n += 1 + sovProducerbills(uint64(m.BillCycleID))
	}
	l = len(m.ProducerDeviceID)
	if l > 0 {
		n += 1 + l + sovProducerbills(uint64(l))
	}
	if m.BillDate != 0 {
		n += 1 + sovProducerbills(uint64(m.BillDate))
	}
	if m.BillTotalWh != 0 {
		n += 1 + sovProducerbills(uint64(m.BillTotalWh))
	}
	if m.BillTotalPrice != 0 {
		n += 1 + sovProducerbills(uint64(m.BillTotalPrice))
	}
	l = len(m.BillCurrency)
	if l > 0 {
		n += 1 + l + sovProducerbills(uint64(l))
	}
	if m.BillValid {
		n += 2
	}
	if m.Paid {
		n += 2
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProducerbills(uint64(l))
	}
	return n
}

func sovProducerbills(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProducerbills(x uint64) (n int) {
	return sovProducerbills(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Producerbills) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProducerbills
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
			return fmt.Errorf("proto: Producerbills: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Producerbills: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillCycleID", wireType)
			}
			m.BillCycleID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BillCycleID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerDeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
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
				return ErrInvalidLengthProducerbills
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbills
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProducerDeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillDate", wireType)
			}
			m.BillDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BillDate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillTotalWh", wireType)
			}
			m.BillTotalWh = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BillTotalWh |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillTotalPrice", wireType)
			}
			m.BillTotalPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BillTotalPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillCurrency", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
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
				return ErrInvalidLengthProducerbills
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbills
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BillCurrency = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillValid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
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
			m.BillValid = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
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
			m.Paid = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbills
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
				return ErrInvalidLengthProducerbills
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbills
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProducerbills(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProducerbills
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
func skipProducerbills(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProducerbills
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
					return 0, ErrIntOverflowProducerbills
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
					return 0, ErrIntOverflowProducerbills
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
				return 0, ErrInvalidLengthProducerbills
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProducerbills
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProducerbills
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProducerbills        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProducerbills          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProducerbills = fmt.Errorf("proto: unexpected end of group")
)
