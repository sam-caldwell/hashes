package hashes

import "testing"

func TestSha512_HashBytes(t *testing.T) {

	t.Run("Test String() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "a7096bad62963e64706aebbbf0558d4435373a5a98ee58200f29674938ccff461403a1c76febb3ebb88eddbf642a04d65bf328f97d9fab807afc5d88088f1b75"
		var hash Sha512
		hash.HashString(testInput)
		if actual := hash.String(); actual != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})

}
