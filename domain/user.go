package domain

import (
	"errors"
	"fmt"
	"strings"
)

type Name string

func (name Name) EnsureValid() error {
	if strings.TrimSpace(string(name)) == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

type LastName string

func (lastName LastName) EnsureValid() error {
	if strings.TrimSpace(string(lastName)) == "" {
		return errors.New("last_name cannot be empty")
	}

	return nil
}

type Password string

func (password Password) ensureValid() error {
	if strings.TrimSpace(string(password)) == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

type User struct {
	ID       ID
	Name     Name
	LastName LastName
	Password Password

	DomainEvents []DomainEvent
}

func NewUser(id ID, name Name, lastName LastName, password Password) (User, error) {

	err := id.EnsureValid()
	if err != nil {
		return User{}, err
	}

	err = name.EnsureValid()
	if err != nil {
		return User{}, err
	}

	err = lastName.EnsureValid()
	if err != nil {
		return User{}, err
	}

	return User{
		ID:           id,
		Name:         name,
		LastName:     lastName,
		Password:     password,
		DomainEvents: []DomainEvent{UserCreated{}},
	}, nil

}

func (u User) Login(password Password) error {
	if string(u.Password) == strings.TrimSpace(string(password)) {
		return nil
	}

	return errors.New("incorrect password")
}

func (u User) ToString() string {
	return fmt.Sprintf("ID: %s, NAME: %s, LASTNAME: %s, PASSWORD: %s", u.ID, u.Name, u.LastName, u.Password)
}

func (u *User) Poll() []DomainEvent {
	events := make([]DomainEvent, 0)

	for _, e := range u.DomainEvents {
		events = append(events, e)
	}
	u.DomainEvents = make([]DomainEvent, 0)

	return events
}

type DomainEvent interface {
	GetEventInfo() string
}
type UserCreated struct {
}

func (u UserCreated) GetEventInfo() string {
	return "user_created"
}

type EventBus interface {
	Publish(events []DomainEvent)
	Subscribe(topic string, fn interface{}) error
}
