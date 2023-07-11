// Package main implements code snippets from Adam Freeman's book
package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

// PrintHello writes a greeting using the fmt.Println function
func PrintHello() {
	fmt.Println("Hello, Go!")
}

// PrintNumber writes the number sent in arg using the fmt.Println function
func PrintNumber(number int) {
	fmt.Println(number)
}