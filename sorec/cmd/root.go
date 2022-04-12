package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var locale string
var products string

func init() {
	rootCmd.AddCommand(listCmd, checkCmd, versionCmd)
	listCmd.AddCommand(productsCmd, localeCmd)
	productsCmd.PersistentFlags().StringVarP(&locale, "locale", "l", "en-us", "locale of sonos shop, use command list locale for supported locale")
	checkCmd.PersistentFlags().StringVarP(&locale, "locale", "l", "en-us", "locale of sonos shop, use command list locale for supported locale")
	checkCmd.PersistentFlags().StringVarP(&products, "products", "p", "", "comma seperated list of products you want to request, defaults to all products available in sonos shop by locale")
}

var rootCmd = &cobra.Command{
	Use:   "sorec",
	Short: "sorec is a CLI-Tool to get prices and availabilities for refurbished sonos products",
	Long: `sorec is a CLI-Tool to get prices and availabilities for refurbished sonos products.
sorec is built with love by spf13 and friends in Go.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
