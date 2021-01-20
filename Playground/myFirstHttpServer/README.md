# Ref
- https://www.youtube.com/watch?v=5BIylxkudaE
Port forward
- https://code.visualstudio.com/docs/remote/containers#_temporarily-forwarding-a-port
- https://qiita.com/tnk4on/items/7ccb1921907bdb99002a

# メモ

## go get
コンテナの中でgo get <package> すれば良い
https://qiita.com/maroKanatani/items/75c38f6ab2f474ef7d37

## ローカルとコンテナのポート接続
=>　以下のポート番号をポートに合わせて修正する
```
## .devcontainer/devcontainer.json
// Use 'forwardPorts' to make a list of ports inside the container available locally.
	"forwardPorts": [8080],
```

```
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
}
```

## stringと[]byteについて
https://qiita.com/ikawaha/items/2d58f58e4ab12918e8c9

