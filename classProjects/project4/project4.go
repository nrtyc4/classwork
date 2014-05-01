package main

import (
  "net/http"
  "log"
)

func main() {  
    
  //start simple static web server
  http.Handle("/", http.FileServer(http.Dir("./")))
  err := http.ListenAndServe(":8004", nil)
  
	//if errors, log them
  if err != nil { log.Fatal("ListenAndServe: ", err) }
}
