package main

import "fmt"

func main(){
	var x [3]int
	fmt.Println(x)
	var y = [4]int{1,2,4,3}
	fmt.Println(y)
	var z = [12]int{1, 5: 4, 6, 10:100, 15}
	fmt.Println(z)
	x[0] = 1
	fmt.Println(x)
	fmt.Println(z[5])
	fmt.Println(len(z))
	var n [2][3]int
	fmt.Println(n)
}

//Arrays:
// arrays are sequential data structures 
// with index value pairs. They are of a
// fixed size that is determined when you
// declare them. All elements in an array
// are of the same datatype, which is 
// specified when it's created. Arrays
// of different length are considered to 
// be of different types.

var x[3]int 
//^^ creates an array of ints with 3 elements
// since no values are specified, it defaults
// to zero.

var y = [4]int{1,2,4,3}
//^^ creates an array of 4 ints, whose values
// are specified in the {}

var z = [12]int{1, 5:4, 6, 10:100, 15}
//^^ Allows you to set values only for specific
// indexes in the array. The index value is
// specified first, followed by the value.
// Subsequent values are placed in consecutive
// indexes, unless another index is specified.

//x[0] = 1
//^^ Arrays are read and written using
// bracket notation. You cannot negative index
// and array or read or write past the lenght
// of an array.

//len(z)
// ^the len function returns the length
// of an array.

var n [2][3]int

// We can simulate a multi-dimensional array
// by creating an array of arrays.

//Arrays are not often used in go due to their
// strict limitations. Instead, slices are used more
// often.







