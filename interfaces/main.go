package main

import "fmt"

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

type Bot interface {
	getGretting() string
}

type EnglishBot struct{}

type SpanishBot struct{}

func (eb EnglishBot) getGretting() string {
	return "Hello there!"
}

func (SpanishBot) getGretting() string {
	return "Hola!"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGretting())
}
