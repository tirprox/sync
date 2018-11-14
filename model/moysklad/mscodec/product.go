package mscodec

type Product struct {
	Meta         Meta   `json:"meta"`
	ID           string `json:"id"`
	AccountID    string `json:"accountId"`
	Owner        Owner  `json:"owner"`
	Shared       bool   `json:"shared"`
	Group        Group  `json:"group"`
	Version      int    `json:"version"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	PathName     string `json:"pathName"`
	Uom          struct {
		Meta Meta `json:"meta"`
	} `json:"uom"`
	MinPrice   float64 `json:"minPrice"`
	SalePrices []struct {
		Value    float64 `json:"value"`
		Currency struct {
			Meta Meta `json:"meta"`
		} `json:"currency"`
		PriceType string `json:"priceType"`
	} `json:"salePrices"`
	BuyPrice struct {
		Value    float64 `json:"value"`
		Currency struct {
			Meta Meta `json:"meta"`
		} `json:"currency"`
	} `json:"buyPrice"`
	Article            string   `json:"article"`
	Weight             float64  `json:"weight"`
	Volume             float64  `json:"volume"`
	Barcodes           []string `json:"barcodes"`
	ModificationsCount int      `json:"modificationsCount"`
	IsSerialTrackable  bool     `json:"isSerialTrackable"`
	Stock              float64  `json:"stock"`
	Reserve            float64  `json:"reserve"`
	InTransit          float64  `json:"inTransit"`
	Quantity           float64  `json:"quantity"`
}