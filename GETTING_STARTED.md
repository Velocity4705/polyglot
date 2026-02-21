# Getting Started with Polyglot Development

Welcome! This guide will help you set up your development environment and make your first contribution to Polyglot.

## Prerequisites

- Go 1.21 or higher
- Git
- A code editor (VS Code, GoLand, vim, etc.)
- At least one language toolchain (python3, node, go, ruby, or php)

## Quick Setup (5 minutes)

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/polyglot.git
cd polyglot
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build the Project

```bash
make build
```

You should see:
```
Building polyglot...
✓ Build complete: bin/polyglot
```

### 4. Run Tests

```bash
make test
```

### 5. Try It Out

```bash
./bin/polyglot run test/fixtures/hello.py
./bin/polyglot run test/fixtures/hello.go
./bin/polyglot run test/fixtures/hello.js
```

## Development Workflow

### Making Changes

1. Create a new branch:
```bash
git checkout -b feature/my-feature
```

2. Make your changes

3. Format your code:
```bash
make fmt
```

4. Run quality checks:
```bash
make check
```

5. Run tests:
```bash
make test
```

6. Test manually:
```bash
make build
./bin/polyglot run test/fixtures/hello.py
```

### Common Tasks

#### Build
```bash
make build
```

#### Run Tests
```bash
make test
```

#### Format Code
```bash
make fmt
```

#### Check Code Quality
```bash
make check
```

#### Clean Build Artifacts
```bash
make clean
```

#### See All Commands
```bash
make help
```

## Project Structure Overview

```
polyglot/
├── cmd/polyglot/              # Entry point (main.go)
├── internal/
│   ├── cli/                   # CLI commands
│   ├── detector/              # Language detection
│   ├── executor/              # Process execution
│   └── language/              # Language handlers
├── pkg/types/                 # Shared types
├── test/fixtures/             # Test files
└── docs/                      # Documentation
```

## Adding Your First Language

Let's add support for Perl as an example!

### Step 1: Add Language Type

Edit `pkg/types/types.go`:

```go
const (
    // ... existing languages
    Perl
)

// In Name() method:
case Perl:
    return "Perl"
```

### Step 2: Add Detection

Edit `internal/detector/detector.go`:

```go
case ".pl":
    return types.Perl, nil
```

### Step 3: Create Handler

Create `internal/language/perl.go`:

```go
package language

import (
    "os/exec"
    "github.com/yourusername/polyglot/pkg/types"
)

type PerlHandler struct{}

func (h *PerlHandler) Name() string {
    return "Perl"
}

func (h *PerlHandler) Extensions() []string {
    return []string{".pl"}
}

func (h *PerlHandler) Type() types.LanguageType {
    return types.Interpreted
}

func (h *PerlHandler) NeedsCompilation() bool {
    return false
}

func (h *PerlHandler) Compile(source string, output string) error {
    return nil
}

func (h *PerlHandler) Run(file string, args []string) ([]byte, error) {
    cmdArgs := append([]string{file}, args...)
    cmd := exec.Command("perl", cmdArgs...)
    return cmd.CombinedOutput()
}
```

### Step 4: Register Handler

Edit `internal/language/registry.go`:

```go
var handlers = map[types.Language]types.LanguageHandler{
    // ... existing handlers
    types.Perl: &PerlHandler{},
}
```

### Step 5: Add Test Fixture

Create `test/fixtures/hello.pl`:

```perl
#!/usr/bin/env perl
print "Hello from Perl!\n";
print "Polyglot is working! 🐪\n";
```

### Step 6: Test It

```bash
make build
./bin/polyglot run test/fixtures/hello.pl
```

### Step 7: Update Documentation

Add Perl to:
- `README.md` (supported languages)
- `docs/LANGUAGES.md` (language matrix)
- `internal/cli/check.go` (toolchains map)

### Step 8: Commit

```bash
git add .
git commit -m "Add Perl support"
git push origin feature/add-perl
```

Then open a Pull Request!

## Understanding the Code

### How Language Detection Works

1. User runs: `polyglot run hello.py`
2. CLI parses command and file path
3. Detector extracts extension: `.py`
4. Detector maps to language: `Python`
5. Registry returns handler: `PythonHandler`
6. Executor runs: `python3 hello.py`
7. Output is captured and displayed

### Key Interfaces

#### LanguageHandler Interface

```go
type LanguageHandler interface {
    Name() string
    Extensions() []string
    Type() LanguageType
    Compile(source string, output string) error
    Run(file string, args []string) ([]byte, error)
    NeedsCompilation() bool
}
```

Every language must implement this interface.

### Code Flow

```
main.go
  → cli/root.go (cobra setup)
    → cli/run.go (run command)
      → detector/detector.go (detect language)
        → language/registry.go (get handler)
          → executor/executor.go (execute)
            → language/python.go (language-specific)
```

## Debugging Tips

### Enable Verbose Mode

```bash
./bin/polyglot run -v hello.py
```

### Check What's Installed

```bash
./bin/polyglot check
```

### Run with Go Directly

```bash
go run ./cmd/polyglot run test/fixtures/hello.py
```

### Use Print Debugging

Add to your code:
```go
fmt.Printf("DEBUG: variable = %v\n", variable)
```

### Run Specific Tests

```bash
go test ./internal/detector -v
```

## Common Issues

### "command not found: polyglot"

You're trying to run `polyglot` but it's not in your PATH. Use:
```bash
./bin/polyglot run hello.py
```

Or install it:
```bash
make install
```

### "unsupported file extension"

The language isn't supported yet. Check:
```bash
./bin/polyglot list
```

### Build Errors

Clean and rebuild:
```bash
make clean
make build
```

### Test Failures

Make sure you have the language toolchains installed:
```bash
./bin/polyglot check
```

## Resources

### Documentation
- [README.md](README.md) - Main documentation
- [CONTRIBUTING.md](CONTRIBUTING.md) - Detailed contribution guide
- [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - Architecture details
- [ROADMAP.md](ROADMAP.md) - Future plans

### Examples
- [DEMO.md](DEMO.md) - Visual demonstration
- [EXAMPLES.md](docs/EXAMPLES.md) - Usage examples
- [test/fixtures/](test/fixtures/) - Example programs

### Go Resources
- [Go Documentation](https://go.dev/doc/)
- [Cobra Documentation](https://github.com/spf13/cobra)
- [Effective Go](https://go.dev/doc/effective_go)

## Getting Help

- 💬 Open an [Issue](https://github.com/yourusername/polyglot/issues)
- 📖 Read the [Contributing Guide](CONTRIBUTING.md)
- 🎯 Check the [Roadmap](ROADMAP.md)
- 📚 Browse [Examples](docs/EXAMPLES.md)

## Next Steps

1. ✅ Set up your development environment
2. ✅ Build and test the project
3. ✅ Understand the code structure
4. 📝 Pick an issue to work on
5. 🚀 Make your first contribution!

### Good First Issues

- Add support for a new language (Perl, Lua, etc.)
- Improve error messages
- Add more test fixtures
- Improve documentation
- Add examples

### Bigger Projects

- Implement Milestone 2 features (Java, C, C++, Rust)
- Add configuration file support
- Implement watch mode
- Add colored output

## Welcome to the Team!

We're excited to have you contribute to Polyglot. Don't hesitate to ask questions or request help. Happy coding! 🚀

---

**Ready to contribute?** Check out [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.
