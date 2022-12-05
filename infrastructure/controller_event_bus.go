package infrastructure

import (
	"DDDEventBus/application"
	"fmt"
	"github.com/asaskevich/EventBus"
)

type InMemoryControllerEventBus struct {
	topic        string
	bus          EventBus.Bus
	counterAdder application.CounterAdder
}

func (b *InMemoryControllerEventBus) Subscribe() error {
	err := b.bus.Subscribe(b.topic, b.callback)
	if err != nil {
		return err
	}

	return nil
}

func (b *InMemoryControllerEventBus) callback() {

	err := b.counterAdder.Execute()
	if err != nil {
		fmt.Printf("hubo un error incrementando el contador")
	}
}

func NewInMemoryControllerEventBus(topic string, bus EventBus.Bus, counterAdder application.CounterAdder) InMemoryControllerEventBus {
	return InMemoryControllerEventBus{
		topic:        topic,
		bus:          bus,
		counterAdder: counterAdder,
	}
}

// TODO: segregar las interfaces
func (b *InMemoryControllerEventBus) Publish() {
	b.bus.Publish(b.topic)
}
