// Code generated by protoc-gen-go.
// source: output_plugin.proto
// DO NOT EDIT!

package apiclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OutputPluginBatchProcessingStatus_Status int32

const (
	OutputPluginBatchProcessingStatus_SUCCESS OutputPluginBatchProcessingStatus_Status = 0
	OutputPluginBatchProcessingStatus_ERROR   OutputPluginBatchProcessingStatus_Status = 1
)

var OutputPluginBatchProcessingStatus_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR",
}
var OutputPluginBatchProcessingStatus_Status_value = map[string]int32{
	"SUCCESS": 0,
	"ERROR":   1,
}

func (x OutputPluginBatchProcessingStatus_Status) Enum() *OutputPluginBatchProcessingStatus_Status {
	p := new(OutputPluginBatchProcessingStatus_Status)
	*p = x
	return p
}
func (x OutputPluginBatchProcessingStatus_Status) String() string {
	return proto.EnumName(OutputPluginBatchProcessingStatus_Status_name, int32(x))
}
func (x *OutputPluginBatchProcessingStatus_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OutputPluginBatchProcessingStatus_Status_value, data, "OutputPluginBatchProcessingStatus_Status")
	if err != nil {
		return err
	}
	*x = OutputPluginBatchProcessingStatus_Status(value)
	return nil
}
func (OutputPluginBatchProcessingStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor13, []int{1, 0}
}

type OutputPluginVerificationResult_Status int32

const (
	OutputPluginVerificationResult_N_A     OutputPluginVerificationResult_Status = 0
	OutputPluginVerificationResult_SUCCESS OutputPluginVerificationResult_Status = 1
	OutputPluginVerificationResult_WARNING OutputPluginVerificationResult_Status = 2
	OutputPluginVerificationResult_FAILURE OutputPluginVerificationResult_Status = 3
)

var OutputPluginVerificationResult_Status_name = map[int32]string{
	0: "N_A",
	1: "SUCCESS",
	2: "WARNING",
	3: "FAILURE",
}
var OutputPluginVerificationResult_Status_value = map[string]int32{
	"N_A":     0,
	"SUCCESS": 1,
	"WARNING": 2,
	"FAILURE": 3,
}

func (x OutputPluginVerificationResult_Status) Enum() *OutputPluginVerificationResult_Status {
	p := new(OutputPluginVerificationResult_Status)
	*p = x
	return p
}
func (x OutputPluginVerificationResult_Status) String() string {
	return proto.EnumName(OutputPluginVerificationResult_Status_name, int32(x))
}
func (x *OutputPluginVerificationResult_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OutputPluginVerificationResult_Status_value, data, "OutputPluginVerificationResult_Status")
	if err != nil {
		return err
	}
	*x = OutputPluginVerificationResult_Status(value)
	return nil
}
func (OutputPluginVerificationResult_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor13, []int{2, 0}
}

type OutputPluginDescriptor struct {
	PluginName       *string `protobuf:"bytes,1,opt,name=plugin_name,json=pluginName" json:"plugin_name,omitempty"`
	PluginArgs       []byte  `protobuf:"bytes,2,opt,name=plugin_args,json=pluginArgs" json:"plugin_args,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OutputPluginDescriptor) Reset()                    { *m = OutputPluginDescriptor{} }
func (m *OutputPluginDescriptor) String() string            { return proto.CompactTextString(m) }
func (*OutputPluginDescriptor) ProtoMessage()               {}
func (*OutputPluginDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *OutputPluginDescriptor) GetPluginName() string {
	if m != nil && m.PluginName != nil {
		return *m.PluginName
	}
	return ""
}

func (m *OutputPluginDescriptor) GetPluginArgs() []byte {
	if m != nil {
		return m.PluginArgs
	}
	return nil
}

// Next id: 7
type OutputPluginBatchProcessingStatus struct {
	Status           *OutputPluginBatchProcessingStatus_Status `protobuf:"varint,1,opt,name=status,enum=OutputPluginBatchProcessingStatus_Status,def=0" json:"status,omitempty"`
	PluginDescriptor *OutputPluginDescriptor                   `protobuf:"bytes,6,opt,name=plugin_descriptor,json=pluginDescriptor" json:"plugin_descriptor,omitempty"`
	Summary          *string                                   `protobuf:"bytes,3,opt,name=summary" json:"summary,omitempty"`
	BatchIndex       *uint64                                   `protobuf:"varint,4,opt,name=batch_index,json=batchIndex" json:"batch_index,omitempty"`
	BatchSize        *uint64                                   `protobuf:"varint,5,opt,name=batch_size,json=batchSize" json:"batch_size,omitempty"`
	XXX_unrecognized []byte                                    `json:"-"`
}

func (m *OutputPluginBatchProcessingStatus) Reset()         { *m = OutputPluginBatchProcessingStatus{} }
func (m *OutputPluginBatchProcessingStatus) String() string { return proto.CompactTextString(m) }
func (*OutputPluginBatchProcessingStatus) ProtoMessage()    {}
func (*OutputPluginBatchProcessingStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{1}
}

const Default_OutputPluginBatchProcessingStatus_Status OutputPluginBatchProcessingStatus_Status = OutputPluginBatchProcessingStatus_SUCCESS

func (m *OutputPluginBatchProcessingStatus) GetStatus() OutputPluginBatchProcessingStatus_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_OutputPluginBatchProcessingStatus_Status
}

func (m *OutputPluginBatchProcessingStatus) GetPluginDescriptor() *OutputPluginDescriptor {
	if m != nil {
		return m.PluginDescriptor
	}
	return nil
}

func (m *OutputPluginBatchProcessingStatus) GetSummary() string {
	if m != nil && m.Summary != nil {
		return *m.Summary
	}
	return ""
}

func (m *OutputPluginBatchProcessingStatus) GetBatchIndex() uint64 {
	if m != nil && m.BatchIndex != nil {
		return *m.BatchIndex
	}
	return 0
}

func (m *OutputPluginBatchProcessingStatus) GetBatchSize() uint64 {
	if m != nil && m.BatchSize != nil {
		return *m.BatchSize
	}
	return 0
}

type OutputPluginVerificationResult struct {
	PluginId         *string                                `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId" json:"plugin_id,omitempty"`
	PluginDescriptor *OutputPluginDescriptor                `protobuf:"bytes,2,opt,name=plugin_descriptor,json=pluginDescriptor" json:"plugin_descriptor,omitempty"`
	Status           *OutputPluginVerificationResult_Status `protobuf:"varint,3,opt,name=status,enum=OutputPluginVerificationResult_Status" json:"status,omitempty"`
	StatusMessage    *string                                `protobuf:"bytes,4,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
	Timestamp        *uint64                                `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *OutputPluginVerificationResult) Reset()                    { *m = OutputPluginVerificationResult{} }
func (m *OutputPluginVerificationResult) String() string            { return proto.CompactTextString(m) }
func (*OutputPluginVerificationResult) ProtoMessage()               {}
func (*OutputPluginVerificationResult) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *OutputPluginVerificationResult) GetPluginId() string {
	if m != nil && m.PluginId != nil {
		return *m.PluginId
	}
	return ""
}

func (m *OutputPluginVerificationResult) GetPluginDescriptor() *OutputPluginDescriptor {
	if m != nil {
		return m.PluginDescriptor
	}
	return nil
}

func (m *OutputPluginVerificationResult) GetStatus() OutputPluginVerificationResult_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return OutputPluginVerificationResult_N_A
}

func (m *OutputPluginVerificationResult) GetStatusMessage() string {
	if m != nil && m.StatusMessage != nil {
		return *m.StatusMessage
	}
	return ""
}

func (m *OutputPluginVerificationResult) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type OutputPluginVerificationResultsList struct {
	Results          []*OutputPluginVerificationResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *OutputPluginVerificationResultsList) Reset()         { *m = OutputPluginVerificationResultsList{} }
func (m *OutputPluginVerificationResultsList) String() string { return proto.CompactTextString(m) }
func (*OutputPluginVerificationResultsList) ProtoMessage()    {}
func (*OutputPluginVerificationResultsList) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{3}
}

func (m *OutputPluginVerificationResultsList) GetResults() []*OutputPluginVerificationResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type EmailOutputPluginArgs struct {
	EmailAddress     *string `protobuf:"bytes,1,opt,name=email_address,json=emailAddress" json:"email_address,omitempty"`
	EmailsLimit      *uint64 `protobuf:"varint,2,opt,name=emails_limit,json=emailsLimit,def=100" json:"emails_limit,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EmailOutputPluginArgs) Reset()                    { *m = EmailOutputPluginArgs{} }
func (m *EmailOutputPluginArgs) String() string            { return proto.CompactTextString(m) }
func (*EmailOutputPluginArgs) ProtoMessage()               {}
func (*EmailOutputPluginArgs) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

const Default_EmailOutputPluginArgs_EmailsLimit uint64 = 100

func (m *EmailOutputPluginArgs) GetEmailAddress() string {
	if m != nil && m.EmailAddress != nil {
		return *m.EmailAddress
	}
	return ""
}

func (m *EmailOutputPluginArgs) GetEmailsLimit() uint64 {
	if m != nil && m.EmailsLimit != nil {
		return *m.EmailsLimit
	}
	return Default_EmailOutputPluginArgs_EmailsLimit
}

type BigQueryOutputPluginArgs struct {
	ExportOptions    *ExportOptions `protobuf:"bytes,2,opt,name=export_options,json=exportOptions" json:"export_options,omitempty"`
	ConvertValues    *bool          `protobuf:"varint,3,opt,name=convert_values,json=convertValues,def=1" json:"convert_values,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *BigQueryOutputPluginArgs) Reset()                    { *m = BigQueryOutputPluginArgs{} }
func (m *BigQueryOutputPluginArgs) String() string            { return proto.CompactTextString(m) }
func (*BigQueryOutputPluginArgs) ProtoMessage()               {}
func (*BigQueryOutputPluginArgs) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

const Default_BigQueryOutputPluginArgs_ConvertValues bool = true

func (m *BigQueryOutputPluginArgs) GetExportOptions() *ExportOptions {
	if m != nil {
		return m.ExportOptions
	}
	return nil
}

func (m *BigQueryOutputPluginArgs) GetConvertValues() bool {
	if m != nil && m.ConvertValues != nil {
		return *m.ConvertValues
	}
	return Default_BigQueryOutputPluginArgs_ConvertValues
}

type CSVOutputPluginArgs struct {
	ExportOptions    *ExportOptions `protobuf:"bytes,2,opt,name=export_options,json=exportOptions" json:"export_options,omitempty"`
	ConvertValues    *bool          `protobuf:"varint,3,opt,name=convert_values,json=convertValues,def=1" json:"convert_values,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *CSVOutputPluginArgs) Reset()                    { *m = CSVOutputPluginArgs{} }
func (m *CSVOutputPluginArgs) String() string            { return proto.CompactTextString(m) }
func (*CSVOutputPluginArgs) ProtoMessage()               {}
func (*CSVOutputPluginArgs) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

const Default_CSVOutputPluginArgs_ConvertValues bool = true

func (m *CSVOutputPluginArgs) GetExportOptions() *ExportOptions {
	if m != nil {
		return m.ExportOptions
	}
	return nil
}

func (m *CSVOutputPluginArgs) GetConvertValues() bool {
	if m != nil && m.ConvertValues != nil {
		return *m.ConvertValues
	}
	return Default_CSVOutputPluginArgs_ConvertValues
}

func init() {
	proto.RegisterType((*OutputPluginDescriptor)(nil), "OutputPluginDescriptor")
	proto.RegisterType((*OutputPluginBatchProcessingStatus)(nil), "OutputPluginBatchProcessingStatus")
	proto.RegisterType((*OutputPluginVerificationResult)(nil), "OutputPluginVerificationResult")
	proto.RegisterType((*OutputPluginVerificationResultsList)(nil), "OutputPluginVerificationResultsList")
	proto.RegisterType((*EmailOutputPluginArgs)(nil), "EmailOutputPluginArgs")
	proto.RegisterType((*BigQueryOutputPluginArgs)(nil), "BigQueryOutputPluginArgs")
	proto.RegisterType((*CSVOutputPluginArgs)(nil), "CSVOutputPluginArgs")
	proto.RegisterEnum("OutputPluginBatchProcessingStatus_Status", OutputPluginBatchProcessingStatus_Status_name, OutputPluginBatchProcessingStatus_Status_value)
	proto.RegisterEnum("OutputPluginVerificationResult_Status", OutputPluginVerificationResult_Status_name, OutputPluginVerificationResult_Status_value)
}

func init() { proto.RegisterFile("output_plugin.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x55, 0xcb, 0x72, 0x1b, 0x45,
	0x14, 0xcd, 0x58, 0x8a, 0x1d, 0xb5, 0x6c, 0x21, 0xda, 0x05, 0x19, 0x58, 0x84, 0x46, 0x21, 0x8e,
	0x80, 0x64, 0x30, 0x4e, 0x85, 0x87, 0x81, 0x50, 0x92, 0x25, 0xbb, 0x94, 0xb2, 0x65, 0x33, 0x72,
	0x1c, 0x76, 0x43, 0x6b, 0xe6, 0x4a, 0xea, 0x30, 0x2f, 0xba, 0x7b, 0x42, 0x94, 0x62, 0xc3, 0x96,
	0x1d, 0x1f, 0xc1, 0x0f, 0xf0, 0x0d, 0xfc, 0x01, 0x1b, 0xd6, 0x50, 0x7c, 0x02, 0x3b, 0x16, 0x54,
	0x3f, 0x24, 0xab, 0x6c, 0x97, 0xbd, 0x67, 0xa5, 0xee, 0xd3, 0xf7, 0x9c, 0xdb, 0xf7, 0x9e, 0x3b,
	0x2d, 0xb4, 0x9e, 0x15, 0x32, 0x2f, 0x64, 0x90, 0xc7, 0xc5, 0x98, 0xa5, 0x5e, 0xce, 0x33, 0x99,
	0xbd, 0xb9, 0x0a, 0x2f, 0xf2, 0x8c, 0x4b, 0xbb, 0xab, 0x09, 0x48, 0x68, 0x2a, 0x59, 0x68, 0xf6,
	0x8d, 0x7f, 0x1c, 0xf4, 0xfa, 0xa1, 0x66, 0x1d, 0x69, 0x52, 0x07, 0x44, 0xc8, 0x59, 0x2e, 0x33,
	0x8e, 0xf7, 0x50, 0xd5, 0x08, 0x05, 0x29, 0x4d, 0xc0, 0x75, 0x88, 0xd3, 0xac, 0xb4, 0x37, 0xfe,
	0xfc, 0xf7, 0xaf, 0xdf, 0x1c, 0x82, 0x6f, 0x1d, 0x4f, 0x80, 0x28, 0x9c, 0x64, 0x23, 0x22, 0x27,
	0x40, 0x4c, 0x6a, 0x62, 0x53, 0xfb, 0xc8, 0x2c, 0xfa, 0x34, 0x01, 0xfc, 0xb3, 0x33, 0x57, 0xa2,
	0x7c, 0x2c, 0xdc, 0x25, 0xe2, 0x34, 0x57, 0xdb, 0xb9, 0x56, 0x7a, 0x86, 0x8f, 0x95, 0x52, 0x4e,
	0x39, 0x4d, 0x40, 0x02, 0x17, 0x64, 0x94, 0x71, 0x22, 0x27, 0x4c, 0xcc, 0xa4, 0xc8, 0x41, 0x21,
	0x24, 0x19, 0x02, 0xa1, 0x29, 0x61, 0xa9, 0x90, 0x34, 0x0d, 0xe7, 0x49, 0xd5, 0x05, 0x22, 0x1b,
	0x78, 0x57, 0x10, 0xa5, 0x1e, 0xc8, 0x69, 0x0e, 0xde, 0x7b, 0x78, 0x0f, 0x6c, 0x45, 0x2d, 0x3e,
	0x16, 0x3b, 0x31, 0x15, 0x62, 0x76, 0x27, 0x05, 0x34, 0x7e, 0x29, 0xa3, 0xb7, 0x17, 0xeb, 0x6e,
	0x53, 0x19, 0x4e, 0x8e, 0x78, 0x16, 0x82, 0x10, 0x2c, 0x1d, 0x0f, 0x24, 0x95, 0x85, 0xc0, 0x3f,
	0x3a, 0x68, 0x59, 0xe8, 0xa5, 0x2e, 0xbf, 0xb6, 0xf5, 0xae, 0x77, 0x25, 0xc9, 0x33, 0x3f, 0xdb,
	0x2b, 0x83, 0x27, 0x3b, 0x3b, 0xdd, 0xc1, 0xa0, 0xfd, 0x50, 0x17, 0xfa, 0x01, 0xbe, 0xff, 0x94,
	0x0a, 0x53, 0xd9, 0x50, 0xf1, 0x88, 0x28, 0x42, 0x45, 0x1c, 0x15, 0x71, 0x3c, 0x25, 0xb9, 0x51,
	0x81, 0x88, 0x64, 0x9c, 0xa4, 0x99, 0xfc, 0xd2, 0xb7, 0x89, 0xf1, 0x4f, 0x0e, 0x7a, 0xd5, 0x76,
	0x2f, 0x9a, 0x9b, 0xe3, 0x2e, 0x13, 0xa7, 0x59, 0xdd, 0xba, 0xe9, 0x5d, 0xec, 0x5d, 0xbb, 0xa3,
	0x73, 0x3e, 0xc2, 0x9f, 0x9f, 0x62, 0x17, 0x1a, 0x45, 0x38, 0x88, 0x3c, 0x4b, 0x05, 0x1b, 0xc6,
	0x70, 0xda, 0x79, 0x0e, 0xa2, 0x88, 0xa5, 0xe7, 0xd7, 0xf3, 0xb3, 0x33, 0xb1, 0x85, 0x56, 0x44,
	0x91, 0x24, 0x94, 0x4f, 0xdd, 0x92, 0x9e, 0x07, 0x57, 0x27, 0xc2, 0xb8, 0x3e, 0x30, 0x30, 0x49,
	0x40, 0x08, 0x3a, 0x06, 0xcf, 0x9f, 0x05, 0xe2, 0x03, 0x54, 0xd5, 0x05, 0x07, 0x2c, 0x8d, 0xe0,
	0x85, 0x5b, 0x26, 0x4e, 0xb3, 0xdc, 0xbe, 0xa7, 0x79, 0x1b, 0xf8, 0x9d, 0x7e, 0x91, 0x0c, 0x61,
	0x7e, 0x39, 0xd3, 0x99, 0x21, 0xb0, 0x74, 0x7c, 0xda, 0x12, 0xcf, 0x47, 0x1a, 0xef, 0x29, 0x3e,
	0x7e, 0x8c, 0xcc, 0x2e, 0x10, 0xec, 0x25, 0xb8, 0xd7, 0xb5, 0xda, 0xfb, 0x5a, 0xed, 0x0e, 0xbe,
	0x3d, 0x60, 0x2f, 0xe1, 0xac, 0x96, 0x3c, 0x23, 0x56, 0xd1, 0x07, 0x2a, 0xb2, 0x41, 0xd0, 0xb2,
	0x75, 0xba, 0x8a, 0x66, 0x7e, 0xd5, 0xaf, 0xe1, 0x0a, 0xba, 0xde, 0xf5, 0xfd, 0x43, 0xbf, 0xee,
	0x34, 0xfe, 0x2e, 0xa1, 0x5b, 0x8b, 0x3d, 0x3e, 0x01, 0xce, 0x46, 0x2c, 0xa4, 0x92, 0x65, 0xa9,
	0xaf, 0xdb, 0x84, 0xbf, 0x46, 0x15, 0xeb, 0x0f, 0x8b, 0xec, 0x57, 0xf2, 0x99, 0xbe, 0xcf, 0x43,
	0xfc, 0x60, 0x20, 0xb9, 0xca, 0xce, 0x22, 0x48, 0x25, 0x1b, 0x4d, 0xd5, 0xfa, 0xbc, 0x07, 0xa6,
	0xdc, 0x70, 0x02, 0xe1, 0xb7, 0xea, 0x7e, 0x37, 0x0c, 0xdc, 0x8b, 0xf0, 0x0f, 0x17, 0x39, 0xbf,
	0x74, 0xb9, 0xf3, 0x1f, 0xe9, 0xd4, 0x9b, 0xd8, 0x3b, 0xef, 0xbc, 0x4d, 0x27, 0x27, 0x54, 0xde,
	0x15, 0x67, 0xb3, 0x9e, 0xf7, 0xfa, 0xd1, 0x7c, 0xf6, 0x4b, 0x7a, 0xf6, 0x37, 0xbc, 0xcb, 0x1b,
	0x61, 0x07, 0x7f, 0x3e, 0xb8, 0x77, 0x50, 0xcd, 0xac, 0x02, 0x3b, 0x13, 0xda, 0xfa, 0x8a, 0xbf,
	0x66, 0xd0, 0x03, 0x03, 0xe2, 0xc7, 0xa8, 0x22, 0x59, 0x02, 0x42, 0xd2, 0x24, 0xb7, 0x76, 0xda,
	0xe1, 0x40, 0x55, 0xbf, 0xb3, 0xdb, 0xa1, 0x12, 0xd4, 0x39, 0xbe, 0x79, 0x3c, 0x8b, 0x9a, 0xd5,
	0xa3, 0xef, 0xee, 0xf9, 0xa7, 0xf4, 0xc6, 0x27, 0x73, 0x3f, 0x57, 0x50, 0xa9, 0x1f, 0xb4, 0xea,
	0xd7, 0x16, 0x8d, 0x75, 0xd4, 0xe6, 0x69, 0xcb, 0xef, 0xf7, 0xfa, 0x7b, 0xf5, 0x25, 0xb5, 0xd9,
	0x6d, 0xf5, 0xf6, 0x9f, 0xf8, 0xdd, 0x7a, 0xa9, 0xf1, 0x0d, 0xba, 0x7d, 0x79, 0x75, 0x62, 0x9f,
	0x09, 0x89, 0x3f, 0x45, 0x2b, 0xe6, 0xe3, 0x50, 0x0f, 0x42, 0xa9, 0x59, 0xdd, 0x7a, 0xeb, 0x8a,
	0xa6, 0xf8, 0xb3, 0xf8, 0xc6, 0xaf, 0x0e, 0x7a, 0xad, 0x9b, 0x50, 0x16, 0x2f, 0x12, 0xd4, 0x5b,
	0x84, 0xbf, 0x43, 0x6b, 0xa0, 0x0e, 0x02, 0x1a, 0x45, 0x1c, 0x84, 0xb0, 0x43, 0xb4, 0xaf, 0xbb,
	0xb0, 0x8b, 0x70, 0x27, 0x4b, 0x28, 0x4b, 0x35, 0xb7, 0x65, 0x22, 0xf0, 0xa6, 0x7a, 0x34, 0x35,
	0x89, 0x58, 0x92, 0xb6, 0x75, 0xf6, 0xf5, 0x09, 0xf2, 0x3d, 0x8b, 0x63, 0xf5, 0x68, 0x0a, 0x48,
	0x25, 0x91, 0x99, 0xe7, 0xaf, 0xc2, 0x22, 0x7f, 0x03, 0x99, 0xbd, 0x08, 0x62, 0x96, 0x30, 0xa9,
	0x87, 0xaa, 0xbc, 0x5d, 0xfa, 0x70, 0x73, 0xd3, 0xaf, 0x9a, 0x83, 0x7d, 0x85, 0x37, 0xfe, 0x70,
	0x90, 0xdb, 0x66, 0xe3, 0xaf, 0x0a, 0xe0, 0xd3, 0x73, 0xf7, 0x3e, 0x42, 0x35, 0xf3, 0xdf, 0x12,
	0x64, 0xb9, 0xaa, 0x58, 0xd8, 0xd9, 0xac, 0x79, 0x5d, 0x0d, 0x1f, 0x1a, 0xb4, 0xfd, 0x86, 0x2e,
	0x64, 0x1d, 0xbf, 0x62, 0x60, 0x62, 0xa3, 0x3d, 0xd7, 0xf1, 0xd7, 0x60, 0x31, 0x12, 0x3f, 0x43,
	0xb5, 0x30, 0x4b, 0x9f, 0x03, 0x97, 0xc1, 0x73, 0x1a, 0x17, 0x60, 0x46, 0xef, 0xc6, 0x76, 0x59,
	0xf2, 0x02, 0xda, 0x5f, 0x68, 0x9d, 0x8f, 0xf1, 0x83, 0xde, 0x88, 0x28, 0xe0, 0x1e, 0xb1, 0xc1,
	0xc4, 0x04, 0xeb, 0x07, 0xcc, 0x08, 0xde, 0x1f, 0x71, 0x06, 0x69, 0x14, 0x4f, 0x15, 0x96, 0x50,
	0xe9, 0xb9, 0x4b, 0xfe, 0x9a, 0x8d, 0x3e, 0xd1, 0xc1, 0x8d, 0xdf, 0x1d, 0xb4, 0xbe, 0x33, 0x38,
	0xf9, 0x7f, 0x55, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x27, 0x06, 0x33, 0x03, 0x08,
	0x00, 0x00,
}
