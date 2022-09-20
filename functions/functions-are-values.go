package main

import (
	"fmt"
	"strconv"
)

func main(){

	// The code below generates and tests a 
	// simple calculator application.
	var opMap = map[string]func(int, int) int{ // because all of the functions have the same funciton integer, we can declare the type of this map to be the funciton signature, func(int,int) int
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
        []string{"2", "+", "3"},
        []string{"2", "-", "3"},
        []string{"2", "*", "3"},
        []string{"2", "/", "3"},
        []string{"2", "%", "3"},
        []string{"two", "+", "three"},
        []string{"5"},
    }
    for _, expression := range expressions {
        if len(expression) != 3 {
            fmt.Println("invalid expression:", expression)
            continue
        }
        p1, err := strconv.Atoi(expression[0])
        if err != nil {
            fmt.Println(err)
            continue
        }
        op := expression[1]
        opFunc, ok := opMap[op]
        if !ok {
            fmt.Println("unsupported operator:", op)
            continue
        }
        p2, err := strconv.Atoi(expression[2])
        if err != nil {
            fmt.Println(err)
            continue
        }
        result := opFunc(p1, p2)
        fmt.Println(result)
    }


	hasInner()

}

// FUNCTIONS ARE VALUES

// Functions have "signatures"
// which are built using the keyword "func"
// and the types of a function's parameters
// and return values. 

// Any function that has the exact same number
// and types of parameters and return values 
// generates the same type signature.

// All of these functions have the same signature:

func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

// Function type declarations:

// We can create new types using the type keyword
// similar to the manner we would with a struct:

// type opFuncType func(int,int) int


// This then allows us to use this type - opFuncType -
// as a type in our code (in a map for instance).

// ANONYMOUS FUNCTIONS

// In Go, you're able to declare new functions within
// another function. These inner functions do not
// need a name and can be assigned to variables or
// directly invoked.

func hasInner(){
	for i:=0; i<3; i++{
		func(j int){
			fmt.Println("printing", j, "from within the anonymous function")
		}(i)
	}
	// If you use the approach above, it throws a compile time erro
	// if you try to assign the function a name

	anonFunc := func (param string){
		fmt.Println(param)
	}
	anonFunc("Hello")

	//As does this technique
}

// The use of declaring anonymous functions and immediately
// invoking them is limited. Two use cases involve
// the defer keyword and launching a go routine.