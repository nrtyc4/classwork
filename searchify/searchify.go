package main

import (
<<<<<<< HEAD
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
=======
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func artistHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req)
	fmt.Println(req.RequestURI)

	artisturl := []string{"https://ws.spotify.com/search/1"}
	appendedURL := strings.Join(artisturl, req.RequestURI)

	res, err := http.Get(appendedURL)
	if err != nil { log.Fatal( fmt.Println("Couldn't append URL")) }
	buf, err := json.Marshal(res)
	if err != nil { log.Fatal( fmt.Println("Couldn't marshal JSON")) }
	w.Write(buf)
>>>>>>> bccb342486aa256a19a8cc4e03bc4bb727e3d4ee
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
<<<<<<< HEAD

	http.HandleFunc("/artist", artistHandler)

	log.fatal(http.ListenAndServe(":8080", nil))
=======

	http.HandleFunc("/artist.json", artistHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
>>>>>>> bccb342486aa256a19a8cc4e03bc4bb727e3d4ee
}
