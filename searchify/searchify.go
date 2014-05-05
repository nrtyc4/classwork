package main

import (
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
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/artist.json", artistHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
