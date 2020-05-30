package features

import (
	"fmt"
	"math/rand"
	"time"
)

func imageSearch(ch chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		ch <- "image content"
	}
}

func webSearch(ch chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		ch <- "web content"
	}
}

func newsSearch(ch chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		//time.Sleep(10)
		ch <- "news content"
	}
}

func SelectDemo() {

	imageChan := make(chan string)
	webChan := make(chan string)
	newsChan := make(chan string)
	go imageSearch(imageChan)
	go webSearch(webChan)
	go newsSearch(newsChan)

	for {
		time.Sleep(1 * time.Second)
		select {
		case content := <-imageChan:
			fmt.Printf("Image channel: %s\n", content)

		case content := <-webChan:
			fmt.Printf("Web channel: %s\n", content)

		case content := <-newsChan:
			fmt.Printf("News channel: %s\n", content)

		default:
			fmt.Printf("Still searching\n")

		}
	}
}
