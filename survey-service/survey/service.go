package survey

// Service contains functions to create and fetch surveys
type Service interface {
	// Insert stores a new survey
	Insert(survey *Survey) error

	// LoadById loads a survey by ID
	LoadByID(id string) (*Survey, error)

	// Load loads all surveys
	Load() (*Surveys, error)
}
