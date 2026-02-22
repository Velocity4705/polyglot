# Polyglot - Frequently Asked Questions

## General Questions

### What is Polyglot?

Polyglot is a universal compiler and interpreter wrapper that automatically detects programming languages and executes them with the appropriate toolchain using a single, consistent command.

Instead of remembering different commands for different languages, just use:
```bash
polyglot run <any-file>
```

### How many languages does it support?

Currently **16 languages**:
- 7 interpreted: Python, JavaScript, Ruby, PHP, Perl, Lua, Shell
- 9 compiled: Go, Java, C, C++, Rust, Zig, Nim, Crystal, D

### Is it free?

Yes! Polyglot is open source under the MIT License.

---

## Installation & Setup

### Do I need to install all 16 language toolchains?

**No!** You only need to install the languages you want to use.

Polyglot is a **wrapper tool** - it uses your system's installed compilers and interpreters. It doesn't bundle them.

```bash
# Check what you have installed
polyglot check

# Install only what you need
sudo dnf install python3 gcc nodejs  # Example
```

### Why doesn't Polyglot bundle the compilers?

**Current Design (Wrapper Model):**
- ✅ Small binary (~9MB)
- ✅ Uses system-installed, up-to-date compilers
- ✅ Users control versions
- ✅ No licensing issues
- ✅ Respects system package management

**Yes!** Polyglot can auto-install missing toolchains using your system package manager.

When you try to run code in a language that's not installed, Polyglot will offer to install it for you.

See [AUTO_UPDATE.md](docs/AUTO_UPDATE.md) for details.

### Can Polyglot auto-install missing languages?

**Yes!** Use the `polyglot install` command or let Polyglot prompt you when a language is missing.

**Current behavior:**
```bash
$ polyglot run hello.py
Error: python3 not found
```

**Future behavior (planned):**
```bash
$ polyglot run hello.py
⚠️  Python not found. Download? [Y/n] y
📦 Downloading Python 3.12... ✓
🚀 Running hello.py...
```

### How do I install language toolchains?

**Fedora/RHEL:**
```bash
sudo dnf install python3 nodejs gcc g++ rustc ruby php perl lua
```

**Ubuntu/Debian:**
```bash
sudo apt install python3 nodejs build-essential rustc ruby php perl lua
```

**macOS:**
```bash
brew install python3 node gcc rust ruby php perl lua
```

Then verify:
```bash
polyglot check
```

---

## Usage Questions

### How do I run a program?

```bash
polyglot run <file>
```

Examples:
```bash
polyglot run hello.py
polyglot run main.go
polyglot run app.js
polyglot run hello.c
```

### How do I pass arguments to my program?

```bash
polyglot run script.py --args arg1,arg2,arg3
```

### How do I suppress Polyglot's messages?

Use quiet mode:
```bash
polyglot run -q script.py
```

Perfect for piping:
```bash
polyglot run -q data.py | jq .
```

### How do I see detailed execution info?

Use verbose mode:
```bash
polyglot run -v hello.py
```

### What if my language isn't supported?

1. Check if it's planned: [ROADMAP.md](ROADMAP.md)
2. Request it: Open an issue
3. Add it yourself: See [CONTRIBUTING.md](CONTRIBUTING.md) - takes ~15 minutes!

---

## Compilation Questions

### Does Polyglot compile my code?

**For compiled languages (C, C++, Rust, Java):** Yes!

Polyglot automatically:
1. Compiles your code
2. Runs the compiled binary
3. Cleans up the binary after execution

```bash
$ polyglot run hello.c
Detected: C
Compiling: C hello.c
Compilation successful
Executing: C hello.c
Hello from C!
# Binary automatically cleaned up
```

### Where do compiled binaries go?

They're created in the same directory as your source file and automatically deleted after execution.

If you want to keep the binary, use:
```bash
polyglot compile hello.c
# Creates 'hello' binary without running it
```

### Can I see the compilation command?

Use verbose mode:
```bash
polyglot run -v hello.c
```

---

## Technical Questions

### How does language detection work?

Polyglot looks at the file extension:
- `.py` → Python
- `.js` → JavaScript
- `.c` → C
- `.rs` → Rust
- etc.

See the full list:
```bash
polyglot list
```

### What if I have multiple versions of a language?

Polyglot uses whatever is in your PATH:
```bash
which python3  # Shows which Python Polyglot will use
```

### Does Polyglot work on Windows?

**Yes!** Polyglot supports Windows, macOS, and Linux.

### How big is the Polyglot binary?

About 9MB. It's a single, self-contained binary with no dependencies (except the language toolchains you want to use).

### Is Polyglot fast?

Yes! Polyglot adds minimal overhead:
- Detection: < 1ms
- Execution: Direct process spawn (no overhead)
- Compilation: Same speed as running the compiler directly

---

## Comparison Questions

### How is this different from Docker?

**Polyglot:**
- ✅ Lightweight wrapper
- ✅ Uses system toolchains
- ✅ Fast startup
- ✅ No containerization overhead

**Docker:**
- ✅ Complete isolation
- ✅ Reproducible environments
- ❌ Slower startup
- ❌ More complex setup

### How is this different from asdf/mise?

**Polyglot:**
- ✅ Automatic language detection
- ✅ Single command for all languages
- ✅ No configuration needed
- ❌ No version management (yet)

**asdf/mise:**
- ✅ Version management
- ✅ Per-project versions
- ❌ Requires configuration
- ❌ Separate commands per language

**Future:** We're planning to combine the best of both!

### How is this different from Nix?

**Polyglot:**
- ✅ Simple wrapper
- ✅ Easy to understand
- ✅ Works with system packages
- ❌ No reproducibility guarantees

**Nix:**
- ✅ Reproducible builds
- ✅ Declarative configuration
- ❌ Steep learning curve
- ❌ Complex setup

---

## Contributing Questions

### How do I add a new language?

It takes about 15 minutes! See [GETTING_STARTED.md](GETTING_STARTED.md) for a step-by-step tutorial.

Basic steps:
1. Add language type to `pkg/types/types.go`
2. Add detection to `internal/detector/detector.go`
3. Create handler in `internal/language/<lang>.go`
4. Register in `internal/language/registry.go`
5. Add test fixture
6. Update docs

### Can I contribute?

Yes! We welcome contributions. See [CONTRIBUTING.md](CONTRIBUTING.md).

Good first issues:
- Add a new language
- Improve error messages
- Add more examples
- Improve documentation

### How is the project organized?

See [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) for a complete overview.

---

## Troubleshooting

### "command not found: polyglot"

Polyglot isn't in your PATH. Either:
```bash
# Use full path
./bin/polyglot run hello.py

# Or install it
make install
```

### "unsupported file extension"

The language isn't supported yet. Check:
```bash
polyglot list
```

### "executable file not found in $PATH"

The language toolchain isn't installed. Check:
```bash
polyglot check
```

Then install the missing toolchain.

### Compilation fails

Check if the compiler is installed:
```bash
which gcc    # For C
which rustc  # For Rust
which javac  # For Java
```

Use verbose mode to see the exact error:
```bash
polyglot run -v hello.c
```

---

## Future Plans

### What's coming next?

See [ROADMAP.md](ROADMAP.md) for future plans.

**Upcoming features:**
- Multi-file project support
- Build system integration
- Dependency management
- Container-based execution
- Performance optimizations

### Will Polyglot support my favorite language?

Probably! Check the [ROADMAP.md](ROADMAP.md) or open an issue to request it.

---

## Getting Help

### Where can I get help?

- 📖 Read the [Quick Start Guide](docs/QUICKSTART.md)
- 💬 Open an [Issue](https://github.com/Velocity4705/polyglot/issues)
- 📚 Check [Examples](docs/EXAMPLES.md)
- 🎯 See the [Demo](DEMO.md)

### How do I report a bug?

Open an issue with:
- Polyglot version (`polyglot version`)
- Operating system
- Language/toolchain versions
- Steps to reproduce
- Expected vs actual behavior

### How do I request a feature?

Open an issue describing:
- What you want to do
- Why it's useful
- How you envision it working

---

## Philosophy

### Why build Polyglot?

**Problem:** Developers working with multiple languages face friction:
- Different commands for different languages
- Need to remember language-specific flags
- Inconsistent interfaces

**Solution:** One command to run them all!
```bash
polyglot run <any-file>
```

### What's the vision?

Make it effortless to run programs in any programming language with a single, consistent command.

**Current:** Wrapper for system toolchains  
**Future:** Smart toolchain manager with auto-download  
**Ultimate:** Universal development environment

---

## More Questions?

Open an issue or check the documentation:
- [README.md](README.md)
- [DOCUMENTATION_INDEX.md](DOCUMENTATION_INDEX.md)
- [All Docs](docs/)
