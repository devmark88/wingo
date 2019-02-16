package model

import (
	"time"
)

// ContestMeta : ContestMetaModel
type ContestMeta struct {
	Base
	AppID                      string
	Title                      string
	Prize                      uint
	BeginTime                  time.Time
	Duration                   uint
	NeededCorrectors           uint
	AllowedCorrectorUsageTimes uint
	AllowCorrectTilQuestion    uint
	NeededTickets              uint
}
