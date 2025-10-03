//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func writer() <-chan int {
//	ch := make(chan int)
//	wg := sync.WaitGroup{}
//
//	wg.Add(2)
//
//	go func() {
//		defer wg.Done()
//		for i := range 3 {
//			ch <- i
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		for i := range 3 {
//			ch <- i + 100
//		}
//	}()
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	return ch
//}
//
//func main() {
//	writer := writer()
//
//	for v := range writer {
//		fmt.Println(v)
//	}
//
//}
//
////func main() {
////	writer := writer()
////
////	for {
////		v, ok := <-writer
////		if !ok {
////			break
////		}
////
////		fmt.Println(v)
////	}
////
////}
