# Release Notes - Polyglot v1.0.1

**Release Date:** February 22, 2026  
**Status:** Production Ready

## What's New

### Auto-Update Feature
- Automatic update checking and installation
- GitHub release integration
- Platform-specific binary downloads
- Automatic backup and rollback on failure
- Package manager integration (Homebrew, APT, Chocolatey)
- Background update checks every 24 hours
- Manual update command: `polyglot update`

### Bug Fixes
- Fixed type conversion issues in run and watch commands
- Fixed test expectations for language handlers
- Fixed CI/CD configuration to properly exclude test fixtures
- Updated integration tests to reflect current version

### CI/CD Improvements
- Fixed GitHub Actions workflow to properly handle test fixtures
- Updated golangci-lint configuration with non-deprecated options
- Added proper dependency verification steps
- Improved test coverage reporting

## Installation

### New Installations
```bash
curl -sSL https://raw.githubusercontent.com/Velocity4705/polyglot/main/scripts/install.sh | bash
```

### Upgrading from v1.0.0
```bash
polyglot update
```

Or manually:
```bash
curl -sSL https://raw.githubusercontent.com/Velocity4705/polyglot/main/scripts/install.sh | bash
```

## Documentation

- [Auto-Update Documentation](docs/AUTO_UPDATE.md)
- [Configuration Guide](docs/CONFIGURATION.md)
- [Quick Start Guide](docs/QUICKSTART.md)
- [Upgrade Guide](UPGRADE_GUIDE.md)

## Testing

All 53 automated tests passing:
- ✓ Unit tests (config, detector, language handlers, updater)
- ✓ Integration tests (run, compile, watch, config commands)
- ✓ Build verification
- ✓ Code quality checks

## Known Issues

None at this time.

## Contributors

Thank you to all contributors who made this release possible!

## Next Steps

See [ROADMAP.md](ROADMAP.md) for upcoming features and improvements.
