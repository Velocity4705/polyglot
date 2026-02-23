package language

import (
	"github.com/Velocity4705/polyglot/pkg/types"
)

var handlers = map[types.Language]types.LanguageHandler{
	types.Python:     &PythonHandler{},
	types.Go:         &GoHandler{},
	types.JavaScript: &JavaScriptHandler{},
	types.Ruby:       &RubyHandler{},
	types.PHP:        &PHPHandler{},
	types.Java:       &JavaHandler{},
	types.C:          &CHandler{},
	types.Cpp:        &CppHandler{},
	types.Rust:       &RustHandler{},
	types.Perl:       &PerlHandler{},
	types.Lua:        &LuaHandler{},
	types.Zig:        &ZigHandler{},
	types.Nim:        &NimHandler{},
	types.Crystal:    &CrystalHandler{},
	types.D:          &DHandler{},
	types.Shell:      &ShellHandler{},
	types.Haskell:    &HaskellHandler{},
	types.OCaml:      &OCamlHandler{},
	types.Elixir:     &ElixirHandler{},
	types.Erlang:     &ErlangHandler{},
	types.FSharp:     &FSharpHandler{},
	types.Kotlin:     &KotlinHandler{},
	types.Scala:      &ScalaHandler{},
	types.Groovy:     &GroovyHandler{},
	types.Swift:      &SwiftHandler{},
	types.Dart:       &DartHandler{},
	types.TypeScript: &TypeScriptHandler{},
	types.R:          &RHandler{},
	types.Julia:      &JuliaHandler{},
	types.Brainfuck:  &BrainfuckHandler{},
}

// GetHandler returns the language handler for a given language
func GetHandler(lang types.Language) types.LanguageHandler {
	return handlers[lang]
}

// GetAllLanguages returns all supported languages
func GetAllLanguages() []types.Language {
	langs := make([]types.Language, 0, len(handlers))
	for lang := range handlers {
		langs = append(langs, lang)
	}
	return langs
}
