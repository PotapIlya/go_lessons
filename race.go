package main

import (
	"fmt"
	"sync"
)

func workerRace(counters map[int]int, idx int, mu *sync.Mutex) {
	mu.Lock()
	counters[idx]++
	mu.Unlock()
}

func main() {
	mu := &sync.Mutex{}
	counters := map[int]int{}

	for idx := range 5 {
		go workerRace(counters, idx, mu)
	}

	fmt.Scanln()
	mu.Lock()
	fmt.Println(counters)
	mu.Unlock()
}
