# Sets

- [Sets](#sets)
  - [Types](#types)
  - [Usage](#usage)
  - [Interfaces](#interfaces)

## Types

Sets currently support the following types:

- Strings

## Usage

The following example creates a new set from two string slices. The resulting set is then output to a string slice.

```go
package main

import (
    "fmt"

    "github.com/teamjorge/core-go/v1/sets"
)

func main() {
    myStringSlice := []string{"val1", "val2", "val3"}
    mySecondSlice := []string{"val2", "val3", "val4"}
    s := sets.NewString(myStringSlice, mySecondSlice)

    fmt.Print(s.ToSlice())
}

>   ["val1", "val2", "val3", "val4"]
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

## Interfaces

The `sets` package does provide a generic interface called `Set` which implements the following methods:

- `Empty() bool`

Using the `Set` interface gives you the ability to use generic functions, such as `sets.IsEmpty(mySet)` to determine if the given set is empty.
