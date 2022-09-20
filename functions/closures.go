package main

import "fmt"

func main(){

	closure:=func(){ // cannot create a typical named function within another function - must use anonymous function variable declaration.
		fmt.Println("From within a closure")
	}
	closure()

	takesFunc(1, isCallback)

	makeFunc(2)(3)

	innerFunc := makeFunc(4)
	innerFunc(6)

}

// CLOSURES

// Functions declared within functions are called closures.
// Closures start to get interesting when you 
// pass them as values to other functions or are returned
// from functions.

// Closures in go are very similar to closures in 
// javascript - the concept of "callback functions"
// is present and utilized, although not by that name.

// PASSING FUNCTIONS AS PARAMETERS

// you can pass functions as parameters:

func takesFunc(num int, someFunc func(int) int ){
	someFunc(num)
}

func isCallback(num int) int {
	fmt.Println(num)
	return num
}

// RETURNING FUNCTIONS AS FUNCTIONS

// You can return closures from another function (this
// is all very similar to javascript).

func makeFunc(num1 int) func(int) int{ // notice that the specified return type is a function
	return func(num2 int) int {
		fmt.Println(num1 * num2)
		return num1 * num2
	}
}

// This inner function can be saved to a variable or immediately invoked

// Returning closures is often used when we build middleware for a web
// server. Go also uses closure for resource cleanup via the defer
// keyword.