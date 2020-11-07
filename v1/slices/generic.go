package slices

// Slice is a generic interface for slices of all types
type Slice interface {
	Empty() bool
}

// IsEmpty determines if the given slice is empty or not
func IsEmpty(s Slice) bool {
	return s.Empty()
}

// Generic provides an interface for types that are not supported by this library that wish
// to use the slice related functions in this package
//
// Common uses would be to implement the Generic methods on a slice of a defined struct
// to provide methods defined in this package
//
// Given the following struct:
// 	type Person struct {
// 		Name string
// 	}
//
// And type definition for it's slice:
//
// 	type Persons []Person
//
// You would define the following methods:
//
//	func (m Persons) Unpack() []interface{} {
//		res := make([]interface{}, 0)
//		for _, i := range m {
//		  res = append(res, i)
//		}
//		return res
//	}
//
//	func (m Persons) Pack(replace []interface{}) Generic {
//		res := make([]Person, 0)
//		for _, value := range replace {
//			res = append(res, value.(Person))
//		}
//		return Persons(res)
//	}
// Unpack converts the Generic instance into a interface slice.
// This method is called during
// generic slice functions to ensure the correct type is passed to modifier
// functions
//
// Pack repackages a interface slice into a Generic instance.
// This method is called during generic slice functions to ensure returning of the correct type
type Generic interface {
	Unpack() []interface{}
	Pack([]interface{}) Generic
}

// ForEach iterates each item in the given Generic and executes
// the given modifier function with it's index and value
func ForEach(g Generic, modifier func(index int, val interface{})) {
	items := g.Unpack()
	for index, val := range items {
		modifier(index, val)
	}
}

// Map iterates each item in the given Generic. On each iteration,
// the current index and value will be passed to the modifier function.
// The value returned will overwrite the current value at the specific
// index in the Generic.
// This method does modify the Generic in place.
func Map(g Generic, modifier func(index int, val interface{}) interface{}) Generic {
	items := g.Unpack()
	res := make([]interface{}, 0)
	for index, val := range items {
		res = append(res, modifier(index, val))
	}
	return g.Pack(res)
}

// Filter iterates each item in the given Generic. On each iteration,
// the current index and value will be passed to the modifier function.
// Only iterations that returned true when passed to the modifier thing
// will be returned.
// This method does not modify the Generic in place and will return
// the modified version
func Filter(g Generic, modifier func(index int, val interface{}) bool) Generic {
	items := g.Unpack()
	res := make([]interface{}, 0)
	for index, val := range items {
		if modifier(index, val) {
			res = append(res, val)
		}
	}
	return g.Pack(res)
}
