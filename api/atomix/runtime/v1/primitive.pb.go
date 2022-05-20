// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/runtime/v1/primitive.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
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

type Primitive struct {
	PrimitiveMeta `protobuf:"bytes,1,opt,name=meta,proto3,embedded=meta" json:"meta"`
	Spec          PrimitiveSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec"`
}

func (m *Primitive) Reset()         { *m = Primitive{} }
func (m *Primitive) String() string { return proto.CompactTextString(m) }
func (*Primitive) ProtoMessage()    {}
func (*Primitive) Descriptor() ([]byte, []int) {
	return fileDescriptor_79864d6f09561b9d, []int{0}
}
func (m *Primitive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Primitive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Primitive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Primitive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Primitive.Merge(m, src)
}
func (m *Primitive) XXX_Size() int {
	return m.Size()
}
func (m *Primitive) XXX_DiscardUnknown() {
	xxx_messageInfo_Primitive.DiscardUnknown(m)
}

var xxx_messageInfo_Primitive proto.InternalMessageInfo

func (m *Primitive) GetSpec() PrimitiveSpec {
	if m != nil {
		return m.Spec
	}
	return PrimitiveSpec{}
}

type PrimitiveId struct {
	Type PrimitiveType `protobuf:"bytes,1,opt,name=type,proto3,casttype=PrimitiveType" json:"type,omitempty"`
	Name string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *PrimitiveId) Reset()         { *m = PrimitiveId{} }
func (m *PrimitiveId) String() string { return proto.CompactTextString(m) }
func (*PrimitiveId) ProtoMessage()    {}
func (*PrimitiveId) Descriptor() ([]byte, []int) {
	return fileDescriptor_79864d6f09561b9d, []int{1}
}
func (m *PrimitiveId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimitiveId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimitiveId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimitiveId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveId.Merge(m, src)
}
func (m *PrimitiveId) XXX_Size() int {
	return m.Size()
}
func (m *PrimitiveId) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveId.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveId proto.InternalMessageInfo

func (m *PrimitiveId) GetType() PrimitiveType {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PrimitiveId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PrimitiveMeta struct {
	PrimitiveID PrimitiveId       `protobuf:"bytes,1,opt,name=primitive_id,json=primitiveId,proto3" json:"primitive_id"`
	Labels      map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *PrimitiveMeta) Reset()         { *m = PrimitiveMeta{} }
func (m *PrimitiveMeta) String() string { return proto.CompactTextString(m) }
func (*PrimitiveMeta) ProtoMessage()    {}
func (*PrimitiveMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_79864d6f09561b9d, []int{2}
}
func (m *PrimitiveMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimitiveMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimitiveMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimitiveMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveMeta.Merge(m, src)
}
func (m *PrimitiveMeta) XXX_Size() int {
	return m.Size()
}
func (m *PrimitiveMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveMeta.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveMeta proto.InternalMessageInfo

func (m *PrimitiveMeta) GetPrimitiveID() PrimitiveId {
	if m != nil {
		return m.PrimitiveID
	}
	return PrimitiveId{}
}

func (m *PrimitiveMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type PrimitiveSpec struct {
	Config *types.Any `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *PrimitiveSpec) Reset()         { *m = PrimitiveSpec{} }
func (m *PrimitiveSpec) String() string { return proto.CompactTextString(m) }
func (*PrimitiveSpec) ProtoMessage()    {}
func (*PrimitiveSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_79864d6f09561b9d, []int{3}
}
func (m *PrimitiveSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimitiveSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimitiveSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimitiveSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveSpec.Merge(m, src)
}
func (m *PrimitiveSpec) XXX_Size() int {
	return m.Size()
}
func (m *PrimitiveSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveSpec proto.InternalMessageInfo

func (m *PrimitiveSpec) GetConfig() *types.Any {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Primitive)(nil), "atomix.runtime.v1.Primitive")
	proto.RegisterType((*PrimitiveId)(nil), "atomix.runtime.v1.PrimitiveId")
	proto.RegisterType((*PrimitiveMeta)(nil), "atomix.runtime.v1.PrimitiveMeta")
	proto.RegisterMapType((map[string]string)(nil), "atomix.runtime.v1.PrimitiveMeta.LabelsEntry")
	proto.RegisterType((*PrimitiveSpec)(nil), "atomix.runtime.v1.PrimitiveSpec")
}

func init() { proto.RegisterFile("atomix/runtime/v1/primitive.proto", fileDescriptor_79864d6f09561b9d) }

var fileDescriptor_79864d6f09561b9d = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0xcb, 0xd3, 0x40,
	0x10, 0xcd, 0xb6, 0xb1, 0xd8, 0x8d, 0x85, 0x76, 0xed, 0x21, 0xf6, 0x90, 0xd4, 0x80, 0xd0, 0x43,
	0xd9, 0xd0, 0x7a, 0xd1, 0x82, 0x82, 0xa1, 0x82, 0x05, 0x05, 0x59, 0xc5, 0xab, 0xa4, 0xe9, 0x36,
	0x2c, 0x26, 0xd9, 0x90, 0x6e, 0x83, 0xf9, 0x05, 0x5e, 0xfd, 0x59, 0x3d, 0xf6, 0xe8, 0x29, 0x48,
	0x7a, 0xf1, 0x37, 0x78, 0x92, 0x6c, 0xd2, 0x18, 0x11, 0xbe, 0xde, 0xde, 0xee, 0xbc, 0x37, 0xf3,
	0xde, 0x30, 0xf0, 0xb1, 0x2b, 0x78, 0xc8, 0xbe, 0xda, 0xc9, 0x31, 0x12, 0x2c, 0xa4, 0x76, 0xba,
	0xb0, 0xe3, 0x84, 0x85, 0x4c, 0xb0, 0x94, 0xe2, 0x38, 0xe1, 0x82, 0xa3, 0x51, 0x45, 0xc1, 0x35,
	0x05, 0xa7, 0x8b, 0xc9, 0x23, 0x9f, 0x73, 0x3f, 0xa0, 0xb6, 0x24, 0x6c, 0x8f, 0x7b, 0xdb, 0x8d,
	0xb2, 0x8a, 0x3d, 0x19, 0xfb, 0xdc, 0xe7, 0x12, 0xda, 0x25, 0xaa, 0x7e, 0xad, 0x6f, 0x00, 0xf6,
	0xdf, 0x5f, 0xfb, 0xa2, 0x97, 0x50, 0x0d, 0xa9, 0x70, 0x75, 0x30, 0x05, 0x33, 0x6d, 0x39, 0xc5,
	0xff, 0x0d, 0xc0, 0x0d, 0xf7, 0x1d, 0x15, 0xae, 0x73, 0xff, 0x94, 0x9b, 0xca, 0x39, 0x37, 0x01,
	0x91, 0x3a, 0xb4, 0x82, 0xea, 0x21, 0xa6, 0x9e, 0xde, 0xb9, 0xad, 0xff, 0x10, 0x53, 0xcf, 0x51,
	0x4b, 0x3d, 0x91, 0x1a, 0xeb, 0x0d, 0xd4, 0x9a, 0xe2, 0x66, 0x87, 0x9e, 0x40, 0x55, 0x64, 0x31,
	0x95, 0x56, 0xfa, 0xce, 0xe8, 0x77, 0x6e, 0x0e, 0x9a, 0xf2, 0xc7, 0x2c, 0xa6, 0x44, 0x96, 0x11,
	0x82, 0x6a, 0xe4, 0x86, 0x54, 0x4e, 0xec, 0x13, 0x89, 0xad, 0x5f, 0x00, 0x0e, 0xfe, 0xf1, 0x89,
	0x3e, 0xc1, 0x07, 0xcd, 0xf2, 0x3e, 0xb3, 0x5d, 0x9d, 0xcf, 0xb8, 0xcb, 0xdf, 0x66, 0xe7, 0x3c,
	0x2c, 0xdd, 0x15, 0xb9, 0xd9, 0xf2, 0xb5, 0x26, 0x5a, 0xdc, 0x32, 0xb9, 0x86, 0xbd, 0xc0, 0xdd,
	0xd2, 0xe0, 0xa0, 0x77, 0xa6, 0xdd, 0x99, 0xb6, 0x9c, 0xdf, 0xda, 0x18, 0x7e, 0x2b, 0xe9, 0xaf,
	0x23, 0x91, 0x64, 0xa4, 0xd6, 0x4e, 0x9e, 0x43, 0xad, 0xf5, 0x8d, 0x86, 0xb0, 0xfb, 0x85, 0x66,
	0x55, 0x70, 0x52, 0x42, 0x34, 0x86, 0xf7, 0x52, 0x37, 0x38, 0x5e, 0x53, 0x56, 0x8f, 0x55, 0xe7,
	0x19, 0xb0, 0x5e, 0xb4, 0x92, 0x96, 0x1b, 0x45, 0x73, 0xd8, 0xf3, 0x78, 0xb4, 0x67, 0x7e, 0x9d,
	0x71, 0x8c, 0xab, 0x8b, 0xc0, 0xd7, 0x8b, 0xc0, 0xaf, 0xa2, 0x8c, 0xd4, 0x9c, 0x25, 0x82, 0xc3,
	0xbf, 0x72, 0x9a, 0xa4, 0xcc, 0xa3, 0x8e, 0x7e, 0x2a, 0x0c, 0x70, 0x2e, 0x0c, 0xf0, 0xb3, 0x30,
	0xc0, 0xf7, 0x8b, 0xa1, 0x9c, 0x2f, 0x86, 0xf2, 0xe3, 0x62, 0x28, 0xdb, 0x9e, 0xec, 0xf1, 0xf4,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xd3, 0x0f, 0x03, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrimitiveServiceClient is the client API for PrimitiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrimitiveServiceClient interface {
}

type primitiveServiceClient struct {
	cc *grpc.ClientConn
}

func NewPrimitiveServiceClient(cc *grpc.ClientConn) PrimitiveServiceClient {
	return &primitiveServiceClient{cc}
}

// PrimitiveServiceServer is the server API for PrimitiveService service.
type PrimitiveServiceServer interface {
}

// UnimplementedPrimitiveServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrimitiveServiceServer struct {
}

func RegisterPrimitiveServiceServer(s *grpc.Server, srv PrimitiveServiceServer) {
	s.RegisterService(&_PrimitiveService_serviceDesc, srv)
}

var _PrimitiveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.runtime.v1.PrimitiveService",
	HandlerType: (*PrimitiveServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "atomix/runtime/v1/primitive.proto",
}

func (m *Primitive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Primitive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Primitive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPrimitive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.PrimitiveMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPrimitive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PrimitiveId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimitiveId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimitiveId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PrimitiveMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimitiveMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimitiveMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintPrimitive(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintPrimitive(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintPrimitive(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.PrimitiveID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPrimitive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PrimitiveSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimitiveSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimitiveSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPrimitive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrimitive(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrimitive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Primitive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PrimitiveMeta.Size()
	n += 1 + l + sovPrimitive(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovPrimitive(uint64(l))
	return n
}

func (m *PrimitiveId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	return n
}

func (m *PrimitiveMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PrimitiveID.Size()
	n += 1 + l + sovPrimitive(uint64(l))
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovPrimitive(uint64(len(k))) + 1 + len(v) + sovPrimitive(uint64(len(v)))
			n += mapEntrySize + 1 + sovPrimitive(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *PrimitiveSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovPrimitive(uint64(l))
	}
	return n
}

func sovPrimitive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrimitive(x uint64) (n int) {
	return sovPrimitive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Primitive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
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
			return fmt.Errorf("proto: Primitive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Primitive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimitiveMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrimitiveMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *PrimitiveId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
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
			return fmt.Errorf("proto: PrimitiveId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimitiveId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = PrimitiveType(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *PrimitiveMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
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
			return fmt.Errorf("proto: PrimitiveMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimitiveMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimitiveID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrimitiveID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPrimitive
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPrimitive
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthPrimitive
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthPrimitive
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPrimitive
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthPrimitive
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthPrimitive
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPrimitive(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthPrimitive
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *PrimitiveSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
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
			return fmt.Errorf("proto: PrimitiveSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimitiveSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &types.Any{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func skipPrimitive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrimitive
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
					return 0, ErrIntOverflowPrimitive
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
					return 0, ErrIntOverflowPrimitive
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
				return 0, ErrInvalidLengthPrimitive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrimitive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrimitive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrimitive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrimitive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrimitive = fmt.Errorf("proto: unexpected end of group")
)