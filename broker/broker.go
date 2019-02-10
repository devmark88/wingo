package broker

import (
	"net/http"

	"gitlab.com/mt-api/wingo/broker/emq"
)

type MessageBroker interface {
	DelayPublish(t string, d int, payload interface{}) (interface{}, error)
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
