package vote

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)

var (
	// ErrInvalidRequest indicates that an invalid vote was provided
	ErrInvalidRequest = errors.New("Invalid vote input")
)

// ErrorResponse provides a structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

type voteService struct {
	writer    WriterRepository
	reader    ReaderRepository
	validator *validator.Validate
}

// NewService creates a new survey service
func NewService(w WriterRepository, r ReaderRepository) Service {
	return &voteService{
		writer:    w,
		reader:    r,
		validator: validator.New(),
	}
}

func (s *voteService) Insert(vote *Vote) error {
	if err := s.validator.Struct(vote); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
	}

	// TODO: Validate that the survey and question exist!

	vote.ID = uuid.NewV4().String()
	vote.Timestamp = time.Now().UTC().Unix()

	return s.writer.Insert(vote)
}
