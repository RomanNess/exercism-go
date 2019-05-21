package hamming

import "errors"

// Distance calculates the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("dna strands have different lengths")
	}

	distance := 0
	for idx := range a {
		if a[idx] != b[idx] {
			distance++
		}
	}
	return distance, nil
}
