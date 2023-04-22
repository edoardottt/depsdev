/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command when called with deps subcommand.
var depsCmd = &cobra.Command{
	Use:   "deps package-manager package-name version",
	Short: "Get info about a package's dependencies",
	Long:  `Get information about a resolved dependency graph for the given package version.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return fmt.Errorf("%s %w", "three", input.ErrArgumentsLeast)
		}
		if !input.IsValidPackageManager(args[0]) {
			return input.ErrInvalidPackageManager
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		d, err := depsdev.DepsHandler(args)
		if err != nil {
			log.Fatal(err)
		}

		dJSON, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(string(dJSON))
	},
}
