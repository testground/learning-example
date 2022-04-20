package rabbit

import (
	"github.com/testground/learning-example/pkg/util"

	"github.com/streadway/amqp"
)

func GetQueueChannel(queueName string) <-chan amqp.Delivery {
	conn := GetConnection()
	AssertQueue(conn, queueName)

	ch, err := conn.Channel()
	util.FailOnError(err, "Error opening rabbitMq channel")

	deliveryChan, err := ch.Consume(
		queueName, // queue
		"",        // zconsumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	return deliveryChan
}
