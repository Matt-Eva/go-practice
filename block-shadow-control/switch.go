package main

import "fmt"

func main(){
	words := []string{"a", "cow", "smile", "gopher",
    "octopus", "anthropologist"}
for _, word := range words {
    switch size := len(word); size { // can declare switch variable inline (size in this case)
    case 1, 2, 3, 4:
        fmt.Println(word, "is a short word!")
    case 5:
        wordLen := len(word)
        fmt.Println(word, "is exactly the right length:", wordLen)
    case 6, 7, 8, 9:
    default:
        fmt.Println(word, "is a long word!")
    }
}

moreWords := []string{"hi", "salutations", "hello"}
for _, word := range moreWords {
    switch wordLen := len(word); { // haven't declared value for comparison
    case wordLen < 5:
        fmt.Println(word, "is a short word!")
    case wordLen > 10:
        fmt.Println(word, "is a long word!")
    default:
        fmt.Println(word, "is exactly the right length.")
    }
}
}

// Note - you can use a break statement to exit a case.
// Keep this in mind when using a switch statement inside
// a for loop. If you want to exit the outer for loop from
// within the switch statement, give your outer for
// loop a label.

// BLANK SWITCHES
// You write a blank switch by not specifying the value
// for comparison. This allows you to use any comparison
// operation, not just strict equality (as with regular
// switch statements).