package hashes

import (
	"encoding/hex"
	"testing"
)

func TestSha256_HashBytes(t *testing.T) {

	t.Run("Test HashBytes() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "6e5d6a1041fe8a13fe5f8062e1c8d2741d71aae225b46c2efc539230f2113d25"
		var hash Sha256
		hash.HashBytes([]byte(testInput))
		actual := hash.ToSlice()
		if hex.EncodeToString(actual) != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})

}
