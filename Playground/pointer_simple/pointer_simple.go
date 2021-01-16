package main

import "fmt"

func changeValue(str *string) { //変数str := 引数(メモリのアドレス) を作成
	fmt.Println(&str, str, *str)
	*str = "changed!" // 変数strがValueとして保存している参照先アドレスの値を変更する
	fmt.Println(str, &str, *str, &*str)
}

func main() {
	value := "Hello"
	fmt.Println("adress:", &value, "value:", value)
	changeValue(&value)
	fmt.Println("adress:", &value, "value:", value)
}
