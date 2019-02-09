package emq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type Broker struct {
	Client *http.Client
}

func (b Broker) DelayPublish(t string, d int, payload interface{}) (*http.Response, error) {
	dtn := fmt.Sprintf("$delayed/%v/%s", d, t)
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("EMQ:PUBLISH => error while parsing payload: %v", err)
	}
	m := PublishModel{Topic: dtn, ClientID: "test_client_id", Payload: string(p), QOS: 2, Retain: false}
	pe := "http://localhost:8080/api/v3/mqtt/publish"
	bytesRepresentation, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("EMQ:PUBLISH => error while parsing publish model: %v", err)
	}
	req, _ := http.NewRequest("POST", pe, bytes.NewBuffer(bytesRepresentation))
	req.SetBasicAuth(viper.GetString("emq.auth.username"), viper.GetString("emq.auth.password"))
	cl := http.Client{}
	return cl.Do(req)
}
