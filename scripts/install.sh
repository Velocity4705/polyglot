#!/bin/bash
# Polyglot Quick Installer
# Usage: curl -sSL https://raw.githubusercontent.com/Velocity4705/polyglot/main/scripts/install.sh | bash

set -e

echo "Installing Polyglot..."

# Clone repository to temp directory
TEMP_DIR=$(mktemp -d)
echo "→ Using temp directory: $TEMP_DIR"
cd "$TEMP_DIR" || exit 1

echo "→ Cloning repository..."
if ! git clone --depth 1 https://github.com/Velocity4705/polyglot.git; then
    echo "✗ Failed to clone repository"
    exit 1
fi

cd polyglot || exit 1
echo "→ Current directory: $(pwd)"

# Build
echo "→ Building..."
if ! make build; then
    echo "✗ Build failed"
    exit 1
fi

# Install
echo "→ Installing to /usr/local/bin..."
if [ -w "/usr/local/bin" ]; then
    mv bin/polyglot /usr/local/bin/polyglot
else
    sudo mv bin/polyglot /usr/local/bin/polyglot
fi

# Cleanup
cd / || exit 1
rm -rf "$TEMP_DIR"

echo "✓ Polyglot installed successfully!"
echo ""
echo "Try it out:"
echo "  polyglot version"
echo "  polyglot list"
echo "  polyglot run hello.py"
