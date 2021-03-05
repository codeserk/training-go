package greetings

import "fmt"

// SayHello creates a message that greets the person with the given name.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v, how are you?", name)
}
