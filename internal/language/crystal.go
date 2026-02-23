package language

import (
	"os/exec"

	"github.com/Velocity4705/polyglot/pkg/types"
)

type CrystalHandler struct{}

func (h *CrystalHandler) Name() string {
	return "Crystal"
}

func (h *CrystalHandler) Extensions() []string {
	return []string{".cr"}
}

func (h *CrystalHandler) Type() types.LanguageType {
	return types.Compiled
}

func (h *CrystalHandler) NeedsCompilation() bool {
	// Crystal can use 'crystal run' which compiles and runs
	return false
}

func (h *CrystalHandler) Compile(source string, output string) error {
	cmd := exec.Command("crystal", "build", source, "-o", output)
	return cmd.Run()
}

func (h *CrystalHandler) Run(file string, args []string) ([]byte, error) {
	cmdArgs := append([]string{"run", file}, args...)
	cmd := exec.Command("crystal", cmdArgs...)
	return cmd.CombinedOutput()
}
