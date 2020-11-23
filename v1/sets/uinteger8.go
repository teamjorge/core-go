package sets

/*
ToUinteger8Slice generates a uint8 slice of distinct uint8i from the given uint8i.

Example Usage:

 someUinteger8i := []uint8{"this", "is", "unique", "this", "is", "distinct"}
 result := Uinteger8(someUinteger8i)
 result
 > []uint8{"this", "is", "unique", "distinct"}
*/
func ToUinteger8Slice(in ...[]uint8) []uint8 {
	set := NewUinteger8(in...)
	res := set.ToSlice()

	return res
}

// Uinteger8 provides a type for uint8 sets
type Uinteger8 map[uint8]struct{}

// NewUinteger8 creates a new set of uint8 type
func NewUinteger8(in ...[]uint8) *Uinteger8 {
	set := make(Uinteger8, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Uinteger8) Add(in ...uint8) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Uinteger8) Remove(element ...uint8) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a uint8 slice
func (i Uinteger8) ToSlice() []uint8 {
	res := make([]uint8, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Uinteger8) Empty() bool {
	return len(i) == 0
}
