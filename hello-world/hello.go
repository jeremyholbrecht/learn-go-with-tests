package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const dutch = "dutch"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const dutchHelloPrefix = "Goeiendag, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case dutch:
		prefix = dutchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Jere", dutch))
}
