package q

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func Start(n int) (*machinery.Server, *[]machinery.Worker, error) {
	if n == 0 {
		n = 1
	}
	var cnf = &config.Config{
		Broker:        "amqp://guest:guest@localhost:5672/",
		DefaultQueue:  "machinery_tasks",
		ResultBackend: "redis://localhost:6379",
		AMQP: &config.AMQPConfig{
			Exchange:     "machinery_exchange",
			ExchangeType: "direct",
			BindingKey:   "machinery_task",
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, nil, err
	}
	var workers []machinery.Worker
	for i := 0; i < n; i++ {
		worker := server.NewWorker(fmt.Sprintf("worker_%v", i), 10)
		go worker.Launch()

		workers = append(workers, *worker)
	}
	registerTasks(server)
	return server, &workers, nil
}

func registerTasks(s *machinery.Server) {
	s.RegisterTasks(GetTasks())
}
