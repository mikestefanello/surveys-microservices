package survey

// Repository contains functions to store and fetch surveys from a repository
type Repository interface {
	// Insert stores a new survey
	Insert(survey *Survey) error

	// LoadById loads a survey by ID
	LoadByID(id string) (*Survey, error)

	// Load loads all surveys
	Load() (*Surveys, error)
}
