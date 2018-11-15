package mscodec

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
