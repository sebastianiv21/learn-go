package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please input a file name")
		os.Exit(1)
	}

	file, err := os.Open(args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
