#!/bin/bash
# Comprehensive test runner for Polyglot

set -e

echo "=========================================="
echo "Polyglot Test Suite"
echo "=========================================="
echo ""

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Build first
echo -e "${BLUE}Building Polyglot...${NC}"
go build -o bin/polyglot ./cmd/polyglot
echo -e "${GREEN}✓ Build successful${NC}"
echo ""

# Run unit tests
echo -e "${BLUE}Running Unit Tests...${NC}"
echo "----------------------------------------"
go test ./internal/... ./pkg/... -v -cover
UNIT_EXIT=$?
echo ""

if [ $UNIT_EXIT -eq 0 ]; then
    echo -e "${GREEN}✓ Unit tests passed${NC}"
else
    echo -e "${RED}✗ Unit tests failed${NC}"
fi
echo ""

# Run integration tests
echo -e "${BLUE}Running Integration Tests...${NC}"
echo "----------------------------------------"
go test ./test/... -v
INTEGRATION_EXIT=$?
echo ""

if [ $INTEGRATION_EXIT -eq 0 ]; then
    echo -e "${GREEN}✓ Integration tests passed${NC}"
else
    echo -e "${RED}✗ Integration tests failed${NC}"
fi
echo ""

# Run benchmarks
echo -e "${BLUE}Running Benchmarks...${NC}"
echo "----------------------------------------"
go test ./test/... -bench=. -benchmem
echo ""

# Code quality checks
echo -e "${BLUE}Running Code Quality Checks...${NC}"
echo "----------------------------------------"

echo "Running go vet..."
go vet ./...
VET_EXIT=$?
if [ $VET_EXIT -eq 0 ]; then
    echo -e "${GREEN}✓ go vet passed${NC}"
else
    echo -e "${RED}✗ go vet failed${NC}"
fi

echo "Checking formatting..."
UNFORMATTED=$(gofmt -l .)
if [ -z "$UNFORMATTED" ]; then
    echo -e "${GREEN}✓ Code is properly formatted${NC}"
else
    echo -e "${RED}✗ Code needs formatting:${NC}"
    echo "$UNFORMATTED"
    FMT_EXIT=1
fi
echo ""

# Test coverage
echo -e "${BLUE}Generating Test Coverage...${NC}"
echo "----------------------------------------"
go test ./internal/... ./pkg/... -coverprofile=coverage.out
go tool cover -func=coverage.out | tail -1
echo ""

# Summary
echo "=========================================="
echo "Test Summary"
echo "=========================================="

TOTAL_FAILURES=0

if [ $UNIT_EXIT -eq 0 ]; then
    echo -e "${GREEN}✓ Unit Tests: PASSED${NC}"
else
    echo -e "${RED}✗ Unit Tests: FAILED${NC}"
    TOTAL_FAILURES=$((TOTAL_FAILURES + 1))
fi

if [ $INTEGRATION_EXIT -eq 0 ]; then
    echo -e "${GREEN}✓ Integration Tests: PASSED${NC}"
else
    echo -e "${RED}✗ Integration Tests: FAILED${NC}"
    TOTAL_FAILURES=$((TOTAL_FAILURES + 1))
fi

if [ $VET_EXIT -eq 0 ]; then
    echo -e "${GREEN}✓ Code Quality: PASSED${NC}"
else
    echo -e "${RED}✗ Code Quality: FAILED${NC}"
    TOTAL_FAILURES=$((TOTAL_FAILURES + 1))
fi

if [ -n "$UNFORMATTED" ]; then
    echo -e "${RED}✗ Code Formatting: FAILED${NC}"
    TOTAL_FAILURES=$((TOTAL_FAILURES + 1))
else
    echo -e "${GREEN}✓ Code Formatting: PASSED${NC}"
fi

echo ""

if [ $TOTAL_FAILURES -eq 0 ]; then
    echo -e "${GREEN}=========================================="
    echo "✓ ALL TESTS PASSED!"
    echo -e "==========================================${NC}"
    exit 0
else
    echo -e "${RED}=========================================="
    echo "✗ $TOTAL_FAILURES TEST SUITE(S) FAILED"
    echo -e "==========================================${NC}"
    exit 1
fi
