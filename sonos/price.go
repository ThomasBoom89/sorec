package sonos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type AdditionalPrice struct {
	Id             string  `json:"id"`
	CurrencyCode   string  `json:"currencyCode,omitempty"`
	CurrencySymbol string  `json:"currencySymbol,omitempty"`
	Price          float64 `json:"price,omitempty"`
	PriceStr       string  `json:"priceStr,omitempty"`
	Master         bool    `json:"master"`
	Variant        bool    `json:"variant"`
	Simple         bool    `json:"simple"`
	Variants       []struct {
		Id              string `json:"id"`
		CurrencyCode    string `json:"currencyCode"`
		CurrencySymbol  string `json:"currencySymbol"`
		Price           int    `json:"price"`
		PriceStr        string `json:"priceStr"`
		PromoPrice      int    `json:"promoPrice,omitempty"`
		PromoPriceStr   string `json:"promoPriceStr,omitempty"`
		PromoCalloutMsg string `json:"promoCalloutMsg,omitempty"`
		PromoPhrase     string `json:"promoPhrase,omitempty"`
	} `json:"variants,omitempty"`
}

func GetAdditionalPrices(locale string, products []string) (error, []AdditionalPrice) {
	url := "https://www.sonos.com/" + locale + "/getproductpricejson?pid=" + strings.Join(products, ",")
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return err, nil
	}

	body, _ := ioutil.ReadAll(response.Body)
	var additionalPrices []AdditionalPrice
	err = json.Unmarshal(body, &additionalPrices)
	if err != nil {
		return err, nil
	}

	return nil, additionalPrices
}
