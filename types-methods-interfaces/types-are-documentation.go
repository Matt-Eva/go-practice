// TYPES ARE EXECUTABLE DOCUMENTATION

// While it's clear that you should declare a struct type to hold a set
// of related data, it's less clear when you should declare a user-defined
// type based on other built-in types or other user-defined types.

// These kinds of types are basically documentation. They make code 
// more explicit by describing the kind of data that's expected in 
// as certain context. It's clearer to someone reading your code what's
// going on when a method has a parameter of type Percentage rather
// than of type int, and it's harder for it to be invoked with an invalid
// value.

// The same rationale applies when declaring a user-defined type based on
// another user-defined type. When you have the same underlying data, but
// different set of operations to perform, you should make two types. Basing
// one of these types on the other type avoids repetition and makes it clear
// that the two types are related.
