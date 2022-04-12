package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sonos-refurbished-checker/sonos"
	"sonos-refurbished-checker/utils"
)

var localeCmd = &cobra.Command{
	Use:   "locale",
	Short: "list available locales",
	Long:  `list available locales provided by sonos store`,
	Run: func(cmd *cobra.Command, args []string) {
		err, locales := sonos.GetLocale()
		if err != nil {
			fmt.Println("could not fetch locales")
			os.Exit(1)
		}
		var rows [][]string
		for _, locale := range locales {
			rows = append(rows, []string{locale.CountryLangName, locale.PathCode})
		}
		utils.BuildTable([]string{"COUNTRY", "LOCALE"}, rows)
	},
}
