package apiclient

import (
	"encoding/hex"
	"testing"
)

func TestAnyValue(t *testing.T) {
	var duration uint64 = 1800
	av, err := ToAnyValue(&KeepAliveArgs{Duration: &duration})
	if err != nil {
		t.Errorf("ToAnyValue: %s", err)
	}
	t.Log("TypeURL: ", *av.TypeUrl)
	t.Log("Value:   ", hex.EncodeToString(av.Value))

	pb, err := av.ToProtoMessage()
	if err != nil {
		t.Errorf("ToProtoMessage: %s", err)
	}
	t.Logf("Type of Regenerated value: %t", pb)
	v, ok := pb.(*KeepAliveArgs)
	if !ok {
		t.Fatal("Unexpected type!")
	}
	t.Logf("- Duration: %d", v.GetDuration())
}
