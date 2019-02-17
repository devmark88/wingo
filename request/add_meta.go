package request

import (
	"time"

	"gitlab.com/mt-api/wingo/model"
)

// AddMetaRequest : request model for add meta request
type AddMetaRequest struct {
	AppID                      string    `json:"app" xml:"app" binding:"required"`
	Title                      string    `json:"title" xml:"title" binding:"required"`
	Prize                      uint      `json:"prize" xml:"prize" binding:"required"`
	BeginTime                  time.Time `json:"beginDateTime" xml:"beginDateTime" binding:"required"`
	Duration                   uint      `json:"duration" xml:"duration" binding:"required"`
	NeededCorrectors           uint      `json:"corrector" xml:"corrector" binding:"required"`
	AllowedCorrectorUsageTimes uint      `json:"correctorUsageLimit" xml:"correctorUsageLimit" binding:"required"`
	AllowCorrectTilQuestion    uint      `json:"allowCorrectTilQuestion" xml:"allowCorrectTilQuestion" binding:"required"`
	NeededTickets              uint      `json:"incomingCost" xml:"incomingCost" binding:"required"`
}

// ToModel : map requests dto to the model
func (r *AddMetaRequest) ToModel() *model.ContestMeta {
	m := model.ContestMeta{}
	m.AppID = r.AppID
	m.Title = r.Title
	m.Prize = r.Prize
	m.BeginTime = r.BeginTime
	m.Duration = r.Duration
	m.NeededCorrectors = r.NeededCorrectors
	m.AllowedCorrectorUsageTimes = r.AllowedCorrectorUsageTimes
	m.AllowCorrectTilQuestion = r.AllowCorrectTilQuestion
	m.NeededTickets = r.NeededTickets
	return &m
}
