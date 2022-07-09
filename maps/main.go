package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#FFF",
	}

	// colors := make(map[string]string)

	// colors["a"] = "a"

	// fmt.Println(colors)

	// delete(colors, "a")

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Key: %v. Value: %v", color, hex)
		fmt.Println()
	}
}
