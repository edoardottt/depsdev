/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package cmd

import (
	"fmt"
	"log"

	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command when called with deps subcommand.
var depsCmd = &cobra.Command{
	Use:   "deps package-manager package-name version",
	Short: "Get info about a package's dependencies",
	Long:  `Get information about a resolved dependency graph for the given package version.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < minArgsThree {
			return fmt.Errorf("%s %w", "three", input.ErrArgumentsLeast)
		}
		if !input.IsValidPackageManager(args[0], input.DepsValidPackageManagers) {
			return input.ErrInvalidPackageManager
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		d, err := api.GetDependencies(args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		dJSON, err := output.IndentJSON(d)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(dJSON)
	},
}
