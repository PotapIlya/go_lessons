package main

import (
	"example.com/lessons/user"
	"fmt"
)

func main() {

	newUser := user.NewUser("1", 1)

	newUser.SetName("Potap")

	fmt.Println(*newUser)

}
