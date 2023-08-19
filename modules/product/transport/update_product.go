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

type ProductUpdate struct {
	Data bson.M `json:"data"`
}

func UpdateOneProduct(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		dbClient := appCtx.GetDbClient()
		productColl := db.GetCollection(dbClient, productmodel.GetCollectionName())

		productStore := productstorage.NewDbStore(productColl)
		productBiz := productbiz.NewUpdateProductBiz(productStore)
	
		// Get Id from URI
		id := ctx.Param("id")
		log.Println("ID: ", id)

		// Fetch updated data
		var data ProductUpdate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		log.Println("data: ", data)

		// Update Data
		result, err := productBiz.UpdateOneProduct(ctx, &id, &data.Data)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H {
				"message": common.NewErrorResponse(err, err.Error(), "", ""),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
		return
	}
}