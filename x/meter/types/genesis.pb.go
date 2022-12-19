// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: electra/meter/genesis.proto

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

// GenesisState defines the meter module's genesis state.
type GenesisState struct {
	Params                    Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	MeterreadingsList         []Meterreadings         `protobuf:"bytes,2,rep,name=meterreadingsList,proto3" json:"meterreadingsList"`
	MeterdirectoryList        []Meterdirectory        `protobuf:"bytes,3,rep,name=meterdirectoryList,proto3" json:"meterdirectoryList"`
	PowerPurchaseContractList []PowerPurchaseContract `protobuf:"bytes,4,rep,name=powerPurchaseContractList,proto3" json:"powerPurchaseContractList"`
	PpaMapList                []PpaMap                `protobuf:"bytes,5,rep,name=ppaMapList,proto3" json:"ppaMapList"`
	BillingcyclesList         []Billingcycles         `protobuf:"bytes,6,rep,name=billingcyclesList,proto3" json:"billingcyclesList"`
	CustomerbillinglineList   []Customerbillingline   `protobuf:"bytes,7,rep,name=customerbillinglineList,proto3" json:"customerbillinglineList"`
	CustomerbillsList         []Customerbills         `protobuf:"bytes,8,rep,name=customerbillsList,proto3" json:"customerbillsList"`
	ProducerbillinglineList   []Producerbillingline   `protobuf:"bytes,9,rep,name=producerbillinglineList,proto3" json:"producerbillinglineList"`
	ProducerbillsList         []Producerbills         `protobuf:"bytes,10,rep,name=producerbillsList,proto3" json:"producerbillsList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca23e331b4d34eb0, []int{0}
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

func (m *GenesisState) GetMeterreadingsList() []Meterreadings {
	if m != nil {
		return m.MeterreadingsList
	}
	return nil
}

func (m *GenesisState) GetMeterdirectoryList() []Meterdirectory {
	if m != nil {
		return m.MeterdirectoryList
	}
	return nil
}

func (m *GenesisState) GetPowerPurchaseContractList() []PowerPurchaseContract {
	if m != nil {
		return m.PowerPurchaseContractList
	}
	return nil
}

func (m *GenesisState) GetPpaMapList() []PpaMap {
	if m != nil {
		return m.PpaMapList
	}
	return nil
}

func (m *GenesisState) GetBillingcyclesList() []Billingcycles {
	if m != nil {
		return m.BillingcyclesList
	}
	return nil
}

func (m *GenesisState) GetCustomerbillinglineList() []Customerbillingline {
	if m != nil {
		return m.CustomerbillinglineList
	}
	return nil
}

func (m *GenesisState) GetCustomerbillsList() []Customerbills {
	if m != nil {
		return m.CustomerbillsList
	}
	return nil
}

func (m *GenesisState) GetProducerbillinglineList() []Producerbillingline {
	if m != nil {
		return m.ProducerbillinglineList
	}
	return nil
}

func (m *GenesisState) GetProducerbillsList() []Producerbills {
	if m != nil {
		return m.ProducerbillsList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "electra.meter.GenesisState")
}

func init() { proto.RegisterFile("electra/meter/genesis.proto", fileDescriptor_ca23e331b4d34eb0) }

var fileDescriptor_ca23e331b4d34eb0 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x13, 0x77, 0x8d, 0x3a, 0xab, 0x07, 0x07, 0x17, 0xd7, 0xa8, 0xb1, 0x16, 0xc1, 0x82,
	0x90, 0x40, 0x7b, 0xf4, 0xd6, 0x1e, 0xbc, 0x58, 0x08, 0xf6, 0xe6, 0xa5, 0x4c, 0xa7, 0x43, 0x1a,
	0x48, 0x32, 0xc3, 0x64, 0x8a, 0xf6, 0x5b, 0xf8, 0xb1, 0x7a, 0xb3, 0x47, 0x4f, 0x22, 0xed, 0x17,
	0x91, 0xbc, 0x4e, 0x4a, 0x26, 0x99, 0xc0, 0x5e, 0x4a, 0xe1, 0xfd, 0xde, 0x2f, 0xff, 0x79, 0x6f,
	0x06, 0xbd, 0x66, 0x19, 0xa3, 0x4a, 0x92, 0x28, 0x67, 0x8a, 0xc9, 0x28, 0x61, 0x05, 0x2b, 0xd3,
	0x32, 0x14, 0x92, 0x2b, 0x8e, 0x9f, 0xe9, 0x62, 0x08, 0x45, 0xff, 0x45, 0xc2, 0x13, 0x0e, 0x95,
	0xa8, 0xfa, 0x77, 0x86, 0x7c, 0xdf, 0x34, 0x08, 0x22, 0x49, 0xae, 0x05, 0xfe, 0x7b, 0xb3, 0x06,
	0xbf, 0x92, 0x91, 0x75, 0x5a, 0x24, 0x35, 0x32, 0xb4, 0x20, 0xeb, 0x54, 0x32, 0xaa, 0xb8, 0xdc,
	0x69, 0xe6, 0x53, 0xeb, 0x13, 0xfc, 0x07, 0x93, 0x4b, 0xb1, 0x95, 0x74, 0x43, 0x4a, 0xb6, 0xa4,
	0xbc, 0x50, 0x92, 0x50, 0xa5, 0xe1, 0xd6, 0x89, 0x84, 0x20, 0xcb, 0x9c, 0x08, 0x7b, 0xa0, 0x55,
	0x9a, 0x65, 0x69, 0x91, 0xd0, 0x1d, 0xcd, 0x58, 0x1d, 0xe8, 0xa3, 0x89, 0xd0, 0x6d, 0xa9, 0x78,
	0xce, 0xa4, 0x46, 0xb3, 0xb4, 0x60, 0x76, 0x57, 0x13, 0xec, 0x71, 0x09, 0xc9, 0xd7, 0x5b, 0x7a,
	0x0f, 0x57, 0x13, 0xd4, 0xae, 0xe1, 0x6f, 0x0f, 0x3d, 0xfd, 0x72, 0x5e, 0xcf, 0x42, 0x11, 0xc5,
	0xf0, 0x04, 0x79, 0xe7, 0x61, 0xdf, 0xb9, 0x03, 0x77, 0x74, 0x33, 0xbe, 0x0d, 0x8d, 0x75, 0x85,
	0x31, 0x14, 0xa7, 0xd7, 0xfb, 0xbf, 0xef, 0x9c, 0x6f, 0x1a, 0xc5, 0x31, 0x7a, 0x6e, 0x6c, 0xe1,
	0x6b, 0x5a, 0xaa, 0xbb, 0x07, 0x83, 0xab, 0xd1, 0xcd, 0xf8, 0x4d, 0xab, 0x7f, 0xde, 0xe4, 0xb4,
	0xa6, 0xdb, 0x8c, 0x17, 0x08, 0x9b, 0x4b, 0x03, 0xe5, 0x15, 0x28, 0xdf, 0xda, 0x94, 0x17, 0x50,
	0x3b, 0x2d, 0xed, 0x78, 0x83, 0x5e, 0xc1, 0x96, 0x63, 0xbd, 0xe4, 0x99, 0xde, 0x31, 0xb8, 0xaf,
	0xc1, 0xfd, 0xa1, 0x7d, 0x5c, 0x1b, 0xaf, 0x3f, 0xd1, 0x2f, 0xc3, 0x9f, 0x11, 0x12, 0x82, 0xcc,
	0x89, 0x00, 0xf5, 0x43, 0x50, 0x77, 0x26, 0x09, 0x80, 0x76, 0x35, 0xf0, 0x6a, 0x9a, 0xc6, 0x15,
	0x02, 0x87, 0x67, 0x9d, 0xe6, 0xb4, 0xc9, 0xd5, 0xd3, 0xec, 0x34, 0xe3, 0x15, 0x7a, 0x69, 0xb9,
	0x71, 0xe0, 0x7d, 0x04, 0xde, 0x61, 0xcb, 0x3b, 0xeb, 0xd2, 0xda, 0xde, 0x27, 0xaa, 0x52, 0x1b,
	0x97, 0x15, 0xec, 0x8f, 0xad, 0xa9, 0x9b, 0xf6, 0x4b, 0xea, 0x4e, 0x73, 0x95, 0xda, 0x72, 0xb7,
	0xc1, 0xfb, 0xc4, 0x9a, 0x3a, 0xee, 0xd2, 0x75, 0xea, 0x1e, 0x51, 0x95, 0xda, 0x78, 0x16, 0x60,
	0x47, 0xd6, 0xd4, 0x4d, 0xfb, 0x25, 0x75, 0xa7, 0x79, 0x1a, 0xed, 0x8f, 0x81, 0x7b, 0x38, 0x06,
	0xee, 0xbf, 0x63, 0xe0, 0xfe, 0x3a, 0x05, 0xce, 0xe1, 0x14, 0x38, 0x7f, 0x4e, 0x81, 0xf3, 0xfd,
	0xb6, 0x7e, 0x8e, 0x3f, 0xf5, 0x83, 0x54, 0x3b, 0xc1, 0xca, 0x95, 0x07, 0x2f, 0x71, 0xf2, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0xe8, 0x64, 0x37, 0x35, 0x05, 0x00, 0x00,
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
	if len(m.ProducerbillsList) > 0 {
		for iNdEx := len(m.ProducerbillsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProducerbillsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.ProducerbillinglineList) > 0 {
		for iNdEx := len(m.ProducerbillinglineList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProducerbillinglineList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.CustomerbillsList) > 0 {
		for iNdEx := len(m.CustomerbillsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CustomerbillsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CustomerbillinglineList) > 0 {
		for iNdEx := len(m.CustomerbillinglineList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CustomerbillinglineList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.BillingcyclesList) > 0 {
		for iNdEx := len(m.BillingcyclesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BillingcyclesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PpaMapList) > 0 {
		for iNdEx := len(m.PpaMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PpaMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PowerPurchaseContractList) > 0 {
		for iNdEx := len(m.PowerPurchaseContractList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PowerPurchaseContractList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.MeterdirectoryList) > 0 {
		for iNdEx := len(m.MeterdirectoryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MeterdirectoryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.MeterreadingsList) > 0 {
		for iNdEx := len(m.MeterreadingsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MeterreadingsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.MeterreadingsList) > 0 {
		for _, e := range m.MeterreadingsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MeterdirectoryList) > 0 {
		for _, e := range m.MeterdirectoryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PowerPurchaseContractList) > 0 {
		for _, e := range m.PowerPurchaseContractList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PpaMapList) > 0 {
		for _, e := range m.PpaMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BillingcyclesList) > 0 {
		for _, e := range m.BillingcyclesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CustomerbillinglineList) > 0 {
		for _, e := range m.CustomerbillinglineList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CustomerbillsList) > 0 {
		for _, e := range m.CustomerbillsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProducerbillinglineList) > 0 {
		for _, e := range m.ProducerbillinglineList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProducerbillsList) > 0 {
		for _, e := range m.ProducerbillsList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field MeterreadingsList", wireType)
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
			m.MeterreadingsList = append(m.MeterreadingsList, Meterreadings{})
			if err := m.MeterreadingsList[len(m.MeterreadingsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeterdirectoryList", wireType)
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
			m.MeterdirectoryList = append(m.MeterdirectoryList, Meterdirectory{})
			if err := m.MeterdirectoryList[len(m.MeterdirectoryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PowerPurchaseContractList", wireType)
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
			m.PowerPurchaseContractList = append(m.PowerPurchaseContractList, PowerPurchaseContract{})
			if err := m.PowerPurchaseContractList[len(m.PowerPurchaseContractList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PpaMapList", wireType)
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
			m.PpaMapList = append(m.PpaMapList, PpaMap{})
			if err := m.PpaMapList[len(m.PpaMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillingcyclesList", wireType)
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
			m.BillingcyclesList = append(m.BillingcyclesList, Billingcycles{})
			if err := m.BillingcyclesList[len(m.BillingcyclesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerbillinglineList", wireType)
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
			m.CustomerbillinglineList = append(m.CustomerbillinglineList, Customerbillingline{})
			if err := m.CustomerbillinglineList[len(m.CustomerbillinglineList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerbillsList", wireType)
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
			m.CustomerbillsList = append(m.CustomerbillsList, Customerbills{})
			if err := m.CustomerbillsList[len(m.CustomerbillsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerbillinglineList", wireType)
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
			m.ProducerbillinglineList = append(m.ProducerbillinglineList, Producerbillingline{})
			if err := m.ProducerbillinglineList[len(m.ProducerbillinglineList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerbillsList", wireType)
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
			m.ProducerbillsList = append(m.ProducerbillsList, Producerbills{})
			if err := m.ProducerbillsList[len(m.ProducerbillsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
