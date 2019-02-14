package request

// PostAnswer : post answer request
type PostAnswer struct {
	QuestionID    uint `json:"id"`
	ContestID     uint `json:"cid"`
	SelectedIndex int  `json:"index"`
}
