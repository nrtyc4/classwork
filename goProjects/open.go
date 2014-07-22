package main

import (
	"os"
	"log"
)

func main() {
	file, err := os.Open("sort.go")
	if err != nil { log.Fatal(err) }
}
