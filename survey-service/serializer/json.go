package serializer

import (
	"encoding/json"

	"github.com/mikestefanello/surveys-microservices/survey-service/survey"
)

type surveyJSONSerializer struct{}

// NewSurveyJSONSerializer creates a new survey JSON serializer.
func NewSurveyJSONSerializer() survey.Serializer {
	return &surveyJSONSerializer{}
}

func (s *surveyJSONSerializer) Encode(survey *survey.Survey) ([]byte, error) {
	return json.Marshal(survey)
}

func (s *surveyJSONSerializer) EncodeMultiple(surveys *survey.Surveys) ([]byte, error) {
	return json.Marshal(surveys)
}

func (s *surveyJSONSerializer) EncodeErrorResponse(er survey.ErrorResponse) ([]byte, error) {
	return json.Marshal(er)
}

func (s *surveyJSONSerializer) Decode(data []byte) (*survey.Survey, error) {
	survey := survey.Survey{}
	err := json.Unmarshal(data, &survey)
	return &survey, err
}

func (s *surveyJSONSerializer) DecodeMultiple(data []byte) (*survey.Surveys, error) {
	surveys := survey.Surveys{}
	err := json.Unmarshal(data, &surveys)
	return &surveys, err
}

func (s *surveyJSONSerializer) GetContentType() string {
	return "application/json"
}
