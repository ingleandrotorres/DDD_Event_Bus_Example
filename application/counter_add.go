package application

import (
	"DDDEventBus/domain"
)

type CounterAdder struct {
	Repository domain.CounterRepository
}

func (c CounterAdder) Execute() error {
	counter := c.Repository.Get()

	counter.Increment()

	err := c.Repository.Save(counter)
	if err != nil {
		return err
	}

	return nil
}

func NewCounterAdder(repository domain.CounterRepository) CounterAdder {
	return CounterAdder{Repository: repository}
}
