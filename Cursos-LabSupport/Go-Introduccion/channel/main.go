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
		"http://gsgvzxolang.orgxc",
	}

	c := make(chan string) // Canal de tipo string

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		//fmt.Println(<-c)
		//go checkLink(<-c, c) == a lo de abajo
		//go checkLink(l, c)
		go func(l string) {
			time.Sleep(5 * time.Second)
			go checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "migth be down!")
		c <- "Hola"
		return
	}

	fmt.Println(link, "is up!")
}
