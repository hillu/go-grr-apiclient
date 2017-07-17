package apiclient

import (
	"encoding/hex"
	"testing"
)

func TestAnyValue(t *testing.T) {
	var duration uint64 = 1800
	av, err := NewAnyValue(&KeepAliveArgs{Duration: &duration})
	if err != nil {
		t.Errorf("NewAnyValue: %s", err)
	}
	t.Log("TypeURL: ", *av.TypeUrl)
	t.Log("Value:   ", hex.EncodeToString(av.Value))

	pb, err := av.GetProtoMessage()
	if err != nil {
		t.Errorf("GetProtoMessage: %s", err)
	}
	t.Logf("Type of Regenerated value: %t", pb)
	v, ok := pb.(*KeepAliveArgs)
	if !ok {
		t.Fatal("Unexpected type!")
	}
	t.Logf("- Duration: %d", v.GetDuration())
}
