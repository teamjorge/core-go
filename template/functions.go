package main

import (
	"strconv"
	"text/template"
)

func add(a, b int) string {
	return strconv.Itoa(a + b)
}

func subtract(a, b int) string {
	return strconv.Itoa(a - b)
}

var funcMap = template.FuncMap{
	"add":      add,
	"subtract": subtract,
}
