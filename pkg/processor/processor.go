package processor

import (
	"github.com/testground/learning-example/pkg/consumer"
	"github.com/testground/learning-example/pkg/listener"
	"github.com/testground/learning-example/pkg/message"
	"github.com/testground/learning-example/pkg/producer"
)

// A processor encapsulates a producer, listener, and consumer, and serves to control
// the creation, movement, and handling of data messages
type Processor struct {
	Producer producer.Producer
	Listener listener.Listener
	Consumer consumer.Consumer
}

func (proc *Processor) StartProcessor() {
	listenerChan := make(chan *message.DataMessage)

	forever := make(chan bool)

	proc.Listener.SetNotifyChannel(listenerChan)

	go func() {
		for msg := range listenerChan {
			if msg != nil {
				proc.Consumer.ConsumeMessage(msg)
			}
		}
	}()

	go func() {
		proc.Listener.ListenForMessages()
	}()

	<-forever
}
