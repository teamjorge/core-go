package sets

/*
String generates a list of distinct strings from the given strings.

This is comparable to
the Set type in other languages, such as Python, Java
etc


Example Usage:

 someStrings := []string{"this", "is", "unique", "this", "is", "distinct"}
 result := String(someStrings)
 result
 > []string{"this", "is", "unique", "distinct"}
*/
func String(in ...[]string) []string {
	set := make(map[string]bool, 0)
	for _, arr := range in {
		for _, s := range arr {
			set[s] = false
		}
	}
	res := make([]string, 0)
	for k := range set {
		res = append(res, k)
	}
	return res
}
