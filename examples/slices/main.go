package main

import (
	"fmt"

	"github.com/teamjorge/core-go/v1/slices"
)

func main() {
	fmt.Println("Slices Example")
	fmt.Println("")

	// Create a new StringSlice
	stringSlice := slices.String([]string{"this", "is", "a", "string", "slice"})

	// Print each element in the StringSlice
	fmt.Println("stringSlice.ForEach():")
	stringSlice.ForEach(func(i int, val string) {
		fmt.Print(val, " ")
	})
	fmt.Println()
	fmt.Println()

	// Change the values in the StringSlice.
	// We'll change each entry to be "[index]. [value]"
	fmt.Println("stringSlice.Map():")
	changedStringSlice := stringSlice.Map(func(i int, val string) string {
		return fmt.Sprintf("%d. %s", i, val)
	})

	fmt.Println(changedStringSlice)
	fmt.Println()

	// Filter the values in a slice
	fmt.Println("stringSlice.Filter():")
	filteredStringSlice := stringSlice.Filter(func(i int, val string) bool {
		return len(val) > 3
	})

	fmt.Println(filteredStringSlice)
	fmt.Println()
}
