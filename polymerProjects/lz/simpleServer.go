package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
})
    panic(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><head></head><body><h1>Welcome Home!</h1></body></html>")
}