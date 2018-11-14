package moysklad

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/tirprox/sync/model/moysklad/mscodec"
	"net/url"
)

const BASE_ASSORTMENT = "https://online.moysklad.ru/api/remap/1.1/entity/assortment"

type Assortment struct {
	Meta mscodec.Meta `json:"meta"`
	/*ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Owner     mscodec.Owner `json:"owner,omitempty"`
	Shared bool `json:"shared,omitempty"`
	Group  mscodec.Group `json:"group,omitempty"`
	Version      int    `json:"version"`
	Updated      string `json:"updated"`*/
	Name string `json:"name"`
	/*Code         string `json:"code"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	PathName     string `json:"pathName,omitempty"`
	Uom          struct {
		Meta mscodec.Meta `json:"meta"`
	} `json:"uom,omitempty"`
	MinPrice   float64 `json:"minPrice,omitempty"`
	SalePrices []struct {
		Value    float64 `json:"value"`
		Currency struct {
			Meta mscodec.Meta `json:"meta"`
		} `json:"currency"`
		PriceType string `json:"priceType"`
	} `json:"salePrices"`
	BuyPrice struct {
		Value    float64 `json:"value"`
		Currency struct {
			Meta mscodec.Meta `json:"meta"`
		} `json:"currency"`
	} `json:"buyPrice,omitempty"`
	Article            string   `json:"article,omitempty"`
	Weight             float64  `json:"weight,omitempty"`
	Volume             float64  `json:"volume,omitempty"`
	Barcodes           []string `json:"barcodes,omitempty"`
	ModificationsCount int      `json:"modificationsCount,omitempty"`
	IsSerialTrackable  bool     `json:"isSerialTrackable,omitempty"`
	Stock              float64  `json:"stock,omitempty"`
	Reserve            float64  `json:"reserve,omitempty"`
	InTransit          float64  `json:"inTransit,omitempty"`
	Quantity           float64  `json:"quantity,omitempty"`
	Characteristics    []struct {
		Meta  mscodec.Meta   `json:"meta"`
		ID    string `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"characteristics,omitempty"`
	Product struct {
		Meta mscodec.Meta `json:"meta"`
	} `json:"product,omitempty"`*/
}

func DecodeAssortment(responseBody []byte) Assortment {
	assortment := Assortment{}
	json.Unmarshal(responseBody, &assortment)

	return assortment
}

func GetAssortment(folder ProductFolderWrapper, store Store) []mscodec.Assortment {

	u, err := url.Parse(BASE_ASSORTMENT)
	if err != nil {
	}

	query := u.Query()
	query.Set("limit", "100")
	query.Set("filter", "productFolder="+folder.ProductFolder.Meta.Href)
	query.Set("stockstore", store.Meta.Href)

	u.RawQuery = query.Encode()

	data := GetAll(u.String())

	assortments := []mscodec.Assortment{}

	for _, item := range data {
		assortment := mscodec.Assortment{}
		mapstructure.Decode(item, &assortment)
		assortments = append(assortments, assortment)
	}

	return assortments
}
