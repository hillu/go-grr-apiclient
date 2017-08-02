package apiclient

import (
	"github.com/golang/protobuf/proto"

	"encoding/json"
	"fmt"
	"reflect"
)

// NewAnyValue creates and returns a new AnyValue in which the type
// information and the serialized form of a proto.Message are stored.
func NewAnyValue(pb proto.Message) (*AnyValue, error) {
	var av AnyValue
	typ := proto.MessageName(pb)
	if pb == nil || typ == "" {
		return &av, nil
	}
	av.TypeUrl = &typ
	value, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}
	av.Value = value
	return &av, nil
}

// MustNewAnyValue behaves like NewAnyValue but panics if an error
// occurs.
func MustNewAnyValue(pb proto.Message) *AnyValue {
	av, err := NewAnyValue(pb)
	if err != nil {
		panic(err)
	}
	return av
}

// GetProtoMessage deserializes and returns the proto.Message stored
// in the AnyValue
func (av *AnyValue) GetProtoMessage() (pb proto.Message, err error) {
	typ := proto.MessageType(av.GetTypeUrl())
	if typ == nil {
		return nil, fmt.Errorf("%s has not been registered with Protobuf library", av.TypeUrl)
	}
	pb = reflect.New(typ.Elem()).Interface().(proto.Message)
	if err = proto.Unmarshal(av.Value, pb); err != nil {
		return nil, err
	}
	return
}

// MarshalJSON implements a JSON marshaller that marshals the value
// stored in av.Value.
func (av *AnyValue) MarshalJSON() ([]byte, error) {
	pb, err := av.GetProtoMessage()
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
