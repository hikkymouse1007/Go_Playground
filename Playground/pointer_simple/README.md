# ポインタの仕組みメモ
## 参考サイト
- https://qiita.com/Sekky0905/items/447efa04a95e3fec217f#%E3%83%9D%E3%82%A4%E3%83%B3%E3%82%BF%E5%9E%8B%E3%81%A8%E3%83%9D%E3%82%A4%E3%83%B3%E3%82%BF%E5%A4%89%E6%95%B0
- https://codetree.dev/golang-basics/pointers/

以下のコードを参考に解説
value: string型の変数("hello")
&value : 変数valueのアドレス
str: string型のポインタ変数(メモリの値にアドレスのみを受け取る)
&str: ポインタ変数strが参照先アドレスを格納しているRAM内のアドレス(varのアドレス)
*str: strが参照しているアドレス先に格納されている値(varの値)

```
package main

import "fmt"

func changeValue(str *string) { //str := 引数(メモリのアドレス
	fmt.Println("adress of (var)value: ", &str, "value of (var)value(=adress): ", str)
	*str = "changed!" // *strがValueに保存している参照先アドレスの値を変更する
	fmt.Println(str, &str, *str, &*str)
}

func main() {
	value := "Hello"
	fmt.Println("adress:", &value, "value:", value)
	changeValue(&value)
	fmt.Println("adress:", &value, "value:", value)
}

```

![ram_value](https://user-images.githubusercontent.com/54907440/104816010-aba64180-585b-11eb-8833-faa2b587f914.png)
