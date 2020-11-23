package sets

/*
ToIntegerSlice generates a int slice of distinct inti from the given inti.

Example Usage:

 someIntegeri := []int{"this", "is", "unique", "this", "is", "distinct"}
 result := Integer(someIntegeri)
 result
 > []int{"this", "is", "unique", "distinct"}
*/
func ToIntegerSlice(in ...[]int) []int {
	set := NewInteger(in...)
	res := set.ToSlice()

	return res
}

// Integer provides a type for int sets
type Integer map[int]struct{}

// NewInteger creates a new set of int type
func NewInteger(in ...[]int) *Integer {
	set := make(Integer, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Integer) Add(in ...int) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Integer) Remove(element ...int) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a int slice
func (i Integer) ToSlice() []int {
	res := make([]int, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Integer) Empty() bool {
	return len(i) == 0
}
