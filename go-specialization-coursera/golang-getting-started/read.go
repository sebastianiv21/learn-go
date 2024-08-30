package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	var names []Name

	fmt.Printf("Enter a file name: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fullName := strings.Split(scanner.Text(), " ")
		names = append(names, Name{fullName[0], fullName[1]})
	}

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}
