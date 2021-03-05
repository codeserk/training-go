package hello

import (
	"fmt"

	"codeserk.es/greetings"
)

func main() {
	message := greetings.SayHello("me")
	fmt.Print(message)
}
