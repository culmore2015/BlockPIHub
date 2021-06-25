// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bpihub/provider.proto

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

type Provider struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Details string `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	Website string `protobuf:"bytes,5,opt,name=website,proto3" json:"website,omitempty"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b82431537075bef, []int{0}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return m.Size()
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Provider) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Provider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Provider) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Provider) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func init() {
	proto.RegisterType((*Provider)(nil), "madtechworks.bpihub.bpihub.Provider")
}

func init() { proto.RegisterFile("bpihub/provider.proto", fileDescriptor_1b82431537075bef) }

var fileDescriptor_1b82431537075bef = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2a, 0xc8, 0xcc,
	0x28, 0x4d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0xca, 0x4d, 0x4c, 0x29, 0x49, 0x4d, 0xce, 0x28, 0xcf, 0x2f, 0xca, 0x2e, 0xd6,
	0x83, 0xa8, 0x81, 0x52, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x65, 0xfa, 0x20, 0x16, 0x44,
	0x87, 0x52, 0x0d, 0x17, 0x47, 0x00, 0xd4, 0x0c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4,
	0x92, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x88, 0x8f, 0x8b, 0x29,
	0x33, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x25, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25,
	0x2f, 0x31, 0x37, 0x55, 0x82, 0x19, 0xac, 0x0c, 0xcc, 0x06, 0xe9, 0x4e, 0x49, 0x2d, 0x49, 0xcc,
	0xcc, 0x29, 0x96, 0x60, 0x81, 0xe8, 0x86, 0x72, 0x41, 0x32, 0xe5, 0xa9, 0x49, 0xc5, 0x99, 0x25,
	0xa9, 0x12, 0xac, 0x10, 0x19, 0x28, 0xd7, 0xc9, 0xfd, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33, 0x4b, 0x40, 0x8e, 0x4f, 0xce, 0xcf, 0xd5, 0x47, 0xf6, 0x94,
	0x3e, 0xd4, 0xe3, 0x15, 0x30, 0x46, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x37, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x99, 0xec, 0x3c, 0x18, 0x01, 0x00, 0x00,
}

func (m *Provider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Provider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Provider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintProvider(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProvider(dAtA []byte, offset int, v uint64) int {
	offset -= sovProvider(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Provider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovProvider(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	return n
}

func sovProvider(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProvider(x uint64) (n int) {
	return sovProvider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Provider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: Provider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Provider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func skipProvider(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
				return 0, ErrInvalidLengthProvider
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProvider
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProvider
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProvider        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProvider          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProvider = fmt.Errorf("proto: unexpected end of group")
)
