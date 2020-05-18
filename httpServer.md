# httpServerとして動かす

とりあえずポート8080と、8080/helloが機能するようにする。

## goを以下のコードとする。

```go
package main

import(
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    dump, err := httputil.DumpRequest(r, true)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
    fmt.Println(string(dump))
    fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World</h1>")
}

func main() {
    var httpServer http.Server
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", helloHandler)
    log.Println("start http listening :8080")
    httpServer.Addr = ":8080"
    log.Println(httpServer.ListenAndServe())

}
```

## dockerから実行

```
# docker-compose exec app go run main.go
```


## 確認
以下で確認する。
- http://localhost:8080
- http://localhost:8080/hello
