package cmd

import (
	"errors"

	"github.com/edoardottt/depsdev/pkg/depsdev"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command when called with query subcommand
var queryCmd = &cobra.Command{
	Use:   "query \"query-value\"",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		depsdev.QueryHandler(args)
	},
}
