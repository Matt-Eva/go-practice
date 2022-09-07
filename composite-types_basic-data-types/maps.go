package main

import "fmt"

func main(){
	newMap := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(newMap)
	fmt.Println(newMap["key1"])
	newMap["key1"] = 3
	fmt.Println(newMap)
	newMap["key3"] = 4
	fmt.Println(newMap)

	sliceMap := map[string][]string{
		"Names": []string{"Kevin", "Julia", "Aaron"},
		"Animals": []string{"Dog", "Cat", "Human"},
	}
	fmt.Println(sliceMap)
	ages := make(map[int][]string, 10)
}

// Maps in go are similar to hashes in 
// ruby or python or objects in javascript.

// Maps are declared with the following syntax:
// map[keyType]valueType

var newMap = map[string]int{
	"key1": 1,
	"key2": 2,
}

// To read or write to a map, you can use the following
// syntax:

// newMap["key1"] - reads the value

// newMap["key1"] = 3 - writes the value. Can also
// be used to create new key value pairs in the map

// the zero value of map is nil

var nilMap map[string]int

// This is different than an empty map literal:

var notNil = map[string]int{}

// both have a length of zero, but we can read and
// write to notNil. We cannot read or write to nilMap

// Example of creating a map of slices:

var sliceMap = map[string][]string{
	"Names": []string{"Kevin", "Julia", "Aaron"},
	"Animals": []string{"Dog", "Cat", "Human"},
}


// If you know how many key-value pairs you intend to put in the map, but don’t know the exact values, you can use make to create a map with a default size:

var ages = make(map[int][]string, 10)
// Maps created with make still have a length of 0, and they can grow past the initially specified size.

// Maps are like slices in several ways:

// Maps automatically grow as you add key-value pairs to them.

// If you know how many key-value pairs you plan to insert into a map, you can use make to create a map with a specific initial size.

// Passing a map to the len function tells you the number of key-value pairs in a map.

// The zero value for a map is nil.

// Maps are not comparable. You can check if they are equal to nil, but you cannot check if two maps have identical keys and values using == or differ using !=.

// Note that reading from a map is not 0(1). Maps
// in go are set up as hash maps, and are actually
// arrays. The hash map uses a hashing algorithm
// to convert the key in the map into a number, which
// is then used to index into the array. Different
// keys can result in the same number. 
// 
// At the specified index in the array, a "bucket"
// (also an array) stores both the key and value. 
// A bucket can hold multiple keys and values.
// When reading from a map, the key is converted
// into a number, the index is found with that number
// and then all the keys in the bucket are iterated over
// until the correct one is found.

// Two keys mapping to the same bucket is known
// as a "collision" - it's best to avoid collisions
// when possible (they slow things down).