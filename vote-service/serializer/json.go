package serializer

import (
	"encoding/json"

	"github.com/mikestefanello/surveys-microservices/vote-service/vote"
)

type voteJSONSerializer struct{}

// NewVoteJSONSerializer creates a new vote JSON serializer.
func NewVoteJSONSerializer() vote.Serializer {
	return &voteJSONSerializer{}
}

func (s *voteJSONSerializer) Encode(v *vote.Vote) ([]byte, error) {
	return json.Marshal(v)
}

func (s *voteJSONSerializer) EncodeErrorResponse(er vote.ErrorResponse) ([]byte, error) {
	return json.Marshal(er)
}

func (s *voteJSONSerializer) Decode(data []byte) (*vote.Vote, error) {
	v := vote.Vote{}
	err := json.Unmarshal(data, &v)
	return &v, err
}

func (s *voteJSONSerializer) GetContentType() string {
	return "application/json"
}
