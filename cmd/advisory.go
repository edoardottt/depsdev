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

// advisoryCmd represents the advisory command when called with advisory subcommand.
var advisoryCmd = &cobra.Command{
	Use:   "advisory advisory-code",
	Short: "Get info about an (OSV) advisory",
	Long:  `Get information about security advisories hosted by OSV.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("%s %w", "one", input.ErrArgumentLeast)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		a, err := api.GetAdvisory(args[0])
		if err != nil {
			log.Fatal(err)
		}

		aJSON, err := output.IndentJSON(a)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(aJSON)
	},
}
