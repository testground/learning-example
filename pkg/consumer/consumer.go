package consumer

import "github.com/testground/learning-example/pkg/message"

// A consumer consumes given data messages. What this means will vary:
// a consumer could discard a message, write the data to a database, etc.
type Consumer interface {
	ConsumeMessage(*message.DataMessage)
}
