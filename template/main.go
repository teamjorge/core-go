package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("-> core-go template builder")
	pkg := flag.String("pkg", "",
		"Package type you want to generate for. For example: -pkg sets or -pkg slices")
	name := flag.String("name", "",
		"Name of the item. For example: -name Integer when used with -pkg Slice will generate and IntegerSlice and a file called v1/slices/integer.go")
	tp := flag.String("type", "",
		"Golang type for the item. If generating a new Slice, all references will use the given type. This field is CASE-SENSITIVE")
	mod := flag.String("modifier", "",
		"Modifier to use when referencing your type. For example: If -mod s, all method receivers and other references will be named s")
	testData := flag.String("test-data", "",
		"Comma separated list of test data to use. Minimum of 5 items. For example: -testdata '1,231,51,23,54,111'")
	nilValue := flag.String("nil-value", "",
		"Nil value for the given type")
	noTest := flag.Bool("no-tests", false,
		"Generates without Unit Tests")

	flag.Parse()

	var err error
	switch *pkg {
	case "":
		fmt.Println("-pkg is required. Available options: -pkg sets | -pkg slices")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	case "sets":
		err = genSet(*name, *tp, *mod, *testData, *nilValue, *noTest)
		break
	case "slices":
		err = genSlice(*name, *tp, *mod, *testData, *nilValue, *noTest)
		break
	default:
		flag.Usage()
		return
	}

	if err != nil {
		log.Panic(err)
	}
}
