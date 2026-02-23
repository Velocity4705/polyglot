package language

import (
	"testing"

	"github.com/Velocity4705/polyglot/pkg/types"
)

func TestPythonHandler(t *testing.T) {
	handler := &PythonHandler{}

	t.Run("Name", func(t *testing.T) {
		if handler.Name() != "Python" {
			t.Errorf("Expected name 'Python', got '%s'", handler.Name())
		}
	})

	t.Run("Extensions", func(t *testing.T) {
		exts := handler.Extensions()
		if len(exts) != 1 {
			t.Errorf("Expected 1 extension, got %d", len(exts))
		}
		if exts[0] != ".py" {
			t.Errorf("Expected extension '.py', got '%s'", exts[0])
		}
	})

	t.Run("Type", func(t *testing.T) {
		if handler.Type() != types.Interpreted {
			t.Errorf("Expected type Interpreted, got %v", handler.Type())
		}
	})

	t.Run("NeedsCompilation", func(t *testing.T) {
		if handler.NeedsCompilation() {
			t.Error("Python should not need compilation")
		}
	})
}
