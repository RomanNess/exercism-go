// Package proverb provides a proverbial rhyme generator
package proverb

import "fmt"

// Proverb provides proverbial rhymes for a list of nouns
func Proverb(nouns []string) (rhymes []string) {

	if len(nouns) == 0 {
		return rhymes
	}

	rhymes = make([]string, len(nouns))
	for i := 0; i < len(nouns)-1; i++ {
		rhymes[i] = fmt.Sprintf("For want of a %s the %s was lost.", nouns[i], nouns[i+1])
	}

	rhymes[len(nouns)-1] = fmt.Sprintf("And all for the want of a %s.", nouns[0])
	return
}
