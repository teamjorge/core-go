package slices

import (
	"errors"
	"fmt"
)

// {{ .SliceName }} wraps a normal {{ .SliceType }} slice to provide
// additional helper methods
type {{ .SliceName }} []{{ .SliceType }}

// ForEach iterates each item in the given {{ .SliceName }}Slice and executes
// the given modifier function with it'{{ .SliceModifier }} index and value
func ({{ .SliceModifier }} {{ .SliceName }}) ForEach(modifier func(index int, val {{ .SliceType }})) {
	for index, value := range {{ .SliceModifier }} {
		modifier(index, value)
	}
}

// Map iterates each item in the given {{ .SliceName }}Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the {{ .SliceName }}.
// This method does modify the {{ .SliceName }} in place.
func ({{ .SliceModifier }} {{ .SliceName }}) Map(modifier func(index int, val {{ .SliceType }}) {{ .SliceType }}) {{ .SliceName }} {
	res := make([]{{ .SliceType }}, 0)
	for index, value := range {{ .SliceModifier }} {
		res = append(res, modifier(index, value))
	}
	return res
}

// Filter iterates each item in the given {{ .SliceName }}Slice.
//
// On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
//
// This method does not modify the {{ .SliceName }} in place and will return
// the modified version
func ({{ .SliceModifier }} {{ .SliceName }}) Filter(modifier func(index int, val {{ .SliceType }}) bool) {{ .SliceName }} {
	if len({{ .SliceModifier }}) == 0 {
		return {{ .SliceModifier }}
	}

	res := make([]{{ .SliceType }}, 0)
	for index, value := range {{ .SliceModifier }} {
		if modifier(index, value) {
			res = append(res, value)
		}
	}

	return res
}

// Pop removes an item from the given {{ .SliceName }}Slice at given index.
//
// The removed element is return by Pop. An error will be returned
// if the given index is out of bounds for the given {{ .SliceName }}
func ({{ .SliceModifier }} {{ .SliceName }}) Pop(index int) ({{ .SliceType }}, {{ .SliceName }}, error) {
	var item {{ .SliceType }}
	res := make([]{{ .SliceType }}, 0)

	if len({{ .SliceModifier }}) == 0 {
		err := "Pop on empty slice failed"
		return item, res, errors.New(err)
	}

	if index < 0 || index > (len({{ .SliceModifier }})-1) {
		err := fmt.Sprintf("Pop on index %d not available on slice of length %d", index, len({{ .SliceModifier }}))
		return item, res, errors.New(err)
	}

	item = {{ .SliceModifier }}[index]
	res = append({{ .SliceModifier }}[:index], {{ .SliceModifier }}[index+1:]...)

	return item, res, nil
}

// Empty determines whether the slice is empty
func ({{ .SliceModifier }} {{ .SliceName }}) Empty() bool {
	return len({{ .SliceModifier }}) == 0
}
