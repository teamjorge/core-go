package slices

import (
	"errors"
	"fmt"
)

// Integer32 wraps a normal int32 slice to provide
// additional helper methods
type Integer32 []int32

// ForEach iterates each item in the given Integer32Slice and executes
// the given modifier function with it'i index and value
func (i Integer32) ForEach(modifier func(index int, val int32)) {
	for index, value := range i {
		modifier(index, value)
	}
}

// Map iterates each item in the given Integer32Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Integer32.
// This method does modify the Integer32 in place.
func (i Integer32) Map(modifier func(index int, val int32) int32) Integer32 {
	res := make([]int32, 0)
	for index, value := range i {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given Integer32Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the Integer32 in place and will return
// the modified version
func (i Integer32) Filter(modifier func(index int, val int32) bool) Integer32 {
	if len(i) == 0 {
		return i
	}

	res := make([]int32, 0)
	for index, value := range i {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given Integer32Slice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given Integer32
func (i Integer32) Pop(index int) (int32, Integer32, error) {
	var item int32
	res := make([]int32, 0)

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
func (i Integer32) Empty() bool {
	return len(i) == 0
}

// Contains determines whether the Integer32Slice contains the given value.
func (i Integer32) Contains(value int32) bool {
	for _, val := range i {
		if val == value {
			return true
		}
	}

	return false
}
