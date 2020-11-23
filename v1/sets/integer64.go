package sets

/*
ToInteger64Slice generates a int64 slice of distinct int64i from the given int64i.

Example Usage:

 someInteger64i := []int64{"this", "is", "unique", "this", "is", "distinct"}
 result := Integer64(someInteger64i)
 result
 > []int64{"this", "is", "unique", "distinct"}
*/
func ToInteger64Slice(in ...[]int64) []int64 {
	set := NewInteger64(in...)
	res := set.ToSlice()

	return res
}

// Integer64 provides a type for int64 sets
type Integer64 map[int64]struct{}

// NewInteger64 creates a new set of int64 type
func NewInteger64(in ...[]int64) *Integer64 {
	set := make(Integer64, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Integer64) Add(in ...int64) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Integer64) Remove(element ...int64) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a int64 slice
func (i Integer64) ToSlice() []int64 {
	res := make([]int64, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Integer64) Empty() bool {
	return len(i) == 0
}
