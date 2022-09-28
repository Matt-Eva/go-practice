package main

import (
	"fmt"
	"encoding/json"
)

func main(){

	hum := struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}{}
	err:= json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &hum)
	fmt.Println(err)
}

// POINTERS ARE A LAST RESORT

// Pointers make it difficult to understand
// data flow, and can create extra work
// for the garbage collector (which is Go's #1
// performance bottleneck).

// For example, rather than populating a struct
// by passing a pointer to it into a function
// simply have the function instantiate and
// return the struct.

type Goo struct {
	Field1 string
	Field2 int
}

func MakeGoo(f *Goo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}
// ^^ Don't do this

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo() (Foo, error){
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

// The only time you should use pointer parameters
// to modify a variable is when a function
// expects an interface. This pattern occurs when
// working with JSON (Go's standard library includes
// JSON support in the encoding/json package).

// The json Unmarshal function - json.Unmarshal - populate
// a variable from a slice of bytes containing json (example
// in main func). It takes a slice of bytes and an interface param.
// The value passed in for the interface param must be a pointer.
// If it isn't, an error is returned.

// Go doesn't have generics, which is why this pattern is employed.
// It's an exception, rather than a common case (as a new Go developer might thinkg)


// When returning values from a function, it's best
// to return value types, rather than pointer types.
// There are certain circumstances in which we'll need
// to pass pointers, but generally it's best to use values.

// POINTER PASSING PERFORMANCE

// If a struct is large enough, there are performance improvements
// from using a pointer to the struct rather than the struct itself
// as an input parameter or a return value. Passing a pointer
// is constant time for all datatypes (around one nanosecond), since the
// size of a pointer is the same for all data types.
// Passing a value into a function takes longer as the data gets larger - 
// once data gets to be around 10 megabytes of data, it takes about one
// millisecond (that's a long time!)

// POINTER RETURNING PERFORMANCE
// For data structures that are smaller than a megabyte
// it's actually slower to return a pointer type than a value type.
// However, once you're data structure is larger than a megabyte,
// the performance advantage flips.

// These are very short times, so generally this won't be much of a concern.
// However, this is a good reason to use a pointer even if data is meant
// to be immutable.

// ZERO VALUE VS NULL VALUE

// A common usage of pointers in Go is to indicate
// the difference between a variable or field that's
// been assigned the zero value and a variable
// or field that hasn't been assigned a value at all.

// If this is an important distinction for your program,
// use a nil pointer to represent an unassigned variable
// or struct field.

// Because pointers also indicate mutability, be careful when
//  using this pattern. Rather than return a pointer set to nil 
//  from a function, use the comma ok idiom that we saw for maps 
//  and return a value type and a boolean.

// // Remember, if a nil pointer is passed into a function via a 
// parameter or a field on a parameter, you cannot set the value 
// within the function as there’s nowhere to store the value. If 
// a non-nil value is passed in for the pointer, do not modify it 
// unless you document the behavior.

// // Again, JSON conversions are the exception that proves the 
// rule. When converting data back and forth from JSON (yes, we’ll 
// 	talk more about the JSON support in Go’s standard library in 
// 	“encoding/json”), you often need a way to differentiate 
// 	between the zero value and not having a value assigned at all. 
// 	Use a pointer value for fields in the struct that are nullable.

// When not working with JSON (or other external protocols), 
// resist the temptation to use a pointer field to indicate no value. 
// While a pointer does provide a handy way to indicate no value, 
// if you are not going to modify the value, you should use a value 
// type instead, paired with a boolean.