package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const russianHelloPrefix = "Привет, "

func Hello(name string, language string) string {
	if name == "" {
		name = "bro"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "russian": 
		prefix = russianHelloPrefix
	case "french": 	
		prefix = frenchHelloPrefix
	default: 		
		prefix = englishHelloPrefix
	}
	
	return prefix
}

func main() {
	fmt.Println(Hello("Sasisa", "english"))
}