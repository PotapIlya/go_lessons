package main

import "fmt"

func main() {

	type User struct {
		name string
		age  int
	}

	var users []User = []User{
		{name: "Bob", age: 20},
		{name: "Bob1", age: 21},
		{name: "Bob2", age: 22},
	}

	fmt.Println(users)

}
