package practice8

import (
	"errors"
)

func Sum(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("You can't divide by zero")
	} else {
		return a / b, nil
	}
}
