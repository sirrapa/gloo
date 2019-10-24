// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation/transformation.proto

package transformation

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FilterTransformations struct {
	// Specifies transformations based on the route matches. The first matched transformation will be
	// applied. If there are overlapped match conditions, please put the most specific match first.
	Transformations      []*TransformationRule `protobuf:"bytes,1,rep,name=transformations,proto3" json:"transformations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FilterTransformations) Reset()         { *m = FilterTransformations{} }
func (m *FilterTransformations) String() string { return proto.CompactTextString(m) }
func (*FilterTransformations) ProtoMessage()    {}
func (*FilterTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{0}
}
func (m *FilterTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterTransformations.Unmarshal(m, b)
}
func (m *FilterTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterTransformations.Marshal(b, m, deterministic)
}
func (m *FilterTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterTransformations.Merge(m, src)
}
func (m *FilterTransformations) XXX_Size() int {
	return xxx_messageInfo_FilterTransformations.Size(m)
}
func (m *FilterTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_FilterTransformations proto.InternalMessageInfo

func (m *FilterTransformations) GetTransformations() []*TransformationRule {
	if m != nil {
		return m.Transformations
	}
	return nil
}

type TransformationRule struct {
	// The route matching parameter. Only when the match is satisfied, the "requires" field will
	// apply.
	//
	// For example: following match will match all requests.
	//
	// .. code-block:: yaml
	//
	//    match:
	//      prefix: /
	//
	Match *route.RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// transformation to perform
	RouteTransformations *RouteTransformations `protobuf:"bytes,2,opt,name=route_transformations,json=routeTransformations,proto3" json:"route_transformations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TransformationRule) Reset()         { *m = TransformationRule{} }
func (m *TransformationRule) String() string { return proto.CompactTextString(m) }
func (*TransformationRule) ProtoMessage()    {}
func (*TransformationRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{1}
}
func (m *TransformationRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformationRule.Unmarshal(m, b)
}
func (m *TransformationRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformationRule.Marshal(b, m, deterministic)
}
func (m *TransformationRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformationRule.Merge(m, src)
}
func (m *TransformationRule) XXX_Size() int {
	return xxx_messageInfo_TransformationRule.Size(m)
}
func (m *TransformationRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformationRule.DiscardUnknown(m)
}

var xxx_messageInfo_TransformationRule proto.InternalMessageInfo

func (m *TransformationRule) GetMatch() *route.RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *TransformationRule) GetRouteTransformations() *RouteTransformations {
	if m != nil {
		return m.RouteTransformations
	}
	return nil
}

type RouteTransformations struct {
	RequestTransformation *Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// clear the route cache if the request transformation was applied
	ClearRouteCache        bool            `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	ResponseTransformation *Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *RouteTransformations) Reset()         { *m = RouteTransformations{} }
func (m *RouteTransformations) String() string { return proto.CompactTextString(m) }
func (*RouteTransformations) ProtoMessage()    {}
func (*RouteTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{2}
}
func (m *RouteTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTransformations.Unmarshal(m, b)
}
func (m *RouteTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTransformations.Marshal(b, m, deterministic)
}
func (m *RouteTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTransformations.Merge(m, src)
}
func (m *RouteTransformations) XXX_Size() int {
	return xxx_messageInfo_RouteTransformations.Size(m)
}
func (m *RouteTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTransformations proto.InternalMessageInfo

func (m *RouteTransformations) GetRequestTransformation() *Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *RouteTransformations) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *RouteTransformations) GetResponseTransformation() *Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

// [#proto-status: experimental]
type Transformation struct {
	// Template is in the transformed request language domain
	// currently both are JSON
	//
	// Types that are valid to be assigned to TransformationType:
	//	*Transformation_TransformationTemplate
	//	*Transformation_HeaderBodyTransform
	TransformationType   isTransformation_TransformationType `protobuf_oneof:"transformation_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{3}
}
func (m *Transformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transformation.Unmarshal(m, b)
}
func (m *Transformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transformation.Marshal(b, m, deterministic)
}
func (m *Transformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transformation.Merge(m, src)
}
func (m *Transformation) XXX_Size() int {
	return xxx_messageInfo_Transformation.Size(m)
}
func (m *Transformation) XXX_DiscardUnknown() {
	xxx_messageInfo_Transformation.DiscardUnknown(m)
}

var xxx_messageInfo_Transformation proto.InternalMessageInfo

type isTransformation_TransformationType interface {
	isTransformation_TransformationType()
	Equal(interface{}) bool
}

type Transformation_TransformationTemplate struct {
	TransformationTemplate *TransformationTemplate `protobuf:"bytes,1,opt,name=transformation_template,json=transformationTemplate,proto3,oneof" json:"transformation_template,omitempty"`
}
type Transformation_HeaderBodyTransform struct {
	HeaderBodyTransform *HeaderBodyTransform `protobuf:"bytes,2,opt,name=header_body_transform,json=headerBodyTransform,proto3,oneof" json:"header_body_transform,omitempty"`
}

func (*Transformation_TransformationTemplate) isTransformation_TransformationType() {}
func (*Transformation_HeaderBodyTransform) isTransformation_TransformationType()    {}

func (m *Transformation) GetTransformationType() isTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (m *Transformation) GetTransformationTemplate() *TransformationTemplate {
	if x, ok := m.GetTransformationType().(*Transformation_TransformationTemplate); ok {
		return x.TransformationTemplate
	}
	return nil
}

func (m *Transformation) GetHeaderBodyTransform() *HeaderBodyTransform {
	if x, ok := m.GetTransformationType().(*Transformation_HeaderBodyTransform); ok {
		return x.HeaderBodyTransform
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Transformation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Transformation_TransformationTemplate)(nil),
		(*Transformation_HeaderBodyTransform)(nil),
	}
}

type Extraction struct {
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// what information to extract. if extraction fails the result is
	// an empty value.
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	Subgroup             uint32   `protobuf:"varint,3,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extraction) Reset()         { *m = Extraction{} }
func (m *Extraction) String() string { return proto.CompactTextString(m) }
func (*Extraction) ProtoMessage()    {}
func (*Extraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{4}
}
func (m *Extraction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extraction.Unmarshal(m, b)
}
func (m *Extraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extraction.Marshal(b, m, deterministic)
}
func (m *Extraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extraction.Merge(m, src)
}
func (m *Extraction) XXX_Size() int {
	return xxx_messageInfo_Extraction.Size(m)
}
func (m *Extraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Extraction.DiscardUnknown(m)
}

var xxx_messageInfo_Extraction proto.InternalMessageInfo

func (m *Extraction) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Extraction) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

func (m *Extraction) GetSubgroup() uint32 {
	if m != nil {
		return m.Subgroup
	}
	return 0
}

type TransformationTemplate struct {
	AdvancedTemplates bool `protobuf:"varint,1,opt,name=advanced_templates,json=advancedTemplates,proto3" json:"advanced_templates,omitempty"`
	// Extractors are in the origin request language domain
	Extractors map[string]*Extraction   `protobuf:"bytes,2,rep,name=extractors,proto3" json:"extractors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Headers    map[string]*InjaTemplate `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to BodyTransformation:
	//	*TransformationTemplate_Body
	//	*TransformationTemplate_Passthrough
	//	*TransformationTemplate_MergeExtractorsToBody
	BodyTransformation   isTransformationTemplate_BodyTransformation `protobuf_oneof:"body_transformation"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *TransformationTemplate) Reset()         { *m = TransformationTemplate{} }
func (m *TransformationTemplate) String() string { return proto.CompactTextString(m) }
func (*TransformationTemplate) ProtoMessage()    {}
func (*TransformationTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{5}
}
func (m *TransformationTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformationTemplate.Unmarshal(m, b)
}
func (m *TransformationTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformationTemplate.Marshal(b, m, deterministic)
}
func (m *TransformationTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformationTemplate.Merge(m, src)
}
func (m *TransformationTemplate) XXX_Size() int {
	return xxx_messageInfo_TransformationTemplate.Size(m)
}
func (m *TransformationTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformationTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TransformationTemplate proto.InternalMessageInfo

type isTransformationTemplate_BodyTransformation interface {
	isTransformationTemplate_BodyTransformation()
	Equal(interface{}) bool
}

type TransformationTemplate_Body struct {
	Body *InjaTemplate `protobuf:"bytes,4,opt,name=body,proto3,oneof" json:"body,omitempty"`
}
type TransformationTemplate_Passthrough struct {
	Passthrough *Passthrough `protobuf:"bytes,5,opt,name=passthrough,proto3,oneof" json:"passthrough,omitempty"`
}
type TransformationTemplate_MergeExtractorsToBody struct {
	MergeExtractorsToBody *MergeExtractorsToBody `protobuf:"bytes,6,opt,name=merge_extractors_to_body,json=mergeExtractorsToBody,proto3,oneof" json:"merge_extractors_to_body,omitempty"`
}

func (*TransformationTemplate_Body) isTransformationTemplate_BodyTransformation()                  {}
func (*TransformationTemplate_Passthrough) isTransformationTemplate_BodyTransformation()           {}
func (*TransformationTemplate_MergeExtractorsToBody) isTransformationTemplate_BodyTransformation() {}

func (m *TransformationTemplate) GetBodyTransformation() isTransformationTemplate_BodyTransformation {
	if m != nil {
		return m.BodyTransformation
	}
	return nil
}

func (m *TransformationTemplate) GetAdvancedTemplates() bool {
	if m != nil {
		return m.AdvancedTemplates
	}
	return false
}

func (m *TransformationTemplate) GetExtractors() map[string]*Extraction {
	if m != nil {
		return m.Extractors
	}
	return nil
}

func (m *TransformationTemplate) GetHeaders() map[string]*InjaTemplate {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *TransformationTemplate) GetBody() *InjaTemplate {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Body); ok {
		return x.Body
	}
	return nil
}

func (m *TransformationTemplate) GetPassthrough() *Passthrough {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Passthrough); ok {
		return x.Passthrough
	}
	return nil
}

func (m *TransformationTemplate) GetMergeExtractorsToBody() *MergeExtractorsToBody {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_MergeExtractorsToBody); ok {
		return x.MergeExtractorsToBody
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransformationTemplate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransformationTemplate_Body)(nil),
		(*TransformationTemplate_Passthrough)(nil),
		(*TransformationTemplate_MergeExtractorsToBody)(nil),
	}
}

//
//custom functions:
//header_value(name) -> from the original headers
//extracted_value(name, index) -> from the extracted values
type InjaTemplate struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InjaTemplate) Reset()         { *m = InjaTemplate{} }
func (m *InjaTemplate) String() string { return proto.CompactTextString(m) }
func (*InjaTemplate) ProtoMessage()    {}
func (*InjaTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{6}
}
func (m *InjaTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjaTemplate.Unmarshal(m, b)
}
func (m *InjaTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjaTemplate.Marshal(b, m, deterministic)
}
func (m *InjaTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjaTemplate.Merge(m, src)
}
func (m *InjaTemplate) XXX_Size() int {
	return xxx_messageInfo_InjaTemplate.Size(m)
}
func (m *InjaTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_InjaTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_InjaTemplate proto.InternalMessageInfo

func (m *InjaTemplate) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Passthrough struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Passthrough) Reset()         { *m = Passthrough{} }
func (m *Passthrough) String() string { return proto.CompactTextString(m) }
func (*Passthrough) ProtoMessage()    {}
func (*Passthrough) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{7}
}
func (m *Passthrough) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Passthrough.Unmarshal(m, b)
}
func (m *Passthrough) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Passthrough.Marshal(b, m, deterministic)
}
func (m *Passthrough) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Passthrough.Merge(m, src)
}
func (m *Passthrough) XXX_Size() int {
	return xxx_messageInfo_Passthrough.Size(m)
}
func (m *Passthrough) XXX_DiscardUnknown() {
	xxx_messageInfo_Passthrough.DiscardUnknown(m)
}

var xxx_messageInfo_Passthrough proto.InternalMessageInfo

type MergeExtractorsToBody struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergeExtractorsToBody) Reset()         { *m = MergeExtractorsToBody{} }
func (m *MergeExtractorsToBody) String() string { return proto.CompactTextString(m) }
func (*MergeExtractorsToBody) ProtoMessage()    {}
func (*MergeExtractorsToBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{8}
}
func (m *MergeExtractorsToBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergeExtractorsToBody.Unmarshal(m, b)
}
func (m *MergeExtractorsToBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergeExtractorsToBody.Marshal(b, m, deterministic)
}
func (m *MergeExtractorsToBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergeExtractorsToBody.Merge(m, src)
}
func (m *MergeExtractorsToBody) XXX_Size() int {
	return xxx_messageInfo_MergeExtractorsToBody.Size(m)
}
func (m *MergeExtractorsToBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MergeExtractorsToBody.DiscardUnknown(m)
}

var xxx_messageInfo_MergeExtractorsToBody proto.InternalMessageInfo

type HeaderBodyTransform struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderBodyTransform) Reset()         { *m = HeaderBodyTransform{} }
func (m *HeaderBodyTransform) String() string { return proto.CompactTextString(m) }
func (*HeaderBodyTransform) ProtoMessage()    {}
func (*HeaderBodyTransform) Descriptor() ([]byte, []int) {
	return fileDescriptor_af1c648bbaa251cf, []int{9}
}
func (m *HeaderBodyTransform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderBodyTransform.Unmarshal(m, b)
}
func (m *HeaderBodyTransform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderBodyTransform.Marshal(b, m, deterministic)
}
func (m *HeaderBodyTransform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderBodyTransform.Merge(m, src)
}
func (m *HeaderBodyTransform) XXX_Size() int {
	return xxx_messageInfo_HeaderBodyTransform.Size(m)
}
func (m *HeaderBodyTransform) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderBodyTransform.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderBodyTransform proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FilterTransformations)(nil), "envoy.api.v2.filter.http.FilterTransformations")
	proto.RegisterType((*TransformationRule)(nil), "envoy.api.v2.filter.http.TransformationRule")
	proto.RegisterType((*RouteTransformations)(nil), "envoy.api.v2.filter.http.RouteTransformations")
	proto.RegisterType((*Transformation)(nil), "envoy.api.v2.filter.http.Transformation")
	proto.RegisterType((*Extraction)(nil), "envoy.api.v2.filter.http.Extraction")
	proto.RegisterType((*TransformationTemplate)(nil), "envoy.api.v2.filter.http.TransformationTemplate")
	proto.RegisterMapType((map[string]*Extraction)(nil), "envoy.api.v2.filter.http.TransformationTemplate.ExtractorsEntry")
	proto.RegisterMapType((map[string]*InjaTemplate)(nil), "envoy.api.v2.filter.http.TransformationTemplate.HeadersEntry")
	proto.RegisterType((*InjaTemplate)(nil), "envoy.api.v2.filter.http.InjaTemplate")
	proto.RegisterType((*Passthrough)(nil), "envoy.api.v2.filter.http.Passthrough")
	proto.RegisterType((*MergeExtractorsToBody)(nil), "envoy.api.v2.filter.http.MergeExtractorsToBody")
	proto.RegisterType((*HeaderBodyTransform)(nil), "envoy.api.v2.filter.http.HeaderBodyTransform")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation/transformation.proto", fileDescriptor_af1c648bbaa251cf)
}

var fileDescriptor_af1c648bbaa251cf = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6a, 0xdb, 0x48,
	0x14, 0xb6, 0xec, 0x38, 0x9b, 0x1c, 0x27, 0x9b, 0x64, 0xe2, 0x1f, 0xe1, 0x8b, 0x10, 0xc4, 0xee,
	0x62, 0x96, 0x8d, 0xb4, 0x78, 0x6f, 0x96, 0x90, 0x85, 0xe0, 0x25, 0xc5, 0xb9, 0x08, 0x94, 0x21,
	0xa4, 0x50, 0x0a, 0xee, 0x58, 0x9e, 0x48, 0x4a, 0x64, 0x8d, 0x3a, 0x33, 0x32, 0xf6, 0x0b, 0xf4,
	0xa2, 0x4f, 0xd2, 0x67, 0x68, 0x6f, 0xfa, 0x16, 0xbd, 0xef, 0x5d, 0x1f, 0xa1, 0x77, 0x45, 0x23,
	0xf9, 0x47, 0x8a, 0x0c, 0xce, 0x8d, 0x99, 0x73, 0xe6, 0x9c, 0xef, 0x3b, 0xe7, 0x3b, 0x67, 0x2c,
	0x70, 0x1c, 0x4f, 0xba, 0xd1, 0xd0, 0xb4, 0xd9, 0xd8, 0x12, 0xcc, 0x67, 0x67, 0x1e, 0xb3, 0x1c,
	0x9f, 0x31, 0x2b, 0xe4, 0xec, 0x81, 0xda, 0x52, 0x24, 0x16, 0x09, 0x3d, 0x8b, 0x4e, 0x25, 0xe5,
	0x01, 0xf1, 0x2d, 0x1a, 0x4c, 0xd8, 0x4c, 0x99, 0x81, 0xf0, 0x58, 0x20, 0x2c, 0xc9, 0x49, 0x20,
	0xee, 0x19, 0x1f, 0x13, 0xe9, 0xb1, 0x20, 0x67, 0x9a, 0x21, 0x67, 0x92, 0x21, 0x5d, 0x65, 0x99,
	0x24, 0xf4, 0xcc, 0x49, 0xd7, 0xbc, 0xf7, 0x7c, 0x49, 0xb9, 0xe9, 0x4a, 0x19, 0xb6, 0x4f, 0x12,
	0xbc, 0x98, 0x62, 0xd2, 0xb5, 0x38, 0x8b, 0x24, 0x4d, 0x7e, 0x93, 0xcc, 0x76, 0x6b, 0x42, 0x7c,
	0x6f, 0x44, 0x24, 0xb5, 0xe6, 0x87, 0xf4, 0xa2, 0xee, 0x30, 0x87, 0xa9, 0xa3, 0x15, 0x9f, 0x12,
	0xaf, 0xc1, 0xa0, 0xf1, 0x42, 0xa1, 0xdf, 0x66, 0xca, 0x10, 0xe8, 0x0e, 0x0e, 0xb2, 0x95, 0x09,
	0x5d, 0x3b, 0xad, 0x74, 0x6a, 0xdd, 0xbf, 0xcc, 0x75, 0xb5, 0x99, 0x59, 0x0c, 0x1c, 0xf9, 0x14,
	0xe7, 0x41, 0x8c, 0xcf, 0x1a, 0xa0, 0xa7, 0x71, 0xe8, 0x12, 0xaa, 0x63, 0x22, 0x6d, 0x57, 0xd7,
	0x4e, 0xb5, 0x4e, 0xad, 0x7b, 0x92, 0x25, 0x49, 0x1a, 0xc4, 0xf1, 0xef, 0x4d, 0x1c, 0xd5, 0x83,
	0x4f, 0xdf, 0xbf, 0x54, 0xaa, 0x1f, 0xb4, 0xf2, 0xa1, 0x86, 0x93, 0x44, 0x64, 0x43, 0x43, 0x85,
	0x0d, 0xf2, 0x65, 0x97, 0x15, 0xa2, 0xb9, 0xbe, 0x6c, 0x85, 0x9b, 0xeb, 0x1f, 0xd7, 0x79, 0x81,
	0xd7, 0x78, 0x5f, 0x86, 0x7a, 0x51, 0x38, 0x1a, 0x40, 0x93, 0xd3, 0x77, 0x11, 0x15, 0x32, 0xc7,
	0x9f, 0x36, 0xd4, 0xd9, 0x58, 0xb5, 0x46, 0x8a, 0x93, 0x75, 0xa3, 0x3f, 0xe1, 0xc8, 0xf6, 0x29,
	0xe1, 0x83, 0xa4, 0x49, 0x9b, 0xd8, 0x2e, 0xd5, 0x2b, 0xa7, 0x5a, 0x67, 0x07, 0x1f, 0xa8, 0x0b,
	0x55, 0xd6, 0xff, 0xb1, 0x1b, 0x11, 0x68, 0x71, 0x2a, 0x42, 0x16, 0x88, 0xbc, 0x1a, 0xa9, 0x18,
	0x9b, 0x57, 0xd3, 0x9c, 0x03, 0x65, 0xfd, 0xc6, 0x0f, 0x0d, 0x7e, 0xcd, 0x55, 0xf8, 0x08, 0xad,
	0x2c, 0xd9, 0x40, 0xd2, 0x71, 0xe8, 0x13, 0x49, 0x53, 0x0d, 0xfe, 0xde, 0x94, 0xf5, 0x36, 0xcd,
	0xeb, 0x97, 0x70, 0x53, 0x16, 0xde, 0xc4, 0xd3, 0x76, 0x29, 0x19, 0x51, 0x3e, 0x18, 0xb2, 0xd1,
	0x6c, 0xd9, 0x65, 0xda, 0xe0, 0xd9, 0x7a, 0xaa, 0xbe, 0x4a, 0xeb, 0xb1, 0xd1, 0x6c, 0x41, 0xda,
	0x2f, 0xe1, 0x63, 0xf7, 0xa9, 0xbb, 0xd7, 0x80, 0xe3, 0x7c, 0x47, 0xb3, 0x90, 0x1a, 0x77, 0x00,
	0x57, 0x53, 0xc9, 0x89, 0xad, 0xda, 0x6e, 0xc2, 0x76, 0x92, 0xab, 0xba, 0xdc, 0xc5, 0xa9, 0x85,
	0xea, 0x50, 0xe5, 0xd4, 0xa1, 0x53, 0x55, 0xd1, 0x2e, 0x4e, 0x0c, 0xd4, 0x86, 0x1d, 0x11, 0x0d,
	0x1d, 0xce, 0xa2, 0x50, 0x4d, 0x6f, 0x1f, 0x2f, 0x6c, 0xe3, 0x6b, 0x15, 0x9a, 0xc5, 0x42, 0xa0,
	0x33, 0x40, 0x64, 0x34, 0x21, 0x81, 0x4d, 0x47, 0x0b, 0x55, 0x85, 0x22, 0xdc, 0xc1, 0x47, 0xf3,
	0x9b, 0x79, 0xb4, 0x40, 0x6f, 0x01, 0x68, 0x52, 0x21, 0xe3, 0xf1, 0x03, 0x88, 0xdf, 0xed, 0xe5,
	0x73, 0xd5, 0x37, 0xaf, 0x16, 0x10, 0x57, 0x81, 0xe4, 0x33, 0xbc, 0x82, 0x89, 0x5e, 0xc1, 0x2f,
	0x49, 0x9f, 0x42, 0xaf, 0x28, 0xf8, 0xff, 0x9e, 0x0d, 0x9f, 0x0c, 0x22, 0xc5, 0x9e, 0xa3, 0xa1,
	0x0b, 0xd8, 0x8a, 0x27, 0xaa, 0x6f, 0xa9, 0x39, 0xfe, 0xb1, 0x1e, 0xf5, 0x3a, 0x78, 0x20, 0x2b,
	0x8b, 0xa2, 0xb2, 0xd0, 0x35, 0xd4, 0x42, 0x22, 0x84, 0x74, 0x39, 0x8b, 0x1c, 0x57, 0xaf, 0x2a,
	0x90, 0xdf, 0xd7, 0x83, 0xbc, 0x5c, 0x06, 0xf7, 0x4b, 0x78, 0x35, 0x17, 0x3d, 0x80, 0x3e, 0xa6,
	0xdc, 0xa1, 0x83, 0x65, 0xd7, 0x03, 0xc9, 0xd4, 0xba, 0xe9, 0xdb, 0x0a, 0xd7, 0x5a, 0x8f, 0x7b,
	0x13, 0x67, 0x2e, 0xf5, 0xbb, 0x65, 0xf1, 0x62, 0xf5, 0x4b, 0xb8, 0x31, 0x2e, 0xba, 0x68, 0xdb,
	0x70, 0x90, 0x13, 0x1b, 0x1d, 0x42, 0xe5, 0x91, 0xce, 0xd2, 0x9d, 0x8a, 0x8f, 0xe8, 0x1c, 0xaa,
	0x13, 0xe2, 0x47, 0x34, 0x5d, 0xf1, 0xdf, 0xd6, 0xb3, 0x2f, 0xb7, 0x13, 0x27, 0x29, 0xe7, 0xe5,
	0x7f, 0xb5, 0xf6, 0x10, 0xf6, 0x56, 0x25, 0x2f, 0x60, 0xb8, 0xc8, 0x32, 0x6c, 0x28, 0xfe, 0x0a,
	0x47, 0xfc, 0x62, 0xb2, 0xef, 0x31, 0xf9, 0xb7, 0x30, 0x60, 0x6f, 0x35, 0x03, 0x21, 0xd8, 0x92,
	0x74, 0x2a, 0x53, 0x6e, 0x75, 0x36, 0xf6, 0xa1, 0xb6, 0x32, 0x0d, 0xa3, 0x05, 0x8d, 0x42, 0x11,
	0x8d, 0x06, 0x1c, 0x17, 0x3c, 0xe1, 0xde, 0x9b, 0x8f, 0xdf, 0x4e, 0xb4, 0xd7, 0x77, 0x9b, 0x7d,
	0xa0, 0xc3, 0x47, 0xe7, 0x59, 0x1f, 0xe9, 0xe1, 0xb6, 0xfa, 0x5a, 0xfe, 0xf3, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x85, 0x4e, 0x2d, 0x34, 0x01, 0x08, 0x00, 0x00,
}

func (this *FilterTransformations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FilterTransformations)
	if !ok {
		that2, ok := that.(FilterTransformations)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TransformationRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationRule)
	if !ok {
		that2, ok := that.(TransformationRule)
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
	if !this.Match.Equal(that1.Match) {
		return false
	}
	if !this.RouteTransformations.Equal(that1.RouteTransformations) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteTransformations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTransformations)
	if !ok {
		that2, ok := that.(RouteTransformations)
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
	if !this.RequestTransformation.Equal(that1.RequestTransformation) {
		return false
	}
	if this.ClearRouteCache != that1.ClearRouteCache {
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
func (this *Transformation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformation)
	if !ok {
		that2, ok := that.(Transformation)
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
	if that1.TransformationType == nil {
		if this.TransformationType != nil {
			return false
		}
	} else if this.TransformationType == nil {
		return false
	} else if !this.TransformationType.Equal(that1.TransformationType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Transformation_TransformationTemplate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformation_TransformationTemplate)
	if !ok {
		that2, ok := that.(Transformation_TransformationTemplate)
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
	if !this.TransformationTemplate.Equal(that1.TransformationTemplate) {
		return false
	}
	return true
}
func (this *Transformation_HeaderBodyTransform) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformation_HeaderBodyTransform)
	if !ok {
		that2, ok := that.(Transformation_HeaderBodyTransform)
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
	if !this.HeaderBodyTransform.Equal(that1.HeaderBodyTransform) {
		return false
	}
	return true
}
func (this *Extraction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Extraction)
	if !ok {
		that2, ok := that.(Extraction)
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
	if this.Header != that1.Header {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	if this.Subgroup != that1.Subgroup {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TransformationTemplate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate)
	if !ok {
		that2, ok := that.(TransformationTemplate)
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
	if this.AdvancedTemplates != that1.AdvancedTemplates {
		return false
	}
	if len(this.Extractors) != len(that1.Extractors) {
		return false
	}
	for i := range this.Extractors {
		if !this.Extractors[i].Equal(that1.Extractors[i]) {
			return false
		}
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if that1.BodyTransformation == nil {
		if this.BodyTransformation != nil {
			return false
		}
	} else if this.BodyTransformation == nil {
		return false
	} else if !this.BodyTransformation.Equal(that1.BodyTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TransformationTemplate_Body) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate_Body)
	if !ok {
		that2, ok := that.(TransformationTemplate_Body)
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
	if !this.Body.Equal(that1.Body) {
		return false
	}
	return true
}
func (this *TransformationTemplate_Passthrough) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate_Passthrough)
	if !ok {
		that2, ok := that.(TransformationTemplate_Passthrough)
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
	if !this.Passthrough.Equal(that1.Passthrough) {
		return false
	}
	return true
}
func (this *TransformationTemplate_MergeExtractorsToBody) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate_MergeExtractorsToBody)
	if !ok {
		that2, ok := that.(TransformationTemplate_MergeExtractorsToBody)
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
	if !this.MergeExtractorsToBody.Equal(that1.MergeExtractorsToBody) {
		return false
	}
	return true
}
func (this *InjaTemplate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*InjaTemplate)
	if !ok {
		that2, ok := that.(InjaTemplate)
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
	if this.Text != that1.Text {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Passthrough) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Passthrough)
	if !ok {
		that2, ok := that.(Passthrough)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MergeExtractorsToBody) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MergeExtractorsToBody)
	if !ok {
		that2, ok := that.(MergeExtractorsToBody)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HeaderBodyTransform) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderBodyTransform)
	if !ok {
		that2, ok := that.(HeaderBodyTransform)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}