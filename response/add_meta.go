package response

// AddMeta : AddMeta response
type AddMeta struct {
	ID                         uint   `json:"id"`
	AppID                      string `json:"app"`
	Title                      string `json:"title"`
	Prize                      uint   `json:"prize"`
	BeginTime                  string `json:"beginDateTime"`
	Duration                   uint   `json:"duration"`
	NeededCorrectors           uint   `json:"corrector"`
	AllowedCorrectorUsageTimes uint   `json:"correctorUsageLimit"`
	AllowCorrectTilQuestion    uint   `json:"allowCorrectTilQuestion"`
	NeededTickets              uint   `json:"incomingCost"`
}
