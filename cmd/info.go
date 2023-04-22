/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package cmd

import (
	"fmt"

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command when called with info subcommand.
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
			return fmt.Errorf("%s %w", "two", input.ErrArgumentsLeast)
		}

		if !input.IsValidPackageManager(args[0]) {
			return input.ErrInvalidPackageManager
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
