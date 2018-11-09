package moysklad

const API_BASE = "https://online.moysklad.ru/api/remap/1.1/entity/"

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

const BASE_GROUP = "https://online.moysklad.ru/api/remap/1.1/entity/group"

type Group struct {
	Meta `json:"meta"`
}

type Owner struct {
	Meta Meta `json:"meta"`
}
