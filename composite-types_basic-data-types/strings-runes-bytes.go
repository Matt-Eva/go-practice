package main

import "fmt"

func main(){
	str := "I am a string"
	str = str + " another string"
	fmt.Println(str)
}


// Strings

// strings are composed of sequences of bytes.
// However, several Go library functions 
// assume that a string is composed of a sequence of
// UTF-8-encoded code points.

// We can extract a single value from a string
// using the index of the string.
// However, when dealing with strings which
// include characters that are multiple bytes long
// in UTF-8, we have to be careful with indexing,
// since we may index partway through the bytes depicting
// the character, rather than getting the whole 
// character as intended.

// Note - converting an integer into a string 
// using type conversion will cause problems due to
// the way strings work under the hood.

// You can convert bytes and runes to strings, and can
// convert strings to slices of runes and / or bytes.


// Rather than using slice or index expressions with
// strings, we should extract substrings and code
// points from strings using functiosn in the 
// strings and unicode/utf8 packages.