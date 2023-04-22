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

// advisoryCmd represents the advisory command when called with advisory subcommand.
var advisoryCmd = &cobra.Command{
	Use:   "advisory advisory-code",
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

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		depsdev.AdvisoryHandler(args)
	},
}
