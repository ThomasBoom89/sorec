package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sonos-refurbished-checker/sonos"
	"sonos-refurbished-checker/utils"
	"strconv"
	"strings"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check current prices and availabilities from the sonos refurbished store.",
	Long: `check current prices and availabilities from the sonos refurbished store.
can be set to the respective country store depending on provided locale. 
see also command "list locale" for available locale.`,
	Run: func(cmd *cobra.Command, args []string) {
		var productNames []string
		var err error
		if products == "" {
			err, productNames = sonos.GetRefurbishedProductNames(locale)
			if err != nil {
				fmt.Println("Could not fetch product names")
				os.Exit(1)
			}
		} else {
			productNames = strings.Split(products, ",")
		}
		if productNames == nil {
			fmt.Println("No products listed")
			os.Exit(0)
		}
		err, additionalPrices := sonos.GetAdditionalPrices(locale, productNames)
		if err != nil {
			fmt.Println("Could not fetch additional prices")
			os.Exit(1)
		}
		additionalPriceMap := buildPriceMap(additionalPrices)
		err, responses := sonos.GetProducts(locale, productNames)
		if err != nil {
			fmt.Println("Could not fetch product data")
			os.Exit(1)
		}
		fmt.Println("")
		var rows [][]string
		for requestedName, response := range responses {

			data := response.Data.GetCommerce.Products.Data[0]
			if len(data.Id) == 0 {
				fmt.Println(requestedName, "no answer!")
				continue
			}
			if len(data.Variants) == 0 {
				rows = append(rows, []string{
					data.Content[0].Name,
					"",
					additionalPriceMap[data.Id],
					strconv.Itoa(data.Inventory.StockLevel),
				})
				continue
			}
			colorMap := buildColorMap(data.VariationAttributes)
			for _, variant := range data.Variants {
				color := colorMap[variant.VariationValues.Color]
				rows = append(rows, []string{
					data.Content[0].Name,
					color,
					additionalPriceMap[data.Id],
					strconv.Itoa(variant.Product.Data[0].Inventory.StockLevel),
				})
			}

		}
		utils.BuildTable([]string{"NAME", "COLOR", "PRICE", "STOCK"}, rows)
	},
}

func buildColorMap(variationAttributes []sonos.VariationAttribute) map[string]string {
	colorMap := make(map[string]string)
	for _, variationAttribute := range variationAttributes {
		if variationAttribute.Id == "color" {
			for _, value := range variationAttribute.Values {
				colorMap[value.Value] = value.Name
			}
		}
	}
	return colorMap
}

func buildPriceMap(additionalPrices []sonos.AdditionalPrice) map[string]string {
	priceMap := make(map[string]string)
	for _, additionalPrice := range additionalPrices {
		if additionalPrice.PriceStr == "" {
			priceMap[additionalPrice.Id] = additionalPrice.Variants[0].PriceStr
		} else {
			priceMap[additionalPrice.Id] = additionalPrice.PriceStr
		}
	}

	return priceMap
}
