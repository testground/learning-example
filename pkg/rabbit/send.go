package rabbit

import (
	"github.com/testground/learning-example/pkg/util"

	"github.com/streadway/amqp"
)

func PublishMessage(queueName string, data []byte) {
	conn := GetConnection()
	AssertQueue(conn, queueName)

	ch, err := conn.Channel()
	util.FailOnError(err, "Error opening rabbitMq channel")

	defer ch.Close()

	err = ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)

	util.FailOnError(err, "Error publishing message")
}
