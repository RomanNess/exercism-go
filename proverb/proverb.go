// Package proverb provides a proverbial rhyme generator
package proverb

import "fmt"

// Proverb provides proverbial rhymes for a list of nouns
func Proverb(rhyme []string) (phrases []string) {

	if len(rhyme) == 0 {
		return phrases
	}

	for i := 1; i < len(rhyme); i++ {
		phrases = append(phrases, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}

	phrases = append(phrases, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return
}
