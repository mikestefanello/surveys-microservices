package survey

// Serializer contains functions to encode and decode surveys
type Serializer interface {
	// Encode encodes a survey
	Encode(survey *Survey) ([]byte, error)

	// EncodeMultiple encodes surveys
	EncodeMultiple(surveys *Surveys) ([]byte, error)

	// EncodeErrorResponse encodes an error response
	EncodeErrorResponse(er ErrorResponse) ([]byte, error)

	// Decode decodes a survey
	Decode(data []byte) (*Survey, error)

	// DecodeMultiple decodes surveys
	DecodeMultiple(data []byte) (*Surveys, error)

	// GetContentType returns the content-type
	GetContentType() string
}
