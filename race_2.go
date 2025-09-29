package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalOperation int32

func main() {
	for range 1000 {
		go func() {
			atomic.AddInt32(&totalOperation, 1)
		}()
	}

	time.Sleep(time.Millisecond)
	fmt.Println(totalOperation)
}
