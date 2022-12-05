package domain

import "fmt"

type Counter struct {
	Quantity int
}

func (c *Counter) Increment() {
	c.Quantity += 1
	fmt.Printf("Se ha incrementado en una unidad el counter")
}

type CounterRepository interface {
	Save(counter Counter) error
	Get() Counter
}

type CounterAdded struct {
}

func (u CounterAdded) GetEventInfo() interface{} {
	return nil
}
