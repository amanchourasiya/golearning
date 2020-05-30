package features

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sender(ch chan<- string, wg *sync.WaitGroup) {
	msg := []string{
		"Openstack",
		"KVM",
		"AWS",
		"Kubernetes",
		"GCP",
		"H/W",
		"VMWare",
	}

	for _, v := range msg {
		ch <- v
		fmt.Printf("Sender: Filling channel with platform: %s\n", v)
		//time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
	fmt.Println("Sender: Closing channel")
	close(ch)
	wg.Done()
}

func receiver(ch <-chan string, wg *sync.WaitGroup) {

	for {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		platform, ok := <-ch
		fmt.Printf("%v\n", ok)
		if ok {
			fmt.Printf("Received: platform:%s\n", platform)
		} else {
			fmt.Println("Channel closed. Exiting ...")
			break
		}

	}
	wg.Done()
}

func BufferChannelDemo() {
	ch := make(chan string, 5)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(ch, &wg)
	go receiver(ch, &wg)
	wg.Wait()
}
