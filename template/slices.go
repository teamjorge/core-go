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

func genSlice(name, tp, mod, testData, nilValue string, noTests bool) error {
	sliceConfig, err := genSliceTemplate(name, tp, mod)
	if err != nil {
		return err
	}

	if noTests {
		return nil
	}

	err = genSliceTestTemplate(sliceConfig, testData, nilValue)

	return err
}

func genSliceTemplate(name, tp, mod string) (Slice, error) {
	sliceConfig := Slice{
		SliceName:     strings.Title(name),
		SliceType:     strings.ToLower(tp),
		SliceModifier: strings.ToLower(mod),
	}

	outpath := fmt.Sprintf("%s/%s.go", sliceFilePath, strings.ToLower(sliceConfig.SliceName))
	err := produceTemplate(sliceConfig, sliceTemplatePath, sliceTemplateBuilder, outpath)

	return sliceConfig, err
}

func genSliceTestTemplate(sliceConfig Slice, testData, nilValue string) error {
	sliceTestConfig := SliceTest{
		Slice:          sliceConfig,
		TestItems:      testData,
		TestItemsSplit: strings.Split(testData, ","),
		NilValue:       nilValue,
	}

	outpath := fmt.Sprintf("%s/%s_test.go", sliceFilePath, strings.ToLower(sliceTestConfig.SliceName))
	err := produceTemplate(sliceTestConfig, sliceTestTemplatePath, sliceTestTemplateBuilder, outpath)

	return err
}
