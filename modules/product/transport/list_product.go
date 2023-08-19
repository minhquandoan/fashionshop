package producttransport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	productbiz "github.com/minhquandoan/fashionshop/modules/product/biz"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	productstorage "github.com/minhquandoan/fashionshop/modules/product/storage"
	"go.mongodb.org/mongo-driver/bson"
)

type RequestBody struct {
	Paging common.Paging `json:"paging"`
	Filter bson.M	`json:"filter"`
}

func ListProduct(appCtx component.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		dbClient := appCtx.GetDbClient()
		productColl := db.GetCollection(dbClient, productmodel.GetCollectionName())

		productStore := productstorage.NewDbStore(productColl)
		productBiz := productbiz.NewListProductBiz(productStore)

		var data []productmodel.Product
		if err := productBiz.ListProducts(ctx, &data); err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "can not retrieve any data",
			})

			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func ListProductsByFilters(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		dbClient := appCtx.GetDbClient()
		productColl := db.GetCollection(dbClient, productmodel.GetCollectionName())

		productStore := productstorage.NewDbStore(productColl)
		productBiz := productbiz.NewListProductBiz(productStore)

		// Fetch filter from request
		var requestBody RequestBody
		if err := ctx.ShouldBind(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": common.ErrInvalidRequest(err),
			})
			return
		}

		log.Println("Filter: ", requestBody)

		// Fullfill the paging 
		requestBody.Paging.Fullfill()

		var data []productmodel.Product
		if err := productBiz.ListProductByFilters(ctx, &requestBody.Filter, &requestBody.Paging, &data); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": common.ErrCannotListEntity("Product", err),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
		return
	}
}
