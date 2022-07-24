package cmd

import (
	"pswd/pkg/commands"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add password",
	Long:  `add password`,
	Run: func(_ *cobra.Command, _ []string) {
		commands.Add()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
