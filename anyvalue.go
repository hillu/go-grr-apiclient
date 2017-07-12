package apiclient

import (
	"github.com/golang/protobuf/proto"

	"encoding/json"
	"fmt"
	"reflect"
)

// ToAnyValue stores a proto.Message in an AnyValue
func ToAnyValue(pb proto.Message) (av *AnyValue, err error) {
	var value []byte
	typ := proto.MessageName(pb)
	value, err = proto.Marshal(pb)
	if err != nil {
		return
	}
	av = &AnyValue{
		TypeUrl: &typ,
		Value:   value,
	}
	return
}

// ToProtoMessage returnes the stored proto.Message
func (av *AnyValue) ToProtoMessage() (pb proto.Message, err error) {
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
	pb, err := av.ToProtoMessage()
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
