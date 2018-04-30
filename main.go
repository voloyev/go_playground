package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://cafephilo.lviv.ua",
		"http://http://10.6.193.162:3000/",
		"http://10.25.12.143:3000/",
	}

	c := make(chan string)

	for _, el := range links {
		go checkLink(el, c)
	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "it might be down")
		c <- "Might be down"
		return
	}

	fmt.Println(link, "is up")
	c <- "Its up"
}
