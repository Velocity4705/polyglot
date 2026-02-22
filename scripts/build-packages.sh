#!/bin/bash
# Build all distribution packages for Polyglot

set -e

VERSION="0.6.0"
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BUILD_DIR="$PROJECT_ROOT/dist"

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Polyglot Package Builder v${VERSION}${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Clean and create build directory
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

# Build binaries for all platforms
build_binaries() {
    echo -e "${BLUE}Building binaries...${NC}"
    
    platforms=(
        "linux/amd64"
        "linux/arm64"
        "darwin/amd64"
        "darwin/arm64"
        "windows/amd64"
    )
    
    for platform in "${platforms[@]}"; do
        IFS='/' read -r -a parts <<< "$platform"
        GOOS="${parts[0]}"
        GOARCH="${parts[1]}"
        
        output_name="polyglot-${GOOS}-${GOARCH}"
        if [ "$GOOS" = "windows" ]; then
            output_name="${output_name}.exe"
        fi
        
        echo -e "${YELLOW}→${NC} Building $output_name..."
        
        GOOS=$GOOS GOARCH=$GOARCH go build \
            -ldflags="-s -w -X main.version=${VERSION}" \
            -o "$BUILD_DIR/$output_name" \
            ./cmd/polyglot
        
        echo -e "${GREEN}✓${NC} Built $output_name"
    done
    
    echo ""
}

# Create archives
create_archives() {
    echo -e "${BLUE}Creating archives...${NC}"
    
    cd "$BUILD_DIR"
    
    for binary in polyglot-*; do
        if [[ $binary == *.exe ]]; then
            # Windows: create zip
            archive_name="${binary%.exe}.zip"
            echo -e "${YELLOW}→${NC} Creating $archive_name..."
            zip -q "$archive_name" "$binary" ../README.md ../LICENSE
        else
            # Unix: create tar.gz
            archive_name="${binary}.tar.gz"
            echo -e "${YELLOW}→${NC} Creating $archive_name..."
            tar czf "$archive_name" "$binary" ../README.md ../LICENSE
        fi
        echo -e "${GREEN}✓${NC} Created $archive_name"
    done
    
    cd "$PROJECT_ROOT"
    echo ""
}

# Generate checksums
generate_checksums() {
    echo -e "${BLUE}Generating checksums...${NC}"
    
    cd "$BUILD_DIR"
    
    # SHA256 checksums
    sha256sum polyglot-* > SHA256SUMS
    echo -e "${GREEN}✓${NC} Generated SHA256SUMS"
    
    # MD5 checksums
    md5sum polyglot-* > MD5SUMS
    echo -e "${GREEN}✓${NC} Generated MD5SUMS"
    
    cd "$PROJECT_ROOT"
    echo ""
}

# Build Docker image
build_docker() {
    echo -e "${BLUE}Building Docker image...${NC}"
    
    docker build -t polyglot:${VERSION} -t polyglot:latest .
    echo -e "${GREEN}✓${NC} Built Docker image"
    
    # Save image
    docker save polyglot:${VERSION} | gzip > "$BUILD_DIR/polyglot-${VERSION}-docker.tar.gz"
    echo -e "${GREEN}✓${NC} Saved Docker image to dist/"
    
    echo ""
}

# Create Debian package
build_deb() {
    echo -e "${BLUE}Building Debian package...${NC}"
    
    DEB_DIR="$BUILD_DIR/deb"
    mkdir -p "$DEB_DIR/DEBIAN"
    mkdir -p "$DEB_DIR/usr/local/bin"
    mkdir -p "$DEB_DIR/usr/share/doc/polyglot"
    
    # Copy files
    cp "$BUILD_DIR/polyglot-linux-amd64" "$DEB_DIR/usr/local/bin/polyglot"
    chmod +x "$DEB_DIR/usr/local/bin/polyglot"
    cp README.md LICENSE "$DEB_DIR/usr/share/doc/polyglot/"
    cp packaging/debian/control "$DEB_DIR/DEBIAN/"
    
    # Build package
    dpkg-deb --build "$DEB_DIR" "$BUILD_DIR/polyglot_${VERSION}_amd64.deb" 2>/dev/null || true
    
    if [ -f "$BUILD_DIR/polyglot_${VERSION}_amd64.deb" ]; then
        echo -e "${GREEN}✓${NC} Built Debian package"
    else
        echo -e "${YELLOW}⚠${NC} Debian package build skipped (dpkg-deb not available)"
    fi
    
    rm -rf "$DEB_DIR"
    echo ""
}

# Create RPM package
build_rpm() {
    echo -e "${BLUE}Building RPM package...${NC}"
    
    if command -v rpmbuild &> /dev/null; then
        RPM_DIR="$BUILD_DIR/rpm"
        mkdir -p "$RPM_DIR"/{BUILD,RPMS,SOURCES,SPECS,SRPMS}
        
        # Create source tarball
        tar czf "$RPM_DIR/SOURCES/polyglot-${VERSION}.tar.gz" \
            --transform "s,^,polyglot-${VERSION}/," \
            cmd/ internal/ pkg/ go.mod go.sum README.md LICENSE
        
        # Copy spec file
        cp packaging/rpm/polyglot.spec "$RPM_DIR/SPECS/"
        
        # Build RPM
        rpmbuild --define "_topdir $RPM_DIR" -ba "$RPM_DIR/SPECS/polyglot.spec" 2>/dev/null || true
        
        # Copy RPM to dist
        if [ -f "$RPM_DIR/RPMS/x86_64/polyglot-${VERSION}-1.x86_64.rpm" ]; then
            cp "$RPM_DIR/RPMS/x86_64/polyglot-${VERSION}-1.x86_64.rpm" "$BUILD_DIR/"
            echo -e "${GREEN}✓${NC} Built RPM package"
        else
            echo -e "${YELLOW}⚠${NC} RPM package build skipped (rpmbuild not available)"
        fi
        
        rm -rf "$RPM_DIR"
    else
        echo -e "${YELLOW}⚠${NC} RPM package build skipped (rpmbuild not available)"
    fi
    
    echo ""
}

# Summary
show_summary() {
    echo -e "${BLUE}========================================${NC}"
    echo -e "${GREEN}Build Complete!${NC}"
    echo -e "${BLUE}========================================${NC}"
    echo ""
    echo "Artifacts created in: $BUILD_DIR"
    echo ""
    echo "Binaries:"
    ls -lh "$BUILD_DIR"/polyglot-* 2>/dev/null | awk '{print "  " $9 " (" $5 ")"}'
    echo ""
    echo "Archives:"
    ls -lh "$BUILD_DIR"/*.{tar.gz,zip} 2>/dev/null | awk '{print "  " $9 " (" $5 ")"}'
    echo ""
    echo "Packages:"
    ls -lh "$BUILD_DIR"/*.{deb,rpm} 2>/dev/null | awk '{print "  " $9 " (" $5 ")"}'
    echo ""
}

# Main execution
main() {
    build_binaries
    create_archives
    generate_checksums
    
    # Optional: build Docker image
    if command -v docker &> /dev/null; then
        build_docker
    else
        echo -e "${YELLOW}⚠${NC} Docker not available, skipping Docker image build"
        echo ""
    fi
    
    # Optional: build packages
    build_deb
    build_rpm
    
    show_summary
}

# Run main
main
