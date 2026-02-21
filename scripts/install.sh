#!/bin/bash

set -e

echo "Installing Polyglot..."

# Build the binary
echo "Building..."
go build -o polyglot ./cmd/polyglot

# Determine install location
INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
    echo "Need sudo permissions to install to $INSTALL_DIR"
    sudo mv polyglot "$INSTALL_DIR/polyglot"
    sudo chmod +x "$INSTALL_DIR/polyglot"
else
    mv polyglot "$INSTALL_DIR/polyglot"
    chmod +x "$INSTALL_DIR/polyglot"
fi

echo "✓ Polyglot installed successfully!"
echo ""
echo "Try it out:"
echo "  polyglot --help"
echo "  polyglot check"
echo "  polyglot list"
