package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number of sorec",
	Long:  `all software has versions. This is sorec's'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0")
	},
}
