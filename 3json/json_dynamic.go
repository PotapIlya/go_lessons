package main

import (
	"encoding/json"
	"fmt"
)

const userDynamicJson1 = `[
{ "id": 1, "name": "Potap" },
{ "id2": 1, "name2": "Potap" }
]`

func main() {

	var user1 interface{}

	err := json.Unmarshal([]byte(userDynamicJson1), &user1)
	if err != nil {
		return
	}

	fmt.Println(user1)

	user2 := map[string]interface{}{
		"id":   42,
		"name": "test",
	}

	res, _ := json.Marshal(user2)

	fmt.Println(string(res))

}
