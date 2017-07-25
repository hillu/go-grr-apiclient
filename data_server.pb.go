// Code generated by protoc-gen-go.
// source: data_server.proto
// DO NOT EDIT!

package apiclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DataStoreCommand_Command int32

const (
	DataStoreCommand_MULTI_SET DataStoreCommand_Command = 0
	// Deprecated
	// MULTI_RESOLVE_REGEX = 1;
	DataStoreCommand_RESOLVE_MULTI     DataStoreCommand_Command = 2
	DataStoreCommand_DELETE_SUBJECT    DataStoreCommand_Command = 3
	DataStoreCommand_DELETE_ATTRIBUTES DataStoreCommand_Command = 4
	// Deprecated.
	// DELETE_ATTRIBUTES_REGEX = 5;
	DataStoreCommand_LOCK_SUBJECT         DataStoreCommand_Command = 6
	DataStoreCommand_UNLOCK_SUBJECT       DataStoreCommand_Command = 7
	DataStoreCommand_EXTEND_SUBJECT       DataStoreCommand_Command = 8
	DataStoreCommand_MULTI_RESOLVE_PREFIX DataStoreCommand_Command = 9
	DataStoreCommand_SCAN_ATTRIBUTES      DataStoreCommand_Command = 10
)

var DataStoreCommand_Command_name = map[int32]string{
	0:  "MULTI_SET",
	2:  "RESOLVE_MULTI",
	3:  "DELETE_SUBJECT",
	4:  "DELETE_ATTRIBUTES",
	6:  "LOCK_SUBJECT",
	7:  "UNLOCK_SUBJECT",
	8:  "EXTEND_SUBJECT",
	9:  "MULTI_RESOLVE_PREFIX",
	10: "SCAN_ATTRIBUTES",
}
var DataStoreCommand_Command_value = map[string]int32{
	"MULTI_SET":            0,
	"RESOLVE_MULTI":        2,
	"DELETE_SUBJECT":       3,
	"DELETE_ATTRIBUTES":    4,
	"LOCK_SUBJECT":         6,
	"UNLOCK_SUBJECT":       7,
	"EXTEND_SUBJECT":       8,
	"MULTI_RESOLVE_PREFIX": 9,
	"SCAN_ATTRIBUTES":      10,
}

func (x DataStoreCommand_Command) Enum() *DataStoreCommand_Command {
	p := new(DataStoreCommand_Command)
	*p = x
	return p
}
func (x DataStoreCommand_Command) String() string {
	return proto.EnumName(DataStoreCommand_Command_name, int32(x))
}
func (x *DataStoreCommand_Command) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataStoreCommand_Command_value, data, "DataStoreCommand_Command")
	if err != nil {
		return err
	}
	*x = DataStoreCommand_Command(value)
	return nil
}
func (DataStoreCommand_Command) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

type DataServerState_Status int32

const (
	DataServerState_AVAILABLE DataServerState_Status = 0
	DataServerState_OFFLINE   DataServerState_Status = 1
)

var DataServerState_Status_name = map[int32]string{
	0: "AVAILABLE",
	1: "OFFLINE",
}
var DataServerState_Status_value = map[string]int32{
	"AVAILABLE": 0,
	"OFFLINE":   1,
}

func (x DataServerState_Status) Enum() *DataServerState_Status {
	p := new(DataServerState_Status)
	*p = x
	return p
}
func (x DataServerState_Status) String() string {
	return proto.EnumName(DataServerState_Status_name, int32(x))
}
func (x *DataServerState_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataServerState_Status_value, data, "DataServerState_Status")
	if err != nil {
		return err
	}
	*x = DataServerState_Status(value)
	return nil
}
func (DataServerState_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{2, 0} }

type DataStoreCommand struct {
	Command          *DataStoreCommand_Command `protobuf:"varint,1,opt,name=command,enum=DataStoreCommand_Command" json:"command,omitempty"`
	Request          *DataStoreRequest         `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *DataStoreCommand) Reset()                    { *m = DataStoreCommand{} }
func (m *DataStoreCommand) String() string            { return proto.CompactTextString(m) }
func (*DataStoreCommand) ProtoMessage()               {}
func (*DataStoreCommand) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *DataStoreCommand) GetCommand() DataStoreCommand_Command {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return DataStoreCommand_MULTI_SET
}

func (m *DataStoreCommand) GetRequest() *DataStoreRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type DataServerInterval struct {
	// Range of hashes used by the server.
	// Represented as an interval [start, end[.
	Start            *uint64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End              *uint64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DataServerInterval) Reset()                    { *m = DataServerInterval{} }
func (m *DataServerInterval) String() string            { return proto.CompactTextString(m) }
func (*DataServerInterval) ProtoMessage()               {}
func (*DataServerInterval) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *DataServerInterval) GetStart() uint64 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *DataServerInterval) GetEnd() uint64 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

type DataServerState struct {
	Status           *DataServerState_Status `protobuf:"varint,1,opt,name=status,enum=DataServerState_Status" json:"status,omitempty"`
	Load             *uint64                 `protobuf:"varint,2,opt,name=load" json:"load,omitempty"`
	Size             *uint64                 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	NumComponents    *uint64                 `protobuf:"varint,4,opt,name=num_components,json=numComponents" json:"num_components,omitempty"`
	AvgComponent     *uint64                 `protobuf:"varint,5,opt,name=avg_component,json=avgComponent" json:"avg_component,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *DataServerState) Reset()                    { *m = DataServerState{} }
func (m *DataServerState) String() string            { return proto.CompactTextString(m) }
func (*DataServerState) ProtoMessage()               {}
func (*DataServerState) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *DataServerState) GetStatus() DataServerState_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return DataServerState_AVAILABLE
}

func (m *DataServerState) GetLoad() uint64 {
	if m != nil && m.Load != nil {
		return *m.Load
	}
	return 0
}

func (m *DataServerState) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *DataServerState) GetNumComponents() uint64 {
	if m != nil && m.NumComponents != nil {
		return *m.NumComponents
	}
	return 0
}

func (m *DataServerState) GetAvgComponent() uint64 {
	if m != nil && m.AvgComponent != nil {
		return *m.AvgComponent
	}
	return 0
}

type DataServerInformation struct {
	Index            *uint64             `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Address          *string             `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Port             *uint64             `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	State            *DataServerState    `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Interval         *DataServerInterval `protobuf:"bytes,5,opt,name=interval" json:"interval,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *DataServerInformation) Reset()                    { *m = DataServerInformation{} }
func (m *DataServerInformation) String() string            { return proto.CompactTextString(m) }
func (*DataServerInformation) ProtoMessage()               {}
func (*DataServerInformation) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DataServerInformation) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *DataServerInformation) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *DataServerInformation) GetPort() uint64 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *DataServerInformation) GetState() *DataServerState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *DataServerInformation) GetInterval() *DataServerInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

type DataServerMapping struct {
	// Version of the mapping.
	Version *uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Number of data servers.
	NumServers *uint64 `protobuf:"varint,2,opt,name=num_servers,json=numServers" json:"num_servers,omitempty"`
	// Information about each server.
	Servers []*DataServerInformation `protobuf:"bytes,3,rep,name=servers" json:"servers,omitempty"`
	// Pathing information for subject paths.
	Pathing          []string `protobuf:"bytes,4,rep,name=pathing" json:"pathing,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DataServerMapping) Reset()                    { *m = DataServerMapping{} }
func (m *DataServerMapping) String() string            { return proto.CompactTextString(m) }
func (*DataServerMapping) ProtoMessage()               {}
func (*DataServerMapping) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *DataServerMapping) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *DataServerMapping) GetNumServers() uint64 {
	if m != nil && m.NumServers != nil {
		return *m.NumServers
	}
	return 0
}

func (m *DataServerMapping) GetServers() []*DataServerInformation {
	if m != nil {
		return m.Servers
	}
	return nil
}

func (m *DataServerMapping) GetPathing() []string {
	if m != nil {
		return m.Pathing
	}
	return nil
}

type DataServerClientInformation struct {
	// Client username.
	Username *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	// Client password.
	Password *string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	// Client permissions (r, rw, w).
	Permissions      *string `protobuf:"bytes,3,opt,name=permissions" json:"permissions,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DataServerClientInformation) Reset()                    { *m = DataServerClientInformation{} }
func (m *DataServerClientInformation) String() string            { return proto.CompactTextString(m) }
func (*DataServerClientInformation) ProtoMessage()               {}
func (*DataServerClientInformation) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *DataServerClientInformation) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *DataServerClientInformation) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *DataServerClientInformation) GetPermissions() string {
	if m != nil && m.Permissions != nil {
		return *m.Permissions
	}
	return ""
}

type DataServerEncryptedCreds struct {
	InitVector       []byte `protobuf:"bytes,1,opt,name=init_vector,json=initVector" json:"init_vector,omitempty"`
	Ciphertext       []byte `protobuf:"bytes,2,opt,name=ciphertext" json:"ciphertext,omitempty"`
	Sha256           []byte `protobuf:"bytes,3,opt,name=sha256" json:"sha256,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DataServerEncryptedCreds) Reset()                    { *m = DataServerEncryptedCreds{} }
func (m *DataServerEncryptedCreds) String() string            { return proto.CompactTextString(m) }
func (*DataServerEncryptedCreds) ProtoMessage()               {}
func (*DataServerEncryptedCreds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *DataServerEncryptedCreds) GetInitVector() []byte {
	if m != nil {
		return m.InitVector
	}
	return nil
}

func (m *DataServerEncryptedCreds) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

func (m *DataServerEncryptedCreds) GetSha256() []byte {
	if m != nil {
		return m.Sha256
	}
	return nil
}

type DataServerClientCredentials struct {
	Users            []*DataServerClientInformation `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *DataServerClientCredentials) Reset()                    { *m = DataServerClientCredentials{} }
func (m *DataServerClientCredentials) String() string            { return proto.CompactTextString(m) }
func (*DataServerClientCredentials) ProtoMessage()               {}
func (*DataServerClientCredentials) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *DataServerClientCredentials) GetUsers() []*DataServerClientInformation {
	if m != nil {
		return m.Users
	}
	return nil
}

type DataServerRebalance struct {
	// ID of operation.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// New mapping information.
	Mapping *DataServerMapping `protobuf:"bytes,2,opt,name=mapping" json:"mapping,omitempty"`
	// Number of files need to move.
	Moving           []uint64 `protobuf:"varint,3,rep,name=moving" json:"moving,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DataServerRebalance) Reset()                    { *m = DataServerRebalance{} }
func (m *DataServerRebalance) String() string            { return proto.CompactTextString(m) }
func (*DataServerRebalance) ProtoMessage()               {}
func (*DataServerRebalance) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *DataServerRebalance) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *DataServerRebalance) GetMapping() *DataServerMapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

func (m *DataServerRebalance) GetMoving() []uint64 {
	if m != nil {
		return m.Moving
	}
	return nil
}

type DataServerFileCopy struct {
	// Rebalance operation.
	RebalanceId *string `protobuf:"bytes,1,opt,name=rebalance_id,json=rebalanceId" json:"rebalance_id,omitempty"`
	// Directory where the file will be copied.
	Directory *string `protobuf:"bytes,2,opt,name=directory" json:"directory,omitempty"`
	// Filename for the file to copy.
	Filename *string `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	// Size of file.
	Size             *uint64 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DataServerFileCopy) Reset()                    { *m = DataServerFileCopy{} }
func (m *DataServerFileCopy) String() string            { return proto.CompactTextString(m) }
func (*DataServerFileCopy) ProtoMessage()               {}
func (*DataServerFileCopy) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *DataServerFileCopy) GetRebalanceId() string {
	if m != nil && m.RebalanceId != nil {
		return *m.RebalanceId
	}
	return ""
}

func (m *DataServerFileCopy) GetDirectory() string {
	if m != nil && m.Directory != nil {
		return *m.Directory
	}
	return ""
}

func (m *DataServerFileCopy) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *DataServerFileCopy) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

type DataStoreAuthToken struct {
	Username         *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Nonce            *string `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DataStoreAuthToken) Reset()                    { *m = DataStoreAuthToken{} }
func (m *DataStoreAuthToken) String() string            { return proto.CompactTextString(m) }
func (*DataStoreAuthToken) ProtoMessage()               {}
func (*DataStoreAuthToken) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

func (m *DataStoreAuthToken) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *DataStoreAuthToken) GetNonce() string {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return ""
}

func (m *DataStoreAuthToken) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

type DataStoreRegistrationRequest struct {
	Port             *uint32             `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	Token            *DataStoreAuthToken `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *DataStoreRegistrationRequest) Reset()                    { *m = DataStoreRegistrationRequest{} }
func (m *DataStoreRegistrationRequest) String() string            { return proto.CompactTextString(m) }
func (*DataStoreRegistrationRequest) ProtoMessage()               {}
func (*DataStoreRegistrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{11} }

func (m *DataStoreRegistrationRequest) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *DataStoreRegistrationRequest) GetToken() *DataStoreAuthToken {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*DataStoreCommand)(nil), "DataStoreCommand")
	proto.RegisterType((*DataServerInterval)(nil), "DataServerInterval")
	proto.RegisterType((*DataServerState)(nil), "DataServerState")
	proto.RegisterType((*DataServerInformation)(nil), "DataServerInformation")
	proto.RegisterType((*DataServerMapping)(nil), "DataServerMapping")
	proto.RegisterType((*DataServerClientInformation)(nil), "DataServerClientInformation")
	proto.RegisterType((*DataServerEncryptedCreds)(nil), "DataServerEncryptedCreds")
	proto.RegisterType((*DataServerClientCredentials)(nil), "DataServerClientCredentials")
	proto.RegisterType((*DataServerRebalance)(nil), "DataServerRebalance")
	proto.RegisterType((*DataServerFileCopy)(nil), "DataServerFileCopy")
	proto.RegisterType((*DataStoreAuthToken)(nil), "DataStoreAuthToken")
	proto.RegisterType((*DataStoreRegistrationRequest)(nil), "DataStoreRegistrationRequest")
	proto.RegisterEnum("DataStoreCommand_Command", DataStoreCommand_Command_name, DataStoreCommand_Command_value)
	proto.RegisterEnum("DataServerState_Status", DataServerState_Status_name, DataServerState_Status_value)
}

func init() { proto.RegisterFile("data_server.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 954 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xde, 0x89, 0x9d, 0x78, 0x5d, 0x76, 0xb2, 0xe3, 0x4e, 0x76, 0x31, 0x4b, 0x04, 0xc3, 0xf0,
	0xa3, 0x20, 0xd0, 0x2c, 0x18, 0x81, 0x38, 0x20, 0x21, 0xc7, 0x99, 0x20, 0xb3, 0xde, 0x04, 0xda,
	0x4e, 0xb4, 0x42, 0x42, 0x56, 0xe3, 0xe9, 0xd8, 0xad, 0xf5, 0xf4, 0x0c, 0xdd, 0x6d, 0xef, 0x86,
	0x3b, 0x2f, 0xc1, 0x73, 0x20, 0xae, 0x5c, 0xb8, 0xf3, 0x02, 0x9c, 0xe0, 0x35, 0x38, 0xa0, 0xfe,
	0x99, 0x19, 0x93, 0xa0, 0x3d, 0xb9, 0xeb, 0xab, 0xaf, 0xbb, 0xaa, 0xbe, 0x2a, 0xd7, 0x40, 0x27,
	0x21, 0x8a, 0x4c, 0x25, 0x15, 0x6b, 0x2a, 0xa2, 0x5c, 0x64, 0x2a, 0x7b, 0xe8, 0x5b, 0x48, 0x65,
	0x82, 0x3a, 0x64, 0x4f, 0xd2, 0x94, 0x70, 0xc5, 0x66, 0xd6, 0x0e, 0x7f, 0xdd, 0x02, 0xff, 0x84,
	0x28, 0x32, 0xd6, 0x9c, 0x41, 0x96, 0xa6, 0x84, 0x27, 0xe8, 0x63, 0x68, 0xcc, 0xec, 0xb1, 0xeb,
	0x05, 0xde, 0xd1, 0x5e, 0xef, 0xd5, 0xe8, 0x26, 0x27, 0x72, 0xbf, 0xb8, 0x60, 0xa2, 0xf7, 0xa1,
	0x21, 0xe8, 0x0f, 0x2b, 0x2a, 0x55, 0x77, 0x2b, 0xf0, 0x8e, 0x5a, 0xbd, 0x4e, 0x75, 0x09, 0x5b,
	0x07, 0x2e, 0x18, 0xe1, 0x6f, 0x1e, 0x34, 0x8a, 0x68, 0xbb, 0xd0, 0x7c, 0x72, 0x31, 0x9a, 0x0c,
	0xa7, 0xe3, 0x78, 0xe2, 0xdf, 0x41, 0x1d, 0xd8, 0xc5, 0xf1, 0xf8, 0x7c, 0x74, 0x19, 0x4f, 0x0d,
	0xec, 0x6f, 0x21, 0x04, 0x7b, 0x27, 0xf1, 0x28, 0x9e, 0xc4, 0xd3, 0xf1, 0xc5, 0xf1, 0x57, 0xf1,
	0x60, 0xe2, 0xd7, 0xd0, 0x7d, 0xe8, 0x38, 0xac, 0x3f, 0x99, 0xe0, 0xe1, 0xf1, 0xc5, 0x24, 0x1e,
	0xfb, 0x75, 0xe4, 0x43, 0x7b, 0x74, 0x3e, 0x78, 0x5c, 0x12, 0x77, 0xf4, 0xe5, 0x8b, 0xb3, 0xff,
	0x60, 0x0d, 0x8d, 0xc5, 0x4f, 0x27, 0xf1, 0xd9, 0x49, 0x89, 0xdd, 0x45, 0x5d, 0x38, 0xb0, 0x69,
	0x14, 0xd1, 0xbf, 0xc6, 0xf1, 0xe9, 0xf0, 0xa9, 0xdf, 0x44, 0xfb, 0x70, 0x6f, 0x3c, 0xe8, 0x9f,
	0x6d, 0x06, 0x82, 0xf0, 0x73, 0x40, 0xa6, 0x3c, 0x23, 0xf7, 0x90, 0x2b, 0x2a, 0xd6, 0x64, 0x89,
	0x0e, 0x60, 0x5b, 0x2a, 0x22, 0x94, 0xd1, 0xad, 0x8e, 0xad, 0x81, 0x7c, 0xa8, 0x51, 0x9e, 0x18,
	0x59, 0xea, 0x58, 0x1f, 0xc3, 0x3f, 0x3d, 0xb8, 0x57, 0x5d, 0x1f, 0x2b, 0xa2, 0x28, 0x7a, 0x04,
	0x3b, 0x52, 0x11, 0xb5, 0x92, 0x4e, 0xf4, 0x57, 0xa2, 0x1b, 0x8c, 0x68, 0x6c, 0xdc, 0xd8, 0xd1,
	0x10, 0x82, 0xfa, 0x32, 0x23, 0xc5, 0xbb, 0xe6, 0xac, 0x31, 0xc9, 0x7e, 0xa4, 0xdd, 0x9a, 0xc5,
	0xf4, 0x19, 0xbd, 0x03, 0x7b, 0x7c, 0x95, 0x4e, 0x67, 0x59, 0x9a, 0x67, 0x9c, 0x72, 0x25, 0xbb,
	0x75, 0xe3, 0xdd, 0xe5, 0xab, 0x74, 0x50, 0x82, 0xe8, 0x2d, 0xd8, 0x25, 0xeb, 0x79, 0x45, 0xeb,
	0x6e, 0x1b, 0x56, 0x9b, 0xac, 0xe7, 0x25, 0x2b, 0x7c, 0x1b, 0x76, 0x6c, 0x16, 0xba, 0x6d, 0xfd,
	0xcb, 0xfe, 0x70, 0xd4, 0x3f, 0x1e, 0xc5, 0xfe, 0x1d, 0xd4, 0x82, 0xc6, 0xf9, 0xe9, 0xe9, 0x68,
	0x78, 0x16, 0xfb, 0x5e, 0xf8, 0x8b, 0x07, 0xf7, 0x37, 0xd5, 0xb9, 0xca, 0x44, 0x4a, 0x14, 0xcb,
	0xb8, 0x16, 0x88, 0xf1, 0x84, 0xbe, 0x28, 0x04, 0x32, 0x06, 0xea, 0x42, 0x83, 0x24, 0x89, 0xa0,
	0x52, 0x9a, 0x62, 0x9a, 0xb8, 0x30, 0x75, 0x3d, 0x79, 0x26, 0x54, 0x51, 0x8f, 0x3e, 0xa3, 0x77,
	0x8d, 0xc8, 0x8a, 0x9a, 0x32, 0x5a, 0x3d, 0xff, 0xa6, 0x4e, 0xd8, 0xba, 0xd1, 0x23, 0xb8, 0xcb,
	0x5c, 0x63, 0x4c, 0x2d, 0xad, 0xde, 0x7e, 0x74, 0xbb, 0x67, 0xb8, 0x24, 0x85, 0x3f, 0x7b, 0xd0,
	0xa9, 0x08, 0x4f, 0x48, 0x9e, 0x33, 0x3e, 0xd7, 0xc9, 0xad, 0xa9, 0x90, 0x2c, 0xe3, 0x2e, 0xe9,
	0xc2, 0x44, 0x6f, 0x40, 0x4b, 0x0b, 0x6b, 0xff, 0x72, 0xd2, 0xf5, 0x01, 0xf8, 0x2a, 0xb5, 0x0f,
	0x48, 0xf4, 0x21, 0x34, 0x0a, 0x67, 0x2d, 0xa8, 0x1d, 0xb5, 0x7a, 0x0f, 0xa2, 0xff, 0x95, 0x05,
	0x17, 0x34, 0x1d, 0x2c, 0x27, 0x6a, 0xc1, 0xf8, 0xbc, 0x5b, 0x0f, 0x6a, 0x5a, 0x09, 0x67, 0x86,
	0xcf, 0xe1, 0xb5, 0xea, 0xee, 0x60, 0xc9, 0x28, 0x57, 0x9b, 0xc2, 0x3e, 0x84, 0xbb, 0x2b, 0x49,
	0x05, 0x27, 0x29, 0x35, 0x69, 0x36, 0x71, 0x69, 0x6b, 0x5f, 0x4e, 0xa4, 0x7c, 0x9e, 0x89, 0xc4,
	0xe9, 0x5b, 0xda, 0x28, 0x80, 0x56, 0x4e, 0x45, 0xca, 0xa4, 0xae, 0x48, 0x1a, 0x9d, 0x9b, 0x78,
	0x13, 0x0a, 0xff, 0xf0, 0xa0, 0x5b, 0x45, 0x8e, 0xf9, 0x4c, 0x5c, 0xe7, 0x8a, 0x26, 0x03, 0x41,
	0x13, 0x89, 0x7a, 0xd0, 0x62, 0x9c, 0xa9, 0xe9, 0x9a, 0xce, 0x54, 0x26, 0x4c, 0xe4, 0xf6, 0x71,
	0xe7, 0xaf, 0x7f, 0xfe, 0xfe, 0xdd, 0x6b, 0x41, 0xb3, 0x1f, 0x8f, 0x3f, 0xea, 0x7d, 0xf6, 0x98,
	0x5e, 0x63, 0xd0, 0xac, 0x4b, 0x43, 0x42, 0xaf, 0x03, 0xcc, 0x58, 0xbe, 0xa0, 0x42, 0xd1, 0x17,
	0x76, 0x59, 0xb4, 0xf1, 0x06, 0x82, 0xa6, 0xb0, 0x23, 0x17, 0xa4, 0xf7, 0xc9, 0xa7, 0x26, 0x9b,
	0xf6, 0xf1, 0x97, 0xe6, 0xb9, 0x3e, 0xfa, 0x62, 0xb2, 0xa0, 0x81, 0xf5, 0x04, 0xd9, 0x55, 0xa0,
	0x16, 0x34, 0xc8, 0x97, 0x84, 0xf1, 0x40, 0xdf, 0x0a, 0x68, 0x91, 0x58, 0xa0, 0x91, 0x05, 0x0d,
	0xaa, 0x17, 0x83, 0x2b, 0x46, 0x97, 0x49, 0x84, 0xdd, 0xb3, 0xe1, 0x37, 0xb7, 0xa5, 0xd4, 0xd5,
	0x50, 0xae, 0x18, 0x59, 0xea, 0x9a, 0xb6, 0xb5, 0x74, 0xfa, 0x7f, 0xa8, 0x7b, 0x76, 0x18, 0xbd,
	0x44, 0x77, 0x6c, 0xa9, 0xe1, 0x33, 0xd8, 0xaf, 0x58, 0x98, 0x7e, 0x4f, 0x96, 0x84, 0xcf, 0x28,
	0xda, 0x83, 0x2d, 0x96, 0xb8, 0x7e, 0x6c, 0xb1, 0x04, 0x7d, 0x00, 0x8d, 0xd4, 0x8e, 0x95, 0x5b,
	0x92, 0x28, 0xba, 0x35, 0x70, 0xb8, 0xa0, 0xa0, 0x07, 0xb0, 0x93, 0x66, 0x6b, 0x4d, 0xd6, 0xd3,
	0x53, 0xc7, 0xce, 0x0a, 0x7f, 0xf2, 0x36, 0x97, 0xcf, 0x29, 0x5b, 0xd2, 0x41, 0x96, 0x5f, 0xa3,
	0x37, 0xa1, 0x2d, 0x8a, 0xc8, 0xd3, 0x32, 0x6c, 0xab, 0xc4, 0x86, 0x09, 0x3a, 0x84, 0x66, 0xc2,
	0x84, 0x69, 0xc3, 0xb5, 0x1b, 0x85, 0x0a, 0xd0, 0x73, 0x72, 0xc5, 0x96, 0xd4, 0xcc, 0x90, 0x1d,
	0x84, 0xd2, 0x2e, 0x17, 0x4b, 0xbd, 0x5a, 0x2c, 0xe1, 0xb7, 0x2e, 0x0d, 0xbd, 0xe2, 0xfb, 0x2b,
	0xb5, 0x98, 0x64, 0xcf, 0xe8, 0xcb, 0x27, 0xf1, 0x00, 0xb6, 0x79, 0xc6, 0x67, 0xd4, 0xc5, 0xb6,
	0x86, 0x7e, 0x7b, 0x41, 0xe4, 0xc2, 0xc5, 0x34, 0xe7, 0xf0, 0x3b, 0x38, 0xdc, 0xf8, 0x7c, 0xcc,
	0x99, 0x54, 0xc2, 0x0a, 0x6e, 0xbf, 0x20, 0xe5, 0x62, 0xd0, 0x11, 0x76, 0xdd, 0x62, 0x78, 0x0f,
	0xb6, 0x95, 0x4e, 0xc1, 0x69, 0xbb, 0x1f, 0xdd, 0xce, 0x0e, 0x5b, 0xc6, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xb8, 0x1c, 0xc4, 0x94, 0x2d, 0x07, 0x00, 0x00,
}
