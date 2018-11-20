package main

import "fmt"

// All different ways to make a map.
// Map is like an object from JS
func main() {
	// All different ways to make a map below
	// --------------------------------------
	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red": "#ff0000"
	// }

	// How to add things to a Map
	colors["white"] = "#ffffff"

	// Syntax for map differes from struct because in the structs we have . notation syntax but in maps we use bracket notation

	// How to delete things from a map
	// delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hexcode for", key, "is", value)
	}

}
