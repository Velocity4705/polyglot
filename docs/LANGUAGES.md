# Supported Languages

This document lists all languages currently supported by Polyglot and their status.

## Milestone 1, 2 & 3 (Current) - 16 Languages! ✅

### Interpreted Languages

| Language   | Extension | Type        | Command    | Status |
|------------|-----------|-------------|------------|--------|
| Python     | .py       | Interpreted | python3    | ✅     |
| JavaScript | .js       | Interpreted | node       | ✅     |
| Ruby       | .rb       | Interpreted | ruby       | ✅     |
| PHP        | .php      | Interpreted | php        | ✅     |
| Perl       | .pl       | Interpreted | perl       | ✅     |
| Lua        | .lua      | Interpreted | lua        | ✅     |
| Shell      | .sh .bash | Interpreted | bash       | ✅     |

### Compiled Languages

| Language | Extension | Type     | Compiler/Runner | Status |
|----------|-----------|----------|-----------------|--------|
| Go       | .go       | Compiled | go run          | ✅     |
| Java     | .java     | Compiled | javac + java    | ✅     |
| C        | .c        | Compiled | gcc             | ✅     |
| C++      | .cpp .cc .cxx .C | Compiled | g++      | ✅     |
| Rust     | .rs       | Compiled | rustc           | ✅     |
| Zig      | .zig      | Compiled | zig run         | ✅     |
| Nim      | .nim      | Compiled | nim c           | ✅     |
| Crystal  | .cr       | Compiled | crystal run     | ✅     |
| D        | .d        | Compiled | dmd             | ✅     |

**Total: 16 languages supported!**

## Milestone 2 (Completed) ✅ - Compiled Languages

All Milestone 2 languages have been implemented and tested.

## Milestone 3 (Completed) ✅ - More Languages

All Milestone 3 languages have been implemented and tested.

## Milestone 4 (Planned) - Esoteric & Specialty

| Language | Extension | Type        | Command | Status |
|----------|-----------|-------------|---------|--------|
| Perl     | .pl       | Interpreted | perl    | 📋     |
| Lua      | .lua      | Interpreted | lua     | 📋     |
| Zig      | .zig      | Compiled    | zig     | 📋     |
| Nim      | .nim      | Compiled    | nim     | 📋     |
| Crystal  | .cr       | Compiled    | crystal | 📋     |
| D        | .d        | Compiled    | dmd     | 📋     |
| Shell    | .sh       | Interpreted | bash    | 📋     |

## Milestone 4 (Planned) - Esoteric & Specialty

| Language   | Extension | Type        | Command  | Status |
|------------|-----------|-------------|----------|--------|
| Haskell    | .hs       | Compiled    | ghc      | 📋     |
| OCaml      | .ml       | Compiled    | ocamlc   | 📋     |
| Erlang     | .erl      | Compiled    | erlc     | 📋     |
| Elixir     | .ex       | Interpreted | elixir   | 📋     |
| Kotlin     | .kt       | Compiled    | kotlinc  | 📋     |
| Scala      | .scala    | Compiled    | scalac   | 📋     |
| Clojure    | .clj      | Interpreted | clojure  | 📋     |
| Brainfuck  | .bf       | Interpreted | bf       | 📋     |
| Assembly   | .asm      | Compiled    | nasm     | 📋     |

## Legend

- ✅ Implemented and tested
- 📋 Planned for future milestone
- 🚧 Work in progress
- ❌ Not supported

## Testing Language Support

To check which languages are available on your system:

```bash
polyglot check
```

To see all languages polyglot can detect:

```bash
polyglot list
```

## Adding New Languages

See [CONTRIBUTING.md](../CONTRIBUTING.md) for instructions on adding support for new languages.

## Language-Specific Notes

### Python
- Uses `python3` by default
- Supports all Python 3.x versions

### Go
- Uses `go run` for quick execution
- Use `polyglot compile` for standalone binaries

### JavaScript
- Requires Node.js
- Does not support browser-specific APIs

### Ruby
- Requires Ruby interpreter
- Tested with Ruby 2.7+

### PHP
- Requires PHP CLI
- Tested with PHP 7.4+

## Requesting New Languages

To request support for a new language, please open an issue with:
- Language name
- File extension(s)
- Compiler/interpreter command
- Example hello world program
- Use case / why you need it
