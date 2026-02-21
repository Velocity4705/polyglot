# Polyglot - Universal Compiler Roadmap

## Project Overview
A universal compiler/interpreter wrapper that detects programming languages by file extension and executes them with the appropriate toolchain.

---

## Milestone 1: Core Foundation (MVP)
**Goal:** Basic CLI that can detect and run 3-5 popular languages

### Features
- [ ] CLI argument parsing (run, compile commands)
- [ ] File extension detection
- [ ] Language registry system
- [ ] Process execution wrapper
- [ ] Support for: Python, Go, JavaScript (Node.js)
- [ ] Basic error handling
- [ ] Exit code propagation

### Deliverable
Working `polyglot run <file>` command for 3-5 languages

---

## Milestone 2: Compiled Languages Support
**Goal:** Add languages that require compilation step

### Features
- [ ] Two-step execution (compile → run)
- [ ] Temporary build artifact management
- [ ] Support for: Java, C, C++, Rust
- [ ] `polyglot compile` command (compile only, no execution)
- [ ] Cleanup of intermediate files

### Deliverable
Full compile + run workflow for compiled languages

---

## Milestone 3: Advanced Language Support
**Goal:** Expand to more languages and edge cases

### Features
- [ ] Support for: Ruby, PHP, Perl, Lua
- [ ] Support for: Zig, Nim, Crystal, D
- [ ] Support for: Assembly (NASM, FASM)
- [ ] Support for: Shell scripts (bash, zsh)
- [ ] Language version detection (python2 vs python3)
- [ ] Configurable compiler flags per language

### Deliverable
20+ language support with robust detection

---

## Milestone 4: Esoteric & Specialty Languages
**Goal:** Support for unusual and esoteric languages

### Features
- [ ] Support for: Brainfuck, Whitespace, Befunge
- [ ] Support for: LOLCODE, Chicken, Malbolge
- [ ] Support for: Haskell, OCaml, F#, Erlang, Elixir
- [ ] Support for: Kotlin, Scala, Clojure
- [ ] Custom interpreter path configuration

### Deliverable
40+ languages including esoteric ones

---

## Milestone 5: Auto-Install System ⭐ NEW!
**Goal:** Automatic system-wide toolchain installation

### Features
- [ ] Detect missing toolchains when running code
- [ ] Prompt user to install (with details)
- [ ] Auto-install using system package manager (dnf/apt/brew/pacman)
- [ ] `polyglot install <language>` command
- [ ] Package name mapping for all 16+ languages
- [ ] Sudo handling and security prompts
- [ ] Installation progress indicators
- [ ] Error handling and fallbacks
- [ ] Config file: `~/.polyglot/config.yaml`
- [ ] Auto-install mode (skip prompts)
- [ ] Manual mode (show instructions only)
- [ ] `polyglot status` command

### User Experience
```bash
$ polyglot run hello.py
⚠️  Python not found
📦 Install Python 3.12 system-wide? [Y/n] y
   Installing via dnf... ✓
✅ Python installed!
🚀 Running hello.py...
```

### Benefits
- ✅ Zero-config experience - just run your code
- ✅ Languages installed system-wide (works with all tools)
- ✅ One-time installation per language
- ✅ Integrates with system package manager
- ✅ Natural learning curve - install as you need

### Deliverable
Truly universal experience - Polyglot handles everything!

See [AUTO_INSTALL_DESIGN.md](docs/AUTO_INSTALL_DESIGN.md) for complete design.

---

## Milestone 6: Configuration & Customization
**Goal:** Advanced configuration and customization

### Features
- [ ] Enhanced config file (~/.polyglot/config.yaml)
- [ ] Custom language definitions
- [ ] Per-project config (.polyglot.yaml)
- [ ] Environment variable support
- [ ] Compiler flag templates
- [ ] Custom file extension mappings
- [ ] Version management per language
- [ ] Alternative installation methods (scripts, binaries)

### Deliverable
Fully customizable language toolchain configuration

---

## Milestone 7: Developer Experience
**Goal:** Polish the UX and add helpful features

### Features
- [ ] Verbose mode (`-v, --verbose`)
- [ ] Quiet mode (`-q, --quiet`)
- [ ] Dry run mode (`--dry-run`)
- [ ] Pass arguments to programs (`--args`)
- [ ] Watch mode for development (`polyglot watch`)
- [ ] Colored output
- [ ] Progress indicators for long compilations
- [ ] Better error messages with suggestions

### Deliverable
Production-ready CLI with excellent UX

---

## Milestone 7: Performance & Optimization
**Goal:** Make it fast and efficient

### Features
- [ ] Parallel compilation support
- [ ] Build caching
- [ ] Incremental compilation detection
- [ ] Benchmark mode
- [ ] Memory usage optimization
- [ ] Startup time optimization

### Deliverable
Fast, efficient execution with caching

---

## Milestone 8: Testing & Quality
**Goal:** Comprehensive testing and reliability

### Features
- [ ] Unit tests for all language handlers
- [ ] Integration tests
- [ ] CI/CD pipeline
- [ ] Cross-platform testing (Linux, macOS, Windows)
- [ ] Error scenario testing
- [ ] Performance benchmarks

### Deliverable
95%+ test coverage, reliable across platforms

---

## Milestone 9: Distribution & Packaging
**Goal:** Easy installation and distribution

### Features
- [ ] Homebrew formula (macOS/Linux)
- [ ] apt/yum packages (Linux)
- [ ] Chocolatey package (Windows)
- [ ] Docker image
- [ ] Pre-built binaries for all platforms
- [ ] Auto-update mechanism
- [ ] Installation script

### Deliverable
Easy installation on all major platforms

---

## Milestone 10: Documentation & Community
**Goal:** Complete documentation and community building

### Features
- [ ] Comprehensive README
- [ ] Language support matrix
- [ ] Contributing guide
- [ ] API documentation
- [ ] Example projects
- [ ] Video tutorials
- [ ] Website/landing page
- [ ] Plugin system for community extensions

### Deliverable
Complete documentation, active community

---

## Future Ideas (Post 1.0)
- Language server protocol (LSP) integration
- IDE plugins (VSCode, IntelliJ, etc.)
- Remote execution support
- Container-based isolated execution
- Language playground web interface
- Package manager integration
- Dependency management
- Multi-file project support
- Build system integration (Make, CMake, etc.)
