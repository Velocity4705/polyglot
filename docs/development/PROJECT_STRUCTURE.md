# Polyglot - Project File Structure

```
polyglot/
│
├── cmd/
│   └── polyglot/
│       └── main.go                 # Entry point, CLI setup
│
├── internal/
│   ├── detector/
│   │   ├── detector.go             # Language detection logic
│   │   └── detector_test.go
│   │
│   ├── language/
│   │   ├── language.go             # Language interface & registry
│   │   ├── registry.go             # Language registry management
│   │   ├── python.go               # Python handler
│   │   ├── go.go                   # Go handler
│   │   ├── javascript.go           # JavaScript/Node.js handler
│   │   ├── java.go                 # Java handler
│   │   ├── c.go                    # C handler
│   │   ├── cpp.go                  # C++ handler
│   │   ├── rust.go                 # Rust handler
│   │   ├── ruby.go                 # Ruby handler
│   │   ├── zig.go                  # Zig handler
│   │   ├── assembly.go             # Assembly handler
│   │   └── ...                     # More language handlers
│   │
│   ├── executor/
│   │   ├── executor.go             # Process execution wrapper
│   │   ├── compiler.go             # Compilation logic
│   │   ├── runner.go               # Runtime execution logic
│   │   └── executor_test.go
│   │
│   ├── config/
│   │   ├── config.go               # Configuration loading
│   │   ├── defaults.go             # Default language configs
│   │   └── config_test.go
│   │
│   └── cli/
│       ├── run.go                  # 'run' command implementation
│       ├── compile.go              # 'compile' command implementation
│       ├── list.go                 # 'list' command (show supported languages)
│       └── version.go              # 'version' command
│
├── pkg/
│   └── types/
│       └── types.go                # Shared types and interfaces
│
├── configs/
│   ├── languages.yaml              # Default language definitions
│   └── example.polyglot.yaml       # Example project config
│
├── scripts/
│   ├── install.sh                  # Installation script
│   └── build.sh                    # Build script
│
├── test/
│   ├── fixtures/                   # Test files in various languages
│   │   ├── hello.py
│   │   ├── hello.go
│   │   ├── Hello.java
│   │   ├── hello.c
│   │   ├── hello.rs
│   │   └── ...
│   └── integration/
│       └── integration_test.go     # Integration tests
│
├── docs/
│   ├── LANGUAGES.md                # Supported languages matrix
│   ├── CONFIGURATION.md            # Configuration guide
│   ├── CONTRIBUTING.md             # Contribution guidelines
│   └── examples/                   # Usage examples
│
├── .github/
│   └── workflows/
│       ├── ci.yml                  # CI pipeline
│       └── release.yml             # Release automation
│
├── go.mod                          # Go module definition
├── go.sum                          # Go dependencies
├── Makefile                        # Build automation
├── README.md                       # Project README
├── LICENSE                         # License file
├── ROADMAP.md                      # Development roadmap
└── PROJECT_STRUCTURE.md            # This file
```

## Key Components Explained

### `/cmd/polyglot/`
Entry point for the application. Handles CLI initialization and command routing.

### `/internal/detector/`
Language detection based on file extensions. Maps `.py` → Python, `.go` → Go, etc.

### `/internal/language/`
Individual language handlers. Each file implements the Language interface with compile/run logic.

### `/internal/executor/`
Low-level process execution. Handles spawning compilers/interpreters, capturing output, and error handling.

### `/internal/config/`
Configuration management. Loads user configs, merges with defaults, handles environment variables.

### `/internal/cli/`
Command implementations (run, compile, list, etc.). Uses cobra or similar CLI framework.

### `/pkg/types/`
Public types and interfaces that could be used by external packages.

### `/configs/`
Default language definitions and example configurations.

### `/test/`
Test files and fixtures for all supported languages.

### `/docs/`
Documentation, guides, and examples.

## Language Handler Interface

Each language handler implements:

```go
type Language interface {
    Name() string
    Extensions() []string
    Detect(filename string) bool
    Compile(source string, output string) error
    Run(file string, args []string) error
    CompileAndRun(source string, args []string) error
}
```

## Configuration File Format

```yaml
# ~/.polyglot/config.yaml or .polyglot.yaml

languages:
  python:
    command: python3
    extensions: [.py]
    type: interpreted
    
  java:
    compiler: javac
    runner: java
    extensions: [.java]
    type: compiled
    compile_flags: []
    
  rust:
    compiler: rustc
    extensions: [.rs]
    type: compiled
    compile_flags: [-O]
```
