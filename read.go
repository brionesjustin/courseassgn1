package main 

import (
	"fmt"
	"io/ioutil"
)

type name struct {
	fname string
	lname string
}

var textfile string

func rdfl() {
	fmt.Println("Please enter text file:")
	_, err := fmt.Scan(&textfile)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Open(textfile)
	barr := make([]byte, 10)
	nb, err := f.Read(barr)
	f.Close()

}


//func main() {

}