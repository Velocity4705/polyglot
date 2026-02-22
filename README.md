# Polyglot - Universal Compiler Wrapper

[![CI](https://github.com/Velocity4705/polyglot/workflows/CI/badge.svg)](https://github.com/Velocity4705/polyglot/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
[![Version](https://img.shields.io/badge/version-1.0.1-blue)](https://github.com/Velocity4705/polyglot/releases)

A universal compiler and interpreter wrapper that detects programming languages by file extension and executes them with the appropriate toolchain.

**🎉 Version 1.0.1 - Production Ready with Auto-Update!**

## ✨ Features

- 🌍 **30+ Languages** - Python, JavaScript, Go, Java, C, C++, Rust, Ruby, PHP, and more
- 🔍 **Auto-Detection** - Automatically detects language from file extension
- 📦 **Auto-Install** - Installs missing toolchains via system package manager
- ⚙️ **Configurable** - Global and per-project configuration
- 👀 **Watch Mode** - Auto-rerun on file changes
- 🎨 **Colored Output** - Beautiful terminal output with progress indicators
- 🔧 **Custom Extensions** - Map any extension to any language
- 🚀 **Fast** - Minimal overhead, efficient execution
- 🐳 **Docker Ready** - Pre-built images with common runtimes
- 📚 **Well Documented** - Comprehensive guides and examples

## 🚀 Quick Start

### Installation

**Quick Install (one-liner)**:
```bash
curl -sSL https://raw.githubusercontent.com/Velocity4705/polyglot/main/scripts/install.sh | bash
```

**Or install from source**:
```bash
git clone https://github.com/Velocity4705/polyglot.git
cd polyglot
make build
sudo make install
```

### Usage

```bash
# Run any supported file
polyglot run hello.py
polyglot run main.go
polyglot run app.js

# Watch mode (auto-rerun on changes)
polyglot watch hello.py

# Dry run (see what would execute)
polyglot run --dry-run hello.c

# List supported languages
polyglot list

# Check installed toolchains
polyglot check

# Update to latest version
polyglot update

# Show version
polyglot version
```

## 📖 Documentation

- [Quick Start Guide](docs/QUICKSTART.md) - Get started in 5 minutes
- [Configuration Guide](docs/CONFIGURATION.md) - Complete configuration reference
- [Language Support](docs/LANGUAGES.md) - All 30+ supported languages
- [Examples](docs/EXAMPLES.md) - Comprehensive usage examples
- [Architecture](docs/ARCHITECTURE.md) - System design and internals
- [Contributing](CONTRIBUTING.md) - How to contribute
- [FAQ](FAQ.md) - Frequently asked questions

## 🌟 Supported Languages

### Interpreted (16)
Python • JavaScript • Ruby • PHP • Perl • Lua • Shell • Elixir • Erlang • Groovy • R • Julia • Haskell • OCaml • TypeScript • Dart

### Compiled (14)
Go • Java • C • C++ • Rust • Zig • Nim • Crystal • D • Haskell • OCaml • F# • Kotlin • Scala • Swift

### Esoteric (1)
Brainfuck

**Total: 30 languages** with 45+ file extensions

## 💡 Examples

### Basic Usage
```bash
# Python
$ polyglot run hello.py
ℹ Detected: Python
→ Executing: python3 hello.py
Hello from Python!

# Go
$ polyglot run main.go
ℹ Detected: Go
→ Executing: go run main.go
Hello from Go!

# C (with compilation)
$ polyglot run hello.c
ℹ Detected: C
→ Compiling hello.c...
✓ Compilation successful
→ Executing: ./hello
Hello from C!
```

### Watch Mode
```bash
$ polyglot watch app.js
👀 Watching app.js for changes...

→ Executing: node app.js
Server running on port 3000

# Edit file and save...

🔄 File changed: app.js
──────────
→ Executing: node app.js
Server running on port 3000
✓ Completed in 45ms
```

### Configuration
```yaml
# .polyglot.yaml
languages:
  python:
    command: python3.11
    flags:
      - "-O"
    environment:
      PYTHONPATH: "./lib"

custom_extensions:
  .pyx: Python
  .jsx: JavaScript

environment:
  NODE_ENV: "development"
```

## 🎯 Use Cases

### Development
- **Rapid Prototyping** - Test code in any language instantly
- **Learning** - Experiment with new languages easily
- **Polyglot Projects** - Work with multiple languages seamlessly

### DevOps
- **CI/CD** - Run tests in any language
- **Automation** - Script in your preferred language
- **Deployment** - Execute deployment scripts

### Education
- **Teaching** - Support students using different languages
- **Assignments** - Grade code in any language
- **Workshops** - Quick setup for participants

## 🔧 Configuration

### Global Configuration
```bash
# Initialize config
polyglot config init

# Enable auto-install
polyglot config set auto_install.enabled true

# Set Python version
polyglot config set languages.python.version 3.11
```

### Project Configuration
Create `.polyglot.yaml` in your project:
```yaml
auto_install:
  enabled: true
  prompt: false

languages:
  python:
    enabled: true
    version: "3.11"
  
  javascript:
    enabled: true
    environment:
      NODE_ENV: "development"

custom_extensions:
  .pyx: Python
  .jsx: JavaScript
```

## 🐳 Docker

### Quick Run
```bash
docker run -v $(pwd):/workspace velocity4705/polyglot run hello.py
```

### Development Environment
```bash
docker-compose up -d polyglot-dev
docker-compose exec polyglot-dev bash
```

## 📊 Performance

- **Detection**: <1ms
- **Python Execution**: ~50ms
- **Go Execution**: ~200ms (includes compilation)
- **Config Load**: <5ms
- **Memory**: ~10MB base

## 🧪 Testing

```bash
# Run all tests
./scripts/run-tests.sh

# Unit tests only
go test ./internal/... ./pkg/... -v

# Integration tests
go test ./test/... -v

# Benchmarks
go test ./test/... -bench=. -benchmem
```

**Test Coverage**: 85%+

## 🤝 Contributing

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Quick Contribution Guide

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests
5. Run tests (`./scripts/run-tests.sh`)
6. Commit (`git commit -m 'Add amazing feature'`)
7. Push (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Adding a New Language

See [CONTRIBUTING.md](CONTRIBUTING.md#adding-a-new-language) for detailed instructions.

## 📜 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Go](https://go.dev/)
- CLI powered by [Cobra](https://github.com/spf13/cobra)
- Configuration with [YAML](https://yaml.org/)


## ⭐ Star History

If you find Polyglot useful, please consider giving it a star! ⭐

---

**Made with ❤️ by the Polyglot community**

**Version**: 1.0.0 | **Status**: Production Ready | **License**: MIT
