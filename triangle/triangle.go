// Package triangle provides triangle utility methods
package triangle

import (
	"errors"
	"math"
	"sort"
)

type Kind int

const (
    NaT = iota// not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides determines the kind of the triangle from the length of its sides
func KindFromSides(a, b, c float64) Kind {

	var err error
	a, b, c, err = validateAndSortTriangleSides(a, b, c)
	if err != nil || a + b < c {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

func validateAndSortTriangleSides(a float64, b float64, c float64) (float64, float64, float64, error) {
	sides := []float64{a, b, c}
	for _, side := range sides {
		if triangleSideIsInvalid(side) {
			return 0, 0, 0, errors.New("at least one triangle side is invalid")
		}
	}
	sort.Float64s(sides)
	return sides[0], sides[1], sides[2], nil
}

func triangleSideIsInvalid(side float64) bool {
	return side <= 0 || math.IsNaN(side) || math.IsInf(side, 0)
}
