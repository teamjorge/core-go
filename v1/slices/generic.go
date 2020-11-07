package slices

// Slice is a generic interface for slices of all types
type Slice interface {
	Empty() bool
}

// IsEmpty determines if the given slice is empty or not
func IsEmpty(s Slice) bool {
	return s.Empty()
}
