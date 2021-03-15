package main

import (
	"fmt"
	"log"

	"codeserk.es/greetings"
)

func main() {
	// Get a greeting message and print it.
	names := []string{"Coosto", "Me", "Other"}
	messages, err := greetings.HelloMultiple(names)
	if err != nil {
		log.Fatal("Invalid message: ", err)
	}

	fmt.Println(messages)
}

func init() {
	log.SetPrefix("Coosto - training> ")
	log.SetFlags(0)
}
