package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	const englishHelloPrefix = "Hello, "

	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}
