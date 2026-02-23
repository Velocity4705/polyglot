package detector

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Velocity4705/polyglot/internal/config"
	"github.com/Velocity4705/polyglot/pkg/types"
)

// DetectLanguage detects the programming language based on file extension
func DetectLanguage(filename string) (types.Language, error) {
	return DetectLanguageWithConfig(filename, nil)
}

// DetectLanguageWithConfig detects language with custom configuration support
func DetectLanguageWithConfig(filename string, cfg *config.Config) (types.Language, error) {
	ext := strings.ToLower(filepath.Ext(filename))

	// Check custom extensions first
	if cfg != nil {
		if customLang := cfg.GetCustomExtension(ext); customLang != "" {
			lang := types.FromString(customLang)
			if lang != types.Unknown {
				return lang, nil
			}
		}
	}

	switch ext {
	case ".py":
		return types.Python, nil
	case ".go":
		return types.Go, nil
	case ".js":
		return types.JavaScript, nil
	case ".java":
		return types.Java, nil
	case ".c":
		return types.C, nil
	case ".cpp", ".cc", ".cxx", ".C":
		return types.Cpp, nil
	case ".rs":
		return types.Rust, nil
	case ".rb":
		return types.Ruby, nil
	case ".php":
		return types.PHP, nil
	case ".pl":
		return types.Perl, nil
	case ".lua":
		return types.Lua, nil
	case ".zig":
		return types.Zig, nil
	case ".nim":
		return types.Nim, nil
	case ".cr":
		return types.Crystal, nil
	case ".d":
		return types.D, nil
	case ".sh", ".bash":
		return types.Shell, nil
	case ".hs":
		return types.Haskell, nil
	case ".ml", ".mli":
		return types.OCaml, nil
	case ".ex", ".exs":
		return types.Elixir, nil
	case ".erl":
		return types.Erlang, nil
	case ".fs", ".fsx":
		return types.FSharp, nil
	case ".kt", ".kts":
		return types.Kotlin, nil
	case ".scala", ".sc":
		return types.Scala, nil
	case ".groovy", ".gvy":
		return types.Groovy, nil
	case ".swift":
		return types.Swift, nil
	case ".dart":
		return types.Dart, nil
	case ".ts":
		return types.TypeScript, nil
	case ".r", ".R":
		return types.R, nil
	case ".jl":
		return types.Julia, nil
	case ".bf", ".b":
		return types.Brainfuck, nil
	default:
		return types.Unknown, fmt.Errorf("unsupported file extension: %s", ext)
	}
}
