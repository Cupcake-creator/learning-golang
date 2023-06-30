package main

import (
	"fmt"
	"net/http"
	"time"
)

// func main() {
// 	links := []string{
// 		"http://google.com",
// 		"http://facebook.com",
// 		"http://stackoverflow.com",
// 		"http://golang.org",
// 		"http://amazon.com",
// 	}

// 	c := make(chan string)

// 	for _, link := range links {
// 		go checkLink(link, c)
// 	}

// 	for l := range c {
// 		go func(link string) {
// 			time.Sleep(5 * time.Second)
// 			checkLink(link, c)
// 		}(l)
// 	}
// }

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		c <- link
// 		return
// 	}

// 	fmt.Println(link, "is up!")
// 	c <- link
// }

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	// First time that checklink is called
	for _, l := range links {
		go checkLink(l, c)
	}

	// When get value from Routine do checklink it again
	//(like ping that host continously)
	// for {
	// 	//Sleep for not ping that site so fast
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(<-c, c)
	// }
	// ||
	// ||

	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, "is up!")
	}
	c <- link
}
