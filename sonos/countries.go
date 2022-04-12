package sonos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Locale struct {
	CountryLangName string
	Currency        string
	PathCode        string
}

type localeResponse struct {
	CountryName string
	Locales     []Locale
}

func GetLocale() (error, []Locale) {
	url := "https://www.sonos.com/api/countries"
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return err, nil
	}
	body, _ := ioutil.ReadAll(response.Body)
	var localeResponses []localeResponse
	err = json.Unmarshal(body, &localeResponses)
	if err != nil {
		return err, nil
	}
	var locales []Locale
	for _, localeResponse := range localeResponses {
		locales = append(locales, localeResponse.Locales...)
	}

	return nil, locales
}
