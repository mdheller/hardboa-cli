package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show hba-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootCmd.Use + " v" + VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
