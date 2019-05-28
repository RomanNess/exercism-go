package strand

import "strings"

// ToRNA return the RNA complement for a given DNA strand
func ToRNA(dna string) string {
	toRnaReplacer := strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	)
	return toRnaReplacer.Replace(dna);
}
