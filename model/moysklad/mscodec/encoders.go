package mscodec

import (
	"github.com/jinzhu/copier"
)

func ToProduct(data interface{}) (product Product) {

	copier.Copy(&product, &data)

	return product
}
