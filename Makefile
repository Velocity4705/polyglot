.PHONY: build install test clean run-tests help version update-check

VERSION := 1.0.1

help:
	@echo "Polyglot - Universal Compiler Wrapper v$(VERSION)"
	@echo ""
	@echo "Available targets:"
	@echo "  make build      - Build the polyglot binary"
	@echo "  make install    - Install polyglot to /usr/local/bin"
	@echo "  make test       - Run unit tests"
	@echo "  make run-tests  - Run integration tests"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make check      - Check code formatting and vet"
	@echo "  make fmt        - Format code"
	@echo "  make version    - Show version information"
	@echo "  make help       - Show this help message"

build:
	@echo "Building polyglot..."
	@go build -o bin/polyglot ./cmd/polyglot
	@echo "✓ Build complete: bin/polyglot"

install: build
	@echo "Installing polyglot..."
	@cp bin/polyglot /usr/local/bin/polyglot
	@echo "✓ Installed to /usr/local/bin/polyglot"

test:
	@echo "Running tests..."
	@go test ./internal/... ./pkg/... ./cmd/...

check:
	@echo "Running go vet..."
	@go vet ./...
	@echo "Checking formatting..."
	@test -z "$$(gofmt -l .)" || (echo "Code needs formatting. Run 'make fmt'" && exit 1)
	@echo "✓ All checks passed"

fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Code formatted"

clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean
	@echo "✓ Clean complete"

run-tests: build
	@echo "Testing Python..."
	@./bin/polyglot run test/fixtures/hello.py
	@echo ""
	@echo "Testing Go..."
	@./bin/polyglot run test/fixtures/hello.go
	@echo ""
	@echo "Testing JavaScript..."
	@./bin/polyglot run test/fixtures/hello.js
	@echo ""
	@echo "Testing with arguments..."
	@./bin/polyglot run test/fixtures/args.py --args arg1,arg2,arg3

version:
	@echo "Polyglot v$(VERSION)"
	@echo "Production Ready - Auto-Update Feature Added!"

update-check: build
	@echo "Checking for updates..."
	@./bin/polyglot update --check

all: clean fmt check build test
	@echo "✓ All tasks complete"
