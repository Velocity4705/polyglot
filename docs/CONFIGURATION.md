# Polyglot Configuration Guide

Complete guide to configuring and customizing Polyglot.

## Table of Contents

- [Overview](#overview)
- [Configuration Files](#configuration-files)
- [Configuration Structure](#configuration-structure)
- [Managing Configuration](#managing-configuration)
- [Examples](#examples)
- [Best Practices](#best-practices)

## Overview

Polyglot supports two levels of configuration:

1. **Global Configuration** (`~/.polyglot/config.yaml`) - System-wide settings
2. **Project Configuration** (`.polyglot.yaml`) - Project-specific settings

Project configuration overrides global configuration, allowing for flexible customization.

## Configuration Files

### Global Configuration

**Location**: `~/.polyglot/config.yaml`

**Purpose**: 
- User-wide preferences
- Default settings for all projects
- Personal toolchain preferences

**Example**:
```yaml
auto_install:
  enabled: true
  method: system
  prompt: false

languages:
  python:
    command: python3.11
    version: "3.11"

environment:
  EDITOR: "vim"
```

### Project Configuration

**Location**: `.polyglot.yaml` (in project root)

**Purpose**:
- Project-specific settings
- Team-wide standards
- Override global settings
- Committed to version control

**Example**:
```yaml
languages:
  python:
    flags:
      - "-O"
    environment:
      PYTHONPATH: "./lib"

custom_extensions:
  .pyx: Python
```

## Configuration Structure

### Auto-Install Settings

Control automatic toolchain installation:

```yaml
auto_install:
  enabled: false      # Enable/disable auto-install
  method: system      # Installation method: system, script, managed
  prompt: true        # Prompt before installing
```

**Options**:
- `enabled`: `true` to auto-install, `false` to prompt
- `method`: 
  - `system`: Use system package manager (dnf, apt, brew, etc.)
  - `script`: Use official installation scripts
  - `managed`: Use Polyglot's managed toolchains (future)
- `prompt`: `true` to ask before installing, `false` to install automatically

### Language Configuration

Configure individual languages:

```yaml
languages:
  python:
    enabled: true                    # Enable/disable language
    command: python3.11              # Custom command path
    compile_cmd: ""                  # Custom compile command
    run_cmd: ""                      # Custom run command
    flags:                           # Compiler/interpreter flags
      - "-O"
      - "-W ignore"
    version: "3.11"                  # Preferred version
    install_method: system           # Installation method
    install_script: ""               # Custom install script URL
    binary_url: ""                   # Direct binary download URL
    environment:                     # Language-specific env vars
      PYTHONPATH: "/custom/path"
      PYTHONDONTWRITEBYTECODE: "1"
```

**Properties**:
- `enabled`: Enable or disable the language
- `command`: Override default command (e.g., `python3.11` instead of `python3`)
- `compile_cmd`: Custom compilation command for compiled languages
- `run_cmd`: Custom run command
- `flags`: Array of flags to pass to compiler/interpreter
- `version`: Preferred version (informational)
- `install_method`: How to install this language
- `install_script`: URL to custom installation script
- `binary_url`: URL to download pre-built binary
- `environment`: Language-specific environment variables

### Custom Extensions

Map file extensions to languages:

```yaml
custom_extensions:
  .pyx: Python          # Cython files
  .jsx: JavaScript      # React JSX
  .tsx: TypeScript      # React TypeScript
  .mjs: JavaScript      # ES modules
  .cjs: JavaScript      # CommonJS
  .vue: JavaScript      # Vue.js
  .svelte: JavaScript   # Svelte
```

**Use Cases**:
- Framework-specific extensions
- Dialect extensions (Cython, JSX, etc.)
- Module system extensions
- Custom project extensions

### Environment Variables

Set global environment variables:

```yaml
environment:
  PATH: "/custom/bin:$PATH"
  LANG: "en_US.UTF-8"
  EDITOR: "vim"
  PAGER: "less"
```

These are applied to all Polyglot executions. Language-specific environment variables (in `languages.<lang>.environment`) override these.

## Managing Configuration

### Initialize Configuration

Create a new configuration file:

```bash
polyglot config init
```

This creates `~/.polyglot/config.yaml` with default settings.

### View Configuration

Show current configuration (merged global + project):

```bash
polyglot config show
```

### Get Configuration Values

Get a specific configuration value:

```bash
# Auto-install settings
polyglot config get auto_install.enabled
polyglot config get auto_install.method

# Language settings
polyglot config get languages.python.version
polyglot config get languages.python.command

# Custom extensions
polyglot config get custom_extensions..pyx

# Environment variables
polyglot config get environment.EDITOR
```

### Set Configuration Values

Set configuration values:

```bash
# Auto-install settings
polyglot config set auto_install.enabled true
polyglot config set auto_install.method system
polyglot config set auto_install.prompt false

# Language settings
polyglot config set languages.python.enabled true
polyglot config set languages.python.version 3.11
polyglot config set languages.python.command python3.11

# Custom extensions (note the double dot)
polyglot config set custom_extensions..pyx Python
polyglot config set custom_extensions..jsx JavaScript

# Environment variables
polyglot config set environment.EDITOR vim
polyglot config set environment.PYTHONPATH /custom/path
```

**Note**: For custom extensions, use double dots (`..`) because the extension starts with a dot.

### Show Config Paths

Show where configuration files are located:

```bash
polyglot config path
```

Output:
```
Global config: /home/user/.polyglot/config.yaml
Project config: /home/user/project/.polyglot.yaml
```

## Examples

### Example 1: Python Development

**Global Config** (`~/.polyglot/config.yaml`):
```yaml
auto_install:
  enabled: true
  method: system
  prompt: false

languages:
  python:
    enabled: true
    command: python3.11
    version: "3.11"
    environment:
      PYTHONDONTWRITEBYTECODE: "1"

custom_extensions:
  .pyx: Python

environment:
  EDITOR: "vim"
```

**Usage**:
```bash
# Runs with Python 3.11, no .pyc files
polyglot run script.py

# Cython files work automatically
polyglot run module.pyx
```

### Example 2: Multi-Language Web Project

**Project Config** (`.polyglot.yaml`):
```yaml
languages:
  javascript:
    enabled: true
    command: node
    environment:
      NODE_ENV: "development"
  
  typescript:
    enabled: true
    command: ts-node
  
  python:
    enabled: true
    flags:
      - "-O"

custom_extensions:
  .jsx: JavaScript
  .tsx: TypeScript
  .mjs: JavaScript

environment:
  PORT: "3000"
```

**Usage**:
```bash
# All these work automatically
polyglot run server.js
polyglot run app.tsx
polyglot run api.py
polyglot run utils.mjs
```

### Example 3: Restricted Environment

Only allow specific languages:

```yaml
languages:
  python:
    enabled: true
  
  go:
    enabled: true
  
  # Disable everything else
  javascript:
    enabled: false
  ruby:
    enabled: false
  php:
    enabled: false
  java:
    enabled: false
```

**Result**:
```bash
$ polyglot run script.py
✓ Works

$ polyglot run app.js
✗ Error: language JavaScript is disabled in configuration
```

### Example 4: Custom Compiler Flags

**Rust with optimization**:
```yaml
languages:
  rust:
    enabled: true
    flags:
      - "--edition=2021"
      - "-O"
      - "--cfg=feature=\"custom\""
```

**C with warnings**:
```yaml
languages:
  c:
    enabled: true
    compile_cmd: gcc
    flags:
      - "-Wall"
      - "-Wextra"
      - "-O2"
      - "-std=c11"
```

### Example 5: Alternative Python Versions

**Per-project Python version**:

**Project A** (`.polyglot.yaml`):
```yaml
languages:
  python:
    command: python3.9
    version: "3.9"
```

**Project B** (`.polyglot.yaml`):
```yaml
languages:
  python:
    command: python3.11
    version: "3.11"
```

Each project uses its specified Python version automatically.

### Example 6: Framework-Specific Extensions

**React/Vue/Svelte project**:
```yaml
custom_extensions:
  .jsx: JavaScript
  .tsx: TypeScript
  .vue: JavaScript
  .svelte: JavaScript

languages:
  javascript:
    environment:
      NODE_ENV: "development"
  
  typescript:
    command: ts-node
```

## Best Practices

### 1. Use Project Config for Team Settings

Commit `.polyglot.yaml` to version control so the entire team uses the same settings:

```yaml
# .polyglot.yaml - committed to git
languages:
  python:
    version: "3.11"
    flags:
      - "-O"

custom_extensions:
  .pyx: Python
```

### 2. Use Global Config for Personal Preferences

Keep personal preferences in global config:

```yaml
# ~/.polyglot/config.yaml - not in git
auto_install:
  enabled: true
  prompt: false

environment:
  EDITOR: "vim"
  PAGER: "less"
```

### 3. Disable Unused Languages

Improve performance by disabling languages you don't use:

```yaml
languages:
  # Only enable what you need
  python:
    enabled: true
  javascript:
    enabled: true
  
  # Disable the rest
  ruby:
    enabled: false
  php:
    enabled: false
  # ... etc
```

### 4. Use Environment Variables for Secrets

Don't put secrets in config files. Use environment variables:

```yaml
# Bad - secrets in config
languages:
  python:
    environment:
      API_KEY: "secret123"

# Good - reference environment variables
languages:
  python:
    environment:
      API_KEY: "$API_KEY"
```

### 5. Document Custom Extensions

Add comments to explain custom extensions:

```yaml
custom_extensions:
  # Cython files
  .pyx: Python
  
  # React components
  .jsx: JavaScript
  .tsx: TypeScript
  
  # ES modules
  .mjs: JavaScript
```

### 6. Test Configuration Changes

After changing configuration, test it:

```bash
# View merged configuration
polyglot config show

# Test with a simple file
polyglot run test.py

# Check what's enabled
polyglot list
```

### 7. Use Version Control for Project Config

```bash
# Add project config to git
git add .polyglot.yaml
git commit -m "Add Polyglot configuration"

# Ignore global config
echo ".polyglot/config.yaml" >> ~/.gitignore_global
```

## Troubleshooting

### Configuration Not Loading

Check file locations:
```bash
polyglot config path
```

Verify file syntax:
```bash
# YAML syntax is strict about indentation
# Use 2 spaces, not tabs
```

### Custom Extension Not Working

Verify the extension is mapped:
```bash
polyglot config get custom_extensions..pyx
```

Check the language is enabled:
```bash
polyglot config get languages.python.enabled
```

### Language Disabled Error

Enable the language:
```bash
polyglot config set languages.python.enabled true
```

Or remove the language config to use defaults.

### Environment Variables Not Applied

Check environment configuration:
```bash
polyglot config show
```

Verify the syntax:
```yaml
# Correct
environment:
  VAR: "value"

# Wrong
environment:
  - VAR: "value"
```

## Advanced Usage

### Conditional Configuration

Use different configs for different environments:

```bash
# Development
cp .polyglot.dev.yaml .polyglot.yaml

# Production
cp .polyglot.prod.yaml .polyglot.yaml
```

### Configuration Templates

Create templates for common setups:

```bash
# Python data science template
cat > ~/.polyglot/templates/datascience.yaml <<EOF
languages:
  python:
    version: "3.11"
    environment:
      PYTHONPATH: "./lib"
  
  r:
    enabled: true
  
  julia:
    enabled: true
EOF

# Use template
cp ~/.polyglot/templates/datascience.yaml .polyglot.yaml
```

### Scripted Configuration

Automate configuration with scripts:

```bash
#!/bin/bash
# setup-polyglot.sh

polyglot config init
polyglot config set auto_install.enabled true
polyglot config set languages.python.version 3.11
polyglot config set custom_extensions..pyx Python
```

## Summary

Polyglot's configuration system provides:
- ✅ Global and project-level customization
- ✅ Language-specific settings
- ✅ Custom extension mapping
- ✅ Environment variable management
- ✅ Easy CLI management
- ✅ Team collaboration via version control

For more information, see:
- [MILESTONE6_COMPLETE.md](../MILESTONE6_COMPLETE.md) - Implementation details
- [QUICKSTART.md](QUICKSTART.md) - Getting started guide
- [EXAMPLES.md](EXAMPLES.md) - Usage examples
