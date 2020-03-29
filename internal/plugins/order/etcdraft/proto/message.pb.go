// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_meshplus_bitxhub_kit_types "github.com/meshplus/bitxhub-kit/types"
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

type RaftMessage_Type int32

const (
	RaftMessage_CONSENSUS    RaftMessage_Type = 0
	RaftMessage_BROADCAST_TX RaftMessage_Type = 1
	RaftMessage_GET_TX       RaftMessage_Type = 2
	RaftMessage_GET_TX_ACK   RaftMessage_Type = 3
)

var RaftMessage_Type_name = map[int32]string{
	0: "CONSENSUS",
	1: "BROADCAST_TX",
	2: "GET_TX",
	3: "GET_TX_ACK",
}

var RaftMessage_Type_value = map[string]int32{
	"CONSENSUS":    0,
	"BROADCAST_TX": 1,
	"GET_TX":       2,
	"GET_TX_ACK":   3,
}

func (x RaftMessage_Type) String() string {
	return proto.EnumName(RaftMessage_Type_name, int32(x))
}

func (RaftMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type RaftMessage struct {
	Type                 RaftMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=proto.RaftMessage_Type" json:"type,omitempty"`
	FromId               uint64           `protobuf:"varint,2,opt,name=fromId,proto3" json:"fromId,omitempty"`
	Data                 []byte           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RaftMessage) Reset()         { *m = RaftMessage{} }
func (m *RaftMessage) String() string { return proto.CompactTextString(m) }
func (*RaftMessage) ProtoMessage()    {}
func (*RaftMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *RaftMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaftMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaftMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaftMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftMessage.Merge(m, src)
}
func (m *RaftMessage) XXX_Size() int {
	return m.Size()
}
func (m *RaftMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RaftMessage proto.InternalMessageInfo

func (m *RaftMessage) GetType() RaftMessage_Type {
	if m != nil {
		return m.Type
	}
	return RaftMessage_CONSENSUS
}

func (m *RaftMessage) GetFromId() uint64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *RaftMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Ready struct {
	TxHashes             []github_com_meshplus_bitxhub_kit_types.Hash `protobuf:"bytes,1,rep,name=txHashes,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"txHashes"`
	Height               uint64                                       `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *Ready) Reset()         { *m = Ready{} }
func (m *Ready) String() string { return proto.CompactTextString(m) }
func (*Ready) ProtoMessage()    {}
func (*Ready) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}
func (m *Ready) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ready) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ready.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ready) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ready.Merge(m, src)
}
func (m *Ready) XXX_Size() int {
	return m.Size()
}
func (m *Ready) XXX_DiscardUnknown() {
	xxx_messageInfo_Ready.DiscardUnknown(m)
}

var xxx_messageInfo_Ready proto.InternalMessageInfo

func (m *Ready) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.RaftMessage_Type", RaftMessage_Type_name, RaftMessage_Type_value)
	proto.RegisterType((*RaftMessage)(nil), "proto.RaftMessage")
	proto.RegisterType((*Ready)(nil), "proto.Ready")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x97, 0xad, 0x1b, 0xfa, 0xb7, 0x1b, 0x25, 0x07, 0x2d, 0x1e, 0xb6, 0xb2, 0x53, 0x51,
	0xd6, 0xc2, 0xfc, 0x04, 0x5b, 0x1d, 0x2a, 0xe2, 0x06, 0xe9, 0x04, 0x6f, 0x23, 0x75, 0x59, 0x53,
	0xb4, 0xa4, 0x2c, 0x29, 0xac, 0x9f, 0xc9, 0x2f, 0xb2, 0xa3, 0x67, 0x0f, 0x43, 0xfa, 0x49, 0xa4,
	0xa9, 0xca, 0x4e, 0x79, 0xbf, 0xf0, 0xde, 0x23, 0x2f, 0xd0, 0x4d, 0x99, 0x94, 0x34, 0x66, 0x5e,
	0xb6, 0x15, 0x4a, 0xe0, 0xb6, 0x3e, 0x2e, 0x47, 0x71, 0xa2, 0x78, 0x1e, 0x79, 0xaf, 0x22, 0xf5,
	0x63, 0x11, 0x0b, 0x5f, 0x5f, 0x47, 0xf9, 0x46, 0x93, 0x06, 0xad, 0xea, 0xd4, 0xf0, 0x03, 0xc1,
	0x19, 0xa1, 0x1b, 0xf5, 0x54, 0x77, 0xe1, 0x6b, 0x30, 0x54, 0x91, 0x31, 0x1b, 0x39, 0xc8, 0xed,
	0x8d, 0x2f, 0x6a, 0x97, 0x77, 0xe4, 0xf0, 0x96, 0x45, 0xc6, 0x88, 0x36, 0xe1, 0x73, 0xe8, 0x6c,
	0xb6, 0x22, 0x7d, 0x58, 0xdb, 0x4d, 0x07, 0xb9, 0x06, 0xf9, 0x25, 0x8c, 0xc1, 0x58, 0x53, 0x45,
	0xed, 0x96, 0x83, 0x5c, 0x93, 0x68, 0x3d, 0x0c, 0xc0, 0xa8, 0x92, 0xb8, 0x0b, 0xa7, 0xc1, 0x62,
	0x1e, 0xce, 0xe6, 0xe1, 0x73, 0x68, 0x35, 0xb0, 0x05, 0xe6, 0x94, 0x2c, 0x26, 0xb7, 0xc1, 0x24,
	0x5c, 0xae, 0x96, 0x2f, 0x16, 0xc2, 0x00, 0x9d, 0xbb, 0x99, 0xd6, 0x4d, 0xdc, 0x03, 0xa8, 0xf5,
	0x6a, 0x12, 0x3c, 0x5a, 0xad, 0xa1, 0x80, 0x36, 0x61, 0x74, 0x5d, 0xe0, 0x39, 0x9c, 0xa8, 0xdd,
	0x3d, 0x95, 0x9c, 0x49, 0x1b, 0x39, 0x2d, 0xd7, 0x9c, 0x8e, 0xf7, 0x87, 0x41, 0xe3, 0xeb, 0x30,
	0xb8, 0x3a, 0xda, 0x9f, 0x32, 0xc9, 0xb3, 0xf7, 0x5c, 0xfa, 0x51, 0xa2, 0x76, 0x3c, 0x8f, 0x46,
	0x6f, 0x89, 0xf2, 0xab, 0x97, 0x4b, 0xaf, 0xca, 0x92, 0xff, 0x8e, 0x6a, 0x09, 0x67, 0x49, 0xcc,
	0xd5, 0xdf, 0x92, 0x9a, 0xa6, 0xe6, 0xbe, 0xec, 0xa3, 0xcf, 0xb2, 0x8f, 0xbe, 0xcb, 0x3e, 0x8a,
	0x3a, 0xfa, 0x37, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xb0, 0xe7, 0x45, 0x7a, 0x01,
	0x00, 0x00,
}

func (m *RaftMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaftMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.FromId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.FromId))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Ready) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ready) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ready) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Height != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TxHashes) > 0 {
		for iNdEx := len(m.TxHashes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.TxHashes[iNdEx].Size()
				i -= size
				if _, err := m.TxHashes[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RaftMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMessage(uint64(m.Type))
	}
	if m.FromId != 0 {
		n += 1 + sovMessage(uint64(m.FromId))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Ready) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TxHashes) > 0 {
		for _, e := range m.TxHashes {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: RaftMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= RaftMessage_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromId", wireType)
			}
			m.FromId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ready) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Ready: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ready: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_meshplus_bitxhub_kit_types.Hash
			m.TxHashes = append(m.TxHashes, v)
			if err := m.TxHashes[len(m.TxHashes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
