package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum) // prints 12345678.9 but if it was float32
	// it would be 12345679.0
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result) // prints 12.1

	var intNum1 int = 10
	var intNum2 int = 3
	fmt.Println(intNum1 / intNum2) // div example
	fmt.Println(intNum1 % intNum2) // mod example

	var myString string = "Testing" + " " + "concatenation"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("Δ"))
	// with this import and function 1 is printed instead of 2 for the length

	var myRune rune = 'a'
	fmt.Println(myRune) // prints 97

	var myBoolean bool = false
	var myTBoolean bool = true
	fmt.Println(myBoolean)
	fmt.Println(myTBoolean)

	// Default Values

	var intNum3 rune
	fmt.Println(intNum3) // 0
	// for strings it's empty string ""
	// for boolean it's false

	// ways to initialize variables

	var myVar = "text"
	//inferred
	myVar1 := "text"

	var1, var2 := 1, 2
	fmt.Println(myVar, myVar1, var1, var2) //text text 1 2

	// Constants
	const myConst string = "never changing value once created"
	// no default value, always need to initialize when declared

	fmt.Println(myConst)

}
