package hashes

import "encoding/hex"

// String - return the contents of the hash object as a hex string
func (block *Sha512) String() string {
	return hex.EncodeToString(block[:])
}
