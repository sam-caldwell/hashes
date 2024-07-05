package hashes

import "testing"

func TestSha1_SizeOf(t *testing.T) {

	const expectedSize = Sha1Length

	t.Run("Test ToSlice() Happy 1", func(t *testing.T) {
		var hash Sha1
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})

	t.Run("Test ToSlice() Happy 2", func(t *testing.T) {
		var hash Sha1
		expected := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
		}
		hash.FromSlice(expected)
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
}
