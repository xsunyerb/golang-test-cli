package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TestCLI",
	Long:  `All software has versions. This is TestCLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TestCLI v0.1 -- HEAD")
	},
}
