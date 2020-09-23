package vote

// Service contains functions to create and load votes
type Service interface {
	// Insert stores a new vote
	Insert(vote *Vote) error
}
