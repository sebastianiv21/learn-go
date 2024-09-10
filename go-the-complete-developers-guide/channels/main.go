package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	// create a channel that shares data of type string
	c := make(chan string)

	for _, link := range links {
		// we only use go keywords in front of function calls
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// the channel waits for any go routine that completes to print its output
	// fmt.Println(<-c)

	// wait for the channel to return some value, after that, assign that value to "l"
	// then run the instructions inside the for loop
	for l := range c {
		// function literal == anonymous function
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // we pass l to the function literal as an argument
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
