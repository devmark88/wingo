package q

import (
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"

	"github.com/spf13/viper"

	"gitlab.com/mt-api/wingo/logger"

	"gitlab.com/mt-api/wingo/response"

	"gitlab.com/mt-api/wingo/model"
)

//QueueManager ...
type QueueManager struct{}

// PushQuestions : Publish all questions of a contest
func (q QueueManager) PushQuestions(c *model.Contest, srv *machinery.Server) error {
	pub := Pub{Server: srv}
	if len(c.Questions) == 0 {
		return fmt.Errorf("No question to publish for contest meta: %v", c.Meta.ID)
	}
	bt := c.Meta.BeginTime.UTC()
	now := time.Now().UTC()
	logger.Debug(fmt.Sprintf("%v - %v", bt.String(), now.String()))

	d := bt.Sub(now).Seconds()
	logger.Debug(fmt.Sprintf("DIFF SEC | INT = %v|%v", d, int(d)))
	if d <= 0 {
		return fmt.Errorf("begin time is %s, you can not add question to this contest anymore", c.Meta.BeginTime.String())
	}
	itemDuration := int(c.Meta.Duration / uint(len(c.Questions)))
	answerWaiting := viper.GetInt("app.answer_delay")
	tpc := getQuestionTopic(c.Meta.ID)

	for idx, q := range c.Questions {
		p := response.QuestionPayload{}
		p.ID = q.ID
		p.Body = q.Body
		p.Index = idx
		p.Options = q.Answers
		p.ContestID = c.ID
		p.Type = response.QuestionPayloadEnum
		delay := 0
		if idx == 0 {
			delay = int(d)
		} else {
			delay = int(d) + idx*(itemDuration+answerWaiting)
		}
		logger.Debug(fmt.Sprintf("publishing: index=%v, delay=%v, topic: %s", idx, delay, tpc))
		e := pub.PublishDelayed(tpc, delay, p)
		if e != nil {
			return e
		}
	}
	return nil
}

// PushDeadline : push job to deadline queue
func (q QueueManager) PushDeadline(c *model.Contest, srv *machinery.Server) error {
	pub := Pub{Server: srv}
	bt := c.Meta.BeginTime.UTC()
	now := time.Now().UTC()
	logger.Debug(fmt.Sprintf("%v - %v", bt.String(), now.String()))
	d := bt.Sub(now).Seconds()
	logger.Debug(fmt.Sprintf("DIFF SEC | INT = %v|%v", d, int(d)))
	if d <= 0 {
		return fmt.Errorf("begin time is %s, you can not add question to this contest anymore", c.Meta.BeginTime.String())
	}
	vd := viper.GetInt("app.video_duration")
	tpc := getDeadlineTopic(c.ID)
	p := response.DeadlinePayload{}
	p.ContestID = c.ID
	p.Type = response.DeadlinePayloadEnum
	delay := int(d) - vd
	e := pub.PublishDelayed(tpc, delay, p)
	return e
}
