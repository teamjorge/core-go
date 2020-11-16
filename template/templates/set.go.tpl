package sets

/*
To{{ .SetName }}Slice generates a {{ .SetType }} slice of distinct {{ .SetType }}{{ .SetModifier }} from the given {{ .SetType }}{{ .SetModifier }}.

Example Usage:

 some{{ .SetName }}{{ .SetModifier }} := []{{ .SetType }}{"this", "is", "unique", "this", "is", "distinct"}
 result := {{ .SetName }}(some{{ .SetName }}{{ .SetModifier }})
 result
 > []{{ .SetType }}{"this", "is", "unique", "distinct"}
*/
func To{{ .SetName }}Slice(in ...[]{{ .SetType }}) []{{ .SetType }} {
	set := New{{ .SetName }}(in...)
	res := set.ToSlice()

	return res
}

// {{ .SetName }} provides a type for {{ .SetType }} sets
type {{ .SetName }} map[{{ .SetType }}]struct{}

// New{{ .SetName }} creates a new set of {{ .SetType }} type
func New{{ .SetName }}(in ...[]{{ .SetType }}) *{{ .SetName }} {
	set := make({{ .SetName }}, 0)
	for _, arr := range in {
		for _, key := range arr {
			set[key] = struct{}{}
		}
	}
	return &set
}

// Add new element({{ .SetModifier }}) to the set
func ({{ .SetModifier }} {{ .SetName }}) Add(in ...{{ .SetType }}) {
	for _, val := range in {
		{{ .SetModifier }}[val] = struct{}{}
	}
}

// Remove element({{ .SetModifier }}) from the set
func ({{ .SetModifier }} {{ .SetName }}) Remove(element ...{{ .SetType }}) {
	for _, elem := range element {
		delete({{ .SetModifier }}, elem)
	}
}

// ToSlice returns the values in the set as a {{ .SetType }} slice
func ({{ .SetModifier }} {{ .SetName }}) ToSlice() []{{ .SetType }} {
	res := make([]{{ .SetType }}, 0)
	for key := range {{ .SetModifier }} {
		res = append(res, key)
	}
	return res
}

// Empty determines whether the set is empty
func ({{ .SetModifier }} {{ .SetName }}) Empty() bool {
	return len({{ .SetModifier }}) == 0
}
