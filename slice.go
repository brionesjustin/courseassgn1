package main

import (
	"fmt"
)

var inputint int



func sliceloop() {
	a := make([]int, 0)
	for {
		fmt.Println("Please enter an integer:")
		_, err := fmt.Scan(&inputint)
		if err != nil {
			fmt.Println(err)
		}
		if inputint == 0 {
			break
		}
		a = append(a, inputint)
	}
	fmt.Println(a)
}



func main() {
	sliceloop()
}