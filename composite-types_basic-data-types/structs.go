package main

import "fmt"

func main(){

	type person struct {
		name string
		age int
		pet string
	}
	var fred person
	fmt.Println(fred)

	george:=person{} // fred and george result in same thing
	fmt.Println(george)

	fmt.Println(fred == george)

	jane:= person{
		"Jane",
		20,
		"lion",
	}
	fmt.Println(jane)

	bob:= person{
		age: 50,
		name: "Robert",
	}
	fmt.Println(bob)
	fmt.Println(bob.name)
	fmt.Println(jane == bob)

	var cat struct {
		name string
		age int
		weight int
	}
	cat.name="Chim"
	cat.age=20
	cat.weight=30
	fmt.Println(cat)
	// fmt.Println(cat == jane) // throws an error - comparing structs of different types

	dog := struct{
		name string
		age int
	}{
		name: "Goof",
		age: 10,
	}

	fmt.Println(dog)
}


// STRUCTS

// Structs are similar to classes in true
// object-oriented languages, in that they
// group certain types of data together.

// To declare a struct, write:

type person struct {
	name string
	age int
	pet string
}

// We can now create things of the "person" struct

var fred person

// Since we aren'st specifying any values
// for fred, every field - name, age, pet - is
// set to that field's data type's zero value

var jane = person{
	"Jane",
	20,
	"lion",
}

// ^^ One way to create a structure with filled-in
// fields. If used this way, fields must be
// put in in the order entered into the struct
// declaration

var bob = person{
	age: 50,
	name: "Robert",
}

// ^^ Way to create structs using the field names
// As you may have noticed, we don't have to 
// specify a value for every field - any field not
// specified is set to zero value.

// ANONYMOUS STRUCTS
// Can declare a variable set to a struct
// that doesn't have a specific name

 var cat struct {
	name string
	weight int
	age int
 }

 // cat is a variable set to an anonymous struct


// COMPARING STRUCTS

// whether or not a struct can be compared
// depends on the fields on the struct.
// Structs whose fields are slices or maps
// are not comparable.

// Go doesn't allow comparisons between variables
// that represent structs of different types.

// However, it does allow you to perform a type
// conversion from one struct to another if the fields
// of both structs have the same names, order, and types.


// Anonymous structs add a small twist to this:
// if two struct variables are being compared and 
// at least one of them has a type that’s an 
// anonymous struct, you can compare them without
// a type conversion if the fields of both 
// structs have the same names, order, and types.
// You can also assign between named and anonymous
// struct types if the fields of both structs have 
// the same names, order, and types:

