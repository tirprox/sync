package moysklad

const BASE_PRODUCT = "https://online.moysklad.ru/api/remap/1.1/entity/product"

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
	MinPrice   int `json:"minPrice"`
	SalePrices []struct {
		Value    int `json:"value"`
		Currency struct {
			Meta Meta `json:"meta"`
		} `json:"currency"`
		PriceType string `json:"priceType"`
	} `json:"salePrices"`
	BuyPrice struct {
		Value    int `json:"value"`
		Currency struct {
			Meta Meta `json:"meta"`
		} `json:"currency"`
	} `json:"buyPrice"`
	Article            string   `json:"article"`
	Weight             int      `json:"weight"`
	Volume             int      `json:"volume"`
	Barcodes           []string `json:"barcodes"`
	ModificationsCount int      `json:"modificationsCount"`
	IsSerialTrackable  bool     `json:"isSerialTrackable"`
	Stock              int      `json:"stock"`
	Reserve            int      `json:"reserve"`
	InTransit          int      `json:"inTransit"`
	Quantity           int      `json:"quantity"`
}

type ProductWrapper struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Product  Product           `json:"productFolder"`
	Stock    map[string]string `json:"stock"`
	Variants []VariantWrapper  `json:"products"`
	Other    struct {
	} `json:"other"`
}

type MegaProduct struct {
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
	MinPrice   int `json:"minPrice"`
	SalePrices []struct {
		Value    int `json:"value"`
		Currency struct {
			Meta Meta `json:"meta"`
		} `json:"currency"`
		PriceType string `json:"priceType"`
	} `json:"salePrices"`
	BuyPrice struct {
		Value    int `json:"value"`
		Currency struct {
			Meta Meta `json:"meta"`
		} `json:"currency"`
	} `json:"buyPrice"`
	Article            string           `json:"article"`
	Weight             int              `json:"weight"`
	Volume             int              `json:"volume"`
	Barcodes           []string         `json:"barcodes"`
	ModificationsCount int              `json:"modificationsCount"`
	IsSerialTrackable  bool             `json:"isSerialTrackable"`
	Stock              int              `json:"stock"`
	Reserve            int              `json:"reserve"`
	InTransit          int              `json:"inTransit"`
	Quantity           int              `json:"quantity"`
	Variants           []VariantWrapper `json:"variants"`
	Other              struct {
	} `json:"other"`
	Stocks map[string]string `json:"stocks"`
}
