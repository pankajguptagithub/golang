package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p Person
	data := `{"Name":"Ramesh", "Age":30}`
	json.Unmarshal([]byte(data), &p)
	fmt.Println(p.Name, p.Age)
}
