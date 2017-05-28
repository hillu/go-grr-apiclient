// Code generated by protoc-gen-go.
// source: semantic.proto
// DO NOT EDIT!

package apiclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SemanticDescriptor_Labels int32

const (
	SemanticDescriptor_ADVANCED SemanticDescriptor_Labels = 1
	// only.
	SemanticDescriptor_HIDDEN SemanticDescriptor_Labels = 2
)

var SemanticDescriptor_Labels_name = map[int32]string{
	1: "ADVANCED",
	2: "HIDDEN",
}
var SemanticDescriptor_Labels_value = map[string]int32{
	"ADVANCED": 1,
	"HIDDEN":   2,
}

func (x SemanticDescriptor_Labels) Enum() *SemanticDescriptor_Labels {
	p := new(SemanticDescriptor_Labels)
	*p = x
	return p
}
func (x SemanticDescriptor_Labels) String() string {
	return proto.EnumName(SemanticDescriptor_Labels_name, int32(x))
}
func (x *SemanticDescriptor_Labels) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SemanticDescriptor_Labels_value, data, "SemanticDescriptor_Labels")
	if err != nil {
		return err
	}
	*x = SemanticDescriptor_Labels(value)
	return nil
}
func (SemanticDescriptor_Labels) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor14, []int{1, 0}
}

type AnyValue struct {
	TypeUrl          *string `protobuf:"bytes,1,opt,name=type_url" json:"type_url,omitempty"`
	Value            []byte  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AnyValue) Reset()                    { *m = AnyValue{} }
func (m *AnyValue) String() string            { return proto.CompactTextString(m) }
func (*AnyValue) ProtoMessage()               {}
func (*AnyValue) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *AnyValue) GetTypeUrl() string {
	if m != nil && m.TypeUrl != nil {
		return *m.TypeUrl
	}
	return ""
}

func (m *AnyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SemanticDescriptor struct {
	// The semantic name of the SemanticValue contained in this field.
	Type *string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// The type of this field can be specified dynamically. This is the name of
	// the attribute to use to retrieve the SemanticValue class to be used for
	// parsing this field.
	DynamicType *string                     `protobuf:"bytes,5,opt,name=dynamic_type" json:"dynamic_type,omitempty"`
	Description *string                     `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Label       []SemanticDescriptor_Labels `protobuf:"varint,3,rep,name=label,enum=SemanticDescriptor_Labels" json:"label,omitempty"`
	// A friendly name for this field - to be used in GUIs etc.
	FriendlyName     *string `protobuf:"bytes,4,opt,name=friendly_name" json:"friendly_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SemanticDescriptor) Reset()                    { *m = SemanticDescriptor{} }
func (m *SemanticDescriptor) String() string            { return proto.CompactTextString(m) }
func (*SemanticDescriptor) ProtoMessage()               {}
func (*SemanticDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *SemanticDescriptor) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *SemanticDescriptor) GetDynamicType() string {
	if m != nil && m.DynamicType != nil {
		return *m.DynamicType
	}
	return ""
}

func (m *SemanticDescriptor) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *SemanticDescriptor) GetLabel() []SemanticDescriptor_Labels {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *SemanticDescriptor) GetFriendlyName() string {
	if m != nil && m.FriendlyName != nil {
		return *m.FriendlyName
	}
	return ""
}

type SemanticMessageDescriptor struct {
	// Describe the purpose of this protobuf.
	Description *string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// Certain RDFValues have union-like semantics. I.e. they are essentially
	// selectors of a number of other predefined RDFValues. They have a single
	// field that identifies the type of the selected "subvalue" and
	// nested fields corresponding to different selections. For examples:
	//
	// message FileFinderFilter {
	//   option (semantic) = {
	//     union_field: "filter_type"
	//   };
	//   enum Type {
	//     MODIFICATION_TIME = 0 [(description) = "Modification time"];
	//     ACCESS_TIME = 1 [(description) = "Access time"];
	//   }
	//
	//   optional Type filter_type = 1;
	//   optional FileFinderModificationTimeFilter modification_time = 2;
	//   optional FileFinderAccessTimeFilter access_time = 3;
	// }
	//
	// Field specified in union_field is used to determine which nested fields is
	// actually used. Field identified by union_field has to be Enum.
	// Corresponding nested values have to have names equal to enum values names,
	// but in lower case.
	//
	// At the moment setting union_field attribute on an RDFValue only affects the
	// UI rendering. RDFValues with union_field set are rendered with a select
	// box that allows users to switch between different available types.
	//
	// TODO(user): Suggestion from bgalehouse: We can just have "is_union"
	// boolean attribute instead of a string "union_field", because for
	// union-like structures we can understand the type of the structure by
	// inspecting its data fields.
	UnionField       *string `protobuf:"bytes,2,opt,name=union_field" json:"union_field,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SemanticMessageDescriptor) Reset()                    { *m = SemanticMessageDescriptor{} }
func (m *SemanticMessageDescriptor) String() string            { return proto.CompactTextString(m) }
func (*SemanticMessageDescriptor) ProtoMessage()               {}
func (*SemanticMessageDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *SemanticMessageDescriptor) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *SemanticMessageDescriptor) GetUnionField() string {
	if m != nil && m.UnionField != nil {
		return *m.UnionField
	}
	return ""
}

var E_SemType = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*SemanticDescriptor)(nil),
	Field:         51584972,
	Name:          "sem_type",
	Tag:           "bytes,51584972,opt,name=sem_type",
	Filename:      "semantic.proto",
}

var E_Semantic = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*SemanticMessageDescriptor)(nil),
	Field:         51584971,
	Name:          "semantic",
	Tag:           "bytes,51584971,opt,name=semantic",
	Filename:      "semantic.proto",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         48651165,
	Name:          "description",
	Tag:           "bytes,48651165,opt,name=description",
	Filename:      "semantic.proto",
}

var E_Label = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: ([]SemanticDescriptor_Labels)(nil),
	Field:         48651166,
	Name:          "label",
	Tag:           "varint,48651166,rep,name=label,enum=SemanticDescriptor_Labels",
	Filename:      "semantic.proto",
}

func init() {
	proto.RegisterType((*AnyValue)(nil), "AnyValue")
	proto.RegisterType((*SemanticDescriptor)(nil), "SemanticDescriptor")
	proto.RegisterType((*SemanticMessageDescriptor)(nil), "SemanticMessageDescriptor")
	proto.RegisterEnum("SemanticDescriptor_Labels", SemanticDescriptor_Labels_name, SemanticDescriptor_Labels_value)
	proto.RegisterExtension(E_SemType)
	proto.RegisterExtension(E_Semantic)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Label)
}

func init() { proto.RegisterFile("semantic.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0xff, 0x48, 0x6f, 0xd3, 0x52, 0xa6, 0x8a, 0xb1, 0x20, 0xc6, 0xac, 0x22, 0x42,
	0x0a, 0xdd, 0xd9, 0x5d, 0x31, 0xf1, 0x07, 0xb4, 0x15, 0x84, 0x6e, 0x43, 0xda, 0x4c, 0x4b, 0x60,
	0x32, 0x53, 0x32, 0x89, 0x90, 0x47, 0x71, 0xa1, 0x2f, 0xe1, 0xc2, 0xa5, 0x0f, 0x60, 0x1f, 0x4a,
	0x32, 0xcd, 0xf8, 0x17, 0x41, 0x97, 0xb9, 0x39, 0xe7, 0x9b, 0x73, 0xee, 0x85, 0x0e, 0xc7, 0x91,
	0x4f, 0x93, 0x70, 0x61, 0xaf, 0x63, 0x96, 0xb0, 0xbe, 0xb1, 0x62, 0x6c, 0x45, 0xf0, 0x40, 0x7c,
	0xcd, 0xd3, 0xe5, 0x20, 0xc0, 0x7c, 0x11, 0x87, 0xeb, 0x84, 0xc5, 0x5b, 0x85, 0x79, 0x02, 0xea,
	0x98, 0x66, 0x33, 0x9f, 0xa4, 0x18, 0x75, 0x41, 0x4d, 0xb2, 0x35, 0xf6, 0xd2, 0x98, 0xe8, 0x8a,
	0xa1, 0x58, 0x4d, 0xd4, 0x86, 0xfa, 0x7d, 0xfe, 0x4b, 0xaf, 0x18, 0x8a, 0xa5, 0x99, 0xcf, 0x0a,
	0xa0, 0xbb, 0xe2, 0x05, 0xe7, 0x83, 0x84, 0x34, 0xa8, 0xe5, 0xbe, 0xc2, 0xb3, 0x03, 0x5a, 0x90,
	0x51, 0x3f, 0x0a, 0x17, 0x9e, 0x98, 0xd6, 0xc5, 0xb4, 0x07, 0x2d, 0xf9, 0x76, 0xc8, 0xa8, 0xe0,
	0x35, 0xd1, 0x31, 0xd4, 0x89, 0x3f, 0xc7, 0x44, 0xaf, 0x1a, 0x55, 0xab, 0x33, 0xec, 0xdb, 0x65,
	0xb8, 0x7d, 0x9d, 0x0b, 0x38, 0xda, 0x85, 0xf6, 0x32, 0x0e, 0x31, 0x0d, 0x48, 0xe6, 0x51, 0x3f,
	0xc2, 0x7a, 0x2d, 0x27, 0x98, 0x26, 0x34, 0x0a, 0x81, 0x06, 0xea, 0xd8, 0x99, 0x8d, 0x27, 0x67,
	0xae, 0xd3, 0x55, 0x10, 0x40, 0xe3, 0xf2, 0xca, 0x71, 0xdc, 0x49, 0xb7, 0x62, 0xba, 0xb0, 0x2f,
	0xb9, 0x37, 0x98, 0x73, 0x7f, 0x85, 0xbf, 0x64, 0xff, 0x91, 0x4b, 0x91, 0x61, 0x53, 0x1a, 0x32,
	0xea, 0x2d, 0x43, 0x4c, 0x82, 0x6d, 0xd8, 0xd1, 0x05, 0xa8, 0x1c, 0x47, 0xa2, 0x13, 0x3a, 0xb0,
	0xb7, 0x8b, 0xb5, 0xe5, 0x62, 0xed, 0xf3, 0x5c, 0x39, 0x15, 0x10, 0xae, 0x6f, 0x5e, 0x37, 0xba,
	0xa1, 0x58, 0xad, 0x61, 0xef, 0x97, 0x46, 0xa3, 0x5b, 0x01, 0x12, 0x53, 0x74, 0x58, 0x02, 0x15,
	0x11, 0x25, 0xea, 0x4d, 0xa2, 0x3e, 0x97, 0x53, 0x2a, 0x31, 0x3a, 0xfd, 0x56, 0x02, 0x1d, 0x95,
	0xa0, 0x2e, 0x4d, 0x23, 0x71, 0x63, 0x89, 0x7d, 0x7c, 0x79, 0xd8, 0x13, 0xad, 0xa6, 0xc5, 0x09,
	0xfe, 0x63, 0x7a, 0xca, 0x4d, 0x7f, 0x1c, 0xea, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xda, 0x3e, 0x9a,
	0x5b, 0x83, 0x02, 0x00, 0x00,
}