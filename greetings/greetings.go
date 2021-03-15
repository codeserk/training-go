package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var messages = [3]string{
	"Hello %v, how are you?",
	"Nice to meet you %v",
	"Is everything okay, %v?",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("really?")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(getRandomMessage(), name)

	return message, nil
}

// HelloMultiple returns a greeting for the named person.
func HelloMultiple(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		if name == "" {
			return nil, errors.New("really?")
		}

		messages[name] = fmt.Sprintf(getRandomMessage(), name)
	}

	return messages, nil
}

func getRandomMessage() string {
	return messages[rand.Intn(len(messages))]
}
