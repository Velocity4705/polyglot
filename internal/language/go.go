package language

import (
	"os/exec"

	"github.com/Velocity4705/polyglot/pkg/types"
)

type GoHandler struct{}

func (h *GoHandler) Name() string {
	return "Go"
}

func (h *GoHandler) Extensions() []string {
	return []string{".go"}
}

func (h *GoHandler) Type() types.LanguageType {
	return types.Compiled
}

func (h *GoHandler) NeedsCompilation() bool {
	// Go can use 'go run' which compiles and runs in one step
	return false
}

func (h *GoHandler) Compile(source string, output string) error {
	cmd := exec.Command("go", "build", "-o", output, source)
	return cmd.Run()
}

func (h *GoHandler) Run(file string, args []string) ([]byte, error) {
	cmdArgs := append([]string{"run", file}, args...)
	cmd := exec.Command("go", cmdArgs...)
	return cmd.CombinedOutput()
}
