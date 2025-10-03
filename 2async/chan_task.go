package main

import (
	"fmt"
	"time"
)

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func doubler(inputCh <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for v := range inputCh {
			time.Sleep(500 * time.Millisecond)
			ch <- v * 2
		}
		close(ch)
	}()

	return ch
}

func reader(inputCh <-chan int) {
	for v := range inputCh {
		fmt.Println(v)
	}
}

func main() {
	reader(doubler(writer()))
}
