package mscodec

type Meta struct {
	Href         string `json:"href,omitempty"`
	MetadataHref string `json:"metadataHref,omitempty"`
	Type         string `json:"type,omitempty"`
	MediaType    string `json:"mediaType,omitempty"`
	Size         int    `json:"size,omitempty" bson:"size,omitempty"`
	Limit        int    `json:"limit,omitempty" bson:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty" bson:"offset,omitempty"`
	NextHref     string `json:"nextHref,omitempty" bson:"nextHref,omitempty"`
	UUIDHref     string `json:"uuidHref,omitempty" bson:"uuidHref,omitempty"`
}

type Context struct {
	Employee struct {
		Meta Meta `json:"meta"`
	} `json:"employee"`
}

type Group struct {
	Meta `json:"meta"`
}

type Owner struct {
	Meta Meta `json:"meta"`
}

type SalePrices []struct {
	Value    float64 `json:"value"`
	Currency struct {
		Meta Meta `json:"meta"`
	} `json:"currency"`
	PriceType string `json:"priceType"`
}

type BuyPrice struct {
	Value    float64 `json:"value"`
	Currency struct {
		Meta Meta `json:"meta"`
	} `json:"currency"`
}

type Characteristics []struct {
	Meta  Meta   `json:"meta"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Uom struct {
	Meta Meta `json:"meta"`
}

type Stock struct {
	Stock map[string]string `json:"stock"`
}
