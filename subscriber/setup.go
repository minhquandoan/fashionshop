package subscriber

import (
	"context"

	"github.com/minhquandoan/fashionshop/component"
)

func SetUp(appCtx component.AppContext) {
	IncreaseLikedCountForShop(appCtx, context.Background())
}