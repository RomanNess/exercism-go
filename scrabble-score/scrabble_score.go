package scrabble

import (
	"unicode"
)

// Score calculates the Scrabble score for a given word
func Score(word string) (score int) {
	for _, r := range word {
		score += scoreForRune(unicode.ToLower(r))
	}
	return score
}

func scoreForRune(r rune) int {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	}
	return 0
}
