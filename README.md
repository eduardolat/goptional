<p align="center">
  <h1 align="center">Goptional</h1>
  <p align="center">
    <img align="center" width="70" src="./goptional.png"/>
  </p>
  <p align="center">
    Provides a generic <code>Optional[T]</code> type for handling nullable or absent values in a type-safe manner. It is designed for general-purpose use and works with any type <code>T</code>
  </p>
</p>

<p align="center">
  <a href="https://github.com/eduardolat/goptional/actions/workflows/ci.yaml?query=branch%3Amain">
    <img src="https://github.com/eduardolat/goptional/actions/workflows/ci.yaml/badge.svg" alt="CI Status"/>
  </a>
  <a href="https://goreportcard.com/report/eduardolat/goptional">
    <img src="https://goreportcard.com/badge/eduardolat/goptional" alt="Go Report Card"/>
  </a>
  <a href="https://github.com/eduardolat/goptional/releases/latest">
    <img src="https://img.shields.io/github/release/eduardolat/goptional.svg" alt="Release Version"/>
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/eduardolat/goptional.svg" alt="License"/>
  </a>
  <a href="https://github.com/eduardolat/goptional">
    <img src="https://img.shields.io/github/stars/eduardolat/goptional?style=flat&label=github+stars"/>
  </a>
</p>

## Features

- **Zero Dependencies**: No external dependencies, making it lightweight and easy to integrate.
- **Type Safety**: Generic implementation for any type `T`.
- **JSON Compatibility**: Implements `json.Marshaler` and `json.Unmarshaler` for integration with JSON serialization and deserialization.
- **Pointer Avoidance**: Uses a struct-based approach to handle optionality without pointers.
- **Zero Value Distinction**: Clearly distinguishes between zero values (like `0`, `""`, `false`) and absent values.
- **General Purpose**: Can be used in any context where you need to represent optional values (it does not need to be JSON-specific).

## When to Use Optional vs Pointers

Use `Optional[T]` instead of `*T` when:

- You want to avoid nil pointer dereference issues
- You need to distinguish between zero values and absent values
- You want explicit type safety for optional fields

## Usage

### Installation

```bash
go get github.com/eduardolat/goptional
```

### Basic Usage

```go
import "github.com/eduardolat/goptional"

// Create optional values (the direct way)
optionalValue := goptional.Optional[string]{
  Value:  "hello",
  Present: true,
}

// Create optional values using helpers
presentValue := goptional.Some("hello")
absentValue := goptional.None[string]()

// Check presence
if presentValue.Present {
  fmt.Println("Value:", presentValue.Value) // Direct access to Value field
}

// Get value or default
defaultValue := absentValue.Or("default")
fmt.Println(defaultValue) // "default"
```

### JSON Serialization

Any absent value is always serialized as `null` in JSON. This makes it predictable and consistent for API responses.

```go
type User struct {
  Name  goptional.Optional[string] `json:"name"`
  Age   goptional.Optional[int]    `json:"age"`
  Email goptional.Optional[string] `json:"email"`
}

// Unmarshaling JSON with null values
data := `{"name": "John", "age": null, "email": "john@example.com"}`
var user User
json.Unmarshal([]byte(data), &user)

fmt.Println(user.Name.Present)  // true
fmt.Println(user.Name.Value)    // "John"
fmt.Println(user.Age.Present)   // false
fmt.Println(user.Email.Present) // true
fmt.Println(user.Email.Value)   // "john@example.com"

// Marshaling back to JSON - absent values become null
jsonData, _ := json.Marshal(user)
// {"name":"John","age":null,"email":"john@example.com"}
```

### Working with Different Types

```go
// Numbers
age := goptional.Some(25)
height := goptional.None[float64]()

// Booleans
isActive := goptional.Some(true)
isVerified := goptional.None[bool]()

// Slices
tags := goptional.Some([]string{"go", "optional"})
emptyTags := goptional.None[[]string]()

// Custom types
type Address struct {
    Street string
    City   string
}

homeAddress := goptional.Some(Address{Street: "123 Main St", City: "Anytown"})
workAddress := goptional.None[Address]()
```

### Method Reference

#### Constructors

- `Some[T](value T) Optional[T]` - Creates an optional with a present value
- `None[T]() Optional[T]` - Creates an optional with no value (absent)

#### Value Access

- `.Or(defaultValue T) T` - Returns the value if present, otherwise returns the default

#### Direct Field Access

- `.Present bool` - Boolean flag indicating if value is present
- `.Value T` - The actual value (zero value of T when Present is false)
