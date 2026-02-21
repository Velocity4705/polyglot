package types

// LanguageType represents whether a language is compiled or interpreted
type LanguageType int

const (
	Interpreted LanguageType = iota
	Compiled
)

// Language represents a programming language
type Language int

const (
	Unknown Language = iota
	Python
	Go
	JavaScript
	Java
	C
	Cpp
	Rust
	Ruby
	PHP
	Perl
	Lua
	Zig
	Nim
	Crystal
	D
	Shell
	Haskell
	OCaml
	Elixir
	Erlang
	FSharp
	Kotlin
	Scala
	Groovy
	Swift
	Dart
	TypeScript
	R
	Julia
	Brainfuck
)

func (l Language) String() string {
	return l.Name()
}

func (l Language) Name() string {
	switch l {
	case Python:
		return "Python"
	case Go:
		return "Go"
	case JavaScript:
		return "JavaScript"
	case Java:
		return "Java"
	case C:
		return "C"
	case Cpp:
		return "C++"
	case Rust:
		return "Rust"
	case Ruby:
		return "Ruby"
	case PHP:
		return "PHP"
	case Perl:
		return "Perl"
	case Lua:
		return "Lua"
	case Zig:
		return "Zig"
	case Nim:
		return "Nim"
	case Crystal:
		return "Crystal"
	case D:
		return "D"
	case Shell:
		return "Shell"
	case Haskell:
		return "Haskell"
	case OCaml:
		return "OCaml"
	case Elixir:
		return "Elixir"
	case Erlang:
		return "Erlang"
	case FSharp:
		return "F#"
	case Kotlin:
		return "Kotlin"
	case Scala:
		return "Scala"
	case Groovy:
		return "Groovy"
	case Swift:
		return "Swift"
	case Dart:
		return "Dart"
	case TypeScript:
		return "TypeScript"
	case R:
		return "R"
	case Julia:
		return "Julia"
	case Brainfuck:
		return "Brainfuck"
	default:
		return "Unknown"
	}
}

// LanguageHandler defines the interface for language-specific operations
type LanguageHandler interface {
	Name() string
	Extensions() []string
	Type() LanguageType
	Compile(source string, output string) error
	Run(file string, args []string) ([]byte, error)
	NeedsCompilation() bool
}
