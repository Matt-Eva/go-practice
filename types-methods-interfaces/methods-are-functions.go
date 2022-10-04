package main

import "fmt"

func main(){

	myAdder := Adder{start: 10}

	fmt.Println(myAdder.AddTo(5))
	fmt.Println(myAdder)

	f1 := myAdder.AddTo // this is called on myAdder
	fmt.Println(f1(10))

	f2 := Adder.AddTo // this is called on the type Adder itself
	fmt.Println(f2(myAdder, 15)) // we then pass the value itself - myAdder - to f2.
}

// METHODS ARE FUNCTIONS

// Methods in Go are so similar to functions in Go that you can often
// replace a function with a method any time there is a variable or parameter
// of a function type.

// Ex:

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// we'll create an instance of the Adder type in the usual manner,
// then invoke its method:

var myAdder = Adder{start: 10}

// fmt.Println(myAdder.AddTo(5))

// We can also assign the method to a variable itself or pass it to 
// a paramter of the proper function type (in this case func(int) int).
// This is called a method value

//Ex:

var f1 = myAdder.AddTo
// fmt.Println(f1(10))

// A method value is similar to a closure, in that it can access the
// values in the fields of the instace from which it was created.

// We can also create a function from the type itself, which is called a 
// method expression:

// Ex:

var f2 = Adder.AddTo 
// note that this is called on the Adder type itself, 
// rather than a specific variable
// (like myAdder)

// fmt.Println(f2(myAdder, 15))

// Specific use cases for method values and method expressions are used
// in certain contexts - specifically dependency injection.


// FUNCTIONS VS METHODS

// Since methods can be used as functions, how do you differentiate between 
// when you should use either option?

// The determining factor is whether or not your function depends
// on other data. Package-level state should be effectively immutable,
// which means that any time your logic depends on values that are configured
// at startup or changed while your program is running, those values should
// be stored in a struct and the logic should be implemented as a method.
// If your logic only depends on the input parameters, then it should be
// a function.

// Types, packages, modules, testing, and dependency injection are 
// interrelated concepts.