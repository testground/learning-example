package listener

import "github.com/testground/learning-example/pkg/message"

// A listener listens for incoming data messages, and sends them to a given
// channel to be processed by consumers
type Listener interface {
	ListenForMessages()
	SetNotifyChannel(chan<- *message.DataMessage)
}
