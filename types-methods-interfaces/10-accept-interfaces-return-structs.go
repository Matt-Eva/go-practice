package main

import "fmt"

func main (){

}

// ACCEPT INTERFACES, RETURN STRUCTS

// The above dictates that the business logic of functions should be invoked
// via interfaces, while the output of functions should be concrete types.

// As discussed, functions should accept interfaces because they make code
// more flexible and explicitly declare what functionality is being used.

// There are a couple reasons we don't want to return instances. 

// The first
// is known as 'decoupling', which is one of the main advantages provided
// by implicit interfaces.

// If we build an API that returns interfaces, our client code now depends
// on the module that contains those interfaces, as well as any dependencies
// of that module. This limits future flexibility - it's better to return a concrete
// type. While this can still lead to dependencies, we can limit this effect
// using "dependency injection" (more on that later).

// Another reason has to do with versioning. If a concrete type is return,
// we can add new methods and fields without breaking existing code. This is
// not true for an interface. Adding new methods and fields to an interface
// means that we would need to update all existing implementations of an interface
// or else our code will break. If there is a backward-breaking change to an API,
// the major version number needs to be incremented.

// When writing functions, it's better to write separate functions for each
// concrete type, rather than a single function that returns different instances
// based on input parameters.

// There are certain circumstances where returning an interface is simply 
// unavoidable.

// Errors, however, are the one exception to this rule. Go functions and
// methods declare a return parameter of the error interface type. When it
// comes to error, it's very likely that different implementation of the interface
// could be returned, so you need to use an interface to handle all possible
// options, since interfaces are the only abstract type in Go.

//The potential drawback to this pattern comes from garbage collection.
// Reducing heap allocations improves performance by reducing garbage collection
// work. Return a struct avoids this heap allocation, which is good. 

// However, invoking a function with parameters of interface types causes
// a heap allocation to occur for each of the interface parameters. Optimizing
// betewen abstraction and increased performance will be determined over the 
// life of your program. Code should be written to be readable and maintainble.
// If interfaces are determined to be bogging down performance, then a function
// should be written to use a concrete parameter. If multiple implementations of
// an interface are passed into a function, that will mean creating multiple
// functions with repeated logic.