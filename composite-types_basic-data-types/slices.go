package main

import "fmt"

func main(){
	var x = []int{10, 20, 30}
	fmt.Println(x)
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)
	fmt.Println(len(y))
	var z = []int{19: 2}
	fmt.Println(z)
	fmt.Println(len(z))
	var n = [][]int{{1, 2}, {2, 4}}
	fmt.Println(n)
	n[0] = []int{8, 9}
	fmt.Println(n)
	var q []int
	fmt.Println(q == nil) // prints true
	fmt.Println(len(q)) // prints 0
	q = append(q, 2, 3)
	fmt.Println(q)
	fmt.Println(cap(q))
	f := make([]int, 5, 10)
	fmt.Println(f)
}

//Slices:
// Slices are declared a lot like 
// arrays, except we do not specify
// a set length when creating them:

var x = []int{10, 20, 30}

var y = []int{1, 5: 4, 6, 10: 100, 15}
// we can also declare values only for set
// indexes. This generated a slice with
// 12 elements. The code below generates
// a slice of 20 elements:

var z = []int{19: 2}

var n [][]int
//^^ Like arrays, we can mimic 
// multidimensional slices using a slice
// of slices

var q []int
//^^ creating a slice with no values
// gives q the zero value for a slice: nil
// Nil represents a lack of value for some
// types. A nil slice contains nothing

// len(q)
//^^ slices also have the len function, 
// like arrays.

// q = append(q, 1, 2)
//^^ We can add values to the end of a slice
// using the append function. Go passes a copy
// of the original slice to the function, then
// adds the new values to the slice. We need to
// reassign the slice to make sure it references
// the new value.

// Capacity:
// Each slice has a specified capacity, which
// can exceed its lenght. When appending to a
// slice, you increase its length by one. Once
// a length reaches a capacity, there is
// no more room to add values. Appending more
// values will cause the Go runtime to allocate
// a new slice witha larger capcity. The original
// slice is copied to the new slice, the new
// values are added, and the new slice is returned.
// To get a slices capacity, use the cap function:

//cap(q)

//To create a slice with the correct initial capacity,
// we use the make function:

var f = make([]int, 5, 10) // creates a slice with a 
// length of 5 and capacity of 10