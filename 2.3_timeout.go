package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(10 * time.Second)

	select {
	case <-timer.C:
		fmt.Println("timer.C timeout happened")
	case <-time.After(time.Second * 5):
		fmt.Println("time.After timeout happened")
	}
}
