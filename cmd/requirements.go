package cmd

import (
	"fmt"
	"log"

	"github.com/edoardottt/depsdev/pkg/input"
	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/spf13/cobra"
)

var requirementsCmd = &cobra.Command{
	Use:   "requirements package-manager package-name version",
	Short: "Get info about a package's requirements",
	Long: `Returns the requirements for a given version in a system-specific format. 
	Requirements are currently available for Maven, npm and NuGet.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return fmt.Errorf("%s %w", "three", input.ErrArgumentsLeast)
		}
		if !input.Contains(args[0], []string{"npm", "maven", "nuget"}) {
			return input.ErrInvalidPackageManagerForRequirements
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		d, err := api.GetRequirements(args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		dJSON, err := output.IndentJSON(d)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(dJSON)
	},
}
