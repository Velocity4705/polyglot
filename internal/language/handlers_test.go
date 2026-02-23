package language

import (
	"testing"

	"github.com/Velocity4705/polyglot/pkg/types"
)

// TestAllHandlers tests all language handlers
func TestAllHandlers(t *testing.T) {
	tests := []struct {
		name             string
		handler          types.LanguageHandler
		expectedName     string
		expectedExts     []string
		expectedType     types.LanguageType
		needsCompilation bool
	}{
		// Interpreted languages
		{"Python", &PythonHandler{}, "Python", []string{".py"}, types.Interpreted, false},
		{"JavaScript", &JavaScriptHandler{}, "JavaScript", []string{".js"}, types.Interpreted, false},
		{"Ruby", &RubyHandler{}, "Ruby", []string{".rb"}, types.Interpreted, false},
		{"PHP", &PHPHandler{}, "PHP", []string{".php"}, types.Interpreted, false},
		{"Perl", &PerlHandler{}, "Perl", []string{".pl"}, types.Interpreted, false},
		{"Lua", &LuaHandler{}, "Lua", []string{".lua"}, types.Interpreted, false},
		{"Shell", &ShellHandler{}, "Shell", []string{".sh", ".bash"}, types.Interpreted, false},

		// Compiled languages
		{"Go", &GoHandler{}, "Go", []string{".go"}, types.Compiled, false}, // go run
		{"Java", &JavaHandler{}, "Java", []string{".java"}, types.Compiled, true},
		{"C", &CHandler{}, "C", []string{".c"}, types.Compiled, true},
		{"C++", &CppHandler{}, "C++", []string{".cpp", ".cc", ".cxx", ".C"}, types.Compiled, true},
		{"Rust", &RustHandler{}, "Rust", []string{".rs"}, types.Compiled, true},
		{"Zig", &ZigHandler{}, "Zig", []string{".zig"}, types.Compiled, false}, // zig run
		{"Nim", &NimHandler{}, "Nim", []string{".nim"}, types.Compiled, true},
		{"Crystal", &CrystalHandler{}, "Crystal", []string{".cr"}, types.Compiled, false}, // crystal run
		{"D", &DHandler{}, "D", []string{".d"}, types.Compiled, true},

		// Functional languages
		{"Haskell", &HaskellHandler{}, "Haskell", []string{".hs"}, types.Compiled, true},
		{"OCaml", &OCamlHandler{}, "OCaml", []string{".ml", ".mli"}, types.Compiled, true},
		{"Elixir", &ElixirHandler{}, "Elixir", []string{".ex", ".exs"}, types.Interpreted, false},
		{"Erlang", &ErlangHandler{}, "Erlang", []string{".erl"}, types.Compiled, true},
		{"F#", &FSharpHandler{}, "F#", []string{".fs", ".fsx"}, types.Interpreted, false},

		// JVM languages
		{"Kotlin", &KotlinHandler{}, "Kotlin", []string{".kt", ".kts"}, types.Compiled, false},
		{"Scala", &ScalaHandler{}, "Scala", []string{".scala", ".sc"}, types.Compiled, false},
		{"Groovy", &GroovyHandler{}, "Groovy", []string{".groovy", ".gvy"}, types.Interpreted, false},

		// Modern languages
		{"Swift", &SwiftHandler{}, "Swift", []string{".swift"}, types.Compiled, false},
		{"Dart", &DartHandler{}, "Dart", []string{".dart"}, types.Compiled, false},
		{"TypeScript", &TypeScriptHandler{}, "TypeScript", []string{".ts"}, types.Compiled, false},

		// Scientific languages
		{"R", &RHandler{}, "R", []string{".r", ".R"}, types.Interpreted, false},
		{"Julia", &JuliaHandler{}, "Julia", []string{".jl"}, types.Interpreted, false},

		// Esoteric languages
		{"Brainfuck", &BrainfuckHandler{}, "Brainfuck", []string{".bf", ".b"}, types.Interpreted, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Name
			if tt.handler.Name() != tt.expectedName {
				t.Errorf("Name: expected '%s', got '%s'", tt.expectedName, tt.handler.Name())
			}

			// Test Extensions
			exts := tt.handler.Extensions()
			if len(exts) != len(tt.expectedExts) {
				t.Errorf("Extensions: expected %d, got %d", len(tt.expectedExts), len(exts))
			}
			for i, ext := range tt.expectedExts {
				if i >= len(exts) || exts[i] != ext {
					t.Errorf("Extension[%d]: expected '%s', got '%s'", i, ext, exts[i])
				}
			}

			// Test Type
			if tt.handler.Type() != tt.expectedType {
				t.Errorf("Type: expected %v, got %v", tt.expectedType, tt.handler.Type())
			}

			// Test NeedsCompilation
			if tt.handler.NeedsCompilation() != tt.needsCompilation {
				t.Errorf("NeedsCompilation: expected %v, got %v", tt.needsCompilation, tt.handler.NeedsCompilation())
			}
		})
	}
}

// TestGetHandler tests the GetHandler function
func TestGetHandler(t *testing.T) {
	tests := []struct {
		lang     types.Language
		expected string
	}{
		{types.Python, "Python"},
		{types.Go, "Go"},
		{types.JavaScript, "JavaScript"},
		{types.Java, "Java"},
		{types.C, "C"},
		{types.Cpp, "C++"},
		{types.Rust, "Rust"},
		{types.Ruby, "Ruby"},
		{types.PHP, "PHP"},
		{types.Unknown, ""},
	}

	for _, tt := range tests {
		t.Run(tt.lang.Name(), func(t *testing.T) {
			handler := GetHandler(tt.lang)
			if tt.expected == "" {
				if handler != nil {
					t.Errorf("Expected nil handler for %s, got %v", tt.lang.Name(), handler)
				}
			} else {
				if handler == nil {
					t.Errorf("Expected handler for %s, got nil", tt.lang.Name())
				} else if handler.Name() != tt.expected {
					t.Errorf("Expected handler name '%s', got '%s'", tt.expected, handler.Name())
				}
			}
		})
	}
}
