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

	fmt.Println("stringSlice.Contains():")
	exists := stringSlice.Contains("this")
	doesNotExist := stringSlice.Contains("randomvalue")

	fmt.Println(exists, doesNotExist)
	fmt.Println()

	// Define our Persons
	// See how to setup your struct below
	friends := Persons{
		Person{Name: "billy"},
		Person{Name: "tommy"},
		Person{Name: "pat"},
	}

	// Iterate each item in the friends slice and print the Person name
	fmt.Println("ForEach():")
	slices.ForEach(friends, func(index int, value interface{}) {
		item := value.(Person)
		fmt.Printf("%s ", item.Name)
	})

	fmt.Println()
	fmt.Println()

	// Iterate the friends slice and add a surname to each element.
	fmt.Println("Map():")
	family := slices.Map(friends, func(index int, value interface{}) interface{} {
		item := value.(Person)
		item.Name = item.Name + " jackson"
		return item
	})

	fmt.Printf("%+v\n", family)
	fmt.Println()

	// Filter the slice for Persons not named "pat"
	fmt.Println("Filter():")
	filteredFriends := slices.Filter(friends, func(index int, val interface{}) bool {
		item := val.(Person)
		return item.Name != "pat"
	})
	fmt.Printf("%+v\n", filteredFriends)
	fmt.Println()

	fmt.Println("Contains():")
	friendDoesExist := slices.Contains(
		friends,
		"billy",
		func(val interface{}) interface{} { return val.(Person).Name },
	)
	friendDoesNotExist := slices.Contains(
		friends,
		"patrick",
		func(val interface{}) interface{} { return val.(Person).Name },
	)

	fmt.Println(friendDoesExist, friendDoesNotExist)
}

// Type definition for our Person
type Person struct {
	Name string
}

// Define a Person slice as our own type
type Persons []Person

// Add the Unpack method to implement the slices.Generic interface
func (m Persons) Unpack() []interface{} {
	res := make([]interface{}, 0)
	for _, i := range m {
		res = append(res, i)
	}
	return res
}

// Add the Unpack method to implement the slices.Generic interface
func (m Persons) Pack(replace []interface{}) slices.Generic {
	res := make([]Person, 0)
	for _, value := range replace {
		res = append(res, value.(Person))
	}
	return Persons(res)
}
