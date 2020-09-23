package vote

// WriterRepository contains functions to write votes to a repository
type WriterRepository interface {
	// Insert stores a new vote
	Insert(v *Vote) error
}

// ReaderRepository contains functions to read votes from a repository
type ReaderRepository interface {
}
