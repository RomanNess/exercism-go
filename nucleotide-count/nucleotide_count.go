package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA []rune

// Counts the nucleotides in a given DNA and provides a histogram.
func (dna DNA) Counts() (hist Histogram, err error) {
	hist = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, n := range dna {
		if count, ok := hist[n]; ok {
			hist[n] = count + 1
		} else {
			return nil, errors.New("strand with invalid nucleotide")
		}
	}

	return hist, nil
}
