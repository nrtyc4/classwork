package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting Go Routines")
	go func() {
		time.Sleep(1 * time.Microsecond)
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting to finish")
	time.Sleep(1 * time.Second)	//give the main goroutine enough time for other goroutines to finish
	fmt.Println("\nTerminating Program")
}
