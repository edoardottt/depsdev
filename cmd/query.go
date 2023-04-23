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

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command when called with query subcommand.
var queryCmd = &cobra.Command{
	Use:   "query \"query-value\"",
	Short: "Get info about multiple package versions using a query",
	Long: `Get information about multiple package versions, which can be specified by name, content hash, or both.
It is typical for hash queries to return many results; 
hashes are matched against multiple artifacts that comprise package versions, 
and any given artifact may appear in many package versions.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("%s %w", "one", input.ErrArgumentLeast)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		p, err := depsdev.QueryHandler(args)
		if err != nil {
			log.Fatal(err)
		}

		pJSON, err := output.IndentJSON(p)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(pJSON)
	},
}
