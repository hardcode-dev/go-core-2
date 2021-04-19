package main

import (
	"fmt"
)

func main() {
	const bufSize = 10

	var bufchan chan int = make(chan int, bufSize) // буферизированный канал
	bufchan <- 10                                  // неблокирующая операция
	val := <-bufchan                               // получение значения в новую переменную
	fmt.Println(val)

	for i := 1; i < bufSize; i++ {
		bufchan <- i * 2
	}
	close(bufchan) // важно всегда закрывать канал!

	for val := range bufchan {
		fmt.Println("получено из канала:", val)
	}

	ch := make(chan int) // переменная типа канал, принимающий int
	ch <- 100            // блок
	val = <-ch           // блок
	fmt.Println(val)
}
