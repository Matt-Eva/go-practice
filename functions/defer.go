package main

import "fmt"

func main(){

	fmt.Println(changeMyReturn(7))

}

// DEFER

// You can set pieces of functionality in your function
// to run after the function itself ends using the defer keyword.
// One example of this is closing a file that was opened by the function
// in order to cleanup up the temporary resource created by the program.
// You can simply tag the close() method with the defer keyword to 
// indicate that you want it to run when your function finishes running


// You can defer closures within a function. However,
// this means that any return value that function might have
// will be unreadable.

// If you use named return values within a function
// a closure within that function is actually able to modify
// the return value after the function itself has run.

// Ex:

func changeMyReturn(num int) (result int){
	defer func(){
		result = 8
	}()
	return num // doesn't matter that you're return something other than the named return value - closure still modifies it.
}

// When using the defer keyword, we need to be sure to actually
// invoke the function we want to defer.