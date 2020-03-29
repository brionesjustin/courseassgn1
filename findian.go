package main

import (
	"fmt"
	"strings"

)

func search(x string) string {
	if strings.HasPrefix(x, "i") {
		fmt.Println("Found!")
	} else if strings.HasSuffix(x, "n") {
		fmt.Println("Found!")
	} else if strings.Contains(x, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found")
	}
	return ""
}

func main() {
	var inputstring string
	fmt.Println("Please enter string:")
	_, err := fmt.Scan(&inputstring)
	if err != nil {
		fmt.Println(err)
	}
	search(inputstring)
}