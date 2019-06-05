package listops

type IntList []int

type predFunc func(int) bool
type unaryFunc func(int) int
type binFunc func(int, int) int


func (IntList) Length() int {
	return -1
}

func (IntList) Reverse() IntList {
	return nil
}

func (IntList) Append(list IntList) IntList {
	return nil
}

func (IntList) Concat(lists []IntList) IntList {
	return nil
}

func (IntList) Filter(condition predFunc) IntList {
	return nil
}

func (IntList) Map(condition unaryFunc) IntList {
	return nil
}

func (IntList) Foldr(condition binFunc, init int) int {
	return -1
}

func (IntList) Foldl(condition binFunc, init int) int {
	return -1
}
