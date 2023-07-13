package main

import (
	"fmt"
	"math"
	"strconv"
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

	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar

	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)

	// Converting, Parsing and Formatting Values
	fmt.Println("--------------------------------")
	fmt.Println("Convert, Parse and Format Values")
	fmt.Println("--------------------------------")

	kayak := 275
	soccerBall := 19.50

	total := float64(kayak) + soccerBall

	fmt.Println(total)

	// Limitations of Explicit Conversion
	fmt.Println("--------------------------------")
	fmt.Println(" Limits of Explicit Conversions ")
	fmt.Println("--------------------------------")

	total2 := kayak + int(soccerBall)

	fmt.Println(total2)
	fmt.Println(int8(total2))

	// Converting floating-Point Values to Integers
	fmt.Println("--------------------------------")
	fmt.Println("Convert Float Values to Integers")
	fmt.Println("--------------------------------")

	total3 := kayak + int(math.Round(soccerBall))

	fmt.Println(total3)

	// Parsing from Strings
	fmt.Println("--------------------------------")
	fmt.Println("      Parsing from Strings      ")
	fmt.Println("--------------------------------")

	val1 := "true"
	val2 := "false"
	val3 := "not true"

	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)

	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)
}	