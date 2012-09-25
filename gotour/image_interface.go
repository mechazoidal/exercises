package main

import (
	"image"
	"code.google.com/p/go-tour/pic"
)

// using http://code.google.com/p/go-wiki/wiki/GoForCPPProgrammers
// to try to grasp Interfaces

type Image struct{
	cm ColorModel
	b Bounds
	a At
}
func (img *Image) ColorModel() RGBA64 {
	return img.cm
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}