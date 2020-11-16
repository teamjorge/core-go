package sets

/*
ToStringSlice generates a string slice of distinct strings from the given strings.

Example Usage:

 someStrings := []string{"this", "is", "unique", "this", "is", "distinct"}
 result := String(someStrings)
 result
 > []string{"this", "is", "unique", "distinct"}
*/
func ToStringSlice(in ...[]string) []string {
	set := NewString(in...)
	res := set.ToSlice()

	return res
}

// String provides a type for string sets
type String map[string]struct{}

// NewString creates a new set of string type
func NewString(in ...[]string) *String {
	set := make(String, 0)
	for _, arr := range in {
		for _, s := range arr {
			set[s] = struct{}{}
		}
	}
	return &set
}

// Add new element(s) to the set
func (s String) Add(in ...string) {
	for _, i := range in {
		s[i] = struct{}{}
	}
}

// Remove element(s) from the set
func (s String) Remove(elem ...string) {
	for _, e := range elem {
		delete(s, e)
	}
}

// ToSlice returns the values in the set as a string slice
func (s String) ToSlice() []string {
	res := make([]string, 0)
	for k := range s {
		res = append(res, k)
	}
	return res
}

// Empty determines whether the set is empty
func (s String) Empty() bool {
	return len(s) == 0
}
