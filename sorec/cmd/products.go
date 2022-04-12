package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sonos-refurbished-checker/sonos"
	"strings"
)

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "list all available products from the sonos refurbished store.",
	Long: `list all available products from the sonos refurbished store.
can be set to the respective country store depending on provided locale. 
see also command "list locale" for available locale.`,
	Run: func(cmd *cobra.Command, args []string) {
		err, products := sonos.GetRefurbishedProductNames(locale)
		if err != nil {
			fmt.Println("Could not fetch product names")
			os.Exit(1)
		}
		if products == nil {
			fmt.Println("No products listed")
			os.Exit(0)
		}
		fmt.Printf("%s\n", strings.Join(products, "\n"))
	},
}
