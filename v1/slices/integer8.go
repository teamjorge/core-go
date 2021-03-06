package slices

import (
	"errors"
	"fmt"
)

// Integer8 wraps a normal int8 slice to provide
// additional helper methods
type Integer8 []int8

// ForEach iterates each item in the given Integer8Slice and executes
// the given modifier function with it'i index and value
func (i Integer8) ForEach(modifier func(index int, val int8)) {
	for index, value := range i {
		modifier(index, value)
	}
}

// Map iterates each item in the given Integer8Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Integer8.
// This method does modify the Integer8 in place.
func (i Integer8) Map(modifier func(index int, val int8) int8) Integer8 {
	res := make([]int8, 0)
	for index, value := range i {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given Integer8Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the Integer8 in place and will return
// the modified version
func (i Integer8) Filter(modifier func(index int, val int8) bool) Integer8 {
	if len(i) == 0 {
		return i
	}

	res := make([]int8, 0)
	for index, value := range i {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given Integer8Slice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Integer8
func (i Integer8) Pop(index int) (int8, Integer8, error) {
	var item int8
	res := make([]int8, 0)

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
func (i Integer8) Empty() bool {
	return len(i) == 0
}

// Contains determines whether the Integer8Slice contains the given value.
func (i Integer8) Contains(value int8) bool {
	for _, val := range i {
		if val == value {
			return true
		}
	}

	return false
}
