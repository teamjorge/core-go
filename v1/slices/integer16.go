package slices

import (
	"errors"
	"fmt"
)

// Integer16 wraps a normal int16 slice to provide
// additional helper methods
type Integer16 []int16

// ForEach iterates each item in the given Integer16Slice and executes
// the given modifier function with it'i index and value
func (i Integer16) ForEach(modifier func(index int, val int16)) {
	for index, value := range i {
		modifier(index, value)
	}
}

// Map iterates each item in the given Integer16Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Integer16.
// This method does modify the Integer16 in place.
func (i Integer16) Map(modifier func(index int, val int16) int16) Integer16 {
	res := make([]int16, 0)
	for index, value := range i {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given Integer16Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the Integer16 in place and will return
// the modified version
func (i Integer16) Filter(modifier func(index int, val int16) bool) Integer16 {
	if len(i) == 0 {
		return i
	}

	res := make([]int16, 0)
	for index, value := range i {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given Integer16Slice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Integer16
func (i Integer16) Pop(index int) (int16, Integer16, error) {
	var item int16
	res := make([]int16, 0)

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
func (i Integer16) Empty() bool {
	return len(i) == 0
}
