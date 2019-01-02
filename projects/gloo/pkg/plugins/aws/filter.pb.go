// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter.proto

package aws // import "github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// AWS Lambda contains the configuration necessary to perform transform regular http calls to
// AWS Lambda invocations.
type LambdaPerRoute struct {
	// The name of the function
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The qualifier of the function (defualts to $LATEST if not specified)
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	// Invocation type - async or regular.
	Async                bool     `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LambdaPerRoute) Reset()         { *m = LambdaPerRoute{} }
func (m *LambdaPerRoute) String() string { return proto.CompactTextString(m) }
func (*LambdaPerRoute) ProtoMessage()    {}
func (*LambdaPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_3d010225fa27eb6d, []int{0}
}
func (m *LambdaPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LambdaPerRoute.Unmarshal(m, b)
}
func (m *LambdaPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LambdaPerRoute.Marshal(b, m, deterministic)
}
func (dst *LambdaPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LambdaPerRoute.Merge(dst, src)
}
func (m *LambdaPerRoute) XXX_Size() int {
	return xxx_messageInfo_LambdaPerRoute.Size(m)
}
func (m *LambdaPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_LambdaPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_LambdaPerRoute proto.InternalMessageInfo

func (m *LambdaPerRoute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LambdaPerRoute) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

func (m *LambdaPerRoute) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

type LambdaProtocolExtension struct {
	// The host header for AWS this cluster
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The region for this cluster
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The access_key for AWS this cluster
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// The secret_key for AWS this cluster
	SecretKey            string   `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LambdaProtocolExtension) Reset()         { *m = LambdaProtocolExtension{} }
func (m *LambdaProtocolExtension) String() string { return proto.CompactTextString(m) }
func (*LambdaProtocolExtension) ProtoMessage()    {}
func (*LambdaProtocolExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_3d010225fa27eb6d, []int{1}
}
func (m *LambdaProtocolExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LambdaProtocolExtension.Unmarshal(m, b)
}
func (m *LambdaProtocolExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LambdaProtocolExtension.Marshal(b, m, deterministic)
}
func (dst *LambdaProtocolExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LambdaProtocolExtension.Merge(dst, src)
}
func (m *LambdaProtocolExtension) XXX_Size() int {
	return xxx_messageInfo_LambdaProtocolExtension.Size(m)
}
func (m *LambdaProtocolExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_LambdaProtocolExtension.DiscardUnknown(m)
}

var xxx_messageInfo_LambdaProtocolExtension proto.InternalMessageInfo

func (m *LambdaProtocolExtension) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *LambdaProtocolExtension) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *LambdaProtocolExtension) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *LambdaProtocolExtension) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func init() {
	proto.RegisterType((*LambdaPerRoute)(nil), "envoy.config.filter.http.aws.v2.LambdaPerRoute")
	proto.RegisterType((*LambdaProtocolExtension)(nil), "envoy.config.filter.http.aws.v2.LambdaProtocolExtension")
}

func init() { proto.RegisterFile("filter.proto", fileDescriptor_filter_3d010225fa27eb6d) }

var fileDescriptor_filter_3d010225fa27eb6d = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0xb6, 0x16, 0x27, 0x88, 0x8b, 0x20, 0xb4, 0x88, 0x62, 0xed, 0xaa, 0x1b, 0x13,
	0x50, 0xc4, 0x7d, 0xc1, 0x55, 0x5d, 0xc8, 0x2c, 0xdd, 0x48, 0x9a, 0xde, 0xc9, 0xc4, 0x66, 0x72,
	0xc7, 0x24, 0x33, 0x75, 0x9e, 0xca, 0xbd, 0x2b, 0x5f, 0xc7, 0xb7, 0x90, 0xf9, 0x11, 0x15, 0xdd,
	0xe5, 0x9c, 0xf3, 0x85, 0x0f, 0x2e, 0x39, 0x48, 0xb5, 0x09, 0xe0, 0x58, 0xe1, 0x30, 0x20, 0x3d,
	0x03, 0x5b, 0x61, 0xcd, 0x24, 0xda, 0x54, 0x2b, 0xd6, 0x4f, 0x59, 0x08, 0x05, 0x13, 0x3b, 0xcf,
	0xaa, 0xcb, 0xe3, 0x49, 0x25, 0x8c, 0xde, 0x88, 0x00, 0xfc, 0xeb, 0xd1, 0xfd, 0x9c, 0x4b, 0x72,
	0x78, 0x27, 0xf2, 0xf5, 0x46, 0xdc, 0x83, 0x4b, 0xb0, 0x0c, 0x40, 0x4f, 0xc9, 0xc8, 0x8a, 0x1c,
	0xa6, 0xd1, 0x2c, 0x5a, 0xc4, 0xcb, 0xf8, 0xed, 0xe3, 0x7d, 0x38, 0x72, 0x83, 0x59, 0x94, 0xb4,
	0x35, 0x3d, 0x21, 0xf1, 0x73, 0x29, 0x8c, 0x4e, 0x35, 0xb8, 0xe9, 0xa0, 0x61, 0x92, 0xef, 0x82,
	0x1e, 0x91, 0x3d, 0xe1, 0x6b, 0x2b, 0xa7, 0xc3, 0x59, 0xb4, 0xd8, 0x4f, 0xba, 0x30, 0x7f, 0x8d,
	0xc8, 0xa4, 0xb7, 0x34, 0x52, 0x89, 0xe6, 0xf6, 0x25, 0x80, 0xf5, 0x1a, 0x6d, 0xa3, 0xcb, 0xd0,
	0x87, 0x7f, 0x74, 0x4d, 0x4d, 0xcf, 0xc9, 0xd8, 0x81, 0xd2, 0x68, 0x3b, 0xd7, 0x4f, 0xa0, 0x1f,
	0xe8, 0x82, 0x10, 0x21, 0x25, 0x78, 0xff, 0xb8, 0x85, 0xba, 0x15, 0xff, 0xc2, 0xe2, 0x6e, 0x5c,
	0x41, 0xdd, 0x90, 0x1e, 0xa4, 0x83, 0xd0, 0x92, 0xa3, 0x3f, 0x64, 0x37, 0xae, 0xa0, 0x5e, 0xde,
	0x3c, 0x5c, 0x2b, 0x1d, 0xb2, 0x72, 0xcd, 0x24, 0xe6, 0xdc, 0xa3, 0xc1, 0x0b, 0x8d, 0x5c, 0x19,
	0x44, 0x5e, 0x38, 0x7c, 0x02, 0x19, 0x7c, 0x9f, 0xb6, 0x8a, 0x17, 0xa6, 0x54, 0xda, 0x7a, 0x2e,
	0x76, 0x7e, 0x3d, 0x6e, 0xcf, 0x7a, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x41, 0x6f, 0xc7, 0x8b,
	0xa0, 0x01, 0x00, 0x00,
}
