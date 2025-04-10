# json-render

A Go library for automatically generating TypeScript type definitions from Go structs

## Features

- Automatically generate TypeScript type definitions from Go structs
- Generate as interface or type
- Generate properties as readonly
- Choose naming convention (camelCase/snake_case)
- Support for `//go:generate` directive

## Installation

```bash
go get github.com/kaitobq/json-render
```

## Usage

### Using as a CLI tool

```bash
go run github.com/kaitobq/json-render/cmd/typegen -input ./path/to/go/files -output ./path/to/output -naming snake_case
```

Options:
- `-input`: Input directory (contains Go files)
- `-output`: Output directory
- `-interface`: Generate as interface (default: false)
- `-readonly`: Generate as readonly properties (default: false)
- `-naming`: Naming convention (camelCase/snake_case) (default: camelCase)

### Using with go:generate

Add the following comment to your Go file:

```go
//go:generate go run github.com/kaitobq/json-render/cmd/typegen -input . -output ./generated
```

## Example

```go
// request.go
package api

type CreateUserRequest struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}
```

Generated TypeScript:

```typescript
export type CreateUserRequest = {
    first_name: string;
    last_name: string;
    email: string;
    password: string;
}
```

## License

MIT