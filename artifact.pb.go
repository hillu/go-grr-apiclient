// Code generated by protoc-gen-go.
// source: artifact.proto
// DO NOT EDIT!

package apiclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ArtifactSource_SourceType int32

const (
	ArtifactSource_COLLECTOR_TYPE_UNKNOWN ArtifactSource_SourceType = 0
	ArtifactSource_FILE                   ArtifactSource_SourceType = 1
	ArtifactSource_REGISTRY_KEY           ArtifactSource_SourceType = 2
	ArtifactSource_REGISTRY_VALUE         ArtifactSource_SourceType = 3
	ArtifactSource_WMI                    ArtifactSource_SourceType = 4
	// ARTIFACT has been deprecated in favor of ARTIFACT_GROUP.
	ArtifactSource_ARTIFACT       ArtifactSource_SourceType = 5
	ArtifactSource_PATH           ArtifactSource_SourceType = 6
	ArtifactSource_DIRECTORY      ArtifactSource_SourceType = 7
	ArtifactSource_ARTIFACT_GROUP ArtifactSource_SourceType = 8
	// TODO(user): these types will likely be separated out from artifacts in
	// the future
	ArtifactSource_GRR_CLIENT_ACTION ArtifactSource_SourceType = 40
	ArtifactSource_LIST_FILES        ArtifactSource_SourceType = 41
	ArtifactSource_ARTIFACT_FILES    ArtifactSource_SourceType = 42
	ArtifactSource_GREP              ArtifactSource_SourceType = 43
	ArtifactSource_COMMAND           ArtifactSource_SourceType = 45
	ArtifactSource_REKALL_PLUGIN     ArtifactSource_SourceType = 46
)

var ArtifactSource_SourceType_name = map[int32]string{
	0:  "COLLECTOR_TYPE_UNKNOWN",
	1:  "FILE",
	2:  "REGISTRY_KEY",
	3:  "REGISTRY_VALUE",
	4:  "WMI",
	5:  "ARTIFACT",
	6:  "PATH",
	7:  "DIRECTORY",
	8:  "ARTIFACT_GROUP",
	40: "GRR_CLIENT_ACTION",
	41: "LIST_FILES",
	42: "ARTIFACT_FILES",
	43: "GREP",
	45: "COMMAND",
	46: "REKALL_PLUGIN",
}
var ArtifactSource_SourceType_value = map[string]int32{
	"COLLECTOR_TYPE_UNKNOWN": 0,
	"FILE":              1,
	"REGISTRY_KEY":      2,
	"REGISTRY_VALUE":    3,
	"WMI":               4,
	"ARTIFACT":          5,
	"PATH":              6,
	"DIRECTORY":         7,
	"ARTIFACT_GROUP":    8,
	"GRR_CLIENT_ACTION": 40,
	"LIST_FILES":        41,
	"ARTIFACT_FILES":    42,
	"GREP":              43,
	"COMMAND":           45,
	"REKALL_PLUGIN":     46,
}

func (x ArtifactSource_SourceType) Enum() *ArtifactSource_SourceType {
	p := new(ArtifactSource_SourceType)
	*p = x
	return p
}
func (x ArtifactSource_SourceType) String() string {
	return proto.EnumName(ArtifactSource_SourceType_name, int32(x))
}
func (x *ArtifactSource_SourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ArtifactSource_SourceType_value, data, "ArtifactSource_SourceType")
	if err != nil {
		return err
	}
	*x = ArtifactSource_SourceType(value)
	return nil
}
func (ArtifactSource_SourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

// Proto representation of an ArtifactSource.
type ArtifactSource struct {
	Type             *ArtifactSource_SourceType `protobuf:"varint,1,opt,name=type,enum=ArtifactSource_SourceType" json:"type,omitempty"`
	Attributes       *Dict                      `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
	Conditions       []string                   `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
	ReturnedTypes    []string                   `protobuf:"bytes,4,rep,name=returned_types,json=returnedTypes" json:"returned_types,omitempty"`
	SupportedOs      []string                   `protobuf:"bytes,5,rep,name=supported_os,json=supportedOs" json:"supported_os,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ArtifactSource) Reset()                    { *m = ArtifactSource{} }
func (m *ArtifactSource) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSource) ProtoMessage()               {}
func (*ArtifactSource) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ArtifactSource) GetType() ArtifactSource_SourceType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ArtifactSource_COLLECTOR_TYPE_UNKNOWN
}

func (m *ArtifactSource) GetAttributes() *Dict {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *ArtifactSource) GetConditions() []string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *ArtifactSource) GetReturnedTypes() []string {
	if m != nil {
		return m.ReturnedTypes
	}
	return nil
}

func (m *ArtifactSource) GetSupportedOs() []string {
	if m != nil {
		return m.SupportedOs
	}
	return nil
}

// Proto representation of an artifact.
type Artifact struct {
	Name             *string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Conditions       []string          `protobuf:"bytes,2,rep,name=conditions" json:"conditions,omitempty"`
	Doc              *string           `protobuf:"bytes,3,opt,name=doc" json:"doc,omitempty"`
	Labels           []string          `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty"`
	SupportedOs      []string          `protobuf:"bytes,5,rep,name=supported_os,json=supportedOs" json:"supported_os,omitempty"`
	Urls             []string          `protobuf:"bytes,6,rep,name=urls" json:"urls,omitempty"`
	Provides         []string          `protobuf:"bytes,8,rep,name=provides" json:"provides,omitempty"`
	Sources          []*ArtifactSource `protobuf:"bytes,9,rep,name=sources" json:"sources,omitempty"`
	ErrorMessage     *string           `protobuf:"bytes,10,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Artifact) Reset()                    { *m = Artifact{} }
func (m *Artifact) String() string            { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()               {}
func (*Artifact) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Artifact) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Artifact) GetConditions() []string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *Artifact) GetDoc() string {
	if m != nil && m.Doc != nil {
		return *m.Doc
	}
	return ""
}

func (m *Artifact) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Artifact) GetSupportedOs() []string {
	if m != nil {
		return m.SupportedOs
	}
	return nil
}

func (m *Artifact) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *Artifact) GetProvides() []string {
	if m != nil {
		return m.Provides
	}
	return nil
}

func (m *Artifact) GetSources() []*ArtifactSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Artifact) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

type ArtifactProcessorDescriptor struct {
	Name             *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description      *string  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	OutputTypes      []string `protobuf:"bytes,3,rep,name=output_types,json=outputTypes" json:"output_types,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ArtifactProcessorDescriptor) Reset()                    { *m = ArtifactProcessorDescriptor{} }
func (m *ArtifactProcessorDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ArtifactProcessorDescriptor) ProtoMessage()               {}
func (*ArtifactProcessorDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ArtifactProcessorDescriptor) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ArtifactProcessorDescriptor) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ArtifactProcessorDescriptor) GetOutputTypes() []string {
	if m != nil {
		return m.OutputTypes
	}
	return nil
}

type ArtifactDescriptor struct {
	Artifact         *Artifact                      `protobuf:"bytes,1,opt,name=artifact" json:"artifact,omitempty"`
	Dependencies     []string                       `protobuf:"bytes,2,rep,name=dependencies" json:"dependencies,omitempty"`
	PathDependencies []string                       `protobuf:"bytes,3,rep,name=path_dependencies,json=pathDependencies" json:"path_dependencies,omitempty"`
	ArtifactSource   *string                        `protobuf:"bytes,4,opt,name=artifact_source,json=artifactSource" json:"artifact_source,omitempty"`
	Processors       []*ArtifactProcessorDescriptor `protobuf:"bytes,5,rep,name=processors" json:"processors,omitempty"`
	IsCustom         *bool                          `protobuf:"varint,6,opt,name=is_custom,json=isCustom" json:"is_custom,omitempty"`
	ErrorMessage     *string                        `protobuf:"bytes,7,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *ArtifactDescriptor) Reset()                    { *m = ArtifactDescriptor{} }
func (m *ArtifactDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ArtifactDescriptor) ProtoMessage()               {}
func (*ArtifactDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ArtifactDescriptor) GetArtifact() *Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

func (m *ArtifactDescriptor) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *ArtifactDescriptor) GetPathDependencies() []string {
	if m != nil {
		return m.PathDependencies
	}
	return nil
}

func (m *ArtifactDescriptor) GetArtifactSource() string {
	if m != nil && m.ArtifactSource != nil {
		return *m.ArtifactSource
	}
	return ""
}

func (m *ArtifactDescriptor) GetProcessors() []*ArtifactProcessorDescriptor {
	if m != nil {
		return m.Processors
	}
	return nil
}

func (m *ArtifactDescriptor) GetIsCustom() bool {
	if m != nil && m.IsCustom != nil {
		return *m.IsCustom
	}
	return false
}

func (m *ArtifactDescriptor) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*ArtifactSource)(nil), "ArtifactSource")
	proto.RegisterType((*Artifact)(nil), "Artifact")
	proto.RegisterType((*ArtifactProcessorDescriptor)(nil), "ArtifactProcessorDescriptor")
	proto.RegisterType((*ArtifactDescriptor)(nil), "ArtifactDescriptor")
	proto.RegisterEnum("ArtifactSource_SourceType", ArtifactSource_SourceType_name, ArtifactSource_SourceType_value)
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x96, 0xed, 0x72, 0xdb, 0x44,
	0x17, 0xc7, 0x1f, 0xc7, 0x79, 0xb1, 0xd7, 0x8e, 0xab, 0x9e, 0x67, 0xfa, 0x3c, 0xa6, 0x2d, 0x65,
	0x27, 0x40, 0xeb, 0xd2, 0x56, 0x9d, 0x66, 0xfa, 0x42, 0x69, 0xa7, 0x1d, 0xbf, 0xd5, 0x35, 0x71,
	0x6c, 0x8f, 0xe2, 0xb4, 0x84, 0x19, 0xd0, 0xc8, 0xd2, 0xda, 0xde, 0x22, 0x6b, 0x55, 0xed, 0x2a,
	0x19, 0xf3, 0x0d, 0xee, 0x80, 0x2b, 0x80, 0x8b, 0xe1, 0x1a, 0xb8, 0x00, 0xb8, 0x0d, 0x3e, 0x30,
	0xbb, 0x6b, 0x59, 0x56, 0x60, 0xf2, 0x8d, 0x4f, 0x6d, 0xbc, 0xe7, 0xfc, 0xfe, 0x67, 0xcf, 0xdb,
	0x0a, 0x55, 0x9c, 0x48, 0xd0, 0x89, 0xe3, 0x0a, 0x33, 0x8c, 0x98, 0x60, 0x57, 0xd1, 0x3b, 0x36,
	0xe6, 0xcb, 0xff, 0x57, 0x38, 0x99, 0x3b, 0x81, 0xa0, 0xae, 0xfe, 0x7b, 0xef, 0xb7, 0x6d, 0x54,
	0xa9, 0x2f, 0xcd, 0x8f, 0x58, 0x1c, 0xb9, 0x04, 0x0e, 0xd0, 0xa6, 0x58, 0x84, 0xa4, 0x9a, 0xc3,
	0xb9, 0x5a, 0x65, 0xff, 0xaa, 0x99, 0x3d, 0x36, 0xf5, 0x3f, 0xa3, 0x45, 0x48, 0x1a, 0xd7, 0x7e,
	0xff, 0xf3, 0x8f, 0x5f, 0x73, 0x57, 0xe0, 0xbf, 0xa3, 0x19, 0xc1, 0xd2, 0x07, 0xb3, 0x09, 0xe6,
	0xda, 0xcc, 0x52, 0x10, 0x98, 0x21, 0xe4, 0x08, 0x11, 0xd1, 0x71, 0x2c, 0x08, 0xaf, 0x6e, 0xe0,
	0x5c, 0xad, 0xb4, 0xbf, 0x65, 0xb6, 0xa8, 0x2b, 0x1a, 0x75, 0xe5, 0xfd, 0x0c, 0x9e, 0x4a, 0xef,
	0xd4, 0x08, 0x8b, 0x99, 0x23, 0xb0, 0x47, 0xb8, 0x1b, 0xd1, 0x31, 0xc1, 0x62, 0x46, 0x12, 0x24,
	0x26, 0xe6, 0xd4, 0xc4, 0x13, 0xea, 0x13, 0x1c, 0x3a, 0x62, 0xc6, 0x4d, 0x6b, 0x8d, 0x0d, 0xa7,
	0x08, 0xb9, 0x2c, 0xf0, 0xa8, 0xa0, 0x2c, 0xe0, 0xd5, 0x3c, 0xce, 0xd7, 0x8a, 0x8d, 0x37, 0x4a,
	0x62, 0x08, 0xfd, 0xc1, 0xf8, 0x1d, 0x71, 0x85, 0xf4, 0x15, 0x24, 0xc2, 0xa9, 0x5d, 0x22, 0xe6,
	0x52, 0x8f, 0x60, 0x3a, 0xc1, 0x62, 0x46, 0xf9, 0x52, 0x0e, 0x3b, 0x61, 0xe8, 0x53, 0x19, 0x0f,
	0xc3, 0x0e, 0x9e, 0xd2, 0x53, 0x12, 0x60, 0xbe, 0xe0, 0x82, 0xcc, 0x4d, 0x6b, 0x4d, 0x09, 0x7e,
	0xc9, 0xa1, 0x4a, 0x44, 0x44, 0x1c, 0x05, 0xc4, 0xb3, 0xe5, 0x9d, 0x79, 0x75, 0x53, 0x89, 0x2f,
	0x94, 0x38, 0x87, 0xf7, 0x75, 0xec, 0x53, 0x2e, 0x64, 0x6e, 0xd4, 0xb1, 0xd6, 0x9c, 0x3b, 0x0b,
	0x3c, 0x26, 0x38, 0xf1, 0xc4, 0xe3, 0xc5, 0xba, 0xb8, 0x89, 0xeb, 0xc1, 0x42, 0xcc, 0x68, 0x30,
	0x4d, 0x2d, 0x94, 0x1b, 0xe5, 0x38, 0x60, 0x02, 0xd3, 0x40, 0x5b, 0x2b, 0xf0, 0x19, 0xf5, 0x7d,
	0xc9, 0xd2, 0xb7, 0x23, 0x9e, 0x69, 0xed, 0x26, 0x4e, 0xb2, 0x4c, 0x1c, 0x02, 0x54, 0xe6, 0x71,
	0x18, 0xb2, 0x48, 0x10, 0xcf, 0x66, 0xbc, 0xba, 0xa5, 0xe2, 0x3b, 0x50, 0xf1, 0xb5, 0xa1, 0x99,
	0xc6, 0xc7, 0x42, 0x12, 0x39, 0x42, 0xaa, 0xea, 0x9b, 0xaa, 0xcb, 0x9f, 0xcd, 0xa8, 0x3b, 0xcb,
	0x64, 0x86, 0xcf, 0x58, 0xec, 0x7b, 0x52, 0x52, 0xe7, 0xc8, 0x33, 0xad, 0xd2, 0x4a, 0x60, 0xc0,
	0xf7, 0x7e, 0xd8, 0x40, 0x28, 0x6d, 0x13, 0xb8, 0x8a, 0xfe, 0xd7, 0x1c, 0xf4, 0x7a, 0xed, 0xe6,
	0x68, 0x60, 0xd9, 0xa3, 0x93, 0x61, 0xdb, 0x3e, 0xee, 0x1f, 0xf4, 0x07, 0x6f, 0xfb, 0xc6, 0x7f,
	0xa0, 0x80, 0x36, 0x5f, 0x75, 0x7b, 0x6d, 0x23, 0x07, 0x06, 0x2a, 0x5b, 0xed, 0x4e, 0xf7, 0x68,
	0x64, 0x9d, 0xd8, 0x07, 0xed, 0x13, 0x63, 0x03, 0x00, 0x55, 0x56, 0xbf, 0xbc, 0xa9, 0xf7, 0x8e,
	0xdb, 0x46, 0x1e, 0x76, 0x50, 0xfe, 0xed, 0x61, 0xd7, 0xd8, 0x84, 0x32, 0x2a, 0xd4, 0xad, 0x51,
	0xf7, 0x55, 0xbd, 0x39, 0x32, 0xb6, 0x24, 0x66, 0x58, 0x1f, 0xbd, 0x36, 0xb6, 0x61, 0x17, 0x15,
	0x5b, 0x5d, 0x4b, 0x69, 0x9d, 0x18, 0x3b, 0x92, 0x91, 0x98, 0xd9, 0x1d, 0x6b, 0x70, 0x3c, 0x34,
	0x0a, 0x70, 0x05, 0x5d, 0xee, 0x58, 0x96, 0xdd, 0xec, 0x75, 0xdb, 0xfd, 0x91, 0x5d, 0x6f, 0x8e,
	0xba, 0x83, 0xbe, 0x51, 0x83, 0x0a, 0x42, 0xbd, 0xee, 0xd1, 0xc8, 0x96, 0xf1, 0x1c, 0x19, 0xb7,
	0x33, 0xae, 0xfa, 0xb7, 0xcf, 0xa4, 0x4e, 0xc7, 0x6a, 0x0f, 0x8d, 0x3b, 0x50, 0x42, 0x3b, 0xcd,
	0xc1, 0xe1, 0x61, 0xbd, 0xdf, 0x32, 0xee, 0xc1, 0x65, 0xb4, 0x6b, 0xb5, 0x0f, 0xea, 0xbd, 0x9e,
	0x3d, 0xec, 0x1d, 0x77, 0xba, 0x7d, 0xc3, 0xdc, 0xfb, 0x71, 0x1b, 0x15, 0x92, 0xc9, 0x81, 0x01,
	0xda, 0x0c, 0x9c, 0xb9, 0x1e, 0xa9, 0x62, 0xe3, 0x99, 0x4a, 0xfc, 0x23, 0x54, 0x4e, 0xce, 0xfb,
	0xce, 0x9c, 0xc0, 0xa7, 0x1d, 0x9f, 0x8d, 0x1d, 0xdf, 0x5f, 0xe0, 0x38, 0xa0, 0xef, 0x63, 0x82,
	0xa5, 0x87, 0x6a, 0x1a, 0x39, 0x1e, 0xc9, 0x44, 0x5b, 0x0a, 0x04, 0x5e, 0xa6, 0xd9, 0x37, 0x54,
	0x3d, 0x5b, 0x0a, 0xfb, 0x02, 0x9e, 0xa7, 0xf5, 0xbc, 0xa8, 0xd1, 0x53, 0x68, 0x52, 0xcf, 0x28,
	0x0e, 0xb2, 0xad, 0xfd, 0x18, 0xe5, 0x3d, 0xe6, 0x56, 0xf3, 0x2a, 0xea, 0x4f, 0x14, 0xfe, 0x06,
	0x5c, 0x6f, 0x31, 0x17, 0x73, 0x11, 0xc9, 0x2e, 0x99, 0xb0, 0xe8, 0x5c, 0x78, 0xd2, 0x01, 0xba,
	0x68, 0xdb, 0x77, 0xc6, 0xc4, 0x4f, 0x26, 0xe1, 0x81, 0x72, 0xbd, 0x03, 0xb7, 0xd3, 0xc8, 0xf4,
	0x79, 0x36, 0x8c, 0x31, 0xf1, 0x59, 0x30, 0x95, 0x3d, 0x67, 0x5a, 0x4b, 0x00, 0x7c, 0xf3, 0x8f,
	0xad, 0xfb, 0x85, 0x02, 0x3e, 0x84, 0xfd, 0x0b, 0x5b, 0x37, 0x73, 0x45, 0x0d, 0xe1, 0xd9, 0x4e,
	0x95, 0xbb, 0x2e, 0x8e, 0x7c, 0x5e, 0xdd, 0x56, 0xd8, 0x27, 0x0a, 0xfb, 0x00, 0xee, 0xa7, 0x58,
	0x79, 0xaa, 0x73, 0x37, 0x23, 0x7e, 0x88, 0x3d, 0xe6, 0xc6, 0x73, 0x12, 0x88, 0xf3, 0x45, 0x91,
	0x66, 0xf0, 0x35, 0x2a, 0x84, 0x11, 0x3b, 0xa5, 0x1e, 0xe1, 0xd5, 0x82, 0x02, 0xbe, 0x50, 0xc0,
	0xcf, 0xe1, 0x71, 0x0a, 0xfc, 0x2e, 0x60, 0x67, 0x3e, 0xf1, 0xa6, 0x64, 0xec, 0x70, 0x82, 0x4f,
	0x1d, 0x3f, 0x56, 0x1b, 0x81, 0xf2, 0x34, 0xd6, 0x04, 0x62, 0x5a, 0x2b, 0x1e, 0xf4, 0xd1, 0x8e,
	0x1e, 0x3e, 0x5e, 0x2d, 0xe2, 0x7c, 0xad, 0xb4, 0x7f, 0xe9, 0xdc, 0x5e, 0x6e, 0x7c, 0xac, 0xb4,
	0x3e, 0x84, 0x6b, 0xa9, 0x56, 0x7a, 0x7b, 0xed, 0x6a, 0x5a, 0x09, 0x04, 0x3c, 0xb4, 0x4b, 0xa2,
	0x88, 0x45, 0xf6, 0x9c, 0x70, 0xee, 0x4c, 0x49, 0x15, 0xa9, 0x22, 0xbf, 0x54, 0x90, 0xa7, 0xf0,
	0x44, 0xee, 0x64, 0x65, 0x80, 0x97, 0x06, 0xaa, 0xd6, 0x09, 0x70, 0x99, 0x92, 0x89, 0x43, 0x7d,
	0xe2, 0xc9, 0x4b, 0x50, 0xcf, 0x91, 0x2d, 0x63, 0x5a, 0x65, 0xe5, 0x74, 0xa8, 0x7d, 0xf6, 0x7e,
	0xda, 0x40, 0xd7, 0x92, 0x30, 0x87, 0x11, 0x73, 0x09, 0xe7, 0x2c, 0x6a, 0xa9, 0xd5, 0x1e, 0x0a,
	0x16, 0x41, 0x23, 0x33, 0x17, 0xa6, 0x12, 0xaf, 0xc1, 0xcd, 0x95, 0xe9, 0x2d, 0xae, 0xa7, 0xc0,
	0xe1, 0x38, 0x22, 0x53, 0xca, 0xd5, 0x7a, 0x93, 0xcb, 0xaf, 0x63, 0x59, 0xc9, 0x28, 0xbc, 0x46,
	0x25, 0x6f, 0x49, 0xa4, 0x2c, 0x50, 0x4f, 0x4c, 0xb1, 0x71, 0x53, 0xa1, 0x30, 0xdc, 0x68, 0xa5,
	0x47, 0x7a, 0x96, 0x28, 0x97, 0xf9, 0xd5, 0x78, 0xd3, 0x5a, 0x77, 0x05, 0x0f, 0x95, 0x59, 0x2c,
	0xc2, 0x58, 0x2c, 0xd7, 0xb8, 0x7e, 0x43, 0x32, 0xcf, 0x54, 0xf2, 0x7a, 0xae, 0x6f, 0x72, 0xd7,
	0x09, 0xe4, 0x2a, 0x0c, 0x23, 0xe6, 0xc5, 0x6e, 0xb2, 0xc9, 0x49, 0x46, 0x45, 0x63, 0xd5, 0x32,
	0xde, 0xfb, 0x79, 0x0b, 0x41, 0x92, 0x93, 0xb5, 0x54, 0xbc, 0x44, 0x85, 0x24, 0xbb, 0x2a, 0x1d,
	0xa5, 0xfd, 0xe2, 0xaa, 0xc2, 0x8d, 0xaa, 0x8a, 0x01, 0xc0, 0x48, 0x7e, 0xc1, 0x54, 0x70, 0xe2,
	0x4f, 0x4c, 0x6b, 0xe5, 0x04, 0x5f, 0xa1, 0xb2, 0x47, 0x42, 0x12, 0x78, 0x24, 0x70, 0x29, 0x49,
	0x96, 0xc2, 0x43, 0xe5, 0x69, 0xc2, 0x5d, 0xb9, 0x63, 0xf8, 0x7a, 0x53, 0x9c, 0xef, 0x3a, 0xed,
	0xcc, 0xb1, 0xaa, 0xe2, 0x3a, 0x09, 0xc6, 0xe8, 0xb2, 0x7c, 0x6f, 0xed, 0x0c, 0x5e, 0x27, 0xe7,
	0x91, 0xc2, 0xdf, 0x87, 0x7b, 0x2b, 0xfc, 0x41, 0x03, 0x33, 0xf5, 0xd8, 0x5e, 0xc4, 0x37, 0x24,
	0xaf, 0xb5, 0xae, 0xf1, 0x1a, 0x5d, 0x4a, 0x2c, 0x6d, 0xdd, 0xa3, 0xd5, 0x4d, 0x55, 0xc9, 0x8f,
	0x94, 0xc2, 0x07, 0xf0, 0xff, 0xe4, 0xea, 0xb7, 0x38, 0xfe, 0xf2, 0x68, 0xd0, 0x5f, 0x7d, 0x67,
	0xac, 0xbe, 0x76, 0x96, 0x9f, 0x2f, 0xdf, 0x23, 0xb4, 0x4a, 0xbd, 0xde, 0x17, 0xa5, 0xfd, 0xeb,
	0xe6, 0x05, 0x5d, 0xd8, 0x78, 0xae, 0x24, 0x1e, 0xc3, 0xc3, 0xd5, 0xe1, 0xb2, 0xb2, 0xea, 0x61,
	0x5d, 0xc2, 0xb2, 0xb7, 0xb9, 0xc5, 0xb1, 0xae, 0xa9, 0x69, 0xad, 0xa9, 0xc1, 0xb7, 0xa8, 0x48,
	0xb9, 0xed, 0xc6, 0x5c, 0xb0, 0x79, 0x75, 0x1b, 0xe7, 0x6a, 0x85, 0xb4, 0x7d, 0xba, 0x13, 0x3c,
	0x8a, 0x62, 0x72, 0xf7, 0x5c, 0x5a, 0xce, 0x1c, 0x8e, 0xe7, 0x4e, 0x10, 0xeb, 0xcd, 0x1f, 0xfa,
	0xcc, 0xf1, 0xd2, 0x1e, 0x8a, 0x39, 0x89, 0x4c, 0xab, 0x40, 0x79, 0x53, 0x21, 0xff, 0x3e, 0xb5,
	0x3b, 0xff, 0xc2, 0xd4, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x70, 0xc1, 0x1d, 0x73, 0x40, 0x0a,
	0x00, 0x00,
}
