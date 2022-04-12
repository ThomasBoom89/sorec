package sonos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func GetRefurbishedProductNames(locale string) (error, []string) {
	url := "https://www.sonos.com/" + locale + "/shop/certified-refurbished"
	client := http.Client{}

	response, err := client.Get(url)
	if err != nil {
		return err, nil
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err, nil
	}
	re, err := regexp.Compile(url)
	valid := re.Find(data)
	if valid == nil {
		return err, nil
	}
	productExpression := fmt.Sprintf("/%s/shop/([^\\s\"]*)\"\\w*><", locale)
	re, err = regexp.Compile(productExpression)
	if err != nil {
		return err, nil
	}
	matches := re.FindAllStringSubmatch(string(data), -1)
	var productNames []string
	for _, match := range matches {
		if match[1] != "" {
			productNames = append(productNames, match[1])
		}
	}

	return nil, productNames
}
