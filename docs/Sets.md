# Sets

- [Sets](#sets)
  - [Types](#types)
  - [Usage](#usage)
  - [Interfaces](#interfaces)
    - [Generic](#generic)
      - [Functions](#functions)
    - [Set](#set)

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

### Generic

The `sets` package allows for `Generic` sets to be created using the `Generic` interface:

```go
type Generic interface {
    Add(...interface{})
    Delete(...interface{})
    ToSlice() interface{}
}
```

This interface provides set functionality for types that are not explicitly defined by the `sets` package.

For example, if we had a defined struct called `User`:

```go
type User struct {
    Name string
    Age  int
}
```

If we wanted distinct users to form part of a group based on their names, we could use a `set` for that. 

First, we'll start by created the underlying type:

```go
type UserSet map[User]struct{}
```

The pattern will always be `map[T]struct{}`.

Next, we'll implement the three methods required by the `Generic` interface:

```go
func (i UserSet) Add(in ...interface{}) {
    for _, x := range in {
        i[x.(User)] = struct{}{} // Cast the value of the item to your type
    }
}

func (i UserSet) Delete(elems ...interface{}) {
    for _, x := range elems {
        delete(i, x.(User)) // Cast the value of the element to your type
    }
}

func (i UserSet) ToSlice() interface{} {
    res := make([]User, 0)  // Create a slice of your type for the results
    for key := range i {
        res = append(res, key)
    }
    return res
}
```

Let's initialize our new set:

```go
g := make(UserSet, 0)
```

#### Functions

We are now able to use the `Generic` functions implemented by the package:

`Add`:

Adds new element(s) to the set.

```go
_ = Add(
    g,
    User{Name: "billy", Age: 9},
    User{Name: "billy", Age: 9},
    User{Name: "bobby", Age: 12},
    User{Name: "bobby", Age: 12},
    User{Name: "tommy", Age: 15},
    User{Name: "tommy", Age: 15},
    User{Name: "jane", Age: 17},
    User{Name: "jane", Age: 17},
)
```

`g` will no contain 4 items:

- billy 9
- bobby 12
- tommy 15
- jane 17

The duplicate entries will be disregarded. If we were to change the age of a single entry, another entry will be added to the set.

If a value of a different type is passed to the `Add` function, the changes will be rolled back and an error returned.

`Delete`:

Deletes element(s) from the set.

Let's say we want to now delete some of the User(s) we added using the `Add` function:

```go
_ = Delete(
    g,
    User{Name: "billy", Age: 9},
    User{Name: "jane", Age: 17},
)
```

Our User set will now only contain:

- bobby 12
- tommy 15

If a value of a different type is passed to the `Delete` function, the changes will be rolled back and an error returned.

`ToSlice`:

Converts the set to a slice

After getting our set exactly how we want it, we probably want to retrieve those values. By using the `ToSlice` function, we can create a new slice of the set values:

```go
uniqueUsers := ToSlice(g)
```

Since the `ToSlice` function always returns `interface{}`, we will need to cast it to the appropriate type:

```go
userSlice := uniqueUsers.([]User)
```

The type casting is an unfortunate part, but it is necessary to avoid the use of the `reflect` package.

`AddUnsafe` and `DeleteUnsafe`:

The `AddUnsafe` and `DeleteUnsafe` methods exist to provide an alternative to rollbacks being performed when trying to add an unsuitable type to a `Generic`.

For `AddUnsafe`, any elements added before an error occurs will remain in the `Generic` and the error will be returned.

For `DeleteUnsafe`, any elements deleted before an error occurs will remain deleted from the `Generic` and the error will be returned

### Set

The `sets` package does provide another generic interface called `Set` which implements the following methods:

- `Empty() bool`

Using the `Set` interface gives you the ability to use generic functions, such as `sets.IsEmpty(mySet)` to determine if the given set is empty.

The type safe `sets` implemented by `core-go` will have this interface implemented by default. Any `Generic` sets will need to manually implement this interface.
