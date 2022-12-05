package infrastructure

import (
	"DDDEventBus/domain"
	"github.com/asaskevich/EventBus"
)

type InMemoryEventBus struct {
	topic string
	bus   EventBus.Bus
}

func (b InMemoryEventBus) Publish(events []domain.DomainEvent) {
	//for _, event := range events {
	b.bus.Publish(b.topic)
	//}
}

func NewInMemoryEventBus(topic string, bus EventBus.Bus) domain.EventBus {
	return InMemoryEventBus{
		topic: topic,
		bus:   bus,
	}
}

// TODO: segregar las interfaces
func (b InMemoryEventBus) Subscribe(topic string, fn interface{}) error {
	err := b.bus.Subscribe(topic, fn)
	if err != nil {
		return err
	}
	return nil
}
