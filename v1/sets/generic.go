package sets

import (
	"errors"
	"fmt"
)

// Set is a generic interface for sets of all types.
type Set interface {
	Empty() bool
}

// IsEmpty determines if the given set is empty or not.
func IsEmpty(s Set) bool {
	return s.Empty()
}

// Generic provides an interface for wrapping map[T]bool
// structures to provide set functionality.
type Generic interface {
	Add(...interface{})
	Delete(...interface{})
	ToSlice() interface{}
}

// resolveError attempts to resolve the given recover value
// and return a standardized error.
func resolveError(r interface{}) error {
	var err error
	switch x := r.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		err = fmt.Errorf("%v", x)
	}
	return err
}

// Add adds items to the given Generic Set.
//
// Add will attempt to add the items to the given set and return an
// error if there is type mismatch.
func Add(g Generic, items ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = resolveError(r)
			// Remove items that were added.
			// rollbackAdd function should panic and recover on the
			// same item that caused this panic.
			rollbackAdd(g, items...)
		}
	}()
	g.Add(items...)
	return nil
}

// rollbackAdd removes any newly added items to the set by running
// the Delete method with the same set of items.
func rollbackAdd(g Generic, items ...interface{}) {
	defer func() {
		recover()
	}()
	g.Delete(items...)
}

// AddUnsafe removes items from the given Generic Set.
//
// As opposed to Add, AddUnsafe will not rollback changes
// if an error has occurred. Thus, any items added before
// that are added before the error occurred will be persisted
// in the given Generic.
//
// Any errors that occurred will still be returned by this function
// to indicate that a items were only partially added.
func AddUnsafe(g Generic, items ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = resolveError(r)
		}
	}()
	g.Add(items...)
	return nil
}

// Delete removes items from the given Generic Set.
//
// Delete will attempt to remove the items from the given set and return an
// error if there is type mismatch.
func Delete(g Generic, items ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = resolveError(r)
			// Add items that were removed.
			// rollbackDelete function should panic and recover on the
			// same item that caused this panic.
			rollbackDelete(g, items...)
		}
	}()
	g.Delete(items...)
	return nil
}

// rollbackDelete adds any newly removed items to the set by running
// the Add method with the same set of items.
func rollbackDelete(g Generic, items ...interface{}) {
	defer func() {
		recover()
	}()
	g.Add(items...)
}

// DeleteUnsafe removes items from the given Generic Set.
//
// As opposed to Delete, DeleteUnsafe will not rollback changes
// if an error has occurred. Thus, any items that are deleted
// before the error occurred will remain deleted.
//
// Any errors that occurred will still be returned by this function
// to indicate that a items were only partially deleted.
func DeleteUnsafe(g Generic, items ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = resolveError(r)
		}
	}()
	g.Delete(items...)
	return nil
}

// ToSlice returns a slice from the items of the given Generic set.
//
// The returned type is interface{} and will need to be type casted to
// the correct slice type. For example, if we have an Generic has an
// underlying type of map[int]bool, we will need to cast the slice with:
//
//  mySlice, _ := ToSlice(myGeneric)
//  mySlice.([]int)
//
// An error will be returned if the given Generic is nil
func ToSlice(g Generic) (interface{}, error) {
	if g == nil {
		return nil, errors.New("Given Generic is nil, can't convert to slice")
	}
	res := g.ToSlice()
	return res, nil
}
