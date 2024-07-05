package hashes

import "encoding/hex"

// HexEncodedString - given a hexadecimal-encoded string, return the hash bytes
func (block *Sha1) HexEncodedString(s *string) {

	hash, err := hex.DecodeString(*s)

	if err != nil {

		panic(err)

	}

	*block = [Sha1Length]byte(hash[:Sha1Length])
}
