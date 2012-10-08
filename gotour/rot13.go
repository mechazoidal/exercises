package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
// Instead of farting with runes, slices, or strconv,
// just define a function that works on a single byte
// (and probably doesn't work on runes, but whatevs)
func rot13(b byte) byte {
	switch {
		case b >= 'A' && b <= 'Z':
	        	return ('A' + (b-'A'+13)%26)
	   	case b >= 'a' && b <= 'z':
        		return ('a' + (b-'a'+13)%26)
	}
	return b
}
// Original solution from https://syslog.warten.de/2011/11/solutions-for-tour-of-go-exercises/
// used os.Error, which is changed to a default system 'error' type in Go 1
// Note also that you need to provide the reader itself as a reference in the function declaration
func (b *rot13Reader) Read(p []byte) (n int, err error){
	n, err = b.r.Read(p)
	// no idea where this '[:n]' syntax comes from, don't see it in Effective Go
	for i:=range(p[:n]) {
		p[i]=rot13(p[i])
	}
	return
}
func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}