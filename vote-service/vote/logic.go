package vote

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	protos "github.com/mikestefanello/surveys-microservices/survey-service/protos/survey"
	uuid "github.com/satori/go.uuid"
)

var (
	// ErrInvalidRequest indicates that an invalid vote was provided
	ErrInvalidRequest = errors.New("Invalid vote input")

	// ErrResultsNotFound indicates that results could not be found for a given survey
	ErrResultsNotFound = errors.New("Results not found")
)

// ErrorResponse provides a structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

type voteService struct {
	writer    WriterRepository
	results   ResultsRepository
	validator *validator.Validate
	surveys   protos.SurveyClient
}

// NewService creates a new survey service
func NewService(w WriterRepository, r ResultsRepository, cli protos.SurveyClient) Service {
	return &voteService{
		writer:    w,
		results:   r,
		validator: validator.New(),
		surveys:   cli,
	}
}

func (s *voteService) Insert(v *Vote) error {
	if err := s.validator.Struct(v); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
	}

	// Fetch the survey
	req := &protos.SurveyRequest{Id: v.Survey}
	surv, err := s.surveys.GetSurvey(context.Background(), req)
	if err != nil {
		return ErrInvalidRequest
	}

	// The survey ID is validated, but we need to validate the question ID
	valid := false
	for _, q := range surv.GetQuestions() {
		if q.GetId() == int32(v.Question) {
			valid = true
			break
		}
	}
	if !valid {
		return ErrInvalidRequest
	}

	v.ID = uuid.NewV4().String()
	v.Timestamp = time.Now().UTC().Unix()

	return s.writer.Insert(v)
}

func (s *voteService) GetResults(surveyID string) (Results, error) {
	return s.results.GetResults(surveyID)
}
