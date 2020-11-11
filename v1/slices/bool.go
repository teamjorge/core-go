package slices

import (
	"errors"
	"fmt"
)

// Boolean wraps a normal bool slice to provide
// additional helper methods.
type Boolean []bool

// ForEach iterates each item in the given BooleanSlice and executes
// the given modifier function with it's index and value.
func (b Boolean) ForEach(modifier func(index int, val bool)) {
	for index, value := range b {
		modifier(index, value)
	}
}

// Map iterates each item in the given BooleanSlice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Boolean.
//
// This method does modify the Boolean in place.
func (b Boolean) Map(modifier func(index int, val bool) bool) Boolean {
	res := make([]bool, 0)
	for index, value := range b {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given BooleanSlice.
//
// On each iteration, the current index and value will be passed to the
// modifier function. Only iterations that returned true when passed to
// the modifier thing will be returned.
//
// This method does not modify the Boolean in place and will return
// the modified version.
func (b Boolean) Filter(modifier func(index int, val bool) bool) Boolean {
	if len(b) == 0 {
		return b
	}

	res := make([]bool, 0)
	for index, value := range b {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given BooleanSlice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Boolean.
func (b Boolean) Pop(index int) (bool, Boolean, error) {
	var item bool
	res := make([]bool, 0)

	if len(b) == 0 {
		err := "Pop on empty slice failed"
		return item, res, errors.New(err)
	}

	if index < 0 || index > (len(b)-1) {
		err := fmt.Sprintf("Pop on index %d not available on slice of length %d", index, len(b))
		return item, res, errors.New(err)
	}

	item = b[index]
	res = append(b[:index], b[index+1:]...)

	return item, res, nil
}

// Any determines if any of the indices in the given Boolean have a
// value of true.
//
// Empty slices will always return false.
func Any(in Boolean) bool {
	res := false
	for _, val := range in {
		if val {
			return true
		}
	}
	return res
}

// All determines if all of the values in the given Boolean have a
// value of true.
//
// Empty slices will always return True.
func All(in Boolean) bool {
	res := true
	for _, val := range in {
		if !val {
			return false
		}
	}
	return res
}

// Empty determines whether the slice is empty
func (b Boolean) Empty() bool {
	return len(b) == 0
}
