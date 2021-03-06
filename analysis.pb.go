// Code generated by protoc-gen-go.
// source: analysis.proto
// DO NOT EDIT!

package apiclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A graph is a way to represent data.
type Sample struct {
	Label            *string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	XValue           *uint64 `protobuf:"varint,2,opt,name=x_value,json=xValue" json:"x_value,omitempty"`
	YValue           *uint64 `protobuf:"varint,3,opt,name=y_value,json=yValue" json:"y_value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Sample) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Sample) GetXValue() uint64 {
	if m != nil && m.XValue != nil {
		return *m.XValue
	}
	return 0
}

func (m *Sample) GetYValue() uint64 {
	if m != nil && m.YValue != nil {
		return *m.YValue
	}
	return 0
}

type SampleFloat struct {
	Label            *string  `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	XValue           *float32 `protobuf:"fixed32,2,opt,name=x_value,json=xValue" json:"x_value,omitempty"`
	YValue           *float32 `protobuf:"fixed32,3,opt,name=y_value,json=yValue" json:"y_value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SampleFloat) Reset()                    { *m = SampleFloat{} }
func (m *SampleFloat) String() string            { return proto.CompactTextString(m) }
func (*SampleFloat) ProtoMessage()               {}
func (*SampleFloat) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SampleFloat) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *SampleFloat) GetXValue() float32 {
	if m != nil && m.XValue != nil {
		return *m.XValue
	}
	return 0
}

func (m *SampleFloat) GetYValue() float32 {
	if m != nil && m.YValue != nil {
		return *m.YValue
	}
	return 0
}

type Graph struct {
	Title            *string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Xtitle           *string   `protobuf:"bytes,2,opt,name=xtitle" json:"xtitle,omitempty"`
	Ytitle           *string   `protobuf:"bytes,3,opt,name=ytitle" json:"ytitle,omitempty"`
	Data             []*Sample `protobuf:"bytes,4,rep,name=data" json:"data,omitempty"`
	XScale           *uint32   `protobuf:"varint,5,opt,name=x_scale,json=xScale,def=1" json:"x_scale,omitempty"`
	YScale           *uint32   `protobuf:"varint,6,opt,name=y_scale,json=yScale,def=1" json:"y_scale,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Graph) Reset()                    { *m = Graph{} }
func (m *Graph) String() string            { return proto.CompactTextString(m) }
func (*Graph) ProtoMessage()               {}
func (*Graph) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

const Default_Graph_XScale uint32 = 1
const Default_Graph_YScale uint32 = 1

func (m *Graph) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Graph) GetXtitle() string {
	if m != nil && m.Xtitle != nil {
		return *m.Xtitle
	}
	return ""
}

func (m *Graph) GetYtitle() string {
	if m != nil && m.Ytitle != nil {
		return *m.Ytitle
	}
	return ""
}

func (m *Graph) GetData() []*Sample {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Graph) GetXScale() uint32 {
	if m != nil && m.XScale != nil {
		return *m.XScale
	}
	return Default_Graph_XScale
}

func (m *Graph) GetYScale() uint32 {
	if m != nil && m.YScale != nil {
		return *m.YScale
	}
	return Default_Graph_YScale
}

type GraphFloat struct {
	Title            *string        `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Xtitle           *string        `protobuf:"bytes,2,opt,name=xtitle" json:"xtitle,omitempty"`
	Ytitle           *string        `protobuf:"bytes,3,opt,name=ytitle" json:"ytitle,omitempty"`
	Data             []*SampleFloat `protobuf:"bytes,4,rep,name=data" json:"data,omitempty"`
	XScale           *uint32        `protobuf:"varint,5,opt,name=x_scale,json=xScale,def=1" json:"x_scale,omitempty"`
	YScale           *uint32        `protobuf:"varint,6,opt,name=y_scale,json=yScale,def=1" json:"y_scale,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GraphFloat) Reset()                    { *m = GraphFloat{} }
func (m *GraphFloat) String() string            { return proto.CompactTextString(m) }
func (*GraphFloat) ProtoMessage()               {}
func (*GraphFloat) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

const Default_GraphFloat_XScale uint32 = 1
const Default_GraphFloat_YScale uint32 = 1

func (m *GraphFloat) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *GraphFloat) GetXtitle() string {
	if m != nil && m.Xtitle != nil {
		return *m.Xtitle
	}
	return ""
}

func (m *GraphFloat) GetYtitle() string {
	if m != nil && m.Ytitle != nil {
		return *m.Ytitle
	}
	return ""
}

func (m *GraphFloat) GetData() []*SampleFloat {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GraphFloat) GetXScale() uint32 {
	if m != nil && m.XScale != nil {
		return *m.XScale
	}
	return Default_GraphFloat_XScale
}

func (m *GraphFloat) GetYScale() uint32 {
	if m != nil && m.YScale != nil {
		return *m.YScale
	}
	return Default_GraphFloat_YScale
}

// The following relate to the timelining functionality.
type Event struct {
	Timestamp *uint64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Source    *string `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	// The name of the object which this event is talking about.
	Subject *string `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	// The type of this event.
	Type *string    `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Stat *StatEntry `protobuf:"bytes,5,opt,name=stat" json:"stat,omitempty"`
	// The sequential number of this event as stored in the time series.
	Id               *uint32 `protobuf:"varint,6,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Event) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Event) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *Event) GetSubject() string {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return ""
}

func (m *Event) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Event) GetStat() *StatEntry {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *Event) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Sample)(nil), "Sample")
	proto.RegisterType((*SampleFloat)(nil), "SampleFloat")
	proto.RegisterType((*Graph)(nil), "Graph")
	proto.RegisterType((*GraphFloat)(nil), "GraphFloat")
	proto.RegisterType((*Event)(nil), "Event")
}

func init() { proto.RegisterFile("analysis.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x90, 0x41, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0xe5, 0xc9, 0x64, 0xa2, 0xbe, 0x40, 0x16, 0x56, 0x05, 0xa3, 0x22, 0x21, 0x2b, 0x6c,
	0xc2, 0x66, 0xa4, 0x76, 0x09, 0x2b, 0xaa, 0x34, 0xdd, 0x21, 0xe4, 0x50, 0x24, 0x56, 0xe8, 0x25,
	0x71, 0x1b, 0x57, 0xce, 0x78, 0x34, 0x7e, 0x89, 0xe2, 0x23, 0x70, 0x0d, 0xd6, 0x1c, 0x83, 0x93,
	0xc0, 0x35, 0x58, 0x20, 0xdb, 0x93, 0x00, 0x0b, 0x24, 0x24, 0xd8, 0xf9, 0xff, 0x7f, 0xfb, 0xf3,
	0xff, 0x1e, 0x8c, 0xb0, 0x46, 0xe3, 0x9d, 0x76, 0x55, 0xd3, 0x5a, 0xb2, 0x67, 0x70, 0x6f, 0x17,
	0x87, 0xf3, 0xc8, 0xa9, 0x0d, 0xd6, 0xa4, 0x97, 0x49, 0x8f, 0xdf, 0x40, 0x31, 0xc7, 0x4d, 0x63,
	0x14, 0x3f, 0x85, 0xbe, 0xc1, 0x85, 0x32, 0x25, 0x13, 0x6c, 0x72, 0x22, 0x93, 0xe0, 0x8f, 0x61,
	0xb0, 0xff, 0xb0, 0x43, 0xb3, 0x55, 0x65, 0x26, 0xd8, 0x24, 0x97, 0xc5, 0xfe, 0x5d, 0x50, 0x21,
	0xf0, 0x5d, 0xd0, 0x4b, 0x81, 0x8f, 0xc1, 0xf8, 0x06, 0x86, 0x89, 0x38, 0x33, 0x16, 0xe9, 0xef,
	0xb0, 0xd9, 0x9f, 0xb0, 0xd9, 0x11, 0xfb, 0x89, 0x41, 0xff, 0xba, 0xc5, 0x66, 0x1d, 0x88, 0xa4,
	0xc9, 0xa8, 0x03, 0x31, 0x0a, 0xfe, 0x08, 0x8a, 0x7d, 0xb2, 0xb3, 0x68, 0x77, 0x2a, 0xf8, 0x3e,
	0xf9, 0xbd, 0xe4, 0x27, 0xc5, 0x9f, 0x40, 0xbe, 0x42, 0xc2, 0x32, 0x17, 0xbd, 0xc9, 0xf0, 0x62,
	0x50, 0xa5, 0xce, 0x32, 0x9a, 0xfc, 0x2c, 0xd4, 0x73, 0x4b, 0x34, 0xaa, 0xec, 0x0b, 0x36, 0x79,
	0xf8, 0x82, 0x9d, 0xcb, 0x62, 0x3f, 0x0f, 0x46, 0xc8, 0x7c, 0x97, 0x15, 0xc7, 0xcc, 0xc7, 0x6c,
	0xfc, 0x99, 0x01, 0xc4, 0x92, 0xc7, 0xd9, 0xff, 0x43, 0x53, 0xf1, 0x5b, 0xd3, 0x07, 0xd5, 0x2f,
	0xdb, 0xfd, 0xc7, 0xba, 0x1f, 0x33, 0xe8, 0x5f, 0xed, 0x54, 0x4d, 0xfc, 0x1a, 0x4e, 0x48, 0x6f,
	0x94, 0x23, 0xdc, 0x34, 0xb1, 0x6d, 0x7e, 0xf9, 0xfc, 0xeb, 0xf7, 0x6f, 0x5f, 0xd8, 0x33, 0x18,
	0xca, 0xe9, 0x6c, 0x8a, 0xa4, 0x42, 0xce, 0x4f, 0xdf, 0xae, 0x95, 0x50, 0xe1, 0x89, 0x38, 0xde,
	0xaf, 0xe4, 0xcf, 0xb7, 0xfc, 0x3d, 0x14, 0xce, 0x6e, 0xdb, 0x65, 0x37, 0xdc, 0xe5, 0xab, 0x48,
	0x79, 0x09, 0x85, 0x9c, 0xce, 0x6e, 0xe4, 0x6b, 0x7e, 0x1e, 0x00, 0xdb, 0xb6, 0x16, 0xf6, 0x56,
	0xd0, 0x5a, 0x09, 0xdb, 0xea, 0x3b, 0x5d, 0x23, 0xe9, 0xfa, 0x4e, 0xd8, 0xc5, 0xbd, 0x5a, 0x92,
	0xb8, 0xb5, 0xad, 0xa0, 0xb5, 0x76, 0xe9, 0x9f, 0x4a, 0x76, 0x40, 0x5e, 0xc2, 0xc0, 0x6d, 0xe3,
	0x95, 0x6e, 0x41, 0x07, 0xc9, 0x39, 0xe4, 0xe4, 0x1b, 0x55, 0xe6, 0xd1, 0x8e, 0x67, 0xfe, 0x14,
	0x72, 0x47, 0x48, 0x71, 0x21, 0xc3, 0x0b, 0xa8, 0xe6, 0x84, 0x74, 0x55, 0x53, 0xeb, 0x65, 0xf4,
	0xf9, 0x08, 0x32, 0xbd, 0x4a, 0x2b, 0x91, 0x99, 0x5e, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0x20, 0x34, 0xfd, 0x35, 0x03, 0x00, 0x00,
}
