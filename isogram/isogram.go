package isogram

import (
	"unicode"
)

// IsIsogram determines whether each letter appears once at most in a given string
func IsIsogram(input string) bool {

	discoveredRunes := make(map[rune]bool)

	for _, r := range input {

		// skip on some runes
		if r == '-' || r == ' ' {
			continue
		}

		rLower := unicode.ToLower(r)
		if discoveredRunes[rLower] {
			return false
		}
		discoveredRunes[rLower] = true
	}

	return true
}