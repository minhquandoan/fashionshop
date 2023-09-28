package subscriber

import (

	"github.com/minhquandoan/fashionshop/component"
)

func SetUp(appCtx component.AppContext) {
	IncreaseLikedCountForShop(appCtx)
}