package emq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/mt-api/wingo/logger"

	"github.com/spf13/viper"
)

// Broker : Send messages to the broker
type Broker struct {
	Client *http.Client
}

const pulishEndpoint = "/api/v3/mqtt/publish"

// DelayPublish : send delayed message with $delayed tag
func (b Broker) DelayPublish(t string, d int, payload interface{}) (*http.Response, error) {
	dtn := fmt.Sprintf("$delayed/%v/%s", d, t)
	logger.Debug("EMQ: delayed topic => " + dtn)
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("EMQ:PUBLISH => error while parsing payload: %v", err)
	}
	m := PublishModel{Topic: dtn, ClientID: "test_client_id", Payload: string(p), QOS: 1, Retain: false}

	pe := fmt.Sprintf("%s%s", viper.GetString("emq.base"), pulishEndpoint)
	bytesRepresentation, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("EMQ:PUBLISH => error while parsing publish model: %v", err)
	}
	req, _ := http.NewRequest("POST", pe, bytes.NewBuffer(bytesRepresentation))
	req.SetBasicAuth(viper.GetString("emq.auth.username"), viper.GetString("emq.auth.password"))
	return b.Client.Do(req)
}

// Publish : Publish message to the emq by HTTP endpoint
func (b Broker) Publish(t string, payload string) (*http.Response, error) {
	m := PublishModel{Topic: t, ClientID: "test_client_id", Payload: payload, QOS: 1, Retain: false}

	pe := fmt.Sprintf("%s%s", viper.GetString("emq.base"), pulishEndpoint)
	bytesRepresentation, err := json.Marshal(m)
	logger.Debug(string(bytesRepresentation))
	if err != nil {
		return nil, fmt.Errorf("EMQ:PUBLISH => error while parsing publish model: %v", err)
	}
	req, _ := http.NewRequest("POST", pe, bytes.NewBuffer(bytesRepresentation))
	req.SetBasicAuth(viper.GetString("emq.auth.username"), viper.GetString("emq.auth.password"))
	res, e := b.Client.Do(req)
	if e != nil {
		logger.Error(fmt.Errorf("error while calling EMQx: %v", e))
		return nil, e
	}
	logger.Debug(fmt.Sprintf("STATUS : %v", res.Status))
	return res, nil
}
