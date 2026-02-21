package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.2.0"
const milestone = "Milestone 2 & 3 (Compiled + Extended Languages)"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Polyglot v%s\n", version)
		fmt.Printf("%s\n", milestone)
		fmt.Println("Universal compiler and interpreter wrapper")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
