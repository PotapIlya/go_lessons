package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//testJson1()
	testJson2()
}

func testJson2() {

	type UserJsonStruct struct {
		Id   int    `json:"user_id,string"`
		Name string `json:"user_name,string"`
	}

	user := UserJsonStruct{
		Id:   1,
		Name: "Potap",
	}

	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))

}
func testJson1() {

	type UserJsonStruct struct {
		Id   int
		Name string
	}

	const userJson = `{ "id": 1, "name": "Potap" }`

	data := []byte(userJson)
	u := UserJsonStruct{}

	err := json.Unmarshal(data, &u)
	if err != nil {
		_ = fmt.Errorf("EROR Unmarshal: %s", err.Error())
		return
	}

	fmt.Println(u)

	res, _ := json.Marshal(u)

	fmt.Println(res)

}
