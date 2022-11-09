// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: electra/meter/customerbills.proto

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

type Customerbills struct {
	BillCycleID      uint64 `protobuf:"varint,1,opt,name=billCycleID,proto3" json:"billCycleID,omitempty"`
	CustomerDeviceID string `protobuf:"bytes,2,opt,name=customerDeviceID,proto3" json:"customerDeviceID,omitempty"`
	BillDate         uint64 `protobuf:"varint,3,opt,name=billDate,proto3" json:"billDate,omitempty"`
	BillTotalWh      uint64 `protobuf:"varint,4,opt,name=billTotalWh,proto3" json:"billTotalWh,omitempty"`
	BillTotalPrice   uint64 `protobuf:"varint,5,opt,name=billTotalPrice,proto3" json:"billTotalPrice,omitempty"`
	BillCurrency     string `protobuf:"bytes,6,opt,name=billCurrency,proto3" json:"billCurrency,omitempty"`
	BillValid        bool   `protobuf:"varint,7,opt,name=billValid,proto3" json:"billValid,omitempty"`
	Paid             bool   `protobuf:"varint,8,opt,name=paid,proto3" json:"paid,omitempty"`
	Creator          string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Customerbills) Reset()         { *m = Customerbills{} }
func (m *Customerbills) String() string { return proto.CompactTextString(m) }
func (*Customerbills) ProtoMessage()    {}
func (*Customerbills) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c7956d3e54163af, []int{0}
}
func (m *Customerbills) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Customerbills) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Customerbills.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Customerbills) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customerbills.Merge(m, src)
}
func (m *Customerbills) XXX_Size() int {
	return m.Size()
}
func (m *Customerbills) XXX_DiscardUnknown() {
	xxx_messageInfo_Customerbills.DiscardUnknown(m)
}

var xxx_messageInfo_Customerbills proto.InternalMessageInfo

func (m *Customerbills) GetBillCycleID() uint64 {
	if m != nil {
		return m.BillCycleID
	}
	return 0
}

func (m *Customerbills) GetCustomerDeviceID() string {
	if m != nil {
		return m.CustomerDeviceID
	}
	return ""
}

func (m *Customerbills) GetBillDate() uint64 {
	if m != nil {
		return m.BillDate
	}
	return 0
}

func (m *Customerbills) GetBillTotalWh() uint64 {
	if m != nil {
		return m.BillTotalWh
	}
	return 0
}

func (m *Customerbills) GetBillTotalPrice() uint64 {
	if m != nil {
		return m.BillTotalPrice
	}
	return 0
}

func (m *Customerbills) GetBillCurrency() string {
	if m != nil {
		return m.BillCurrency
	}
	return ""
}

func (m *Customerbills) GetBillValid() bool {
	if m != nil {
		return m.BillValid
	}
	return false
}

func (m *Customerbills) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

func (m *Customerbills) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Customerbills)(nil), "electra.meter.Customerbills")
}

func init() { proto.RegisterFile("electra/meter/customerbills.proto", fileDescriptor_4c7956d3e54163af) }

var fileDescriptor_4c7956d3e54163af = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x35, 0xb6, 0xc9, 0x6a, 0x45, 0x06, 0x84, 0x45, 0x64, 0x89, 0x3d, 0x48, 0xf0,
	0xd0, 0x1c, 0x7c, 0x03, 0x9b, 0x8b, 0x37, 0x09, 0xa2, 0xe0, 0x6d, 0xbb, 0x1d, 0x30, 0xb0, 0x75,
	0xc3, 0x66, 0x2b, 0xe6, 0x2d, 0x7c, 0x17, 0x5f, 0xc2, 0x63, 0x8f, 0x1e, 0x25, 0x79, 0x11, 0xc9,
	0x92, 0xb4, 0xd5, 0xde, 0xe6, 0xff, 0xf6, 0xdf, 0x7f, 0x86, 0x19, 0x7a, 0x89, 0x0a, 0xa5, 0x35,
	0x22, 0x59, 0xa2, 0x45, 0x93, 0xc8, 0x55, 0x69, 0xf5, 0x12, 0xcd, 0x3c, 0x57, 0xaa, 0x9c, 0x16,
	0x46, 0x5b, 0x0d, 0xe3, 0xce, 0x32, 0x75, 0x96, 0xc9, 0xe7, 0x80, 0x8e, 0x67, 0xbb, 0x36, 0x88,
	0xe8, 0x51, 0x5b, 0xcc, 0x2a, 0xa9, 0xf0, 0x2e, 0x65, 0x24, 0x22, 0xb1, 0x9f, 0xed, 0x22, 0xb8,
	0xa6, 0xa7, 0x7d, 0x72, 0x8a, 0x6f, 0xb9, 0x6c, 0x6d, 0x83, 0x88, 0xc4, 0x61, 0xb6, 0xc7, 0xe1,
	0x9c, 0x06, 0xed, 0xd7, 0x54, 0x58, 0x64, 0x07, 0x2e, 0x6a, 0xa3, 0xfb, 0x4e, 0x0f, 0xda, 0x0a,
	0xf5, 0xf4, 0xc2, 0xfc, 0x6d, 0xa7, 0x0e, 0xc1, 0x15, 0x3d, 0xd9, 0xc8, 0x7b, 0x93, 0x4b, 0x64,
	0x87, 0xce, 0xf4, 0x8f, 0xc2, 0x84, 0x1e, 0xbb, 0x01, 0x57, 0xc6, 0xe0, 0xab, 0xac, 0xd8, 0xd0,
	0x4d, 0xf3, 0x87, 0xc1, 0x05, 0x0d, 0x5b, 0xfd, 0x28, 0x54, 0xbe, 0x60, 0xa3, 0x88, 0xc4, 0x41,
	0xb6, 0x05, 0x00, 0xd4, 0x2f, 0x44, 0xbe, 0x60, 0x81, 0x7b, 0x70, 0x35, 0x30, 0x3a, 0x92, 0x06,
	0x85, 0xd5, 0x86, 0x85, 0x2e, 0xb0, 0x97, 0xb7, 0xc9, 0x57, 0xcd, 0xc9, 0xba, 0xe6, 0xe4, 0xa7,
	0xe6, 0xe4, 0xa3, 0xe1, 0xde, 0xba, 0xe1, 0xde, 0x77, 0xc3, 0xbd, 0xe7, 0xb3, 0xfe, 0x02, 0xef,
	0xdd, 0x0d, 0x6c, 0x55, 0x60, 0x39, 0x1f, 0xba, 0xe5, 0xdf, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xcc, 0x27, 0x2d, 0xe0, 0xa1, 0x01, 0x00, 0x00,
}

func (m *Customerbills) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Customerbills) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Customerbills) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCustomerbills(dAtA, i, uint64(len(m.Creator)))
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
		i = encodeVarintCustomerbills(dAtA, i, uint64(len(m.BillCurrency)))
		i--
		dAtA[i] = 0x32
	}
	if m.BillTotalPrice != 0 {
		i = encodeVarintCustomerbills(dAtA, i, uint64(m.BillTotalPrice))
		i--
		dAtA[i] = 0x28
	}
	if m.BillTotalWh != 0 {
		i = encodeVarintCustomerbills(dAtA, i, uint64(m.BillTotalWh))
		i--
		dAtA[i] = 0x20
	}
	if m.BillDate != 0 {
		i = encodeVarintCustomerbills(dAtA, i, uint64(m.BillDate))
		i--
		dAtA[i] = 0x18
	}
	if len(m.CustomerDeviceID) > 0 {
		i -= len(m.CustomerDeviceID)
		copy(dAtA[i:], m.CustomerDeviceID)
		i = encodeVarintCustomerbills(dAtA, i, uint64(len(m.CustomerDeviceID)))
		i--
		dAtA[i] = 0x12
	}
	if m.BillCycleID != 0 {
		i = encodeVarintCustomerbills(dAtA, i, uint64(m.BillCycleID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCustomerbills(dAtA []byte, offset int, v uint64) int {
	offset -= sovCustomerbills(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Customerbills) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BillCycleID != 0 {
		n += 1 + sovCustomerbills(uint64(m.BillCycleID))
	}
	l = len(m.CustomerDeviceID)
	if l > 0 {
		n += 1 + l + sovCustomerbills(uint64(l))
	}
	if m.BillDate != 0 {
		n += 1 + sovCustomerbills(uint64(m.BillDate))
	}
	if m.BillTotalWh != 0 {
		n += 1 + sovCustomerbills(uint64(m.BillTotalWh))
	}
	if m.BillTotalPrice != 0 {
		n += 1 + sovCustomerbills(uint64(m.BillTotalPrice))
	}
	l = len(m.BillCurrency)
	if l > 0 {
		n += 1 + l + sovCustomerbills(uint64(l))
	}
	if m.BillValid {
		n += 2
	}
	if m.Paid {
		n += 2
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCustomerbills(uint64(l))
	}
	return n
}

func sovCustomerbills(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCustomerbills(x uint64) (n int) {
	return sovCustomerbills(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Customerbills) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomerbills
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
			return fmt.Errorf("proto: Customerbills: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Customerbills: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillCycleID", wireType)
			}
			m.BillCycleID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerbills
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
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerDeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerbills
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
				return ErrInvalidLengthCustomerbills
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomerbills
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustomerDeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillDate", wireType)
			}
			m.BillDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerbills
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
					return ErrIntOverflowCustomerbills
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
					return ErrIntOverflowCustomerbills
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
					return ErrIntOverflowCustomerbills
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
				return ErrInvalidLengthCustomerbills
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomerbills
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
					return ErrIntOverflowCustomerbills
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
					return ErrIntOverflowCustomerbills
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
					return ErrIntOverflowCustomerbills
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
				return ErrInvalidLengthCustomerbills
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomerbills
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomerbills(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCustomerbills
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
func skipCustomerbills(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCustomerbills
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
					return 0, ErrIntOverflowCustomerbills
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
					return 0, ErrIntOverflowCustomerbills
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
				return 0, ErrInvalidLengthCustomerbills
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCustomerbills
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCustomerbills
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCustomerbills        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCustomerbills          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCustomerbills = fmt.Errorf("proto: unexpected end of group")
)
