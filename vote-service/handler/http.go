package handler

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mikestefanello/surveys-microservices/vote-service/middleware"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
	"github.com/rs/zerolog"
)

// VoteHTTPHandler handles HTTP requests for votes
type VoteHTTPHandler struct {
	service vote.Service
	log     *zerolog.Logger
}

// NewVoteHTTPHandler creates a new vote HTTP handler
func NewVoteHTTPHandler(service vote.Service, log *zerolog.Logger) *VoteHTTPHandler {
	return &VoteHTTPHandler{
		service,
		log,
	}
}

// Vote handles post requests to cast a vote
func (h *VoteHTTPHandler) Vote(w http.ResponseWriter, r *http.Request) {
	h.log.Info().Msg("POST request received: Vote")

	// Read the request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to read vote POST body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Decode the request
	serializer := h.GetSerializer(r)
	v, err := serializer.Decode(requestBody)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to decode vote POST body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Save the vote
	err = h.service.Insert(v)
	if err != nil {
		if errors.Is(err, vote.ErrInvalidRequest) {
			h.log.Debug().Err(err).Msg("Invalid vote data in POST")
			h.Error(w, r, err.Error(), http.StatusUnprocessableEntity)
		} else {
			h.log.Error().Err(err).Msg("Unable to save vote")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("id", v.ID).Msg("Vote created")

	// Encode the vote to be returned
	res, err := serializer.Encode(v)
	if err != nil {
		h.log.Error().Str("id", v.ID).Err(err).Msg("Unable to encode vote")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, res, http.StatusCreated)
}

// GetResults handles get requests to get results for a given survey
func (h *VoteHTTPHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("GET request received: GetResults")

	results, err := h.service.GetResults(id)
	if err != nil {
		if errors.Is(err, vote.ErrResultsNotFound) {
			h.log.Debug().Str("id", id).Msg("Invalid survey requested for results")
			h.Error(w, r, err.Error(), http.StatusNotFound)
		} else {
			h.log.Error().Str("id", id).Err(err).Msg("Unable to load results")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Encode the results to be returned
	res, err := h.GetSerializer(r).EncodeResults(&results)
	if err != nil {
		h.log.Error().Str("id", id).Err(err).Msg("Unable to encode results")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, res, http.StatusOK)
}

// GetSerializer gets a serializer from the request, which is added via middleware
func (h *VoteHTTPHandler) GetSerializer(r *http.Request) vote.Serializer {
	return r.Context().Value(middleware.SerializerKey).(vote.Serializer)
}

// Response sends an HTTP response
func (h *VoteHTTPHandler) Response(w http.ResponseWriter, r *http.Request, output []byte, code int) {
	w.Header().Set("Content-Type", h.GetSerializer(r).GetContentType())
	w.WriteHeader(code)
	w.Write(output)
}

// Error sends an HTTP error response
func (h *VoteHTTPHandler) Error(w http.ResponseWriter, r *http.Request, message string, code int) {
	serializer := h.GetSerializer(r)
	output, err := serializer.EncodeErrorResponse(vote.ErrorResponse{Error: message})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, output, code)
}
