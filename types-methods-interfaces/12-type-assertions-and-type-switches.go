package main

import (
	"fmt"
	"reflect"
)

func main(){
	type MyInt int

	var i interface{}
	fmt.Println(reflect.TypeOf(i))
	var mine MyInt = 20
	i = mine
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i.(MyInt))
	i2 := i.(MyInt) // i2 is of type MyInt
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2 + 1)
	// i3 :=i.(string) //<-- this throws a panic, due to wrong type assertion.
	// fmt.Println(i3)
	// i4 := i.(int) // <-- this also throws a panic - wrong type assertion again.
	// fmt.Println(i4 + 1)

	i5, ok := i.(int) // <-- this is how we can handle a possible panic do to incorrect type assertion.
	if !ok {
		fmt.Println(fmt.Errorf("unexpected type for %v",i))
	} else {
		fmt.Println(i5 + 1)
	}
	

}

// TYPE ASSERTIONS AND TYPE SWITCHES

// Go gives us two ways to see if a variable with an 
// interface type has a specific concrete type
// or if the concrete type implements another interface.

// The first of these is type assertions:

// A type assertion names the concrete type that implemented the interface
// or names another interface that is also
// implemented by the concrete type underlying the interface.

// Ex:

// type MyInt int

// func main() {
//     var i interface{}
//     var mine MyInt = 20
//     i = mine
//     i2 := i.(MyInt)
//     fmt.Println(i2 + 1)
// }

// The variable i2 is of type MyInt. Note - what is this notation, i.(MyInt)?
// I guess this is type assertion notation itself?

// If the type assertion is wrong, the code panices:

// i3 :=i.(string)
// fmt.Println(3)

//Go is very careful about concrete types. Even though the underlying type
// of MyInt is int, the following code will still throw a panic:

// i2 := i.(int)
// fmt.Println(i2 + 1)

// since we don't want our code to crash, we can use the comma ok idiom
// to handle this type of behavior:

// i5, ok := i.(int)
// if !ok {
//     return fmt.Errorf("unexpected type for %v",i)
// }
// fmt.Println(i5 + 1)


// Note: Type assertions and type conversions are very different
// A type conversion can be applied to both concrete types and interfaces
// and are checked at compile time. Type assertions can only be applied
// to interface types and are checked at runtime.