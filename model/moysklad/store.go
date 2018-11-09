package moysklad

import (
	"github.com/mitchellh/mapstructure"
)

const BASE_STORE = "https://online.moysklad.ru/api/remap/1.1/entity/store"

var ALLOWED_STORES = []string{"Флигель new", "Арма Мск"}

type Store struct {
	Meta      Meta   `json:"meta"`
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Owner     struct {
		Meta Meta `json:"meta"`
	} `json:"owner"`
	Shared bool `json:"shared"`
	Group  struct {
		Meta `json:"meta"`
	} `json:"group"`
	Version      int    `json:"version"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	PathName     string `json:"pathName"`
	Address      string `json:"address"`
}

func GetStores() []Store {

	data := GetAll(BASE_STORE)

	filterable := []Filterable{}

	for _, item := range data {

		store := Store{}
		mapstructure.Decode(item, &store)
		filterable = append(filterable, store)
	}

	filtered := FilterSlice(filterable, ALLOWED_STORES)

	storeSlice := []Store{}
	for _, item := range filtered {
		storeSlice = append(storeSlice, item.(Store))
	}

	return storeSlice
}

func (s Store) GetName() string {
	return s.Name
}
