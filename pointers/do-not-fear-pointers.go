package main

import "fmt"

func main(){

}

// POINTERS AND CLASSES

// While Java, JavaScript, Python, and Ruby don't
// have pointers available for programmers to use / play
// around with, they do still use them for basic 
// object-oriented operations, like keeping track of 
// instances of a class.

// The languages are all pass by value, like Go, which
// means that parameters in a function are variables
// whose values are copied from whatever is passed 
// in as the argument. 

//Every instance of a class in these languages is implemented as a pointer.
// When a class instance is pass to a function, the value
// being copied to the parameter is the pointer to the instance.

// If you use your parameter to modify the fields of an 
// instance of a class, that chance will be reflected in the
// outer variable as well, since the parameter and the outer
// variable are pointing to the same location in memory.

// However, if you choose to reassign the value of your parameter
// - to a new class instance, for example - this creates a 
// separate instance pointing to a different location in memory
// which does not affect the outer variable.

// An example of this phenomenon in Python:

// class Foo:
//     def __init__(self, x):
//         self.x = x


// def outer():
//     f = Foo(10)
//     inner1(f)
//     print(f.x)
//     inner2(f)
//     print(f.x)
//     g = None
//     inner2(g)
//     print(g is None)


// def inner1(f):
//     f.x = 20


// def inner2(f):
//     f = Foo(30)


// outer()


// Running this code results in:

// 20
// 20
// True

// inner1 modifies f because it is modifying a field
// within the instance while still referencing the same
// location in memory.

// inner2 does not modify f because it assigns the parameter
// value to a new instance - since the parameter is a copy of
// the original, reassigning its value will not affect the value
// of the original.