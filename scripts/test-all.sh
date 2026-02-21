#!/bin/bash

set -e

echo "================================"
echo "Polyglot Test Suite"
echo "================================"
echo ""

# Build first
echo "Building polyglot..."
make build
echo "✓ Build successful"
echo ""

# Run unit tests
echo "Running unit tests..."
go test ./...
echo "✓ Unit tests passed"
echo ""

# Check toolchains
echo "Checking toolchains..."
./bin/polyglot check
echo ""

# Test each language
echo "================================"
echo "Testing Language Support"
echo "================================"
echo ""

test_language() {
    local file=$1
    local lang=$2
    
    if [ -f "$file" ]; then
        echo "Testing $lang..."
        if ./bin/polyglot run "$file" > /dev/null 2>&1; then
            echo "✓ $lang works"
        else
            echo "✗ $lang failed (toolchain may not be installed)"
        fi
    fi
}

test_language "test/fixtures/hello.py" "Python"
test_language "test/fixtures/hello.go" "Go"
test_language "test/fixtures/hello.js" "JavaScript"
test_language "test/fixtures/hello.rb" "Ruby"
test_language "test/fixtures/hello.php" "PHP"

echo ""
echo "================================"
echo "All tests completed!"
echo "================================"
