package main

import "fmt"

func main(){
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	if q := rand.Intn(10); q == 0 {
		fmt.Println("That's too low")
	} else if q > 5 {
		fmt.Println("That's too big:", q)
	} else {
		fmt.Println("That's a good number:", q)
	}
	//fmt.Println(q) <-- throws error; q is only defined in if / else block
}

// IF STATEMENTS
// If statements in Go are basically identical to if
// statements in JavaScript, except we don't put
// parentheses around the condition.

// n := rand.Intn(10)
// if n == 0 {
//     fmt.Println("That's too low")
// } else if n > 5 {
//     fmt.Println("That's too big:", n)
// } else {
//     fmt.Println("That's a good number:", n)
// }

// One cool feature of Go is that it allows
// us to declare variables that are scoped
// to a specific if / else block:

// if n := rand.Intn(10); n == 0 {
//     fmt.Println("That's too low")
// } else if n > 5 {
//     fmt.Println("That's too big:", n)
// } else {
//     fmt.Println("That's a good number:", n)
// }
