package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, Operations!")

	// Arithmetic Operators
	fmt.Println("--------------------------------")
	fmt.Println("      Arithmetic Operators      ")
	fmt.Println("--------------------------------")
	price, tax := 275.00, 27.40

	sum 		:= price + tax
	difference 	:= price - tax
	product 	:= price * tax
	quotient 	:= price / tax

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)

	// Arithmecit Overflow
	fmt.Println("--------------------------------")
	fmt.Println("       Arithmetic Overflow      ")
	fmt.Println("--------------------------------")
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))

	// Remainder Operator
	fmt.Println("--------------------------------")
	fmt.Println("       Remainder Operator       ")
	fmt.Println("--------------------------------")
	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))

	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)

	// Increment and Decrement Operators
	fmt.Println("--------------------------------")
	fmt.Println(" Increment & Decrement Operators")
	fmt.Println("--------------------------------")
	value := 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)

	// Concatenating Strings
	fmt.Println("--------------------------------")
	fmt.Println("      Concatenating Strings     ")
	fmt.Println("--------------------------------")
	greetings 		:= "Hello"
	language 		:= "Go"
	punctuation 	:= "!"
	combinedString 	:= greetings + ", " + language + punctuation

	fmt.Println(combinedString)

	// Comparison Operators
	fmt.Println("--------------------------------")
	fmt.Println("      Comparison Operators      ")
	fmt.Println("--------------------------------")
	first := 100
	const second = 200.00

	equal 				:= first == second
	notEqual 			:= first != second
	lessThan 			:= first < second
	lessThanOrEqual		:= first <= second
	greaterThan 		:= first > second
	greaterThanOrEqual 	:= first >= second

	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)

	// Comparing pointers
	fmt.Println("--------------------------------")
	fmt.Println("       Comparing pointers       ")
	fmt.Println("--------------------------------")
	// first := 100

	second1 := &first
	third := &first

	alpha := 100
	beta := &alpha

	fmt.Println(second1 == third)
	fmt.Println(second1 == beta)
	fmt.Println(*second1 == *third)
	fmt.Println(*second1 == *beta)

	// Logical Operators
	fmt.Println("--------------------------------")
	fmt.Println("        Logical Operators       ")
	fmt.Println("--------------------------------")

	maxMph := 50
	passengerCapacity := 4
	airbags := true
}	