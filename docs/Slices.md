# Slices

- [Slices](#slices)
  - [Types](#types)
  - [Usage](#usage)
    - [ForEach](#foreach)
    - [Map](#map)
    - [Filter](#filter)
    - [Contains](#contains)
  - [Interfaces](#interfaces)
  - [Boolean Specific](#boolean-specific)

## Types

Slices currently support the following types:

- `string`
- `boolean`

Additionally, support is provided for user defined types, such as `structs`, via the `Generic` slice interface.

## Usage

For the following examples, we'll assume these variables exist:

```go
mySlice := []string{"this", "is", "a", "string", "slice"}

stringSlice := slices.String(mySlice)
```

### ForEach

```go
stringSlice.ForEach(func(index int, val string) {
    fmt.Printf("%d. %s\n", index, val)
})
```

Output:

```text
0. this
1. is
2. a
3. string
4. slice
```

### Map

```go
mappedSlice := stringSlice.Map(func(index int, val string) string {
    return fmt.Sprintf("%d. %s", index, val)
})

fmt.Println(mappedSlice)
```

Output:

```text
[0. this 1. is 2. a 3. string 4. slice]
```

### Filter

```go
filteredSlice := stringSlice.Filter(func(index int, val string) bool {
    return len(val) < 5
})

fmt.Println(filteredSlice)
```

Output:

```text
[this is a]
```

### Contains

```go
exists := stringSlice.Contains("this")
doesNotExist := stringSlice.Contains("randomvalue")

fmt.Println(exists, doesNotExist)
```

Output:

```text
true false
```

## Interfaces

```go
type Generic interface {
    Unpack() []interface{}
    Pack([]interface{}) Generic
}
```

The `Generic` interface for provides slice functionality for all types that are not outright supported by this library. An example of this would be a user defined struct, such as:

```go
type Person struct {
    Name string
}
```

Let's say we have a slice of `Person` (`[]Person`) that we would like to iterate and modify or filter. To comply with the `Generic` interface, we will need to define our own type to add methods to:

```go
type Persons []Person
```

Next, we will need to add the two methods to `Persons`:

`Unpack` - Unpacks your data as `interface{}`. This is usually called at the start of the generic slice functions.

```go
func (m Persons) Unpack() []interface{} {
    res := make([]interface{}, 0)
    for _, i := range m {
        res = append(res, i)
    }
    return res
}
```

`Pack` - Repackages your data into the correct type after running the generic slice functions.

```go
func (m Persons) Pack(replace []interface{}) Generic {
    res := make([]Person, 0)
    for _, value := range replace {
        res = append(res, value.(Person))
    }
    return Persons(res)
}
```

It is important to note that you will need to add your specified types to three parts of the `Pack` method:

1. Declaring the new slice - `make([]Person, 0)`
1. Casting value to your type - `value.(Person)`
1. Casting the result to your slice type - `Persons(res)`

First we define our Persons data:

```go
friends := Persons{
    Person{Name: "billy"},
    Person{Name: "tommy"},
    Person{Name: "pat"},
}
```

We can now use `ForEach`, `Map`, and `Filter` functions to iterate our `Persons` slice:

```go
slices.ForEach(friends, func(index int, value interface{}) {
    item := value.(Person)
    fmt.Printf("%d. %s \n", index, item.Name)
})
```

Output:

```text
0. billy
1. tommy
2. pat
```

```go
mappedFriends := slices.Map(friends, func(index int, value interface{}) interface{} {
    item := value.(Person)
    item.Name = item.Name + "_jackson"
    return item
})

fmt.Printf("%+v\n", mappedFriends)
```

Output:

```text
[{Name:billy_jackson} {Name:tommy_jackson} {Name:pat_jackson}]
```

```go
filteredFriends := slices.Filter(friends, func(index int, val interface{}) bool {
    item := val.(Person)
    return item.Name != "pat"
})

fmt.Printf("%+v\n", filteredFriends)
```

Output:

```text
[{Name:billy} {Name:tommy}]
```

We can check if our Slice contains specific names:

```go
doesExist := slices.Contains(
    friends,
    "billy",
    func(val interface{}) interface{} { return val.(Person).Name },
)
doesNotExist := slices.Contains(
    friends,
    "patrick",
    func(val interface{}) interface{} { return val.(Person).Name },
)

fmt.Println(doesExist, doesNotExist)
```

Output:

```text
true false
```

## Boolean Specific

`Boolean` slices support two additional methods:

- `Any` - Returns `true` if at least one of indices in the slice is `true`.
- `All` - Returns `true` if all of the indices in the slice are `true`.
