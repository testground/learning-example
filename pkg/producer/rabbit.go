package producer

import (
	"encoding/json"
	"fmt"

	"github.com/testground/learning-example/pkg/message"
	"github.com/testground/learning-example/pkg/rabbit"
	"github.com/testground/learning-example/pkg/util"
)

// Simple producer that produces JSON encoded messages on a specified RabbitMq queue
type RabbitProducer struct {
	IdGen     int32
	QueueName string
}

func (prod *RabbitProducer) ProduceMessage(data string) {
	prod.IdGen++
	msg := &message.DataMessage{
		Id:   fmt.Sprint(prod.IdGen),
		Data: data,
	}

	jsonData, err := json.Marshal(msg)
	util.FailOnError(err, "Error producing message")

	rabbit.PublishMessage(prod.QueueName, jsonData)
}
