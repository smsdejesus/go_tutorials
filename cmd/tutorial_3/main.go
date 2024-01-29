package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "tutorial 3"
	printMe(printValue)

	var num int = 12
	var den int = 3
	var result, remainder, err = intDivision(num, den)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The Result is just %v", result)
	} else {
		fmt.Printf("Result is %v, with a Remainder of %v", result, remainder)
	}
	// when doing checks if it is 1==1 || 2==2 then only 1==1 is checked
	// same for 1==2 && 2==2 only 1==2 is checked since it is false for an AND

	// SWITCH VERSION
	switch {
	case err != nil:
		fmt.Printf(err.Error())
		//break not needed
	case remainder == 0:
		fmt.Printf("The Result is just %v", result)
	default:
		fmt.Printf("Result is %v, with a Remainder of %v", result, remainder)
	}
	// SWITCH W/ Variable
	switch remainder {
	case 0:
		fmt.Printf("no remainder")
	case 1, 2:
		fmt.Printf("twas close")
	default:
		fmt.Printf("not close")
	}
}
func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error // default value = nil
	if denominator == 0 {
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
