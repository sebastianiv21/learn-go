package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	delete(colors, "green")

	for color, hex := range colors {
		fmt.Println("The color is", color, "and the hex is", hex)
	}

	fmt.Println(colors)
}
