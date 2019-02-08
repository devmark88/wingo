package response

type GetMetaResponse struct {
	HoursToNextContest       string `json:"hoursToNextContest"`
	MinToNextContest         string `json:"minToNextContest"`
	Prize                    uint   `json:"nextRoundPrice"`
	Tickets                  uint   `json:"userCredit"`
	NextContestNeededTickets byte   `json:"incomeTicketCount"`
	Correctors               uint   `json:"correctorCount"`
	SeekSeconds              uint   `json:"seekSeconds"`
	IsInDeadline             bool   `json:"isInDeadline"`
	SecondsToStart           uint   `json:"secondsToStart"`
	VideoUrl                 string `json:"videoUrl"`
}
