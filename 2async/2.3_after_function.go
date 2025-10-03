package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second*1, func() {
		fmt.Println("hello")
	})

	_, _ = fmt.Scanln()
	timer.Stop()
	_, _ = fmt.Scanln()
}
