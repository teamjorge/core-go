package slices

import (
	"errors"
	"fmt"
)

// Integer64 wraps a normal int64 slice to provide
// additional helper methods
type Integer64 []int64

// ForEach iterates each item in the given Integer64Slice and executes
// the given modifier function with it'i index and value
func (i Integer64) ForEach(modifier func(index int, val int64)) {
	for index, value := range i {
		modifier(index, value)
	}
}

// Map iterates each item in the given Integer64Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Integer64.
// This method does modify the Integer64 in place.
func (i Integer64) Map(modifier func(index int, val int64) int64) Integer64 {
	res := make([]int64, 0)
	for index, value := range i {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given Integer64Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the Integer64 in place and will return
// the modified version
func (i Integer64) Filter(modifier func(index int, val int64) bool) Integer64 {
	if len(i) == 0 {
		return i
	}

	res := make([]int64, 0)
	for index, value := range i {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given Integer64Slice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Integer64
func (i Integer64) Pop(index int) (int64, Integer64, error) {
	var item int64
	res := make([]int64, 0)

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
func (i Integer64) Empty() bool {
	return len(i) == 0
}

// Contains determines whether the Integer64Slice contains the given value.
func (i Integer64) Contains(value int64) bool {
	for _, val := range i {
		if val == value {
			return true
		}
	}

	return false
}
