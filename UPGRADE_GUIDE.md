# Upgrade Guide

Guide for upgrading Polyglot between versions.

## Upgrading to v1.0.1 (Auto-Update Feature)

### From v1.0.0

The easiest way to upgrade is using the new auto-update feature:

```bash
polyglot update
```

That's it! Polyglot will:
1. Check for the latest version
2. Download the appropriate binary
3. Backup your current version
4. Install the new version
5. Verify the installation

### Manual Upgrade

If you prefer to upgrade manually:

#### Via Package Manager

**Homebrew (macOS/Linux)**:
```bash
brew upgrade polyglot
```

**APT (Debian/Ubuntu)**:
```bash
sudo apt update
sudo apt upgrade polyglot
```

**Chocolatey (Windows)**:
```bash
choco upgrade polyglot
```

#### Via Direct Download

1. Download the latest release:
```bash
# Linux amd64
curl -LO https://github.com/yourusername/polyglot/releases/download/v1.0.1/polyglot-linux-amd64

# macOS amd64 (Intel)
curl -LO https://github.com/yourusername/polyglot/releases/download/v1.0.1/polyglot-darwin-amd64

# macOS arm64 (Apple Silicon)
curl -LO https://github.com/yourusername/polyglot/releases/download/v1.0.1/polyglot-darwin-arm64

# Windows amd64
curl -LO https://github.com/yourusername/polyglot/releases/download/v1.0.1/polyglot-windows-amd64.exe
```

2. Make it executable (Unix only):
```bash
chmod +x polyglot-*
```

3. Replace your current binary:
```bash
# Linux/macOS
sudo mv polyglot-* /usr/local/bin/polyglot

# Windows (run as Administrator)
move polyglot-windows-amd64.exe C:\Program Files\polyglot\polyglot.exe
```

4. Verify the installation:
```bash
polyglot version
```

### Verify Upgrade

After upgrading, verify you're running v1.0.1:

```bash
$ polyglot version
Polyglot v1.0.1
Production Ready - Auto-Update Feature Added!
Universal compiler and interpreter wrapper
```

### New Features in v1.0.1

After upgrading, you can now:

1. **Update automatically**:
```bash
polyglot update
```

2. **Check for updates**:
```bash
polyglot update --check
```

3. **Get automatic update notifications** (every 24 hours)

### Configuration Changes

No configuration changes are required for v1.0.1. All existing configurations remain compatible.

### Breaking Changes

None! v1.0.1 is fully backward compatible with v1.0.0.

## Upgrading from Earlier Versions

### From v0.x to v1.0.1

1. First upgrade to v1.0.0 (if not already):
```bash
# Download v1.0.0
curl -LO https://github.com/yourusername/polyglot/releases/download/v1.0.0/polyglot-linux-amd64
chmod +x polyglot-linux-amd64
sudo mv polyglot-linux-amd64 /usr/local/bin/polyglot
```

2. Then use auto-update to get v1.0.1:
```bash
polyglot update
```

### Configuration Migration

If you have custom configurations from v0.x:

1. **Global config** (`~/.polyglot/config.yaml`) - No changes needed
2. **Project config** (`.polyglot.yaml`) - No changes needed
3. **Environment variables** - No changes needed

All configurations are forward compatible.

## Troubleshooting

### Update Fails

If `polyglot update` fails:

1. **Check internet connection**:
```bash
curl -I https://api.github.com
```

2. **Check permissions**:
```bash
ls -la $(which polyglot)
```

If you don't have write permissions, use sudo:
```bash
sudo polyglot update
```

3. **Try manual update** (see above)

### Version Mismatch

If `polyglot version` shows the wrong version after updating:

1. **Check which binary is being used**:
```bash
which polyglot
```

2. **Check if there are multiple installations**:
```bash
find /usr -name polyglot 2>/dev/null
```

3. **Remove old installations** and reinstall

### Configuration Issues

If you encounter configuration issues after upgrading:

1. **Validate your config**:
```bash
polyglot config show
```

2. **Reset to defaults** (if needed):
```bash
mv ~/.polyglot/config.yaml ~/.polyglot/config.yaml.backup
polyglot config init
```

3. **Restore custom settings** from backup

## Rollback

If you need to rollback to a previous version:

### Automatic Rollback

If an update fails, Polyglot automatically rolls back to the previous version.

### Manual Rollback

To manually rollback to v1.0.0:

```bash
# Download v1.0.0
curl -LO https://github.com/yourusername/polyglot/releases/download/v1.0.0/polyglot-linux-amd64
chmod +x polyglot-linux-amd64
sudo mv polyglot-linux-amd64 /usr/local/bin/polyglot

# Verify
polyglot version
```

## Version History

| Version | Release Date | Key Features |
|---------|--------------|--------------|
| v1.0.1 | 2026-02-22 | Auto-update feature |
| v1.0.0 | 2026-02-22 | Production ready, all 10 milestones |
| v0.5.0 | 2026-02-22 | Developer experience (colors, watch mode) |
| v0.4.0 | 2026-02-22 | Configuration system |
| v0.3.0 | 2026-02-22 | Auto-install system |
| v0.2.0 | 2026-02-22 | 30 languages support |
| v0.1.0 | 2026-02-22 | MVP with 5 languages |

## Getting Help

If you encounter issues during upgrade:

1. Check the [FAQ](FAQ.md)
2. Search [existing issues](https://github.com/yourusername/polyglot/issues)
3. Open a [new issue](https://github.com/yourusername/polyglot/issues/new)
4. Ask in [discussions](https://github.com/yourusername/polyglot/discussions)

## Best Practices

### Before Upgrading

1. **Backup your configuration**:
```bash
cp ~/.polyglot/config.yaml ~/.polyglot/config.yaml.backup
```

2. **Note your current version**:
```bash
polyglot version > /tmp/polyglot-version-before.txt
```

3. **Test in a non-production environment** (if applicable)

### After Upgrading

1. **Verify the version**:
```bash
polyglot version
```

2. **Test basic functionality**:
```bash
polyglot list
polyglot check
polyglot run test/fixtures/hello.py
```

3. **Check your configuration**:
```bash
polyglot config show
```

### Staying Up-to-Date

1. **Enable automatic checks** (enabled by default in v1.0.1)
2. **Subscribe to releases** on GitHub
3. **Check for updates regularly**:
```bash
polyglot update --check
```

## Future Upgrades

Starting with v1.0.1, all future upgrades can be done with:

```bash
polyglot update
```

No more manual downloads or installations needed!

---

**Last Updated**: 2026-02-22  
**Current Version**: v1.0.1
