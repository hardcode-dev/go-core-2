package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(ch chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		ch <- rand.Intn(1000)
	}
	close(ch)
}

func processor(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int) // трубка между потоками generator и ?
	go generator(ch)     // как называется такой шаблон?
	processor(ch)
}
