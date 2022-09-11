package main

import "fmt"
import "math/rand"

func notMain(){
	fmt.Println("Coming from notMain")
}

func main(){
	notMain()

	x := 12
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	if x> 5{
		fmt.Println(x)
	}

	fmt.Println(true)
	true := 7
	fmt.Println(true)
}




// BLOCKS
// Blocks are essentially like scopes in JavaScript
// Except there is also a 'file block' which is where
// the names of our imports are 'scoped' to.
// More local blocks can access values in more outer blocks

// SHADOWING VARIABLES

// While we can access variables declared in 
// more global blocks within more local blocks,
// if we create a new variable of the same name 
// within the local block, the local block will
// refer to the local version of that variable
// (following the new declaration).
// The more global variable remains unchanged.

// This more local variable is known as a
// "shadowing variable".

// Be careful about accidentally shadowing
// package names in particular - if you try 
// to access functionality in the original package
// after having created a shadowing variable,
// you'll get an error. Ex:

// fmt := "oops"
// fmt.Println(fmt) <-- this will throw an error

// Basically it's best to never use shadow variables

// THE UNIVERSE BLOCK
// Go only has 25 actual keywords
// A lot of the other keywords that are used 
// - including built in types and constants (true, false)-
// are considered 'predeclared identifiers' and
// are defined in the univers block. As such, these can
// all be shadowed:

//fmt.Println(true)
// true := 10
// fmt.Println(true)

// ^^ NEVER SHADOW THESE ^^

