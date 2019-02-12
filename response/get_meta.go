package response

// GetMetaResponse : GetMetaResponse response
type GetMetaResponse struct {
	NextContestInHours       string     `json:"hoursToNextContest"`
	NextContestInMinutes     string     `json:"minToNextContest"`
	NextContestInSeconds     string     `json:"secToNextContest"`
	Prize                    uint       `json:"nextRoundPrice"`
	NextContestNeededTickets byte       `json:"incomeTicketCount"`
	Tickets                  uint       `json:"userCredit"`
	Correctors               uint       `json:"correctorCount"`
	SeekSeconds              uint       `json:"seekSeconds"`
	IsInDeadline             bool       `json:"isInDeadline"`
	SecondsToStart           uint       `json:"secondsToStart"`
	VideoURL                 string     `json:"videoUrl"`
	ShareData                ShareData  `json:"shareData"`
	Timeline                 []Timeline `json:"timeline"`
}

// ShareData : ShareData response
type ShareData struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Message     string `json:"message"`
	DialogTitle string `json:"dialogTitle"`
}

// Timeline : Timeline response
type Timeline struct {
	Text      string `json:"text"`
	IsPast    bool   `json:"isPast"`
	IsCurrent bool   `json:"isCurrent"`
	StartTime string `json:"startTime"`
	Prize     uint   `json:"prize"`
	Currency  string `json:"currency"`
}
