# Goptional Package

This package provides a generic `Optional[T]` type for handling nullable or absent values in a type-safe manner. It is designed for general-purpose use and works with any type `T`.

## Features

- **Zero Dependencies**: No external dependencies, making it lightweight and easy to integrate.
- **Type Safety**: Generic implementation for any type `T`.
- **JSON Compatibility**: Implements `json.Marshaler` and `json.Unmarshaler` for seamless integration with JSON serialization and deserialization.
- **Pointer Avoidance**: Uses a struct-based approach to handle optionality without pointers.
- **Zero Value Distinction**: Clearly distinguishes between zero values (like `0`, `""`, `false`) and absent values.
- **General Purpose**: Can be used in any context where you need to represent optional values.

## When to Use Optional vs Pointers

Use `Optional[T]` instead of `*T` when:

- You want to avoid nil pointer dereference issues
- You need to distinguish between zero values and absent values
- You want explicit type safety for optional fields

## Usage

### Basic Usage

```go
import "github.com/uforg/ufoconnect/core/internal/util/optional"

// Create optional values
optionalValue := optional.Optional[string]{
  Value:  "hello",
  Present: true,
}

// Create optional values using helpers
presentValue := optional.Some("hello")
absentValue := optional.None[string]()

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
  Name  optional.Optional[string] `json:"name"`
  Age   optional.Optional[int]    `json:"age"`
  Email optional.Optional[string] `json:"email"`
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
age := optional.Some(25)
height := optional.None[float64]()

// Booleans
isActive := optional.Some(true)
isVerified := optional.None[bool]()

// Slices
tags := optional.Some([]string{"go", "optional"})
emptyTags := optional.None[[]string]()

// Custom types
type Address struct {
    Street string
    City   string
}

homeAddress := optional.Some(Address{Street: "123 Main St", City: "Anytown"})
workAddress := optional.None[Address]()
```

### Method Reference

#### Constructors

- `Some[T](value T) Optional[T]` - Creates an optional with a present value
- `None[T]() Optional[T]` - Creates an optional with no value (absent)

#### Value Access

- `Or(defaultValue T) T` - Returns the value if present, otherwise returns the default

#### Direct Field Access

- `.Present bool` - Boolean flag indicating if value is present
- `.Value T` - The actual value (zero value of T when absent)
