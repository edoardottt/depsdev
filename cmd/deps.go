package cmd

import (
	"errors"
	"fmt"

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command when called with deps subcommand
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
			return errors.New("requires at least one argument")
		}
		if input.IsValidPackageManager(args[0]) {
			return nil
		}

		if input.IsValidPackageManager(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid package name specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		depsdev.DepsHandler(args)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
