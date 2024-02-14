package cmd

import (
	"fmt"
	"log"

	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/spf13/cobra"
)

var packagesCmd = &cobra.Command{
	Use:   "packages project",
	Short: "Get info about a project's package versions (GitHub, GitLab, or BitBucket)",
	Long: `Returns known mappings between the requested project and package versions. 
	At most 1500 package versions are returned. (GitHub, GitLab, or BitBucket)`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("%s %w", "one", input.ErrArgumentLeast)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		v, err := api.GetPackageVersions(args[0])
		if err != nil {
			log.Fatal(err)
		}

		vJSON, err := output.IndentJSON(v)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(vJSON)
	},
}
