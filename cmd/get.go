package cmd

import (
	"pswd/pkg/commands"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get password",
	Long:  `get password`,
	Run: func(_ *cobra.Command, _ []string) {
		commands.Get()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
