package main

import "fmt"

type chatBot interface {
	getGreeting(string) string
}

type englishBot struct{}
type frenchBot struct{}

func main() {
	firstname := "Fred"

	eb := englishBot{}
	printGreeting(eb, firstname)

	fb := frenchBot{}
	printGreeting(fb, firstname)

}

func (eb englishBot) getGreeting(firstname string) string {
	return "Hello " + firstname
}

func (fb frenchBot) getGreeting(firstname string) string {
	return "Bonjour " + firstname
}

func printGreeting(b chatBot, firstname string) {
	fmt.Println(b.getGreeting(firstname))
}
