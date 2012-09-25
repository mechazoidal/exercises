package main

import (
	"fmt";
	"math";
)

type ErrNegativeSqrt float64 


func (e ErrNegativeSqrt) Error() string {
	msg := ""
	if e < 0 {
		msg = fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
		//msg = fmt.Sprintf(float64(e))
		//msg = fmt.Print("cannot Sqrt negative number")
	}
	return msg
}
	
	
func Sqrt(f float64) (float64, error) {
// Not totally happy with this. Unsure if the Error func should do the <0 check, or do it here.
	e := ErrNegativeSqrt(f)
	
	
	last_z, z := f, 1.0
	if(f > 0) {
		for math.Abs(z - last_z) >= 1.0e-6 {
			last_z, z = z, z - (z * z - f) / (2 * z);
		}
	}
	return z, e
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}