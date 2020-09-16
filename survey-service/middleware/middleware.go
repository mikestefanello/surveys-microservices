package middleware

import (
	"context"
	"net/http"

	"github.com/mikestefanello/surveys-microservices/survey-service/serializer"
	"github.com/mikestefanello/surveys-microservices/survey-service/survey"
)

type key int

// SerializerKey is used as a key to store the serializer in the context
const SerializerKey key = 0

// AddSerializer adds a serializer to the request context
func AddSerializer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s survey.Serializer

		// Determine which serializer to use based on the content-type.
		switch r.Header.Get("Content-Type") {
		// TODO: One day add more serializers
		default:
			s = serializer.NewSurveyJSONSerializer()
		}

		// Store the serializer in context so it can be used by handlers.
		ctx := context.WithValue(r.Context(), SerializerKey, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
