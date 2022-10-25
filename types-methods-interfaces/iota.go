package main

import "fmt"

func main(){
	const (
		Uncategorized MailCategory = iota
		Personal
		MineCraft
		Spam
		Social
		Advertisements
	)
	fmt.Println(Uncategorized)
	fmt.Println(Personal)
	fmt.Println(Spam)

	const (
		WakeUp MorningRoutine = iota
		DrinkSmoothie
		DrinkCoffee
		BrushTeeth
	)

	fmt.Println(WakeUp)
	fmt.Println(BrushTeeth)


}

// iota

// Many programming languages support enumerations, which allow
// you to specify that a certain type can only have a limited set
// of values.

// Go doesn't have an enumeration type - instead, it uses "iota", which
// allows you to assign an increasing value to a set of constants.

// To use iota, best practice dictates that you first define a type
// based on int that will represent all of the valid values:

type MailCategory int

type MorningRoutine int

// Then, you would use a const block to define a set of values for
// your type

// const (
// 	Uncategorized MailCategory = iota
// 	Personal
// 	Spam
// 	Social
// 	Advertisements
// )

// Let's examine this syntax: the first constant has its type
// and value declared - the type is the type based on int set previously,
// and the value is set to iota, which starts at 0 by default.

// If you don't want your first constant to have a value of zero,
// you can assign the first value in your const block to _ or to
// a constant that indicates the value is invalid.

// The other constants don't have their types or values explicitly set.
// Instead, Go automatically assigns them the same type as the first
// constant, and gives them a value of iota, with the value of iota
// incrementing on each line.


// USING IOTA

// Don't use iota for defining constans whose values are explicitly defined
// elsewhere. For example, when implementing portions of a specification
// which says which values are assigned to which constants, you should
// explicitly write the constant values. Iota should be ussed for "internal"
// purposes only - i.e., where the constants are referred to by name rather
// than by value. That way you can enjoy using iota by inserting new 
// constants into your list at any time / location without the risk
// of breaking everything.


// To clarify - nothing in Go stops you or any other programmer from 
// creating additional values of your type. Moreover, if you insert a new
// identifier in the middle of your list of literals, all of the subsequent
// literals will be renumbered (by inserting MineCraft, we bumped the value
// of Spam from 2 to 3, along with all subsequent consts). This will
// cause subtle breakages if those constants represented values in
// another system or database. 

// For these reasons, it only makes sesne to use iota when you need
// to differentiate between a set of values, but don't particularly care
// what the values are behind the scenes. If the value itself actually
// matters, specify it explicitly.