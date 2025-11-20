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

// infoCmd represents the info command when called with info subcommand.
var infoCmd = &cobra.Command{
	Use:   "info package-manager package-name [version]",
	Short: "Get info about a package or a specific version of that",
	Long: `Get information about a package, including a list of its available versions, 
with the default version marked if known.
If version is also specified, returns information about a specific package version 
including its licenses and any security advisories known to affect it.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < minArgsTwo {
			return fmt.Errorf("%s %w", "two", input.ErrArgumentsLeast)
		}

		if !input.IsValidPackageManager(args[0]) {
			return input.ErrInvalidPackageManager
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= minArgsThree {
			v, err := api.GetVersion(args[0], args[1], args[2])
			if err != nil {
				log.Fatal(err)
			}

			vJSON, err := output.IndentJSON(v)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(vJSON)
		} else {
			p, err := api.GetInfo(args[0], args[1])
			if err != nil {
				log.Fatal(err)
			}

			pJSON, err := output.IndentJSON(p)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(pJSON)
		}
	},
}
