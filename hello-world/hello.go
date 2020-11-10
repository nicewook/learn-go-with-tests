package main

import "fmt"

const (
	English        = "English"
	Spanish        = "Spanish"
	French         = "French"
	engHelloPrefix = "Hello, "
	spaHelloPrefix = "Hola, "
	freHelloPrefix = "Hola, "
)

func prefix(language string) string {
	var prefix string
	switch language {
	case English:
		prefix = engHelloPrefix
	case Spanish:
		prefix = spaHelloPrefix
	case French:
		prefix = freHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return prefix
}

// Hello better out from main(), so easy to test
func Hello(lang, name string) string {
	if name == "" {
		name = "world"
	}

	return prefix(lang) + name
}

func main() {
	fmt.Println(Hello(English, "world"))
}
