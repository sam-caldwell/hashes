package hashes

import (
	"testing"
)

func TestStateMetric_StateType_Sha256(t *testing.T) {

	var s Sha256
	if len(s) != Sha256Length {
		t.Fatal("size mismatch")
	}

}
