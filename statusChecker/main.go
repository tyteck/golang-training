package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.apple.com",
		"https://www.golang.org",
		"https://www.podmytube.com",
	}

	c := make(chan string)

	/*
		this part will call checkWebsiteStatus
		in a go routine. The go routine will check the url
		but at the same time others go routines will be launched
		concurrently.
	*/
	for _, url := range urls {
		go checkWebsiteStatus(url, c)
	}

	/*
		in this loop we are telling go to sit and wait a channel return
		the channel will get a string(the url) that we are sending again
		to another go routine.
	*/

	// for i := 0; i < len(urls); i++ {
	//	go checkWebsiteStatus(<-c, c)
	//}
	// for {
	//	go checkWebsiteStatus(<-c, c)
	//}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkWebsiteStatus(link, c)
		}(l)
	}

}

func checkWebsiteStatus(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(url + " is down")
		c <- url
		return
	}
	fmt.Println(url + " is down")
	c <- url

}
