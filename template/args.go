package main

import "errors"

type templateArgs struct {
	// Base
	pkg      *string
	name     *string
	tp       *string
	modifier *string

	// Testing
	testData    *string
	nilValue    *string
	noTests     *bool
	randomValue *string
}

func (t templateArgs) validate() error {
	if *t.pkg == "" || *t.name == "" || *t.tp == "" || *t.modifier == "" {
		return errors.New("")
	}

	if !*t.noTests && (*t.testData == "" || *t.nilValue == "") {
		return errors.New("-test-data and -nil-value is required when tests are being generated")
	}
	return nil
}
