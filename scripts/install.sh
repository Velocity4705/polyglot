#!/bin/bash
# Polyglot Quick Installer
# Usage: curl -sSL https://raw.githubusercontent.com/Velocity4705/polyglot/main/scripts/install.sh | bash

set -e

echo "Installing Polyglot..."

# Clone repository to temp directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

echo "→ Cloning repository..."
git clone --depth 1 https://github.com/Velocity4705/polyglot.git
cd polyglot

# Build
echo "→ Building..."
make build

# Install
echo "→ Installing to /usr/local/bin..."
if [ -w "/usr/local/bin" ]; then
    mv bin/polyglot /usr/local/bin/polyglot
else
    sudo mv bin/polyglot /usr/local/bin/polyglot
fi

# Cleanup
cd /
rm -rf "$TEMP_DIR"

echo "✓ Polyglot installed successfully!"
echo ""
echo "Try it out:"
echo "  polyglot version"
echo "  polyglot list"
echo "  polyglot run hello.py"
