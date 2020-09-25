package vote

// Serializer contains functions to encode and decode votes
type Serializer interface {
	// Encode encodes a vote
	Encode(v *Vote) ([]byte, error)

	// EncodeResults encodes results
	EncodeResults(r *Results) ([]byte, error)

	// EncodeErrorResponse encodes an error response
	EncodeErrorResponse(err ErrorResponse) ([]byte, error)

	// Decode decodes a vote
	Decode(data []byte) (*Vote, error)

	// DecodeResults decodes results
	DecodeResults(data []byte) (*Results, error)

	// GetContentType returns the content-type
	GetContentType() string
}
