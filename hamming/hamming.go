package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	aRunes := []rune(a)
	bRunes := []rune(b)

	if len(aRunes) != len(bRunes) {
		return 0, errors.New("dna strands have different lengths")
	}

	distance := 0
	for idx := range aRunes {
		if aRunes[idx] != bRunes[idx] {
			distance++
		}
	}
	return distance, nil
}
