# Contributing to Polyglot

Thanks for your interest in contributing! This guide will help you get started.

## Development Setup

1. Clone the repository
2. Ensure you have Go 1.21+ installed
3. Run `go mod tidy` to install dependencies
4. Build with `make build`

## Project Structure

- `cmd/polyglot/` - Entry point
- `internal/cli/` - CLI commands
- `internal/detector/` - Language detection
- `internal/executor/` - Process execution
- `internal/language/` - Language handlers
- `pkg/types/` - Shared types
- `test/fixtures/` - Test files

## Adding a New Language

To add support for a new language, follow these steps:

### 1. Add the language to types

Edit `pkg/types/types.go`:

```go
const (
    // ... existing languages
    YourLanguage
)

// Add to Name() method
case YourLanguage:
    return "YourLanguage"
```

### 2. Add detection

Edit `internal/detector/detector.go`:

```go
case ".ext":
    return types.YourLanguage, nil
```

### 3. Create a handler

Create `internal/language/yourlanguage.go`:

```go
package language

import (
    "os/exec"
    "github.com/yourusername/polyglot/pkg/types"
)

type YourLanguageHandler struct{}

func (h *YourLanguageHandler) Name() string {
    return "YourLanguage"
}

func (h *YourLanguageHandler) Extensions() []string {
    return []string{".ext"}
}

func (h *YourLanguageHandler) Type() types.LanguageType {
    return types.Interpreted // or types.Compiled
}

func (h *YourLanguageHandler) NeedsCompilation() bool {
    return false // or true for compiled languages
}

func (h *YourLanguageHandler) Compile(source string, output string) error {
    // Implement compilation if needed
    return nil
}

func (h *YourLanguageHandler) Run(file string, args []string) ([]byte, error) {
    cmdArgs := append([]string{file}, args...)
    cmd := exec.Command("your-interpreter", cmdArgs...)
    return cmd.CombinedOutput()
}
```

### 4. Register the handler

Edit `internal/language/registry.go`:

```go
var handlers = map[types.Language]types.LanguageHandler{
    // ... existing handlers
    types.YourLanguage: &YourLanguageHandler{},
}
```

### 5. Add test fixtures

Create `test/fixtures/hello.ext` with a simple hello world program.

### 6. Update documentation

- Add to README.md supported languages list
- Add to check.go toolchains map

### 7. Test

```bash
make build
./bin/polyglot run test/fixtures/hello.ext
go test ./...
```

## Testing

Run all tests:
```bash
make test
```

Run integration tests:
```bash
./scripts/test-all.sh
```

## Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Add comments for exported functions
- Keep functions small and focused

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests if applicable
5. Ensure all tests pass
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## Reporting Issues

When reporting issues, please include:

- Polyglot version (`polyglot version`)
- Operating system
- Language/toolchain versions
- Steps to reproduce
- Expected vs actual behavior

## Questions?

Feel free to open an issue for questions or discussions!
