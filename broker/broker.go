package broker

import (
	"gitlab.com/mt-api/wingo/broker/emq"
)

type MessageBroker interface {
	DelayPublish(t string, d int, payload interface{}) error
}
type Broker struct{}

func (b Broker) DelayPublish(t string, d int, payload interface{}) error {
	e := emq.Broker{}
	return e.DelayPublish(t, d, payload)
}
