package request

// PostAnswer : post answer request
type PostAnswer struct {
	QuestionID    int `json:"id"`
	SelectedIndex int `json:"index"`
}
