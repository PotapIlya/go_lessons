package main

import "fmt"

func main() {

	var users map[string]string = map[string]string{
		"k-test1": "v-test1",
		"k-test2": "v-test2",
	}

	//if _, ifExist := users["test3"]; !ifExist {
	//	fmt.Print("Key don't exist")
	//	return
	//}

	for key, value := range users {
		fmt.Println(key, value)
	}

}
