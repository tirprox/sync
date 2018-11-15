package mscodec

type Counterparty struct {
	Meta         Meta   `json:"meta"`
	SyncID       string `json:"syncId"`
	ID           string `json:"id"`
	AccountID    string `json:"accountId"`
	Owner        Owner  `json:"owner"`
	Shared       bool   `json:"shared"`
	Group        Group  `json:"group"`
	Version      int    `json:"version"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	Created      string `json:"created"`
	CompanyType  string `json:"companyType"`
	Phone        string `json:"phone"`
	Accounts     struct {
		Meta Meta `json:"meta"`
	} `json:"accounts"`
	Tags           []string `json:"tags"`
	Contactpersons struct {
		Meta Meta `json:"meta"`
	} `json:"contactpersons"`
	Notes struct {
		Meta Meta `json:"meta"`
	} `json:"notes"`
	State struct {
		Meta Meta `json:"meta"`
	} `json:"state"`
	SalesAmount int `json:"salesAmount"`
}
