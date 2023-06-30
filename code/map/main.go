package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)

	// Like Dict in Python
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Update Map
	colors["orange"] = "#FFA500"
	fmt.Println(colors)

	// Delete colors white key
	delete(colors, "orange")
	fmt.Println(colors)

	printMap(colors)
	fmt.Println(colors)

}

func printMap(m map[string]string) {
	m["orange"] = "#FFA500"
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
