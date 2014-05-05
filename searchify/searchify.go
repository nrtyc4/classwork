package main

import (
	"encoding/json"
	"fmt" //debugging purposes for now
	"log"
	"net/http"
	"strings"
)

func artistHandler(w http.ResponseWriter, req *http.Request) {
	/* switch req.Method {
	case "GET":
	  buf, err := json.Marshal(someMapofStrings)
	  w.Write(buf)
	default:
	  w.WriteHeader(400)
	} */

	//fmt.Println(req)
	fmt.Println(req.RequestURI)

	artisturl := []string{"https://ws.spotify.com/search/1"}
	appendedURL := strings.Join(artisturl, req.RequestURI)

	res, _ := http.Get(appendedURL)
	buf, _ := json.Marshal(res)
	w.Write(buf)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/artist.json", artistHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
