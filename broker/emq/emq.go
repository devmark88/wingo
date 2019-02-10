package emq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/mt-api/wingo/logger"

	"github.com/spf13/viper"
)

type Broker struct {
	Client *http.Client
}

const PUBLISH_ENDPOINT = "/api/v3/mqtt/publish"

func (b Broker) DelayPublish(t string, d int, payload interface{}) (*http.Response, error) {
	dtn := fmt.Sprintf("$delayed/%v/%s", d, t)
	logger.Debug("EMQ: delayed topic => " + dtn)
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("EMQ:PUBLISH => error while parsing payload: %v", err)
	}
	m := PublishModel{Topic: dtn, ClientID: "test_client_id", Payload: string(p), QOS: 1, Retain: false}

	pe := fmt.Sprintf("%s%s", viper.GetString("emq.base"), PUBLISH_ENDPOINT)
	bytesRepresentation, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("EMQ:PUBLISH => error while parsing publish model: %v", err)
	}
	req, _ := http.NewRequest("POST", pe, bytes.NewBuffer(bytesRepresentation))
	req.SetBasicAuth(viper.GetString("emq.auth.username"), viper.GetString("emq.auth.password"))
	return b.Client.Do(req)
}
