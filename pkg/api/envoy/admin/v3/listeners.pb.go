// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v3/listeners.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
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

// Admin endpoint uses this wrapper for `/listeners` to display listener status information.
// See :ref:`/listeners <operations_admin_interface_listeners>` for more information.
type Listeners struct {
	// List of listener statuses.
	ListenerStatuses     []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Listeners) Reset()         { *m = Listeners{} }
func (m *Listeners) String() string { return proto.CompactTextString(m) }
func (*Listeners) ProtoMessage()    {}
func (*Listeners) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32da9886d1d0ff3, []int{0}
}
func (m *Listeners) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Listeners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Listeners.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Listeners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listeners.Merge(m, src)
}
func (m *Listeners) XXX_Size() int {
	return m.Size()
}
func (m *Listeners) XXX_DiscardUnknown() {
	xxx_messageInfo_Listeners.DiscardUnknown(m)
}

var xxx_messageInfo_Listeners proto.InternalMessageInfo

func (m *Listeners) GetListenerStatuses() []*ListenerStatus {
	if m != nil {
		return m.ListenerStatuses
	}
	return nil
}

// Details an individual listener's current status.
type ListenerStatus struct {
	// Name of the listener
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The actual local address that the listener is listening on. If a listener was configured
	// to listen on port 0, then this address has the port that was allocated by the OS.
	LocalAddress         *v3.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListenerStatus) Reset()         { *m = ListenerStatus{} }
func (m *ListenerStatus) String() string { return proto.CompactTextString(m) }
func (*ListenerStatus) ProtoMessage()    {}
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32da9886d1d0ff3, []int{1}
}
func (m *ListenerStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListenerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListenerStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListenerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerStatus.Merge(m, src)
}
func (m *ListenerStatus) XXX_Size() int {
	return m.Size()
}
func (m *ListenerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerStatus proto.InternalMessageInfo

func (m *ListenerStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenerStatus) GetLocalAddress() *v3.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*Listeners)(nil), "envoy.admin.v3.Listeners")
	proto.RegisterType((*ListenerStatus)(nil), "envoy.admin.v3.ListenerStatus")
}

func init() { proto.RegisterFile("envoy/admin/v3/listeners.proto", fileDescriptor_d32da9886d1d0ff3) }

var fileDescriptor_d32da9886d1d0ff3 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x65, 0x40, 0x48, 0x75, 0xa1, 0x82, 0x4c, 0x51, 0x45, 0x4d, 0x88, 0x3a, 0x84, 0xc5,
	0x96, 0x92, 0xad, 0x62, 0x21, 0x2b, 0x0c, 0x55, 0x38, 0x40, 0x65, 0x12, 0x53, 0x2c, 0xa5, 0x76,
	0x64, 0xbb, 0x11, 0x5d, 0x98, 0xd9, 0xd9, 0xb8, 0x0f, 0x12, 0x23, 0x47, 0x40, 0x39, 0x09, 0x8a,
	0x9d, 0x54, 0x0a, 0x03, 0x9b, 0x7f, 0xfd, 0xef, 0xf9, 0x7d, 0xff, 0x83, 0x88, 0x89, 0x5a, 0xee,
	0x08, 0x2d, 0x36, 0x5c, 0x90, 0x3a, 0x21, 0x25, 0xd7, 0x86, 0x09, 0xa6, 0x34, 0xae, 0x94, 0x34,
	0xd2, 0x9b, 0xd8, 0x3d, 0xb6, 0x7b, 0x5c, 0x27, 0xd3, 0xd0, 0xe9, 0x73, 0x29, 0x9e, 0xf8, 0x9a,
	0xe4, 0x52, 0xb1, 0xd6, 0x45, 0x8b, 0x42, 0x31, 0xdd, 0x79, 0xa6, 0x57, 0xdb, 0xa2, 0xa2, 0x84,
	0x0a, 0x21, 0x0d, 0x35, 0x5c, 0x0a, 0x4d, 0x6a, 0xa6, 0x34, 0x97, 0x82, 0x8b, 0xb5, 0x93, 0x84,
	0xaf, 0x70, 0x74, 0xdf, 0x27, 0x79, 0x77, 0xf0, 0xbc, 0x8f, 0x5d, 0x69, 0x43, 0xcd, 0x56, 0x33,
	0xed, 0x83, 0xe0, 0x30, 0x1a, 0xc7, 0x08, 0x0f, 0xf3, 0x71, 0xef, 0x7a, 0xb0, 0xba, 0xec, 0xac,
	0x1c, 0xcc, 0x4c, 0x2f, 0xe6, 0x1f, 0x9f, 0x6f, 0xe8, 0x12, 0xce, 0x06, 0xbe, 0x98, 0x96, 0xd5,
	0x33, 0xdd, 0x9b, 0x75, 0xf8, 0x0e, 0xe0, 0x64, 0xf8, 0x95, 0xe7, 0xc1, 0x23, 0x41, 0x37, 0xcc,
	0x07, 0x01, 0x88, 0x46, 0x99, 0x7d, 0x7b, 0x29, 0x3c, 0x2d, 0x65, 0x4e, 0xcb, 0x55, 0x77, 0xa0,
	0x7f, 0x10, 0x80, 0x68, 0x1c, 0xcf, 0x3a, 0x2a, 0xd7, 0x02, 0x6e, 0x5b, 0x68, 0xd9, 0x6e, 0x9d,
	0x28, 0x3b, 0xb1, 0x9e, 0x6e, 0x5a, 0x5c, 0xb7, 0x40, 0x73, 0x18, 0xfe, 0x07, 0xe4, 0x10, 0xd2,
	0x9b, 0xaf, 0x06, 0x81, 0xef, 0x06, 0x81, 0x9f, 0x06, 0x01, 0x78, 0xc1, 0xa5, 0xcb, 0xa9, 0x94,
	0x7c, 0xd9, 0xfd, 0x29, 0x22, 0xdd, 0xe3, 0xeb, 0x65, 0xdb, 0xe8, 0x12, 0x3c, 0x1e, 0xdb, 0x6a,
	0x93, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xc2, 0xf5, 0x76, 0xd3, 0x01, 0x00, 0x00,
}

func (m *Listeners) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listeners) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Listeners) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ListenerStatuses) > 0 {
		for iNdEx := len(m.ListenerStatuses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ListenerStatuses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintListeners(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListenerStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListenerStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListenerStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LocalAddress != nil {
		{
			size, err := m.LocalAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListeners(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintListeners(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintListeners(dAtA []byte, offset int, v uint64) int {
	offset -= sovListeners(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Listeners) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ListenerStatuses) > 0 {
		for _, e := range m.ListenerStatuses {
			l = e.Size()
			n += 1 + l + sovListeners(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListenerStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovListeners(uint64(l))
	}
	if m.LocalAddress != nil {
		l = m.LocalAddress.Size()
		n += 1 + l + sovListeners(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovListeners(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListeners(x uint64) (n int) {
	return sovListeners(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Listeners) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListeners
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
			return fmt.Errorf("proto: Listeners: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Listeners: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenerStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
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
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenerStatuses = append(m.ListenerStatuses, &ListenerStatus{})
			if err := m.ListenerStatuses[len(m.ListenerStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListeners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthListeners
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
func (m *ListenerStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListeners
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
			return fmt.Errorf("proto: ListenerStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenerStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
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
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
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
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LocalAddress == nil {
				m.LocalAddress = &v3.Address{}
			}
			if err := m.LocalAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListeners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthListeners
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
func skipListeners(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListeners
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
					return 0, ErrIntOverflowListeners
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowListeners
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
				return 0, ErrInvalidLengthListeners
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthListeners
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowListeners
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipListeners(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthListeners
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthListeners = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListeners   = fmt.Errorf("proto: integer overflow")
)