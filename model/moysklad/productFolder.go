package moysklad

const BASE_PRODUCTFOLDER = "https://online.moysklad.ru/api/remap/1.1/entity/productfolder"

// MoySklad ProductFolder type
type ProductFolder struct {
	Meta          Meta   `json:"meta"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Owner         Owner  `json:"owner"`
	Shared        bool   `json:"shared"`
	Group         Group  `json:"group"`
	Version       int    `json:"version"`
	Updated       string `json:"updated"`
	Name          string `json:"name"`
	ExternalCode  string `json:"externalCode"`
	Archived      bool   `json:"archived"`
	PathName      string `json:"pathName"`
	Code          string `json:"code,omitempty"`
	ProductFolder struct {
		Meta Meta `json:"meta"`
	} `json:"productFolder,omitempty"`
}

// MoySklad ProductFolder List type
type ProductFolders struct {
	Context Context         `json:"context"`
	Meta    Meta            `json:"meta"`
	Rows    []ProductFolder `json:"rows"`
}

type ProductFolderContainer struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	ProductFolder ProductFolder `json:"productFolder"`
	Other         struct {
		Products []Product `json:"products"`
	} `json:"other"`
}
