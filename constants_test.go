package hashes

import (
	"testing"
)

func TestHashes_constants(t *testing.T) {
	t.Run("Test Sha1Length", func(t *testing.T) {
		// Expect 20 byte hash
		const expectedSize = uint(20)
		if expectedSize != Sha1Length {
			t.Fatal("length is wrong")
		}
	})
	t.Run("Test Sha256Length", func(t *testing.T) {
		// Expect 32 byte hash
		const expectedSize = uint(32)
		if expectedSize != Sha256Length {
			t.Fatal("length is wrong")
		}
	})
	t.Run("Test Sha512Length", func(t *testing.T) {
		// Expect 64 byte hash
		const expectedSize = uint(64)
		if expectedSize != Sha512Length {
			t.Fatal("length is wrong")
		}
	})
}
