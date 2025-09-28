//package main
//
//import (
//	"context"
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func worker(ctx context.Context, idx int, ch chan<- int) {
//	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
//	fmt.Println(idx, "sleep", waitTime)
//
//	select {
//	case <-ctx.Done():
//		return
//	case <-time.After(waitTime):
//		fmt.Println("worker", idx, "done")
//		ch <- idx
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	result := make(chan int, 1)
//
//	for i := range 10 {
//		go worker(ctx, i, result)
//	}
//
//	fmt.Println("result found by", <-result)
//}
