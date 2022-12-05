package main

import (
	"bufio"
	"fmt"
	"github.com/asaskevich/EventBus"
	"os"
	"strings"
	"time"
)

const createUserTopic = "Crear Usuario"
const logUserTopic = "Loguear Usuario"
const giveTokensToUserTopic = "Dar tokens a un usuario"

func CreateUser(s string) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("CrearUsuario done %s \n", s)
}

func LogUser(s string) {
	time.Sleep(1500 * time.Millisecond)
	fmt.Printf("LogUser done %s", s)
}

func GiveTokensToUser(id, numberOfTokens string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("GiveTokensToUser done " + id + " " + numberOfTokens + "\n")
}

func main2() {
	/*
		bus := EventBus.New() //TODO cambiar libreria reactiveX para golang

		_ = bus.Subscribe(createUserTopic, CreateUser)
		_ = bus.Subscribe(logUserTopic, LogUser)
		_ = bus.Subscribe(giveTokensToUserTopic, GiveTokensToUser)

		go func() {
			bus.Publish(logUserTopic, "loguear1")
			bus.Publish(logUserTopic, "loguear2")
			bus.Publish(logUserTopic, "loguear3")
		}()

		go func() {
			bus.Publish(giveTokensToUserTopic, "1", "10")
			bus.Publish(giveTokensToUserTopic, "2", "20")
			bus.Publish(giveTokensToUserTopic, "3", "30")
		}()

		bus.Publish(createUserTopic, "createuser1")
		bus.Publish(createUserTopic, "createuser2")
		bus.Publish(createUserTopic, "createuser3")

		time.Sleep(4000 * time.Millisecond)

		for i := 0; i < 5; i++ {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("\nSuscribe : create :1 ,loguser :2, give tokens :3\n")
			text, _ := reader.ReadString('\n')
			fmt.Println("you chose" + text)
			text = strings.Trim(text, "\n")

			if text == "1" {
				go func() {
					bus.Publish(createUserTopic, "createuserconsole")
				}()
				go func() {
					bus.Publish(giveTokensToUserTopic, "3", "30")
				}()

			}
			if text == "2" {
				bus.Publish(logUserTopic, "loguear console")
			}
			if text == "3" {
				bus.Publish(giveTokensToUserTopic, "1", "400")
			}
			if text == "4" {
				ExampleDDDEventBus()
			}
		}
	*/
	ExampleDDDEventBus()
	/*bus := EventBus.New()
	_ = bus.Subscribe(createUserTopic, CreateUser)
	bus.Publish(giveTokensToUserTopic)*/
}
func main() {

	bus := EventBus.New() //TODO cambiar libreria reactiveX para golang

	_ = bus.Subscribe(createUserTopic, CreateUser)
	_ = bus.Subscribe(logUserTopic, LogUser)
	_ = bus.Subscribe(giveTokensToUserTopic, GiveTokensToUser)

	go func() {
		bus.Publish(logUserTopic, "loguear1")
		bus.Publish(logUserTopic, "loguear2")
		bus.Publish(logUserTopic, "loguear3")
	}()

	go func() {
		bus.Publish(giveTokensToUserTopic, "1", "10")
		bus.Publish(giveTokensToUserTopic, "2", "20")
		bus.Publish(giveTokensToUserTopic, "3", "30")
	}()

	bus.Publish(createUserTopic, "createuser1")
	bus.Publish(createUserTopic, "createuser2")
	bus.Publish(createUserTopic, "createuser3")

	time.Sleep(4000 * time.Millisecond)

	for i := 0; i < 5; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nSuscribe : create :1 ,loguser :2, give tokens :3\n")
		text, _ := reader.ReadString('\n')
		fmt.Println("you chose" + text)
		text = strings.Trim(text, "\n")

		if text == "1" {
			go func() {
				bus.Publish(createUserTopic, "createuserconsole")
			}()
			go func() {
				bus.Publish(giveTokensToUserTopic, "3", "30")
			}()

		}
		if text == "2" {
			bus.Publish(logUserTopic, "loguear console")
		}
		if text == "3" {
			bus.Publish(giveTokensToUserTopic, "1", "400")
		}
		if text == "4" {
			ExampleDDDEventBus()
		}
	}

}
