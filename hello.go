package main

import "fmt"

const yoruba = "yoruba"
const french = "french"
const englishHelloPrefix = "Hello, "
const yorubaHelloPrefix = "Pele o, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

// In Go, public functions start with a capital letter and private 
// ones start with a lowercase

func greetingPrefix(lang string) (prefix string) { // "prefix" is a named return value
	switch lang {
	case yoruba:
        prefix = yorubaHelloPrefix
	case french:
        prefix = frenchHelloPrefix
    default:
        prefix = englishHelloPrefix
	}
    return
}

func main() {
	fmt.Println(Hello("world", ""))
}
