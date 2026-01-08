# Examples

Runnable examples demonstrating `go-xerrs` features.

## Table of Contents

| Example | Description | Run |
| ------- | ----------- | --- |
| [basic](./basic/) | Basic error creation and configuration | `cd basic && go run main.go` |
| [chaining](./chaining/) | Fluent error type conversion and chaining | `cd chaining && go run main.go` |
| [wrapping](./wrapping/) | Error wrapping and automatic detection | `cd wrapping && go run main.go` |

## Quick Start

```bash
# Clone the repository
git clone https://github.com/hotfixfirst/go-xerrs.git
cd go-xerrs/_examples

# Run a specific example
cd basic && go run main.go
```

## Adding New Examples

When adding a new feature example:

1. Create a new directory: `_examples/{feature}/`
2. Add `main.go` with runnable code
3. Add `README.md` with documentation
4. Update this file's table of contents
