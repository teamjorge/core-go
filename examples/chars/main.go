package main

import (
	"fmt"
	"reflect"

	"github.com/teamjorge/core-go/v1/chars"
)

func main() {
	fmt.Println("Chars Example")

	fmt.Println("Char.Reverse()")

	myString := chars.Char("uoy evol i")

	reversedString := myString.Reverse()

	fmt.Println("reversedString value: ", reversedString)

	fmt.Println("reversedString is type: ", reflect.TypeOf(reversedString))
}
