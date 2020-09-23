package vote

// Serializer contains functions to encode and decode votes
type Serializer interface {
	// Encode encodes a vote
	Encode(vote *Vote) ([]byte, error)

	// EncodeErrorResponse encodes an error response
	EncodeErrorResponse(er ErrorResponse) ([]byte, error)

	// Decode decodes a vote
	Decode(data []byte) (*Vote, error)

	// GetContentType returns the content-type
	GetContentType() string
}
