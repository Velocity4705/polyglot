# Polyglot Architecture

## Overview

Polyglot is built with a modular architecture that separates concerns and makes it easy to extend.

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         CLI Layer                            │
│  (cmd/polyglot/main.go + internal/cli/*.go)                 │
│                                                              │
│  Commands: run, compile, list, check, version, install,     │
│            config (init, show, get, set, path)              │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    Configuration Layer                       │
│                  (internal/config/)                          │
│                                                              │
│  • Load global config (~/.polyglot/config.yaml)             │
│  • Load project config (.polyglot.yaml)                     │
│  • Merge configurations                                      │
│  • Provide configuration access                             │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                     Detection Layer                          │
│                  (internal/detector/)                        │
│                                                              │
│  • Detect language by file extension                        │
│  • Support custom extension mappings                        │
│  • Return language type                                     │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    Language Registry                         │
│                  (internal/language/)                        │
│                                                              │
│  • 30 language handlers                                     │
│  • Each implements LanguageHandler interface                │
│  • Provides compile/run logic per language                  │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                     Execution Layer                          │
│                  (internal/executor/)                        │
│                                                              │
│  • Execute language handlers                                │
│  • Handle compilation workflow                              │
│  • Manage build artifacts                                   │
│  • Apply environment variables                              │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    Installation Layer                        │
│                  (internal/installer/)                       │
│                                                              │
│  • Detect package manager                                   │
│  • Install missing toolchains                               │
│  • Interactive prompts                                       │
│  • System-wide installation                                 │
└─────────────────────────────────────────────────────────────┘
```

## Configuration System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    User Invokes Polyglot                     │
│                   $ polyglot run hello.py                    │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                  Load Configuration                          │
│                                                              │
│  1. Load default config (built-in)                          │
│  2. Load global config (~/.polyglot/config.yaml)            │
│  3. Load project config (.polyglot.yaml)                    │
│  4. Merge: default < global < project                       │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                  Detect Language                             │
│                                                              │
│  1. Check custom extensions first                           │
│  2. Fall back to built-in mappings                          │
│  3. Return language type                                    │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                  Check Language Enabled                      │
│                                                              │
│  • Check config.languages.<lang>.enabled                    │
│  • Default to true if not configured                        │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                  Get Language Handler                        │
│                                                              │
│  • Retrieve handler from registry                           │
│  • Apply custom commands if configured                      │
│  • Apply custom flags if configured                         │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                  Check Toolchain                             │
│                                                              │
│  • Check if toolchain is installed                          │
│  • If not, check auto_install.enabled                       │
│  • Install if enabled, prompt if not                        │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                  Execute Program                             │
│                                                              │
│  • Apply environment variables (global + language)          │
│  • Compile if needed                                        │
│  • Run program                                              │
│  • Clean up artifacts                                       │
└─────────────────────────────────────────────────────────────┘
```

## Module Breakdown

### cmd/polyglot/
Entry point for the application. Initializes Cobra CLI.

### internal/cli/
CLI command implementations:
- `root.go` - Root command setup
- `run.go` - Run command (main functionality)
- `compile.go` - Compile-only command
- `list.go` - List supported languages
- `check.go` - Check installed toolchains
- `version.go` - Version information
- `install.go` - Install toolchains
- `config.go` - Configuration management

### internal/config/
Configuration system:
- `config.go` - Configuration structure and loading
  - `Config` struct
  - `LanguageConfig` struct
  - Load/Save functions
  - Merge logic
  - Helper methods

### internal/detector/
Language detection:
- `detector.go` - File extension detection
  - `DetectLanguage()` - Basic detection
  - `DetectLanguageWithConfig()` - With custom extensions
- `detector_test.go` - Unit tests

### internal/executor/
Program execution:
- `executor.go` - Execution logic
  - `Executor` struct
  - `Run()` - Execute program
  - `Compile()` - Compile only
  - Artifact cleanup

### internal/installer/
Toolchain installation:
- `installer.go` - Installation logic
  - `Installer` struct
  - `Install()` - Install toolchain
  - `InstallIfMissing()` - Conditional install
- `package_manager.go` - Package manager detection
  - Detect system package manager
  - Map languages to packages
  - Execute installation commands

### internal/language/
Language handlers (30 files):
- `registry.go` - Language registry
- `<language>.go` - Individual handlers
  - Each implements `LanguageHandler` interface
  - Provides compile/run logic

### pkg/types/
Shared types:
- `types.go` - Type definitions
  - `Language` enum
  - `LanguageType` enum
  - `LanguageHandler` interface
  - `FromString()` converter

## Data Flow

### Run Command Flow

```
User Input
    ↓
Parse CLI Args
    ↓
Load Configuration (global + project)
    ↓
Detect Language (with custom extensions)
    ↓
Check Language Enabled
    ↓
Get Language Handler
    ↓
Check Toolchain Installed
    ↓
Install if Missing (auto-install)
    ↓
Apply Environment Variables
    ↓
Compile (if needed)
    ↓
Execute Program
    ↓
Clean Up Artifacts
    ↓
Return Exit Code
```

### Configuration Loading Flow

```
Default Config (built-in)
    ↓
Merge Global Config (~/.polyglot/config.yaml)
    ↓
Merge Project Config (.polyglot.yaml)
    ↓
Final Merged Config
```

### Custom Extension Detection Flow

```
File Extension (.pyx)
    ↓
Check Custom Extensions Map
    ↓
Found? → Return Mapped Language (Python)
    ↓
Not Found? → Check Built-in Mappings
    ↓
Found? → Return Language
    ↓
Not Found? → Return Error
```

## Key Design Patterns

### 1. Interface-Based Design

All language handlers implement the `LanguageHandler` interface:

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

This makes it easy to add new languages without modifying core logic.

### 2. Configuration Merging

Configurations are merged in order of precedence:
1. Default (lowest priority)
2. Global
3. Project (highest priority)

This allows for flexible customization at different levels.

### 3. Registry Pattern

Languages are registered in a central registry, making it easy to:
- Add new languages
- Look up handlers by language type
- List all supported languages

### 4. Command Pattern

Each CLI command is a separate module that implements Cobra's command interface.

### 5. Strategy Pattern

Different installation methods (system, script, binary) can be selected via configuration.

## Extension Points

### Adding a New Language

1. Create `internal/language/<language>.go`
2. Implement `LanguageHandler` interface
3. Register in `internal/language/registry.go`
4. Add to `pkg/types/types.go` enum
5. Add to `internal/detector/detector.go` mappings
6. Create test fixture in `test/fixtures/`

### Adding a New Package Manager

1. Add detection logic to `internal/installer/package_manager.go`
2. Add package mappings for all languages
3. Add installation command logic

### Adding a New Config Option

1. Add field to `Config` or `LanguageConfig` struct
2. Add getter/setter methods
3. Add to config command handlers
4. Update documentation

## Performance Considerations

### Fast Startup
- Minimal initialization
- Lazy loading of language handlers
- No unnecessary file I/O

### Efficient Detection
- Simple file extension matching
- O(1) lookup in maps
- No regex or complex parsing

### Clean Compilation
- Automatic artifact cleanup
- Temporary files removed after execution
- No disk space waste

## Security Considerations

### Safe Execution
- No shell injection (uses exec.Command properly)
- No eval or dynamic code execution
- Proper argument escaping

### Configuration Safety
- YAML parsing with validation
- No arbitrary code execution from config
- Safe default values

### Installation Safety
- Prompts before installing (by default)
- Uses system package managers
- No downloading arbitrary binaries (unless configured)

## Testing Strategy

### Unit Tests
- Detector tests (100% coverage)
- Configuration loading tests
- Language handler tests

### Integration Tests
- End-to-end execution tests
- Configuration merging tests
- Custom extension tests

### Manual Tests
- Test all 30 languages
- Test all package managers
- Test configuration scenarios

## Future Architecture Improvements

### Milestone 7+
- Plugin system for community extensions
- Language server protocol (LSP) integration
- Remote execution support
- Container-based isolation
- Build caching system
- Parallel compilation

---

**Version**: 0.4.0  
**Last Updated**: 2026-02-22
