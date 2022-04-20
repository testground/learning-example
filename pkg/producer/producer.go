package producer

// A producer creates new DataMessages and pushes them to a specified queue, channel, etc. where
// the messages will be received and handled by listeners and consumers
type Producer interface {
	// Produces a new message from the given string data, and pushes it to queue
	ProduceMessage(string)
}
