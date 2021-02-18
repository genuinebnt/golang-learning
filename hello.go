package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//Hello world function
func Hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}
