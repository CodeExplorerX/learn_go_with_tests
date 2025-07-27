package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPreix  = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPreifx(language) + name
}

func greetingPreifx(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPreix
	}

	return
}

func main() {
	fmt.Println(Hello("World", "English"))
}
