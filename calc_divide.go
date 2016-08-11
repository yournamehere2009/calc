package calc

import (
	"errors"
)

// Divide takes two numbers and adds them together
func Divide(a float64, b float64) (float64, error) {

    if b != 0 {
        return a/b, nil;
    }

    return 0, errors.New("Why you trying to divide by zero fucker?");
}