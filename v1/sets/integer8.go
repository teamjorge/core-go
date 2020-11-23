package sets

/*
ToInteger8Slice generates a int8 slice of distinct int8i from the given int8i.

Example Usage:

 someInteger8i := []int8{"this", "is", "unique", "this", "is", "distinct"}
 result := Integer8(someInteger8i)
 result
 > []int8{"this", "is", "unique", "distinct"}
*/
func ToInteger8Slice(in ...[]int8) []int8 {
	set := NewInteger8(in...)
	res := set.ToSlice()

	return res
}

// Integer8 provides a type for int8 sets
type Integer8 map[int8]struct{}

// NewInteger8 creates a new set of int8 type
func NewInteger8(in ...[]int8) *Integer8 {
	set := make(Integer8, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Integer8) Add(in ...int8) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Integer8) Remove(element ...int8) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a int8 slice
func (i Integer8) ToSlice() []int8 {
	res := make([]int8, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Integer8) Empty() bool {
	return len(i) == 0
}
