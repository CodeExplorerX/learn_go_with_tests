package main

import "fmt"

const englishHelloPreix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPreix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
