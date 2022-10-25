package main

import "fmt"

func main(){

	m:= Manager {
		Employee: Employee{
			Name: "Jene Wean",
			ID: "98098",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID) // even though ID is an attribute of the embedded
	// struct, it can be accessed because Employee is an embedded field
	fmt.Println(m.Description())
	// Similarly, the Description method is available for m as well (it's
	// part of it's method set). Note that these are still run on the 
	// embedded struct.
	fmt.Println(m.Employee.ID) // this also works.

	o := Outer {
		Inner: Inner{
			X:10,
		},
		X: 20,
	}
	fmt.Println(o.X) // prints 20
	fmt.Println(o.Inner.X) // prints 10

	j := Iggle{
        Oogle: Oogle{
            A: 10,
        },
        S: "Hello",
    }
    fmt.Println(j.Double()) // this method is invoked on Oogle, which 
	// runs its method of IntPrinter, even though there is a method of the
	// same name for Iggle.
	
}

// EMBEDDING FOR COMPOSITION

// While Go doesn't have inheritance, it does encourage
// code reuse via built in support for composition and
// promotion:

type Employee struct {
	Name string
	ID string
}

func (e Employee) Description()string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

// in the Manger struct, note that Manager contains a field
// with the Employee type, but that no name has been assigned to that field.
// This makes Employee an embedded field. Any fields or methods declared
// on an embedded field are promoted to the containing struct and can
// be invoked directly on it. (see main() for example)


// Note that you can embed any type within a struct, not just another
// struct.


// If the containing struct and the embedded struct have the fields
// or methods of the same name, you need to use the embedded field's
// type to refer to the obscured fields or methods. Ex:

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

// EMBEDDING IS NOT INHERITANCE

// Built-in embedding support is rare in programming languages. Developers
// coming from more popular object oriented languages will likely be
// inclined to think of embedding as inheritance - it's not. For example,
// You can't assign a variable of type Manager to a variable of type Employee.

// If you want to access the Employe field in Manager, you must explicitly
// do so. Ex:

var m = Manager {
	Employee: Employee{
		Name: "Jene Wean",
		ID: "98098",
	},
	Reports: []Employee{},
}

// var eFail Employee = m // doesn't work - can't assign a variable of type
// Employee to a variable of type Manager
var eOK Employee = m.Employee // But this is ok, since you're accessing
// the embedded Employee field, which is of the Employee type.

// METHODS DON'T KNOW THEY'RE EMBEDDED

// The methods on an embedded field don't know that they are embedded.
// That means that if you have a method on an embedded field that calls
// another method of that field, that method will be called, even if 
// it has the same name as a method on the containing struct (which
// will not be invoked). See main for an example:

type Oogle struct {
    A int
}

func (i Oogle) IntPrinter(val int) string {
    return fmt.Sprintf("Oogle: %d", val)
}

func (i Oogle) Double() string {
    return i.IntPrinter(i.A * 2)
}

type Iggle struct {
    Oogle
    S string
}

func (o Iggle) IntPrinter(val int) string {
    return fmt.Sprintf("Iggle: %d", val)
}