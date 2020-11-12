# template

Generate `core-go` `v1` source code without hassle.

## Description

`template` provides a way to quickly generate source code for some of the concepts in the `v1` package. The aim of this package is to improve the velocity of adding new types to:

* `sets`
* `slices`

Using this package, it will hopefully remove some of the grunt work required to add type specific implementations of the features. Generating source code are automatically formatted using the built-in `go/format` package.

## Get Started

First step would to ensure that `core-go/template` is your `PWD`.

Next, we need to build the `template` binary:

```bash
go build
```

The template binary has some `flags` that are **required**:

|Flag        |Description                                                                 |
|------------|----------------------------------------------------------------------------|
|`-pkg`| Options are `sets` or `slices`, depending on what you would like to generate.|
|`-name`| Type name. For example: `-name integer` will generate a `IntegerSlice` or `IntegerSet`.|
|`-type`| Actual type name in `go`. For example: `-type int` or `-type float64` or even `-type 'time.Time'`. **This field is case-sensitive!**.|
|`-modifier`| Modifier to use when referencing your type. For example: If `-mod s`, all method receivers and other references will be named `s`|
|`-test-data`| Comma separated list of data to use in tests. For example: `-testdata '1,231,51,23,54,111'`. You should aim to have at least 5 data items.|
|`-nil-value`| The `nil` value for your type. For example, if we are generating for `int`, we will have a `nil` value of `0`.|

**Optional** `flags`:

|Flag        |Description                                                                 |
|------------|----------------------------------------------------------------------------|
|`-no-tests`| Generate source code without generating Unit Tests.|

## Examples

```bash
./template -pkg slices -name integer -type int -modifier i -test-data '1,2,3,4,5' -nil-value 0
```

Will generate a `v1/slices/integer.go` and `v1/slices/integer_test.go`.

```bash
./template -pkg sets -name integer -type int -modifier i -no-tests -test-data '1,2,3,4,5' -nil-value 0
```

Will generate a `v1/sets/integer.go` and `v1/sets/integer_test.go`.
