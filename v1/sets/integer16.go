package sets

/*
ToInteger16Slice generates a int16 slice of distinct int16i from the given int16i.

Example Usage:

 someInteger16i := []int16{"this", "is", "unique", "this", "is", "distinct"}
 result := Integer16(someInteger16i)
 result
 > []int16{"this", "is", "unique", "distinct"}
*/
func ToInteger16Slice(in ...[]int16) []int16 {
	set := NewInteger16(in...)
	res := set.ToSlice()

	return res
}

// Integer16 provides a type for int16 sets
type Integer16 map[int16]struct{}

// NewInteger16 creates a new set of int16 type
func NewInteger16(in ...[]int16) *Integer16 {
	set := make(Integer16, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Integer16) Add(in ...int16) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Integer16) Remove(element ...int16) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a int16 slice
func (i Integer16) ToSlice() []int16 {
	res := make([]int16, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Integer16) Empty() bool {
	return len(i) == 0
}
