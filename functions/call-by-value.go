package main

import "fmt"

func main(){
	num := 1
	name := "Ben"
	noChange(num, name)
	fmt.Println(num, name)

	

	fido := dog{
		age: 10,
		name: "fido",
	}
	noChangeStruct(fido)
	fmt.Println(fido)
}

// CALL BY VALUE

// Go is "call by value" in that if you supply a variable
// for a parameter to a function, Go always makes a copy of the
// value of that variable.

// That means that passing the variables as arguments to a function
// then attempting to change their values within the function
// will not work.

func noChange(num int, name string){
	num = 0
	name = "Jorge"
}

// This behavior applies to things structs as well:

type dog struct {
	age int
	name string
}

func noChangeStruct(d dog){
	d.name="skippy"
}

// It isn't true for maps and slices, however. The values within
// both can be modified (although the length of a slice cannot be
// increased). This is due to the fact that both maps and slices 
// use pointers. This includes struct fields that are maps and slices.