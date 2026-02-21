# Changelog

All notable changes to Polyglot will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2026-02-21

### Added - Milestone 2 & 3 (Compiled + Extended Languages)

#### New Languages (11 added!)

**Milestone 2 - Compiled Languages:**
- Java (.java) - Full compile and run support with automatic cleanup
- C (.c) - Compiled with gcc
- C++ (.cpp, .cc, .cxx, .C) - Compiled with g++
- Rust (.rs) - Compiled with rustc

**Milestone 3 - Extended Languages:**
- Perl (.pl) - Interpreted scripting language
- Lua (.lua) - Lightweight scripting language
- Zig (.zig) - Modern systems programming language
- Nim (.nim) - Efficient compiled language
- Crystal (.cr) - Ruby-like compiled language
- D (.d) - Systems programming language
- Shell (.sh, .bash) - Bash shell scripts

#### Improvements
- Enhanced compilation workflow with automatic cleanup
- Better handling of compiled binaries
- Automatic removal of Java .class files
- Improved error messages for compilation failures
- Support for multiple file extensions per language

#### Total Language Count
- **16 languages** now supported (up from 5)
- **7 interpreted** languages
- **9 compiled** languages

### Technical Changes
- Updated language registry with 11 new handlers
- Enhanced executor to handle compilation cleanup
- Improved detector to support multiple extensions
- Updated check command to show all 16 toolchains

## [0.1.0] - 2026-02-21

### Added - Milestone 1 (MVP)

#### Core Features
- Automatic language detection from file extensions
- Process execution with proper output handling
- Exit code propagation from executed programs
- Error handling and user-friendly error messages

#### Language Support
- Python (.py) - interpreted via python3
- Go (.go) - compiled and run via go run
- JavaScript (.js) - interpreted via node
- Ruby (.rb) - interpreted via ruby
- PHP (.php) - interpreted via php

#### CLI Commands
- `polyglot run <file>` - Detect and run any supported file
- `polyglot compile <file>` - Compile without running (for compiled languages)
- `polyglot list` - Show all supported languages and their extensions
- `polyglot check` - Verify which language toolchains are installed
- `polyglot version` - Display version and milestone information

#### CLI Flags
- `-v, --verbose` - Show detailed execution information
- `-q, --quiet` - Suppress polyglot messages (show only program output)
- `--args` - Pass comma-separated arguments to the executed program

#### Documentation
- Comprehensive README with installation and usage instructions
- Quick Start Guide for getting started in 5 minutes
- Language Support Matrix showing all supported languages
- Examples document with real-world usage scenarios
- Contributing Guide for adding new languages
- 10-milestone Roadmap for future development
- Project Structure documentation
- Visual Demo of all features
- Milestone 1 Summary and Checklist

#### Build & Distribution
- Makefile with targets: build, install, test, check, fmt, clean, help
- Installation script (scripts/install.sh)
- Comprehensive test script (scripts/test-all.sh)
- GitHub Actions CI workflow for automated testing
- GitHub Actions Release workflow for multi-platform binaries
- Cross-platform support (Linux, macOS, Windows)

#### Testing
- Unit tests for language detector
- Test fixtures for all 5 supported languages
- Integration test suite
- Code formatting checks
- Go vet checks

#### Project Infrastructure
- Clean modular architecture with clear separation of concerns
- Interface-based design for easy extensibility
- MIT License
- .gitignore for Go projects
- go.mod with dependencies (cobra for CLI)

### Technical Details

#### Architecture
- `cmd/polyglot/` - Application entry point
- `internal/cli/` - CLI command implementations
- `internal/detector/` - Language detection logic
- `internal/executor/` - Process execution wrapper
- `internal/language/` - Language-specific handlers
- `pkg/types/` - Shared types and interfaces

#### Dependencies
- github.com/spf13/cobra v1.8.0 - CLI framework
- Go 1.21+ required

### Metrics
- ~1,500 lines of code
- 30+ files
- 5 languages supported
- 5 CLI commands
- 7 documentation files
- ~8MB binary size
- ~2 second build time

---

## [Unreleased]

### Planned for Milestone 2
- Java support with compile + run workflow
- C support with gcc
- C++ support with g++
- Rust support with rustc
- Two-step compilation workflow
- Automatic cleanup of build artifacts
- Better error messages with suggestions

### Planned for Milestone 3
- Perl, Lua, Zig, Nim, Crystal, D support
- Shell script support
- Language version detection
- Configurable compiler flags

### Planned for Milestone 4
- Esoteric language support (Brainfuck, Whitespace, etc.)
- Functional language support (Haskell, OCaml, Erlang, Elixir)
- JVM language support (Kotlin, Scala, Clojure)

### Planned for Milestone 5
- Configuration file support (~/.polyglot/config.yaml)
- Custom language definitions
- Per-project configuration
- Environment variable support

### Planned for Milestone 6
- Watch mode for development
- Colored output
- Progress indicators
- Better error messages with suggestions
- Dry run mode

---

## Version History

- **v0.1.0** (2026-02-21) - Milestone 1: MVP with 5 languages
- More versions coming soon...

---

## Links

- [GitHub Repository](https://github.com/yourusername/polyglot)
- [Issue Tracker](https://github.com/yourusername/polyglot/issues)
- [Roadmap](ROADMAP.md)
- [Contributing](CONTRIBUTING.md)
