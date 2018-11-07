package moysklad

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

type Stores struct {
	Context Context `json:"context"`
	Meta    Meta    `json:"meta"`
	Rows    []Group `json:"rows"`
}
