package producttransport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	productbiz "github.com/minhquandoan/fashionshop/modules/product/biz"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	productstorage "github.com/minhquandoan/fashionshop/modules/product/storage"
)

func ListProduct(appCtx component.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		dbClient := appCtx.GetDbClient()
		productColl := db.GetCollection(dbClient, productmodel.GetCollectionName())

		productStore := productstorage.NewDbStore(productColl)
		productBiz := productbiz.NewListProductBiz(productStore)

		var data []productmodel.Product
		if err := productBiz.ListProducts(ctx, &data); err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusNotFound, gin.H {
				"error": "can not retrieve any data",
			})

			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}