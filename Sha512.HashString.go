package hashes

// HashString - hash the input and store as state
func (block *Sha512) HashString(data string) {
	block.HashBytes([]byte(data))
}
