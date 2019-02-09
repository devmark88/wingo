package broker

type MessageBroker interface {
	DelayPublish(t string, d int)
}
