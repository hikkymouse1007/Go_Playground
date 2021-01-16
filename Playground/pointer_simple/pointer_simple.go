package main

import "fmt"

func changeValue(str *string) {
	fmt.Println(str)
	*str = "changed!"
	fmt.Println(&str)
}

func main() {
	fmt.Println("Hello, playground")
	value := "Hello"
	fmt.Println("adress:", &value, "value:", value)
	changeValue(&value)
}
