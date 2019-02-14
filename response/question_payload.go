package response

// QuestionPayload : payload data for Question message
type QuestionPayload struct {
	ID        uint   `json:"id"`
	ContestID uint   `josn:"cid"`
	Index     int    `json:"index"`
	Body      string `json:"text"`
	Options   string `json:"options"`
	Type      string `json:"type"`
}
