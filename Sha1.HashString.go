package hashes

// HashString - hash the input and store as state
func (block *Sha1) HashString(data string) {

	block.HashBytes([]byte(data))

}
