package main

import (
<<<<<<< HEAD
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
=======
  "net/http"
  "log"
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

func main(){
	http.Handle("/", http.FileServer(http.Dir("./")) )

  http.HandleFunc("/artist", artistHandler)
  
	log.fatal(http.ListenAndServe(":8080", nil))
}
>>>>>>> d3ca3f1cb120cec177f976561e9ac0a72dce1007
