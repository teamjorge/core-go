package sets

/*
ToUintegerSlice generates a uint slice of distinct uinti from the given uinti.

Example Usage:

 someUintegeri := []uint{"this", "is", "unique", "this", "is", "distinct"}
 result := Uinteger(someUintegeri)
 result
 > []uint{"this", "is", "unique", "distinct"}
*/
func ToUintegerSlice(in ...[]uint) []uint {
	set := NewUinteger(in...)
	res := set.ToSlice()

	return res
}

// Uinteger provides a type for uint sets
type Uinteger map[uint]struct{}

// NewUinteger creates a new set of uint type
func NewUinteger(in ...[]uint) *Uinteger {
	set := make(Uinteger, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element(i) to the set
func (i Uinteger) Add(in ...uint) {
	for _, val := range in {
		i[val] = struct{}{}
	}
}

// Remove element(i) from the set
func (i Uinteger) Remove(element ...uint) {
	for _, elem := range element {
		delete(i, elem)
	}
}

// ToSlice returns the values in the set as a uint slice
func (i Uinteger) ToSlice() []uint {
	res := make([]uint, 0)
	for key := range i {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func (i Uinteger) Empty() bool {
	return len(i) == 0
}
