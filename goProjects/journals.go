package main

import "fmt"

type Journal struct {
	bay string
	size int
	color string
}

func main() {
	journal := new(Journal)

	journal.bay = "leather"
	journal.size = 7
	journal.color = "light brown"	

	fmt.Printf("This journal is:")
}
