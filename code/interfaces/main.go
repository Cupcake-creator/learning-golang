package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type thaiBot struct{}

func (englishBot) getGreeting() string {
	return "Hello, english!"
}

func (thaiBot) getGreeting() string {
	return "Hello, thai!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	// 	bot1 := englishBot{}
	// 	bot2 := thaiBot{}
	// fmt.Println(bot1.getGreeting())
	// fmt.Println(bot2.getGreeting())
	// printGreeting(bot1)
	// printGreeting(bot2)

	shape1 := triangle{base: 4, height: 4}
	shape2 := square{sideLength: 4}

	printArea(shape1)
	printArea(shape2)
}
