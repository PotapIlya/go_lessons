package main

import "fmt"

func main() {

	defer fmt.Println("end")
	defer fmt.Println("end2")
	fmt.Println("Start")

}
