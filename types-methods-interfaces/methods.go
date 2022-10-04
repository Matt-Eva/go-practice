package main

import (
	"fmt"
	"time"
)

func main (){
	
	var f = Floob {
		Field: "harbo",
	}
	var output = f.String()
	fmt.Println(output)

	var c Counter
	// fmt.Println(c.String())
	// c.Increment()
	// fmt.Println(c.String())

	doUpdateWrong(c)
    fmt.Println("in main:", c.String())
    doUpdateRight(&c)
    fmt.Println("in main:", c.String())

    var it *IntTree
    it = it.Insert(5)
    it = it.Insert(3)
    it = it.Insert(10)
    it = it.Insert(2)
    fmt.Println(it.Contains(2))  // true
    fmt.Println(it.Contains(12)) // false
}

// METHODS

// Like many languages, we can create methods that we can call
// on user defined types (think classes in OOP).

// Methods for a type are defined at the package block level:

type Thing struct {
	Field1 string
	Field2 int
	Field3 string
}

func (t Thing) String() string {
	return fmt.Sprintf("%s %s, age %d", t.Field1, t.Field2, t.Field3)
}

// Methods are very similar to functions in their declarations, except for
// the receiver specification, which appears between the func keyword and the
// the name of the method (in this case its (t Thing)).  What we're actually
// defining here is the receiver variable - by convention, it's shortened
// to be the first letter of the type's name. Go does not use self or this (it
// is considered non-idiomatic).

// You cannot overload methods (as with functions). You can use the same method
// names for different types

type Floob struct {
	Field string
}

func (f Floob) String()string { // it's okay to create this method because it's for a different type than the method above
	return "fnoof"
}

// but you cannot use the same method name for two different methods on the
// same type. (There can be no other String method for Thing or Floob.)
// This helps keep your code clear and clean.

// Methods must be declared in the same package as their associated type;
// Go doesn't allow you to add methods to types you don't control. 
// It's possible to define a method in a different file within the same
// package as the type declaration, but it's best to keep your type definition
// and its associated methods together.

// Method invocation is similar to that of JavaScript, Ruby, and Python

// var f = Floob {
//	Field: "harbo"
// }
// var output = f.String()

// POINTER RECEIVERS AND VALUE RECEIVERS

// Function use parameters of pointer type in a function
// to indicate that a parameter might be modified by that function.

// The same rules apply for method receivers as well. They can be 
// pointer receivers (the type is a pointer) or value receivers (the type is
// a value). These rules help determine which kind of receiver to use:

	// If your method modifies the receiver, you MUST use a pointer receiver.

	// If your method needs to handle nil instances, then it must use a pointer
	// receiver.

	// If your method doesn't modify the reciever, you can use a value receiver.


// Whether or not you should use a value receiver depends on the other
// methods declared for that type. When a type has any pointer receiver
// methods, it's common practice to be consistent and use pointer receivers
// for all methods, even those that don't modify the receiver.

// Example:

type Counter struct {
    total             int
    lastUpdated time.Time
}

func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}
// Uses pointer receiver type

func (c Counter) String() string {
    return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
// Uses value receiver type. However, the code works the same if we use a 
// pointer receiver type.

// While we specify c as a pointer type in Increment, we still call
// it on a Counter type in the normal manner. This is due to the
// fact that Go automatically converts a local variable with a value type
// to a pointer type when this type of method is used on it. Basically,
// c.Increment is converted to (&c).Increment().


// Be aware that passing a value type to a function and calling a pointer
// receiver method on the passed in value will result in you invoking the method
// on the copy, which will not update the original method.

// EX:

func doUpdateWrong(c Counter) {
    c.Increment()
    fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
    c.Increment()
    fmt.Println("in doUpdateRight:", c.String())
}

// Using a pointer type as the parameter remedies this problem.

// Note that although c in doUpdateRight is of pointer type, we can
// still call Increment and String on it, which are methods of thet type
// that the pointer points to. Go considers both pointer an value receiver
// methods to be in the method set for a pointer instance. For a value instance,
// only the value receiver methods are in the method set. (Note - read up 
// more on what this means and why this is different than go
// converting value type variables to pointer type variables
// when )

// NOTE:
// Do not write getter and setter methods for Go structs, unless you need
// them to meet an interface. Instead, directly access a field and save 
// methods for business logic. Exceptions occur when you need to update
// multiple fields as a single operation or when the update isn't a
// straightforward assignment of a new value. (Increment demonstrates both
// of these properties).

// METHODS AND NIL INSTANCES

// Side Note: How to implement a binary tree in go

type IntTree struct {
	val int
	left, right *IntTree
}

// What happens when you call a method on a nil instance in Go? In many
// languages, this results in some sort of error. In Go, it will result
// in a panic if you use a method with a value receiver. However, if the
// method is built with a pointer receiver, it can work so long as the method
// is written to handle the possibility of a nil instance.

// Here's the rest of the code used to implement a binary tree using nil:

func (it *IntTree) Insert(val int) *IntTree {
    if it == nil {
        return &IntTree{val: val}
    }
    if val < it.val {
        it.left = it.left.Insert(val)
    } else if val > it.val {
        it.right = it.right.Insert(val)
    }
    return it
}

func (it *IntTree) Contains(val int) bool {
    switch {
    case it == nil:
        return false
    case val < it.val:
        return it.left.Contains(val)
    case val > it.val:
        return it.right.Contains(val)
    default:
        return true
    }
}

// Note that both methods use a pointer receiver, even though
// the Contains method doesn't actually modify the tree. This
// is so that the Contains method can handle a nil value.

// Being able to call a method on a nil receiver is a handy tool to have
// in your toolkit, but it's often not very useful. However, pointer receivers
// are like pointer function parameters - they're just copies of the pointer
// that's passed into the method. If you change the copy of the pointer, 
// that doesn't change the original, which means you can't write a pointer 
// receiver method that handles nil and converts the original pointer to 
// something non-nil.