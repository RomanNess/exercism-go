// Package leap provides leap year utilities
package leap

// IsLeapYear tells whether a year is a leap year
func IsLeapYear(year int) bool {
	if (year % 400) == 0 {
		return true
	}
	if (year % 100) == 0 {
		return false
	}
	return (year % 4) == 0
}
