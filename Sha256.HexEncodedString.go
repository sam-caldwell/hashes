package hashes

import "encoding/hex"

// HexEncodedString - given a hexadecimal-encoded string, return the hash bytes
func (block *Sha256) HexEncodedString(s *string) {

	hash, err := hex.DecodeString(*s)

	if err != nil {
		panic(err)
	}

	*block = [Sha256Length]byte(hash[:Sha256Length])

}
