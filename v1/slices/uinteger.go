package slices

import (
	"errors"
	"fmt"
)

// Uinteger wraps a normal uint slice to provide
// additional helper methods
type Uinteger []uint

// ForEach iterates each item in the given UintegerSlice and executes
// the given modifier function with it'i index and value
func (i Uinteger) ForEach(modifier func(index int, val uint)) {
	for index, value := range i {
		modifier(index, value)
	}
}

// Map iterates each item in the given UintegerSlice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Uinteger.
// This method does modify the Uinteger in place.
func (i Uinteger) Map(modifier func(index int, val uint) uint) Uinteger {
	res := make([]uint, 0)
	for index, value := range i {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given UintegerSlice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the Uinteger in place and will return
// the modified version
func (i Uinteger) Filter(modifier func(index int, val uint) bool) Uinteger {
	if len(i) == 0 {
		return i
	}

	res := make([]uint, 0)
	for index, value := range i {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given UintegerSlice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Uinteger
func (i Uinteger) Pop(index int) (uint, Uinteger, error) {
	var item uint
	res := make([]uint, 0)

	if len(i) == 0 {
		err := "Pop on empty slice failed"
		return item, res, errors.New(err)
	}

	if index < 0 || index > (len(i)-1) {
		err := fmt.Sprintf("Pop on index %d not available on slice of length %d", index, len(i))
		return item, res, errors.New(err)
	}

	item = i[index]
	res = append(i[:index], i[index+1:]...)

	return item, res, nil
}

// Empty determines whether the slice is empty
func (i Uinteger) Empty() bool {
	return len(i) == 0
}
