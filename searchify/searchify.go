package main

import (
	"log"
	"net/http"
)

func artistHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		buf, err := json.Marshal(someMapofStrings)
		w.write(buf)
	default:
		w.WriteHeader(400)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/artist", artistHandler)

	log.fatal(http.ListenAndServe(":8080", nil))
}
