# Release Notes - Polyglot v1.0.1

**Release Date**: 2026-02-22  
**Type**: Feature Release  
**Status**: Production Ready

## 🎉 What's New

### Auto-Update Feature

Polyglot can now update itself! No more manual downloads or reinstallations.

```bash
# Update to the latest version
polyglot update

# Check for updates without installing
polyglot update --check
```

## ✨ Key Features

### 1. Self-Updating Binary
- One-command updates from GitHub releases
- Automatic platform detection (Linux, macOS, Windows)
- Architecture support (amd64, arm64)
- Safe updates with automatic backup and rollback

### 2. Background Update Checks
- Automatic checks every 24 hours
- Non-blocking notifications
- No performance impact
- Configurable check interval

### 3. Package Manager Integration
- Detects Homebrew installations
- Detects APT installations
- Uses native package manager when available
- Falls back to direct binary update

### 4. Safety Features
- Automatic backup before updating
- Rollback on failure
- HTTPS-only downloads
- Verification of downloaded binaries

## 📦 What's Included

### New Commands
- `polyglot update` - Update to latest version
- `polyglot update --check` - Check for updates only

### New Package
- `internal/updater` - Complete update management system

### Enhanced Documentation
- `docs/AUTO_UPDATE.md` - Comprehensive update guide
- Updated quickstart and README
- Examples and troubleshooting

## 🔧 Technical Details

### Update Process
1. Check GitHub API for latest release
2. Compare versions (current vs latest)
3. Download appropriate binary for platform
4. Backup current binary
5. Replace with new version
6. Verify installation
7. Cleanup temporary files

### Supported Platforms
- Linux: amd64, arm64
- macOS: amd64 (Intel), arm64 (Apple Silicon)
- Windows: amd64

### Security
- HTTPS-only downloads from GitHub
- Automatic backup before update
- Rollback on any failure
- No elevated permissions required (unless installed in system directory)

## 📊 Statistics

- **New Files**: 3
- **Modified Files**: 7
- **Lines Added**: ~1,000+
- **Tests Added**: 5
- **Documentation**: 400+ lines

## 🚀 Getting Started

### Installation

If you're installing for the first time:

```bash
# Homebrew (macOS/Linux)
brew install polyglot

# Universal installer
curl -fsSL https://get.polyglot.dev | bash

# From source
git clone https://github.com/yourusername/polyglot.git
cd polyglot
go build -o polyglot ./cmd/polyglot
sudo mv polyglot /usr/local/bin/
```

### Updating

If you already have Polyglot installed:

```bash
# Update to v1.0.1
polyglot update
```

## 📖 Usage Examples

### Check for Updates
```bash
$ polyglot update --check
ℹ Checking for updates...
New version available: v1.0.2 (current: v1.0.1)
Run 'polyglot update' to upgrade
```

### Install Update
```bash
$ polyglot update
ℹ Checking for updates...
New version available: v1.0.2 (current: v1.0.1)

Updating Polyglot
=================

→ Downloading polyglot-linux-amd64...
✓ Downloaded successfully
→ Backing up current version...
→ Installing new version...
✓ Update complete!

ℹ Polyglot has been updated to version 1.0.2
  Run 'polyglot version' to verify
```

### Already Up-to-Date
```bash
$ polyglot update
ℹ Checking for updates...
✓ You're running the latest version (1.0.1)
```

## 🐛 Bug Fixes

- Fixed type conversion issue in language detection
- Fixed unused import in update command
- Improved error handling in CLI commands

## 🔄 Upgrade Path

### From v1.0.0
Simply run:
```bash
polyglot update
```

### From Earlier Versions
1. Update to v1.0.0 first (if not already)
2. Then run `polyglot update` to get v1.0.1

## ⚠️ Breaking Changes

None! This is a backward-compatible feature release.

## 📝 Notes

### Background Update Checks
Polyglot will automatically check for updates every 24 hours. When a new version is available, you'll see a notification:

```
A new version of Polyglot is available: v1.0.2
Run 'polyglot update' to upgrade
```

This check is non-blocking and won't slow down your commands.

### Package Manager Installations
If you installed Polyglot via Homebrew or APT, the update command will use your package manager:

```bash
# Homebrew
polyglot update
# Runs: brew upgrade polyglot

# APT
polyglot update
# Runs: sudo apt update && sudo apt upgrade polyglot
```

### Manual Updates
You can always update manually if needed:

```bash
# Via package manager
brew upgrade polyglot
sudo apt upgrade polyglot

# Via direct download
curl -LO https://github.com/yourusername/polyglot/releases/latest/download/polyglot-linux-amd64
chmod +x polyglot-linux-amd64
sudo mv polyglot-linux-amd64 /usr/local/bin/polyglot
```

## 🔮 What's Next

### Planned for v1.1.0
- SHA256 checksum verification
- GPG signature verification
- Beta/nightly release channels
- Configurable update check interval
- Update history tracking

### Planned for v1.2.0
- Plugin system
- Custom language plugins
- Language-specific configuration
- Performance optimizations

## 🤝 Contributing

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Areas for Contribution
- Additional language support
- Package manager integrations
- Documentation improvements
- Bug fixes and optimizations
- Test coverage improvements

## 📚 Documentation

- [Auto-Update Guide](docs/AUTO_UPDATE.md) - Complete update documentation
- [Quick Start](docs/QUICKSTART.md) - Get started in 5 minutes
- [Configuration](docs/CONFIGURATION.md) - Configuration reference
- [Architecture](docs/ARCHITECTURE.md) - System design
- [FAQ](FAQ.md) - Frequently asked questions

## 🙏 Acknowledgments

Thanks to all contributors and users who provided feedback and suggestions!

## 📞 Support

- **Issues**: https://github.com/yourusername/polyglot/issues
- **Discussions**: https://github.com/yourusername/polyglot/discussions
- **Documentation**: https://github.com/yourusername/polyglot/tree/main/docs

## 📄 License

MIT License - see [LICENSE](LICENSE) for details

---

**Full Changelog**: https://github.com/yourusername/polyglot/compare/v1.0.0...v1.0.1

**Download**: https://github.com/yourusername/polyglot/releases/tag/v1.0.1
