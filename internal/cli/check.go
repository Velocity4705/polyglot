package cli

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check which language toolchains are installed",
	RunE:  checkToolchains,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func checkToolchains(cmd *cobra.Command, args []string) error {
	fmt.Println("Checking installed toolchains...")
	fmt.Println("================================")

	toolchains := map[string]string{
		"Python":     "python3",
		"Go":         "go",
		"JavaScript": "node",
		"Ruby":       "ruby",
		"PHP":        "php",
		"Java":       "javac",
		"C":          "gcc",
		"C++":        "g++",
		"Rust":       "rustc",
		"Perl":       "perl",
		"Lua":        "lua",
		"Zig":        "zig",
		"Nim":        "nim",
		"Crystal":    "crystal",
		"D":          "dmd",
		"Shell":      "bash",
		"Haskell":    "ghc",
		"OCaml":      "ocamlc",
		"Elixir":     "elixir",
		"Erlang":     "erl",
		"F#":         "dotnet",
		"Kotlin":     "kotlin",
		"Scala":      "scala",
		"Groovy":     "groovy",
		"Swift":      "swift",
		"Dart":       "dart",
		"TypeScript": "ts-node",
		"R":          "Rscript",
		"Julia":      "julia",
		"Brainfuck":  "bf",
	}

	installed := 0
	total := len(toolchains)

	for lang, command := range toolchains {
		_, err := exec.LookPath(command)
		status := "✗ NOT FOUND"
		if err == nil {
			status = "✓ INSTALLED"
			installed++
		}
		fmt.Printf("%-15s %-15s %s\n", lang, command, status)
	}

	fmt.Println("================================")
	fmt.Printf("Summary: %d/%d toolchains installed\n", installed, total)

	if installed < total {
		fmt.Println("\nTo use all languages, install the missing toolchains.")
	}

	return nil
}
