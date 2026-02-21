# Polyglot - Universal Compiler Wrapper

[![CI](https://github.com/yourusername/polyglot/workflows/CI/badge.svg)](https://github.com/yourusername/polyglot/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/polyglot)](https://goreportcard.com/report/github.com/yourusername/polyglot)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)

A universal compiler and interpreter wrapper that detects programming languages by file extension and executes them with the appropriate toolchain.

## Current Status: Milestone 2 & 3 Complete! ✅

Currently supports 16 languages:

**Interpreted Languages:**
- Python (.py)
- JavaScript (.js)
- Ruby (.rb)
- PHP (.php)
- Perl (.pl)
- Lua (.lua)
- Shell (.sh, .bash)

**Compiled Languages:**
- Go (.go) - uses `go run`
- Java (.java)
- C (.c)
- C++ (.cpp, .cc, .cxx)
- Rust (.rs)
- Zig (.zig) - uses `zig run`
- Nim (.nim)
- Crystal (.cr) - uses `crystal run`
- D (.d)

## Installation

### Build from source

```bash
make build
```

### Install globally

```bash
make install
```

Or use the install script:

```bash
./scripts/install.sh
```

Or manually:

```bash
go build -o polyglot ./cmd/polyglot
sudo mv polyglot /usr/local/bin/
```

## Usage

### Run a file

```bash
polyglot run hello.py
polyglot run main.go
polyglot run app.js
```

### Pass arguments to your program

```bash
polyglot run script.py --args arg1,arg2,arg3
```

### Verbose mode

```bash
polyglot run -v hello.py
```

### Quiet mode (suppress polyglot messages)

```bash
polyglot run -q hello.py
```

### List supported languages

```bash
polyglot list
```

### Check installed toolchains

```bash
polyglot check
```

### Show version

```bash
polyglot version
```

### Compile only (for compiled languages)

```bash
polyglot compile main.go
```

## Examples

```bash
# Python
$ polyglot run hello.py
Detected: Python
Executing: python3 hello.py
Hello from Python!
Polyglot is working! 🐍

# Go
$ polyglot run main.go
Detected: Go
Executing: go run main.go
Hello from Go!
Polyglot is working! 🚀

# JavaScript
$ polyglot run app.js
Detected: JavaScript
Executing: node app.js
Hello from JavaScript!
Polyglot is working! ✨

# Ruby
$ polyglot run script.rb
Detected: Ruby
Executing: ruby script.rb
Hello from Ruby!
Polyglot is working! 💎

# PHP
$ polyglot run index.php
Detected: PHP
Executing: php index.php
Hello from PHP!
Polyglot is working! 🐘
```

See [DEMO.md](DEMO.md) for a complete visual demonstration.

## Testing

Run all tests:

```bash
make test
```

Run integration tests:

```bash
make run-tests
```

## Documentation

- [Quick Start Guide](docs/QUICKSTART.md) - Get started in 5 minutes
- [Demo](DEMO.md) - Visual demonstration of all features
- [Language Support Matrix](docs/LANGUAGES.md) - All supported languages
- [Examples](docs/EXAMPLES.md) - Comprehensive usage examples
- [Contributing Guide](CONTRIBUTING.md) - How to contribute
- [Roadmap](ROADMAP.md) - Future development plans
- [Project Structure](PROJECT_STRUCTURE.md) - Codebase organization
- [Summary](SUMMARY.md) - Milestone 1 completion summary

## Requirements

Polyglot is a **wrapper tool** that uses your system's installed language toolchains. You need to install the compilers/interpreters for the languages you want to use.

### Check What's Installed

```bash
polyglot check
```

### Install Language Toolchains

For the languages you want to use, install their toolchains:

**Fedora/RHEL:**
```bash
sudo dnf install python3 nodejs ruby php perl lua gcc g++ rustc
```

**Ubuntu/Debian:**
```bash
sudo apt install python3 nodejs ruby php perl lua build-essential rustc
```

**macOS:**
```bash
brew install python3 node ruby php perl lua gcc rust
```

Use `polyglot check` to see which toolchains you have installed.

### Future: Automatic Toolchain Management

We're planning to add automatic toolchain downloading in Milestone 5! See [TOOLCHAIN_MANAGEMENT.md](docs/TOOLCHAIN_MANAGEMENT.md) for details.

```bash
# Future feature (not yet implemented)
polyglot install python  # Will auto-download Python
polyglot install rust    # Will auto-download Rust
```

## Project Structure

```
polyglot/
├── cmd/polyglot/          # Entry point
├── internal/
│   ├── cli/               # CLI commands
│   ├── detector/          # Language detection
│   ├── executor/          # Process execution
│   └── language/          # Language handlers
├── pkg/types/             # Shared types
└── test/fixtures/         # Test files
```

## Roadmap

See [ROADMAP.md](ROADMAP.md) for the full development plan.

### Coming in Milestone 2
- Java, C, C++, Rust support
- Two-step compilation workflow
- Better error handling

## Contributing

Contributions welcome! See [ROADMAP.md](ROADMAP.md) for planned features.

## License

MIT
