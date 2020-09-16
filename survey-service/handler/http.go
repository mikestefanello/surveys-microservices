package handler

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mikestefanello/surveys-microservices/survey-service/middleware"
	"github.com/mikestefanello/surveys-microservices/survey-service/survey"
	"github.com/rs/zerolog"
)

// SurveyHTTPHandler handles HTTP requests for surveys
type SurveyHTTPHandler struct {
	service survey.Service
	log     *zerolog.Logger
}

// NewSurveyHTTPHandler creates a new survey HTTP handler
func NewSurveyHTTPHandler(service survey.Service, log *zerolog.Logger) *SurveyHTTPHandler {
	return &SurveyHTTPHandler{
		service,
		log,
	}
}

// Get handles get requests to get a single survey
func (h *SurveyHTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("GET request received")

	// Load the requested survey
	item, err := h.service.LoadByID(id)
	if err != nil {
		if errors.Is(err, survey.ErrNotFound) {
			h.log.Debug().Str("id", id).Msg("Survey not found")
			h.Error(w, r, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			h.log.Error().Err(err).Str("id", id).Msg("Unable to load survey")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Encode the survey to be returned
	serializer := h.GetSerializer(r)
	json, err := serializer.Encode(item)
	if err != nil {
		h.log.Error().Str("id", id).Msg("Unable to encode survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, json, http.StatusOK)
}

// Collection handles get requests to return a collection of surveys
func (h *SurveyHTTPHandler) Collection(w http.ResponseWriter, r *http.Request) {
	h.log.Info().Msg("COLLECTION request received")

	// Load all surveys
	items, err := h.service.Load()
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to load surveys")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Encode the surveys to be returned
	serializer := h.GetSerializer(r)
	json, err := serializer.EncodeMultiple(items)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to encode surveys")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, json, http.StatusOK)
}

// Post handles post requests to create a survey
func (h *SurveyHTTPHandler) Post(w http.ResponseWriter, r *http.Request) {
	h.log.Info().Msg("POST request received")

	// Read the request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to read survey POST body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Decode the request
	serializer := h.GetSerializer(r)
	s, err := serializer.Decode(requestBody)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to decode survey POST body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Save the survey
	err = h.service.Insert(s)
	if err != nil {
		if errors.Is(err, survey.ErrInvalidRequest) {
			h.log.Debug().Err(err).Msg("Invalid survey data in POST")
			h.Error(w, r, err.Error(), http.StatusUnprocessableEntity)
		} else {
			h.log.Error().Err(err).Msg("Unable to save survey")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("id", s.ID).Msg("Survey created")

	// Encode the survey to be returned
	json, err := serializer.Encode(s)
	if err != nil {
		h.log.Error().Str("id", s.ID).Err(err).Msg("Unable to encode survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, json, http.StatusCreated)
}

// GetSerializer gets a serializer from the request, which is added via middleware
func (h *SurveyHTTPHandler) GetSerializer(r *http.Request) survey.Serializer {
	return r.Context().Value(middleware.SerializerKey).(survey.Serializer)
}

// Response sends an HTTP response
func (h *SurveyHTTPHandler) Response(w http.ResponseWriter, r *http.Request, output []byte, code int) {
	w.Header().Set("Content-Type", h.GetSerializer(r).GetContentType())
	w.WriteHeader(code)
	w.Write(output)
}

// Error sends an HTTP error response
func (h *SurveyHTTPHandler) Error(w http.ResponseWriter, r *http.Request, message string, code int) {
	serializer := h.GetSerializer(r)
	output, err := serializer.EncodeErrorResponse(survey.ErrorResponse{Error: message})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, output, code)
}
