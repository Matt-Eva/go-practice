package main

import "fmt"

func main(){
	slice := []int{2, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	d := 1
	for d < 100{
		fmt.Println(d)
		d = d * 2
	}

	for {
		fmt.Println(d)
		d = d * 2
		if (d > 1000){
			break
		}
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	oddVals := []int{1,3,5,7,9}
	for _, v := range oddVals {
		fmt.Println(v)
	}
}

// For Loops

// Can write a basic for loop in a manner very
// similar to javascript - only the parentheses
// are missing:

// for i := 0; i < 10; i++ {
//     fmt.Println(i)
// }

// Note: you MUST use := to initialize the variables;
// "var" is not allowed here. You can also shadow variables
// with a for block, as with if statements.

// CONDITION ONLY FOR LOOP

// go allows you to leave off both the variable
// initialization and the increment in a for loop.
// (but you still need a declared variable and a way
// to increase its value)
// This looks similar to a while statement in Python,
// Ruby, or JavaScript

// BREAK STATEMENTS

// the break statements exits the for loop.


// CONTINUE

// go allows the use of the 'continue' keyword
// to skip to the next iteration of a loop. This can
// be used to clean up code:

// for i := 1; i <= 100; i++ {
//     if i%3 == 0 {
//         if i%5 == 0 {
//             fmt.Println("FizzBuzz")
//         } else {
//             fmt.Println("Fizz")
//         }
//     } else if i%5 == 0 {
//         fmt.Println("Buzz")
//     } else {
//         fmt.Println(i)
//     }
// }

// ^^ messy


// for i := 1; i <= 100; i++ {
//     if i%3 == 0 && i%5 == 0 {
//         fmt.Println("FizzBuzz")
//         continue
//     }
//     if i%3 == 0 {
//         fmt.Println("Fizz")
//         continue
//     }
//     if i%5 == 0 {
//         fmt.Println("Buzz")
//         continue
//     }
//     fmt.Println(i)
// }

// ^^ Cleaner, through use of continue.



// FOR-RANGE 
// The for range loop is similar to the
// 'for of' or 'for in' loop in JavaScript, although
// the syntax is significantly different.
// We can use a for range with strings, arrays,
// slices, and maps.
// You can only use for range loops to iterate over built-in
// compound types and user-defined types that are
// based on them.

// Here an example of a for range loop with a slice

// evenVals := []int{2, 4, 6, 8, 10, 12}
// for i, v := range evenVals {
//     fmt.Println(i, v)
// }

// the first variable declared here - i - indicates 
// the position in the compound type. For arrays,
// slices, and strings, that references the index.
// For maps, it would reference the key (and would often
// be declared as k).

// The second variable - v - contains the value associated
// at a specific position.

// If you don't want to use the key or
// index in your for range,  you can 
// simply leave that section blank using an underscore:

// oddVals := []int{1,3,5,7,9}
// for _, v := range oddVals {
// 	fmt.Println(v)
// }

// If you just want the first value - the key, for
// example - you can just declare the first variable

// uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
// for k := range uniqueNames {
//     fmt.Println(k)
// }

// ^^ This is so versatile! Especially for a map
// that is being used as a set. More often
// used with maps than with arrays or slices
