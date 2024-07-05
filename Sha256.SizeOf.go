package hashes

// SizeOf - return the size of the object
// Note that we return an unsigned integer because a negative link would not be rational
func (block *Sha256) SizeOf() uint {

	return uint(len(*block))

}
