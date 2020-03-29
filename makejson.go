package main

import (
	"fmt"
	"encoding/json"
)

var name string
var address string

func main() {
	m := make(map[string]string)
	fmt.Println("Please enter a name and address:")
	_, err := fmt.Scan(&name, &address)
	if err != nil {
		fmt.Println(err)
	}
	m["name"] = name
	m["address"] = address
	//for key, value := range m {
		//fmt.Println("Key:", key, "Value:", value)
	b, err := json.Marshal(m)
	fmt.Println(b)

}