package hashes

// ToSlice - return []byte representation of the internal state
func (block *Sha1) ToSlice() []byte {
	b := make([]byte, Sha1Length)
	copy(b, block[:])
	return b
}
