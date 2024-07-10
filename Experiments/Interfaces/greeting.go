package main

import "fmt"

type Bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string{
	return "Hello buddy"
}

func (eb spanishBot) getGreeting() string{
	return "Hola"
}

func printGreeting(b Bot){
	fmt.Println(b.getGreeting())
}

func main() {
	var bot Bot
	bot = englishBot{}
	printGreeting(bot)

	sb := spanishBot{}
	printGreeting(sb)
}