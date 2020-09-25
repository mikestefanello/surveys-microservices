package vote

// WriterRepository contains functions to write votes to a repository
type WriterRepository interface {
	// Insert stores a new vote
	Insert(v *Vote) error
}

// ResultsRepository contains functions to read votes from a repository
type ResultsRepository interface {
	// GetResults gets the results for a given survey
	GetResults(surveyID string) (Results, error)
}
