package main

import(
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World</h1>")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)

}
