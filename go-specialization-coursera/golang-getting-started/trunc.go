package main

import "fmt"

func main() {
	var number float32

	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&number)

	fmt.Printf("Truncated number: %d\n", int32(number))
}
