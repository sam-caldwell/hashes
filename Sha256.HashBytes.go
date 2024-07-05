package hashes

import "crypto/sha256"

// HashBytes - hash the input and store as state
func (block *Sha256) HashBytes(data []byte) {

	*block = sha256.Sum256(data)

}
