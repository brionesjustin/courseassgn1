package main

import (
	"fmt"
	"math"
)



func truncated(x float64) float64 {
	t := math.Trunc(x)
	return t
}


func main() {
	var inputfloat float64

	fmt.Printf("Please enter float64 value:")
	_, err := fmt.Scan(&inputfloat)
	if err != nil {
		fmt.Println(err)
	}
	newnum := truncated(inputfloat)
	fmt.Println(fmt.Sprint(newnum))
}