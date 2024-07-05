package hashes

// ToSlice - return []byte representation of the internal state
func (block *Sha512) ToSlice() []byte {
	b := make([]byte, Sha512Length)
	copy(b, block[:])
	return b
}
