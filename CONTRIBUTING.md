# Contributing to Polyglot

Thank you for your interest in contributing to Polyglot! This guide will help you get started.

## 🌟 Ways to Contribute

- 🐛 Report bugs
- 💡 Suggest features  
- 📝 Improve documentation
- 🔧 Fix issues
- ✨ Add new languages
- 🧪 Write tests
- 🎨 Improve UX

## 🚀 Quick Start

```bash
# Fork and clone
git clone https://github.com/YOUR_USERNAME/polyglot.git
cd polyglot

# Build
go build -o polyglot ./cmd/polyglot

# Test
./scripts/run-tests.sh
```

## 📋 Development Workflow

1. Create a branch: `git checkout -b feature/amazing-feature`
2. Make changes
3. Add tests
4. Run tests: `./scripts/run-tests.sh`
5. Commit: `git commit -m "feat: add amazing feature"`
6. Push: `git push origin feature/amazing-feature`
7. Open Pull Request

## 🔧 Adding a New Language

See full guide in the documentation. Quick steps:

1. Create handler in `internal/language/yourlang.go`
2. Add to `pkg/types/types.go`
3. Update `internal/detector/detector.go`
4. Register in `internal/language/registry.go`
5. Add test fixture in `test/fixtures/`
6. Write tests
7. Update documentation

## 📝 Code Style

- Follow Go conventions
- Run `go fmt ./...`
- Run `go vet ./...`
- Add comments for exported functions
- Keep functions small (<50 lines)

## 🧪 Testing

All new features must include tests:

```bash
# Unit tests
go test ./internal/... -v

# Integration tests
go test ./test/... -v

# All tests
./scripts/run-tests.sh
```

## 🙏 Thank You!

Every contribution makes Polyglot better!
