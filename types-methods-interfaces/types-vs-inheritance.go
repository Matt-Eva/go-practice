package main

import "fmt"

func main(){

	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	//hs = s // throws an error
	// s = i // throws an error
	s = Score(i)
	hs = HighScore(i)

	fmt.Println(s, hs)

	var q Score = 2
	var j = s + q // adding two Score types together, since they're underlying value is int.
				  // note that you cannot add a score type and a highscore type together without a type conversion.

	fmt.Println(j)

	var m HighScore = 4
	var z = hs + m // this also works, even though there are two levels of
				   // sub-typing away from Int.

	fmt.Println(z)
	

}

// TYPE DECLARATIONS ARE NOT INHERITANCE

// In addition to declaring types based on built-in Go types and struct
// literals, you can also declare a user-defined type based on another
// user-defined type:

type Score int

type HighScore Score

// If you're familiar with object-oriented languages, this seems similar
// to the concept of inheritance. However, this is not inheritance (and Go
// is not object-oriented).

// Instead, these two types simply have the same underlying type - that's all.
// This doesn't establish any sort of hierarchy between these types.

// In object orientation, child instances have access to all the methods and
//data structres of parent instances, and can be used anywhere the parent
// instance is used. This is not the case in Go. For example:

// You can't assign an instance of type HighScore to a variable of type Score
// (or vice versa) without a type conversion.

// Nor can you assign either of them to a variable type of type int
// without a type conversion.

// Furthermore, any methods defined on Score aren't defined on HighScore.

var i int = 300
var s Score = 100
var hs HighScore = 200

//hs = s // throws an error
// s = i // throws an error
// s = Score(i) // ok
// hs = HighScore(i) //ok 

	// fmt.Println(s, hs)

// For user-defined types whose underlying types are built-in types,
// a user-declared type can be used with operators for those types.

var q Score = 2
var j = s + q

// As seen above, they can also be assigned literals and constants 
// compatible with the underlying type. (hs has a value of 200, for example).

// Note: A type conversion between types that share an underlying type
// keeps the same underlying storage but associates different methods.


