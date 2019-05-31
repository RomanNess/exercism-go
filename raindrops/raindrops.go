package raindrops

import "strconv"

// Convert provides a raindrop sound based on the factors of a given number
func Convert(n int) (result string) {
	for _, e := range factorsToSound {
		if (n % e.factor) == 0 {
			result += e.sound
		}
	}

	if len(result) == 0 {
		return strconv.Itoa(n)	// default
	}
	return result
}

type fToS struct {
	factor int
	sound string
}

var factorsToSound = [...]fToS{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}
