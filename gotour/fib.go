package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
// Had to again look at outside. :(
// http://codingfu.com/blog/first-impression-of-go/
// In this case, I didn't remember that there's an iterative solution for Fibonacci numbers.
// And go's syntax allows you to do multiple assignment.
	a,b := 0,1
	return func() int {
		a,b = b, a+1
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}