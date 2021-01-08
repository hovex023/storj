// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ordersmeta.proto

package internalpb

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// OrderLimitMetadata is used to transmit meta information about an order limit.
// This data will be encrypted.
type OrderLimitMetadata struct {
	BucketId                   []byte   `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	ProjectBucketPrefix        []byte   `protobuf:"bytes,2,opt,name=project_bucket_prefix,json=projectBucketPrefix,proto3" json:"project_bucket_prefix,omitempty"`
	CompactProjectBucketPrefix []byte   `protobuf:"bytes,3,opt,name=compact_project_bucket_prefix,json=compactProjectBucketPrefix,proto3" json:"compact_project_bucket_prefix,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *OrderLimitMetadata) Reset()         { *m = OrderLimitMetadata{} }
func (m *OrderLimitMetadata) String() string { return proto.CompactTextString(m) }
func (*OrderLimitMetadata) ProtoMessage()    {}
func (*OrderLimitMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00ef1afe54bb544, []int{0}
}
func (m *OrderLimitMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLimitMetadata.Unmarshal(m, b)
}
func (m *OrderLimitMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLimitMetadata.Marshal(b, m, deterministic)
}
func (m *OrderLimitMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLimitMetadata.Merge(m, src)
}
func (m *OrderLimitMetadata) XXX_Size() int {
	return xxx_messageInfo_OrderLimitMetadata.Size(m)
}
func (m *OrderLimitMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLimitMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLimitMetadata proto.InternalMessageInfo

func (m *OrderLimitMetadata) GetBucketId() []byte {
	if m != nil {
		return m.BucketId
	}
	return nil
}

func (m *OrderLimitMetadata) GetProjectBucketPrefix() []byte {
	if m != nil {
		return m.ProjectBucketPrefix
	}
	return nil
}

func (m *OrderLimitMetadata) GetCompactProjectBucketPrefix() []byte {
	if m != nil {
		return m.CompactProjectBucketPrefix
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderLimitMetadata)(nil), "satellite.ordersmeta.OrderLimitMetadata")
}

func init() { proto.RegisterFile("ordersmeta.proto", fileDescriptor_e00ef1afe54bb544) }

var fileDescriptor_e00ef1afe54bb544 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x2f, 0x4a, 0x49,
	0x2d, 0x2a, 0xce, 0x4d, 0x2d, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x29, 0x4e,
	0x2c, 0x49, 0xcd, 0xc9, 0xc9, 0x2c, 0x49, 0xd5, 0x43, 0xc8, 0x29, 0xad, 0x60, 0xe4, 0x12, 0xf2,
	0x07, 0x71, 0x7d, 0x32, 0x73, 0x33, 0x4b, 0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x85,
	0xa4, 0xb9, 0x38, 0x93, 0x4a, 0x93, 0xb3, 0x53, 0x4b, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x78, 0x82, 0x38, 0x20, 0x02, 0x9e, 0x29, 0x42, 0x46, 0x5c, 0xa2, 0x05, 0x45, 0xf9, 0x59,
	0xa9, 0xc9, 0x25, 0xf1, 0x50, 0x45, 0x05, 0x45, 0xa9, 0x69, 0x99, 0x15, 0x12, 0x4c, 0x60, 0x85,
	0xc2, 0x50, 0x49, 0x27, 0xb0, 0x5c, 0x00, 0x58, 0x4a, 0xc8, 0x91, 0x4b, 0x36, 0x39, 0x3f, 0xb7,
	0x20, 0x31, 0x19, 0xa4, 0x18, 0x9b, 0x5e, 0x66, 0xb0, 0x5e, 0x29, 0xa8, 0xa2, 0x00, 0x4c, 0x23,
	0x9c, 0x54, 0xa3, 0x94, 0x8b, 0x4b, 0xf2, 0x8b, 0xb2, 0xf4, 0x32, 0xf3, 0xf5, 0xc1, 0x0c, 0x7d,
	0xb8, 0x8f, 0xf4, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0xf2, 0x12, 0x73, 0x0a, 0x92, 0x92, 0xd8, 0xc0,
	0xde, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x00, 0xcc, 0x15, 0xb0, 0x02, 0x01, 0x00, 0x00,
}
