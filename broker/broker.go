package broker

import (
	"net/http"

	"gitlab.com/mt-api/wingo/broker/emq"
)

type MessageBroker interface {
	DelayPublish(t string, d int, payload interface{}) (interface{}, error)
	Publish(t string, payload string) (interface{}, error)
}
type Broker struct{}

func (b Broker) DelayPublish(t string, d int, payload interface{}) (interface{}, error) {
	client := http.Client{}
	e := emq.Broker{Client: &client}
	res, err := e.DelayPublish(t, d, payload)
	if err != nil {
		return nil, err
	}
	return res.StatusCode, nil
}
func (b Broker) Publish(t string, payload string) (interface{}, error) {
	client := http.Client{}
	e := emq.Broker{Client: &client}
	res, err := e.Publish(t, payload)
	if err != nil {
		return nil, err
	}
	return res.StatusCode, nil
}
