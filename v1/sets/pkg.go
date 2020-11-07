// Package sets provides helper methods for Sets. A set is a unique collection
// of values for a given type.
//
// Example:
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/teamjorge/core-go/v1/sets"
//	)
//
//	func main() {
//		myStringSlice := []string{"val1", "val2", "val3"}
//		mySecondSlice := []string{"val2", "val3", "val4"}
//		s := sets.NewString(myStringSlice, mySecondSlice)
//
//		fmt.Print(s.ToSlice())
//	}
//
//	>   ["val1", "val2", "val3", "val4"]
package sets
