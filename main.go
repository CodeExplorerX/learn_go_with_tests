package main

import "fmt"

const englishHelloPreix = "Hello, "

func Hello(name string) string {
	return englishHelloPreix + name
}

func main() {
	fmt.Println(Hello("world"))
}
