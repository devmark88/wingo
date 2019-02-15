package model

import (
	"time"

	"github.com/spf13/viper"
	"gitlab.com/mt-api/wingo/helpers"
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
