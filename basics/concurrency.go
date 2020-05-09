package basics

import (
	"fmt"
	"time"
)

func ParallelTask(s string) {
	i := 0
	for {
		if i == 10 {
			break
		}
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%s i: %d\n", s, i)
		i++
	}
}

func ConcurrencyDemo() {
	go ParallelTask("goroutine1")
	go ParallelTask("goroutine2")
}
