package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pswd",
	Short: "pswd - a CLI to store your passwords",
	Long:  `Blah-blah`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not run pswd", err)
		os.Exit(1)
	}
}
