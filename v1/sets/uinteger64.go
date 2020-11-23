package sets

/*
ToUinteger64Slice generates a uint64 slice of distinct uint64i from the given uint64i.

Example Usage:

 someUinteger64i := []uint64{"this", "is", "unique", "this", "is", "distinct"}
 result := Uinteger64(someUinteger64i)
 result
 > []uint64{"this", "is", "unique", "distinct"}
*/
func ToUinteger64Slice(in ...[]uint64) []uint64 {
	set := NewUinteger64(in...)
	res := set.ToSlice()

	return res
}

// Uinteger64 provides a type for uint64 sets
type Uinteger64 map[uint64]struct{}

// NewUinteger64 creates a new set of uint64 type
func NewUinteger64(in ...[]uint64) *Uinteger64 {
	set := make(Uinteger64, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Uinteger64) Add(in ...uint64) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Uinteger64) Remove(element ...uint64) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a uint64 slice
func (i Uinteger64) ToSlice() []uint64 {
	res := make([]uint64, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Uinteger64) Empty() bool {
	return len(i) == 0
}
