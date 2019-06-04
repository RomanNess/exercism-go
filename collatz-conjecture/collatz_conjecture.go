package collatzconjecture

import (
	"errors"
)

// CollatzConjecture calculates the number of steps in the 3n + 1 problem for a given positive number
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("accepts positive integers only")
	}
	// won't return errors for any positive integer
	return collatzConjectureFailsafe(n), nil
}

func collatzConjectureFailsafe(n int) int {
	if n == 1 {
		return 0
	}

	if (n % 2) == 0 {
		return 1 + collatzConjectureFailsafe(n/2)
	}
	return 1 + collatzConjectureFailsafe(3*n+1)
}
