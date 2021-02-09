package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	v1 "github.com/teamjorge/core-go/v1"
)

func main() {
	t := templateArgs{}

	t.pkg = flag.String("pkg", "",
		"Package type you want to generate for. For example: -pkg sets or -pkg slices")
	t.name = flag.String("name", "",
		"Name of the item. For example: -name Integer when used with -pkg Slice will generate and IntegerSlice and a file called v1/slices/integer.go")
	t.tp = flag.String("type", "",
		"Golang type for the item. If generating a new Slice, all references will use the given type. This field is CASE-SENSITIVE")
	t.modifier = flag.String("modifier", "",
		"Modifier to use when referencing your type. For example: If -modifier s, all method receivers and other references will be named s")

	t.testData = flag.String("test-data", "",
		"Comma separated list of test data to use. Minimum of 5 items. For example: -testdata '1,231,51,23,54,111'")
	t.nilValue = flag.String("nil-value", "",
		"Nil value for the given type")
	t.noTests = flag.Bool("no-tests", false,
		"Generates without Unit Tests")
	t.randomValue = flag.String("random-value", "",
		"Random value for tests outside of test data")

	printVersion := flag.Bool("v", false,
		"Print the current version")

	flag.Parse()

	if err := t.validate(); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("-> core-go template builder", v1.Version)

	if *printVersion {
		return
	}

	var err error
	switch *t.pkg {
	case "sets":
		err = genSet(&t)
		break
	case "slices":
		err = genSlice(&t)
		break
	default:
		flag.Usage()
		return
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
