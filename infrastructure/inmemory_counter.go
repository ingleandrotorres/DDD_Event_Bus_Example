package infrastructure

import (
	"DDDEventBus/domain"
	"fmt"
)

type BestCounter struct {
	counter domain.Counter
}

func (b *BestCounter) Save(counter domain.Counter) error {
	fmt.Printf("El contador va en %d", counter)
	b.counter = counter
	return nil
}
func (b *BestCounter) Get() domain.Counter {
	return b.counter
}

func NewBestCounter() domain.CounterRepository {
	return &BestCounter{}
}
