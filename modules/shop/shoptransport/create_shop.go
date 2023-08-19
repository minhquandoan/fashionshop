package shoptransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/shop/shopbiz"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/shop/shopstorage"
)

func AddShop(ctx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		shopColl := db.GetCollection(ctx.GetDbClient(), shopmodel.GetCollectionName())
		shopStorage := shopstorage.NewShopStorage(shopColl)
		shopCreateBiz := shopbiz.NewCreateShopBiz(shopStorage)

		//fetch data from request
		var data shopmodel.ShopModel
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		result, err := shopCreateBiz.AddShop(c, &data)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": common.ErrCannotCreateEntity("shop", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
		return
	}
}
