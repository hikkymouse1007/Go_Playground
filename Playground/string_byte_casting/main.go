package main

import (
	"fmt"
)

func main() {

	test := "sample"

	// string 型 → []byte 型
	b := []byte(test)
	fmt.Println(b) // [115 97 109 112 108 101]

	// []byte 型 → string 型
	s := string(b)
	fmt.Println(s) // sample

}
