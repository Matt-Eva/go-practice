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

	myMap := map[string]int{
		"a":1,
		"b": 2,
		"c":3,
	}

	for k,v := range myMap {
		fmt.Println(k, v)
	}

	myString := "hello" 
	for i, r:= range myString{ // r here represents the numeric value of the character in a string
		fmt.Println(i, r, string(r))
	}

	complexString := "apple_π!"
	
	for i, r := range complexString{ // i represents the index value in the string - notice that 7 is skipped, since π takes up to indexes worth of bytes.
		fmt.Println(i,r,string(r))
	}

	numMap := map[string]int{
		"a": 1,
		"b": 2,
		"c":3,
	}

	for _,v := range numMap{ // doesn't modify source
		v += 1
	}
	fmt.Println(numMap)

	for k := range numMap{ // does modify source
		numMap[k] += 1
	}
	fmt.Println(numMap)

	samples := []string{"hello", "apple_π!"}
outer: // we indent labels to the same level as the block in which they are included
	for _, sample := range samples {
	inner: // while you can label a nested for loop, it's rare to do so
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
				// break outer // <-- also works
				// break inner
			}
		}
		fmt.Println()
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


// String iteration

// Due to the way strings are stored as bytes,
// using a for range loop to iterate over a string
// will iteate over the RUNES in a string, rather than
// the BYTES. 


// Whenever a for-range loop encounters a multibyte
//  rune in a string, 
// it converts the UTF-8 representation into a single
//  32-bit number and assigns it to the value.
//   The offset is incremented by the number of bytes 
//   in the rune. If the for-range loop encounters a 
//   byte that doesn’t represent a valid UTF-8 value, 
//   the Unicode replacement character 
//   (hex value 0xfffd) is returned instead.

// FOR-RANGE IS COPY

// When you use a for-range to iterate over a compound
// type, it copies the value from the compound type
// to the value variable. Changing the value variable
// doesn't change anything within your
// compound type.


// LABELING FOR STATEMENTS

// Labelling our for statments allows us to use the
// continue and break keywords on labelled for statements
// which is usefull when dealing with nested for loops.


// samples := []string{"hello", "apple_π!"}
// outer:
//     for _, sample := range samples {
//         for i, r := range sample {
//             fmt.Println(i, r, string(r))
//             if r == 'l' {
//                 continue outer
//             }
//         }
//         fmt.Println()
//     }