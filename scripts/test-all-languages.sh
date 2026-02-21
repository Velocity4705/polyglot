#!/bin/bash

set -e

echo "================================"
echo "Polyglot - Testing All Languages"
echo "================================"
echo ""

# Build first
echo "Building polyglot..."
make build
echo "✓ Build successful"
echo ""

# Test each language
echo "================================"
echo "Testing All 16 Languages"
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

# Interpreted Languages
echo "--- Interpreted Languages ---"
test_language "test/fixtures/hello.py" "Python"
test_language "test/fixtures/hello.js" "JavaScript"
test_language "test/fixtures/hello.rb" "Ruby"
test_language "test/fixtures/hello.php" "PHP"
test_language "test/fixtures/hello.pl" "Perl"
test_language "test/fixtures/hello.lua" "Lua"
test_language "test/fixtures/hello.sh" "Shell"

echo ""
echo "--- Compiled Languages ---"
test_language "test/fixtures/hello.go" "Go"
test_language "test/fixtures/Hello.java" "Java"
test_language "test/fixtures/hello.c" "C"
test_language "test/fixtures/hello.cpp" "C++"
test_language "test/fixtures/hello.rs" "Rust"
test_language "test/fixtures/hello.zig" "Zig"
test_language "test/fixtures/hello.nim" "Nim"
test_language "test/fixtures/hello.cr" "Crystal"
test_language "test/fixtures/hello.d" "D"

echo ""
echo "================================"
echo "All tests completed!"
echo "================================"
