package hashes

import (
	"testing"
)

func TestStateMetric_StateType_Sha1(t *testing.T) {
	var s Sha1
	if len(s) != Sha1Length {
		t.Fatal("size mismatch")
	}
}
