// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: article.proto

package model

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

// Article
type Article struct {
	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement;comment:id"`
	// aid
	Aid string `protobuf:"bytes,2,opt,name=aid,proto3" json:"aid,omitempty" gorm:"unique;size:128;comment:article id"`
	//  pid
	Pid string `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty" gorm:"index;size:128;not null;default:null;comment:product pid"`
	// title
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty" gorm:"comment:title"`
	// cover
	Cover string `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty" gorm:"comment:cover"`
	// raw
	Raw []byte `protobuf:"bytes,7,opt,name=raw,proto3" json:"raw,omitempty" gorm:"comment:raw"`
	// status
	Status int32 `protobuf:"varint,19,opt,name=status,proto3" json:"status,omitempty" gorm:"default:1;size:2;comment:status 1 normal"`
	// created_at
	CreatedAt int64 `protobuf:"varint,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"index;comment:created at"`
	// updated_at
	UpdatedAt int64 `protobuf:"varint,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"comment:updated at"`
	// deleted_at
	DeletedAt int64 `protobuf:"varint,22,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty" gorm:"comment:deleted at"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}
func (m *Article) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Article.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return m.Size()
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetAid() string {
	if m != nil {
		return m.Aid
	}
	return ""
}

func (m *Article) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *Article) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Article) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Article) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Article) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Article) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Article)(nil), "Article")
}

func init() { proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2) }

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0xeb, 0x86, 0x76, 0x9a, 0x05, 0x17, 0x33, 0x26, 0x83, 0x44, 0x6c, 0x65, 0x97, 0x80,
	0xb4, 0x56, 0xdd, 0x2e, 0xa8, 0x45, 0x4c, 0x0d, 0x27, 0xb4, 0x5b, 0x24, 0x2e, 0x5c, 0x90, 0x1b,
	0x7b, 0xc5, 0x52, 0x12, 0x07, 0xd7, 0x66, 0x8c, 0x5f, 0xc1, 0xcf, 0xe2, 0xb8, 0x23, 0xa7, 0x08,
	0xda, 0x7f, 0x90, 0xd3, 0x8e, 0x28, 0x76, 0x5a, 0xb6, 0x43, 0x4f, 0xc9, 0x1b, 0xbf, 0xcf, 0xa3,
	0xef, 0x4b, 0x02, 0x9f, 0x30, 0x6d, 0x64, 0x96, 0x8b, 0x51, 0xa5, 0x95, 0x51, 0x2f, 0x4e, 0x97,
	0xd2, 0x7c, 0xb1, 0x8b, 0x51, 0xa6, 0x8a, 0xf1, 0x52, 0x2d, 0xd5, 0xd8, 0x3d, 0x5e, 0xd8, 0x2b,
	0x97, 0x5c, 0x70, 0x77, 0xbe, 0x1e, 0xdd, 0x3d, 0x82, 0x07, 0x73, 0x2f, 0x40, 0xef, 0x60, 0x5f,
	0x72, 0x0c, 0x28, 0x88, 0x83, 0x64, 0xd4, 0xd4, 0xe4, 0xf5, 0x52, 0xe9, 0x62, 0x1a, 0x55, 0x5a,
	0x16, 0x4c, 0xdf, 0x5c, 0x8a, 0x9b, 0x19, 0xb3, 0x46, 0x7d, 0x28, 0x33, 0x2d, 0x0a, 0x51, 0x9a,
	0x59, 0xa6, 0x8a, 0xf6, 0x3a, 0x95, 0x3c, 0x4a, 0xfb, 0x92, 0xa3, 0x0b, 0x18, 0x30, 0xc9, 0x71,
	0x9f, 0x82, 0xf8, 0x30, 0x39, 0x6d, 0x6a, 0xf2, 0xca, 0x0b, 0x6c, 0x29, 0xbf, 0x5a, 0x31, 0x5b,
	0xc9, 0x1f, 0x62, 0x3a, 0x39, 0x7b, 0xb3, 0xe3, 0xba, 0xc1, 0x69, 0xcb, 0xb7, 0x24, 0xfa, 0x08,
	0x83, 0x4a, 0x72, 0x1c, 0x38, 0xc1, 0xfb, 0xa6, 0x26, 0x17, 0x5e, 0x20, 0x4b, 0x2e, 0xbe, 0xff,
	0xe7, 0x4b, 0x65, 0x68, 0x69, 0xf3, 0x7c, 0xc6, 0xc5, 0x15, 0xb3, 0xb9, 0x99, 0xba, 0xb0, 0xb5,
	0x56, 0x5a, 0x71, 0x9b, 0x19, 0x5a, 0x39, 0x6d, 0x25, 0x39, 0x1a, 0xc1, 0x81, 0x91, 0x26, 0x17,
	0x78, 0xe0, 0xc4, 0xb8, 0xa9, 0xc9, 0x91, 0x17, 0x6f, 0x19, 0x77, 0x1c, 0xa5, 0xbe, 0xd6, 0xf6,
	0x33, 0xf5, 0x4d, 0x68, 0x3c, 0xdc, 0xd7, 0x77, 0xc7, 0x51, 0xea, 0x6b, 0x28, 0x86, 0x81, 0x66,
	0xd7, 0xf8, 0x80, 0x82, 0xf8, 0x71, 0x72, 0xdc, 0xd4, 0x04, 0x3d, 0x6c, 0x6b, 0x76, 0x1d, 0xa5,
	0x6d, 0x05, 0x5d, 0xc2, 0xe1, 0xca, 0x30, 0x63, 0x57, 0xf8, 0x29, 0x05, 0xf1, 0x20, 0x39, 0x6f,
	0x6a, 0x32, 0xf6, 0xe5, 0xed, 0x2e, 0x13, 0xbf, 0xe7, 0xd9, 0x6e, 0x1f, 0x0f, 0xd0, 0x09, 0x2d,
	0x95, 0x2e, 0x58, 0x1e, 0xa5, 0x9d, 0x02, 0x25, 0x10, 0x66, 0x5a, 0x30, 0x23, 0xf8, 0x67, 0x66,
	0xf0, 0x91, 0xfb, 0x6c, 0x27, 0x4d, 0x4d, 0xc8, 0xfd, 0x97, 0xb6, 0x9b, 0xd8, 0x37, 0x29, 0x33,
	0x51, 0x7a, 0xd8, 0x85, 0xb9, 0x41, 0x6f, 0x21, 0xb4, 0x15, 0xdf, 0x3a, 0x9e, 0x39, 0xc7, 0xcb,
	0xa6, 0x26, 0xcf, 0x1f, 0x6e, 0xd0, 0x75, 0x3c, 0xdd, 0x05, 0x4f, 0x73, 0x91, 0x8b, 0x8e, 0x3e,
	0xde, 0x47, 0x77, 0x1d, 0x4f, 0x77, 0x61, 0x6e, 0x92, 0x93, 0xbb, 0xbf, 0x21, 0xf8, 0xb5, 0x0e,
	0xc1, 0xed, 0x3a, 0x04, 0x7f, 0xd6, 0x21, 0xf8, 0xb9, 0x09, 0x7b, 0xb7, 0x9b, 0xb0, 0xf7, 0x7b,
	0x13, 0xf6, 0x3e, 0x0d, 0x0a, 0xc5, 0x45, 0xbe, 0x18, 0xba, 0xdf, 0xf4, 0xfc, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x85, 0x8d, 0x09, 0x42, 0xe6, 0x02, 0x00, 0x00,
}

func (m *Article) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Article) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Article) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DeletedAt != 0 {
		i = encodeVarintArticle(dAtA, i, uint64(m.DeletedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb0
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintArticle(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa8
	}
	if m.CreatedAt != 0 {
		i = encodeVarintArticle(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.Status != 0 {
		i = encodeVarintArticle(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if len(m.Raw) > 0 {
		i -= len(m.Raw)
		copy(dAtA[i:], m.Raw)
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Raw)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Cover) > 0 {
		i -= len(m.Cover)
		copy(dAtA[i:], m.Cover)
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Cover)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Pid) > 0 {
		i -= len(m.Pid)
		copy(dAtA[i:], m.Pid)
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Pid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Aid) > 0 {
		i -= len(m.Aid)
		copy(dAtA[i:], m.Aid)
		i = encodeVarintArticle(dAtA, i, uint64(len(m.Aid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintArticle(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintArticle(dAtA []byte, offset int, v uint64) int {
	offset -= sovArticle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedArticle(r randyArticle, easy bool) *Article {
	this := &Article{}
	this.Id = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Id *= -1
	}
	this.Aid = string(randStringArticle(r))
	this.Pid = string(randStringArticle(r))
	this.Title = string(randStringArticle(r))
	this.Cover = string(randStringArticle(r))
	v1 := r.Intn(100)
	this.Raw = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Raw[i] = byte(r.Intn(256))
	}
	this.Status = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Status *= -1
	}
	this.CreatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.CreatedAt *= -1
	}
	this.UpdatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.UpdatedAt *= -1
	}
	this.DeletedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.DeletedAt *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyArticle interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneArticle(r randyArticle) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringArticle(r randyArticle) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneArticle(r)
	}
	return string(tmps)
}
func randUnrecognizedArticle(r randyArticle, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldArticle(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldArticle(dAtA []byte, r randyArticle, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateArticle(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateArticle(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateArticle(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateArticle(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateArticle(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateArticle(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateArticle(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Article) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovArticle(uint64(m.Id))
	}
	l = len(m.Aid)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Pid)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Cover)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	l = len(m.Raw)
	if l > 0 {
		n += 1 + l + sovArticle(uint64(l))
	}
	if m.Status != 0 {
		n += 2 + sovArticle(uint64(m.Status))
	}
	if m.CreatedAt != 0 {
		n += 2 + sovArticle(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 2 + sovArticle(uint64(m.UpdatedAt))
	}
	if m.DeletedAt != 0 {
		n += 2 + sovArticle(uint64(m.DeletedAt))
	}
	return n
}

func sovArticle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozArticle(x uint64) (n int) {
	return sovArticle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Article) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArticle
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
			return fmt.Errorf("proto: Article: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Article: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
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
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
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
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
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
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cover", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
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
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cover = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
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
				return ErrInvalidLengthArticle
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthArticle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			m.DeletedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArticle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeletedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipArticle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthArticle
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
func skipArticle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowArticle
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
					return 0, ErrIntOverflowArticle
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
					return 0, ErrIntOverflowArticle
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
				return 0, ErrInvalidLengthArticle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupArticle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthArticle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthArticle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowArticle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupArticle = fmt.Errorf("proto: unexpected end of group")
)
