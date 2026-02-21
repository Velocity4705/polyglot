package language

import (
	"os/exec"

	"github.com/yourusername/polyglot/pkg/types"
)

type RubyHandler struct{}

func (h *RubyHandler) Name() string {
	return "Ruby"
}

func (h *RubyHandler) Extensions() []string {
	return []string{".rb"}
}

func (h *RubyHandler) Type() types.LanguageType {
	return types.Interpreted
}

func (h *RubyHandler) NeedsCompilation() bool {
	return false
}

func (h *RubyHandler) Compile(source string, output string) error {
	// Ruby doesn't need compilation
	return nil
}

func (h *RubyHandler) Run(file string, args []string) ([]byte, error) {
	cmdArgs := append([]string{file}, args...)
	cmd := exec.Command("ruby", cmdArgs...)
	return cmd.CombinedOutput()
}
