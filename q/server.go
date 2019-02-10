package q

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func Start() *machinery.Server {
	var cnf = &config.Config{
		Broker:        "redis://localhost:6379",
		DefaultQueue:  "machinery_tasks",
		ResultBackend: "amqp://guest:guest@localhost:5672/",
		AMQP: &config.AMQPConfig{
			Exchange:     "machinery_exchange",
			ExchangeType: "direct",
			BindingKey:   "machinery_task",
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		// do something with the error
	}
	return server
}
