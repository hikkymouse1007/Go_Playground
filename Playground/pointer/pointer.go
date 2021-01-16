package main

import "fmt"

func changeValue(str string) { //str is "string"
	str = "changed!"
	fmt.Println(&str, str, *&str, "changeValue")
}

// ポインタ型変数(*type)を引数に取る場合は、ポインタ型の(=アドレス)しか代入できない
func changeValue2(str *string) { //str is "pointer_variable" of "string"
	fmt.Println(&str, *&str, str, *str, "before changeValue2")
	*str = "changed!"
	fmt.Println(&str, str, *str, "after changeValue2")
}

// ref: https://www.youtube.com/watch?v=a4HcEsJ1hIE&t=381s
func main() {
	toChange := "hello"
	fmt.Println(&toChange, toChange) // 0xc0000981e0 hello
	changeValue(toChange)            // 0xc000010230 changed! changed! changeValue
	fmt.Println(&toChange, toChange) // 0xc0000981e0 hello, need to return value
	changeValue2(&toChange)          // 0xc00000e030 0xc0000981e0 changed! changeValue2
	fmt.Println(&toChange, toChange) // 0xc0000981e0 changed!
}
