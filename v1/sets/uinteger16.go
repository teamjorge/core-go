package sets

/*
ToUinteger16Slice generates a uint16 slice of distinct uint16i from the given uint16i.

Example Usage:

 someUinteger16i := []uint16{"this", "is", "unique", "this", "is", "distinct"}
 result := Uinteger16(someUinteger16i)
 result
 > []uint16{"this", "is", "unique", "distinct"}
*/
func ToUinteger16Slice(in ...[]uint16) []uint16 {
	set := NewUinteger16(in...)
	res := set.ToSlice()

	return res
}

// Uinteger16 provides a type for uint16 sets
type Uinteger16 map[uint16]struct{}

// NewUinteger16 creates a new set of uint16 type
func NewUinteger16(in ...[]uint16) *Uinteger16 {
	set := make(Uinteger16, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Uinteger16) Add(in ...uint16) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Uinteger16) Remove(element ...uint16) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a uint16 slice
func (i Uinteger16) ToSlice() []uint16 {
	res := make([]uint16, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Uinteger16) Empty() bool {
	return len(i) == 0
}
