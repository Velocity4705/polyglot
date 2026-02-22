package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/yourusername/polyglot/internal/config"
	"github.com/yourusername/polyglot/internal/ui"
	"github.com/yourusername/polyglot/pkg/types"
)

type ExecutionError struct {
	Message  string
	ExitCode int
}

func (e *ExecutionError) Error() string {
	return e.Message
}

type Executor struct {
	verbose bool
	quiet   bool
	config  *config.Config
}

func New(verbose, quiet bool) *Executor {
	return &Executor{
		verbose: verbose,
		quiet:   quiet,
	}
}

func NewWithConfig(verbose, quiet bool, cfg *config.Config) *Executor {
	return &Executor{
		verbose: verbose,
		quiet:   quiet,
		config:  cfg,
	}
}

func (e *Executor) Run(handler types.LanguageHandler, filename string, args []string) error {
	// Check if language is enabled
	if e.config != nil && !e.config.IsLanguageEnabled(string(handler.Name())) {
		return fmt.Errorf("language %s is disabled in configuration", handler.Name())
	}

	if handler.NeedsCompilation() {
		// Compile first
		output := e.getOutputName(filename)

		if !e.quiet {
			ui.Step("Compiling %s...", ui.File(filename))
		}

		spinner := ui.NewSpinner(fmt.Sprintf("Compiling %s", handler.Name()))
		if !e.quiet && !e.verbose {
			spinner.Start()
		}

		err := handler.Compile(filename, output)

		if !e.quiet && !e.verbose {
			spinner.Stop()
		}

		if err != nil {
			if !e.quiet {
				ui.Error("Compilation failed")
			}
			return err
		}

		if !e.quiet {
			ui.Success("Compilation successful")
		}

		// Clean up compiled binary after execution
		defer func() {
			os.Remove(output)
			// Also clean up Java .class files
			if handler.Name() == "Java" {
				dir := filepath.Dir(filename)
				className := strings.TrimSuffix(filepath.Base(filename), ".java")
				classFile := filepath.Join(dir, className+".class")
				os.Remove(classFile)
			}
		}()
	}

	if !e.quiet {
		cmdStr := e.getCommand(handler, filename)
		if len(args) > 0 {
			cmdStr += " " + strings.Join(args, " ")
		}
		ui.Step("Executing: %s", ui.Command(cmdStr))
	}

	output, err := handler.Run(filename, args)

	// Print output (both stdout and stderr)
	if len(output) > 0 {
		fmt.Print(string(output))
	}

	if err != nil {
		// Check if it's an exit error to preserve exit code
		if exitErr, ok := err.(*exec.ExitError); ok {
			return &ExecutionError{
				Message:  "program exited with error",
				ExitCode: exitErr.ExitCode(),
			}
		}
		return fmt.Errorf("execution failed: %w", err)
	}

	return nil
}

func (e *Executor) Compile(handler types.LanguageHandler, filename string) error {
	if !handler.NeedsCompilation() {
		if !e.quiet {
			fmt.Println("No compilation needed")
		}
		return nil
	}

	output := e.getOutputName(filename)

	if !e.quiet {
		fmt.Printf("Compiling: %s %s -o %s\n", handler.Name(), filename, output)
	}

	if err := handler.Compile(filename, output); err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	if !e.quiet {
		fmt.Printf("Compilation successful: %s\n", output)
	}

	return nil
}

func (e *Executor) getOutputName(filename string) string {
	ext := filepath.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}

func (e *Executor) getCommand(handler types.LanguageHandler, filename string) string {
	switch handler.Name() {
	case "Python":
		return "python3"
	case "Go":
		return "go run"
	case "JavaScript":
		return "node"
	default:
		return handler.Name()
	}
}
