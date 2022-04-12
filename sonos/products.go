package sonos

import (
	"context"
	"fmt"
	"github.com/getoutreach/goql"
	"time"
)

type Inventory struct {
	Backorderable bool   `goql:"keep"`
	InStockDate   string `goql:"keep"`
	Orderable     bool   `goql:"keep"`
	Preorderable  bool   `goql:"keep"`
	Ats           int    `goql:"keep"`
	StockLevel    int    `goql:"keep"`
	Perpetual     bool   `goql:"keep"`
}

type Content struct {
	Name       string `goql:"keep"`
	Overview   string `goql:"keep"`
	Descriptor string `goql:"keep"`
}

type VariationValue struct {
	Color string `goql:"keep"`
}

type VariantProductData struct {
	Id                string           `goql:"keep"`
	SKU               string           `goql:"keep,SKU"`
	VPN               string           `goql:"keep,VPN"`
	Gtin              string           `goql:"keep"`
	Price             int              `goql:"keep"`
	ProductPromotions ProductPromotion `goql:"keep"`
	Inventory         Inventory        `goql:"keep"`
}

type VariantProduct struct {
	Data []VariantProductData `goql:"keep"`
}

type Variant struct {
	VariationValues VariationValue `goql:"keep"`
	ProductId       string         `goql:"keep"`
	Product         VariantProduct `goql:"keep"`
}

type ProductPromotion struct {
	PromotionalPrice string `goql:"keep"`
}

type VariationAttributeValue struct {
	Name        string `goql:"keep"`
	Value       string `goql:"keep"`
	Description string `goql:"keep"`
	Orderable   bool   `goql:"keep"`
}

type VariationAttribute struct {
	Id     string                    `goql:"keep"`
	Name   string                    `goql:"keep"`
	Values []VariationAttributeValue `goql:"keep"`
}

type RequestProductData struct {
	Id                  string               `goql:"keep"`
	Price               int                  `goql:"keep"`
	Inventory           Inventory            `goql:"keep"`
	Content             []Content            `goql:"keep"`
	ProductPromotions   ProductPromotion     `goql:"keep"`
	VariationAttributes []VariationAttribute `goql:"keep"`
	Variants            []Variant            `goql:"keep"`
}

type RequestProduct struct {
	Data []RequestProductData `goql:"keep"`
}
type Request struct {
	GetCommerce struct {
		Products RequestProduct `goql:"keep"`
	} `goql:"getCommerce(locale:$locale<String>,pids:$pids<String>)"`
}

type ProductResponse struct {
	Data struct {
		GetCommerce struct {
			Products RequestProduct
		}
	}
}

func GetProducts(locale string, products []string) (error, map[string]ProductResponse) {
	var bal = Request{}
	var fields = goql.Fields{}
	query, err := goql.MarshalQuery(bal, fields)
	if err != nil {
		return err, nil
	}
	client := goql.NewClient("https://www.sonos.com/api/graphql", goql.ClientOptions{})
	var ctx = context.Background()
	variables := make(map[string]interface{})
	variables["locale"] = locale
	responses := make(map[string]ProductResponse)
	for _, product := range products {
		fmt.Print("Checking: ", product)
		time.Sleep(1500 * time.Millisecond)
		variables["pids"] = product
		var resp = ProductResponse{}
		err = client.CustomOperation(ctx, query, variables, &resp)
		if err != nil {
			fmt.Println("\u271C")
			continue
		}
		responses[product] = resp
		fmt.Println("\u2714")
	}

	return nil, responses
}
