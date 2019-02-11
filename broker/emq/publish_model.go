package emq

type PublishModel struct {
	Topic string `json:"topic"`
	Payload string `json:"payload"`
	QOS int `json:"qos"`
	Retain bool `json:"retain"`
	ClientID string `json:"client_id"`
}
