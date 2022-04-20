package rabbit

import (
	"log"
	"time"

	env "github.com/testground/learning-example/pkg/env"
	"github.com/testground/learning-example/pkg/util"

	"github.com/streadway/amqp"
)

var rabbitConnection *amqp.Connection

// Returns the connection to RabbitMq. If the connection is not established, will attempt to
// establish it, with a specified timeout period
func GetConnection() *amqp.Connection {
	if rabbitConnection == nil {
		var err error
		for i := 10; i > 0; i-- {
			rabbitConnection, err = amqp.Dial(env.GetEnv().RabbitBrokerUrl)
			if err == nil {
				break
			} else {
				time.Sleep(time.Second * 2)
			}
		}

		// if err still exists, then we failed to connect to rabbit
		util.FailOnError(err, "Error connecting to rabbitMq")
		log.Println("Connected to rabbitMQ")
	}
	return rabbitConnection
}
