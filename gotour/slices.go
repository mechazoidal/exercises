package main

import "code.google.com/p/go-tour/pic"
// had to go out for this one :(
// http://stackoverflow.com/questions/11585008/go-tour-excerciseslices-index-out-of-range
func Pic(dx, dy int) [][]uint8 {
 	// Secret here: 'make' takes a length, AND A CAPACITY
	// so you don't initialize the whole thing here, you init each item(to another slice) on each step of the loop
	// else you get runtime panics about being out of range
	pixels := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pixels[y] = make([]uint8,dx)
		for x := 0; x < dx; x++ {
			// the formula here will determine what actually shows up
			pixels[y][x] = uint8(x * y)
		}
	}
	return pixels
}

func main() {
// Note that Show expects a _function itself_ to be passed in, NOT the result of one
// (looking at the source code, it calls the func with x=256, y=256)
	pic.Show(Pic)
}
