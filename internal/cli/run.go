package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/polyglot/internal/detector"
	"github.com/yourusername/polyglot/internal/executor"
	"github.com/yourusername/polyglot/internal/language"
)

var (
	verbose bool
	quiet   bool
	args    []string
)

var runCmd = &cobra.Command{
	Use:   "run <file>",
	Short: "Detect and run a source file",
	Args:  cobra.ExactArgs(1),
	RunE:  runFile,
}

func init() {
	runCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	runCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode (suppress polyglot messages)")
	runCmd.Flags().StringSliceVar(&args, "args", []string{}, "Arguments to pass to the program")
}

func runFile(cmd *cobra.Command, cmdArgs []string) error {
	filename := cmdArgs[0]

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filename)
	}

	// Detect language
	lang, err := detector.DetectLanguage(filename)
	if err != nil {
		return err
	}

	if !quiet {
		fmt.Printf("Detected: %s\n", lang.Name())
	}

	// Get language handler
	handler := language.GetHandler(lang)
	if handler == nil {
		return fmt.Errorf("no handler found for language: %s", lang)
	}

	// Execute
	exec := executor.New(verbose, quiet)
	return exec.Run(handler, filename, args)
}
