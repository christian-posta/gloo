// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest.proto

package rest // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/rest"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"

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

type ServiceSpec struct {
	Transformations      map[string]*transformation.TransformationTemplate `protobuf:"bytes,1,rep,name=transformations" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	SwaggerInfo          *ServiceSpec_SwaggerInfo                          `protobuf:"bytes,2,opt,name=swagger_info,json=swaggerInfo" json:"swagger_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_rest_36ddb4411c65025c, []int{0}
}
func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (dst *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(dst, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

func (m *ServiceSpec) GetTransformations() map[string]*transformation.TransformationTemplate {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *ServiceSpec) GetSwaggerInfo() *ServiceSpec_SwaggerInfo {
	if m != nil {
		return m.SwaggerInfo
	}
	return nil
}

type ServiceSpec_SwaggerInfo struct {
	// Types that are valid to be assigned to SwaggerSpec:
	//	*ServiceSpec_SwaggerInfo_Url
	//	*ServiceSpec_SwaggerInfo_Inline
	SwaggerSpec          isServiceSpec_SwaggerInfo_SwaggerSpec `protobuf_oneof:"swagger_spec"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ServiceSpec_SwaggerInfo) Reset()         { *m = ServiceSpec_SwaggerInfo{} }
func (m *ServiceSpec_SwaggerInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec_SwaggerInfo) ProtoMessage()    {}
func (*ServiceSpec_SwaggerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_rest_36ddb4411c65025c, []int{0, 1}
}
func (m *ServiceSpec_SwaggerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Unmarshal(m, b)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Marshal(b, m, deterministic)
}
func (dst *ServiceSpec_SwaggerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec_SwaggerInfo.Merge(dst, src)
}
func (m *ServiceSpec_SwaggerInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec_SwaggerInfo.Size(m)
}
func (m *ServiceSpec_SwaggerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec_SwaggerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec_SwaggerInfo proto.InternalMessageInfo

type isServiceSpec_SwaggerInfo_SwaggerSpec interface {
	isServiceSpec_SwaggerInfo_SwaggerSpec()
	Equal(interface{}) bool
}

type ServiceSpec_SwaggerInfo_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}
type ServiceSpec_SwaggerInfo_Inline struct {
	Inline string `protobuf:"bytes,2,opt,name=inline,proto3,oneof"`
}

func (*ServiceSpec_SwaggerInfo_Url) isServiceSpec_SwaggerInfo_SwaggerSpec()    {}
func (*ServiceSpec_SwaggerInfo_Inline) isServiceSpec_SwaggerInfo_SwaggerSpec() {}

func (m *ServiceSpec_SwaggerInfo) GetSwaggerSpec() isServiceSpec_SwaggerInfo_SwaggerSpec {
	if m != nil {
		return m.SwaggerSpec
	}
	return nil
}

func (m *ServiceSpec_SwaggerInfo) GetUrl() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Url); ok {
		return x.Url
	}
	return ""
}

func (m *ServiceSpec_SwaggerInfo) GetInline() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Inline); ok {
		return x.Inline
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServiceSpec_SwaggerInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServiceSpec_SwaggerInfo_OneofMarshaler, _ServiceSpec_SwaggerInfo_OneofUnmarshaler, _ServiceSpec_SwaggerInfo_OneofSizer, []interface{}{
		(*ServiceSpec_SwaggerInfo_Url)(nil),
		(*ServiceSpec_SwaggerInfo_Inline)(nil),
	}
}

func _ServiceSpec_SwaggerInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServiceSpec_SwaggerInfo)
	// swagger_spec
	switch x := m.SwaggerSpec.(type) {
	case *ServiceSpec_SwaggerInfo_Url:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Url)
	case *ServiceSpec_SwaggerInfo_Inline:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Inline)
	case nil:
	default:
		return fmt.Errorf("ServiceSpec_SwaggerInfo.SwaggerSpec has unexpected type %T", x)
	}
	return nil
}

func _ServiceSpec_SwaggerInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServiceSpec_SwaggerInfo)
	switch tag {
	case 1: // swagger_spec.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.SwaggerSpec = &ServiceSpec_SwaggerInfo_Url{x}
		return true, err
	case 2: // swagger_spec.inline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.SwaggerSpec = &ServiceSpec_SwaggerInfo_Inline{x}
		return true, err
	default:
		return false, nil
	}
}

func _ServiceSpec_SwaggerInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServiceSpec_SwaggerInfo)
	// swagger_spec
	switch x := m.SwaggerSpec.(type) {
	case *ServiceSpec_SwaggerInfo_Url:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case *ServiceSpec_SwaggerInfo_Inline:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Inline)))
		n += len(x.Inline)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// This is only for upstream with REST service spec
type DestinationSpec struct {
	FunctionName           string                                 `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	Parameters             *transformation.Parameters             `protobuf:"bytes,2,opt,name=parameters" json:"parameters,omitempty"`
	ResponseTransformation *transformation.TransformationTemplate `protobuf:"bytes,3,opt,name=response_transformation,json=responseTransformation" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                               `json:"-"`
	XXX_unrecognized       []byte                                 `json:"-"`
	XXX_sizecache          int32                                  `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_rest_36ddb4411c65025c, []int{1}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (dst *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(dst, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

func (m *DestinationSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *DestinationSpec) GetParameters() *transformation.Parameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *DestinationSpec) GetResponseTransformation() *transformation.TransformationTemplate {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "rest.plugins.gloo.solo.io.ServiceSpec")
	proto.RegisterMapType((map[string]*transformation.TransformationTemplate)(nil), "rest.plugins.gloo.solo.io.ServiceSpec.TransformationsEntry")
	proto.RegisterType((*ServiceSpec_SwaggerInfo)(nil), "rest.plugins.gloo.solo.io.ServiceSpec.SwaggerInfo")
	proto.RegisterType((*DestinationSpec)(nil), "rest.plugins.gloo.solo.io.DestinationSpec")
}
func (this *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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
	if len(this.Transformations) != len(that1.Transformations) {
		return false
	}
	for i := range this.Transformations {
		if !this.Transformations[i].Equal(that1.Transformations[i]) {
			return false
		}
	}
	if !this.SwaggerInfo.Equal(that1.SwaggerInfo) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo)
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
	if that1.SwaggerSpec == nil {
		if this.SwaggerSpec != nil {
			return false
		}
	} else if this.SwaggerSpec == nil {
		return false
	} else if !this.SwaggerSpec.Equal(that1.SwaggerSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Url) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Url)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Url)
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
	if this.Url != that1.Url {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Inline) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Inline)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Inline)
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
	if this.Inline != that1.Inline {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest.proto", fileDescriptor_rest_36ddb4411c65025c)
}

var fileDescriptor_rest_36ddb4411c65025c = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0xed, 0x16, 0x17, 0x4c, 0x57, 0x57, 0xc2, 0xa2, 0xb5, 0x17, 0x32, 0xac, 0x37, 0x73,
	0x63, 0x82, 0xe3, 0x8d, 0xb8, 0x78, 0xb3, 0xae, 0xa0, 0x08, 0xfe, 0xe9, 0x8c, 0x20, 0xde, 0x0c,
	0xd9, 0x72, 0x1a, 0xe3, 0xb6, 0x39, 0x21, 0x49, 0x47, 0xf6, 0xca, 0xe7, 0xf0, 0x0d, 0x7c, 0x2e,
	0xdf, 0x43, 0x90, 0x36, 0x1d, 0x67, 0x3a, 0x8c, 0x30, 0xe8, 0xde, 0x0c, 0xe7, 0x4b, 0xe6, 0xfb,
	0x7d, 0x39, 0x27, 0x29, 0x39, 0x93, 0xca, 0x7f, 0x6e, 0xce, 0x59, 0x81, 0x35, 0x77, 0x58, 0xe1,
	0x43, 0x85, 0x5c, 0x56, 0x88, 0xdc, 0x58, 0xfc, 0x02, 0x85, 0x77, 0x41, 0x09, 0xa3, 0xf8, 0xe2,
	0x11, 0x37, 0x55, 0x23, 0x95, 0x76, 0xdc, 0x82, 0xf3, 0xdd, 0x0f, 0x33, 0x16, 0x3d, 0xd2, 0x7b,
	0xa1, 0x0e, 0xbb, 0xac, 0x75, 0xb0, 0x16, 0xc6, 0x14, 0x66, 0x47, 0x12, 0x25, 0x76, 0xff, 0xe2,
	0x6d, 0x15, 0x0c, 0xd9, 0xc7, 0x7f, 0x8a, 0xf5, 0x56, 0x68, 0x57, 0xa2, 0xad, 0x85, 0x57, 0xa8,
	0x37, 0x64, 0x4f, 0x9e, 0x5d, 0x05, 0xd9, 0x08, 0x2b, 0x6a, 0xf0, 0x60, 0x5d, 0xa0, 0x1e, 0x7f,
	0x8f, 0x49, 0x32, 0x05, 0xbb, 0x50, 0x05, 0x4c, 0x0d, 0x14, 0x14, 0xc8, 0xe1, 0xd0, 0xe2, 0xd2,
	0x68, 0x14, 0x8f, 0x93, 0xc9, 0x09, 0xfb, 0xeb, 0x28, 0xd8, 0x1a, 0x80, 0xcd, 0x86, 0xee, 0x17,
	0xda, 0xdb, 0xcb, 0x7c, 0x93, 0x49, 0x3f, 0x90, 0x03, 0xf7, 0x55, 0x48, 0x09, 0x76, 0xae, 0x74,
	0x89, 0xe9, 0xde, 0x28, 0x1a, 0x27, 0x93, 0xc9, 0x8e, 0x19, 0xd3, 0x60, 0x7d, 0xa5, 0x4b, 0xcc,
	0x13, 0xb7, 0x12, 0xd9, 0x37, 0x72, 0xb4, 0x2d, 0x9f, 0xde, 0x26, 0xf1, 0x05, 0x5c, 0xa6, 0xd1,
	0x28, 0x1a, 0xdf, 0xc8, 0xdb, 0x92, 0xbe, 0x27, 0xd7, 0x17, 0xa2, 0x6a, 0xa0, 0x4f, 0x3e, 0x61,
	0x9b, 0x33, 0xdf, 0x76, 0x86, 0x21, 0x7b, 0x06, 0xb5, 0xa9, 0x84, 0x87, 0x3c, 0x90, 0x9e, 0xee,
	0x3d, 0x89, 0xb2, 0xd7, 0x24, 0x59, 0x3b, 0x1c, 0xa5, 0x24, 0x6e, 0x6c, 0x15, 0x72, 0x5f, 0x5e,
	0xcb, 0x5b, 0x41, 0x53, 0xb2, 0xaf, 0x74, 0xa5, 0x74, 0x88, 0x6e, 0x97, 0x7b, 0x7d, 0x7a, 0x6b,
	0x35, 0x14, 0x67, 0xa0, 0x38, 0xfe, 0x15, 0x91, 0xc3, 0x33, 0x70, 0x5e, 0xe9, 0x2e, 0xaf, 0xbb,
	0x9f, 0x07, 0xe4, 0x66, 0xd9, 0xe8, 0xa2, 0xd5, 0x73, 0x2d, 0x6a, 0xe8, 0x7b, 0x3a, 0x58, 0x2e,
	0xbe, 0x11, 0x35, 0xd0, 0xb7, 0x84, 0xac, 0x2e, 0xba, 0xef, 0x90, 0xef, 0xd4, 0xe1, 0xbb, 0x3f,
	0xb6, 0x7c, 0x0d, 0x41, 0x3d, 0xb9, 0x6b, 0xc1, 0x19, 0xd4, 0x0e, 0xe6, 0x43, 0x4c, 0x1a, 0xff,
	0xff, 0xfc, 0xee, 0x2c, 0xd9, 0xc3, 0xfd, 0xd3, 0xe7, 0x3f, 0x7e, 0xde, 0x8f, 0x3e, 0x3d, 0xdb,
	0xed, 0xdd, 0x9b, 0x0b, 0xb9, 0xed, 0x63, 0x3e, 0xdf, 0xef, 0xde, 0xf9, 0xe3, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xee, 0x95, 0x1d, 0x3a, 0x10, 0x04, 0x00, 0x00,
}
