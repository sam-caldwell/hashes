package hashes

import (
	"testing"
)

func TestStateMetric_StateType_Sha512(t *testing.T) {

	var s Sha512
	if len(s) != Sha512Length {
		t.Fatal("size mismatch")
	}

}
