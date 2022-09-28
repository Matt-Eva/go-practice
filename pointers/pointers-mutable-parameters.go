package main

import "fmt"

func main(){
	x := 10
    failedUpdate(&x)
    fmt.Println(x) // prints 10
    update(&x)
    fmt.Println(x) // prints 20
}

// POINTERS MEAN MUTABLE PARAMETERS

// A note on constants: Go constans provide names
// for literal expressions that can be
// calculated at compile time. There is no mechanism
// in the language to declare that other kinds
// of values are immutable, even though
// modern software engineering embraces 
// immutability.

// To work around this, Go developers use pointers
// to indicate that a parameter is mutable.
// Since Go is a call by value language, the
// values passed to functions are copies. For
// nonpointer types like primitives, structs,
// and arrays, this means that the called function
// cannot modify the original. 

// However, if a pointer is passed to a funciton,
// the function recives a copy of the pointer, which
// still points to the original data, meaning
// that the original data can by modified by the
// function.


// IMPLICATIONS

// When you pass a nil pointer to a function
// you cannot make the value non-nil. You can
// only reassign the value if there was a value
// already assigned to the pointer. Since the memory
// location was passed to the function using
// call-by-value, we can't change the memory
// address.

// To actually reassign the value associated
// with a pointer parameter, you must dereference
// the pointer and set the new value.
// If you change the value of the pointer itself,
// you change the copy, not the original. Dereferencing
// puts the new value in the memory location pointed
// to by both the original and the copy. Ex:

func failedUpdate(px *int) {
    x2 := 20
    px = &x2
}
// ^^ reassigns the value of the pointer itself.

func update(px *int) {
    *px = 20
}
// ^^Reassigns the value at the memory location
// that the pointer points to.