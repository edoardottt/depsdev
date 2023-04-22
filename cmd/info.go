package cmd

import (
	"errors"
	"fmt"

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/spf13/cobra"
)

var (
	packageName    string
	packageManager string
)

// infoCmd represents the info command when called with info subcommand
var infoCmd = &cobra.Command{
	Use:   "info package-manager package-name [version]",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires at least two arguments")
		}

		if !input.IsValidPackageManager(args[0]) {
			return fmt.Errorf("invalid package manager specified: %s", args[0])
		}

		if !input.IsValidPackageName(args[1]) {
			return fmt.Errorf("invalid package name specified: %s", args[1])
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 2 {
			depsdev.VersionHandler(args)
		} else {
			depsdev.InfoHandler(args)
		}
	},
}
