package main

import "fmt"

func main(){

var x int32 = 10
var y bool = true


var pointerX = &x
// var pointerY = &y
// var pointerZ *string

var q = 5 + *pointerX

var n *bool
n = &y

fmt.Println(q) // prints 15
fmt.Println(n)

}

// POINTERS

// A poitner is essential a variable
// that holds the location in memory where a
// value is stored.


// Every variable is stored in
// one or more contiguous memory locations called
// addresses. Different types of variables  can take
// up different amounts of memory.

// The smallest amount of memory that can be
// individually addressed is a byte.

// Ex:

var x int32 = 10
var y bool = true

// x is a 32-bit int, which requires four bytes, so
// it takes up 4 addresses. y is a boolean, which 
// only requires a single bit. However, since the smallest
// amount of memory that can be addressed is a byte
// a boolean takes up a single byte and a single address.

// A pointer itself is just a variable whose contents 
// are the address where another variable is stored:

var pointerX = &x
var pointerY = &y
var pointerZ *string

// The size of a pointer is always the same - a number
// that points to the location in memory of a variable.
// Therefore a pointer itself always takes up the same
// amount of memory / number of addresses.

// The value of pointerX would be the starting address
// for x. The value for pointerY would be the starting
// address for y. The value of pointerZ would be 0, since
// it doesn't point to anything.


// ZERO VALUE

// The zero value for a pointer is nil. 
// nil is an untyped idetifier that represents
// the lack of a value for certain types.

// A note on nil: because nil is declared in the
// universal block, it can be shadowed. Never name
// a variable or function nil.

// &

// & is the address operator. It precedes a value type
// and returns the address of the memory location where
// the value is stored.

// *

// * is the indirection operator. It preceds a variable
// of pointer type and returns the pointed-to value. This
// is called dereferencing.

// Ex:

var q = 5 + *pointerX

// q will have a value of 15.
// To dereference a pointer, you msut make sure that
// the pointer is non-nil. You program will panic if
// you attempt to dereference a nil pointer.

// POINTER TYPE

// A pointer type is a type that represents a pointer
// It is written with a * before a type name. A pointer
// type can be based on any type.

// The pointer type is the type of the variable
// to which the pointer is pointing (don't know why
// this is so fucking hard for people to explain 
// clearly).

// new()

// The built-in function new creates a pointer variable
// It returns a pointer to a zero value instance of
// the provided type.

// new() is not often used. Generally speaking,
// it's better to use &.

// However, you can't use & before a primitive literal
// (numbers, booleans, and strings) or a constant
// because they don't have memroy addresses; they
// only exist at compile time. When you need a pointer
// to a primite type, declare a variable and point to
// it.

var l string
var i = &l

// Ex: If you have a struct whose field value is a pointer
 // to a primitive data type, you can't pass a literal directly
 // to the field:

 type person struct {
    FirstName  string
    MiddleName *string
    LastName   string
}

var p = person{
  FirstName:  "Pat",
  // MiddleName: "Perry", // This line won't compile
  LastName:   "Peterson",
}

// You could create a variable to hold the constant value
// or create a function that takes in a primitive type -
// boolean, numeric, or string - and returns a pointer to
// that type.

func stringp(s string) *string{
	return &s
}

var o = person{
	FirstName: "Pat",
	MiddleName: stringp("Perry"), // this works
	LastName: "Peterson",
}

// This works because Go is pass by value - the constant
// is copied to the parameter, which is a variable, and therefore
// has an address in memory. The function then returns the
// variable's memory address.