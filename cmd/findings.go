/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package cmd

/*

Not supported for now.
Maybe will be shipped when stable.

// findingsCmd represents the package command when called with findings subcommand.
var findingsCmd = &cobra.Command{
	Use:   "findings package-manager package-name [version]",
	Short: "Get info about a safe dependency management",
	Long:  `Findings evaluates a specified package or version and returns findings which are relevant to safe dependency management.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < minArgsTwo {
			return fmt.Errorf("%s %w", "two", input.ErrArgumentsLeast)
		}

		if !input.IsValidPackageManager(args[0], input.AllValidPackageManagers) {
			return input.ErrInvalidPackageManager
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= minArgsThree {
			v, err := api.GetFindingsVersion(args[0], args[1], args[2])
			if err != nil {
				log.Fatal(err)
			}

			vJSON, err := output.IndentJSON(v)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(vJSON)
		} else {
			p, err := api.GetFindings(args[0], args[1])
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

*/
