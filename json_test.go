package apiclient

import (
	"testing"
)

func TestATVJSON(t *testing.T) {
	var val int

	if err := unmarshalGrrJSON(
		[]byte(`{"type":"string","age":0,"value":null}`), &val,
	); err != nil {
		t.Error(err)
	}
}
