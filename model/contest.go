package model

import (
	"fmt"
	"time"

	"gitlab.com/mt-api/wingo/logger"
	"gitlab.com/mt-api/wingo/messages"

	"gitlab.com/mt-api/wingo/helpers"

	"github.com/spf13/viper"
)

// Contest : Contest Model
type Contest struct {
	Base
	ContestMetaID         uint64
	Meta                  ContestMeta
	Questions             []Question
	CorrectAnswersIndices string
}

// IsPast : is begin time of contest past?
func (c *Contest) IsPast() bool {
	return helpers.IsPast(c.Meta.BeginTime)
}

// IsQuestionInTime : is valid time for answer this question is past?
func (c *Contest) IsQuestionInTime(id uint) bool {
	d := int(c.Meta.Duration) / len(c.Questions)
	idx := c.GetQuestionIndex(id)
	if idx < 0 {
		return false
	}
	dt := d + viper.GetInt("app.answer_threshould")
	delay := dt
	min := c.Meta.BeginTime
	if idx > 0 {
		delay = dt * idx
	}
	max := c.Meta.BeginTime.Add(time.Second * time.Duration(delay))
	return helpers.DateWithinRange(time.Now(), min, max)
}

// GetQuestionIndex : whats is index of given question id
func (c *Contest) GetQuestionIndex(id uint) int {
	idx := -1
	for i, q := range c.Questions {
		if q.ID == id {
			return i
		}
	}
	return idx
}

// IsItCorrectAnswer : is selectedAnswerIndex is corrected answer
func (c *Contest) IsItCorrectAnswer(selectedAnswerIndex int, qID uint) bool {
	if selectedAnswerIndex < 0 {
		return false
	}
	qidx := c.GetQuestionIndex(qID)
	ca, err := helpers.StringToIntArray(c.CorrectAnswersIndices, ",")
	if err != nil {
		logger.Error(fmt.Sprintf(messages.ErrorInSplitCorrectAnswerIndices, qID, err))
		return false
	}
	if qidx > len(c.Questions)-1 {
		logger.Error(fmt.Sprintf(messages.WrongIndexInSetAnsewr, len(c.Questions)-1, selectedAnswerIndex, c.ID, qID))
		return false
	}
	return selectedAnswerIndex == ca[qidx]
}

// CaneYetUserCorrector : can user use correct in particular question
func (c *Contest) CaneYetUserCorrector(u *UserInfo, ut *UserTrack, qidx uint) bool {
	if u == nil {
		return c.Meta.AllowedCorrectorUsageTimes > 0 && u.Correctors >= c.Meta.NeededCorrectors
	}
	return ut.CanPlay && c.Meta.AllowedCorrectorUsageTimes > ut.CorrectorUsageTimes && c.Meta.AllowCorrectTilQuestion > qidx
}
