package main

import (
	"fmt"
	"reflect"

	"github.com/teamjorge/core-go/v1/sets"
)

func main() {
	fmt.Println("Sets Example")
	fmt.Println()

	fmt.Println("NewString()")
	sliceOne := []string{"val1", "val2", "val3"}
	sliceTwo := []string{"val2", "val1", "val5"}
	mySet := sets.NewString(sliceOne, sliceTwo)
	fmt.Println(mySet.ToSlice())
	fmt.Println()

	fmt.Println("Add()")
	newValues := []string{"val5", "val6", "val7", "val8"}
	mySet.Add(newValues...)
	fmt.Println(mySet.ToSlice())
	fmt.Println()

	fmt.Println("Remove()")
	mySet.Remove("val5", "val7", "val8")
	fmt.Println(mySet.ToSlice())
	fmt.Println()

	fmt.Println("ToSlice()")
	mySetAsASlice := mySet.ToSlice()
	fmt.Println("mySetAsASlice is type: ", reflect.TypeOf(mySetAsASlice))
}
