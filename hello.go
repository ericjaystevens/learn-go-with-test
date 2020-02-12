package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("guy"))
}

// Hello says hi
func Hello(name string) string {
	if name == "" {
		return "hi, you!"
	} else {
		return "hello, " + name
	}
}
