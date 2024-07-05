package hashes

import (
	"encoding/hex"
	"testing"
)

func TestSha1_HashString(t *testing.T) {

	t.Run("Test HashString() Happy", func(t *testing.T) {

		const testInput = "This is my test.  There is no test like my test."

		const expected = "027c4109f56c0dffdcfd2e0dc50fdc97488d5300"

		var hash Sha1

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
