package factorial

import ()

func Factorial(input int) int {
	if input <= 1 {
		return 1
	}
	return input * Factorial(input-1)
}
