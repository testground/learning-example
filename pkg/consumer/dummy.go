package consumer

import (
	"log"

	"github.com/testground/learning-example/pkg/message"
)

// Simple consumer that prints out received data messages
type DummyConsumer struct {
	QueueName string
}

func (consumer *DummyConsumer) ConsumeMessage(msg *message.DataMessage) {
	log.Println("Consumed message #", (*msg).Id, ". Data: ", (*msg).Data)
}
