package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)
	i := 0

	for tickerTime := range ticker.C {
		i++
		fmt.Println("sleep", i, "time", tickerTime)

		if i >= 5 {
			ticker.Stop()
			break
		}
	}

	fmt.Println("Total ticks", i)
}
