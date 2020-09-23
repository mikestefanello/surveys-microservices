package middleware

import (
	"context"
	"net/http"

	"github.com/mikestefanello/surveys-microservices/vote-service/serializer"
	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
)

type key int

// SerializerKey is used as a key to store the serializer in the context
const SerializerKey key = 0

// AddSerializer adds a serializer to the request context
func AddSerializer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s vote.Serializer

		// Determine which serializer to use based on the content-type.
		switch r.Header.Get("Content-Type") {
		// TODO: One day add more serializers
		// TODO: Store the serializer so a new instance doesn't need to be created on each request
		// TODO: Use .GetContentType() to match to the Content-Type header
		default:
			s = serializer.NewVoteJSONSerializer()
		}

		// Store the serializer in context so it can be used by handlers.
		ctx := context.WithValue(r.Context(), SerializerKey, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
