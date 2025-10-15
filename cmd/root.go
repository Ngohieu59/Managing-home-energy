package cmd

import (
	"Managing-home-energy/cmd/api"
	"Managing-home-energy/cmd/migrate"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Managing-home-energy",
	Short: "Managing-home-energy",
	Long:  "Managing-home-energy",
}

func init() {
	rootCmd.AddCommand(api.Cmd)
	rootCmd.AddCommand(migrate.Cmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {

		os.Exit(1)
	}
}
