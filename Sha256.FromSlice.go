package hashes

// FromSlice - Given a block of []byte, store to our internal state
func (block *Sha256) FromSlice(b []byte) {

	copy(block[:], b)

}
