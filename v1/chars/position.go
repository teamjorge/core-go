package chars

// Char wraps the string type to provide additional helper methods.
type Char string

// Reverse the given string. This method does not modify the receiver in place
// and will return reversed string.
func (c Char) Reverse() string {
	rs := []rune(c)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
