package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	countWorkers int = 3
	numJobs      int = 10
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	wg := &sync.WaitGroup{}
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for workerIdx := range countWorkers {
		wg.Add(1)
		go worker(workerIdx, jobs, results, wg)
	}

	for j := range numJobs {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("result is:", result)
	}

}
