package handler

import (
	"context"
	"errors"

	protos "github.com/mikestefanello/surveys-microservices/survey-service/protos/survey"
	"github.com/mikestefanello/surveys-microservices/survey-service/survey"
	"github.com/rs/zerolog"
)

// SurveyGrpcHandler handles gRPC requests for surveys
type SurveyGrpcHandler struct {
	service survey.Service
	log     *zerolog.Logger
}

// NewSurveyGrpcHandler created a new survey gRPC handler
func NewSurveyGrpcHandler(service survey.Service, log *zerolog.Logger) *SurveyGrpcHandler {
	return &SurveyGrpcHandler{service, log}
}

// GetSurvey loads and returns a requested survey
func (g *SurveyGrpcHandler) GetSurvey(ctx context.Context, r *protos.SurveyRequest) (*protos.SurveyResponse, error) {
	id := r.GetId()
	g.log.Info().Str("id", id).Msg("GetSurvey request received")

	s, err := g.service.LoadByID(id)
	if err != nil {
		if errors.Is(err, survey.ErrNotFound) {
			g.log.Debug().Str("id", id).Msg("Survey not found")
		} else {
			g.log.Error().Err(err).Str("id", id).Msg("Unable to load survey")
		}
		return nil, err
	}

	res := &protos.SurveyResponse{
		Id:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
	}

	for _, q := range s.Questions {
		res.Questions = append(res.Questions, &protos.QuestionResponse{
			Id:   int32(q.ID),
			Text: q.Text,
		})
	}

	return res, nil
}
