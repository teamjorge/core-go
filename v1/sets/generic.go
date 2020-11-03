package sets

// Set is a generic interface for sets of all types
type Set interface {
	Empty() bool
}

// IsEmpty determines if the given set is empty or not
func IsEmpty(s Set) bool {
	return s.Empty()
}
