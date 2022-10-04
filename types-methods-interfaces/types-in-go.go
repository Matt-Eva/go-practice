package main

import "fmt"

func main(){

}

// Go is a statically typed language, which means variable types
// are known at compile time. In statically typed languages - Rust,
// Java, C, C++ - types must be expressly indiciated by the programmer.

// However, some statically typed languages use "type inference", which
// allows the the programmer to not indicate their variable types.

// Go also allows you to attach methods to types. It implements type abstraction
// wich allows you to write code that invokes methods without explicitly
// specifying the implementation.

// Go differs greatly from most contemporary, commonly used languages
// in the way it approaches methods, interfaces, and types. It was designed
// to encourage best practices as advocated by software engineers, avoiding
// inheritance while encouraging composition.

// TYPES IN GO

// As we saw in the section on structs, we can define a struct type:

type Dog struct {
	Name string
	Age int
	Breed string
}

// This allows use to create a user defined type of Dog, which has the
// underlying type of the struct literal that follows.

// In addition to struct literals, you can use any primitive type or compound
// type literal to define a concrete type:

type Score int
type Converter func(string)Score
type TeamScroes map[string]Score


// You can declare these kinds of type at any block level, from the
// package block down. Types are only accessible from within their scope,
// with the exception of exported package block level types.

// ABSTRACT VS CONCRETE TYPE (AND HYBRID)

// An abstract type is a type that specifies what a type should do,
// but not how it is done. A concrete type specifies both what and how.
// This means it has a specified way to store its data and provides an
// implementation of any methods declared on the type. All types in Go are either
// abstract or concrete types. Some languages, like Java, allow "hybrid types".