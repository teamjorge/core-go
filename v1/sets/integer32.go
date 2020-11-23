package sets

/*
ToInteger32Slice generates a int32 slice of distinct int32i from the given int32i.

Example Usage:

 someInteger32i := []int32{"this", "is", "unique", "this", "is", "distinct"}
 result := Integer32(someInteger32i)
 result
 > []int32{"this", "is", "unique", "distinct"}
*/
func ToInteger32Slice(in ...[]int32) []int32 {
	set := NewInteger32(in...)
	res := set.ToSlice()

	return res
}

// Integer32 provides a type for int32 sets
type Integer32 map[int32]struct{}

// NewInteger32 creates a new set of int32 type
func NewInteger32(in ...[]int32) *Integer32 {
	set := make(Integer32, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Integer32) Add(in ...int32) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Integer32) Remove(element ...int32) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a int32 slice
func (i Integer32) ToSlice() []int32 {
	res := make([]int32, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Integer32) Empty() bool {
	return len(i) == 0
}
