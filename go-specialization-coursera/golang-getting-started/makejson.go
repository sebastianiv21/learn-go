package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	var jsonMap map[string]string

	jsonMap = make(map[string]string)

	fmt.Printf("Enter a name: ")
	fmt.Scan(&name)
	fmt.Printf("Enter an address: ")
	fmt.Scan(&address)

	jsonMap["name"] = name
	jsonMap["address"] = address

	jsonBytes, _ := json.Marshal(jsonMap)

	fmt.Println(string(jsonBytes))
}
