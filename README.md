# core-go

[![Documentation](https://godoc.org/github.com/teamjorge/core-go/v1?status.svg)](https://godoc.org/github.com/teamjorge/core-go/v1) [![codecov](https://codecov.io/gh/teamjorge/core-go/branch/main/graph/badge.svg?token=08QVKSEPXT)](https://codecov.io/gh/teamjorge/core-go) [![Go Report Card](https://goreportcard.com/badge/github.com/teamjorge/core-go/v1)](https://goreportcard.com/report/github.com/teamjorge/core-go/v1)

Golang library for general utility functions that are common in other languages.

## Install

```bash
go get -u github.com/teamjorge/core-go/v1
```

## Overview

This library aims to provide general functions that can be found in many other languages. No reflection is used and all functions are statically typed. 

## Sets

### Usage

Sets currently support the following types:

* Strings

Example usage:

```go
package main

import (
    "fmt"
    "github.com/teamjorge/core-go/v1"
)

myStringSlice := []string{"val1", "val2", "val3"}
mySecondSlice := []string{"val2", "val3", "val4"}
s := sets.NewString(myStringSlice, mySecondSlice)

fmt.Print(s.ToSlice())
> ["val1", "val2", "val3", "val4"]
```

Each type of set implements the following methods:

|method|usage|
|------|-----|
|`Add`|Adds a new element to the set|
|`Remove`|Removes an element from the set|
|`Empty`|Determines if the given slice is empty|

Additionally, for simple distinct operations, there are functions such as `ToStringSlice`:

```go
myStringSlice := []string{"val1", "val2", "val3"}
mySecondSlice := []string{"val2", "val3", "val4"}

fmt.Print(sets.ToStringSlice(myStringSlice, mySecondSlice))
> ["val1", "val2", "val3", "val4"]
```

### Interfaces

The `sets` package does provide a generic interface called `Set` which implements the following methods:

* `Empty() bool`

Using the `Set` interface gives you the ability to use generic functions, such as `sets.IsEmpty(mySet)` to determine if the given set is empty.

## Versioning

The package is versioned to `v1` for now due to:

* Generics for golang being on the horizon
* In case there are breaking changes introduced (extremely low risk)
