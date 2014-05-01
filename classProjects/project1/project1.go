package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Options struct {
	Path string
	Port string
}

func main() {
	// set default options
	op := &Options{Path: "./", Port: "8000"}

	// read config file into memory
	data, _ := ioutil.ReadFile("./config.json")

	// parse config and log files, store results in "op"
	json.Unmarshal(data, op)

	//For debugging purposes
	//log.Println("Parsed options from config file: ", op)

	//start simple static web server
	http.Handle("/", http.FileServer(http.Dir(op.Path)))
	err := http.ListenAndServe(":"+op.Port, Log(http.DefaultServeMux))

	//if errors, log them
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//Function to get user IP, Method and file/folder they are requesting
//Takes http.DefaultServeMux as input, returns http.Handler
//Function also logs user data to log file (when provided at the command line)
//  and stdout
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
