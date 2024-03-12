package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	encodeStructToJson()
	// var user User
	// var jsonString = `{"name":"Riksa","age":20}`
	// err := json.Unmarshal([]byte(jsonString), &user)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(user.Name)
	// fmt.Println(user.Age)
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func decodeJsonToMap() {
	var user map[string]interface{}
	var jsonString = `{"name":"Riksa","age":20}`
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user["name"])
	fmt.Println(user["age"])
}

func encodeStructToJson() {
	var user = User{
		Name: "Riksa",
		Age:  20,
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
