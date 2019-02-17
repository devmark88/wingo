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

// UserInfoPayload : payload data for Question message
type UserInfoPayload struct {
	ID        string `json:"id"`
	Ticket    uint   `josn:"ticket"`
	Corrector uint   `json:"corrector"`
	CanPlay   bool   `json:"canPlay"`
	Type      string `json:"type"`
}
