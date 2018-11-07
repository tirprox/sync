package moysklad

type Variant struct {
	Meta            Meta   `json:"meta"`
	ID              string `json:"id"`
	AccountID       string `json:"accountId"`
	Version         int    `json:"version"`
	Updated         string `json:"updated"`
	Name            string `json:"name"`
	Code            string `json:"code"`
	ExternalCode    string `json:"externalCode"`
	Archived        bool   `json:"archived"`
	Characteristics []struct {
		Meta  Meta   `json:"meta"`
		ID    string `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"characteristics"`
	SalePrices []struct {
		Value    int `json:"value"`
		Currency struct {
			Meta Meta `json:"meta"`
		} `json:"currency"`
		PriceType string `json:"priceType"`
	} `json:"salePrices"`
	Barcodes []string `json:"barcodes"`
	Product  struct {
		Meta Meta `json:"meta"`
	} `json:"product"`
	Stock     int `json:"stock"`
	Reserve   int `json:"reserve"`
	InTransit int `json:"inTransit"`
	Quantity  int `json:"quantity"`
}
