package repository

import (
	"sort"
	"sync"

	"github.com/mikestefanello/surveys-microservices/survey-service/survey"
)

type surveyMemoryRepository struct {
	storage map[string]*survey.Survey
	mutex   *sync.Mutex
}

// NewSurveyMemoryRepository creates a new survey repository that stores in memory
func NewSurveyMemoryRepository() (survey.Repository, error) {
	return &surveyMemoryRepository{
		storage: make(map[string]*survey.Survey),
		mutex:   &sync.Mutex{},
	}, nil
}

func (r *surveyMemoryRepository) Insert(survey *survey.Survey) error {
	r.mutex.Lock()
	r.storage[survey.ID] = survey
	r.mutex.Unlock()
	return nil
}

func (r *surveyMemoryRepository) LoadByID(id string) (*survey.Survey, error) {
	if _, ok := r.storage[id]; !ok {
		return nil, survey.ErrNotFound
	}
	return r.storage[id], nil
}

func (r *surveyMemoryRepository) Load() (*survey.Surveys, error) {
	surveys := make(survey.Surveys, 0, len(r.storage))

	for _, item := range r.storage {
		surveys = append(surveys, item)
	}

	sort.Slice(surveys, func(i, j int) bool {
		return surveys[i].CreatedAt > surveys[j].CreatedAt
	})

	return &surveys, nil
}
