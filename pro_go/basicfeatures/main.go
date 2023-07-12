package main

import (
	"fmt"
	"sort"
	// "math/rand"
)

func main() {
	
	// first := 100
	// second := &first
	// third := &second

	// fmt.Println(first)
	// fmt.Println(*second)
	// fmt.Println(**third)

	names := [3]string {"Alice", "Charlie", "Bob"}
	secondName := names[1]
	secondPosition := &names[1]

	fmt.Println(secondName)
	fmt.Println(*secondPosition)
	sort.Strings(names[:])
	fmt.Println(secondName)
	fmt.Println(*secondPosition)
}