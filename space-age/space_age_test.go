package space

import (
	"math"
	"testing"
)

func TestAge(t *testing.T) {
	const precision = 0.01
	for _, tc := range testCases {
		actual, err := Age(tc.seconds, tc.planet)
		if err != nil {
			t.Errorf("FAIL: Expected no error\nActual: %v", err)
		}
		if math.IsNaN(actual) || math.Abs(actual-tc.expected) > precision {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestErrorOnUnknownPlanet(t *testing.T) {
	_, err := Age(100000.0, "Pluto")
	expectedErrString := "Pluto is no planet."
	if err == nil {
		t.Fatalf("FAIL: Expected: %#v\nActual: %#v", expectedErrString, err)
	}
	if err.Error() != expectedErrString {
		t.Fatalf("FAIL: Expected: %#v\nActual: %#v", expectedErrString, err)
	}
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Age(tc.seconds, tc.planet)
		}
	}
}
