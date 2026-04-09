# Toolkit

A Go package providing utility functions for common tasks, including cryptographically secure random string generation.

## Features

- **RandomString**: Generate cryptographically secure random strings of specified length
- Extensible design for adding more utility functions

## Installation

```bash
go get github.com/atirek89/toolkit
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/atirek89/toolkit"
)

func main() {
    var tools toolkit.Tools

    // Generate a random string of length 10
    randomStr := tools.RandomString(10)
    fmt.Println("Random String:", randomStr)
}
```

## API Reference

### Tools

The main type for accessing toolkit functions.

#### RandomString(n int) string

Generates a cryptographically secure random string of length `n`.

**Parameters:**
- `n int`: The desired length of the random string

**Returns:**
- `string`: A random string containing letters (upper and lowercase), numbers, and symbols (+, _)

**Example:**
```go
var tools toolkit.Tools
str := tools.RandomString(16) // Generates a 16-character random string
```

**Security Note:** Uses `crypto/rand` for cryptographically secure randomness, making it suitable for generating tokens, passwords, or other security-sensitive random values.

## Testing

Run the test suite:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## License

MIT License - see LICENSE file for details.
