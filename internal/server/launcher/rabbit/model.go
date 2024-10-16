package rabbit

import "Lists-app/internal/broker/rabbit/config"

type consumer struct {
	topic string
}

type Model struct {
	consumers []consumer
}

func NewModel() Model {
	model := Model{}

	model.consumers = append(model.consumers, consumer{topic: config.QueueName})

	return model
}
