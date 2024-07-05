package hashes

import (
	"encoding/hex"
	"testing"
)

func TestSha256_HashString(t *testing.T) {

	t.Run("Test HashString() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "a7096bad62963e64706aebbbf0558d4435373a5a98ee58200f29674938ccff461403a1c76febb3ebb88eddbf642a04d65bf328f97d9fab807afc5d88088f1b75"
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
