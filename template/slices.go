package main

import (
	"fmt"
	"strings"
)

const (
	sliceTestTemplatePath    string = "./templates/slice_test.go.tpl"
	sliceTemplatePath        string = "./templates/slice.go.tpl"
	sliceFilePath            string = "../v1/slices"
	sliceTemplateBuilder     string = "Slice Builder"
	sliceTestTemplateBuilder string = "Slice Test Builder"
)

// Slice defines the template for generating a new slice
type Slice struct {
	SliceName     string
	SliceType     string
	SliceModifier string
}

// SliceTest defines the template for generating a new slice test file
type SliceTest struct {
	Slice
	TestItems      string
	TestItemsSplit []string
	NilValue       string
}

func genSlice(args *templateArgs) error {
	sliceConfig, err := genSliceTemplate(args)
	if err != nil {
		return err
	}

	if *args.noTests {
		return nil
	}

	err = genSliceTestTemplate(sliceConfig, args)

	return err
}

func genSliceTemplate(args *templateArgs) (Slice, error) {
	sliceConfig := Slice{
		SliceName:     strings.Title(*args.name),
		SliceType:     *args.tp,
		SliceModifier: strings.ToLower(*args.modifier),
	}

	outpath := fmt.Sprintf("%s/%s.go", sliceFilePath, strings.ToLower(sliceConfig.SliceName))
	err := produceTemplate(sliceConfig, sliceTemplatePath, sliceTemplateBuilder, outpath)

	return sliceConfig, err
}

func genSliceTestTemplate(sliceConfig Slice, args *templateArgs) error {
	sliceTestConfig := SliceTest{
		Slice:          sliceConfig,
		TestItems:      *args.testData,
		TestItemsSplit: strings.Split(*args.testData, ","),
		NilValue:       *args.nilValue,
	}

	outpath := fmt.Sprintf("%s/%s_test.go", sliceFilePath, strings.ToLower(sliceTestConfig.SliceName))
	err := produceTemplate(sliceTestConfig, sliceTestTemplatePath, sliceTestTemplateBuilder, outpath)

	return err
}
