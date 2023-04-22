/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"os"

	"github.com/edoardottt/depsdev/pkg/output"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "depsdev",
	Short: "A brief description of your application",
	Long:  output.Banner,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(depsCmd)
	rootCmd.AddCommand(advisoryCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(queryCmd)
}
