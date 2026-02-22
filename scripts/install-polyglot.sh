#!/bin/bash
# Polyglot Universal Installation Script
# Supports: Linux, macOS, Windows (WSL)

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# Version
VERSION="0.6.0"

echo -e "${CYAN}╔════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║     Polyglot Installer v${VERSION}        ║${NC}"
echo -e "${CYAN}║  Universal Compiler Wrapper           ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════╝${NC}"
echo ""

# Detect OS and Architecture
OS="$(uname -s)"
ARCH="$(uname -m)"

echo -e "${BLUE}→${NC} Detected OS: ${GREEN}$OS${NC}"
echo -e "${BLUE}→${NC} Architecture: ${GREEN}$ARCH${NC}"
echo ""

# Check prerequisites
check_prerequisites() {
    echo -e "${BLUE}Checking prerequisites...${NC}"
    
    if ! command -v go &> /dev/null; then
        echo -e "${YELLOW}⚠${NC} Go is not installed"
        echo -e "${YELLOW}  Building from source requires Go 1.21+${NC}"
        echo -e "${YELLOW}  Download: https://go.dev/dl/${NC}"
        HAS_GO=false
    else
        GO_VERSION=$(go version | awk '{print $3}')
        echo -e "${GREEN}✓${NC} Go ${GO_VERSION} found"
        HAS_GO=true
    fi
    
    if ! command -v git &> /dev/null; then
        echo -e "${YELLOW}⚠${NC} Git is not installed"
        HAS_GIT=false
    else
        echo -e "${GREEN}✓${NC} Git found"
        HAS_GIT=true
    fi
    
    echo ""
}

# Build from source
build_from_source() {
    echo -e "${BLUE}Building Polyglot from source...${NC}"
    echo ""
    
    if [ "$HAS_GO" = false ]; then
        echo -e "${RED}✗${NC} Go is required to build from source"
        echo -e "${YELLOW}  Install Go: https://go.dev/dl/${NC}"
        exit 1
    fi
    
    # Check if we're in the polyglot directory
    if [ ! -f "go.mod" ]; then
        if [ "$HAS_GIT" = false ]; then
            echo -e "${RED}✗${NC} Git is required to clone the repository"
            exit 1
        fi
        
        echo -e "${BLUE}→${NC} Cloning repository..."
        git clone https://github.com/yourusername/polyglot.git
        cd polyglot
    fi
    
    echo -e "${BLUE}→${NC} Building binary..."
    go build -ldflags="-s -w" -o polyglot ./cmd/polyglot
    
    echo -e "${GREEN}✓${NC} Build successful"
    echo ""
}

# Install binary
install_binary() {
    echo -e "${BLUE}Installing Polyglot...${NC}"
    
    INSTALL_DIR="/usr/local/bin"
    
    if [ ! -w "$INSTALL_DIR" ]; then
        echo -e "${YELLOW}→${NC} Requesting sudo access for installation..."
        sudo mv polyglot "$INSTALL_DIR/polyglot"
        sudo chmod +x "$INSTALL_DIR/polyglot"
    else
        mv polyglot "$INSTALL_DIR/polyglot"
        chmod +x "$INSTALL_DIR/polyglot"
    fi
    
    echo -e "${GREEN}✓${NC} Installed to $INSTALL_DIR/polyglot"
    echo ""
}

# Download pre-built binary
download_binary() {
    echo -e "${BLUE}Downloading pre-built binary...${NC}"
    echo ""
    
    # Determine download URL
    case "$OS" in
        Linux)
            case "$ARCH" in
                x86_64)
                    BINARY="polyglot-linux-amd64"
                    ;;
                aarch64|arm64)
                    BINARY="polyglot-linux-arm64"
                    ;;
                *)
                    echo -e "${RED}✗${NC} Unsupported architecture: $ARCH"
                    exit 1
                    ;;
            esac
            ;;
        Darwin)
            case "$ARCH" in
                x86_64)
                    BINARY="polyglot-darwin-amd64"
                    ;;
                arm64)
                    BINARY="polyglot-darwin-arm64"
                    ;;
                *)
                    echo -e "${RED}✗${NC} Unsupported architecture: $ARCH"
                    exit 1
                    ;;
            esac
            ;;
        *)
            echo -e "${RED}✗${NC} Unsupported OS: $OS"
            exit 1
            ;;
    esac
    
    URL="https://github.com/yourusername/polyglot/releases/download/v${VERSION}/${BINARY}"
    
    echo -e "${BLUE}→${NC} Downloading from: $URL"
    
    if command -v curl &> /dev/null; then
        curl -L -o polyglot "$URL"
    elif command -v wget &> /dev/null; then
        wget -O polyglot "$URL"
    else
        echo -e "${RED}✗${NC} Neither curl nor wget found"
        exit 1
    fi
    
    chmod +x polyglot
    echo -e "${GREEN}✓${NC} Download successful"
    echo ""
}

# Main installation flow
main() {
    check_prerequisites
    
    echo -e "${BLUE}Choose installation method:${NC}"
    echo "  1) Build from source (recommended)"
    echo "  2) Download pre-built binary"
    echo "  3) Cancel"
    echo ""
    read -p "Enter choice [1-3]: " choice
    echo ""
    
    case $choice in
        1)
            build_from_source
            install_binary
            ;;
        2)
            download_binary
            install_binary
            ;;
        3)
            echo "Installation cancelled"
            exit 0
            ;;
        *)
            echo -e "${RED}✗${NC} Invalid choice"
            exit 1
            ;;
    esac
    
    # Verify installation
    if command -v polyglot &> /dev/null; then
        echo -e "${CYAN}╔════════════════════════════════════════╗${NC}"
        echo -e "${CYAN}║   ${GREEN}✓ Installation Successful!${CYAN}        ║${NC}"
        echo -e "${CYAN}╚════════════════════════════════════════╝${NC}"
        echo ""
        echo -e "${BLUE}Verify installation:${NC}"
        echo "  $ polyglot version"
        echo ""
        echo -e "${BLUE}Get started:${NC}"
        echo "  $ polyglot run hello.py"
        echo "  $ polyglot list"
        echo "  $ polyglot --help"
        echo ""
        echo -e "${BLUE}Documentation:${NC}"
        echo "  https://github.com/yourusername/polyglot"
        echo ""
    else
        echo -e "${RED}✗${NC} Installation failed"
        echo -e "${YELLOW}  Please check the error messages above${NC}"
        exit 1
    fi
}

# Run main
main
