package shoptransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/shop/shopbiz"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/shop/shopstorage"
)

func IncreaseLikedCount(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Get Shop ID to increase like
		id := ctx.Param("id")

		shopColl := db.GetCollection(appCtx.GetDbClient(), shopmodel.GetCollectionName())
		shopStorage := shopstorage.NewShopStorage(shopColl)
		shopUpdateBiz := shopbiz.NewUpdateShopBiz(shopStorage)

		// update liked count
		result, err := shopUpdateBiz.IncreaseLikedCount(ctx, &id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
		return
	}
}