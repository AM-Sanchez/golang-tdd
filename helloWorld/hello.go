package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloDefaultSuffix = "World!!!"

func Hello(name string) string {
	if name == "" {
		name = englishHelloDefaultSuffix
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Bub"))
}
