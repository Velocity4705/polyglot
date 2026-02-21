package language

import (
	"os/exec"

	"github.com/yourusername/polyglot/pkg/types"
)

type SwiftHandler struct{}

func (h *SwiftHandler) Name() string {
	return "Swift"
}

func (h *SwiftHandler) Extensions() []string {
	return []string{".swift"}
}

func (h *SwiftHandler) Type() types.LanguageType {
	return types.Compiled
}

func (h *SwiftHandler) NeedsCompilation() bool {
	// Swift can use 'swift run' for quick execution
	return false
}

func (h *SwiftHandler) Compile(source string, output string) error {
	cmd := exec.Command("swiftc", source, "-o", output)
	return cmd.Run()
}

func (h *SwiftHandler) Run(file string, args []string) ([]byte, error) {
	cmdArgs := append([]string{file}, args...)
	cmd := exec.Command("swift", cmdArgs...)
	return cmd.CombinedOutput()
}
