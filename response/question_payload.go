package response

type QuestionPayload struct {
	ID      uint   `json:"id"`
	Index   int    `json:"index"`
	Body    string `json:"text"`
	Options string `json:"options"`
}
type AnswerPayload struct {
	Order uint   `json:"order"`
	Index int    `json:"index"`
	Body  string `json:"text"`
}
