package hashes

import (
	"testing"
)

func TestHashes_constants(t *testing.T) {
	t.Run("Test Sha1Length", func(t *testing.T) {
		const expectedSize = uint(20)
		if expectedSize != Sha1Length {
			t.Fatal("length is wrong")
		}
	})
}
