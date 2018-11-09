package moysklad

import (
	"github.com/mitchellh/mapstructure"
)

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
	Products      []Product     `json:"products"`
	Other         struct {
	} `json:"other"`
}

func GetProductFolders() []ProductFolderContainer {

	data := GetAll(BASE_PRODUCTFOLDER)

	productFolderContainers := []ProductFolderContainer{}

	for _, item := range data {

		folder := ProductFolder{}
		container := ProductFolderContainer{}

		mapstructure.Decode(item, &folder)
		container.ProductFolder = folder
		container.Name = folder.Name
		container.ID = folder.ID

		productFolderContainers = append(productFolderContainers, container)

	}

	return productFolderContainers
}

func (s ProductFolderContainer) GetName() string {
	return s.Name
}
