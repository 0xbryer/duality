// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/deposit_record.proto

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

type DepositRecord struct {
	PairId          string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	SharesOwned     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalShares" yaml:"totalShares"`
	CenterTickIndex int64                                  `protobuf:"varint,3,opt,name=centerTickIndex,proto3" json:"centerTickIndex,omitempty"`
	LowerTickIndex  int64                                  `protobuf:"varint,4,opt,name=lowerTickIndex,proto3" json:"lowerTickIndex,omitempty"`
	UpperTickIndex  int64                                  `protobuf:"varint,5,opt,name=upperTickIndex,proto3" json:"upperTickIndex,omitempty"`
	FeeIndex        uint64                                 `protobuf:"varint,6,opt,name=feeIndex,proto3" json:"feeIndex,omitempty"`
}

func (m *DepositRecord) Reset()         { *m = DepositRecord{} }
func (m *DepositRecord) String() string { return proto.CompactTextString(m) }
func (*DepositRecord) ProtoMessage()    {}
func (*DepositRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d7c76d21a5add72, []int{0}
}
func (m *DepositRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositRecord.Merge(m, src)
}
func (m *DepositRecord) XXX_Size() int {
	return m.Size()
}
func (m *DepositRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DepositRecord proto.InternalMessageInfo

func (m *DepositRecord) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *DepositRecord) GetCenterTickIndex() int64 {
	if m != nil {
		return m.CenterTickIndex
	}
	return 0
}

func (m *DepositRecord) GetLowerTickIndex() int64 {
	if m != nil {
		return m.LowerTickIndex
	}
	return 0
}

func (m *DepositRecord) GetUpperTickIndex() int64 {
	if m != nil {
		return m.UpperTickIndex
	}
	return 0
}

func (m *DepositRecord) GetFeeIndex() uint64 {
	if m != nil {
		return m.FeeIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*DepositRecord)(nil), "nicholasdotsol.duality.dex.DepositRecord")
}

func init() { proto.RegisterFile("duality/dex/deposit_record.proto", fileDescriptor_9d7c76d21a5add72) }

var fileDescriptor_9d7c76d21a5add72 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xb1, 0x4a, 0xc3, 0x40,
	0x1c, 0xc6, 0x73, 0x6d, 0x2d, 0x7a, 0x45, 0x85, 0x43, 0xa4, 0x74, 0x48, 0x42, 0x87, 0x92, 0xa5,
	0xc9, 0xe0, 0xe6, 0x58, 0x0a, 0xd2, 0x45, 0x21, 0x75, 0x72, 0x91, 0x34, 0x77, 0xb6, 0xa1, 0xd7,
	0xfc, 0xc3, 0xdd, 0x85, 0xa6, 0x6f, 0xe1, 0xc3, 0xf8, 0x10, 0x1d, 0x3b, 0x8a, 0x43, 0x90, 0x76,
	0x73, 0xf4, 0x09, 0x24, 0x97, 0x28, 0xb1, 0xd3, 0xdd, 0xff, 0xbb, 0xdf, 0xff, 0x3b, 0xf8, 0x3e,
	0x6c, 0xd3, 0x34, 0xe0, 0x91, 0xda, 0x78, 0x94, 0x65, 0x1e, 0x65, 0x09, 0xc8, 0x48, 0x3d, 0x0b,
	0x16, 0x82, 0xa0, 0x6e, 0x22, 0x40, 0x01, 0xe9, 0xc5, 0x51, 0xb8, 0x00, 0x1e, 0x48, 0x0a, 0x4a,
	0x02, 0x77, 0xab, 0x05, 0x97, 0xb2, 0xac, 0x77, 0x35, 0x87, 0x39, 0x68, 0xcc, 0x2b, 0x6e, 0xe5,
	0x46, 0xff, 0xad, 0x81, 0xcf, 0xc7, 0xa5, 0x95, 0xaf, 0x9d, 0xc8, 0x35, 0x6e, 0x27, 0x41, 0x24,
	0x26, 0xb4, 0x8b, 0x6c, 0xe4, 0x9c, 0xf9, 0xd5, 0x44, 0x52, 0xdc, 0x91, 0x8b, 0x40, 0x30, 0xf9,
	0xb0, 0x8e, 0x19, 0xed, 0x36, 0x8a, 0xc7, 0xd1, 0x74, 0x9b, 0x5b, 0xc6, 0x47, 0x6e, 0x0d, 0xe6,
	0x91, 0x5a, 0xa4, 0x33, 0x37, 0x84, 0x95, 0x17, 0x82, 0x5c, 0x81, 0xac, 0x8e, 0xa1, 0xa4, 0x4b,
	0x4f, 0x6d, 0x12, 0x26, 0xdd, 0x49, 0xac, 0xbe, 0x72, 0xab, 0xa3, 0x40, 0x05, 0x7c, 0xaa, 0x9d,
	0xbe, 0x73, 0x8b, 0x6c, 0x82, 0x15, 0xbf, 0xed, 0xd7, 0xc4, 0xbe, 0x5f, 0xff, 0x87, 0x38, 0xf8,
	0x32, 0x64, 0xb1, 0x62, 0xe2, 0x31, 0x0a, 0x97, 0x93, 0x98, 0xb2, 0xac, 0xdb, 0xb4, 0x91, 0xd3,
	0xf4, 0x8f, 0x65, 0x32, 0xc0, 0x17, 0x1c, 0xd6, 0x75, 0xb0, 0xa5, 0xc1, 0x23, 0xb5, 0xe0, 0xd2,
	0x24, 0xa9, 0x73, 0x27, 0x25, 0xf7, 0x5f, 0x25, 0x3d, 0x7c, 0xfa, 0xc2, 0x58, 0x49, 0xb4, 0x6d,
	0xe4, 0xb4, 0xfc, 0xbf, 0x79, 0x74, 0xb7, 0xdd, 0x9b, 0x68, 0xb7, 0x37, 0xd1, 0xe7, 0xde, 0x44,
	0xaf, 0x07, 0xd3, 0xd8, 0x1d, 0x4c, 0xe3, 0xfd, 0x60, 0x1a, 0x4f, 0xc3, 0x5a, 0x12, 0xf7, 0x55,
	0x1b, 0x63, 0x50, 0x53, 0xe0, 0xde, 0x6f, 0x7d, 0x99, 0x2e, 0x50, 0x87, 0x32, 0x6b, 0xeb, 0x1a,
	0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x02, 0x49, 0x05, 0xdc, 0x01, 0x00, 0x00,
}

func (m *DepositRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeIndex != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.FeeIndex))
		i--
		dAtA[i] = 0x30
	}
	if m.UpperTickIndex != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.UpperTickIndex))
		i--
		dAtA[i] = 0x28
	}
	if m.LowerTickIndex != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.LowerTickIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.CenterTickIndex != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.CenterTickIndex))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDepositRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintDepositRecord(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDepositRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovDepositRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovDepositRecord(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovDepositRecord(uint64(l))
	if m.CenterTickIndex != 0 {
		n += 1 + sovDepositRecord(uint64(m.CenterTickIndex))
	}
	if m.LowerTickIndex != 0 {
		n += 1 + sovDepositRecord(uint64(m.LowerTickIndex))
	}
	if m.UpperTickIndex != 0 {
		n += 1 + sovDepositRecord(uint64(m.UpperTickIndex))
	}
	if m.FeeIndex != 0 {
		n += 1 + sovDepositRecord(uint64(m.FeeIndex))
	}
	return n
}

func sovDepositRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDepositRecord(x uint64) (n int) {
	return sovDepositRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDepositRecord
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
			return fmt.Errorf("proto: DepositRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesOwned", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CenterTickIndex", wireType)
			}
			m.CenterTickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CenterTickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowerTickIndex", wireType)
			}
			m.LowerTickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LowerTickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpperTickIndex", wireType)
			}
			m.UpperTickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpperTickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeIndex", wireType)
			}
			m.FeeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDepositRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDepositRecord
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
func skipDepositRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
				return 0, ErrInvalidLengthDepositRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDepositRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDepositRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDepositRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDepositRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDepositRecord = fmt.Errorf("proto: unexpected end of group")
)