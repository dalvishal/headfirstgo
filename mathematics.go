package mathematics

import (
	"fmt"
)

func Add(first float64 ,second float64) float64 {
	return first + second
}

func Subtract( first float64 ,second float64) float64 {
	return first - second
}

func Multiply( first float64 ,second float64) float64 {
	return first * second
}

func Divide( first float64 ,second float64) (float64,error) {
	
	if second == 0 {
		return 0, fmt.Errorf("Division by %f is not allowed\n", second)
	}
	return first / second, nil
}