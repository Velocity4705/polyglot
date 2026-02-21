# Polyglot Examples

Comprehensive examples showing all features of Polyglot.

## Basic Usage

### Run a Python script

```bash
$ polyglot run hello.py
Detected: Python
Executing: python3 hello.py
Hello from Python!
```

### Run a Go program

```bash
$ polyglot run main.go
Detected: Go
Executing: go run main.go
Hello from Go!
```

### Run JavaScript

```bash
$ polyglot run app.js
Detected: JavaScript
Executing: node app.js
Hello from JavaScript!
```

## Passing Arguments

### Create a script that uses arguments

```python
# greet.py
import sys

if len(sys.argv) < 2:
    print("Usage: polyglot run greet.py --args <name>")
    sys.exit(1)

name = sys.argv[1]
print(f"Hello, {name}!")
```

### Run with arguments

```bash
$ polyglot run greet.py --args Alice
Detected: Python
Executing: python3 greet.py Alice
Hello, Alice!
```

### Multiple arguments

```bash
$ polyglot run script.py --args arg1,arg2,arg3
```

## Quiet Mode

Suppress Polyglot's detection and execution messages:

```bash
$ polyglot run -q hello.py
Hello from Python!
```

Perfect for piping output:

```bash
$ polyglot run -q data.py | jq .
```

## Verbose Mode

See detailed execution information:

```bash
$ polyglot run -v hello.py
Detected: Python
Executing: python3 hello.py
Hello from Python!
```

## Combining Flags

```bash
$ polyglot run -v hello.py --args world
```

## Error Handling

### Script with error

```python
# error.py
import sys
print("This will fail")
sys.exit(1)
```

```bash
$ polyglot run error.py
Detected: Python
Executing: python3 error.py
This will fail
Error: program exited with error
```

### Non-existent file

```bash
$ polyglot run nonexistent.py
Error: file not found: nonexistent.py
```

### Unsupported extension

```bash
$ polyglot run file.xyz
Error: unsupported file extension: .xyz
```

## Utility Commands

### List supported languages

```bash
$ polyglot list
Supported Languages:
-------------------
Python          [.py]
Go              [.go]
JavaScript      [.js]
Ruby            [.rb]
PHP             [.php]
```

### Check installed toolchains

```bash
$ polyglot check
Checking installed toolchains...
================================
Python          python3         ✓ INSTALLED
Go              go              ✓ INSTALLED
JavaScript      node            ✓ INSTALLED
Ruby            ruby            ✗ NOT FOUND
PHP             php             ✗ NOT FOUND
================================
Summary: 3/5 toolchains installed
```

### Show version

```bash
$ polyglot version
Polyglot v0.1.0
Milestone 1 (MVP)
Universal compiler and interpreter wrapper
```

### Get help

```bash
$ polyglot --help
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

## Real-World Examples

### Quick script testing

```bash
# Test a Python script
polyglot run test.py

# Test a Go program
polyglot run main.go

# Test JavaScript
polyglot run index.js
```

### CI/CD Integration

```yaml
# .github/workflows/test.yml
- name: Test with Polyglot
  run: |
    polyglot run tests/test.py
    polyglot run tests/test.go
    polyglot run tests/test.js
```

### Shell scripting

```bash
#!/bin/bash
# Run all test files

for file in tests/*; do
    echo "Testing $file..."
    polyglot run "$file"
done
```

### Data processing pipeline

```bash
# Generate data with Python
polyglot run -q generate_data.py > data.json

# Process with JavaScript
polyglot run -q process.js < data.json > output.json

# Analyze with Go
polyglot run analyze.go --args output.json
```

## Advanced Usage

### Running with different interpreters

Currently Polyglot uses default interpreters:
- Python: `python3`
- JavaScript: `node`
- Ruby: `ruby`
- PHP: `php`
- Go: `go run`

Future versions will support custom interpreter paths via config files.

### Project-specific configuration

Coming in Milestone 5: `.polyglot.yaml` for project-specific settings.

### Watch mode

Coming in Milestone 6: `polyglot watch` for automatic re-execution on file changes.

## Tips & Tricks

### Create an alias

```bash
# Add to ~/.bashrc or ~/.zshrc
alias pg='polyglot run'

# Now you can use:
pg hello.py
pg main.go
```

### Use with shebang

```python
#!/usr/bin/env polyglot run
print("This file can be executed directly!")
```

```bash
chmod +x script.py
./script.py
```

### Quick one-liners

```bash
# Python
echo 'print("Hello!")' | polyglot run /dev/stdin  # (not yet supported)

# For now, use temp files:
echo 'print("Hello!")' > /tmp/test.py && polyglot run /tmp/test.py
```

## Troubleshooting

### Command not found

```bash
$ polyglot run hello.py
bash: polyglot: command not found
```

Solution: Make sure polyglot is installed and in your PATH:

```bash
which polyglot
echo $PATH
```

### Toolchain not found

```bash
$ polyglot run hello.rb
Detected: Ruby
Error: exec: "ruby": executable file not found in $PATH
```

Solution: Install the required toolchain:

```bash
# Check what's installed
polyglot check

# Install missing toolchain (example for Ubuntu)
sudo apt install ruby
```

### Permission denied

```bash
$ polyglot run script.py
Error: permission denied
```

Solution: Check file permissions:

```bash
ls -l script.py
chmod +r script.py
```

## More Examples

See the `test/fixtures/` directory for more example programs in each supported language.
