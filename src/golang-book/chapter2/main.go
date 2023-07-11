package main

import "fmt"

func main() {
	// Integers and floating points
	fmt.Println("1 + 1 =", 1.0 + 1.0)
	// Strings
	fmt.Println(len("Hello, World!"))
	fmt.Println("Hello, World!"[1])
	fmt.Println("Hello, " + "World!")
	// Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}