package main

import "fmt"

// interfaces are not generic types
// interfaces are implicit
// interfaces are a contract to help us manage types
// interfaces are a way to define a function that can be called on different types
type bot interface {
	// now every type that has a function called getGreeting() that returns a string is also of type bot
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
