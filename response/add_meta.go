package response

// AddMeta : AddMeta response
type AddMeta struct {
	ID                         uint   `json:"id"`
	AppID                      string `json:"app"`
	Title                      string `json:"title"`
	Prize                      uint   `json:"prize"`
	BeginTime                  string `json:"beginDateTime"`
	Duration                   uint16 `json:"duration"`
	NeededCorrectors           uint8  `json:"corrector"`
	AllowedCorrectorUsageTimes uint8  `json:"correctorUsageLimit"`
	AllowCorrectTilQuestion    uint8  `json:"allowCorrectTilQuestion"`
	NeededTickets              uint8  `json:"incomingCost"`
}
