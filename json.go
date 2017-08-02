package apiclient

import (
	"github.com/golang/protobuf/proto"

	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"reflect"
)

var protoMessageType = reflect.TypeOf(new(proto.Message)).Elem()
var anyValueType = reflect.TypeOf(new(AnyValue))

// unmarshalGrrJSON is a constrained JSON deserializer that can deal
// with both the type-stripped data and the age/type/value records
// returned by the GRR API.
//
// It supports unmarshalling to the following data types:
//
// - elementary: uint/int/float of any size,bool, string
//   (using json.Unmarshal)
// - proto.Message
// - AnyValue (generic value that stores a proto.Message with type
//   information)
// - pointers and slices of the above
func unmarshalGrrJSON(buf []byte, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.IsNil() {
		return errors.New("output value must not be nil")
	}
	if val.Kind() != reflect.Ptr {
		return errors.New("output type must be a pointer")
	}
	typ := reflect.TypeOf(v).Elem()
	val = val.Elem()

	// Deal with age/type/value records first, they are only relevant
	// for AnyValue.
	var atv struct {
		Age   *uint64
		Type  *string
		Value json.RawMessage
	}
	if json.Unmarshal(buf, &atv) == nil &&
		atv.Age != nil && atv.Type != nil && len(atv.Value) > 0 {

		if typ == anyValueType {
			if typ = proto.MessageType(*atv.Type); typ == nil {
				return nil
			}
			pb := reflect.New(typ.Elem()).Interface().(proto.Message)
			if err := unmarshalGrrJSON(atv.Value, pb); err != nil {
				return err
			}
			av, err := NewAnyValue(pb)
			if err != nil {
				return err
			}
			val.Set(reflect.ValueOf(av))
			return nil
		} else {
			return unmarshalGrrJSON(atv.Value, v)
		}
	}

	// Only deserialize struct types via proto.MessageType
	if val.Addr().Type().Implements(protoMessageType) {
		structMap := make(map[string]json.RawMessage)
		if err := json.Unmarshal(buf, &structMap); err != nil {
			return err
		}
		for _, p := range proto.GetProperties(typ).Prop {
			if fieldbuf, ok := structMap[p.OrigName]; ok {
				fv := val.FieldByName(p.Name)
				fv.Set(reflect.New(fv.Type()).Elem())
				if err := unmarshalGrrJSON(fieldbuf, fv.Addr().Interface()); err != nil {
					return err
				}
			}
		}
		return nil
	}

	switch typ.Kind() {
	case reflect.Ptr:
		val.Set(reflect.New(typ.Elem()))
		return unmarshalGrrJSON(buf, val.Interface())
	case reflect.Slice:
		// special case for []byte
		if typ.Elem().Kind() == reflect.Uint8 {
			var s string
			if err := json.Unmarshal(buf, &s); err != nil {
				return errors.New("Error while decoding " + string(buf) +
					" to string ([]byte): " + err.Error())
			}
			// Note: Base64-decoding should not be blindly attempted
			// but based upon hard facts.
			b, err := base64.StdEncoding.DecodeString(s)
			if err != nil {
				b = []byte(s)
			}
			val.Set(reflect.ValueOf(b))
			return nil
		}
		structSlice := []json.RawMessage{}
		if err := json.Unmarshal(buf, &structSlice); err != nil {
			return err
		}
		for _, fieldbuf := range structSlice {
			elem := reflect.New(typ.Elem())
			if err := unmarshalGrrJSON(fieldbuf, elem.Interface()); err != nil {
				return err
			}
			val.Set(reflect.Append(val, elem.Elem()))
		}
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.String,
		reflect.Bool:
		if err := json.Unmarshal(buf, v); err != nil {
			return errors.New("Error while decoding " + string(buf) + " to " +
				typ.Kind().String() + " : " + err.Error())
		}
		return nil
	default:
		return errors.New("Will not unmarshal to a " + typ.Kind().String())
	}
}

func decodeGrrJSON(r io.Reader, v interface{}) error {
	buf, _ := ioutil.ReadAll(r)
	return unmarshalGrrJSON(buf, v)
}
