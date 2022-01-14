package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", *e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, &e
	}
	var z float64 = 1
	var prev = float64(-1000)

	for {
		//fmt.Println(z)
		if prev == z {
			break
		} else {
			prev = z
			z -= (z*z - x) / (2 * z)
		}
	}

	return z, nil
}
func main() {
	if num, err := Sqrt(49); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
	if num, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
	// These checks for errors aren't necessary
	// fmt.Println(Sqrt(-2)) would print the error if there is one
	fmt.Println("Done!")
}
