package q

import (
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"

	"github.com/spf13/viper"

	"gitlab.com/mt-api/wingo/broker"

	"gitlab.com/mt-api/wingo/logger"

	"gitlab.com/mt-api/wingo/response"

	"gitlab.com/mt-api/wingo/helpers"

	"gitlab.com/mt-api/wingo/model"
)

type Question struct{}

func (q Question) PublishAll2(c model.Contest, srv *machinery.Server) error {
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
	itemDuration := int(c.Meta.Duration / uint16(len(c.Questions)))
	answerWaiting := viper.GetInt("app.answer_delay")
	tpc := fmt.Sprintf("contest%v", c.Meta.ID)

	for idx, q := range c.Questions {
		p := response.QuestionPayload{}
		p.ID = q.ID
		p.Body = q.Body
		p.Index = idx
		p.Options = q.Answers
		delay := 0
		if idx == 0 {
			delay = int(d)
		} else {
			delay = int(d) + idx*(itemDuration+answerWaiting)
		}
		logger.Debug(fmt.Sprintf("publishing: index=%v, delay=%v, topic: %s", idx, delay, tpc))
		e := pub.PublishQuestion(tpc, delay, p)
		if e != nil {
			return e
		}
	}
	return nil
}
func (q Question) PublishAll(c model.Contest) error {
	broker := broker.Broker{}
	if len(c.Questions) == 0 {
		return fmt.Errorf("No question to publish for contest meta: %v", c.Meta.ID)
	}
	bt := helpers.TimeInTehran(c.Meta.BeginTime)
	now := helpers.TimeInTehran(time.Now())

	logger.Debug("Defining Messages")
	logger.Debug(fmt.Sprintf("Begin Time = %v", c.Meta.BeginTime.String()))

	logger.Debug(fmt.Sprintf("%v - %v", bt.String(), now.String()))
	logger.Debug(fmt.Sprintf("SUB = " + bt.Sub(now).String()))
	d := bt.Sub(now).Seconds()
	logger.Debug(fmt.Sprintf("DIFF SEC | INT = %v|%v", d, int(d)))
	if d <= 0 {
		return fmt.Errorf("begin time is %s, you can not add question to this contest anymore", c.Meta.BeginTime.String())
	}
	itemDuration := int(c.Meta.Duration / uint16(len(c.Questions)))
	answerWaiting := viper.GetInt("app.answer_delay")
	logger.Debug(fmt.Sprintf("total duration: %v, question count: %v, item duration: %v, answer waiting: %v", c.Meta.Duration, len(c.Questions), itemDuration, answerWaiting))
	for idx, q := range c.Questions {
		p := response.QuestionPayload{}
		p.ID = q.ID
		p.Body = q.Body
		p.Index = idx
		p.Options = q.Answers

		delay := 0
		if idx == 0 {
			delay = int(d)
		} else {
			delay = int(d) + idx*(itemDuration+answerWaiting)
		}
		tpc := fmt.Sprintf("contest%v", c.Meta.ID)
		logger.Debug(fmt.Sprintf("index=%v, delay=%v, topic: %s", idx, delay, tpc))
		r, e := broker.DelayPublish(tpc, delay, p)
		if e != nil {
			return fmt.Errorf("error while trying to define delayed queue contestMeta:%v, questionIndex:%v: %v", c.Meta.ID, q.ID, e)
		} else {
			logger.Debug(fmt.Sprintf("PublishAll: EMQ delayed publish response: %v", r))
		}
	}
	return nil
}
