package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("guy", "english"))
}

// Hello says hi
func Hello(name string, language string) string {

	const englishGreeting = "Hello, "
	const englishPronoun = "you"

	const spanishGreeting = "Hola, "
	const spanishPronoun = "tu"

	const frenchGreeting = "Bonjour, "
	const frenchPronoun = "vous"

	const spanish = "spanish"
	const french = "french"

	var greeting string
	var pronoun string

	switch language {
	case spanish:
		greeting = spanishGreeting
		pronoun = spanishPronoun
	case french:
		greeting = frenchGreeting
		pronoun = frenchPronoun
	default:
		greeting = englishGreeting
		pronoun = englishPronoun
	}
	if name == "" {
		return greeting + pronoun
	} else {
		return greeting + name
	}
}
