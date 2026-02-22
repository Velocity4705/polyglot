# Auto-Update Feature

Polyglot includes a built-in auto-update mechanism that keeps your installation up-to-date with the latest features, bug fixes, and language support.

## Overview

The auto-update feature provides:

- **Manual Updates**: Update on-demand with `polyglot update`
- **Background Checks**: Automatic update notifications every 24 hours
- **Package Manager Support**: Works with Homebrew, APT, and other package managers
- **Safe Updates**: Automatic backup and rollback on failure
- **Cross-Platform**: Works on Linux, macOS, and Windows

## Usage

### Check and Install Updates

```bash
polyglot update
```

This will:
1. Check GitHub for the latest release
2. Download the appropriate binary for your platform
3. Backup your current version
4. Install the new version
5. Verify the installation

### Check for Updates Only

```bash
polyglot update --check
```

This checks if an update is available without installing it.

## How It Works

### Background Update Checks

Polyglot automatically checks for updates in the background every 24 hours. When a new version is available, you'll see a notification:

```
A new version of Polyglot is available: v1.0.2
Run 'polyglot update' to upgrade
```

This check is non-blocking and won't slow down your commands.

### Update Process

1. **Version Check**: Queries GitHub API for latest release
2. **Platform Detection**: Identifies your OS and architecture
3. **Download**: Fetches the appropriate binary
4. **Backup**: Creates a backup of your current version
5. **Installation**: Replaces the binary with the new version
6. **Verification**: Ensures the update was successful
7. **Cleanup**: Removes temporary files and backups

### Package Manager Integration

If you installed Polyglot via a package manager, the update command will use that:

**Homebrew (macOS/Linux)**:
```bash
polyglot update
# Runs: brew upgrade polyglot
```

**APT (Debian/Ubuntu)**:
```bash
polyglot update
# Runs: sudo apt update && sudo apt upgrade polyglot
```

**Direct Binary**:
If not installed via a package manager, Polyglot will download and install the binary directly from GitHub.

## Supported Platforms

The auto-update feature supports:

- **Linux**: amd64, arm64
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Windows**: amd64

## Update Frequency

- **Background Checks**: Every 24 hours
- **Check File**: `~/.polyglot/last_update_check`

You can manually trigger a check at any time with `polyglot update --check`.

## Safety Features

### Automatic Backup

Before updating, Polyglot creates a backup of your current binary:

```
/usr/local/bin/polyglot.backup
```

If the update fails, the backup is automatically restored.

### Rollback on Failure

If the update process encounters an error:
1. The new binary is removed
2. The backup is restored
3. An error message is displayed
4. Your original version remains intact

### Verification

After updating, verify the new version:

```bash
polyglot version
```

## Troubleshooting

### Update Check Fails

If the update check fails:

```
Failed to check for updates: connection timeout
```

This usually means:
- No internet connection
- GitHub API is unavailable
- Firewall blocking the request

The command will fail gracefully and your current version remains unchanged.

### Permission Denied

If you see:

```
Update failed: permission denied
```

This means Polyglot doesn't have permission to replace the binary. Try:

```bash
sudo polyglot update
```

Or reinstall with proper permissions.

### Binary Not Found

If the update can't find a binary for your platform:

```
No binary found for linux/arm
```

This means your platform isn't supported yet. Please open an issue on GitHub.

## Manual Update

If the auto-update fails, you can always update manually:

### Via Package Manager

**Homebrew**:
```bash
brew upgrade polyglot
```

**APT**:
```bash
sudo apt update
sudo apt upgrade polyglot
```

### Via Direct Download

```bash
# Download latest release
curl -LO https://github.com/Velocity4705/polyglot/releases/latest/download/polyglot-linux-amd64

# Make executable
chmod +x polyglot-linux-amd64

# Replace current binary
sudo mv polyglot-linux-amd64 /usr/local/bin/polyglot
```

### Via Go Install

```bash
go install github.com/Velocity4705/polyglot/cmd/polyglot@latest
```

## Disabling Background Checks

If you want to disable background update checks, you can:

1. **Delete the check file**:
```bash
rm ~/.polyglot/last_update_check
```

2. **Set a very old timestamp**:
```bash
touch -t 197001010000 ~/.polyglot/last_update_check
```

Note: This will cause Polyglot to check on the next run, but you can ignore the notification.

## Release Channels

Currently, Polyglot only supports the stable release channel. Future versions may include:

- **Stable**: Production-ready releases (current)
- **Beta**: Pre-release testing versions
- **Nightly**: Daily builds with latest features

## Security

### HTTPS Only

All update checks and downloads use HTTPS to ensure:
- Authenticity of the source
- Integrity of the downloaded binary
- Protection against man-in-the-middle attacks

### GitHub Releases

Updates are only downloaded from official GitHub releases, ensuring:
- Verified source code
- Signed releases (future)
- Transparent change history

### Checksum Verification

Future versions will include SHA256 checksum verification for downloaded binaries.

## Examples

### Regular Update Workflow

```bash
# Check for updates
$ polyglot update --check
Checking for updates...
New version available: v1.0.2 (current: v1.0.1)
Run 'polyglot update' to upgrade

# Install update
$ polyglot update
Checking for updates...
New version available: v1.0.2 (current: v1.0.1)

Updating Polyglot
=================

Downloading polyglot-linux-amd64...
✓ Downloaded successfully
Backing up current version...
Installing new version...
✓ Update complete!

Polyglot has been updated to version 1.0.2
Run 'polyglot version' to verify

# Verify
$ polyglot version
Polyglot v1.0.2
Production Ready - Bug Fixes and Improvements!
Universal compiler and interpreter wrapper
```

### Already Up-to-Date

```bash
$ polyglot update
Checking for updates...
✓ You're running the latest version (1.0.2)
```

### Package Manager Update

```bash
$ polyglot update
Detected Homebrew installation
Running: brew upgrade polyglot
==> Upgrading polyglot
==> Downloading https://homebrew.bintray.com/bottles/polyglot-1.0.2.tar.gz
==> Pouring polyglot-1.0.2.tar.gz
🍺  /usr/local/Cellar/polyglot/1.0.2: 5 files, 8.2MB
```

## FAQ

### How often should I update?

Update whenever you see a notification or when you need new features. Polyglot follows semantic versioning:

- **Major** (2.0.0): Breaking changes
- **Minor** (1.1.0): New features, backward compatible
- **Patch** (1.0.1): Bug fixes

### Will updates break my scripts?

Patch and minor updates are backward compatible. Major updates may include breaking changes, which will be clearly documented in the release notes.

### Can I downgrade?

Yes, you can install any previous version manually:

```bash
# Download specific version
curl -LO https://github.com/Velocity4705/polyglot/releases/download/v1.0.0/polyglot-linux-amd64

# Install
chmod +x polyglot-linux-amd64
sudo mv polyglot-linux-amd64 /usr/local/bin/polyglot
```

### What if I'm offline?

The update check will fail gracefully and your current version will continue to work. Updates require an internet connection.

### Does it work in CI/CD?

Yes, but you may want to pin specific versions in CI/CD for reproducibility:

```bash
# Pin to specific version
curl -LO https://github.com/Velocity4705/polyglot/releases/download/v1.0.1/polyglot-linux-amd64
```

## Related Commands

- `polyglot version` - Show current version
- `polyglot check` - Check installed toolchains
- `polyglot config show` - Show configuration

## See Also

- [Installation Guide](QUICKSTART.md)
- [Configuration](CONFIGURATION.md)
- [Changelog](../CHANGELOG.md)
- [GitHub Releases](https://github.com/Velocity4705/polyglot/releases)
