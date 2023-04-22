package cmd

import (
	"fmt"

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command when called with deps subcommand.
var depsCmd = &cobra.Command{
	Use:   "deps package-manager package-name version",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("%s %w", "one", input.ErrArgumentLeast)
		}
		if !input.IsValidPackageManager(args[0]) {
			return input.ErrInvalidPackageManager
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		depsdev.DepsHandler(args)
	},
}
