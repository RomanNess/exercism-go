package raindrops

import "strconv"

// Convert provides a raindrop sound based on the factors of a given number
func Convert(n int) (result string) {

	for i := 1; i <= n; i++ {
		if isFactorOf(i, n) {
			result += numberToRaindrop(i)
		}
	}

	if len(result) > 0 {
		return result
	}
	return strconv.Itoa(n)
}

func isFactorOf(i int, n int) bool {
	return (n % i) == 0
}

func numberToRaindrop(i int) string {
	switch i {
	case 3:
		return "Pling"
	case 5:
		return "Plang"
	case 7:
		return "Plong"
	default:
		return ""
	}
}
