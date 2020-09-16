package survey

// Survey describes a survey
type Survey struct {
	ID        string     `json:"id" bson:"id"`
	Name      string     `json:"name" bson:"name" validate:"required"`
	Questions []Question `json:"questions" bson:"questions" validate:"required,min=2"`
	CreatedAt int64      `json:"createdAt" bson:"createdAt"`
}

// Surveys is a slice of survey pointers
type Surveys []*Survey

// Question describes a survey question
type Question struct {
	ID   int    `json:"id" bson:"id"`
	Text string `json:"text" bson:"text" validate:"required"`
}
