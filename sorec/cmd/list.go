package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list has a number of different subcommands",
	Long:  `list has a number of different subcommands`,
}
