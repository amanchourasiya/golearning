package basics

import "time"
import "fmt"

func consumer(ch chan int){
	
	for{
	    x := <- ch
		fmt.Printf("consumer thread x:%d\n",x)
	}
}

func producer(ch chan int){
    for i := 0;i < 10;i++{
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
}

func ChannelDemo(){
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(20 * time.Second)
}