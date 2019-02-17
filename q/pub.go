package q

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"gitlab.com/mt-api/wingo/logger"
)

// Pub : Publish queue server
type Pub struct {
	Server *machinery.Server
}

//PublishDelayed : Send Delayed job task
func (s *Pub) PublishDelayed(topic string, delay int, res interface{}) error {
	var args []tasks.Arg
	topicArg := &tasks.Arg{
		Name:  "t",
		Type:  "string",
		Value: topic,
	}

	ps, _ := json.Marshal(res)

	payloadArg := &tasks.Arg{
		Name:  "payload",
		Type:  "string",
		Value: string(ps),
	}
	args = append(args, *topicArg)
	args = append(args, *payloadArg)

	signature, err := tasks.NewSignature("publish", args)
	eta := time.Now().UTC().Add(time.Second * time.Duration(delay))
	signature.ETA = &eta
	if err != nil {
		return fmt.Errorf("error while creating signature: %v", err)
	}
	logger.Debug(fmt.Sprintf("sending tasks: delay=%v, topic: %s", delay, topic))

	_, err = s.Server.SendTask(signature)
	if err != nil {
		return fmt.Errorf("error while send task signature:%v => %v", signature, err)
	}
	return nil
}

// Publish : Send Job
func (s *Pub) Publish(topic string, res interface{}) error {
	var args []tasks.Arg
	topicArg := &tasks.Arg{
		Name:  "t",
		Type:  "string",
		Value: topic,
	}
	ps, _ := json.Marshal(res)

	payloadArg := &tasks.Arg{
		Name:  "payload",
		Type:  "string",
		Value: string(ps),
	}
	args = append(args, *topicArg)
	args = append(args, *payloadArg)

	signature, err := tasks.NewSignature("publish", args)
	if err != nil {
		return fmt.Errorf("error while creating signature: %v", err)
	}
	logger.Debug(fmt.Sprintf("sending tasks: topic: %s", topic))

	_, err = s.Server.SendTask(signature)
	return err
}
