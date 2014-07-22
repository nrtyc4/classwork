package main

import (
	"fmt"
	"sort"
)

type Journal struct {
	ISBN int
	Color string
	Size int
	Bay string
}

func (j Journal) String() string {
	return fmt.Sprintf("%d: %s", j.ISBN, j.Color)
}

// ByColor implements sort.Interface for []Journal based on
// the Color field
type ByColor []Journal

// These methods are defined in the sort package
func (c ByColor) Len() int	{ return len(c) }
func (c ByColor) Swap(i, j int)	{ c[i], c[j] = c[j], c[i] }
func (c ByColor) Less(i, j int) bool	{ return c[i].Color < c[j].Color }

func main() {
	// Define and create an interface of Journal slices called journals
	journals := []Journal{
		{9780431, "dark brown", 7, "Leather Journal bay"},
		{9789321, "light brown", 6, "Leather Journal bay"},
		{9780062, "tan", 5, "Leather Journal bay"},
		{9784477, "gray", 4, "Leather Journal bay"},
	}
	
	fmt.Println(journals)
	sort.Sort(ByColor(journals))
	fmt.Println(journals)
}
