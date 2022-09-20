package main

import (
	"fmt"
	"errors"
	"os"
)

func main(){
	sum, difference := sumAndDifference(5, 3)
	fmt.Println(sum, difference)

	sum2, _ :=sumAndDifference(1,2) // if we don't need to use our variable for the difference, we can just use _ as a placeholder. This is necessary due to Go preventing any unused variables, and requiring the exact number of variable declarations as return values in a function.
	fmt.Println(sum2) 
	
	anotherSum := add(1,2)

	fmt.Println(anotherSum)

	result, err := errorExampleDivision(4,2)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("result",result)

	fmt.Println(returnsBlanks("string"))
}

// MULTIPLE RETURN VALUES

// One thing that makes go different from
// other programming languages is its ability to return
// multiple values:

func sumAndDifference(num1, num2 int) (int, int){ // specifying multiple return values
	// newVar := "string" <--- can declare variables like this within any function
	// fmt.Println(newVar)
	return num1 + num2, num1 - num2
}

// if a function is specified to return multiple values
// it MUST return all of the those values, separated by commas
// (as shown above).


// CREATING AND RETURNING ERRORS
// Go allows you to return an "error" from a function.
// You have to declare that error as a return value type.
// If your function does not trigger an error, you should
// return a nil value for the error:

//Ex:

func errorExampleDivision(numerator, denominator int) (int, error){
	if denominator == 0 {
		return 0, errors.New("cannot divid by zero")
	}

	return numerator / denominator, nil
}

// NAMED RETURN VALUES

// Go allows to you specify names for your return values:

func add(num1, num2 int) (sum int){
	sum = num1 + num2
	return sum
	// return 5
}

// This creates a local variable - sum - that is
// initialized to its type's zero value. 
// Note that you do not have to return this variable
// explicitly so long as your return type matches
// the specified type.

// This named return is local only to the function.

// PROBLEMS WITH NAMED RETURNS

// Shadowing can occur when using a named return.
// Naming returns a generally unnecessary except in a 
// specific instance using the defer keyword.

// BLANK RETURNS - NEVER TO BE USED

// Another severe downside to named returns
// is that they allow you to simply write the 
// "return" keyword without specifying any return values.
// This returns the last values assigned to the
// named return values.

func returnsBlanks(value string) (result string){
	return
}

// ^^ Returns an empty string as the result.

// Blank returns make your code hard to understand.
// Don't use them.