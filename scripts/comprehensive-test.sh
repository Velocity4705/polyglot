#!/bin/bash

# Comprehensive Test Suite for Polyglot v1.0.1
# Tests all features to catch first-day problems

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

PASSED=0
FAILED=0
SKIPPED=0

echo -e "${BLUE}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║  Polyglot v1.0.1 - Comprehensive Test Suite              ║${NC}"
echo -e "${BLUE}║  Testing all features before launch                       ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════════╝${NC}"
echo ""

# Function to run a test
run_test() {
    local test_name="$1"
    local test_command="$2"
    local expect_fail="${3:-false}"
    
    echo -n "Testing: $test_name... "
    
    if eval "$test_command" > /tmp/test_output.log 2>&1; then
        if [ "$expect_fail" = "true" ]; then
            echo -e "${RED}✗ FAILED${NC} (expected to fail but passed)"
            FAILED=$((FAILED + 1))
            return 1
        else
            echo -e "${GREEN}✓ PASSED${NC}"
            PASSED=$((PASSED + 1))
            return 0
        fi
    else
        if [ "$expect_fail" = "true" ]; then
            echo -e "${GREEN}✓ PASSED${NC} (failed as expected)"
            PASSED=$((PASSED + 1))
            return 0
        else
            echo -e "${RED}✗ FAILED${NC}"
            echo "  Output: $(cat /tmp/test_output.log | head -3)"
            FAILED=$((FAILED + 1))
            return 1
        fi
    fi
}

# Function to skip a test
skip_test() {
    local test_name="$1"
    local reason="$2"
    echo -e "Testing: $test_name... ${YELLOW}⊘ SKIPPED${NC} ($reason)"
    SKIPPED=$((SKIPPED + 1))
}

echo -e "${BLUE}═══ Phase 1: Build & Binary Tests ═══${NC}"
echo ""

run_test "Clean build" "make clean"
run_test "Build binary" "make build"
run_test "Binary exists" "test -f bin/polyglot"
run_test "Binary is executable" "test -x bin/polyglot"
run_test "Binary runs" "./bin/polyglot version"

echo ""
echo -e "${BLUE}═══ Phase 2: Core Commands ═══${NC}"
echo ""

run_test "Help command" "./bin/polyglot --help"
run_test "Version command" "./bin/polyglot version"
run_test "List command" "./bin/polyglot list"
run_test "Check command" "./bin/polyglot check"

echo ""
echo -e "${BLUE}═══ Phase 3: Configuration Commands ═══${NC}"
echo ""

run_test "Config help" "./bin/polyglot config --help"
run_test "Config show" "./bin/polyglot config show"
run_test "Config path" "./bin/polyglot config path"

echo ""
echo -e "${BLUE}═══ Phase 4: Update Command ═══${NC}"
echo ""

run_test "Update help" "./bin/polyglot update --help"
run_test "Update check (expect 404)" "./bin/polyglot update --check" "true"

echo ""
echo -e "${BLUE}═══ Phase 5: Watch Command ═══${NC}"
echo ""

run_test "Watch help" "./bin/polyglot watch --help"

echo ""
echo -e "${BLUE}═══ Phase 6: Run Command Tests ═══${NC}"
echo ""

# Test Python
if command -v python3 &> /dev/null; then
    run_test "Run Python file" "./bin/polyglot run test/fixtures/hello.py"
    run_test "Run Python with quiet" "./bin/polyglot run -q test/fixtures/hello.py"
    run_test "Run Python with verbose" "./bin/polyglot run -v test/fixtures/hello.py"
else
    skip_test "Run Python file" "python3 not installed"
fi

# Test JavaScript
if command -v node &> /dev/null; then
    run_test "Run JavaScript file" "./bin/polyglot run test/fixtures/hello.js"
else
    skip_test "Run JavaScript file" "node not installed"
fi

# Test Go
if command -v go &> /dev/null; then
    run_test "Run Go file" "./bin/polyglot run test/fixtures/hello.go"
else
    skip_test "Run Go file" "go not installed"
fi

# Test Ruby
if command -v ruby &> /dev/null; then
    run_test "Run Ruby file" "./bin/polyglot run test/fixtures/hello.rb"
else
    skip_test "Run Ruby file" "ruby not installed"
fi

# Test Shell
if command -v bash &> /dev/null; then
    run_test "Run Shell file" "./bin/polyglot run test/fixtures/hello.sh"
else
    skip_test "Run Shell file" "bash not installed"
fi

echo ""
echo -e "${BLUE}═══ Phase 7: Compile Command Tests ═══${NC}"
echo ""

# Test C compilation
if command -v gcc &> /dev/null; then
    run_test "Compile C file" "./bin/polyglot compile test/fixtures/hello.c"
    run_test "Cleanup C binary" "rm -f test/fixtures/hello"
else
    skip_test "Compile C file" "gcc not installed"
fi

# Test Rust compilation
if command -v rustc &> /dev/null; then
    run_test "Compile Rust file" "./bin/polyglot compile test/fixtures/hello.rs"
    run_test "Cleanup Rust binary" "rm -f test/fixtures/hello"
else
    skip_test "Compile Rust file" "rustc not installed"
fi

echo ""
echo -e "${BLUE}═══ Phase 8: Error Handling Tests ═══${NC}"
echo ""

run_test "Non-existent file" "./bin/polyglot run nonexistent.py" "true"
run_test "Unsupported extension" "./bin/polyglot run test.xyz" "true"
run_test "Invalid command" "./bin/polyglot invalid" "true"

echo ""
echo -e "${BLUE}═══ Phase 9: Argument Passing Tests ═══${NC}"
echo ""

if command -v python3 &> /dev/null; then
    run_test "Pass arguments to Python" "./bin/polyglot run test/fixtures/args.py --args arg1,arg2,arg3"
else
    skip_test "Pass arguments to Python" "python3 not installed"
fi

echo ""
echo -e "${BLUE}═══ Phase 10: Unit Tests ═══${NC}"
echo ""

run_test "Config tests" "go test ./internal/config -v"
run_test "Detector tests" "go test ./internal/detector -v"
run_test "Updater tests" "go test ./internal/updater -v"

echo ""
echo -e "${BLUE}═══ Phase 11: Code Quality Tests ═══${NC}"
echo ""

run_test "Go vet" "go vet ./cmd/... ./internal/... ./pkg/..."
run_test "Go fmt check" "test -z \"\$(gofmt -l . | grep -v vendor)\""

echo ""
echo -e "${BLUE}═══ Phase 12: Documentation Tests ═══${NC}"
echo ""

run_test "README exists" "test -f README.md"
run_test "CHANGELOG exists" "test -f CHANGELOG.md"
run_test "LICENSE exists" "test -f LICENSE"
run_test "CONTRIBUTING exists" "test -f CONTRIBUTING.md"
run_test "Auto-update docs exist" "test -f docs/AUTO_UPDATE.md"
run_test "Quick start exists" "test -f docs/QUICKSTART.md"
run_test "Configuration docs exist" "test -f docs/CONFIGURATION.md"

echo ""
echo -e "${BLUE}═══ Phase 13: File Structure Tests ═══${NC}"
echo ""

run_test "cmd directory exists" "test -d cmd/polyglot"
run_test "internal directory exists" "test -d internal"
run_test "pkg directory exists" "test -d pkg"
run_test "test directory exists" "test -d test"
run_test "docs directory exists" "test -d docs"
run_test "scripts directory exists" "test -d scripts"

echo ""
echo -e "${BLUE}═══ Phase 14: Version Consistency Tests ═══${NC}"
echo ""

run_test "Version in binary" "./bin/polyglot version | grep -q '1.0.1'"
run_test "Version in version.go" "grep -q '1.0.1' internal/cli/version.go"
run_test "Version in main.go" "grep -q '1.0.1' cmd/polyglot/main.go"
run_test "Version in Makefile" "grep -q '1.0.1' Makefile"

echo ""
echo -e "${BLUE}═══ Phase 15: Integration Tests ═══${NC}"
echo ""

# Test full workflow
if command -v python3 &> /dev/null; then
    echo "print('Integration test')" > /tmp/integration_test.py
    run_test "Full workflow: detect and run" "./bin/polyglot run /tmp/integration_test.py"
    rm -f /tmp/integration_test.py
else
    skip_test "Full workflow test" "python3 not installed"
fi

echo ""
echo -e "${BLUE}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║                    Test Results Summary                   ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "  ${GREEN}✓ Passed:${NC}  $PASSED"
echo -e "  ${RED}✗ Failed:${NC}  $FAILED"
echo -e "  ${YELLOW}⊘ Skipped:${NC} $SKIPPED"
echo ""

TOTAL=$((PASSED + FAILED + SKIPPED))
echo -e "  Total tests: $TOTAL"

if [ $FAILED -eq 0 ]; then
    echo ""
    echo -e "${GREEN}╔════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${GREEN}║  🎉 ALL TESTS PASSED! Ready for launch! 🚀                ║${NC}"
    echo -e "${GREEN}╚════════════════════════════════════════════════════════════╝${NC}"
    exit 0
else
    echo ""
    echo -e "${RED}╔════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║  ⚠️  SOME TESTS FAILED - Review before launch             ║${NC}"
    echo -e "${RED}╚════════════════════════════════════════════════════════════╝${NC}"
    exit 1
fi
