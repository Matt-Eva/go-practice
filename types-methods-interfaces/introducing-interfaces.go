package main

import "fmt"

func main(){
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}

// INTERFACES

// Go is oft touted for its built in concurrency, but another
// key feature that it implements is its implicit interfaces,
// the only abstract type in Go.

// we declare an interface in a manner quite similar to a struct:

type Stringer interface {
	String() int
}

// an interface is a type, so we declare it using "type". 
// When declaring an inteface, we write an interface literal after
// the name of the interface type, which lists methods that must
// be implemented by a concrete type to meet the interface. The methods
// defined herein are called the method set of the interface.

// Interfaces can be declared in blocks, like other types.

// Interfaces are usually named such that their endings finish in "er"


// INTERFACES ARE TYPE-SAFE DUCK TYPING

// What's been listed above is common across many languages - that's
// typically how interfaces work. What makes Go special?

// What's interesting is that concrete type does not declare that it
// implements an interface; rather, if the method set for a concrete type
// contains all of the methods of an interface's method set, the concrete
// type implements the interface. This means that the concrete type
// can be assigned to a variable or field of the type of the interface it
// implements.

// This allows type safety and decoupling, which bridges the funcitonality
// between dynamic and static languages.


// WHY DO WE HAVE INTERFACES?

// A programming paradigm introduced in the book Design Patterns dictates
// that we "Program to an interface, not an implementation." Doing so means
// we can depend on behavior, rather than on implementation, which allows
// us to swap implementations as needed. This makes it easier for code
// to evolve over time, since requirements will inevitably change.


// DUCK-TYPING

// Dynamically typed languages like Python, Ruby, and JavaScript don't use interfaces.
// Instead, developers who use these languages rely on "duck typing", which comes
// from the expression "if it walks like a duck and quacks like a duck, it's a duck".
// The idea behind this is that you can pass an instance of a type as a parameter
// to a funciton as long as the function is able to find a method to invoke that
// it expects.

// This may sound a bit odd, but it has been used to build large and successful
// systems. To a developer who works in a statically typed language, this sounds
// like chaos.

// Languages like Java take a different approach. Java developers will define 
// an interface, create an implementation of that interface, but only
// refer to the interface in the client code (?)

// To developers who work with dynamic typing, this approach makes it seem
// impossible to refactor code when you have explicit dependencies.
// Switching to a new implementation from a different provider means
// rewriting code to depend on a new interface.

// Go compromises - if applications grow and change, programmers need the
// flexibility to change implementation. But, for people to effectively understand
// what code is doing, we need to specify what code depends on. Which is where
// Go's implicit interfaces come in:

type LogicProvider struct {
	field string
}

func (lp LogicProvider) Process(data string) string{
	// business logic
	fmt.Println(data)
	return data
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program(){
	// get data from somewhere
	c.L.Process("data")
}

// c := Client{
// 	L: LogicProvider{},
// }
// c.Program()

// Note that LogicProvider implements the Logic interface without us having
// to declare that LogicProvider meets the interface - on the caller (Client)
// knows about it. This allows us to easily generate a new logic provider in the future
// while also providing executable documentation to ensure that any type passed
// into the client will match the client's needs (have the requisite methods 
// and such).

// Interfaces specify what callers need.

// Notice that we defined the Client struct to have a type of the Logic interface
// to its field L, but that when we actually instantiate a Client struct on 104,
// we give L a field of LogicProvider. This is because LogicProvider actually
// implements the interface.

// Interfaces can actually be shared - there are several interfaces in the 
// standard library that are used for input and output. This is a powerful feature
// - if your code works with io.Reader and io.Writer, it will function
// correctly whether it is writing to a file on local disk or to a value in memory.

// Example:

// r, err := os.Open(fileName)
// if err != nil {
//     return err
// }
// defer r.Close()
// return process(r)

// os.Open returns an os.File isntance, which meets the io.Reader interface
// and can be used in any code that reads in data. You can even handle a
// gzip-compressed file by wrapping the io.Reader an another io.Reader:

// r, err := os.Open(fileName)
// if err != nil {
//     return err
// }
// defer r.Close()
// gz, err = gzip.NewReader(r)
// if err != nil {
//     return err
// }
// defer gz.Close()
// return process(gz)


// Rule of thumb: if there's an interface in the standard library that 
// describes what your code needs, use it!

// Note: types that meet an interface can still specify additional methods
// that aren't part of the interface - this is totally fine.

// For example, the io.File type also meets the io.Writer interface. However,
// if your code only cares about reading from a file, use the io.Reader interface
// to refer to the file instance and ignore the other methods.



