package hashes

// HashString - hash the input and store as state
func (block *Sha256) HashString(data string) {

	block.HashBytes([]byte(data))

}
