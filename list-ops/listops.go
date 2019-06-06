package listops

// IntList wraps a slice of integers
type IntList []int

type predFunc func(int) bool
type unaryFunc func(int) int
type binFunc func(int, int) int

// Length returns the number of elements in an IntList
func (list IntList) Length() (len int) {
	// without using len(list)
	for range list {
		len++
	}
	return len
}

// Reverse provides the reverse of an IntList
func (list IntList) Reverse() IntList {
	for i, j := 0, list.Length()-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// Append appends a given list to an IntList
func (list IntList) Append(other IntList) IntList {

	lLength := list.Length()
	otherLength := other.Length()

	result := make(IntList, lLength+otherLength)
	copy(result, list)

	for i := 0; i < otherLength; i++ {
		result[lLength+i] = other[i]
	}
	return result
}

// Concat appends a given list of lists to an IntList
func (list IntList) Concat(otherLists []IntList) (result IntList) {
	result = list
	for _, otherList := range otherLists {
		result = result.Append(otherList)
	}
	return result
}

// Filter provides a list with only the elements that match a given predicate
func (list IntList) Filter(pred predFunc) (result IntList) {
	result = IntList{}
	for _, e := range list {
		if pred(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies a given function to every element of an IntList
func (list IntList) Map(fn unaryFunc) (result IntList) {
	result = make(IntList, list.Length())
	copy(result, list)

	for i, e := range list {
		result[i] = fn(e)
	}
	return result
}

// Foldr performs a right fold with a given function on an IntList
func (list IntList) Foldr(fn binFunc, init int) int {
	if list.Length() == 0 {
		return init
	}
	return list[:list.Length()-1].Foldr(fn, fn(list[list.Length()-1], init))
}

// Foldl performs a left fold with a given function on an IntList
func (list IntList) Foldl(fn binFunc, init int) int {
	if list.Length() == 0 {
		return init
	}
	return list[1:].Foldl(fn, fn(init, list[0]))
}
