package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("guy", "english"))
}

// Hello says hi
func Hello(name string, language string) string {

	var greeting string

	switch language {
	case "spanish":
		greeting = "Hola, "
	default:
		greeting = "Hello, "
	}
	if name == "" {
		return greeting + "you"
	} else {
		return greeting + name
	}
}
