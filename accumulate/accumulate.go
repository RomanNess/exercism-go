package accumulate

// Accumulate executes a given operation on any element in a given slice
func Accumulate(list []string, converter func(string) string) []string {
	for i, e := range list {
		list[i] = converter(e)
	}
	return list
}
