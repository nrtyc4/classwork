package main

import "fmt"

type Journal struct {
	Bay string
	Size int
	Color string
}

func main() {
	journal := new(Journal)

	journal.Bay = "leather"
	journal.Size = 7
	journal.Color = "light brown"	

	fmt.Println("This journal is:")
}
