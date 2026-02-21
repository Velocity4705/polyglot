# Polyglot - Project Overview

## Vision

Create a universal compiler and interpreter wrapper that makes it effortless to run programs in any programming language with a single, consistent command.

## The Problem

Developers working with multiple programming languages face friction:
- Different commands for different languages (python3, node, go run, javac + java, etc.)
- Need to remember language-specific flags and options
- Inconsistent interfaces across toolchains
- Tedious setup for multi-language projects

## The Solution

Polyglot provides a unified interface:
```bash
polyglot run <any-file>
```

It automatically:
1. Detects the programming language from file extension
2. Finds the appropriate compiler/interpreter
3. Executes with proper flags and options
4. Returns output and exit codes correctly

## Current Status (v0.1.0 - Milestone 1)

✅ **MVP Complete** - 5 languages supported with full CLI

### Supported Languages
- Python (.py)
- Go (.go)
- JavaScript (.js)
- Ruby (.rb)
- PHP (.php)

### Key Features
- Automatic language detection
- Unified run command
- Argument passing
- Verbose and quiet modes
- Toolchain verification
- Comprehensive documentation

## Architecture

### High-Level Design

```
User Input → CLI Parser → Language Detector → Language Handler → Executor → Output
```

### Components

1. **CLI Layer** (`internal/cli/`)
   - Command parsing (cobra framework)
   - Flag handling
   - User interaction

2. **Detector** (`internal/detector/`)
   - File extension analysis
   - Language identification
   - Error handling for unknown types

3. **Language Registry** (`internal/language/`)
   - Handler registration
   - Language-specific implementations
   - Extensible interface design

4. **Executor** (`internal/executor/`)
   - Process spawning
   - Output capture
   - Exit code handling

5. **Type System** (`pkg/types/`)
   - Language enum
   - Handler interface
   - Shared types

### Design Principles

1. **Extensibility**: Adding a new language requires minimal code
2. **Separation of Concerns**: Each component has a single responsibility
3. **Interface-Based**: Language handlers implement a common interface
4. **Error Handling**: Graceful failures with helpful messages
5. **Testability**: Components are independently testable

## File Structure

```
polyglot/
├── cmd/polyglot/              # Entry point
│   └── main.go
├── internal/
│   ├── cli/                   # CLI commands
│   │   ├── root.go
│   │   ├── run.go
│   │   ├── compile.go
│   │   ├── list.go
│   │   ├── check.go
│   │   └── version.go
│   ├── detector/              # Language detection
│   │   ├── detector.go
│   │   └── detector_test.go
│   ├── executor/              # Process execution
│   │   └── executor.go
│   └── language/              # Language handlers
│       ├── registry.go
│       ├── python.go
│       ├── go.go
│       ├── javascript.go
│       ├── ruby.go
│       └── php.go
├── pkg/types/                 # Shared types
│   └── types.go
├── test/fixtures/             # Test files
│   ├── hello.py
│   ├── hello.go
│   ├── hello.js
│   ├── hello.rb
│   ├── hello.php
│   ├── args.py
│   └── error.py
├── docs/                      # Documentation
│   ├── QUICKSTART.md
│   ├── LANGUAGES.md
│   └── EXAMPLES.md
├── scripts/                   # Automation
│   ├── install.sh
│   └── test-all.sh
├── .github/workflows/         # CI/CD
│   ├── ci.yml
│   └── release.yml
└── [documentation files]
```

## Development Workflow

### Adding a New Language

1. Add language constant to `pkg/types/types.go`
2. Add detection rule to `internal/detector/detector.go`
3. Create handler in `internal/language/<language>.go`
4. Register handler in `internal/language/registry.go`
5. Add test fixture in `test/fixtures/`
6. Update documentation

Time to add a new language: ~15 minutes

### Building

```bash
make build          # Build binary
make test           # Run tests
make check          # Code quality checks
make install        # Install globally
```

### Testing

```bash
make test           # Unit tests
make run-tests      # Integration tests
./scripts/test-all.sh  # Full test suite
```

## Technology Stack

- **Language**: Go 1.21+
- **CLI Framework**: Cobra
- **Testing**: Go standard testing
- **CI/CD**: GitHub Actions
- **Build**: Make

## Roadmap

### Milestone 2 (Next)
- Compiled languages (Java, C, C++, Rust)
- Two-step compilation workflow
- Build artifact management

### Milestone 3
- 20+ languages
- Language version detection
- Configurable compiler flags

### Milestone 4
- Esoteric languages
- Functional languages
- JVM languages

### Milestone 5
- Configuration files
- Custom language definitions
- Per-project settings

### Milestone 6-10
- Watch mode
- Colored output
- Performance optimization
- Distribution packages
- Community building

See [ROADMAP.md](ROADMAP.md) for complete plan.

## Contributing

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for:
- How to add a new language
- Code style guidelines
- Pull request process
- Issue reporting

## Documentation

- [README.md](README.md) - Main documentation
- [QUICKSTART.md](docs/QUICKSTART.md) - Get started in 5 minutes
- [DEMO.md](DEMO.md) - Visual demonstration
- [EXAMPLES.md](docs/EXAMPLES.md) - Usage examples
- [LANGUAGES.md](docs/LANGUAGES.md) - Language support matrix
- [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guide
- [ROADMAP.md](ROADMAP.md) - Development plan
- [CHANGELOG.md](CHANGELOG.md) - Version history
- [SUMMARY.md](SUMMARY.md) - Milestone 1 summary

## Metrics

- **Lines of Code**: ~1,500
- **Files**: 30+
- **Languages**: 5 (growing to 40+)
- **Commands**: 5
- **Build Time**: ~2 seconds
- **Binary Size**: ~8MB
- **Test Coverage**: Detector fully tested
- **Documentation**: 7 comprehensive guides

## Success Criteria

### Milestone 1 (Current) ✅
- [x] Support 3-5 languages
- [x] Working CLI with run command
- [x] Basic error handling
- [x] Documentation
- [x] Tests

### Long-term Goals
- [ ] 40+ languages supported
- [ ] 10,000+ GitHub stars
- [ ] Package manager distribution (brew, apt, etc.)
- [ ] Active community contributions
- [ ] IDE integrations
- [ ] 1.0 stable release

## License

MIT License - See [LICENSE](LICENSE) for details.

## Links

- **Repository**: https://github.com/yourusername/polyglot
- **Issues**: https://github.com/yourusername/polyglot/issues
- **Discussions**: https://github.com/yourusername/polyglot/discussions

## Contact

- Open an issue for bugs or feature requests
- Start a discussion for questions or ideas
- See CONTRIBUTING.md for contribution guidelines

---

**Status**: Milestone 1 Complete ✅  
**Version**: 0.1.0  
**Last Updated**: 2026-02-21
