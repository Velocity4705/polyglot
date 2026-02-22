# Polyglot - Roadmap

Future plans and features for Polyglot.

---

## Current Status

**Version**: 1.0.1  
**Status**: Production Ready ✅

Polyglot currently supports:
- ✅ 30 programming languages
- ✅ 9 CLI commands
- ✅ Auto-install system
- ✅ Configuration management
- ✅ Watch mode
- ✅ Colored output
- ✅ Auto-update feature
- ✅ Multi-platform support

---

## Future Plans

### v1.1.0 - Security & Stability
**Focus**: Improve update security and fix bugs

- SHA256 checksum verification for updates
- GPG signature verification
- Beta/nightly release channels
- Bug fixes from user feedback
- Performance optimizations

### v1.2.0 - Project Support
**Focus**: Handle real-world projects

- Multi-file project support (handle imports)
- Build system integration (Makefile, Cargo.toml, package.json)
- Dependency management (auto-install pip, npm, cargo packages)
- Project detection and configuration

### v1.3.0 - Advanced Features
**Focus**: Security and isolation

- Container-based execution (Docker integration)
- Sandboxed execution mode
- Resource limits (CPU, memory, time)
- Remote execution support

### v2.0.0 - Performance & Scale
**Focus**: Speed and efficiency

- Build caching
- Parallel compilation
- Incremental builds
- Performance benchmarks
- Memory optimizations

---

## Community Contributions Welcome

We welcome contributions in these areas:

- **Additional language support** - Submit PRs for new languages
- **IDE integration guides** - Document how to use Polyglot with VSCode, IntelliJ, etc.
- **CI/CD templates** - Pre-built configs for GitHub Actions, GitLab CI
- **Language-specific optimizations** - Improve handling for specific languages
- **Documentation improvements** - Better examples, tutorials, guides

---

## Out of Scope

These features are intentionally NOT planned:

- ❌ **Language Server Protocol (LSP)** - Each language has its own LSP server
- ❌ **Code editing features** - Polyglot is for execution, not editing
- ❌ **Plugin system** - Just add languages to core or fork the project
- ❌ **Web playground** - Just install locally (takes 30 seconds)
- ❌ **IDE-like features** - Use proper IDEs for development

**Why?** Polyglot focuses on doing ONE thing well: executing code in any language.

---

## Philosophy

**"The best CLI tools do one thing well."**

- `git` → Version control
- `docker` → Containers
- `make` → Build automation
- `polyglot` → Universal code execution

We keep Polyglot focused on its core purpose: making it easy to run code in any programming language.

---

## Feedback & Suggestions

Have ideas for Polyglot? We'd love to hear them!

- Open an issue on GitHub
- Start a discussion
- Submit a pull request

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

---

**Last Updated**: 2026-02-22  
**Current Version**: 1.0.1
