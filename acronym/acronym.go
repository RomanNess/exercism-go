package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate provides the acronym for a given string
func Abbreviate(s string) string {
	regex, _ := regexp.Compile("([[:alpha:]])[A-Za-z']*[^A-Za-z']*")
	return strings.ToUpper(regex.ReplaceAllString(s, "$1"))
}
