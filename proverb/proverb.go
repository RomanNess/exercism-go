// Package proverb provides a proverbial rhyme generator
package proverb

import "fmt"

// Proverb provides proverbial rhymes for a list of nouns
func Proverb(nouns []string) (rhymes []string) {

	if len(nouns) == 0 {
		return rhymes
	}

	for i := 0; i < len(nouns)-1; i++ {
		appendProverb(&rhymes, nouns[i:])
	}

	appendLastProverb(&rhymes, nouns[0])
	return
}

func appendProverb(rhymes *[]string, nouns []string) {
	*rhymes = append(*rhymes, fmt.Sprintf("For want of a %s the %s was lost.", nouns[0], nouns[1]))
}

func appendLastProverb(rhymes *[]string, noun string) {
	*rhymes = append(*rhymes, fmt.Sprintf("And all for the want of a %s.", noun))
}
