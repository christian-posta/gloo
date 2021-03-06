// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto

package v1 // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=sec
// @solo-kit:resource.plural_name=secrets
// @solo-kit:resource.resource_groups=api.gloo.solo.io,discovery.gloo.solo.io,translator.supergloo.solo.io
//
// Certain plugins such as the AWS Lambda Plugin require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.
//
// Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:
//
// Kubernetes Secrets
// Hashicorp Vault
// Plaintext files (recommended only for testing)
// Secrets must adhere to a structure, specified by the plugin that requires them.
//
// Gloo's secret backend can be configured in Gloo's bootstrap options
type Secret struct {
	// Types that are valid to be assigned to Kind:
	//	*Secret_Aws
	//	*Secret_Azure
	//	*Secret_Tls
	Kind isSecret_Kind `protobuf_oneof:"kind"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_87775a72bfd2ab43, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (dst *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(dst, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

type isSecret_Kind interface {
	isSecret_Kind()
	Equal(interface{}) bool
}

type Secret_Aws struct {
	Aws *AwsSecret `protobuf:"bytes,1,opt,name=aws,oneof"`
}
type Secret_Azure struct {
	Azure *AzureSecret `protobuf:"bytes,2,opt,name=azure,oneof"`
}
type Secret_Tls struct {
	Tls *TlsSecret `protobuf:"bytes,3,opt,name=tls,oneof"`
}

func (*Secret_Aws) isSecret_Kind()   {}
func (*Secret_Azure) isSecret_Kind() {}
func (*Secret_Tls) isSecret_Kind()   {}

func (m *Secret) GetKind() isSecret_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Secret) GetAws() *AwsSecret {
	if x, ok := m.GetKind().(*Secret_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Secret) GetAzure() *AzureSecret {
	if x, ok := m.GetKind().(*Secret_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Secret) GetTls() *TlsSecret {
	if x, ok := m.GetKind().(*Secret_Tls); ok {
		return x.Tls
	}
	return nil
}

func (m *Secret) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Secret) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Secret_OneofMarshaler, _Secret_OneofUnmarshaler, _Secret_OneofSizer, []interface{}{
		(*Secret_Aws)(nil),
		(*Secret_Azure)(nil),
		(*Secret_Tls)(nil),
	}
}

func _Secret_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Secret)
	// kind
	switch x := m.Kind.(type) {
	case *Secret_Aws:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *Secret_Azure:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case *Secret_Tls:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tls); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Secret.Kind has unexpected type %T", x)
	}
	return nil
}

func _Secret_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Secret)
	switch tag {
	case 1: // kind.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AwsSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Aws{msg}
		return true, err
	case 2: // kind.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AzureSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Azure{msg}
		return true, err
	case 3: // kind.tls
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TlsSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Tls{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Secret_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Secret)
	// kind
	switch x := m.Kind.(type) {
	case *Secret_Aws:
		s := proto.Size(x.Aws)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Azure:
		s := proto.Size(x.Azure)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Tls:
		s := proto.Size(x.Tls)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AwsSecret struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsSecret) Reset()         { *m = AwsSecret{} }
func (m *AwsSecret) String() string { return proto.CompactTextString(m) }
func (*AwsSecret) ProtoMessage()    {}
func (*AwsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_87775a72bfd2ab43, []int{1}
}
func (m *AwsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsSecret.Unmarshal(m, b)
}
func (m *AwsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsSecret.Marshal(b, m, deterministic)
}
func (dst *AwsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsSecret.Merge(dst, src)
}
func (m *AwsSecret) XXX_Size() int {
	return xxx_messageInfo_AwsSecret.Size(m)
}
func (m *AwsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AwsSecret proto.InternalMessageInfo

func (m *AwsSecret) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AwsSecret) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type AzureSecret struct {
	ApiKeys              map[string]string `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys" json:"api_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AzureSecret) Reset()         { *m = AzureSecret{} }
func (m *AzureSecret) String() string { return proto.CompactTextString(m) }
func (*AzureSecret) ProtoMessage()    {}
func (*AzureSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_87775a72bfd2ab43, []int{2}
}
func (m *AzureSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSecret.Unmarshal(m, b)
}
func (m *AzureSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSecret.Marshal(b, m, deterministic)
}
func (dst *AzureSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSecret.Merge(dst, src)
}
func (m *AzureSecret) XXX_Size() int {
	return xxx_messageInfo_AzureSecret.Size(m)
}
func (m *AzureSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSecret proto.InternalMessageInfo

func (m *AzureSecret) GetApiKeys() map[string]string {
	if m != nil {
		return m.ApiKeys
	}
	return nil
}

type TlsSecret struct {
	CertChain            string   `protobuf:"bytes,1,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	PrivateKey           string   `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	RootCa               string   `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsSecret) Reset()         { *m = TlsSecret{} }
func (m *TlsSecret) String() string { return proto.CompactTextString(m) }
func (*TlsSecret) ProtoMessage()    {}
func (*TlsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_87775a72bfd2ab43, []int{3}
}
func (m *TlsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSecret.Unmarshal(m, b)
}
func (m *TlsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSecret.Marshal(b, m, deterministic)
}
func (dst *TlsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSecret.Merge(dst, src)
}
func (m *TlsSecret) XXX_Size() int {
	return xxx_messageInfo_TlsSecret.Size(m)
}
func (m *TlsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSecret proto.InternalMessageInfo

func (m *TlsSecret) GetCertChain() string {
	if m != nil {
		return m.CertChain
	}
	return ""
}

func (m *TlsSecret) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TlsSecret) GetRootCa() string {
	if m != nil {
		return m.RootCa
	}
	return ""
}

func init() {
	proto.RegisterType((*Secret)(nil), "gloo.solo.io.Secret")
	proto.RegisterType((*AwsSecret)(nil), "gloo.solo.io.AwsSecret")
	proto.RegisterType((*AzureSecret)(nil), "gloo.solo.io.AzureSecret")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.AzureSecret.ApiKeysEntry")
	proto.RegisterType((*TlsSecret)(nil), "gloo.solo.io.TlsSecret")
}
func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Secret_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Aws)
	if !ok {
		that2, ok := that.(Secret_Aws)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *Secret_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Azure)
	if !ok {
		that2, ok := that.(Secret_Azure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *Secret_Tls) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Tls)
	if !ok {
		that2, ok := that.(Secret_Tls)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Tls.Equal(that1.Tls) {
		return false
	}
	return true
}
func (this *AwsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsSecret)
	if !ok {
		that2, ok := that.(AwsSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AccessKey != that1.AccessKey {
		return false
	}
	if this.SecretKey != that1.SecretKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AzureSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AzureSecret)
	if !ok {
		that2, ok := that.(AzureSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ApiKeys) != len(that1.ApiKeys) {
		return false
	}
	for i := range this.ApiKeys {
		if this.ApiKeys[i] != that1.ApiKeys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TlsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TlsSecret)
	if !ok {
		that2, ok := that.(TlsSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CertChain != that1.CertChain {
		return false
	}
	if this.PrivateKey != that1.PrivateKey {
		return false
	}
	if this.RootCa != that1.RootCa {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto", fileDescriptor_secret_87775a72bfd2ab43)
}

var fileDescriptor_secret_87775a72bfd2ab43 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x9d, 0x34, 0x33, 0xe9, 0xe4, 0x67, 0x16, 0xc8, 0x1a, 0x31, 0xa1, 0x12, 0x0c, 0xca, 0x02,
	0x8d, 0x84, 0x70, 0xe8, 0x20, 0xa1, 0xa1, 0xbb, 0xb6, 0x42, 0x02, 0x55, 0x6c, 0x02, 0x2b, 0x36,
	0x95, 0xeb, 0x5a, 0xa9, 0x49, 0x5a, 0x47, 0xb6, 0xdb, 0x2a, 0x9c, 0x81, 0x83, 0x70, 0x14, 0x6e,
	0xc0, 0x8e, 0x05, 0x27, 0x41, 0xb6, 0xd3, 0x10, 0xa1, 0x22, 0xcd, 0x2a, 0xfe, 0xff, 0xbd, 0xff,
	0xfc, 0x9e, 0xf3, 0xe1, 0x4d, 0xce, 0xf5, 0x6a, 0xbb, 0xc0, 0x54, 0xac, 0x53, 0x25, 0x4a, 0xf1,
	0x82, 0x8b, 0x34, 0x2f, 0x85, 0x48, 0x2b, 0x29, 0xbe, 0x30, 0xaa, 0x95, 0xab, 0x48, 0xc5, 0xd3,
	0xdd, 0x30, 0x55, 0x8c, 0x4a, 0xa6, 0x71, 0x25, 0x85, 0x16, 0xe8, 0xc2, 0x20, 0xd8, 0x0c, 0x61,
	0x2e, 0x06, 0x97, 0xb9, 0xc8, 0x85, 0x05, 0x52, 0x73, 0x72, 0x9c, 0xc1, 0xf0, 0x88, 0xbc, 0xfd,
	0x16, 0x5c, 0x1f, 0x44, 0xd7, 0x4c, 0x93, 0x25, 0xd1, 0xc4, 0x8d, 0x24, 0x3f, 0x3d, 0x08, 0x3e,
	0xda, 0x7b, 0xd0, 0x73, 0xf0, 0xc9, 0x5e, 0xc5, 0xde, 0x53, 0xef, 0x26, 0xba, 0xbd, 0xc2, 0xdd,
	0xfb, 0xf0, 0x78, 0xaf, 0x1c, 0xeb, 0xdd, 0x49, 0x66, 0x58, 0x68, 0x08, 0x67, 0xe4, 0xeb, 0x56,
	0xb2, 0xb8, 0x67, 0xe9, 0x8f, 0xfe, 0xa1, 0x1b, 0xa8, 0x1d, 0x70, 0x4c, 0xa3, 0xaf, 0x4b, 0x15,
	0xfb, 0xc7, 0xf4, 0x3f, 0x95, 0x1d, 0x7d, 0x5d, 0x2a, 0x74, 0x07, 0xe7, 0x07, 0xa7, 0x71, 0xdf,
	0x4e, 0x3c, 0xc4, 0x54, 0x48, 0xd6, 0x4e, 0x7c, 0x68, 0xd0, 0xc9, 0xe9, 0x8f, 0x5f, 0xd7, 0x27,
	0x59, 0xcb, 0x9e, 0x04, 0x70, 0x5a, 0xf0, 0xcd, 0x32, 0x79, 0x0f, 0x61, 0xeb, 0x1a, 0x3d, 0x06,
	0x20, 0x94, 0x32, 0xa5, 0xe6, 0x05, 0xab, 0x6d, 0xc4, 0x30, 0x0b, 0x5d, 0x67, 0xc6, 0x6a, 0x03,
	0xbb, 0xc7, 0xb6, 0x70, 0xcf, 0xc1, 0xae, 0x33, 0x63, 0x75, 0xf2, 0xcd, 0x83, 0xa8, 0x13, 0x09,
	0x8d, 0xe1, 0x9c, 0x54, 0xdc, 0x70, 0xcd, 0x73, 0xf9, 0x37, 0xd1, 0xed, 0xb3, 0xff, 0xe6, 0xc7,
	0xe3, 0x8a, 0xcf, 0x58, 0xad, 0xde, 0x6e, 0xb4, 0xac, 0xb3, 0x3e, 0x71, 0xd5, 0x60, 0x04, 0x17,
	0x5d, 0x00, 0x3d, 0x00, 0xff, 0xaf, 0x33, 0x73, 0x44, 0x97, 0x70, 0xb6, 0x23, 0xe5, 0x96, 0x35,
	0x76, 0x5c, 0x31, 0xea, 0xdd, 0x79, 0xc9, 0x12, 0xc2, 0xf6, 0xbd, 0x8c, 0x75, 0xca, 0xa4, 0x9e,
	0xd3, 0x15, 0xe1, 0x9b, 0x43, 0x32, 0xd3, 0x99, 0x9a, 0x06, 0xba, 0x86, 0xa8, 0x92, 0x7c, 0x47,
	0x34, 0xeb, 0x44, 0x83, 0xa6, 0x65, 0xa2, 0x5f, 0x41, 0x5f, 0x0a, 0xa1, 0xe7, 0x94, 0xd8, 0x3f,
	0x13, 0x66, 0x81, 0x29, 0xa7, 0x64, 0xf2, 0xfa, 0xfb, 0xef, 0x27, 0xde, 0xe7, 0x97, 0xf7, 0xdb,
	0xd8, 0xaa, 0xc8, 0x9b, 0x05, 0x5b, 0x04, 0x76, 0xb1, 0x5e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x22, 0x9f, 0x20, 0xf9, 0xec, 0x02, 0x00, 0x00,
}
