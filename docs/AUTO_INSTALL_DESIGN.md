# Auto-Install Design: System-Wide Toolchain Installation

## Vision

When a user runs `polyglot run hello.py` without Python installed, Polyglot automatically installs Python **system-wide** so it's available for all tools, not just Polyglot.

## User Experience

### First Time Running Python

```bash
$ polyglot run hello.py

⚠️  Python not found in your system

Polyglot can install Python 3.12 for you:
  • Will be installed system-wide (available to all programs)
  • Uses your system's package manager (dnf/apt/brew)
  • One-time installation (~5 minutes)
  
Install Python now? [Y/n] y

📦 Installing Python 3.12...
   Running: sudo dnf install python3
   [sudo] password for user: ****
   
   Installing packages...
   ████████████████████████ 100%
   
✅ Python 3.12 installed successfully!
💡 Python is now available system-wide

🚀 Running hello.py...
Hello from Python!

---
Next time you run Python code, it will execute immediately!
```

### Second Time (Already Installed)

```bash
$ polyglot run hello.py
Detected: Python
Executing: python3 hello.py
Hello from Python!
```

### Multiple Languages

```bash
$ polyglot run hello.c

⚠️  C compiler (gcc) not found

Install gcc? [Y/n] y
📦 Installing gcc...
✅ gcc installed!
🚀 Running hello.c...

$ polyglot run hello.rs

⚠️  Rust compiler not found

Install Rust? [Y/n] y
📦 Installing Rust...
✅ Rust installed!
🚀 Running hello.rs...

# Now both are available system-wide!
$ gcc --version
gcc (GCC) 13.2.0

$ rustc --version
rustc 1.75.0
```

---

## Implementation Strategy

### 1. Detection Phase

```go
func checkToolchain(language string) (bool, string) {
    // Check if toolchain exists
    path, err := exec.LookPath(getCommand(language))
    if err == nil {
        return true, path  // Found!
    }
    return false, ""  // Not found
}
```

### 2. Installation Phase

```go
func installToolchain(language string) error {
    // Detect package manager
    pkgManager := detectPackageManager()
    
    // Get package name for this language
    pkgName := getPackageName(language, pkgManager)
    
    // Prompt user
    if !promptUser(language, pkgName) {
        return errors.New("installation cancelled")
    }
    
    // Install using system package manager
    return installPackage(pkgManager, pkgName)
}
```

### 3. Package Manager Detection

```go
func detectPackageManager() string {
    managers := []struct {
        name    string
        command string
    }{
        {"dnf", "dnf"},
        {"apt", "apt-get"},
        {"brew", "brew"},
        {"pacman", "pacman"},
        {"zypper", "zypper"},
    }
    
    for _, mgr := range managers {
        if _, err := exec.LookPath(mgr.command); err == nil {
            return mgr.name
        }
    }
    return ""
}
```

### 4. Package Name Mapping

```go
var packageNames = map[string]map[string]string{
    "python": {
        "dnf":    "python3",
        "apt":    "python3",
        "brew":   "python3",
        "pacman": "python",
    },
    "gcc": {
        "dnf":    "gcc",
        "apt":    "build-essential",
        "brew":   "gcc",
        "pacman": "gcc",
    },
    "rust": {
        "dnf":    "rust",
        "apt":    "rustc",
        "brew":   "rust",
        "pacman": "rust",
    },
    // ... more languages
}
```

---

## Installation Methods

### Method 1: System Package Manager (Preferred)

**Pros:**
- ✅ Integrates with system
- ✅ Automatic updates via system
- ✅ Trusted sources
- ✅ Handles dependencies

**Cons:**
- ❌ Requires sudo
- ❌ May not have latest versions

**Implementation:**
```bash
# Fedora/RHEL
sudo dnf install python3

# Ubuntu/Debian
sudo apt install python3

# macOS
brew install python3

# Arch
sudo pacman -S python
```

### Method 2: Official Install Scripts (Fallback)

For languages not in package managers or when user wants latest version:

**Rust:**
```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

**Node.js (via nvm):**
```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
nvm install node
```

**Zig:**
```bash
# Download from ziglang.org
wget https://ziglang.org/download/0.11.0/zig-linux-x86_64-0.11.0.tar.xz
tar -xf zig-linux-x86_64-0.11.0.tar.xz
sudo mv zig-linux-x86_64-0.11.0 /opt/zig
sudo ln -s /opt/zig/zig /usr/local/bin/zig
```

### Method 3: Polyglot-Managed (Last Resort)

If system installation fails, install to `~/.polyglot/toolchains/` but add to PATH:

```bash
# Install to Polyglot directory
~/.polyglot/toolchains/python-3.12/

# Add to shell profile
echo 'export PATH="$HOME/.polyglot/toolchains/python-3.12/bin:$PATH"' >> ~/.bashrc
```

---

## User Prompts

### Interactive Mode (Default)

```bash
⚠️  Python not found in your system

Polyglot can install Python 3.12 for you:
  
  Installation method: System package manager (dnf)
  Package: python3
  Size: ~50MB
  Time: ~2 minutes
  Requires: sudo password
  
  After installation, Python will be available:
    • For Polyglot (polyglot run *.py)
    • For direct use (python3 script.py)
    • For other tools (pip, virtualenv, etc.)

Install Python now? [Y/n] 
```

### Auto-Install Mode

```bash
# Enable auto-install in config
polyglot config set auto-install true

# Now it installs without prompting
$ polyglot run hello.py
📦 Installing Python... ✓
🚀 Running hello.py...
```

### Manual Mode

```bash
# Disable auto-install
polyglot config set auto-install false

# Now it just shows instructions
$ polyglot run hello.py
⚠️  Python not found

To install Python:
  Fedora/RHEL: sudo dnf install python3
  Ubuntu/Debian: sudo apt install python3
  macOS: brew install python3

Or enable auto-install:
  polyglot config set auto-install true
```

---

## Configuration

### Config File: `~/.polyglot/config.yaml`

```yaml
# Auto-installation settings
auto_install:
  enabled: true              # Auto-install missing toolchains
  method: system             # system, script, or managed
  prompt: true               # Prompt before installing
  
# Package manager preferences
package_manager:
  prefer: auto               # auto, dnf, apt, brew, pacman
  
# Per-language settings
languages:
  python:
    auto_install: true
    method: system           # Use system package manager
    package: python3
    
  rust:
    auto_install: true
    method: script           # Use rustup script
    
  zig:
    auto_install: false      # Don't auto-install
```

---

## Commands

### New Commands

```bash
# Install a language
polyglot install python
polyglot install rust
polyglot install gcc

# Install multiple
polyglot install python rust node

# Check installation status
polyglot status

# Configure auto-install
polyglot config set auto-install true
polyglot config set auto-install false

# Show what would be installed
polyglot install --dry-run python
```

### Enhanced Check Command

```bash
$ polyglot check

Checking installed toolchains...
================================
Python          python3         ✓ INSTALLED (system)
Go              go              ✓ INSTALLED (system)
JavaScript      node            ✗ NOT FOUND
                                  → Run: polyglot install node
Rust            rustc           ✗ NOT FOUND
                                  → Run: polyglot install rust
================================
Summary: 2/4 checked toolchains installed

💡 Enable auto-install: polyglot config set auto-install true
```

---

## Security Considerations

### 1. Sudo Requirement

```bash
⚠️  Installing Python requires administrator privileges

This will run: sudo dnf install python3

Continue? [y/N] y
[sudo] password for user: ****
```

### 2. Package Verification

- Use official package managers (trusted sources)
- Verify checksums for downloaded binaries
- Show exact commands before running

### 3. User Control

- Always prompt before installing (unless auto-install enabled)
- Show what will be installed
- Allow cancellation at any time
- Provide manual installation instructions

---

## Error Handling

### Package Manager Not Found

```bash
⚠️  No package manager found

Polyglot supports: dnf, apt, brew, pacman, zypper

Please install Python manually:
  • Download from: https://python.org
  • Or use your system's package manager

After installation, run: polyglot check
```

### Installation Failed

```bash
❌ Failed to install Python

Error: Package python3 not found in repositories

Try:
  1. Update package lists: sudo dnf update
  2. Install manually: https://python.org
  3. Check your internet connection

Need help? https://github.com/yourusername/polyglot/issues
```

### Insufficient Permissions

```bash
❌ Installation requires administrator privileges

Please run one of:
  • polyglot install python (will prompt for sudo)
  • sudo polyglot install python
  • Install manually: sudo dnf install python3
```

---

## Benefits of This Approach

### For Users
- ✅ **Zero configuration** - just run your code
- ✅ **System-wide availability** - works with all tools
- ✅ **One-time setup** - install once, use everywhere
- ✅ **Natural workflow** - learn languages as you use them
- ✅ **No vendor lock-in** - standard system installation

### For System
- ✅ **Proper integration** - uses system package manager
- ✅ **Automatic updates** - via system updates
- ✅ **Dependency management** - handled by package manager
- ✅ **Clean uninstall** - standard system removal

### For Polyglot
- ✅ **Truly universal** - works out of the box
- ✅ **Better UX** - seamless experience
- ✅ **Competitive advantage** - unique feature
- ✅ **Viral growth** - users recommend it

---

## Implementation Phases

### Phase 1: Detection & Prompting (Milestone 5)
- [ ] Detect missing toolchains
- [ ] Prompt user to install
- [ ] Show installation instructions
- [ ] `polyglot install` command

### Phase 2: System Package Manager (Milestone 5)
- [ ] Detect package manager (dnf, apt, brew, etc.)
- [ ] Map languages to package names
- [ ] Execute installation commands
- [ ] Handle sudo prompts

### Phase 3: Alternative Methods (Milestone 6)
- [ ] Official install scripts (rustup, nvm, etc.)
- [ ] Direct binary downloads
- [ ] Fallback to managed installation

### Phase 4: Configuration & Polish (Milestone 6)
- [ ] Config file support
- [ ] Auto-install mode
- [ ] Per-language settings
- [ ] Better error messages

---

## Example Workflows

### Beginner Learning Python

```bash
# Day 1: First Python program
$ polyglot run hello.py
📦 Installing Python... ✓
Hello from Python!

# Day 2: Python is already there
$ polyglot run calculator.py
Executing: python3 calculator.py
2 + 2 = 4

# Day 3: Use Python directly
$ python3 my_script.py
Works!

# Day 4: Install packages
$ pip install requests
Works! (pip came with Python)
```

### Developer Using Multiple Languages

```bash
# Install as needed
$ polyglot run api.py
📦 Installing Python... ✓

$ polyglot run server.js
📦 Installing Node.js... ✓

$ polyglot run main.go
📦 Installing Go... ✓

# All available system-wide now
$ python3 --version
Python 3.12.0

$ node --version
v20.0.0

$ go version
go version go1.21.0
```

### CI/CD Pipeline

```bash
# .github/workflows/test.yml
- name: Setup Polyglot
  run: |
    curl -sSL https://polyglot.sh/install.sh | sh
    polyglot config set auto-install true

- name: Run tests
  run: |
    polyglot run tests/test.py    # Auto-installs Python
    polyglot run tests/test.go    # Auto-installs Go
    polyglot run tests/test.js    # Auto-installs Node
```

---

## Comparison with Other Tools

### vs Manual Installation
- ✅ Automatic (no manual steps)
- ✅ Consistent across systems
- ✅ Integrated with Polyglot

### vs Docker
- ✅ Faster (no container overhead)
- ✅ System-wide availability
- ✅ Simpler setup

### vs asdf/mise
- ✅ Automatic detection
- ✅ No configuration needed
- ✅ System integration

### vs Nix
- ✅ Simpler to understand
- ✅ Uses standard package managers
- ✅ Easier to debug

---

## Conclusion

This approach gives us the **best of all worlds**:

1. **Zero-config experience** - just run your code
2. **System-wide installation** - works with all tools
3. **User control** - can disable auto-install
4. **Standard integration** - uses system package managers
5. **Natural learning** - install languages as you need them

**Next Steps:**
1. Implement in Milestone 5
2. Test across different systems
3. Gather user feedback
4. Iterate and improve

---

**This is the future of Polyglot!** 🚀
