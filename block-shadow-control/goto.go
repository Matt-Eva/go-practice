package main

import "fmt"
import "math/rand"

func main(){

	// 	a := 10
	// 	goto skip
	// 	b := 20
	// skip:
	// 	c := 30
	// 	fmt.Println(a, b, c)
	// 	if c > a {
	// 		goto inner
	// 	}
	// 	if a < b {
	// 	inner:
	// 		fmt.Println("a is less than b")
	// 	}
	// ^^ This produces the following errors:
	// goto skip jumps over declaration of b at ./main.go:8:4
    // goto inner jumps into block starting at ./main.go:15:11

    a := rand.Intn(10)
    for a < 100 {
        if a%5 == 0 {
            goto done
        }
        a = a*2 + 1
    }
    fmt.Println("do something when the loop completes normally")
done:
    fmt.Println("do complicated stuff no matter why we left the loop")
    fmt.Println(a)

	// a contrived but valid use case for goto

}

// GOTO

// The goto statement allows you to jump to any point
// in your code. It is generally considered dangerous
// and has been removed from many modern programming
// languages. However, Go has made some modifications
// for making it safer, and it can be used in specific
// circumstances.


// For example, Go forbids jumps that skip over
// variable declaration and jumps that go into an
// inner or parallel block.


// Sometimes goto makes your code cleaner and more legible
// however, you'll want to avoid using goto more often
// than not.