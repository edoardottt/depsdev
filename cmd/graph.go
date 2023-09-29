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
	"log"

	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/spf13/cobra"
)

// depsCmd represents the deps command when called with deps subcommand.
var graphCmd = &cobra.Command{
	Use:   "graph package-manager package-name version",
	Short: "Generate a Graphviz compatible dependencies graph",
	Long:  `Generate a Graphviz compatible dependencies graph for a specific version of a package.`,
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
		d, err := api.GetDependencies(args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		g, err := output.GenerateGraph(d)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(g)
	},
}
