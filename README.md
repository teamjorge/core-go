# core-go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/teamjorge/core-go)](https://pkg.go.dev/github.com/teamjorge/core-go) [![Documentation](https://godoc.org/github.com/teamjorge/core-go/v1?status.svg)](https://godoc.org/github.com/teamjorge/core-go/v1) [![codecov](https://codecov.io/gh/teamjorge/core-go/branch/main/graph/badge.svg?token=08QVKSEPXT)](https://codecov.io/gh/teamjorge/core-go) [![Go Report Card](https://goreportcard.com/badge/github.com/teamjorge/core-go/v1)](https://goreportcard.com/report/github.com/teamjorge/core-go/v1)

Golang library for general utility functions that are common in other languages.

## Install

```bash
go get -u github.com/teamjorge/core-go/v1
```

## Overview

This library aims to provide general functions that can be found in many other languages. Majority of this package is statically typed, however, some generic implementations are provided for user defined structures.

This package does not use `reflect`.

## Features

* [Sets](./docs/Sets.md)

* [Slices](./docs/Slices.md)

* [Chars](./docs/Chars.md)

## Unit Testing

The unit tests can be run from the root of this repository using:

```bash
go test -cover ./v1/...
```

## Template

The `template` binary provides a way to quickly generate source code for some of the concepts in the `v1` package. This can either be used in your application/package or to contribute to `core-go`.

[Template Documentation](./template/README.md)
