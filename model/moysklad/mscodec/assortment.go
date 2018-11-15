package mscodec

type Assortment struct {
	Meta               Meta            `json:"meta"`
	ID                 string          `json:"id"`
	AccountID          string          `json:"accountId"`
	Owner              Owner           `json:"owner,omitempty"`
	Shared             bool            `json:"shared,omitempty"`
	Group              Group           `json:"group,omitempty"`
	Version            int             `json:"version"`
	Updated            string          `json:"updated"`
	Name               string          `json:"name"`
	Code               string          `json:"code"`
	ExternalCode       string          `json:"externalCode"`
	Archived           bool            `json:"archived"`
	PathName           string          `json:"pathName,omitempty"`
	Uom                Uom             `json:"uom,omitempty"`
	MinPrice           float64         `json:"minPrice,omitempty"`
	SalePrices         SalePrices      `json:"salePrices"`
	BuyPrice           BuyPrice        `json:"buyPrice,omitempty"`
	Article            string          `json:"article,omitempty"`
	Weight             float64         `json:"weight,omitempty"`
	Volume             float64         `json:"volume,omitempty"`
	Barcodes           []string        `json:"barcodes,omitempty"`
	ModificationsCount int             `json:"modificationsCount,omitempty"`
	IsSerialTrackable  bool            `json:"isSerialTrackable,omitempty"`
	Stock              float64         `json:"stock,omitempty"`
	Reserve            float64         `json:"reserve,omitempty"`
	InTransit          float64         `json:"inTransit,omitempty"`
	Quantity           float64         `json:"quantity,omitempty"`
	Characteristics    Characteristics `json:"characteristics,omitempty"`
	Product            struct {
		Meta Meta `json:"meta"`
	} `json:"product,omitempty"`
}
