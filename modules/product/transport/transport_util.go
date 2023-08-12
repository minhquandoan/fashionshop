package producttransport

import (
	"github.com/minhquandoan/fashionshop/component"
	productbiz "github.com/minhquandoan/fashionshop/modules/product/biz"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	productstorage "github.com/minhquandoan/fashionshop/modules/product/storage"
	"github.com/quan-doan/golang-mongo-example/db"
)

var collectionName string = productmodel.GetCollectionName()

func GetProductBiz(ctx component.AppContext) interface{} {
	dbClient := ctx.GetDbClient()
	productColl := db.GetCollection(dbClient, collectionName)
	productStore := productstorage.NewDbStore(productColl)

	return productbiz.NewCreateProductBiz(productStore)
}