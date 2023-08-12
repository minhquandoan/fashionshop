package producttransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	productbiz "github.com/minhquandoan/fashionshop/modules/product/biz"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	productstorage "github.com/minhquandoan/fashionshop/modules/product/storage"
	"github.com/quan-doan/golang-mongo-example/common"
)

func CreateOneProduct(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		dbClient := appCtx.GetDbClient()
		productColl := db.GetCollection(dbClient, collectionName)
		productStore := productstorage.NewDbStore(productColl)
		productBiz := productbiz.NewCreateProductBiz(productStore)
		

		var data productmodel.Product
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": common.ErrInvalidRequest(err),
			})
			return
		}

		id, err := productBiz.CreateProduct(ctx, &data)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": id,
		})
		return
	}
}