package sets

/*
ToUinteger32Slice generates a uint32 slice of distinct uint32i from the given uint32i.

Example Usage:

 someUinteger32i := []uint32{"this", "is", "unique", "this", "is", "distinct"}
 result := Uinteger32(someUinteger32i)
 result
 > []uint32{"this", "is", "unique", "distinct"}
*/
func ToUinteger32Slice(in ...[]uint32) []uint32 {
	set := NewUinteger32(in...)
	res := set.ToSlice()

	return res
}

// Uinteger32 provides a type for uint32 sets
type Uinteger32 map[uint32]struct{}

// NewUinteger32 creates a new set of uint32 type
func NewUinteger32(in ...[]uint32) *Uinteger32 {
	set := make(Uinteger32, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Uinteger32) Add(in ...uint32) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Uinteger32) Remove(element ...uint32) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a uint32 slice
func (i Uinteger32) ToSlice() []uint32 {
	res := make([]uint32, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Uinteger32) Empty() bool {
	return len(i) == 0
}
