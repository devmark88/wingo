package model

import (
	"database/sql"
)

// UserTrack : UserTrackModel
type UserTrack struct {
	Base
	ContestID             uint
	UserID                string
	QuestionID            uint
	QuestionIndex         int
	CanPlay               bool
	CanUseCorrector       bool
	IsSelectCorrectAnswer bool
	CorrectorUsageTimes   int
	State                 TrackStateEnum
	BeforeTickets         uint
	BeforeCorrectors      uint
	AfterTickets          uint
	AfterCorrectors       uint
	MetaData              sql.NullString
}
