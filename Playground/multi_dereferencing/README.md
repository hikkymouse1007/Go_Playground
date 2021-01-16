# 参照の参照
- https://qiita.com/tutuz/items/cc5752740e77bd013a68

```
package main

import "fmt"

func main() {
	str1 := "string1"
	// fmt.Println(*str1) // str1の値はアドレスではないので、このコードは動かない
	v1 := str1
	fmt.Println("&str1:", &str1, "str1:", str1)
	fmt.Println("&v1:", &v1, "v1:", v1)
	// &str1: 0xc0000821e0 str1: string1
	// &v1: 0xc0000821f0 v1: string1

	fmt.Println("")
	p1 := &str1 // p1の"value"に"str1のアドレス"を格納(&によって、メモリの"value"に"アドレス"を格納することを明示)
	fmt.Printf("%T\n", p1)
	// *string

	fmt.Println("")
	fmt.Println("&str1:", &str1, "p1:", p1)
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "str1:", str1) // p1のvalueがアドレスであるので、*p1でアドレスの参照ができる => str1のvalueへアクセス
	*p1 = "updated!!"
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1, "str1:", str1)
	// &str1: 0xc0000821e0 p1: 0xc0000821e0
	// &p1: 0xc0000a6020 p1: 0xc0000821e0 *p1: string1 str1: string1
	// &p1: 0xc0000a6020 p1: 0xc0000821e0 *p1: updated!! str1: updated!!

	p2 := &p1
	p3 := &p2
	fmt.Println("")
	fmt.Println("&str1:", &str1, "str:", str1)
	fmt.Println("&p1:", &p1, "p1:", p1, "*p1:", *p1)
	fmt.Println("&p2:", &p2, "p2:", p2, "*p2:", *p2)
	fmt.Println("&p3:", &p3, "p3:", p3, "*p3:", *p3)
	// &str1: 0xc0000821e0 str: updated!!
	// &p1: 0xc0000a6020 p1: 0xc0000821e0 *p1: updated!!
	// &p2: 0xc0000a6028 p2: 0xc0000a6020 *p2: 0xc0000821e0
	// &p3: 0xc0000a6030 p3: 0xc0000a6028 *p3: 0xc0000a6020

	fmt.Println("")
	fmt.Printf("str1 is %T\n p1 is %T\np2 is %T\np3 is %T\n", str1, p1, p2, p3)
	// str1 is string
	// p1 is *string
	// p2 is **string
	// p3 is ***string

	p1 = **p3
	p2 = *p3
	fmt.Println(p1, p2)
	fmt.Printf("str1 is %T\n p1 is %T\np2 is %T\np3 is %T\n", str1, p1, p2, p3)
}

```
![multi_deref](https://user-images.githubusercontent.com/54907440/104818267-efa04300-5869-11eb-9897-399fa02f787c.png)
