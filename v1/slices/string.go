package slices

import (
	"errors"
	"fmt"
)

// StringSlice wraps a normal string slice to provide
// additional helper methods
type StringSlice []string

// ForEach iterates each item in the given StringSlice and executes
// the given modifier function with it's index and value
func (s StringSlice) ForEach(modifier func(index int, val string)) {
	for index, value := range s {
		modifier(index, value)
	}
}

// Map iterates each item in the given StringSlice. On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the StringSlice.
// This method does modify the StringSlice in place.
func (s StringSlice) Map(modifier func(index int, val string) string) StringSlice {
	res := make([]string, 0)
	for index, value := range s {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given StringSlice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the StringSlice in place and will return
// the modified version
func (s StringSlice) Filter(modifier func(index int, val string) bool) StringSlice {
	if len(s) == 0 {
		return s
	}

	res := make([]string, 0)
	for index, value := range s {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given StringSlice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given StringSlice
func (s StringSlice) Pop(index int) (string, StringSlice, error) {
	var item string
	res := make([]string, 0)

	if len(s) == 0 {
		err := "Pop on empty slice failed"
		return item, res, errors.New(err)
	}

	if index < 0 || index > (len(s)-1) {
		err := fmt.Sprintf("Pop on index %d not available on slice of length %d", index, len(s))
		return item, res, errors.New(err)
	}

	item = s[index]
	res = append(s[:index], s[index+1:]...)

	return item, res, nil
}

// Empty determines whether the slice is empty
func (s StringSlice) Empty() bool {
	return len(s) == 0
}
