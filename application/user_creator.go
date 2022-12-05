package application

import (
	"DDDEventBus/domain"
	"fmt"
)

type UserCreator struct {
	EventBus domain.EventBus
}

func (c UserCreator) Execute(id domain.ID, name domain.Name, lastName domain.LastName, password domain.Password) (domain.User, error) {
	user, err := domain.NewUser(id, name, lastName, password)
	if err != nil {
		return domain.User{}, err
	}
	fmt.Printf("\n El usuario se creó  %s", user.ToString())
	c.EventBus.Publish(user.Poll())
	fmt.Printf("\n El vento se publicó")

	return user, nil
}

func NewUserCreator(eventBus domain.EventBus) UserCreator {
	return UserCreator{EventBus: eventBus}
}
