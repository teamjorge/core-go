package main

import (
	"fmt"
	"reflect"

	"github.com/teamjorge/core-go/v1/sets"
)

func main() {
	fmt.Println("Sets Example")
	fmt.Println()

	// Create a new String Set from string slices
	fmt.Println("NewString()")
	sliceOne := []string{"val1", "val2", "val3"}
	sliceTwo := []string{"val2", "val1", "val5"}
	mySet := sets.NewString(sliceOne, sliceTwo)
	fmt.Println(mySet.ToSlice())
	fmt.Println()

	// Add more strings to the Set
	fmt.Println("Add()")
	newValues := []string{"val5", "val6", "val7", "val8"}
	mySet.Add(newValues...)
	fmt.Println(mySet.ToSlice())
	fmt.Println()

	// Remove some values from the Set
	fmt.Println("Remove()")
	mySet.Remove("val5", "val7", "val8")
	fmt.Println(mySet.ToSlice())
	fmt.Println()

	// Output the values in the Set to a String Slice
	fmt.Println("ToSlice()")
	mySetAsASlice := mySet.ToSlice()
	fmt.Println("mySetAsASlice is type: ", reflect.TypeOf(mySetAsASlice))
	fmt.Println()

	// Generic

	// Create a new empty UserSet
	// See how to setup your struct below
	user := make(UserSet, 0)

	// Add some users to the set
	// We omit the return value since there will be no errors.
	fmt.Println("sets.Add()")
	_ = sets.Add(
		user,
		User{Name: "billy", Age: 9},
		User{Name: "billy", Age: 9},
		User{Name: "bobby", Age: 12},
		User{Name: "tommy", Age: 15},
		User{Name: "jane", Age: 17},
	)
	fmt.Printf("%+v\n", user)
	fmt.Println()

	// Delete some of the users in the set
	// We omit the return value since there will be no errors.
	fmt.Println("sets.Delete()")
	_ = sets.Delete(
		user,
		User{Name: "billy", Age: 9},
	)
	fmt.Printf("%+v\n", user)
	fmt.Println()

	// Output the values in the set back to a slice
	fmt.Println("sets.ToSlice()")
	newUserList, _ := sets.ToSlice(user)
	fmt.Printf("%+v\n", newUserList.([]User)) // Cast it back to the User type
	fmt.Println()
}

// User type definition
type User struct {
	Name string
	Age  int
}

// Initialize a new type for the User Set
type UserSet map[User]struct{}

// Declare the Add function on our UserSet Type
func (i UserSet) Add(in ...interface{}) {
	for _, x := range in {
		i[x.(User)] = struct{}{} // Cast the value of the item to your type
	}
}

// Declare the Delete function on our UserSet Type
func (i UserSet) Delete(elems ...interface{}) {
	for _, x := range elems {
		delete(i, x.(User)) // Cast the value of the element to your type
	}
}

// Declare the ToSlice function on our UserSet Type
func (i UserSet) ToSlice() interface{} {
	res := make([]User, 0) // Create a slice of your type for the results
	for key := range i {
		res = append(res, key)
	}
	return res
}
