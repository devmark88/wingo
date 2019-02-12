package q

import (
	"gitlab.com/mt-api/wingo/broker"
)

// GetTasks : get all tasks
func GetTasks() map[string]interface{} {
	broker := broker.Broker{}
	m := make(map[string]interface{})
	m["publish"] = broker.Publish
	return m
}
