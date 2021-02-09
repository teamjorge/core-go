package slices

import (
	"errors"
	"fmt"
)

// Uinteger8 wraps a normal uint8 slice to provide
// additional helper methods
type Uinteger8 []uint8

// ForEach iterates each item in the given Uinteger8Slice and executes
// the given modifier function with it'i index and value
func (i Uinteger8) ForEach(modifier func(index int, val uint8)) {
	for index, value := range i {
		modifier(index, value)
	}
}

// Map iterates each item in the given Uinteger8Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Uinteger8.
// This method does modify the Uinteger8 in place.
func (i Uinteger8) Map(modifier func(index int, val uint8) uint8) Uinteger8 {
	res := make([]uint8, 0)
	for index, value := range i {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given Uinteger8Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the Uinteger8 in place and will return
// the modified version
func (i Uinteger8) Filter(modifier func(index int, val uint8) bool) Uinteger8 {
	if len(i) == 0 {
		return i
	}

	res := make([]uint8, 0)
	for index, value := range i {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given Uinteger8Slice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Uinteger8
func (i Uinteger8) Pop(index int) (uint8, Uinteger8, error) {
	var item uint8
	res := make([]uint8, 0)

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
func (i Uinteger8) Empty() bool {
	return len(i) == 0
}

// Contains determines whether the Uinteger8Slice contains the given value.
func (i Uinteger8) Contains(value uint8) bool {
	for _, val := range i {
		if val == value {
			return true
		}
	}

	return false
}
