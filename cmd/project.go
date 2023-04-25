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

// projectCmd represents the project command when called with project subcommand.
var projectCmd = &cobra.Command{
	Use:   "project project-name",
	Short: "Get info about a project (GitHub, GitLab, or BitBucket)",
	Long:  `Get information about projects hosted by GitHub, GitLab, or BitBucket (if available).`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("%s %w", "one", input.ErrArgumentLeast)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		p, err := depsdev.GetProject(args[0])
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
