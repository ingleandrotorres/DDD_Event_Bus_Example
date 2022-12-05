package main

import (
	"DDDEventBus/application"
	"DDDEventBus/domain"
	"DDDEventBus/infrastructure"
	"fmt"
	"github.com/asaskevich/EventBus"
)

func ExampleDDDEventBus() {
	eventBus := EventBus.New()

	counterRepo := infrastructure.NewBestCounter()
	counterAdder := application.NewCounterAdder(counterRepo)
	controllerEventBus := infrastructure.NewInMemoryControllerEventBus(topicUserCreated, eventBus, counterAdder)

	go controllerEventBus.Subscribe()
	fmt.Println("El counter se  suscribió a Crear Usuario")

	id := domain.ID("29070fce-62a8-4655-b849-55be86e30692")
	name := domain.Name("Juan")
	lastName := domain.LastName("Perez")
	password := domain.Password("123")

	inMemoryEventBus := infrastructure.NewInMemoryEventBus(topicUserCreated, eventBus)
	usercreator := application.NewUserCreator(inMemoryEventBus)
	_, err := usercreator.Execute(id, name, lastName, password)
	if err != nil {
		fmt.Println("No se puso ejecutar el caso de uso")
		return
	}
}

const topicUserCreated = "top_user_created"
