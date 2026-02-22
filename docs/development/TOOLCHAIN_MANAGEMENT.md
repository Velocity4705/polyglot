# Toolchain Management (Future Feature)

## Current Behavior

Polyglot currently requires users to install language toolchains manually:

```bash
# Install toolchains
sudo dnf install python3 nodejs gcc rustc

# Check what's installed
polyglot check

# Use installed toolchains
polyglot run hello.py
```

---

## Planned Feature: Automatic Toolchain Management

### Vision

Make Polyglot truly universal by automatically managing toolchains.

### Proposed Commands

```bash
# Install a toolchain
polyglot install python
polyglot install rust
polyglot install node

# List available toolchains
polyglot toolchains available

# List installed toolchains
polyglot toolchains list

# Update a toolchain
polyglot update python

# Remove a toolchain
polyglot uninstall python

# Use specific version
polyglot run --python-version 3.11 hello.py
```

### Behavior

**1. Auto-detect system toolchains (current)**
```bash
$ polyglot run hello.py
Detected: Python
Using: /usr/bin/python3 (system)
Executing: python3 hello.py
Hello from Python!
```

**2. Auto-download if not found (new)**
```bash
$ polyglot run hello.py
⚠️  Python not found in system
📦 Download Python 3.12 to ~/.polyglot/toolchains? [Y/n] y
⬇️  Downloading Python 3.12... ✓
✅ Python 3.12 installed
🚀 Running hello.py...
Hello from Python!
```

**3. Prefer system, fallback to managed (new)**
```bash
$ polyglot run hello.py
Detected: Python
Using: ~/.polyglot/toolchains/python-3.12/bin/python3
Executing: python3 hello.py
Hello from Python!
```

---

## Implementation Strategy

### Phase 1: Detection Enhancement
- Detect system toolchains (current)
- Detect managed toolchains in `~/.polyglot/toolchains/`
- Prefer system over managed

### Phase 2: Download System
- Download pre-built binaries
- Extract to `~/.polyglot/toolchains/<lang>-<version>/`
- Update PATH for polyglot execution

### Phase 3: Version Management
- Support multiple versions
- Switch between versions
- Update toolchains

### Phase 4: Advanced Features
- Automatic updates
- Toolchain caching
- Offline mode
- Custom toolchain paths

---

## Directory Structure

```
~/.polyglot/
├── toolchains/
│   ├── python-3.12/
│   │   ├── bin/python3
│   │   └── lib/...
│   ├── node-20.0/
│   │   ├── bin/node
│   │   └── lib/...
│   ├── rust-1.75/
│   │   ├── bin/rustc
│   │   └── lib/...
│   └── ...
├── cache/
│   └── downloads/
└── config.yaml
```

---

## Configuration

```yaml
# ~/.polyglot/config.yaml

toolchains:
  python:
    prefer: system  # system, managed, or specific path
    version: "3.12"
    path: "/usr/bin/python3"  # optional override
  
  rust:
    prefer: managed
    version: "1.75"
  
  node:
    prefer: system

auto_download: true  # Prompt to download if not found
auto_update: false   # Auto-update toolchains
```

---

## Download Sources

### Option 1: Official Binaries
- Python: python.org
- Node.js: nodejs.org
- Rust: rust-lang.org
- Go: golang.org

### Option 2: Package Managers
- Use system package manager (apt, dnf, brew)
- Requires sudo permissions

### Option 3: Pre-built Bundles
- Host pre-built toolchains
- Fast downloads
- Verified binaries

---

## User Experience Examples

### First-time User
```bash
$ polyglot run hello.py
⚠️  Python not found
📦 Polyglot can download Python 3.12 for you
   This will be installed to ~/.polyglot/toolchains/
   Download now? [Y/n] y

⬇️  Downloading Python 3.12 (50MB)... ████████████ 100%
📦 Extracting...
✅ Python 3.12 installed successfully!

🚀 Running hello.py...
Hello from Python!

💡 Tip: Run 'polyglot toolchains' to manage installed toolchains
```

### Experienced User
```bash
# Install multiple toolchains
$ polyglot install python rust node go

# Check what's installed
$ polyglot toolchains list
System Toolchains:
  ✓ Python 3.11 (/usr/bin/python3)
  ✓ GCC 13.2 (/usr/bin/gcc)

Managed Toolchains:
  ✓ Python 3.12 (~/.polyglot/toolchains/python-3.12)
  ✓ Rust 1.75 (~/.polyglot/toolchains/rust-1.75)
  ✓ Node 20.0 (~/.polyglot/toolchains/node-20.0)

# Use specific version
$ polyglot run --python-version 3.12 hello.py
```

---

## Benefits

### For Users
- ✅ Works out of the box
- ✅ No manual installation needed
- ✅ Version management
- ✅ Consistent environment

### For Developers
- ✅ Easy testing across versions
- ✅ Reproducible builds
- ✅ No system pollution

### For Polyglot
- ✅ Truly universal
- ✅ Better user experience
- ✅ Competitive advantage

---

## Challenges

### Technical
- Binary compatibility across systems
- Disk space management
- Download reliability
- Security (verify binaries)

### Legal
- Redistribution licenses
- Trademark usage
- Attribution requirements

### Practical
- Large download sizes
- Network requirements
- Storage requirements

---

## Alternatives

### 1. Docker-based
```bash
polyglot run --docker hello.py
# Runs in Docker container with Python
```

### 2. WebAssembly
```bash
# Bundle WASM interpreters
polyglot run --wasm hello.py
```

### 3. Cloud-based
```bash
# Execute remotely
polyglot run --cloud hello.py
```

---

## Roadmap

### Milestone 5: Basic Toolchain Management
- [ ] Detect managed toolchains
- [ ] Download command
- [ ] List command
- [ ] Prefer system over managed

### Milestone 6: Advanced Management
- [ ] Auto-download on first use
- [ ] Version switching
- [ ] Update command
- [ ] Configuration file

### Milestone 7: Polish
- [ ] Offline mode
- [ ] Caching
- [ ] Verification
- [ ] Documentation

---

## Current Recommendation

**For now**: Users install toolchains manually (current behavior)

**Future**: Implement hybrid approach in Milestone 5-6

**Best of both worlds**:
- Respects system installations
- Auto-downloads when needed
- Gives users control

---

## Feedback Welcome!

What do you think? Should Polyglot:
1. Stay as a wrapper (current)
2. Auto-download toolchains
3. Bundle common toolchains
4. Hybrid approach

Let us know in the issues!
