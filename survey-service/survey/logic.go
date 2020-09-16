package survey

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/teris-io/shortid"
)

var (
	// ErrNotFound indicates that a requested survey was not found
	ErrNotFound = errors.New("Survey not found")

	// ErrInvalidRequest indicates that an invalid survey was provided
	ErrInvalidRequest = errors.New("Invalid survey input")
)

// ErrorResponse provides a structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

type surveyService struct {
	repository Repository
	validator  *validator.Validate
}

// NewService creates a new survey service
func NewService(repository Repository) Service {
	return &surveyService{
		repository: repository,
		validator:  validator.New(),
	}
}

func (s *surveyService) Insert(survey *Survey) error {
	if err := s.validator.Struct(survey); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
	}

	survey.ID = shortid.MustGenerate()
	survey.CreatedAt = time.Now().UTC().Unix()

	for k := range survey.Questions {
		survey.Questions[k].ID = k + 1
	}

	return s.repository.Insert(survey)
}

func (s *surveyService) LoadByID(id string) (*Survey, error) {
	return s.repository.LoadByID(id)
}

func (s *surveyService) Load() (*Surveys, error) {
	return s.repository.Load()
}
