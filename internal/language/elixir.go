package language

import (
	"os/exec"

	"github.com/yourusername/polyglot/pkg/types"
)

type ElixirHandler struct{}

func (h *ElixirHandler) Name() string {
	return "Elixir"
}

func (h *ElixirHandler) Extensions() []string {
	return []string{".ex", ".exs"}
}

func (h *ElixirHandler) Type() types.LanguageType {
	return types.Interpreted
}

func (h *ElixirHandler) NeedsCompilation() bool {
	return false
}

func (h *ElixirHandler) Compile(source string, output string) error {
	return nil
}

func (h *ElixirHandler) Run(file string, args []string) ([]byte, error) {
	cmdArgs := append([]string{file}, args...)
	cmd := exec.Command("elixir", cmdArgs...)
	return cmd.CombinedOutput()
}
