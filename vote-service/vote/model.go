package vote

// Vote describes a vote
type Vote struct {
	ID        string `json:"id"`
	Survey    string `json:"survey" validate:"required"`
	Question  int    `json:"question" validate:"required,min=1"`
	Timestamp int64  `json:"timestamp"`
}
