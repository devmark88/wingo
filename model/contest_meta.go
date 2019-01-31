package model

import (
	"time"
)

//MetaContest => Meta Contest Model
type MetaContest struct {
	appID                      string
	title                      string
	prize                      int
	beginTime                  time.Time
	duration                   int
	itemDuration               int
	neededCorrectors           int
	allowedCorrectorUsageTimes int
	allowCorrectTilQuestion    int
	neededTickets              int
}
