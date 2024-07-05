package hashes

import "crypto/sha1"

// HashBytes - hash the input and store as state
func (block *Sha1) HashBytes(data []byte) {

	*block = sha1.Sum(data)

}
