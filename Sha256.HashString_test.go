package hashes

import (
	"encoding/hex"
	"testing"
)

func TestSha256_HashString(t *testing.T) {

	t.Run("Test HashString() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "6e5d6a1041fe8a13fe5f8062e1c8d2741d71aae225b46c2efc539230f2113d25"
		var hash Sha256
		hash.HashString(testInput)
		result := hash.ToSlice()
		if actual := hex.EncodeToString(result); actual != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
}
