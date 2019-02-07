package model

import "time"

type ContestMeta struct {
	Base
	AppID                      string
	Title                      string
	Prize                      uint
	BeginTime                  time.Time
	Duration                   uint16
	NeededCorrectors           uint8
	AllowedCorrectorUsageTimes uint8
	AllowCorrectTilQuestion    uint8
	NeededTickets              uint8
}
