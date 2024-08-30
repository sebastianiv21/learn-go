package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Printf("Enter a string: ")
	fmt.Scan(&input)

	lower := strings.ToLower(input)

	if lower[0] == 'i' && lower[len(lower)-1] == 'n' && strings.Contains(lower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
