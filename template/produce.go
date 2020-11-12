package main

import (
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

func produceTemplate(config interface{}, templatePath, templateName, outpath string) error {
	t := template.New(templateName).Funcs(funcMap)

	templateInput, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	if _, err := t.Parse(string(templateInput)); err != nil {
		return err
	}

	o, err := os.Create(outpath)
	if err != nil {
		return err
	}

	defer o.Close()

	err = t.Execute(o, config)
	if err != nil {
		return err
	}

	err = formatOutput(outpath)

	return err
}

func formatOutput(path string) error {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	out, err := format.Source(source)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, out, os.ModePerm)
}
