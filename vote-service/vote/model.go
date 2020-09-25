package vote

// Vote describes a vote
type Vote struct {
	ID        string `json:"id"`
	Survey    string `json:"survey" validate:"required"`
	Question  int    `json:"question" validate:"required,min=1"`
	Timestamp int64  `json:"timestamp"`
}

// Results describes the results of a survey
type Results struct {
	Survey    string   `json:"survey"`
	Results   []Result `json:"results"`
	UpdatedAt int64    `json:"updatedAt"`
}

// Result describes the voting results of a given survey question
type Result struct {
	Question int `json:"question"`
	Votes    int `json:"votes"`
}
