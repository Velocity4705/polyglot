# Polyglot Demo

A visual demonstration of all Polyglot features.

## Installation

```bash
$ make build
Building polyglot...
✓ Build complete: bin/polyglot
```

## Version Information

```bash
$ ./bin/polyglot version
Polyglot v0.1.0
Milestone 1 (MVP)
Universal compiler and interpreter wrapper
```

## Check Installed Toolchains

```bash
$ ./bin/polyglot check
Checking installed toolchains...
================================
Python          python3         ✓ INSTALLED
Go              go              ✓ INSTALLED
JavaScript      node            ✓ INSTALLED
Ruby            ruby            ✗ NOT FOUND
PHP             php             ✗ NOT FOUND
================================
Summary: 3/5 toolchains installed

To use all languages, install the missing toolchains.
```

## List Supported Languages

```bash
$ ./bin/polyglot list
Supported Languages:
-------------------
Python          [.py]
Go              [.go]
JavaScript      [.js]
Ruby            [.rb]
PHP             [.php]
```

## Running Programs

### Python

```bash
$ ./bin/polyglot run test/fixtures/hello.py
Detected: Python
Executing: python3 test/fixtures/hello.py
Hello from Python!
Polyglot is working! 🐍
```

### Go

```bash
$ ./bin/polyglot run test/fixtures/hello.go
Detected: Go
Executing: go run test/fixtures/hello.go
Hello from Go!
Polyglot is working! 🚀
```

### JavaScript

```bash
$ ./bin/polyglot run test/fixtures/hello.js
Detected: JavaScript
Executing: node test/fixtures/hello.js
Hello from JavaScript!
Polyglot is working! ✨
```

## Passing Arguments

```bash
$ ./bin/polyglot run test/fixtures/args.py --args hello,world,test
Detected: Python
Executing: python3 test/fixtures/args.py hello world test
Arguments received: 3
  arg1: hello
  arg2: world
  arg3: test
```

## Quiet Mode

```bash
$ ./bin/polyglot run -q test/fixtures/hello.py
Hello from Python!
Polyglot is working! 🐍
```

## Verbose Mode

```bash
$ ./bin/polyglot run -v test/fixtures/hello.py
Detected: Python
Executing: python3 test/fixtures/hello.py
Hello from Python!
Polyglot is working! 🐍
```

## Error Handling

### File Not Found

```bash
$ ./bin/polyglot run nonexistent.py
Error: file not found: nonexistent.py
```

### Unsupported Extension

```bash
$ ./bin/polyglot run file.xyz
Error: unsupported file extension: .xyz
```

### Program Error

```bash
$ ./bin/polyglot run test/fixtures/error.py
Detected: Python
Executing: python3 test/fixtures/error.py
This script will exit with an error code
Error: program exited with error
```

## Help System

```bash
$ ./bin/polyglot --help
Polyglot detects programming languages and executes them with the appropriate toolchain.

Usage:
  polyglot [command]

Available Commands:
  check       Check which language toolchains are installed
  compile     Compile a source file without running it
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all supported languages
  run         Detect and run a source file
  version     Print version information

Flags:
  -h, --help   help for polyglot

Use "polyglot [command] --help" for more information about a command.
```

### Command-Specific Help

```bash
$ ./bin/polyglot run --help
Detect and run a source file

Usage:
  polyglot run <file> [flags]

Flags:
      --args strings   Arguments to pass to the program
  -h, --help           help for run
  -q, --quiet          Quiet mode (suppress polyglot messages)
  -v, --verbose        Verbose output
```

## Testing

### Unit Tests

```bash
$ make test
Running tests...
?   	github.com/yourusername/polyglot/cmd/polyglot	[no test files]
?   	github.com/yourusername/polyglot/internal/cli	[no test files]
ok  	github.com/yourusername/polyglot/internal/detector	0.002s
?   	github.com/yourusername/polyglot/internal/executor	[no test files]
?   	github.com/yourusername/polyglot/internal/language	[no test files]
?   	github.com/yourusername/polyglot/pkg/types	[no test files]
?   	github.com/yourusername/polyglot/test/fixtures	[no test files]
```

### Integration Tests

```bash
$ ./scripts/test-all.sh
================================
Polyglot Test Suite
================================

Building polyglot...
✓ Build successful

Running unit tests...
✓ Unit tests passed

Checking toolchains...
[output from polyglot check]

================================
Testing Language Support
================================

Testing Python...
✓ Python works
Testing Go...
✓ Go works
Testing JavaScript...
✓ JavaScript works
Testing Ruby...
✗ Ruby failed (toolchain may not be installed)
Testing PHP...
✗ PHP failed (toolchain may not be installed)

================================
All tests completed!
================================
```

### Code Quality

```bash
$ make check
Running go vet...
Checking formatting...
✓ All checks passed
```

## Real-World Usage

### Quick Script Testing

```bash
# Write a quick Python script
$ echo 'print("Testing polyglot!")' > test.py

# Run it
$ polyglot run test.py
Detected: Python
Executing: python3 test.py
Testing polyglot!
```

### Multi-Language Project

```bash
# Run Python data generator
$ polyglot run generate_data.py

# Process with JavaScript
$ polyglot run process_data.js

# Analyze with Go
$ polyglot run analyze.go
```

### CI/CD Integration

```bash
# In your CI pipeline
$ polyglot run tests/test.py
$ polyglot run tests/test.go
$ polyglot run tests/test.js
```

## Summary

Polyglot provides a unified interface for running programs in multiple languages:

- ✅ Automatic language detection
- ✅ 5 languages supported (Python, Go, JavaScript, Ruby, PHP)
- ✅ Simple, intuitive CLI
- ✅ Proper error handling
- ✅ Argument passing
- ✅ Quiet and verbose modes
- ✅ Toolchain verification
- ✅ Comprehensive help system

**One command to run them all!** 🚀
