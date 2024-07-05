package hashes

// ToSlice - return []byte representation of the internal state
func (block *Sha256) ToSlice() []byte {

	b := make([]byte, Sha256Length)

	copy(b, block[:])

	return b

}
