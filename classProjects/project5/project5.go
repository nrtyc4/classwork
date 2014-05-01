package main

import (
  "net/http"
  "io"
  "io/ioutil"
)

var messages chan string = make(chan string, 100)

var counter = 0

func PushHandler(w http.ResponseWriter, req *http.Request) {

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        w.WriteHeader(400)
    } 
    counter += 1
    messages <- string(body)
}


func PollResponse(w http.ResponseWriter, req *http.Request) {

    io.WriteString(w, <-messages)
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/poll", PollResponse)
    http.HandleFunc("/push", PushHandler)
    http.ListenAndServe(":8005", nil)
}
