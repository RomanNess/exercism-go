// Package bob provides a simulation of a lackadaisical teenager
package bob

import (
	"strings"
)

// Hey provides a typical teenager comment for a given remark
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	remarkIsLoud := containsNoLowerCase(remark)
	remarkIsQuestion := endsWithQuestionMark(remark)

	if remarkIsQuestion && remarkIsLoud {
		return "Calm down, I know what I'm doing!"
	}
	if remarkIsLoud {
		return "Whoa, chill out!"
	}
	if remarkIsQuestion {
		return "Sure."
	}
	return "Whatever."
}

func containsNoLowerCase(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}

func endsWithQuestionMark(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
