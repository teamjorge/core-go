package slices

import (
	"errors"
	"fmt"
)

// Integer wraps a normal int slice to provide
// additional helper methods
type Integer []int

// ForEach iterates each item in the given IntegerSlice and executes
// the given modifier function with it'i index and value
func (i Integer) ForEach(modifier func(index int, val int)) {
	for index, value := range i {
		modifier(index, value)
	}
}

// Map iterates each item in the given IntegerSlice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Integer.
// This method does modify the Integer in place.
func (i Integer) Map(modifier func(index int, val int) int) Integer {
	res := make([]int, 0)
	for index, value := range i {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given IntegerSlice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the Integer in place and will return
// the modified version
func (i Integer) Filter(modifier func(index int, val int) bool) Integer {
	if len(i) == 0 {
		return i
	}

	res := make([]int, 0)
	for index, value := range i {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given IntegerSlice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Integer
func (i Integer) Pop(index int) (int, Integer, error) {
	var item int
	res := make([]int, 0)

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
func (i Integer) Empty() bool {
	return len(i) == 0
}
