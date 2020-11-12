package main

import (
	"fmt"
	"strings"
)

const (
	setsTestTemplatePath    string = "./templates/set_test.go.tpl"
	setsTemplatePath        string = "./templates/set.go.tpl"
	setsFilePath            string = "../v1/sets"
	setsTemplateBuilder     string = "Set Builder"
	setsTestTemplateBuilder string = "Set Test Builder"
)

// Set defines the template for generating a new Set
type Set struct {
	SetName     string
	SetType     string
	SetModifier string
}

// SetTest defines the template for generating a new Set test file
type SetTest struct {
	Set
	TestItems           string
	TestItemsSplit      []string
	TestItemsMap        string
	TestItemsMapWithNil string
	NilValue            string
}

func getSetTestValueMap(s SetTest) string {
	items := make([]string, 0)

	for _, item := range s.TestItemsSplit {
		items = append(items, fmt.Sprintf("%s: false", item))
	}

	return strings.Join(items, ", ")
}

func genSet(name, tp, mod, testData, nilValue string, noTests bool) error {
	SetConfig, err := genSetTemplate(name, tp, mod)
	if err != nil {
		return err
	}

	if noTests {
		return nil
	}

	err = genSetTestTemplate(SetConfig, testData, nilValue)

	return err
}

func genSetTemplate(name, tp, mod string) (Set, error) {
	SetConfig := Set{
		SetName:     strings.Title(name),
		SetType:     strings.ToLower(tp),
		SetModifier: strings.ToLower(mod),
	}

	outpath := fmt.Sprintf("%s/%s.go", setsFilePath, strings.ToLower(SetConfig.SetName))
	err := produceTemplate(SetConfig, setsTemplatePath, setsTemplateBuilder, outpath)

	return SetConfig, err
}

func genSetTestTemplate(SetConfig Set, testData, nilValue string) error {
	SetTestConfig := SetTest{
		Set:            SetConfig,
		TestItems:      testData,
		TestItemsSplit: strings.Split(testData, ","),
		NilValue:       nilValue,
	}

	SetTestConfig.TestItemsMap = getSetTestValueMap(SetTestConfig)
	SetTestConfig.TestItemsMapWithNil = fmt.Sprintf(
		"%s, %s: false", SetTestConfig.TestItemsMap, SetTestConfig.NilValue,
	)

	outpath := fmt.Sprintf("%s/%s_test.go", setsFilePath, strings.ToLower(SetTestConfig.SetName))
	err := produceTemplate(SetTestConfig, setsTestTemplatePath, setsTemplateBuilder, outpath)

	return err
}
