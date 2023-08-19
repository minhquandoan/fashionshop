package shoptransport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/shop/shopbiz"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/shop/shopstorage"
)

func ListShopById(ctx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		shopColl := db.GetCollection(ctx.GetDbClient(), shopmodel.GetCollectionName())
		shopStorage := shopstorage.NewShopStorage(shopColl)
		shopCreateBiz := shopbiz.NewListShopBiz(shopStorage)

		id := c.Param("id")
		fmt.Println(id)
		result, err := shopCreateBiz.ListShopById(c, &id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": *result,
		})
		return
	}
}
