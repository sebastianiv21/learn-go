package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	intSlice := make([]int, 3)
	var input string

	for input != "X" {
		fmt.Printf("Enter an integer: ")
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		i, _ := strconv.Atoi(input)
		intSlice = append(intSlice, i)
		// sort slice
		sort.Ints(intSlice)
		// print slice
		fmt.Println(intSlice)
	}
}
