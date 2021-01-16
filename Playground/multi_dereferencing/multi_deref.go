package main

import "fmt"

func main() {
	str1 := "string1"
	// fmt.Println(*str1) // str1の値はアドレスではないので、このコードは動かない
	v1 := str1
	fmt.Println("&str1:", &str1, "str1:", str1, "&v1:", &v1, "str1:", str1)

	p1 := &str1 // p1の"value"に"str1のアドレス"を格納(&によって、メモリの"value"に"アドレス"を格納することを明示)
	fmt.Println("&str1:", &str1, "p1:", p1)
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "str1:", str1) // p1のvalueがアドレスであるので、*p1でアドレスの参照ができる => str1のvalueへアクセス
	*p1 = "updated!!"
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "str1:", str1)

	p2 := &p1
	fmt.Println("&str1:", &str1, "p1:", p1, "p2:", p2)
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "&p2:", &p2, "p2:", p2, "*p2:", *p2, "str1:", str1)
	*p2 = &str1
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "&p2:", &p2, "p2:", p2, "*p2:", *p2, "str1:", str1)

	p3 := &p2
	fmt.Println("&str1:", &str1, "p1:", p1, "p2:", p2, "p3:", p3)
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "&p2:", &p2, "p2:", p2, "*p2:", *p2, "&p3:", &p3, "p3:", p3, "*p3:", *p3, "str1:", str1)
	*p3 = &p1
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "&p2:", &p2, "p2:", p2, "*p2:", *p2, "&p3:", &p3, "p3:", p3, "*p3:", *p3, "str1:", str1)
}
