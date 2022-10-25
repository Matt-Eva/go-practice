package main

import "fmt"

func main() {
	var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil)

	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName string
	} {"Fred", "Fredson"}
}

// INTERFACES AND NIL

// As with pointer types, nil is the zero value for interfaces.
// However, this is not quite as simple as it is for concrete types.

// In order for an interface to be nil, both the type and te value must be nil.

// In the Go runtime, interfaces are implemented as a pair of pointers, one to
// the underlying interface type, one to the underlying value. As long as the type
// is non-nil, the interface is non-nil. (Since variables cannot exist without
// a type, the type pointer will always be non-nil if the value pointer is non-nil)

// nil interfaces cannot have methods invoked upon them - doing so introduces
// a panic. Any non-nil interface can have a method invoked on it, although if
// the value is nil this could still introduce a panic if methods aren't designed
// to handle them.

// It can be tricky to tell whether or not a value is nil when the type is non-nil.
// We can use reflections (mentioned later) to find out.

// The Empty Interface Says Nothing

// In statically typed languages, we somtimes need to make a variable
// that can store any value. Go uses interface{} to represent this:

// var i interface{}
// i = 20
// i = "hello"
// i = struct {
//     FirstName string
//     LastName string
// } {"Fred", "Fredson"}

// interface{} isn't special case syntax. This empty interface type simply states
// that the variable can store any value whose type implements zero or more methods,
// which as it happends matches every type in Go. 

// Since an empty interface doesn't really tell you anything about the value
// it represents, there isn't a lot you can do with it. One common use of empty
// interface is as a placeholder for data of uncertain schema that's read from
// an external source, like a JSON file:

// one set of braces for the interface{} type,
// the other to instantiate an instance of the map
// data := map[string]interface{}{}
// contents, err := ioutil.ReadFile("testdata/sample.json")
// if err != nil {
//     return err
// }
// defer contents.Close()
// json.Unmarshal(contents, &data)
// the contents are now in the data map

// Another use of interface{} is as a way to store a value in a user-created data
// structure. This is due to Go's current lack of user-defined generics.
// If you need a data structure beyond a slice, array, or map, and you don't
// want it to work with only a single type, we need a field of type interface{}
// to hold its value.

type LinkedList struct {
	Value interface{}
	Next *LinkedList
}

func (ll *LinkedList) Insert(pos int, value interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &Linkedlist{
			Value: val,
			Next: ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

// ^^Note that this is NOT a good implementation of a linked-list - don't actually
// use this.


// If you do see a function that takes an empty interface, it's probably using
// reflection to populate or read the value. 

// This type of situation should be rare - it's best to generally avoid using
// interface{}. Go is designed to be strongly typed - working around this is
// unidiomatic.

// If you are in a situation in which you had to store a value into an empty
// interface, you may be wondering how to read the value back again. This leads
// into Type Assertions and Type Switches.