package mscodec

type Variant struct {
	Meta            Meta            `json:"meta"`
	ID              string          `json:"id"`
	AccountID       string          `json:"accountId"`
	Version         int             `json:"version"`
	Updated         string          `json:"updated"`
	Name            string          `json:"name"`
	Code            string          `json:"code"`
	ExternalCode    string          `json:"externalCode"`
	Archived        bool            `json:"archived"`
	Characteristics Characteristics `json:"characteristics"`
	SalePrices      SalePrices      `json:"salePrices"`
	Barcodes        []string        `json:"barcodes"`
	Product         struct {
		Meta Meta `json:"meta"`
	} `json:"product"`
	Stock     int `json:"stock"`
	Reserve   int `json:"reserve"`
	InTransit int `json:"inTransit"`
	Quantity  int `json:"quantity"`
}
