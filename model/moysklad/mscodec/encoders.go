package mscodec

import "encoding/json"

func ToProduct(data interface{}) (product Product) {

	j, err := json.Marshal(data)
	if err != nil {
	}
	json.Unmarshal(j, &product)

	return product
}

func Encode(data interface{}, entity string) (encoded interface{}) {

	return
}
