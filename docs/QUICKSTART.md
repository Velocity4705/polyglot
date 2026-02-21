# Quick Start Guide

Get up and running with Polyglot in 5 minutes!

## Installation

### Option 1: Quick Install (Recommended)

```bash
git clone https://github.com/yourusername/polyglot.git
cd polyglot
./scripts/install.sh
```

### Option 2: Manual Build

```bash
git clone https://github.com/yourusername/polyglot.git
cd polyglot
make build
sudo cp bin/polyglot /usr/local/bin/
```

### Option 3: Go Install

```bash
go install github.com/yourusername/polyglot/cmd/polyglot@latest
```

## Verify Installation

```bash
polyglot version
```

You should see:
```
Polyglot v0.1.0
Milestone 1 (MVP)
Universal compiler and interpreter wrapper
```

## Check Your Toolchains

See which language toolchains you have installed:

```bash
polyglot check
```

Example output:
```
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

## Your First Program

### 1. Create a Python script

```bash
echo 'print("Hello from Polyglot!")' > hello.py
```

### 2. Run it with Polyglot

```bash
polyglot run hello.py
```

Output:
```
Detected: Python
Executing: python3 hello.py
Hello from Polyglot!
```

## Try Different Languages

### JavaScript

```bash
echo 'console.log("Hello from JavaScript!");' > hello.js
polyglot run hello.js
```

### Go

```bash
cat > hello.go << 'EOF'
package main
import "fmt"
func main() {
    fmt.Println("Hello from Go!")
}
EOF

polyglot run hello.go
```

### Ruby (if installed)

```bash
echo 'puts "Hello from Ruby!"' > hello.rb
polyglot run hello.rb
```

## Passing Arguments

Create a script that uses arguments:

```python
# args.py
import sys
print(f"You passed {len(sys.argv)-1} arguments:")
for arg in sys.argv[1:]:
    print(f"  - {arg}")
```

Run with arguments:

```bash
polyglot run args.py --args hello,world,test
```

## Quiet Mode

Suppress Polyglot's messages:

```bash
polyglot run -q hello.py
```

Output:
```
Hello from Polyglot!
```

## Verbose Mode

See detailed execution info:

```bash
polyglot run -v hello.py
```

## List Supported Languages

```bash
polyglot list
```

Output:
```
Supported Languages:
-------------------
Python          [.py]
Go              [.go]
JavaScript      [.js]
Ruby            [.rb]
PHP             [.php]
```

## Next Steps

- Read the [full README](../README.md) for more features
- Check the [language support matrix](LANGUAGES.md)
- Learn how to [contribute](../CONTRIBUTING.md)
- See the [roadmap](../ROADMAP.md) for upcoming features

## Common Issues

### "command not found: polyglot"

Make sure `/usr/local/bin` is in your PATH:

```bash
echo $PATH
```

If not, add to your shell config (~/.bashrc or ~/.zshrc):

```bash
export PATH="/usr/local/bin:$PATH"
```

### "unsupported file extension"

Check if the language is supported:

```bash
polyglot list
```

### "execution failed"

Check if the toolchain is installed:

```bash
polyglot check
```

## Getting Help

```bash
polyglot --help
polyglot run --help
```

## Uninstall

```bash
sudo rm /usr/local/bin/polyglot
```
