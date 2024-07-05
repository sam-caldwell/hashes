package hashes

import "crypto/sha512"

// HashBytes - hash the input and store as state
func (block *Sha512) HashBytes(data []byte) {
	*block = sha512.Sum512(data)
}
