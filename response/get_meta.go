package response

type GetMetaResponse struct {
	NextContestInHours       string `json:"hoursToNextContest"`
	NextContestInMinutes     string `json:"minToNextContest"`
	NextContestInSeconds     string `json:"secToNextContest"`
	Prize                    uint   `json:"nextRoundPrice"`
	NextContestNeededTickets byte   `json:"incomeTicketCount"`
	Tickets                  uint8  `json:"userCredit"`
	Correctors               uint8  `json:"correctorCount"`
	SeekSeconds              uint   `json:"seekSeconds"`
	IsInDeadline             bool   `json:"isInDeadline"`
	SecondsToStart           uint   `json:"secondsToStart"`
	VideoUrl                 string `json:"videoUrl"`
}
