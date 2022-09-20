package main

import "fmt"

func main(){ // the main func will never have parameters

	result := divide(2, 2) // we cannot invoke functions outside of main
	fmt.Println(result)

	testingParams(1, 2, "string")
}

func divide(numerator int, denominator int) int { // notice parameter type and return type.
	fmt.Println("running divide")
	if denominator == 0{
		return 0 // like javascript, the return keyword halts a function
	}
	return simpleDiv(numerator, denominator) // functions can invoke other functions
}

func simpleDiv(numerator, denominator int) int{ // if multiple parameters have the same type, you can write your parameters like this
	fmt.Println("running simpleDiv")
	return numerator / denominator
}

func testingParams(num1, num2 int, string1 string) { // can declare parameters like this
	fmt.Println(num1, num2, string1)
}


// DECLARING FUNCTIONS

// Functions have five main parts - the keyword func, the 
// function name, the input parameters (whose type must be
// specified), the return type, and the body of the funciton itself.

// RETURN
// Like javascript, go uses the return keyword
// to return values from a function. It's not
// needed if your function returns nothing.